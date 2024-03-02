package tx

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/math"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryTxsByEvents performs a search for transactions for a given set of events
// via the Tendermint RPC. An event takes the form of:
// "{eventAttribute}.{attributeKey} = '{attributeValue}'". Each event is
// concatenated with an 'AND' operand. It returns a slice of Info object
// containing txs and metadata. An error is returned if the query fails.
// If an empty string is provided it will order txs by asc
func QueryTxsByEvents(clientCtx client.Context, events []string, page, limit int, orderBy string) (*sdk.SearchTxsResult, error) {
	if len(events) == 0 {
		return nil, errors.New("must declare at least one event to search")
	}

	if page <= 0 {
		return nil, errors.New("page must be greater than 0")
	}

	if limit <= 0 {
		return nil, errors.New("limit must be greater than 0")
	}

	// XXX: implement ANY
	query := strings.Join(events, " AND ")

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	// TODO: this may not always need to be proven
	// https://github.com/cosmos/cosmos-sdk/issues/6807
	resTxs, err := node.TxSearch(context.Background(), query, true, &page, &limit, orderBy)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(clientCtx, resTxs.Txs)
	if err != nil {
		return nil, err
	}

	txs, err := formatTxResults(clientCtx.TxConfig, resTxs.Txs, resBlocks)
	if err != nil {
		return nil, err
	}

	result := sdk.NewSearchTxsResult(uint64(resTxs.TotalCount), uint64(len(txs)), uint64(page), uint64(limit), txs)

	return result, nil
}

// QueryTx queries for a single transaction by a hash string in hex format. An
// error is returned if the transaction does not exist or cannot be queried.
func QueryTx(clientCtx client.Context, hashHexStr string) (*sdk.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return nil, err
	}

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	// TODO: this may not always need to be proven
	// https://github.com/cosmos/cosmos-sdk/issues/6807
	resTx, err := node.Tx(context.Background(), hash, true)
	if err != nil {
		return nil, err
	}

	resBlocks, err := getBlocksForTxResults(clientCtx, []*coretypes.ResultTx{resTx})
	if err != nil {
		return nil, err
	}

	out, err := mkTxResult(clientCtx.TxConfig, resTx, resBlocks[resTx.Height])
	if err != nil {
		if strings.Contains(err.Error(), "Mismatched \"*types.MsgLiquidStake\"") {
			return tryParsingLiquidStakeTx(resTx, resBlocks[resTx.Height])
		}
		return out, err
	}

	return out, nil
}

func tryParsingLiquidStakeTx(resTx *coretypes.ResultTx, block *coretypes.ResultBlock) (*sdk.TxResponse, error) {
	var ok bool
	var denom string
	amountInt := math.ZeroInt()
	for _, event := range resTx.TxResult.Events {
		if event.Type == "coin_spent" {
			for _, attribute := range event.Attributes {
				if attribute.Key == "amount" && strings.Contains(attribute.Value, "ibc/") {
					amountString := attribute.Value
					amountSplit := strings.Split(amountString, "ibc/")
					if len(amountSplit) != 2 {
						return nil, fmt.Errorf("error parsing amount %s", amountString)
					}

					amountInt, ok = sdk.NewIntFromString(amountSplit[0])
					if !ok {
						return nil, fmt.Errorf("error parsing amount %s", amountString)
					}
					denom = "ibc/" + amountSplit[1]
				}
			}
		}
	}
	if amountInt.IsZero() {
		return nil, errors.New("Unable to find liquid stake amount from tx")
	}

	return &sdk.TxResponse{
		Height:    resTx.Height,
		TxHash:    resTx.Hash.String(),
		Codespace: resTx.TxResult.Codespace,
		Code:      resTx.TxResult.Code,
		Data:      strings.ToUpper(hex.EncodeToString(resTx.TxResult.Data)),
		RawLog:    resTx.TxResult.Log,
		Info:      resTx.TxResult.Info,
		GasWanted: resTx.TxResult.GasWanted,
		GasUsed:   resTx.TxResult.GasUsed,
		Timestamp: block.Block.Time.String(),
		Events:    resTx.TxResult.Events,
		Tx:        tx,
	}, nil
}

// formatTxResults parses the indexed txs into a slice of TxResponse objects.
func formatTxResults(txConfig client.TxConfig, resTxs []*coretypes.ResultTx, resBlocks map[int64]*coretypes.ResultBlock) ([]*sdk.TxResponse, error) {
	var err error
	out := make([]*sdk.TxResponse, len(resTxs))
	for i := range resTxs {
		out[i], err = mkTxResult(txConfig, resTxs[i], resBlocks[resTxs[i].Height])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

func getBlocksForTxResults(clientCtx client.Context, resTxs []*coretypes.ResultTx) (map[int64]*coretypes.ResultBlock, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*coretypes.ResultBlock)

	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(context.Background(), &resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func mkTxResult(txConfig client.TxConfig, resTx *coretypes.ResultTx, resBlock *coretypes.ResultBlock) (*sdk.TxResponse, error) {
	txb, err := txConfig.TxDecoder()(resTx.Tx)
	if err != nil {
		return nil, err
	}
	p, ok := txb.(intoAny)
	if !ok {
		return nil, fmt.Errorf("expecting a type implementing intoAny, got: %T", txb)
	}
	any := p.AsAny()
	return sdk.NewResponseResultTx(resTx, any, resBlock.Block.Time.Format(time.RFC3339)), nil
}

// Deprecated: this interface is used only internally for scenario we are
// deprecating (StdTxConfig support)
type intoAny interface {
	AsAny() *codectypes.Any
}

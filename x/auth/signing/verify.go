package signing

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

// VerifySignature verifies a transaction signature contained in SignatureData abstracting over different signing modes
// and single vs multi-signatures.
func VerifySignature(pubKey cryptotypes.PubKey, signerData SignerData, sigData signing.SignatureData, handler SignModeHandler, tx sdk.Tx) error {
	switch data := sigData.(type) {
	case *signing.SingleSignatureData:
		signBytes, err := handler.GetSignBytes(data.SignMode, signerData, tx)
		if err != nil {
			return err
		}
		if !pubKey.VerifySignature(signBytes, data.Signature) {
			directSignBytes, err := handler.GetSignBytes(signing.SignMode_SIGN_MODE_DIRECT, signerData, tx)
			if err != nil {
				return errorsmod.Wrapf(err, "unable to conver tx to directSignBytes")
			}
			return fmt.Errorf(
				"unable to verify single signer signature '%s' for signBytes '%s' from tx '%s'",
				hex.EncodeToString(data.Signature),
				hex.EncodeToString(signBytes),
				base64.StdEncoding.EncodeToString(directSignBytes))
		}
		return nil

	case *signing.MultiSignatureData:
		multiPK, ok := pubKey.(multisig.PubKey)
		if !ok {
			return fmt.Errorf("expected %T, got %T", (multisig.PubKey)(nil), pubKey)
		}
		err := multiPK.VerifyMultisignature(func(mode signing.SignMode) ([]byte, error) {
			return handler.GetSignBytes(mode, signerData, tx)
		}, data)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("unexpected SignatureData %T", sigData)
	}
}

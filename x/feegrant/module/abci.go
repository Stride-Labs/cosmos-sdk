package module

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	fmt.Println("FEE GRANT END BLOCKER")

	k.RemoveExpiredAllowances(ctx)
}

package authz

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz/keeper"
)

// BeginBlocker is called at the beginning of every block
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {
	fmt.Println("AUTHZ BEGIN BLOCKER")

	// delete all the mature grants
	if err := keeper.DequeueAndDeleteExpiredGrants(ctx); err != nil {
		panic(err)
	}
}

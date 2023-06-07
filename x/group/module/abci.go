package module

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group/keeper"
)

// EndBlocker called at every block, updates proposal's `FinalTallyResult` and
// prunes expired proposals.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	fmt.Println("GROUP END BLOCKER")

	if err := k.TallyProposalsAtVPEnd(ctx); err != nil {
		panic(err)
	}

	if err := k.PruneProposals(ctx); err != nil {
		panic(err)
	}
}

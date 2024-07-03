package pasans_test

import (
	"testing"

	keepertest "pasans/testutil/keeper"
	"pasans/testutil/nullify"
	pasans "pasans/x/pasans/module"
	"pasans/x/pasans/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PasansKeeper(t)
	pasans.InitGenesis(ctx, k, genesisState)
	got := pasans.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

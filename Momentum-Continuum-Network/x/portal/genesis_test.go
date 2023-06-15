package portal_test

import (
	"testing"

	keepertest "Momentum-Continuum-Network/testutil/keeper"
	"Momentum-Continuum-Network/testutil/nullify"
	"Momentum-Continuum-Network/x/portal"
	"Momentum-Continuum-Network/x/portal/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PortalKeeper(t)
	portal.InitGenesis(ctx, *k, genesisState)
	got := portal.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

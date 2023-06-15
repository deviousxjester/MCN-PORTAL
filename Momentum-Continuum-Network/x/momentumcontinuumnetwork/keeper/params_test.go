package keeper_test

import (
	"testing"

	testkeeper "Momentum-Continuum-Network/testutil/keeper"
	"Momentum-Continuum-Network/x/momentumcontinuumnetwork/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MomentumcontinuumnetworkKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

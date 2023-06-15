package keeper_test

import (
	"context"
	"testing"

	keepertest "Momentum-Continuum-Network/testutil/keeper"
	"Momentum-Continuum-Network/x/portal/keeper"
	"Momentum-Continuum-Network/x/portal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PortalKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}

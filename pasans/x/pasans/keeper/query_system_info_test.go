package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "pasans/testutil/keeper"
	"pasans/testutil/network"
	"pasans/testutil/nullify"
	"pasans/x/pasans/types"
)

func networkWithSystemInfoObjects(t *testing.T) (*network.Network, types.SystemInfo) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	systemInfo := state.SystemInfo
	nullify.Fill(&systemInfo)
	state.SystemInfo = systemInfo
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.SystemInfo
}

func TestSystemInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.PasansKeeper(t)
	item := createTestSystemInfo(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetSystemInfoRequest
		response *types.QueryGetSystemInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSystemInfoRequest{},
			response: &types.QueryGetSystemInfoResponse{SystemInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SystemInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

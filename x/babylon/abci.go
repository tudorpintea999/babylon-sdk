package babylon

import (
	"time"

	"github.com/babylonchain/babylon-sdk/x/babylon/keeper"
	"github.com/babylonchain/babylon-sdk/x/babylon/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// EndBlocker is called after every block
func EndBlocker(ctx sdk.Context, k *keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
}

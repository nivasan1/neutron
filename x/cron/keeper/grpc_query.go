package keeper

import (
	"github.com/neutron-org/neutron/v2/x/cron/types"
)

var _ types.QueryServer = Keeper{}

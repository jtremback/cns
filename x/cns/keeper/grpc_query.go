package keeper

import (
	"github.com/jtremback/cns/x/cns/types"
)

var _ types.QueryServer = Keeper{}

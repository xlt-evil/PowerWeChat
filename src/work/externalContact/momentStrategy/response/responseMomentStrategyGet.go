package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"strategy"`
}

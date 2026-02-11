package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyGetRange struct {
	response.ResponseWork

	Range []*power.HashMap `json:"range"`
}

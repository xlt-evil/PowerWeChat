package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	response.ResponseWork
	Group []*power.HashMap `json:"group"`
}

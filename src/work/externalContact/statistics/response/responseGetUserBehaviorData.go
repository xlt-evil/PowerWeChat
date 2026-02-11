package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}

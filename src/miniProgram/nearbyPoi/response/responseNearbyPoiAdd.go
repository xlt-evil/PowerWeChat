package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseNearbyPoiAdd struct {
	response.ResponseMiniProgram
	Data []*power.HashMap `json:"data"`
}

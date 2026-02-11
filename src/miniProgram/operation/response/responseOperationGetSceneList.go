package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetSceneList struct {
	response.ResponseMiniProgram
	Scene []*power.HashMap `json:"scene"`
}

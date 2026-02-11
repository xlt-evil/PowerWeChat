package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	response.ResponseMiniProgram
	GrayReleasePlan *power.HashMap `json:"gray_release_plan"`
}

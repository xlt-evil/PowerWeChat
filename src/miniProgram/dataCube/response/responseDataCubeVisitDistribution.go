package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeVisit struct {
	response.ResponseMiniProgram

	RefDate string           `json:"ref_date"`
	List    []*power.HashMap `json:"list"`
}

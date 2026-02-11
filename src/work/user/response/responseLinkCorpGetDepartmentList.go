package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetDepartmentList struct {
	response.ResponseWork

	DepartmentList []*power.HashMap `json:"department_list"`
}

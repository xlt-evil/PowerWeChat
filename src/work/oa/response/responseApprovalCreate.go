package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalCreate struct {
	response.ResponseWork
	SpNo string `json:"sp_no"`
}

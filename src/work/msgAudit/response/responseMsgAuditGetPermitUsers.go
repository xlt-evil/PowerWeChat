package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetPermitUsers struct {
	response.ResponseWork
	IDs []string `json:"ids"`
}

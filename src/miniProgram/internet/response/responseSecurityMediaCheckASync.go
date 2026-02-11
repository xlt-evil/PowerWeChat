package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseSecurityMediaCheckASync struct {
	response.ResponseMiniProgram
	TraceID string `json:"trace_id"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageAddTemplate struct {
	response.ResponseMiniProgram
	PriTmplID string `json:"priTmplId"`
}

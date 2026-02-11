package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerServiceMessageGetTempMedia struct {
	response.ResponseMiniProgram
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}

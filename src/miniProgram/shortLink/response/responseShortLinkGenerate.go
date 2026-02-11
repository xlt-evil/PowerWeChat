package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseShortLinkGenerate struct {
	response.ResponseMiniProgram

	Link string `json:"link"`
}

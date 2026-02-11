package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}

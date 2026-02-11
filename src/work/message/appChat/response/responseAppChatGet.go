package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatGet struct {
	response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}

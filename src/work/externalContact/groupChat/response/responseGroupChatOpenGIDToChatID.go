package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}

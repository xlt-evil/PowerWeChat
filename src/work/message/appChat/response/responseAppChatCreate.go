package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatCreate struct {
	response.ResponseWork

	ChatID string `json:"chatid"`
}

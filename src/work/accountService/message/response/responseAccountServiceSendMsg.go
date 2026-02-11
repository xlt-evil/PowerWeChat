package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsg struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseJoinCode struct {
	response.ResponseWork

	JoinCode string `json:"join_qrcode"`
}

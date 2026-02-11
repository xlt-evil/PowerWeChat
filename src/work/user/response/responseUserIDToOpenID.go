package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWork

	OpenID string `json:"openid"`
}

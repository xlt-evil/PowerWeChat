package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingCreate struct {
	response.ResponseWork

	LivingID int `json:"livingid"`
}

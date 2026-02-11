package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseTransferCustomer struct {
	response.ResponseWork

	Customer []*power.HashMap `json:"customer"`
}

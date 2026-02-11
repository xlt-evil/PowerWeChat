package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseUnitfy struct {
	response.ResponsePayment

	PrepayID string `json:"prepay_id"`
}

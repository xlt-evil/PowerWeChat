package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	response.ResponsePayment

	Status string `header:"status"`
}

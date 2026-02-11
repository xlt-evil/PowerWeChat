package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentSendResult struct {
	response.ResponseWork

	CustomerList []*power.HashMap `json:"customer_list"`
	NextCursor   string           `json:"next_cursor"`
}

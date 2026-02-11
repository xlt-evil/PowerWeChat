package response

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"

type ResponseIsTradeManaged struct {
	response.ResponseMiniProgram
	IsTradeManaged bool `json:"is_trade_managed"`
}

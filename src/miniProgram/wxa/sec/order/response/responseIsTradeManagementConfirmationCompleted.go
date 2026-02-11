package response

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"

type ResponseIsTradeManagementConfirmationCompleted struct {
	response.ResponseMiniProgram
	Completed bool `json:"completed"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountList struct {
	response.ResponseWork

	AccountList []*power.HashMap `json:"account_list"`
}

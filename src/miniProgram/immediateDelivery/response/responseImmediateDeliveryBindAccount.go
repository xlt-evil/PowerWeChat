package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseImmediateDeliveryBindAccount struct {
	response.ResponseMiniProgram

	ShopList []*power.HashMap `json:"shop_list"`
}

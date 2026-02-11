package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	response.ResponseMiniProgram

	URL int `json:"url"`
}

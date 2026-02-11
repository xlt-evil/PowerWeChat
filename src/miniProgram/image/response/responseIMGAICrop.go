package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGAICrop struct {
	response.ResponseMiniProgram
	Results []*power.HashMap `json:"results"`
}

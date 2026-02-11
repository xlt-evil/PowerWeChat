package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	response.ResponseWork

	URL string `json:"url"`
}

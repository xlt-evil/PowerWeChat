package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetAPIDomainIP struct {
	IPList []string `json:"ip_list"`
	response.ResponseWork
}

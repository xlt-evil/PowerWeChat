package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMenuCreate struct {
	response.ResponseOfficialAccount
}

type ResponseMenuCreateConditional struct {
	response.ResponseOfficialAccount

	MenuID string `json:"menuid"`
}

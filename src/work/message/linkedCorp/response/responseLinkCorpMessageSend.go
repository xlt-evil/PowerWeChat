package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpMessageSend struct {
	response.ResponseWork

	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []string `json:"invalidparty"`
	InvalidTag   []string `json:"invalidtag"`
}

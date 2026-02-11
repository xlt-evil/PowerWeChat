package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpGroupTransferSession struct {
	response.ResponseWork
	Userid     string `json:"userid"`
	SessionKey string `json:"session_key"`
}

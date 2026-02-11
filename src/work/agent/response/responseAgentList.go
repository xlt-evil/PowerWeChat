package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseAgentList struct {
	response.ResponseWork
	AgentList []ResponseAgentGet `json:"agentlist"`
}

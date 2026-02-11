package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagGetStrategyTagList struct {
	response.ResponseWork

	TagGroups []*StrategyTagGroup `json:"tag_group"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseResignedGetUnassignedList struct {
	response.ResponseWork

	Info       []*power.HashMap `json:"info"`
	IsLast     bool             `json:"is_last"`
	NextCursor string           `json:"next_cursor"`
}

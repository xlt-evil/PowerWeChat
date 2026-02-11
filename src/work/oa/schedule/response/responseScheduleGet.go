package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseScheduleGet struct {
	response.ResponseWork

	ScheduleList []*power.HashMap `json:"schedule_list"`
}

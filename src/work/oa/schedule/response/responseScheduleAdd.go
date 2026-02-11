package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseScheduleAdd struct {
	response.ResponseWork

	ScheduleID string `json:"schedule_id"`
}

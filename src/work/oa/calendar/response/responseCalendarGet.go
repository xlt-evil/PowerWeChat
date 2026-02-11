package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCalendarGet struct {
	response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`
}

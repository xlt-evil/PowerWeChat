package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingRoomList struct {
	response.ResponseWork

	MeetingRoomList []*power.HashMap `json:"meetingroom_list"`
}

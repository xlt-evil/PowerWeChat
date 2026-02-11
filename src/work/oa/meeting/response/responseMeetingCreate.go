package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMeetingCreate struct {
	response.ResponseWork

	MeetingID int `json:"meetingid"`
}

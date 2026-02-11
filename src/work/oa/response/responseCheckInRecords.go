package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRecords struct {
	response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}

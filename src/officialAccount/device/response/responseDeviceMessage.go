package response

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"

type ResponseDeviceMessage struct {
	response.ResponseOfficialAccount

	Ret     int    `json:"ret"`
	RetInfo string `json:"ret_info"`
}

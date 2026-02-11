package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceInfo struct {
	response.ResponseWork

	SpaceInfo *power.HashMap `json:"space_info"`
}

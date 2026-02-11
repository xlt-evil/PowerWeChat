package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceShare struct {
	response.ResponseWork

	SpaceShareURL string `json:"space_share_url"`
}

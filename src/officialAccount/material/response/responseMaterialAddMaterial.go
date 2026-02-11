package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseMaterialAddMaterial struct {
	response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}

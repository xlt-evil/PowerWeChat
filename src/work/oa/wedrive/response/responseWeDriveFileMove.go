package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileMove struct {
	response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}

package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserExportJobs struct {
	response.ResponseWork

	JobID string `json:"jobid"`
}

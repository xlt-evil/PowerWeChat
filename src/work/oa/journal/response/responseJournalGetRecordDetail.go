package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseJournalGetRecordDetail struct {
	response.ResponseWork

	Info *power.HashMap `json:"info"`
}

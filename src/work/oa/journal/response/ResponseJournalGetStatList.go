package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseJournalGetStatList struct {
	response.ResponseWork

	StatList *power.HashMap `json:"stat_list"`
}

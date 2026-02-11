package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponsePluginGetPluginList struct {
	response.ResponseMiniProgram
	PluginList []*power.HashMap `json:"plugin_list"`
}

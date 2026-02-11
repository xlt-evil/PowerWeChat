package request

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}

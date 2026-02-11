package menu

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(app)

}

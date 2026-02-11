package menu

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, error) {

	client, err := NewClient(app)

	return client, err

}

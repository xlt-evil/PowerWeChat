package security

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, err
	}

	return client, nil

}

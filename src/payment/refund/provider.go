package refund

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/payment/kernel"
)

func RegisterProvider(app kernel.ApplicationPaymentInterface) (*Client, error) {

	return NewClient(app)

}

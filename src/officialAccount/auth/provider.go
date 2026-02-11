package auth

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	return NewAccessToken(app)

}

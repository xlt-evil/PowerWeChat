package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
	"github.com/xlt-evil/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseGetAppLicenseInfo struct {
	response.ResponseWork
	model.LicenseInfo
}

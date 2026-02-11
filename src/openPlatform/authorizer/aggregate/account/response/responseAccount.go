package response

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"

type ResponseCreate struct {
	response.ResponseOpenPlatform

	OpenAppID string `json:"open_appid"`
}

type ResponseGetBinding struct {
	response.ResponseOpenPlatform

	OpenAppid string `json:"open_appid"`
}

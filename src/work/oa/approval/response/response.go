package response

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"

type ResponseCreateTemplate struct {
	response.ResponseWork

	TemplateId string `json:"template_id"`
}

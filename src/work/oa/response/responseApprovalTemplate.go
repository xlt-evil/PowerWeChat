package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type Title struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Placeholder struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Property struct {
	Control     string         `json:"control"`
	ID          string         `json:"id"`
	Title       []*Title       `json:"title"`
	Placeholder []*Placeholder `json:"placeholder"`
	Require     int            `json:"require"`
	UnPrint     int            `json:"un_print"`
}

type Selector struct {
	Type    string    `json:"type"`
	Options []*Option `json:"options"`
}

type Value struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type Option struct {
	Key   string   `json:"key"`
	Value []*Value `json:"value"`
}

type DateConfig struct {
	Type string `json:"type"` // day / hour
}

type ContactConfig struct {
	Type string `json:"type"` // single / multi
	Mode string `json:"mode"` // user / department
}

type TableConfig struct {
	Children   []Control     `json:"children"`
	StatFields []interface{} `json:"stat_field"`
}

type Config struct {
	Selector *Selector      `json:"selector"`
	Date     *DateConfig    `json:"date"`
	Contact  *ContactConfig `json:"contact"`
	Table    *TableConfig   `json:"table"`
}

type Control struct {
	Property *Property `json:"property"`
	Config   *Config   `json:"config"`
}

type TemplateContent struct {
	Controls []Control `json:"controls"`
}

type TemplateName struct {
	Text string `json:"text"`
	Lang string `json:"lang"`
}

type ResponseApprovalTemplate struct {
	response.ResponseWork

	TemplateNames   []*TemplateName  `json:"template_names"`
	TemplateContent *TemplateContent `json:"template_content"`
}

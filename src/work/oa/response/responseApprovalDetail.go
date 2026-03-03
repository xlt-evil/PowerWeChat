package response

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalDetail struct {
	response.ResponseWork
	Info ResponseApprovalDetailInfo `json:"info"`
}

// Info 审批单核心信息
type ResponseApprovalDetailInfo struct {
	SpNo        string      `json:"sp_no"`        // 审批单号
	SpName      string      `json:"sp_name"`      // 审批单名称
	SpStatus    int         `json:"sp_status"`    // 审批单状态
	TemplateID  string      `json:"template_id"`  // 模板ID
	ApplyTime   int64       `json:"apply_time"`   // 提交时间（时间戳）
	Applyer     Applyer     `json:"applyer"`      // 提交人信息
	SpRecord    []SpRecord  `json:"sp_record"`    // 审批记录
	Notifyer    []Notifyer  `json:"notifyer"`     // 抄送人列表
	ApplyData   ApplyData   `json:"apply_data"`   // 审批单填写数据
	Comments    []Comment   `json:"comments"`     // 备注信息列表
	ProcessList ProcessList `json:"process_list"` // 审批流程列表
}

// Applyer 提交人信息
type Applyer struct {
	UserID  string `json:"userid"`  // 提交人用户ID
	PartyID string `json:"partyid"` // 提交人所属部门ID
}

// SpRecord 审批记录项
type SpRecord struct {
	SpStatus     int      `json:"sp_status"`    // 审批状态
	ApproverAttr int      `json:"approverattr"` // 审批人属性
	Details      []Detail `json:"details"`      // 审批详情列表
}

// Detail 审批详情
type Detail struct {
	Approver Approver `json:"approver"`  // 审批人信息
	Speech   string   `json:"speech"`    // 审批意见
	SpStatus int      `json:"sp_status"` // 审批状态
	Sptime   int64    `json:"sptime"`    // 审批时间（时间戳）
	MediaID  []string `json:"media_id"`  // 审批附件媒体ID列表
}

// Approver 审批人信息
type Approver struct {
	UserID string `json:"userid"` // 审批人用户ID
}

// Notifyer 抄送人信息
type Notifyer struct {
	UserID string `json:"userid"` // 抄送人用户ID
}

// ApplyData 审批单填写数据
type ApplyData struct {
	Contents []Content `json:"contents"` // 表单内容列表
}

// Content 表单内容项
type Content struct {
	Control string                    `json:"control"` // 控件类型（如Text）
	ID      string                    `json:"id"`      // 控件ID
	Title   []ApprovalDetailInfoTitle `json:"title"`   // 控件标题
	Value   ApprovalDetailInfoValue   `json:"value"`   // 控件填写值
}

// ApprovalDetailInfoTitle Title 控件标题
type ApprovalDetailInfoTitle struct {
	Text string `json:"text"` // 标题文本
	Lang string `json:"lang"` // 语言类型（如zh_CN）
}

// ApprovalDetailInfoValue Value 控件填写值
type ApprovalDetailInfoValue struct {
	Text        []string `json:"text"`        // 文本内容
	Tips        []string `json:"tips"`        // 提示信息
	Members     []string `json:"members"`     // 成员列表
	Departments []string `json:"departments"` // 部门列表
	Files       []string `json:"files"`       // 文件列表
	Children    []string `json:"children"`    // 子控件列表
	StatField   []string `json:"stat_field"`  // 统计字段
}

// Comment 备注信息
type Comment struct {
	CommentUserInfo CommentUserInfo `json:"commentUserInfo"` // 备注人信息
	Commenttime     int64           `json:"commenttime"`     // 备注时间（时间戳）
	Commentcontent  string          `json:"commentcontent"`  // 备注内容
	Commentid       string          `json:"commentid"`       // 备注ID
	MediaID         []string        `json:"media_id"`        // 备注附件媒体ID列表
}

// CommentUserInfo 备注人信息
type CommentUserInfo struct {
	UserID string `json:"userid"` // 备注人用户ID
}

// ProcessList 审批流程列表
type ProcessList struct {
	NodeList []Node `json:"node_list"` // 审批节点列表
}

// Node 审批节点
type Node struct {
	NodeType    int       `json:"node_type"`     // 节点类型
	SpStatus    int       `json:"sp_status"`     // 节点审批状态
	ApvRel      int       `json:"apv_rel"`       // 审批关系（1-或签，2-会签等）
	SubNodeList []SubNode `json:"sub_node_list"` // 子节点列表（审批人）
}

// SubNode 子审批节点（具体审批人）
type SubNode struct {
	UserID   string   `json:"userid"`    // 审批人用户ID
	Speech   string   `json:"speech"`    // 审批意见
	SpYj     int      `json:"sp_yj"`     // 审批意见类型（1-同意等）
	Sptime   int64    `json:"sptime"`    // 审批时间（时间戳）
	MediaIDs []string `json:"media_ids"` // 审批附件媒体ID列表
}

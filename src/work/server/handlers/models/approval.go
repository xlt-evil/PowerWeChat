package models

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/contract"
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/models"
)

type EventSysApprovalChange struct {
	contract.EventInterface
	models.CallbackMessageHeader
	ApprovalInfo ApprovalInfo `xml:"ApprovalInfo"` // 审批信息主体
}

package request

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveFileACLDel struct {
	UserID   string           `json:"userid"`
	FileID   string           `json:"fileid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}

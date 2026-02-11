package request

import "github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"

type RequestWeDriveSpaceACLAdd struct {
	UserID   string           `json:"userid"`
	SpaceID  string           `json:"spaceid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}

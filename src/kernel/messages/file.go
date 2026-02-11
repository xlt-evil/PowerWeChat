package messages

import (
	"github.com/xlt-evil/PowerWeChat/v3/src/kernel/power"
)

type File struct {
	*Media
}

func NewFile(mediaID string, items *power.HashMap) *File {
	m := &File{
		NewMedia(mediaID, "file", items),
	}

	return m
}

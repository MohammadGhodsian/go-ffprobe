// +build !linux,!windows

package ffprobe

import (
	"syscall"
)

func procAttributes() *syscall.SysProcAttr {
	return nil
}
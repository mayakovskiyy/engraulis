package sysmon

import (
	"runtime"
)

func GetCurrMem() int {
	switch runtime.GOOS {
		case "darwin":
			return MtrCurrent()
		case "linux":
			return MtrCurrent()
		default:
			return 0
	}
}
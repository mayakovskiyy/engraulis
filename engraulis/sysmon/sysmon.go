package sysmon

import (
	"time"
	"fmt"
)

func Monitoring(duration time.Duration) string {
	inp := MtrDarwin()

	output := fmt.Sprintf("RAM Usage: %dMB", inp)
	
	return output
}
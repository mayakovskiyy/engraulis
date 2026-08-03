package sysmon

import (
	"os"
	"fmt"
	"time"
)

func SysInfo(duration time.Duration) string {
	hn, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}

	

	output := fmt.Sprintf("Hostname: %s", hn)
	
	return output
}
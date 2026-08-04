package sysmon

import (
	"os/exec"
	"fmt"
	"strconv"
	"strings"
)

func MtrCurrent() int {
	output, err := exec.Command("cat", "/proc/meminfo").Output()
	if err != nil {
		fmt.Println(err)
	}

	rawOutput := strings.TrimSpace(string(output))

	var memTotal, memFree int

	lines := strings.Split(string(rawOutput), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		clearPart := strings.Split(parts[1], "kB")
		if len(clearPart) < 2 {
			continue
		}
		valStr := strings.TrimSpace(clearPart[0])
		val, _ := strconv.Atoi(valStr)

		switch {
			case strings.Contains(parts[0], "MemTotal"):
				memTotal = val
			case strings.Contains(parts[0], "MemFree"):
				memFree = val
			default:
				continue
		}
	}
	
	ramUsage := (memTotal - memFree) / 1024

	fmt.Println(ramUsage)
	return ramUsage
}
package sysmon

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func MtrCurrent() int {
	output, err := exec.Command("memory_pressure").Output()
	if err != nil {
		fmt.Println(err)
	}

	rawOutput := strings.TrimSpace(string(output))

	var pagesActive, pagesWDown, pagesComp int

	lines := strings.Split(string(rawOutput), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		valStr := strings.TrimSpace(parts[1])
		val, _ := strconv.Atoi(valStr)

		switch {
		case strings.Contains(parts[0], "Pages active"):
			pagesActive = val
		case strings.Contains(parts[0], "Pages wired down"):
			pagesWDown = val
		case strings.Contains(parts[0], "Pages active"):
			pagesComp = val
		default:
			continue
		}
	}

	ramUsage := (pagesActive + pagesWDown + pagesComp) / 64

	return ramUsage
}

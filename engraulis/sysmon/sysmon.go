package sysmon

import (
	"fmt"
	"math"
	"time"
)

func Monitoring(delay int, samplingRate int, logging bool) string {
	var current int
	var max int = 0
	var min int = math.MaxInt
	var delta int

	ticker := time.NewTicker(time.Duration(samplingRate) * time.Second) // sampling rate ticker
	defer ticker.Stop()

	tickerDelay := time.NewTicker(time.Duration(delay) * time.Minute) // delay ticker
	defer tickerDelay.Stop()

	for {
		select {
		case <-ticker.C:
			current = GetCurrMem()
			if current > max {
				max = current
			}
			if current < min {
				min = current
			}

		case <-tickerDelay.C:
			delta = max - min
			fmt.Printf("Delta: %dMB\nMin: %dMB, Max: %dMB\n", delta, min, max)
			if logging {
				SaveData(min, max, delta)
			}
			max = 0
			min = math.MaxInt
			delta = 0
		}
	}
}

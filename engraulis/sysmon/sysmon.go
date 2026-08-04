package sysmon

import (
	"fmt"
	"math"
	"time"
)

func Monitoring(delay time.Duration, samplingRate time.Duration) string {
	var current int
	var max int = 0
	var min int = math.MaxInt
	var delta int

	ticker := time.NewTicker(samplingRate * time.Second) // sampling rate ticker
	defer ticker.Stop()

	tickerDelay := time.NewTicker(delay * time.Minute) // delay ticker
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
			SaveData(min, max, delta)
			max = 0
			min = math.MaxInt
			delta = 0
		}
	}
}

package sysmon

import (
	"time"
	"fmt"
	"math"
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
				fmt.Printf("Min: %d, Max: %d\n", min, max)
				
			case <-tickerDelay.C:
				delta = max - min
				fmt.Printf("Delta: %d\nMin: %d, Max: %d\n", delta, min, max)
				max = 0
				min = math.MaxInt
		}
	}
}
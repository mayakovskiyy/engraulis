package main

import (
	// "engraulis/client"
	"engraulis/sysmon"
	"fmt"
)

// usage example
// uncomment the 4th line to use the website monitoring

func main() {
	// website monitoring example ↓
	/* address := "https://example.com/"

	res := client.Req(address, 1, 3, false)
	fmt.Printf("Status: %d \n", res.StatusCode)
	fmt.Printf("Server: %s \n", address) // also you're able to use: res.Header.Get("Server"), but address := ... works better imo
	fmt.Printf("Date: %s \n", res.Header.Get("Date")) */

	// system monitoring example ↓
	sysmon.InitDatabase("./db/engraulis.db")
	sysmt := sysmon.Monitoring(1, 3) // the first value stands for minutes between deltas (in minutes). the second value stands for sampling rate delay (in seconds)
	fmt.Println(sysmt)
}

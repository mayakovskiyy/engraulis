package main

import (
	"engraulis/client"
	"engraulis/sysmon"
	"fmt"
)

// usage example

func main() {
	address := "https://example.com/"

	res := client.Req(address, 1, 3, false)
	fmt.Printf("Status: %d \n", res.StatusCode)
	fmt.Printf("Server: %s \n", address) // also you're able to use: res.Header.Get("Server"), but address := ... works better imo
	fmt.Printf("Date: %s \n", res.Header.Get("Date"))

	sysmonitor := sysmon.SysInfo(3)
	fmt.Println(sysmonitor)
}

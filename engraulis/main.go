package main

import (
	"engraulis/client"
	"fmt"
)

// usage example

func main() {
	address := "https://microsoft.com/"
	
	res := client.Req(address, 1, 3)
	fmt.Printf("Status: %d \n", res.StatusCode)
	fmt.Printf("Server: %s \n", address) // also you're able to use: res.Header.Get("Server"), but address := ... works better imo
	fmt.Printf("Date: %s \n", res.Header.Get("Date"))
}

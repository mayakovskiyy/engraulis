package main

import (
	"engraulis/client"
	"fmt"
)

func main() {
	res := client.Req("https://github.com/", 1, 3)
	fmt.Printf("Status: %d \n", res.StatusCode)
	fmt.Printf("Server: %s \n", res.Header.Get("Server"))
	fmt.Printf("Date: %s \n", res.Header.Get("Date"))
}

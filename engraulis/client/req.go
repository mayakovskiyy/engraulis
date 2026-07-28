package client

import (
	"fmt"
	"net/http"
	"time"
)

func Req(address string, delay time.Duration, amount int) http.Response {
	var res *http.Response
	var err error
	var oks int
	var errs int
	name := address
	
	for i := 0; i < amount; i++ {
		res, err = http.Get(address)

		fmt.Printf("Heat %d, Server: %s, Status: %d\n", i, name, res.StatusCode)
		if err != nil {
			fmt.Println(err)
			res.Body.Close()
			continue
		}
		if address == "" {
			fmt.Println("engraulis: Address mustn't be empty.")
			break
		}
		if res.StatusCode == 200 {
			oks += 1
		} else {
			errs += 1
		}
		
		time.Sleep(delay)
	}
	res.Body.Close()
	if res != nil {
		return *res
	}
	return http.Response{}
}

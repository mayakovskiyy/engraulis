package client

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Req(address string, delay time.Duration, amount int, logging bool) http.Response {
	var res *http.Response
	var oks, errs int
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	name := address
	filename := filepath.Join(homedir, "Documents", fmt.Sprintf("log_%s.txt", time.Now().Format("2006-01-02_15-04-05")))

	if delay < 1 {
		fmt.Println("engraulis: delay must equal 1 or be bigger.")
		res.Body.Close()
	}

	for i := 0; i < amount; i++ {
		res, err = http.Get(address)
		statusIcon := ""
		status := ""

		if res.StatusCode == 200 {
			statusIcon = "🟢"
			status = "(OK)"
		} else {
			statusIcon = "🔴"
			status = "(ERROR)"
		}
		fmt.Printf("%s Heat %d, Server: %s, Status: %d %s\n", statusIcon, i, name, res.StatusCode, status)
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

		time.Sleep(delay * time.Second)
	}

	logtext := fmt.Sprintf("Heat finished!\nOKs: %d, Errors: %d\nServer: %s\nDate: %s", oks, errs, address, time.Now())
	fmt.Printf("Heat finished!\nOKs: %d, Errors: %d\n", oks, errs)

	if logging {
		err := os.WriteFile(filename, []byte(logtext), 0755)
		if err != nil {
			fmt.Println(err)
		}
	}

	res.Body.Close()
	if res != nil {
		return *res
	}

	return http.Response{}
}

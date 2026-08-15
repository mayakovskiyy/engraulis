package client

import (
	"fmt"
	"github.com/mayakovskiyy/engraulis/sysmon"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func Req(address string, delay int, amount int, logging bool, db bool) http.Response {
	var res *http.Response
	var oks, errs int
	homedir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	name := address
	filename := filepath.Join(homedir, "Documents", fmt.Sprintf("log_%s.txt", time.Now().Format("2006-01-02_15-04-05")))

	switch {
	case delay < 0:
		fmt.Println("engraulis: Delay must be bigger or equal 1")
	case delay == 0:
		fmt.Println("engraulis: Delay must be bigger or equal 1")
	}

	for i := 0; i < amount; i++ {
		res, err = http.Get(address)
		statusIcon := ""
		status := ""

		switch {
		case res.StatusCode == 200:
			statusIcon = "🟢"
			status = "(OK)"
			oks += 1
		default:
			statusIcon = "🔴"
			status = "(ERROR)"
			errs += 1
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
		if db {
			sysmon.SaveDataWebMon(address, res.StatusCode)
		}

		time.Sleep(time.Duration(delay) * time.Second)
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

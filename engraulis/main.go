package main

import (
	"flag"
	"fmt"
	"github.com/mayakovskiyy/engraulis/client"
	"github.com/mayakovskiyy/engraulis/sysmon"
	"os"
	"path/filepath"
)

// usage example
// uncomment the 4th line to use the website monitoring

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	webmonPath := filepath.Join(homeDir, "Documents", "engraulisWM.db")
	defaultPath := filepath.Join(homeDir, "Documents", "engraulis.db")

	sysMon := flag.Bool("sysmon", false, "System monitoring enabling.")
	delaySysmon := flag.Int("delay", 10, "Delay between delta calculating.")
	samplingRate := flag.Int("sampling_rate", 3, "Delay between sampling rate.")
	logs := flag.Bool("logging", false, "Sysmon DB (Sqlite3) logging.")
	dbPath := flag.String("db_path", defaultPath, "DB Path. Works only when the logs flag enabled.")
	websiteMonitoring := flag.Bool("web_monitoring", false, "Website monitoring.")
	websiteMonitoringAddress := flag.String("web_address", "https://example.com/", "Website address (for website monitoring).")
	delayWebsiteMonitoring := flag.Int("web_delay", 3, "Website monitoring delay.")
	amountWebsiteMonitoring := flag.Int("web_amount", 3, "Website monitoring amount.")
	logsWebsiteMonitoring := flag.Bool("web_logging", false, "Website Monitoring logs.")
	dbLogsWebsiteMonitoring := flag.Bool("web_db", false, "DB for the Website Monitoring. ")

	flag.Parse()

	switch {
	case *websiteMonitoring:
		res := client.Req(*websiteMonitoringAddress, *delayWebsiteMonitoring, *amountWebsiteMonitoring, *logsWebsiteMonitoring, *dbLogsWebsiteMonitoring)
		fmt.Printf("Status: %d \n", res.StatusCode)
		fmt.Printf("Server: %s \n", *websiteMonitoringAddress) // also you're able to use: res.Header.Get("Server"), but address := ... works better imo
		fmt.Printf("Date: %s \n", res.Header.Get("Date"))
		if *dbLogsWebsiteMonitoring {
			sysmon.InitDatabase(webmonPath, true)
		case *sysMon:
			if *logs == true {
				sysmon.InitDatabase(*dbPath, false)
			}
			sysmt := sysmon.Monitoring(*delaySysmon, *samplingRate, *logs) // the first value stands for minutes between deltas (in minutes). the second value stands for sampling rate delay (in seconds)
			fmt.Println(sysmt)
	}
}

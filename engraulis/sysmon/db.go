package sysmon

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase(dbPath string) {
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println(err)
	}

	db.SetMaxOpenConns(1)

	createSQLTable := `CREATE TABLE IF NOT EXISTS sys_metrics (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    min_usage INTEGER NOT NULL, 
    max_usage INTEGER NOT NULL,
    delta INTEGER NOT NULL,
    date DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createSQLTable)
	if err != nil {
		fmt.Println(err)
	}

}

func SaveData(minval int, maxval int, deltaval int) {
	_, err := db.Exec("INSERT INTO sys_metrics (min_usage, max_usage, delta) VALUES (?, ?, ?)", minval, maxval, deltaval)
	if err != nil {
		fmt.Println(err)
	}
}

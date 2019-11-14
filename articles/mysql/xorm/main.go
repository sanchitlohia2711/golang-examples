package main

import (
	"database/sql"
	"fmt"
	_ "go-sql-driver/mysql"
)

var masterDB = *sql.DB

func main() {
	dataSourceString := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", "root", "yourpassword",  127.0.0.1:3006)
	masterDB, err = sql.Open("mysql", dataSourceString)
	if err != nil {
       panic("Unable to connec to database")
	}
}

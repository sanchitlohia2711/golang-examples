package main

import (
	"go-sql-driver/mysql"

	"github.com/go-sql-driver/mysql"

	"github.com/golang-examples/articles/mysql/gosqldriver"
)

func main() {
	gosqldriver.Initialize()
	_ := mysql.Config

}

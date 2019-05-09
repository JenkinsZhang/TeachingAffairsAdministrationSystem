package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect() {
	tmp, err := sql.Open("mysql", "root:miao@(111.230.97.43:3306)/school?charset=utf8")
	if err != nil {
		fmt.Fprintf(os.Stderr, "connecting: %v\n", err)
		os.Exit(1)
	}
	Db = tmp
}

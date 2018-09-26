package database

import (
"fmt"
"database/sql"
_ "github.com/go-sql-driver/mysql"

)
var (
    // DBCon is the connection handle
    // for the database
    DBCon *sql.DB
)

func Connect() {
	var err error
    DBCon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hopeip")
	if err != nil {
		fmt.Print(err.Error())
	}
	// make sure connection is available
	err = DBCon.Ping()
	
	if err != nil {
		fmt.Print(err.Error())
	}
}


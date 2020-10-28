package connectdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //imported but not used
)

func Connect() *sql.DB {
	datasource := "root:password@tcp(localhost:3306)/go_backend01"
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Println("error connecting to db", err)
	}
	fmt.Println("db is connected")
	err = db.Ping()
	if err != nil {
		log.Println("error pinging db", err)
	}
	return db

}

//DB exported variable
var DB = Connect()

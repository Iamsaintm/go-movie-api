package repo

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	databaseString := os.Getenv("DATABASE")
	db, error := sql.Open("mysql", databaseString)
	if error != nil {
		fmt.Println("SERVER ERROR", error.Error())
	}
	error = db.Ping()
	if error != nil {
		fmt.Println("SERVER ERROR", error.Error())
	}
	fmt.Println("Connected to Database")
	return db
}

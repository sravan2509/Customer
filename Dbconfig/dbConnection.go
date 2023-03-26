package Dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() {
	db, err := sql.Open("mysql", "root:Sravan@2509@tcp(localhost:3306)/golang")
	if err != nil {
		fmt.Println("Error in connecting the DB", err)
	}
	fmt.Println("Database is connected successfully")
	defer db.Close()
}

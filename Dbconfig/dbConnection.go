package Dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Sravan@2509@tcp(localhost:3306)/golang")
	if err != nil {
		fmt.Println("Error in connecting the DB", err)
		return nil, err
	}
	fmt.Println("Database is connected successfully")
	return db, nil
}

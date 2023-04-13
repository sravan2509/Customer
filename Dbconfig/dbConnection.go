package Dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	Config "github.com/sravan2509/Customer/Config"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", Config.GetConnectionString())
	if err != nil {
		fmt.Println("Error in connecting the DB", err)
		return nil, err
	}
	fmt.Println("Database is connected successfully")
	return db, nil
}

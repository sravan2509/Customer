package main

import (
	// "database/sql"
	"fmt"
	"net/http"

	// _ "github.com/go-sql-driver/mysql"
	Dbconfig "github.com/sravan2509/Customer/tree/main/UserAccountManagement/Dbconfig"
	// Handler "github.com/sravan2509/Customer/tree/main/UserAccountManagement/Endpoints"
)

func main() {
	Dbconfig.DBConnection

	// db, err := sql.Open("mysql", "root:Sravan@2509@tcp(localhost:3306)/golang")
	// if err != nil {
	// 	fmt.Println("Error in connecting the DB", err)
	// }
	// fmt.Println("Database is connected successfully")
	// defer db.Close()

	// //handling routes
	// http.HandleFunc("/Signup", Handler.signupHandler)
	// http.HandleFunc("/changePassword", Handler.changePasswordHandler)
	// http.HandleFunc("/login", Handler.loginHandler)
	// http.HandleFunc("/deleteCustomer", Handler.deleteCustomerHandler)

	//hosting the server
	err = http.ListenAndServe(":8080", nil)
	fmt.Println("Local host is servered at port 8080")
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}

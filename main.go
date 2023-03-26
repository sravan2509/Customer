package main

import (
	"fmt"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/tree/main/UserAccountManagement/Dbconfig"
	// Handler "github.com/sravan2509/Customer/tree/main/UserAccountManagement/Endpoints"
)

func main() {
	Dbconfig.DBConnection

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

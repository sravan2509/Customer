package main

import (
	"fmt"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Handler "github.com/sravan2509/Customer/Endpoints"
)

func main() {
	db, err := Dbconfig.DBConnection()
	defer db.Close()

	//handling routes
	http.HandleFunc("/Signup", Handler.SignupHandler)
	http.HandleFunc("/changePassword", Handler.ChangePasswordHandler)
	http.HandleFunc("/login", Handler.LoginHandler)
	http.HandleFunc("/deleteCustomer", Handler.DeleteCustomerHandler)
	http.HandleFunc("/getCustomers", Handler.GetAllCustomersHandler)

	//hosting the server
	fmt.Println("Local host is servered at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}

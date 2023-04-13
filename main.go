package main

import (
	"fmt"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Handler "github.com/sravan2509/Customer/Endpoints"
	Token "github.com/sravan2509/Customer/TokenHandler"
)

func main() {
	db, err := Dbconfig.DBConnection()
	defer db.Close()

	//handling routes
	http.HandleFunc("/Signup", Handler.SignupHandler)
	http.HandleFunc("/login", Handler.LoginHandler)
	http.Handle("/changePassword", Token.AuthMiddleware(http.HandlerFunc(Handler.ChangePasswordHandler)))
	http.Handle("/deleteCustomer", Token.AuthMiddleware(http.HandlerFunc(Handler.DeleteCustomerHandler)))
	http.Handle("/getCustomers", Token.AuthMiddleware(http.HandlerFunc(Handler.GetAllCustomersHandler)))

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}

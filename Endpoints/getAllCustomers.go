package endpoint

import (
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
)

func GetAllCustomersHandler(w http.ResponseWriter, r *http.Request) {

	//valiadte method
	if r.Method != "GET" {
		ResponseFormat(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	//DB connection
	db, err := Dbconfig.DBConnection()
	if err != nil {
		ResponseFormat(w, "DB Connection Failed", http.StatusInternalServerError, nil)
		return
	}
	defer db.Close()

	//Get all customers
	StatusCode, err, Result := Dbconfig.GetCustomers(db)
	if err != nil {
		ResponseFormat(w, err.Error(), StatusCode, nil)
	}

	//formating the result to []Schema.Customer
	var customers []Schema.Customer
	for Result.Next() {
		var customer Schema.Customer
		err := Result.Scan(&customer.Name, &customer.Email, &customer.PhoneNumber, &customer.Password, &customer.Address)
		if err != nil {
			ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
			panic(err.Error())
		}
		customers = append(customers, customer)
	}

	ResponseFormat(w, "Customers", http.StatusOK, customers)

}

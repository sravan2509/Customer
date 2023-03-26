package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	schema "github.com/sravan2509/Customer/Schema"
	validation "github.com/sravan2509/Customer/Validation"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//signup only for post method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := Dbconfig.DBConnection()
	defer db.Close()
	//reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//unmarshal the body
	var newCustomer schema.Customer
	err = json.Unmarshal(body, &newCustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//validating the new customer
	if !validation.IsEmailValid(newCustomer.Email) {
		http.Error(w, "Email is not valid", http.StatusBadRequest)
		return
	}
	if validation.IsCustomerExist(newCustomer.Email) {
		http.Error(w, "Customer already exists", http.StatusBadRequest)
		return
	}
	// if !isPasswordValid(newCustomer.Password) {
	// 	http.Error(w, "Password is not valid", http.StatusBadRequest)
	// 	return
	// }
	if newCustomer.Password == "" || newCustomer.Password == " " {
		http.Error(w, "Password is not valid", http.StatusBadRequest)
		return
	}

	//adding the new customer to DB
	stmt, err := db.Prepare(`INSERT INTO customers(Name,Email,PhoneNumber,Password,Address) Values (?,?,?,?,?)`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(newCustomer.Name, newCustomer.Email, newCustomer.PhoneNumber, newCustomer.Password, newCustomer.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Customer %s created successfully", newCustomer.Email)
	w.WriteHeader(http.StatusCreated)
}

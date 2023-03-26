package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
	Validation "github.com/sravan2509/Customer/Validation"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

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
	var logincustomer Schema.LoginCustomer
	err = json.Unmarshal(body, &logincustomer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//validating the new logincustomer
	if !Validation.IsEmailValid(logincustomer.Email) {
		http.Error(w, "Email is not valid", http.StatusBadRequest)
		return
	}
	// if !isPasswordValid(logincustomer.Password) {
	// 	http.Error(w, "Password is not valid", http.StatusBadRequest)
	// 	return
	// }
	if logincustomer.Password == "" || logincustomer.Password == " " {
		http.Error(w, "Password is not valid", http.StatusBadRequest)
		return
	}

	if !Validation.IsCustomerExist(logincustomer.Email) {
		http.Error(w, "Customer with email address not found", http.StatusBadRequest)
		return
	}

	if !Validation.IsLoginValid(logincustomer.Email, logincustomer.Password) {
		http.Error(w, "Incorrect Password", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Login succesful!!!!")
	w.WriteHeader(http.StatusOK)

}

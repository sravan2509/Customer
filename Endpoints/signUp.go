package endpoint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
	Token "github.com/sravan2509/Customer/TokenHandler"
	Validation "github.com/sravan2509/Customer/Validation"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//signup only for post method
	if r.Method != "POST" {
		ResponseFormat(w, "Method not allowed", http.StatusMethodNotAllowed, nil)
		return
	}

	db, err := Dbconfig.DBConnection()
	if err != nil {
		ResponseFormat(w, "DB Connection Failed", http.StatusInternalServerError, nil)
		return
	}
	defer db.Close()
	//reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	//unmarshal the body
	var newCustomer Schema.Customer
	err = json.Unmarshal(body, &newCustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	//validating the new customer
	if ErrorCode, err := Validation.SignupValidation(db, newCustomer); err != nil {
		ResponseFormat(w, err.Error(), ErrorCode, nil)
		return
	}

	//adding the new customer to DB
	if Errorcode, err := Dbconfig.InsertCustomer(db, newCustomer); err != nil {
		ResponseFormat(w, err.Error(), Errorcode, nil)
		return
	}

	token, err := Token.GenerateToken(newCustomer.Email)

	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
	}
	ResponseFormat(w, fmt.Sprintf("Customer %s created successfully", newCustomer.Email), http.StatusCreated, token)
}

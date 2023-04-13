package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
	Validation "github.com/sravan2509/Customer/Validation"
)

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {

	TokenEmail, ok := r.Context().Value("Email").(string)
	if !ok {
		ResponseFormat(w, "Email not found in context", http.StatusInternalServerError, nil)
		return
	}

	if r.Method != "PUT" {
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
	var changecustomerlogin Schema.ChangeLoginPassword
	err = json.Unmarshal(body, &changecustomerlogin)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	//validating the new logincustomer
	if StatusCode, err := Validation.ChangePasswordValidation(db, changecustomerlogin, TokenEmail); err != nil {
		ResponseFormat(w, err.Error(), StatusCode, nil)
		return
	}

	// updating the password
	if status, err := Dbconfig.UpdateCustomer(db, changecustomerlogin); err != nil {
		ResponseFormat(w, err.Error(), status, nil)
		return
	}
	ResponseFormat(w, "Customer Password is updated successfully!", http.StatusOK, nil)

}

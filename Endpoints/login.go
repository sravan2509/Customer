package endpoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
	Token "github.com/sravan2509/Customer/TokenHandler"
	Validation "github.com/sravan2509/Customer/Validation"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

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
	var logincustomer Schema.Customer
	err = json.Unmarshal(body, &logincustomer)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}

	//validating the  logincustomer
	if statusCode, err := Validation.LoginValidation(db, logincustomer); err != nil {
		ResponseFormat(w, err.Error(), statusCode, nil)
		return
	}

	token, err := Token.GenerateToken(logincustomer.Email)

	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
	}

	ResponseFormat(w, "Login Success!", http.StatusOK, token)

}

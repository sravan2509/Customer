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

func changePasswordHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
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
	var changecustomerlogin schema.ChangeLogin
	err = json.Unmarshal(body, &changecustomerlogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf(changecustomerlogin.NewPassword)

	//validating the new logincustomer
	if !validation.IsEmailValid(changecustomerlogin.Email) {
		http.Error(w, "Email is not valid", http.StatusBadRequest)
		return
	}
	if !validation.IsCustomerExist(changecustomerlogin.Email) {
		http.Error(w, "Customer with email address not found", http.StatusBadRequest)
		return
	}
	// if !isPasswordValid(changecustomerlogin.Password) {
	// 	http.Error(w, "Password is not valid", http.StatusBadRequest)
	// 	return
	// }
	if changecustomerlogin.OldPassword == "" || changecustomerlogin.OldPassword == " " {
		http.Error(w, "Old Password is not valid", http.StatusBadRequest)
		return
	}
	if changecustomerlogin.NewPassword == "" || changecustomerlogin.NewPassword == " " {
		http.Error(w, "New Password is not valid", http.StatusBadRequest)
		return
	}

	if !validation.IsLoginValid(changecustomerlogin.Email, changecustomerlogin.OldPassword) {
		http.Error(w, "oldPassword is Incorrect", http.StatusBadRequest)
		return
	}

	_, err = db.Query(`UPDATE customers SET Password = ? WHERE Email = ?`, changecustomerlogin.NewPassword, changecustomerlogin.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Customer Password is updated successfully!")
	w.WriteHeader(http.StatusOK)

}

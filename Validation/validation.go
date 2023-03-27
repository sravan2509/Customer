package customerValidation

import (
	"database/sql"
	"errors"
	"net/http"

	Dbconfig "github.com/sravan2509/Customer/Dbconfig"
	Schema "github.com/sravan2509/Customer/Schema"
)

func SignupValidation(db *sql.DB, newCustomer Schema.Customer) (int, error) {

	if !IsEmailValid(newCustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if Dbconfig.IsCustomerExist(db, newCustomer.Email) {
		return http.StatusForbidden, errors.New("Customer already exists")
	}
	if !IsPasswordValid(newCustomer.Password) {
		return http.StatusUnauthorized, errors.New("Password is not Valid")
	}
	if newCustomer.Password != newCustomer.ConformPassword {
		return http.StatusBadRequest, errors.New("Passwords Mismatch")
	}

	return http.StatusCreated, nil
}

func LoginValidation(db *sql.DB, logincustomer Schema.Customer) (int, error) {

	if !IsEmailValid(logincustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if !Dbconfig.IsCustomerExist(db, logincustomer.Email) {
		return http.StatusBadRequest, errors.New("Customer with email address not found")
	}
	if !IsPasswordValid(logincustomer.Password) {
		return http.StatusUnauthorized, errors.New("Password is not Valid")
	}
	if status, err := Dbconfig.IsLoginValid(db, logincustomer.Email, logincustomer.Password); err != nil {
		return status, err
	}

	return http.StatusAccepted, nil
}

func DeleteCustomerValidation(db *sql.DB, Email string) (int, error) {

	if !IsEmailValid(Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if !Dbconfig.IsCustomerExist(db, Email) {
		return http.StatusBadRequest, errors.New("Customer with email not found")
	}

	return http.StatusOK, nil
}

func ChangePasswordValidation(db *sql.DB, changecustomerlogin Schema.Customer) (int, error) {

	if !IsEmailValid(changecustomerlogin.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if !Dbconfig.IsCustomerExist(db, changecustomerlogin.Email) {
		return http.StatusBadRequest, errors.New("Customer with email not found")
	}
	if !IsPasswordValid(changecustomerlogin.Password) {
		return http.StatusUnauthorized, errors.New("Password is not Valid")
	}
	if !IsPasswordValid(changecustomerlogin.NewPassword) {
		return http.StatusUnauthorized, errors.New("New Password is not Valid")
	}
	if status, err := Dbconfig.IsLoginValid(db, changecustomerlogin.Email, changecustomerlogin.Password); err != nil {
		return status, err
	}

	return 204, nil
}

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
		return http.StatusBadRequest, errors.New("Password is not Valid")
	}
	if newCustomer.Password != newCustomer.ConformPassword {
		return http.StatusBadRequest, errors.New("Passwords Mismatch")
	}
	if !IsPhoneNumberValid(newCustomer.PhoneNumber) {
		return http.StatusBadRequest, errors.New("Invalid phone number")
	}
	if newCustomer.Name == "" {
		return http.StatusBadRequest, errors.New("Name is required")
	}

	return http.StatusCreated, nil
}

func LoginValidation(db *sql.DB, logincustomer Schema.LoginCustomer) (int, error) {

	if !IsEmailValid(logincustomer.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if !IsPasswordValid(logincustomer.Password) {
		return http.StatusUnauthorized, errors.New("Password is not Valid")
	}
	if status, err := Dbconfig.IsLoginValid(db, logincustomer.Email, logincustomer.Password); err != nil {
		return status, err
	}

	return http.StatusAccepted, nil
}

func DeleteCustomerValidation(db *sql.DB, Email string, TokenEmail string) (int, error) {

	if !IsEmailValid(Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if TokenEmail != Email {
		return http.StatusBadRequest, errors.New("Invalid Email login")
	}
	if !Dbconfig.IsCustomerExist(db, Email) {
		return http.StatusBadRequest, errors.New("Customer with email not found")
	}

	return http.StatusOK, nil
}

func ChangePasswordValidation(db *sql.DB, changecustomerlogin Schema.ChangeLoginPassword, TokenEmail string) (int, error) {

	if !IsEmailValid(changecustomerlogin.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if TokenEmail != changecustomerlogin.Email {
		return http.StatusBadRequest, errors.New("Invalid Email login")
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

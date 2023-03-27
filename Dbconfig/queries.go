package Dbconfig

import (
	"database/sql"
	"errors"
	"net/http"

	Schema "github.com/sravan2509/Customer/Schema"
)

func InsertCustomer(db *sql.DB, newCustomer Schema.Customer) (int, error) {

	stmt, err := db.Prepare(`INSERT INTO customers(Name,Email,PhoneNumber,Password,Address) Values (?,?,?,?,?)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = stmt.Exec(newCustomer.Name, newCustomer.Email, newCustomer.PhoneNumber, newCustomer.Password, newCustomer.Address)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil

}

func DeleteCustomer(db *sql.DB, Email string) (int, error) {
	_, err := db.Query(`DELETE FROM customers WHERE Email = ?`, Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 204, nil
}

func UpdateCustomer(db *sql.DB, changecustomerlogin Schema.Customer) (int, error) {
	_, err := db.Query(`UPDATE customers SET Password = ? WHERE Email = ?`, changecustomerlogin.NewPassword, changecustomerlogin.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func GetCustomers(db *sql.DB) (int, error, *sql.Rows) {
	result, err := db.Query(`SELECT * FROM customers `)
	if err != nil {
		return http.StatusInternalServerError, err, nil
	}
	return http.StatusOK, nil, result
}

func IsLoginValid(db *sql.DB, email string, password string) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE Email = ? AND Password = ?", email, password).Scan(&count)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("Invalid Password")
	}
	return http.StatusOK, nil
}

func IsCustomerExist(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE Email = ?", email).Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

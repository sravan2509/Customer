package Dbconfig

import (
	"database/sql"
	"errors"
	"net/http"

	Schema "github.com/sravan2509/Customer/Schema"
	"golang.org/x/crypto/bcrypt"
)

func InsertCustomer(db *sql.DB, newCustomer Schema.Customer) (int, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newCustomer.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	stmt, err := db.Prepare(`INSERT INTO customers(Name,Email,PhoneNumber,Password,Address) Values (?,?,?,?,?)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = stmt.Exec(newCustomer.Name, newCustomer.Email, newCustomer.PhoneNumber, string(hashedPassword), newCustomer.Address)
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

func UpdateCustomer(db *sql.DB, changecustomerlogin Schema.ChangeLoginPassword) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changecustomerlogin.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = db.Query(`UPDATE customers SET Password = ? WHERE Email = ?`, string(hashedPassword), changecustomerlogin.Email)
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
	var hashedPassword string
	err := db.QueryRow("SELECT Password FROM customers WHERE Email = ?", email).Scan(&hashedPassword)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if hashedPassword == "" {
		return http.StatusBadRequest, errors.New("Email Not Found")
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return http.StatusBadRequest, errors.New("Incorrect Password")
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

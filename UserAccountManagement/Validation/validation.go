package customerValidation

import (
	"database/sql"
	"fmt"
	"regexp"
)

func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		fmt.Println("Error")
		return false
	}
	return match
}

func IsPasswordValid(password string) bool {
	pattern := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	match, err := regexp.MatchString(pattern, password)
	if err != nil {
		return false
	}
	return match
}

func IsLoginValid(db *sql.DB, email string, password string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM customers WHERE Email = ? AND Password = ?", email, password).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
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

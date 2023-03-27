package customerValidation

import (
	"regexp"
)

func IsEmailValid(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, email)
	if err != nil {
		return false
	}
	return match
}

func IsPasswordValid(password string) bool {
	// pattern := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	// match, err := regexp.MatchString(pattern, password)
	// if err != nil {
	// 	return false
	// }
	// return match
	if len(password) < 8 {
		return false
	}
	return true
}

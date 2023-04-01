package Token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	config "github.com/sravan2509/Customer/Config"
)

func GenerateToken(Email string) (string, error) {

	claims := jwt.MapClaims{
		"Email": Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey())
}

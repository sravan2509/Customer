package Token

import (
	"time"

	"github.com/golang-jwt/jwt"
	Config "github.com/sravan2509/Customer/Config"
)

func GenerateToken(Email string) (string, error) {

	claims := jwt.MapClaims{
		"Email": Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Config.SecretKey())
}

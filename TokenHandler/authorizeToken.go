package Token

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	config "github.com/sravan2509/Customer/Config"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get the token
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		fmt.Println(tokenString)
		// parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return config.SecretKey(), nil
		})
		fmt.Println(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate the token and return the correponding error or continue with the handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			Email, ok := claims["Email"].(string)
			if !ok {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			//assigning the email to the context of request
			ctx := context.WithValue(r.Context(), "Email", Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}

package Token

import (
	"net/http"
	"time"
)

func RemoveToken(w http.ResponseWriter, r *http.Request) {
	// Remove the token from the client by setting its expiration time to a past date
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

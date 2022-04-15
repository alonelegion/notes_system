package jwt

import (
	"github.com/cristalhq/jwt"
	"net/http"
	"strings"
)

type UserClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func JWTMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		if len(authHeader) != 2 {

		}
	}
}

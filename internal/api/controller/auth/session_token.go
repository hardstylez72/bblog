package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type MyCustomClaims struct {
	jwt.StandardClaims
	User
}

func ExtractUserFromCookies(cookies *http.Cookie) (*User, error) {
	var claims MyCustomClaims

	_, err := jwt.ParseWithClaims(cookies.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyBytes, nil
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

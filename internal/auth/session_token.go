package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var secret = viper.GetString("oauth.sessionCookie.secret")

type tokenClaims struct {
	jwt.StandardClaims
	User
}

type User struct {
	ExternalUser
	SessionId    string `json:"sessionId"`
	IsAuthorized bool   `json:"isAuthorized"`
	UserId       string `json:"userId"`
}

func GenerateTokenForGuest(sessionId string) (string, error) {
	return GenerateToken(&User{
		ExternalUser: ExternalUser{},
		SessionId:    sessionId,
		IsAuthorized: false,
		UserId:       "",
	})
}

func GenerateToken(user *User) (string, error) {

	var email, login, name string

	if user.Email.Valid {
		email = user.Email.String
	}

	if user.Login.Valid {
		login = user.Login.String
	}

	if user.Name.Valid {
		name = user.Name.String
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"externalId":   user.ExternalId,
		"authType":     user.AuthType,
		"email":        email,
		"login":        login,
		"name":         name,
		"sessionId":    user.SessionId,
		"isAuthorized": user.IsAuthorized,
	})

	return token.SignedString([]byte(secret))
}

func ExtractUserAndTokenFromStringToken(tokenString string) (*User, *jwt.Token, error) {
	var claims tokenClaims

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("invalid signing method, expected HMAC, got: %s", token.Method.Alg())
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, nil, err
	}

	return &claims.User, token, nil
}

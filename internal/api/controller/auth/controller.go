package auth

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	rm          map[string]bool // todo: remove!!
	verifyKey   *rsa.PublicKey
	signKey     *rsa.PrivateKey
	signBytes   []byte
	verifyBytes []byte
)

func init() {
	signBytes, err := ioutil.ReadFile("internal/api/controller/auth/jwtRS256.key")
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile("internal/api/controller/auth/jwtRS256.key.pub")
	if err != nil {
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}

type controller struct {
	github OAuth2
	google OAuth2
}

type Oauth struct {
	Google              Config
	Github              Config
	SessionCookieConfig SessionCookieConfig
}

type UserRedirects struct {
	OnSuccess string
	OnFailure string
}

type Config struct {
	RedirectURL   string
	ClientID      string
	ClientSecret  string
	Scopes        []string
	UserInfoURL   string
	UserRedirects UserRedirects
}

type SessionCookieConfig struct {
	Name   string
	Domain string
	Path   string
	MaxAge int
	Secure bool
}

func NewAuthController(oauth Oauth) *controller {

	rm = make(map[string]bool)

	return &controller{
		github: NewGithubOAuth2Controller(oauth.Github, oauth.SessionCookieConfig),
		google: NewGoogleOAuth2Controller(oauth.Google, oauth.SessionCookieConfig),
	}
}

func (a *controller) Mount(r chi.Router) {

	r.Route("/v1", func(r chi.Router) {
		r.HandleFunc("/google/oauth/login", a.google.HandleLogin)
		r.HandleFunc("/google/oauth/callback", a.google.HandleCallback)

		r.HandleFunc("/github/oauth/login", a.github.HandleLogin)
		r.HandleFunc("/github/oauth/callback", a.github.HandleCallback)
	})
}

func setSessionCookie(w http.ResponseWriter, token string, cookie SessionCookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookie.Name,
		Value:   token,
		MaxAge:  cookie.MaxAge,
		Expires: time.Now().Add(time.Hour),
		Path:    cookie.Path,
		Secure:  cookie.Secure,
	})
}

func GenerateToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userExternalId": user.ExternalId,
		"authType":       user.AuthType,
		"email":          user.Email.String,
		"login":          user.Login.String,
		"name":           user.Name.String,
	})

	return token.SignedString(signBytes)
}

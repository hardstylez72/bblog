package auth

import (
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

var (
	rm map[string]bool // todo: remove!!
)

type controller struct {
	github OAuth2
	google OAuth2
}

type Oauth struct {
	Google              Config
	Github              Config
	SessionCookieConfig SessionCookieConfig
}

type SessionCookieConfig struct {
	Name   string
	Domain string
	Path   string
	MaxAge int
	Secure bool
}

type Config struct {
	RedirectURL   string
	ClientID      string
	ClientSecret  string
	Scopes        []string
	UserInfoURL   string
	UserRedirects UserRedirects
}

type UserRedirects struct {
	OnSuccess string
	OnFailure string
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

func SetSessionCookie(w http.ResponseWriter, token string, cookie SessionCookieConfig) {
	http.SetCookie(w, &http.Cookie{
		Name:    cookie.Name,
		Value:   token,
		MaxAge:  cookie.MaxAge,
		Expires: time.Now().Add(time.Hour),
		Path:    cookie.Path,
		Secure:  cookie.Secure,
	})
}

package controller

import (

	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"golang.org/x/oauth2/github"

	"github.com/hardstylez72/bbckend/internal/store"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/yandex"
	"net/http"
	"text/template"

)

type Oauth struct {
	Google AuthConfig
	Yandex AuthConfig
	Github AuthConfig
}

type AuthConfig struct {
	RedirectURL    string
	ClientID       string
	ClientSecret   string
	Scopes         []string
	GetUserInfoURL string
}

var (
	oauthStateString = "pseudo-random" // todo: repace with uuid
)


func HandleMain(w http.ResponseWriter, r *http.Request) {
	var htmlIndex = `<html>
<body>
	<a href="/api/v1/google/oauth/login">Google Log In</a> <br>
	<a href="/api/v1/yandex/oauth/login">Yandex Log In</a> <br>
	<a href="/api/v1/github/oauth/login">Github Log In</a>
</body>
</html>`
	fmt.Fprintf(w, htmlIndex)
}

type externalAuthTypeMap map[int]string

type authController struct {
	googleOauthConfig   *oauth2.Config
	googleGetUserInfoURLTmpl  *template.Template

	yandexOauthConfig   *oauth2.Config
	yandexGetUserInfoURLTmpl  *template.Template

	githubOauthConfig   *oauth2.Config
	githubGetUserInfoURLTmpl  string

	userStore           store.Store
	externalAuthTypeMap externalAuthTypeMap
	httClient *http.Client
}

const (
	authTypeGoogle = iota
	authTypeYandex
	authTypeGithub
)

func NewAuthController(oauth Oauth, userStore store.Store) (*authController, error) {
	 googleOauthConfig := &oauth2.Config{
		RedirectURL:  oauth.Google.RedirectURL,
		ClientID:     oauth.Google.ClientID,
		ClientSecret: oauth.Google.ClientSecret,
		Scopes:       oauth.Google.Scopes,
		Endpoint:     google.Endpoint,
	}

	githubOauthConfig := &oauth2.Config{
		RedirectURL:  oauth.Github.RedirectURL,
		ClientID:     oauth.Github.ClientID,
		ClientSecret: oauth.Github.ClientSecret,
		Scopes:       oauth.Github.Scopes,
		Endpoint:     github.Endpoint,
	}


	yandexOauthConfig := &oauth2.Config{
		RedirectURL:  oauth.Yandex.RedirectURL,
		ClientID:     oauth.Yandex.ClientID,
		ClientSecret: oauth.Yandex.ClientSecret,
		Scopes:       oauth.Yandex.Scopes,
		Endpoint:     yandex.Endpoint,
	}

	typeMap := make(externalAuthTypeMap)
	typeMap[authTypeGoogle] = "d7c8e05b-156b-4e86-811f-30b26bb81240"
	typeMap[authTypeYandex] = "b65cb084-095a-4b48-9913-6fcbb1afa21e"
	typeMap[authTypeGithub] = "e97892e1-7507-440a-80fa-e43b74aff469"


	googleTmpl, err := template.New("accessTokenURL").Parse(oauth.Google.GetUserInfoURL)
	if err != nil {
		return nil, errors.New("cant parse accessTokenURL: " + oauth.Google.GetUserInfoURL)
	}

	yandexTmpl, err := template.New("accessTokenURL").Parse(oauth.Yandex.GetUserInfoURL)
	if err != nil {
		return nil, errors.New("cant parse accessTokenURL: " + oauth.Yandex.GetUserInfoURL)
	}


	return &authController{
		googleOauthConfig:        googleOauthConfig,
		googleGetUserInfoURLTmpl: googleTmpl,
		yandexOauthConfig:        yandexOauthConfig,
		yandexGetUserInfoURLTmpl: yandexTmpl,
		githubOauthConfig:        githubOauthConfig,
		githubGetUserInfoURLTmpl: oauth.Github.GetUserInfoURL,
		userStore:                userStore,
		externalAuthTypeMap:      typeMap,
		httClient:                new(http.Client),
	}, nil
}

func (a *authController) Mount(r chi.Router) {

	r.Route("/v1",func(r chi.Router) {
		r.HandleFunc("/google/oauth/login", a.HandleGoogleLogin)
		r.HandleFunc("/google/oauth/callback", a.HandleGoogleCallback)

		r.HandleFunc("/yandex/oauth/login", a.HandleYandexLogin)
		r.HandleFunc("/yandex/oauth/callback", a.HandleYandexCallback)

		r.HandleFunc("/github/oauth/login", a.HandleGithubLogin)
		r.HandleFunc("/github/oauth/callback", a.HandleGithubCallback)
	})

}

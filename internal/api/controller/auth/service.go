package auth

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"net/http"
	"time"
)

var rm map[string]bool // todo: remove!!

type controller struct {
	github OAuth2
	google OAuth2
}

type Oauth struct {
	Google Config
	Yandex Config
	Github Config
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

var (
	oauthStateString = "pseudo-random" // todo: replace with uuid
)

func NewAuthController(oauth Oauth, userStore user.Storage) *controller {

	rm = make(map[string]bool)

	return &controller{
		github: NewGithubOAuth2Controller(oauth.Github, userStore),
		google: NewGoogleOAuth2Controller(oauth.Google, userStore),
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

func resolveUser(ctx context.Context, userStore user.Storage, extUser *user.User, authType string) (userId string, err error) {

	u, err := userStore.GetUserByExternalId(ctx, extUser.ExternalId, authType)
	if err != nil {
		if err == user.ErrNotFound {
			err = userStore.SaveUser(ctx, extUser)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	u, err = userStore.GetUserByExternalId(ctx, extUser.ExternalId, authType)
	if err != nil {
		return "", err
	}

	return u.Id, nil
}

func setUserCookie(w http.ResponseWriter, userId string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "user_id",
		Value:   userId,
		MaxAge:  999999999,
		Expires: time.Now().Add(time.Hour),
		Path:    "/",
	})
}

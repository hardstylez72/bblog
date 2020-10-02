package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"time"
)

const (
	authTypeGoogle = "Google"
)

type GoogleOauthUserData struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type googleAuth struct {
	Oauth2Config  *oauth2.Config
	UserInfoURL   string
	UserRedirects UserRedirects
	userStore     user.Storage
	httClient     *http.Client
}

func NewGoogleOAuth2Controller(cfg Config, userStore user.Storage) *googleAuth {

	return &googleAuth{
		Oauth2Config: &oauth2.Config{
			RedirectURL:  cfg.RedirectURL,
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       cfg.Scopes,
			Endpoint:     google.Endpoint,
		},
		UserRedirects: cfg.UserRedirects,
		UserInfoURL:   cfg.UserInfoURL,
		userStore:     userStore,
		httClient:     &http.Client{},
	}
}

func (a *googleAuth) HandleLogin(w http.ResponseWriter, r *http.Request) {
	uniq := uuid.New().String()
	rm[uniq] = true
	url := a.Oauth2Config.AuthCodeURL(uniq)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *googleAuth) HandleCallback(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	googleUser, err := a.GetUser(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	userId, err := resolveUser(ctx, a.userStore, googleUser, authTypeGoogle)
	if err != nil {
		http.Redirect(w, r, a.UserRedirects.OnFailure, http.StatusTemporaryRedirect)
		return
	}

	setUserCookie(w, userId)

	http.Redirect(w, r, a.UserRedirects.OnSuccess, http.StatusTemporaryRedirect)
}

func (a *googleAuth) GetUser(ctx context.Context, state string, code string) (*user.User, error) {
	if !rm[state] {
		return nil, fmt.Errorf("invalid oauth state")
	}
	oauthToken, err := a.GetToken(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	token := oauthToken.(*oauth2.Token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, buildGoogleGetUserUrl(token.AccessToken, a.UserInfoURL), nil)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	response, err := a.httClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	if response.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("authorization error")
	}
	defer response.Body.Close()

	user := new(GoogleOauthUserData)

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(user); err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return convertGoogleUser(user, authTypeGoogle), nil
}

func buildGoogleGetUserUrl(token, baseUrl string) string {
	return baseUrl + "?access_token=" + token
}

func (a *googleAuth) GetToken(ctx context.Context, code string) (interface{}, error) {
	return a.Oauth2Config.Exchange(ctx, code)
}

func convertGoogleUser(googleUser *GoogleOauthUserData, authTypeGoogleId string) *user.User {
	userId := uuid.New().String()

	u := user.User{
		Id:               userId,
		ExternalId:       googleUser.ID,
		ExternalAuthType: authTypeGoogleId,
		RegisteredAt:     time.Now(),
	}

	if googleUser.Email != "" {
		u.Email.String = googleUser.Email
		u.Email.Valid = true
	}

	if googleUser.Name != "" {
		u.Name.String = googleUser.FamilyName + " " + googleUser.Name
		u.Name.Valid = true
	}

	return &u
}

package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hardstylez72/bblog/internal/storage/user"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
	"strconv"
	"time"
)

const (
	authTypeGithub = "Github"
)

type GithubOauthUserData struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              interface{} `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          interface{} `json:"location"`
	Email             string      `json:"email,omitempty"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

type githubAuth struct {
	OauthConfig *oauth2.Config
	UserInfoURL string

	userStore user.Storage
	httClient *http.Client
}

func NewGithubOAuth2Controller(cfg Config, userStore user.Storage) *githubAuth {
	return &githubAuth{
		OauthConfig: &oauth2.Config{
			RedirectURL:  cfg.RedirectURL,
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			Scopes:       cfg.Scopes,
			Endpoint:     github.Endpoint,
		},
		UserInfoURL: cfg.UserInfoURL,
		userStore:   userStore,
		httClient:   &http.Client{},
	}
}

func (a *githubAuth) HandleLogin(w http.ResponseWriter, r *http.Request) {
	uniq := uuid.New().String()
	rm[uniq] = true
	url := a.OauthConfig.AuthCodeURL(uniq)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *githubAuth) HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	extUser, err := a.GetUser(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	_, err = a.userStore.GetUserByExternalId(ctx, extUser.Id, authTypeGithub)
	if err != nil {
		if err == user.ErrNotFound {
			err = a.userStore.SaveUserWithEmail(ctx, extUser)
			if err != nil {
				return
				// todo: handle
			}
		} else {
			return
		}
	}

	http.Redirect(w, r, "/protected", http.StatusTemporaryRedirect)
}

func (a *githubAuth) GetToken(ctx context.Context, code string) (interface{}, error) {
	return a.OauthConfig.Exchange(ctx, code)
}

func (a *githubAuth) GetUser(ctx context.Context, state string, code string) (*user.UserAgr, error) {
	if !rm[state] {
		return nil, fmt.Errorf("invalid oauth state")
	}
	oauthToken, err := a.GetToken(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	token := oauthToken.(*oauth2.Token)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.UserInfoURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	req.Header.Add("Authorization", "token "+token.AccessToken)

	response, err := a.httClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	user := new(GithubOauthUserData)

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(user); err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return convertGithubUser(user, authTypeGithub), nil
}

func convertGithubUser(githubUser *GithubOauthUserData, authTypeYandexId string) *user.UserAgr {
	userId := uuid.New().String()

	u := user.User{
		Id: userId,
		ExternalId: sql.NullString{
			String: strconv.Itoa(githubUser.ID),
			Valid:  true,
		},
		ExternalAuthType: sql.NullString{
			String: authTypeYandexId,
			Valid:  true,
		},
		Login: sql.NullString{
			String: githubUser.Login,
			Valid:  true,
		},
		FirstName:  sql.NullString{},
		LastName:   sql.NullString{},
		MiddleName: sql.NullString{},
		IsBanned:   sql.NullBool{},
	}

	emails := make([]user.Email, 0)
	emails = append(emails, user.Email{
		Id: uuid.New().String(),
		UserId: sql.NullString{
			String: userId,
			Valid:  true,
		},
		Address: sql.NullString{
			String: githubUser.Email,
			Valid:  true,
		},
	})

	return &user.UserAgr{
		User:   u,
		Emails: emails,
	}
}

package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/hardstylez72/bbckend/internal/store"
	"net/http"
	"strconv"
	"time"
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
	Email             string `json:"email,omitempty"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}


func (a *authController) getGithubUserInfo(ctx context.Context, state string, code string) (*GithubOauthUserData, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := a.githubOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.githubGetUserInfoURLTmpl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	req.Header.Add("Authorization", "token " +  token.AccessToken)

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

	return user, nil
}


func (a *authController) HandleGithubLogin(w http.ResponseWriter, r *http.Request) {
	url := a.githubOauthConfig.AuthCodeURL(oauthStateString,)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}


func (a *authController) HandleGithubCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := a.getGithubUserInfo(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	authGoogleTypeId := a.externalAuthTypeMap[authTypeGithub]

	u, err := a.userStore.GetUserByExternalId(ctx, strconv.Itoa(user.ID), authGoogleTypeId)
	if err != nil {
		panic(err)
	}

	if u == nil {
		err = a.userStore.SaveUser(ctx, convertGithubUser(user, authGoogleTypeId))
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/protected", http.StatusTemporaryRedirect)

}



func convertGithubUser(user *GithubOauthUserData, authTypeYandexId string) *store.UserAgr {
	userId := uuid.New().String()

	u :=  store.User{
		Id: userId,
		ExternalId: sql.NullString{
			String: strconv.Itoa(user.ID),
			Valid:  true,
		},
		ExternalAuthType: sql.NullString{
			String: authTypeYandexId,
			Valid:  true,
		},
		Login:    sql.NullString{
			String: user.Login,
			Valid:  true,
		},
		Password: sql.NullString{},
		FirstName: sql.NullString{
		},
		LastName: sql.NullString{
		},
		MiddleName: sql.NullString{},
		IsBanned:   sql.NullBool{},
		LastActivityTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	}

	emails := make([]store.Email, 0)
	emails = append(emails, store.Email{
		Id:        uuid.New().String(),
		UserId:    sql.NullString{
			String: userId,
			Valid:  true,
		},
		Address:   sql.NullString{
			String: user.Email,
			Valid:  true,
		},
	})

	return &store.UserAgr{
		User:   u,
		Emails: emails,
		Phones: nil,
	}
}

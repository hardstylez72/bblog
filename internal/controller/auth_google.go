package controller

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"

	"fmt"

	"github.com/google/uuid"
	"github.com/hardstylez72/bbckend/internal/store"

	"net/http"

	"time"
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

func (a *authController) HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := a.googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *authController) HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	user, err := a.getGoogleUserInfo(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	authGoogleTypeId := a.externalAuthTypeMap[authTypeGoogle]

	u, err := a.userStore.GetUserByExternalId(ctx, user.ID, authGoogleTypeId)
	if err != nil {

	}
	if u == nil {
		err = a.userStore.SaveUser(ctx, convertGoogleUser(user, authGoogleTypeId))
	}

	http.Redirect(w, r, "/protected", http.StatusTemporaryRedirect)
}

func convertGoogleUser(user *GoogleOauthUserData, authTypeGoogleId string) *store.UserAgr {
	userId := uuid.New().String()

	u :=  store.User{
		Id: userId,
		ExternalId: sql.NullString{
			String: user.ID,
			Valid:  true,
		},
		ExternalAuthType: sql.NullString{
			String: authTypeGoogleId,
			Valid:  true,
		},
		Login:    sql.NullString{},
		Password: sql.NullString{},
		FirstName: sql.NullString{
			String: user.GivenName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: user.FamilyName,
			Valid:  true,
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


func (a *authController) getGoogleUserInfo(ctx context.Context, state string, code string) (*GoogleOauthUserData, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := a.googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	buf := bytes.NewBuffer([]byte{})
	err = a.googleGetUserInfoURLTmpl.Execute(buf, token.AccessToken)
	accessTokenValidationURL := buf.String()

	response, err := http.Get(accessTokenValidationURL)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	user := new(GoogleOauthUserData)

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(user); err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return user, nil
}



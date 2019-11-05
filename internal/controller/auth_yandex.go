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

type YandexOauthUserData struct {
	FirstName        string      `json:"first_name"`
	LastName         string      `json:"last_name"`
	DisplayName      string      `json:"display_name"`
	Emails           []string    `json:"emails"`
	DefaultAvatarID  string      `json:"default_avatar_id"`
	DefaultEmail     string      `json:"default_email"`
	RealName         string      `json:"real_name"`
	IsAvatarEmpty    bool        `json:"is_avatar_empty"`
	Birthday         interface{} `json:"birthday"`
	ClientID         string      `json:"client_id"`
	OpenidIdentities []string    `json:"openid_identities"`
	Login            string      `json:"login"`
	Sex              string      `json:"sex"`
	ID               string      `json:"id"`
}


func (a *authController) getYandexUserInfo(ctx context.Context, state string, code string) (*YandexOauthUserData, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := a.yandexOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	buf := bytes.NewBuffer([]byte{})
	err = a.yandexGetUserInfoURLTmpl.Execute(buf, token.AccessToken)
	yandexGetUserInfoURL := buf.String()


	response, err := http.Get(yandexGetUserInfoURL)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	user := new(YandexOauthUserData)

	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(user); err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return user, nil
}


func (a *authController) HandleYandexLogin(w http.ResponseWriter, r *http.Request) {
	url := a.yandexOauthConfig.AuthCodeURL(oauthStateString,)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}


// https://yandex.ru/dev/passport/doc/dg/reference/response-docpage/#response__norights_3
// https://login.yandex.ru/info?format=json&with_openid_identity=1&oauth_token=AgAAAAAu8JN2AAX1EBLIpiy_5E66r4wym7WQZhU

func (a *authController) HandleYandexCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := a.getYandexUserInfo(ctx, r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	authGoogleTypeId := a.externalAuthTypeMap[authTypeYandex]

	u, err := a.userStore.GetUserByExternalId(ctx, user.ID, authGoogleTypeId)
	if err != nil {
		panic(err)
	}

	if u == nil {
		err = a.userStore.SaveUser(ctx, convertYandexUser(user, authGoogleTypeId))
		if err != nil {
			panic(err)
		}
	}

	http.Redirect(w, r, "/protected", http.StatusTemporaryRedirect)

}



func convertYandexUser(user *YandexOauthUserData, authTypeYandexId string) *store.UserAgr {
	userId := uuid.New().String()

	u :=  store.User{
		Id: userId,
		ExternalId: sql.NullString{
			String: user.ID,
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
			String: user.FirstName,
			Valid:  true,
		},
		LastName: sql.NullString{
			String: user.LastName,
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

	for _, email := range user.Emails {
		emails = append(emails, store.Email{
			Id:        uuid.New().String(),
			UserId:    sql.NullString{
				String: userId,
				Valid:  true,
			},
			Address:   sql.NullString{
				String: email,
				Valid:  true,
			},
		})
	}


	return &store.UserAgr{
		User:   u,
		Emails: emails,
		Phones: nil,
	}
}

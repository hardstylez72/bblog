package auth

import (
	"context"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"net/http"
)

type userCtx struct{}

func GuestSession(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenName := viper.GetString("oauth.sessionCookie.name")
		_, err := r.Cookie(tokenName)

		if err == http.ErrNoCookie {
			sessionId := uuid.New().String()
			token, err := GenerateTokenForGuest(sessionId)
			if err != nil {
				// todo: !!!
			}

			SetSessionCookie(w, token, SessionCookieConfig{
				Name:   viper.GetString("oauth.sessionCookie.name"),
				Domain: viper.GetString("oauth.sessionCookie.domain"),
				Path:   viper.GetString("oauth.sessionCookie.path"),
				MaxAge: viper.GetInt("oauth.sessionCookie.maxAge"),
				Secure: viper.GetBool("oauth.sessionCookie.secure"),
			})

			ctx = context.WithValue(ctx, userCtx{}, User{
				SessionId:    sessionId,
				IsAuthorized: false,
			})
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		// todo: add ban check
		w.Header().Set("Access-Control-Expose-Headers", "user_id")
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func GetUser(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenName := viper.GetString("oauth.sessionCookie.name")
		sessionCookie, err := r.Cookie(tokenName)
		if err == nil {
			user, _, errN := ExtractUserAndTokenFromStringToken(sessionCookie.Value)
			if errN == nil {
				ctx = context.WithValue(ctx, userCtx{}, user)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
		}

		// todo: add ban check
		w.Header().Set("Access-Control-Expose-Headers", tokenName)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

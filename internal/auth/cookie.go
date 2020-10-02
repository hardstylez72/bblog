package auth

import "net/http"

func InjectUserIdFromCookies(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cookie, err := r.Cookie("user_id")
		if err == nil {
			w.Header().Set("user_id", cookie.Value)
		}
		// todo: add ban check
		w.Header().Set("Access-Control-Expose-Headers", "user_id")

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

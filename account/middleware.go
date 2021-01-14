package account

import (
	"fmt"
	"net/http"
)

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// generate token
		tokenString, err := createToken()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tokenString)

		// get token and check
		err = validateToken(parseToken(r.Header.Get("X-Token-Gen")))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

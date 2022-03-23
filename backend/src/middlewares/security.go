package middlewares

import (
	"github.com/qndaa/backend/src/token"
	"net/http"
	"regexp"
	"strings"
)

type Methods []string

func (m *Methods) Contains(method string) bool {
	for _, value := range *m {
		if method == value {
			return true
		}
	}
	return false
}

var authRouts = make(map[string]Methods)

func init() {
	authRouts["/api/book"] = []string{
		"POST",
	}
	authRouts["/api/book/[1-9]+"] = []string{
		"GET",
	}
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for pattern, methods := range authRouts {
			b, _ := regexp.MatchString(pattern, r.RequestURI)
			if b == true && methods.Contains(r.Method) == true {
				maker := token.NewJWTMaker("secret")
				bearer := r.Header["Authorization"]
				if len(bearer) == 0 {
					w.WriteHeader(403)
					return

				}

				split := strings.Split(bearer[0], " ")
				verifyToken, err := maker.VerifyToken(split[1])
				if err != nil {
					w.WriteHeader(403)
					return
				}

				err = verifyToken.Valid()

				if err != nil {
					w.WriteHeader(403)
					w.Write([]byte(err.Error()))
					return
				}

				next.ServeHTTP(w, r)
				return

			}
		}
		next.ServeHTTP(w, r)
	})
}

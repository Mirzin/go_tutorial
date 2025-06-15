package middleware

import (
	"errors"
	"net/http"

	"github.com/Mirzin/go_tutorial/api"
	"github.com/Mirzin/go_tutorial/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrorUnauthorizedError = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(ErrorUnauthorizedError)
			api.RequestErrorHandler(w, ErrorUnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := (*database).GetUserLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(ErrorUnauthorizedError)
			api.RequestErrorHandler(w, ErrorUnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}

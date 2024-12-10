package middleware

import (
	"errors"
	"github/abdqadr1/goapi/api"
	"net/http"

	"github.com/abdqadr1/goapi/api"
	"github.com/abdqadr1/goapi/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()

		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails != nil || (token != (*loginDetails).AuthToken) {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})

}

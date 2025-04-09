package middleware

import (
	"errors"
	"net/http"

	"sudoku-server/api"

	log "github.com/sirupsen/logrus"
)

var UnauthorizedError = errors.New("Invalid token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		// var database *tools.DatabaseInterface
		// database, err := tools.NewDatabase()
		// if err != nil {
		// 	log.Error("Failed to open to database: ", err)
		// 	api.InternalErrorHandler(w)
		// 	return
		// }

		// TODO: add a proper token management system, storing tokens in the database
		if token != "admin" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
		// var loginDetails *tools.LoginDetails
		// loginDetails = (*database).GetUserLoginDetails(username)
		//
		// if loginDetails == nil || token != loginDetails.Token {
		// 	log.Error(UnauthorizedError)
		// 	api.RequestErrorHandler(w, UnauthorizedError)
		// 	return
		// }

		next.ServeHTTP(w, r)
	})
}

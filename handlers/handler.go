package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)



type ErrorResponse struct {
    Code int
    Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
    response := ErrorResponse{
        Code: code,
        Message: message,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)

    json.NewEncoder(w).Encode(response)
}

var (
    RequestErrorHandler = func(w http.ResponseWriter, err error) {
        writeError(w, err.Error(), http.StatusBadRequest)
    }
    InternalErrorHandler = func(w http.ResponseWriter) {
        writeError(w, "An Unexpected Internal Error Occured.", http.StatusInternalServerError)
    }
)



func Handler(r *chi.Mux) {

    creds := map[string]string{
        "admin": "password",
    }

    r.Use(chimiddle.StripSlashes)
    r.Use(chimiddle.Logger)
    r.Use(chimiddle.BasicAuth("Restricted", creds))


    r.Get("/solve", SolveHandler)
    r.Get("/valid", ValidHandler)
    r.Get("/gen", GenHandler)
}

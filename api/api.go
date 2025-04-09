package api

import (
    "encoding/json"
    "net/http"
)


type SolveParams struct {
    Input string 
}
type SolveResponse struct {
    Code int
    Solution []string
}

type ValidParams struct {
    Input string
}
type ValidResponse struct {
    Code int
    Valid bool
}


type Error struct {
    Code int
    Message string
}


func writeError(w http.ResponseWriter, message string, code int) {
    response := Error{
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





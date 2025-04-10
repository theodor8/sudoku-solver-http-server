package handlers

import (
	"encoding/json"
	"net/http"
	"sudoku-server/api"
	"sudoku-server/internal/tools"

	log "github.com/sirupsen/logrus"
)

type DbResponse struct {
    Code int
    Data []tools.SolutionData
}

func DbHandler(w http.ResponseWriter, r *http.Request) {

    database, err := tools.NewDatabase()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    solutionDatas := (*database).GetAllSolutionData()

    w.Header().Set("Content-Type", "application/json")
    response := DbResponse{
        Code: http.StatusOK,
        Data: solutionDatas,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("failed to encode response: ", err)
        api.InternalErrorHandler(w)
    }
}

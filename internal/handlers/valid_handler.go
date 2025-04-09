package handlers

import (
	"net/http"
	"sudoku-server/api"
	"sudoku-server/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)



func ValidHandler(w http.ResponseWriter, r *http.Request) {
    params := api.SolveParams{}
    decoder := schema.NewDecoder()
    
    err := decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error("Failed to decode parameters: ", err)
        api.InternalErrorHandler(w)
        return
    }

    database, err := tools.NewDatabase()
    if err != nil {
        log.Error(err)
        api.InternalErrorHandler(w)
        return
    }

    var solutions []string
    solutionData := (*database).GetSolutionData(params.Input)
    if solutionData != nil {
        solutions = solutionData.Solutions
        log.Info("Solution found in database")
    } else {
        solutions, err = solver.Solve(params.Input)
        if err != nil {
            log.Error("Solve failed: ", err)
            api.InternalErrorHandler(w)
            return
        }
        (*database).StoreSolutionData(&tools.SolutionData{
            Input: params.Input,
            Solutions: solutions,
        })
        log.Info("Solution computed and stored in database")
    }

    w.Header().Set("Content-Type", "application/json")
    response := api.SolveResponse{
        Code: http.StatusOK,
        Solution: solutions,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("Failed to encode response: ", err)
        api.InternalErrorHandler(w)
    }
}

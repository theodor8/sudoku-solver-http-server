package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"sudoku-server/api"
	"sudoku-server/internal/solver"
	"sudoku-server/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)


func SolveHandler(w http.ResponseWriter, r *http.Request) {
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

    var solutions string
    solutionDetails := (*database).GetSolutionDetails(params.Input)
    if solutionDetails != nil {
        solutions = solutionDetails.Solutions
        log.Info("Solutions found in database: ", solutions)
    } else {
        solutionsSlice, err := solver.Solve(params.Input)
        if err != nil {
            log.Error("Solve failed: ", err)
            api.InternalErrorHandler(w)
            return
        }
        solutions = strings.Join(solutionsSlice, " ")
        (*database).StoreSolutionDetails(&tools.SolutionDetails{
            Input: params.Input,
            Solutions: solutions,
        })
        log.Info("Solutions computed and stored in database: ", solutions)
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

package handlers

import (
	"encoding/json"
	"net/http"
	"sudoku-server/api"
	"sudoku-server/internal/solver"
	"sudoku-server/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)


type SolveParams struct {
    Input string 
}
type SolveResponse struct {
    Code int
    Solution []string
}


func SolveHandler(w http.ResponseWriter, r *http.Request) {
    params := SolveParams{}
    decoder := schema.NewDecoder()
    
    err := decoder.Decode(&params, r.URL.Query())

    if err != nil {
        log.Error("failed to decode parameters: ", err)
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
        log.Info("solution found in database")
    } else {
        solutions, err = solver.Solve(params.Input)
        if err != nil {
            log.Error("solve failed: ", err)
            api.RequestErrorHandler(w, err)
            return
        }
        (*database).StoreSolutionData(&tools.SolutionData{
            Input: params.Input,
            Solutions: solutions,
        })
        log.Info("solution computed and stored in database")
    }

    w.Header().Set("Content-Type", "application/json")
    response := SolveResponse{
        Code: http.StatusOK,
        Solution: solutions,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("failed to encode response: ", err)
        api.InternalErrorHandler(w)
    }

}

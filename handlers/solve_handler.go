package handlers

import (
	"encoding/json"
	"net/http"
	"sudoku-server/sudoku"
        "sudoku-server/database"

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
        InternalErrorHandler(w)
        return
    }

    db, err := database.NewDatabase()
    if err != nil {
        log.Error(err)
        InternalErrorHandler(w)
        return
    }

    var solutions []string
    solutionData := (*db).GetSolutionData(params.Input)
    if solutionData != nil {
        solutions = solutionData.Solutions
        log.Info("solution found in database")
    } else {
        solutions, err = sudoku.Solve(params.Input)
        if err != nil {
            log.Error("solve failed: ", err)
            RequestErrorHandler(w, err)
            return
        }
        (*db).StoreSolutionData(&database.SolutionData{
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
        InternalErrorHandler(w)
    }

}

package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"sudoku-server/sudoku"
	"time"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)



type GenParams struct {
    Unknowns uint8
}
type GenResponse struct {
    Code int
    Grid string
}

func GenHandler(w http.ResponseWriter, r *http.Request) {

    params := GenParams{}

    if _, ok := r.URL.Query()["unknowns"]; !ok {
        params.Unknowns = 40
    } else {
        decoder := schema.NewDecoder()
        err := decoder.Decode(&params, r.URL.Query())

        if err != nil {
            log.Error("failed to decode parameters: ", err)
            InternalErrorHandler(w)
            return
        }
    }

    rand := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))

    log.Info("generating grid with ", params.Unknowns, " unknowns")
    grid := sudoku.Generate(rand, params.Unknowns)


    w.Header().Set("Content-Type", "application/json")
    response := GenResponse{
        Code: http.StatusOK,
        Grid: grid,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("failed to encode response: ", err)
        InternalErrorHandler(w)
    }
}

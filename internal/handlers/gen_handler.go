package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"sudoku-server/api"
	"sudoku-server/internal/solver"
	"time"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)




func GenHandler(w http.ResponseWriter, r *http.Request) {

    params := api.GenParams{}

    if _, ok := r.URL.Query()["unknowns"]; !ok {
        params.Unknowns = 40
    } else {
        decoder := schema.NewDecoder()
        err := decoder.Decode(&params, r.URL.Query())

        if err != nil {
            log.Error("failed to decode parameters: ", err)
            api.InternalErrorHandler(w)
            return
        }
    }

    rand := rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
    grid := solver.Generate(rand, params.Unknowns)

    w.Header().Set("Content-Type", "application/json")
    response := api.GenResponse{
        Code: http.StatusOK,
        Grid: grid,
    }
    if err := json.NewEncoder(w).Encode(response); err != nil {
        log.Error("failed to encode response: ", err)
        api.InternalErrorHandler(w)
    }
}

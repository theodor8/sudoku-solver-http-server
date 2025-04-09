package tools

import (
    log "github.com/sirupsen/logrus"
)


type SolutionDetails struct {
    Input string
    Solutions string // TODO: make this a slice
}

type DatabaseInterface interface {
    GetSolutionDetails(input string) *SolutionDetails
    StoreSolutionDetails(solutions *SolutionDetails) error
    SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
    var database DatabaseInterface = &sqliteDB{}

    if err := database.SetupDatabase(); err != nil {
        log.Error("Failed to setup database: ", err)
        return nil, err
    }

    return &database, nil
}


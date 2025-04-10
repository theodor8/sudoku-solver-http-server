package tools

import (
    log "github.com/sirupsen/logrus"
)


type Solutions []string

type SolutionData struct {
    Input string
    Solutions Solutions
}

type DatabaseInterface interface {
    SetupDatabase() error
    GetSolutionData(input string) *SolutionData
    StoreSolutionData(solutions *SolutionData) error
    GetAllSolutionData() []SolutionData
}

func NewDatabase() (*DatabaseInterface, error) {
    var database DatabaseInterface = &sqliteDB{}

    if err := database.SetupDatabase(); err != nil {
        log.Error("Failed to setup database: ", err)
        return nil, err
    }

    return &database, nil
}


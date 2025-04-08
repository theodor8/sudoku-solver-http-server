package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var db *gorm.DB


type solutionsType []string

func (s *solutionsType) Scan(value any) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("value value cannot cast to []byte")
    }
    return json.Unmarshal(bytes, s)
}

func (s solutionsType) Value() (driver.Value, error) {
    return json.Marshal(s)
}

type inputSolutions struct {
    gorm.Model
    Input  string
    Solutions solutionsType `gorm:"type:json"`
}


func Init() {
    var err error
    db, err = gorm.Open(sqlite.Open("solutions.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&inputSolutions{})
}

func StoreSolutions(input string, solutions []string) {
    db.Create(&inputSolutions{Input: input, Solutions: solutions})
}

func FindSolutions(input string) []string {
    var solution inputSolutions
    result := db.Limit(1).Find(&solution, "input = ?", input)
    if result.RowsAffected == 0 {
        return nil
    }
    return solution.Solutions
}

func AllSolutions() map[string][]string {
    var inputSolutions []inputSolutions
    db.Find(&inputSolutions)
    solutionsMap := make(map[string][]string)
    for _, solution := range inputSolutions {
        solutionsMap[solution.Input] = solution.Solutions
    }
    return solutionsMap
}


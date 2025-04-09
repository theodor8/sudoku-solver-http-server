package tools

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDB struct {
    db *gorm.DB
}


func (db *sqliteDB) GetSolutionDetails(input string) *SolutionDetails {
    var solution SolutionDetails
    // result := db.db.First(&solution, "input = ?", input)
    result := db.db.Limit(1).Find(&solution, "input = ?", input)
    if result.RowsAffected == 0 {
        return nil
    }
    return &solution
}

func (db *sqliteDB) StoreSolutionDetails(solutions *SolutionDetails) error {
    result := db.db.Create(solutions)
    if result.Error != nil {
        return result.Error
    }
    return nil
}


func (db *sqliteDB) SetupDatabase() error {
    var err error
    db.db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to open database.")
    }
    db.db.AutoMigrate(&SolutionDetails{})
    return nil
}



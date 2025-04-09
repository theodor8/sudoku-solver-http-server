package tools

import (
    "encoding/json"
    "database/sql/driver"
    "errors"

    "gorm.io/driver/sqlite"

    "gorm.io/gorm"
)

func (s *Solutions) Scan(value any) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New("value value cannot cast to []byte")
    }
    return json.Unmarshal(bytes, s)
}

func (s Solutions) Value() (driver.Value, error) {
    return json.Marshal(s)
}


type sqliteDB struct {
    db *gorm.DB
}


func (db *sqliteDB) GetSolutionData(input string) *SolutionData {
    var solution SolutionData
    result := db.db.Limit(1).Find(&solution, "input = ?", input)
    if result.RowsAffected == 0 {
        return nil
    }
    return &solution
}

func (db *sqliteDB) StoreSolutionData(solutions *SolutionData) error {
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
    db.db.AutoMigrate(&SolutionData{})
    return nil
}



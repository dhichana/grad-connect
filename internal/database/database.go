package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbPath string) (*gorm.DB, error) {
    var err error

    DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = DB.AutoMigrate()
    if err != nil {
        return nil, err
    }

    return DB, nil
}
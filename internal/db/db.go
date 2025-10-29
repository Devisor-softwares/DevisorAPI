package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "devisor/internal/models"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&models.User{}, &models.Server{})

    return db, nil
}

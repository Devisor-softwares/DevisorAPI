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

    // Here, we essentially migrate our models.
    // We take the models.User, which simply get's
    // the user model from the modules module.
    // We do the same for server, which are both
    // not implemented yet.
    db.AutoMigrate(&models.User{}, &models.Server{})

    return db, nil
}

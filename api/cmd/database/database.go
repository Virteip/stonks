package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewPostgresDatabase creates a new Database instance using PostgreSQL/CockroachDB
func NewPostgresDatabase(connectionString string) (Database, error) {
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	db, err := gorm.Open(postgres.Open(connectionString), gormConfig)
	if err != nil {
		return nil, err
	}

	// Wrap the GORM DB with our adapter
	return NewGormAdapter(db), nil
}

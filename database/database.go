package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresService struct {
	DB *gorm.DB
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password= dbname=postgres port=5433"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewPostgresService(db *gorm.DB) *PostgresService {
	return &PostgresService{
		DB: db,
	}
}

package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jmoiron/sqlx"
	"os"

	_ "github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Database struct {
	gClient *gorm.DB
}

func NewDatabase() (*Database, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	dbConn, err := gorm.Open("postgres", connString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to database: %w", err)
	}

	return &Database{
		gClient: dbConn,
	}, nil
}

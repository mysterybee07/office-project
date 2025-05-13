package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mysterybee07/office-project-setup/internal/config"
)

type Service interface {
	Close() error
}

type service struct {
	db *sql.DB
}

var dbInstance *service

var once sync.Once

func NewDatabase(cfg *config.DatabaseConfig) Service {
	once.Do(func() {

		if err := MigrateFromConfig(cfg); err != nil {
			log.Fatalf("Database migration failed: %v", err)
		}

		connStr := cfg.GetDatabaseURL()
		db, err := sql.Open("pgx", connStr)
		if err != nil {
			log.Fatal(err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Could not establish database connection: %v", err)
		}

		dbInstance = &service{db: db}
	})
	return dbInstance
}

func (s *service) Close() error {
	fmt.Println("Disconnected from database...")
	return s.db.Close()
}

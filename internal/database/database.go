package database

import (
	"log"
	"strings"
	"sync"

	"github.com/mysterybee07/office-project-setup/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
	GetDB() *gorm.DB
	Close() error
}

type service struct {
	db *gorm.DB
}

var dbInstance *service
var once sync.Once

func NewDatabase(cfg *config.DatabaseConfig) Service {
	once.Do(func() {
		if err := MigrateFromConfig(cfg); err != nil {
			log.Fatalf("Database migration failed: %v", err)
		}

		dsn := cfg.GetDSN()

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}

		// Set the search path after connection is established
		if cfg.Schema != "" {
			schemas := strings.Split(cfg.Schema, ",")
			searchPath := "SET search_path = " + strings.Join(schemas, ", ")

			if err := db.Exec(searchPath).Error; err != nil {
				log.Fatalf("Failed to set search path: %v", err)
			}
		}

		// Auto migrate models
		// if err := db.AutoMigrate(&model.User{}, &model.Auth{}); err != nil {
		// 	log.Fatalf("Failed to migrate database models: %v", err)
		// }

		// log.Println("Database migrations completed successfully")

		dbInstance = &service{db: db}
		log.Println("Database connected successfully")
	})
	return dbInstance
}

func (s *service) GetDB() *gorm.DB {
	return s.db
}

func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	log.Println("Disconnected from database...")
	return sqlDB.Close()
}

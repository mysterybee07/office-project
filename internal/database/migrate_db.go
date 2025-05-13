package database

import (
	"fmt"
	"log"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/mysterybee07/office-project-setup/internal/config"
)

// MigrateDB performs database migrations from the migrations directory
func MigrateDB(databaseURL string) error {
	m, err := migrate.New("file://internal/database/migrations", databaseURL)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}
	defer m.Close()

	// Handle dirty database state
	currentVersion, dirty, err := m.Version()
	if err != nil {
		// If no version is found, this is expected for a new database
		if err == migrate.ErrNilVersion {
			log.Println("No existing migration version found")
		} else {
			log.Printf("Error checking migration version: %v", err)
		}
	}

	// If database is in a dirty state, attempt to reset
	if dirty {
		log.Println("Database is in a dirty state. Attempting to reset.")

		// First, try to roll down to reset the dirty state
		if err := m.Down(); err != nil {
			// If down migration fails, force the version
			log.Printf("Down migration failed: %v. Attempting to force reset.", err)
			if forceErr := m.Force(int(currentVersion)); forceErr != nil {
				return fmt.Errorf("error forcing migration version: %w", forceErr)
			}
		}
	}

	// Run up migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No new migrations to apply")
			return nil
		}

		// Handle specific migration errors
		if strings.Contains(err.Error(), "already exists") {
			log.Println("Some database objects already exist. Continuing...")
			return nil
		}

		return fmt.Errorf("error running migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

// ForceCleanMigration provides a method to manually clean up migration state
func ForceCleanMigration(databaseURL string) error {
	m, err := migrate.New("file://internal/database/migrations", databaseURL)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}
	defer m.Close()

	// Force version to 0 (clean slate)
	if err := m.Force(0); err != nil {
		return fmt.Errorf("error forcing migration to version 0: %w", err)
	}

	// Run all migrations from scratch
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply after force reset")
			return nil
		}
		return fmt.Errorf("error running migrations after force reset: %w", err)
	}

	log.Println("Database migrations forcefully reset and reapplied")
	return nil
}

// RollbackMigration rolls back the last applied migration
func RollbackMigration(databaseURL string) error {
	m, err := migrate.New("file://internal/database/migrations", databaseURL)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Steps(-1); err != nil {
		return fmt.Errorf("error rolling back migration: %w", err)
	}

	log.Println("Last migration rolled back successfully")
	return nil
}

// MigrateFromConfig performs database migrations using the provided database configuration
func MigrateFromConfig(cfg *config.DatabaseConfig) error {
	return MigrateDB(cfg.GetDatabaseURL())
}

// RollbackFromConfig rolls back the last migration using the provided database configuration
func RollbackFromConfig(cfg *config.DatabaseConfig) error {
	return RollbackMigration(cfg.GetDatabaseURL())
}

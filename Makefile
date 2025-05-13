# Usage: make migrate-create name=create_table_name

# Load environment variables from .env file
include .env
export $(shell sed 's/=.*//' .env)

# Construct DB URL (with schema)
DB_URL=postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable&search_path=$(DB_SCHEMA)

# Migration directory
MIGRATIONS_DIR=internal/database/migrations

# Create a new migration file
migrate-create:
ifndef name
	$(error You must specify a migration name using `make migrate-create name=your_migration_name`)
endif
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

## Run all migrations
migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

## Rollback last migration
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

# Roll back N steps
# Usage: make migrate-down-n n=3
migrate-down-n:
ifndef n
	$(error You must specify the number of steps using `make migrate-down-n n=number`)
endif
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down $(n)


## Show current migration version
migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version

# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Live reload using Air
watch:
	@if command -v air > /dev/null; then \
		echo "Starting Air..."; \
		air; \
	else \
		read -p "Go's 'air' is not installed. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
		else \
			echo "You chose not to install Air. Exiting..."; \
			exit 1; \
		fi; \
	fi

.PHONY: all build run watch migrate-up migrate-down migrate-version migrate-create

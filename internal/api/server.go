package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mysterybee07/office-project-setup/domain"
	"github.com/mysterybee07/office-project-setup/infrastructure"
	"github.com/mysterybee07/office-project-setup/internal/config"
	"github.com/mysterybee07/office-project-setup/internal/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// NewServer creates and returns an HTTP server with all dependencies wired using fx
func NewServer() *http.Server {
	// Load configuration first, as we need it before fx
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	// Create a channel to receive the HTTP handler from fx
	var handler http.Handler

	// Create and start the fx application
	app := fx.New(
		// Provide the database connection
		fx.Provide(
			func() *config.DatabaseConfig {
				return cfg
			},
			database.NewDatabase,
			func(db database.Service) *gorm.DB {
				return db.GetDB()
			},
			// Provide router implementation
			infrastructure.NewRouter,
		),

		// Include domain modules
		domain.Module,

		// Extract the configured router/handler
		fx.Invoke(func(router *infrastructure.Router) {
			handler = router.Handler()
		}),
	)

	// Start the fx application in the background
	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatalf("failed to start fx application: %v", err)
	}

	// Create the HTTP server with the handler from fx
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

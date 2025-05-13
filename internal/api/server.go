package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mysterybee07/office-project-setup/internal/api/route"
	"github.com/mysterybee07/office-project-setup/internal/config"
	"github.com/mysterybee07/office-project-setup/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	srv := &Server{
		port: port,
		db:   database.NewDatabase(cfg),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", srv.port),
		Handler:      route.SetupRouter(srv.db.GetDB()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return server
}

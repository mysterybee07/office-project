package main

import (
	"fmt"
	"net/http"

	"github.com/mysterybee07/office-project-setup/internal/api"
)

// @title       Office Project API
// @version     1.0
// @description Learning swagger api.
// @host        localhost:8080
// @BasePath    /api/v1
func main() {
	server := api.NewServer()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error:%s", err))
	}

}

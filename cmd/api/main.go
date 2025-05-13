package main

import (
	"fmt"
	"net/http"

	"github.com/mysterybee07/office-project-setup/internal/api"
)

func main() {
	server := api.NewServer()

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error:%s", err))
	}

}

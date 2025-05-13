package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/internal/api/handler"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello World")
	})

	router.POST("/login", handler.Login(db))
	return router
}

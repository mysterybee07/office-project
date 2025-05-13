package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/internal/api/handler"

	// "github.com/mysterybee07/office-project-setup/internal/middleware"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello World")
	})

	// Public routes
	router.POST("/login", handler.Login(db))

	// // protected route group
	// authorized := router.Group("/")
	// authorized.Use(middleware.JWTAuthMiddleware())
	// {

	// }

	return router
}

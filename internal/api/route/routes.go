package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/internal/api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) http.Handler {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello World")
	})

	// Public routes
	router.POST("/auth/login", handler.Login(db))

	// // protected route group
	// authorized := router.Group("/")
	// authorized.Use(middleware.JWTAuthMiddleware())
	// {

	// }

	return router
}

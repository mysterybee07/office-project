package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/infrastructure"
)

type UserRoute struct {
	router     *infrastructure.Router
	controller *UserController
}

func NewUserRoute(router *infrastructure.Router, controller *UserController) *UserRoute {
	return &UserRoute{
		router:     router,
		controller: controller,
	}
}

func RegisterRoute(r *UserRoute) {
	r.router.POST("/user", r.controller.CreateUser)
	r.router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello world")
	})
}

package infrastructure

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()
	return &Router{
		router,
	}
}

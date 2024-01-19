package api

import (
	"back/internal/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer(handler http.Handler) *http.Server {
	return &http.Server{
		MaxHeaderBytes: 2048,
		Handler:        handler,
		Addr:           fmt.Sprintf(":%d", config.JSON.Port),
	}
}

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", handler)

	return router
}

func handler(ctx *gin.Context) {
	ctx.JSON(200, "hello")
}

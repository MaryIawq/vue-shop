package api

import (
	"fmt"
	"net/http"

	authAPI "back/internal/api/handlers/auth"
	paymentAPI "back/internal/api/handlers/payment"
	productsAPI "back/internal/api/handlers/products"
	"back/internal/api/middleware"
	"back/internal/config"

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
	router.Use(middleware.CORS)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authAPI.Login)
		auth.POST("/register", authAPI.Register)

		auth.GET("/logout", authAPI.Logout)
		auth.GET("/getUser", authAPI.GetUser)
	}

	admin := router.Group("/admin")
	{
		admin.GET("/admin", productsAPI.GetAll)
	}

	user := router.Group("/user")
	{
		user.GET("/user", productsAPI.GetAll)
	}

	products := router.Group("/products")
	{
		products.GET("/products/getAll", productsAPI.GetAll)
	}

	payment := router.Group("/payment", middleware.Authentication)
	{
		payment.POST("/getToken", paymentAPI.GetToken)
	}

	return router
}

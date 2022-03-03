package router

import (
	"github.com/labstack/echo/v4"
	"rest-api-go-lang/internal/app/domain"
	"rest-api-go-lang/internal/app/web/handlers"
)

func Configure(server *echo.Echo, services *domain.Services) {

	// products
	server.GET("/products", handlers.ListProducts(services))
	server.POST("/products", handlers.CreateProduct(services))
}

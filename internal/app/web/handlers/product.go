package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rest-api-go-lang/internal/app/domain"
	"rest-api-go-lang/internal/app/web/data/mapper"
	"rest-api-go-lang/internal/app/web/data/presenter"
	"rest-api-go-lang/internal/app/web/data/request"
)

func ListProducts(services *domain.Services) func(c echo.Context) error {
	return func(c echo.Context) error {
		products, err := services.ProductService.ListProducts(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, presenter.ListProducts(products))
	}
}

func CreateProduct(services *domain.Services) func(c echo.Context) error {
	return func(c echo.Context) error {
		productRequest := new(request.CreateProduct)
		if err := c.Bind(productRequest); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(productRequest); err != nil {
			return err
		}

		product := mapper.CreateProductRequestToProduct(productRequest)
		err := services.ProductService.CreateProduct(c.Request().Context(), product)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, presenter.CreateProduct(product))
	}
}

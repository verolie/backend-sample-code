package server

import (
	"net/http"

	"github.com/code-sample/app/login"
	"github.com/code-sample/app/product"

	// "github.com/code-sample/migrate"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	e := gin.Default()

	// Initialize migrations
	// migrate.Init()

	// Seed initial data for Users
	// seed.SeedUsers()

	// Register routes
	registerServer(e)

	// Start the server
	e.Run()
}


func registerServer(e *gin.Engine) {
	// Login
	e.POST("/login", LoginHandler)

	// Stock product
	e.GET("/stock-product", StockProductHandler)
	e.POST("/stock-product", StockProductHandler)
	e.PUT("/stock-product/:id", usersHandlerParamStockProduct)
	e.DELETE("/stock-product/:id", usersHandlerParamStockProduct)
}

func LoginHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPost:
		login.PostLogin(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}


func StockProductHandler(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodGet:
		product.GetProduct(c)
	case http.MethodPost:
		product.CreateProduct(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}


func usersHandlerParamStockProduct(c *gin.Context) {
	switch c.Request.Method {
	case http.MethodPut:
		product.PutProduct(c)
	case http.MethodDelete:
		product.DeleteProduct(c)
	default:
		c.String(http.StatusMethodNotAllowed, "Method not allowed")
	}
}
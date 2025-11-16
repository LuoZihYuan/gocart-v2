package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/gocart-v2/cart-service/docs"
	"github.com/gocart-v2/cart-service/internal/handler"
	"github.com/gocart-v2/cart-service/internal/repository"
	"github.com/gocart-v2/cart-service/internal/router"
	"github.com/gocart-v2/cart-service/internal/service"
)

// @title E-commerce API
// @version 1.0.0
// @description API for managing products, shopping carts, warehouse operations, and credit card processing
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key
// @securityDefinitions.bearer BearerAuth
// @tag.name Shopping Cart
// @tag.description Shopping cart operations
func main() {

	rh := handler.NewRootHandler()

	cr := repository.NewCartRepository()
	cs := service.NewCartService(cr)
	ch := handler.NewCartHandler(cs)

	e := gin.Default()
	router.SetupRoutes(e, &router.AllHandlers{
		RootHandler:    rh,
		CartHandler:    ch,
		SwaggerHandler: swaggerFiles.Handler,
	})

	log.Println("Starting server on :8081")
	if err := e.Run(":8081"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

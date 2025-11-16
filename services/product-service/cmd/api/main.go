package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"

	_ "github.com/gocart-v2/product-service/docs"
	"github.com/gocart-v2/product-service/internal/handler"
	"github.com/gocart-v2/product-service/internal/repository"
	"github.com/gocart-v2/product-service/internal/router"
	"github.com/gocart-v2/product-service/internal/service"
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
// @tag.name Product
// @tag.description Product management operations
func main() {

	rh := handler.NewRootHandler()

	pr := repository.NewProductRepository()
	ps := service.NewProductService(pr)
	ph := handler.NewProductHandler(ps)

	r := gin.Default()
	router.SetupRoutes(r, &router.AllHandlers{
		RootHandler:    rh,
		ProductHandler: ph,
		SwaggerHandler: swaggerFiles.Handler,
	})

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

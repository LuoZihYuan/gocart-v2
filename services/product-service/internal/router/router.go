package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/webdav"

	"github.com/gocart-v2/product-service/internal/handler"
)

type AllHandlers struct {
	RootHandler    *handler.RootHandler
	ProductHandler *handler.ProductHandler
	SwaggerHandler *webdav.Handler
}

func SetupRoutes(e *gin.Engine, h *AllHandlers) {
	root := e.Group("")
	{
		root.GET("/health", h.RootHandler.GetHealthStatus)
	}

	v1 := e.Group("/v1")
	{
		// Product routes
		product := v1.Group("/product")
		{
			product.GET("/:productId", h.ProductHandler.GetProduct)
			product.POST("/:productId/details", h.ProductHandler.AddProductDetails)
		}
	}

	swagger := e.Group("/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(h.SwaggerHandler))
	}
}

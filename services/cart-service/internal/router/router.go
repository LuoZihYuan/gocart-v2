package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/net/webdav"

	"github.com/gocart-v2/cart-service/internal/handler"
)

type AllHandlers struct {
	RootHandler    *handler.RootHandler
	CartHandler    *handler.CartHandler
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
		carts := v1.Group("/shopping-cart")
		{
			carts.POST("", h.CartHandler.CreateCart)
			carts.GET("/:shoppingCartId", h.CartHandler.GetCart)
			carts.POST("/:shoppingCartId/items", h.CartHandler.AddItemsToCart)
			carts.POST("/:shoppingCartId/checkout", h.CartHandler.CheckoutCart)
		}
	}

	swagger := e.Group("/swagger")
	{
		swagger.GET("/*any", ginSwagger.WrapHandler(h.SwaggerHandler))
	}
}

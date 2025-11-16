package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/gocart-v2/product-service/internal/service"
	"github.com/gocart-v2/shared/model"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetProduct handles GET /product/{productId}
// @Summary Get product by ID
// @Description Retrieve a product's details using its unique identifier
// @ID getProduct
// @Tags Product
// @Accept json
// @Produce json
// @Param productId path int true "Unique identifier for the product" minimum(1)
// @Success 200 {object} model.Product
// @Failure 404 {object} model.Error
// @Failure 500 {object} model.Error
// @Router /product/{productId} [get]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *ProductHandler) GetProduct(c *gin.Context) {
	// Parse productId from URL parameter
	productIDStr := c.Param("productId")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID < 1 {
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   "INVALID_INPUT",
			Message: "Invalid product ID",
			Details: "Product ID must be a positive integer",
		})
		return
	}

	// Get product from service
	product, err := h.service.GetProduct(productID)
	if err == service.ErrProductNotFound {
		c.JSON(http.StatusNotFound, model.Error{
			Error:   "NOT_FOUND",
			Message: "Product not found",
			Details: "No product exists with the specified ID",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, model.Error{
			Error:   "INTERNAL_ERROR",
			Message: "Internal server error",
			Details: err.Error(),
		})
		return
	}

	// Return product
	c.JSON(http.StatusOK, product)
}

// AddProductDetails handles POST /product/{productId}/details
// @Summary Add product details
// @Description Add or update detailed information for a specific product
// @ID addProductDetails
// @Tags Product
// @Accept json
// @Produce json
// @Param productId path int true "Unique identifier for the product" minimum(1)
// @Param product body model.Product true "Product details"
// @Success 204 "Product details added successfully"
// @Failure 400 {object} model.Error
// @Failure 404 {object} model.Error
// @Failure 500 {object} model.Error
// @Router /product/{productId}/details [post]
// @Security ApiKeyAuth
// @Security BearerAuth
func (h *ProductHandler) AddProductDetails(c *gin.Context) {
	// Parse productId from URL parameter
	productIDStr := c.Param("productId")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID < 1 {
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   "INVALID_INPUT",
			Message: "Invalid product ID",
			Details: "Product ID must be a positive integer",
		})
		return
	}

	// Parse request body
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   "INVALID_INPUT",
			Message: "Invalid input data",
			Details: err.Error(),
		})
		return
	}

	// Add product details through service
	if err := h.service.AddProductDetails(productID, &product); err != nil {
		if err == service.ErrProductNotFound {
			c.JSON(http.StatusNotFound, model.Error{
				Error:   "NOT_FOUND",
				Message: "Product not found",
				Details: "No product exists with the specified ID",
			})
			return
		}
		if err == service.ErrInvalidProduct || err.Error() == "product ID mismatch" {
			c.JSON(http.StatusBadRequest, model.Error{
				Error:   "INVALID_INPUT",
				Message: "Invalid input data",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, model.Error{
			Error:   "INTERNAL_ERROR",
			Message: "Internal server error",
			Details: err.Error(),
		})
		return
	}

	// Return 204 No Content on success
	c.Status(http.StatusNoContent)
}

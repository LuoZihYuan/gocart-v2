package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RootHandler struct {
}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

// GetHealthStatus godoc
// @Summary Health check endpoint
// @Description Returns HTTP 200 OK if the service is running and healthy
// @Tags Root
// @Produce plain
// @Success 200 {string} string "Service is healthy"
// @Router /health [get]
func (h *RootHandler) GetHealthStatus(c *gin.Context) {
	c.Status(http.StatusOK)
}

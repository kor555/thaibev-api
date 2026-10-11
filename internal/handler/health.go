package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpHealthHandler struct {
}

func NewHealthHandler() *HttpHealthHandler {
	return &HttpHealthHandler{}
}

func (h *HttpHealthHandler) Healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, "Healthy")
}

func (h *HttpHealthHandler) Version(c *gin.Context) {
	version := "1.2.0"
	buildDate := "2026-01-27"

	response := map[string]interface{}{
		"version":   version,
		"buildDate": buildDate,
		"status":    "ok",
	}

	c.JSON(http.StatusOK, response)
}

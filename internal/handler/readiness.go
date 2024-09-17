package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Readiness .
func Readiness(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}

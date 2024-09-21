package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Dummy .
func Dummy(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
}
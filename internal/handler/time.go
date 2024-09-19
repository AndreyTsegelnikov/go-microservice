package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Time .
func Time(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"currentTime": time.Now().Format("15:04"),
	})
}
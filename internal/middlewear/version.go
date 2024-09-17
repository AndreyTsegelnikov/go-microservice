package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/AndreyTsegelnikov/go-microservice/internal/vars"
	
)

func Version(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(vars.HeaderServiceVersion, version)
		c.Next()
	}
}
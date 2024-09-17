package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// initiates a gin Engine with the default logger and recovery middleware
	router := gin.Default()

	// sets up a GET API in route /time that returns the current time
	router.GET("/time", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"currentTime": time.Now().Format("15:04"),
		})
	})

	// Run implements a http.ListenAndServe() and takes in an optional Port number
	// The default port is :8080
	if err := router.Run(); err != nil {
		panic(err)
	}
}

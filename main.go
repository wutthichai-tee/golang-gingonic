package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	// http://localhost:8080/ping
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

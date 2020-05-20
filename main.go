package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	S "./src/api/services"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "API running"})
	})

	r.POST("/auth", S.Login)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
package main

import (
	"net/http"

	gorm "./Config"
	S "./src/api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	db := gorm.GetDB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message": "API running"})
	})

	r.POST("/login", S.Login)

	r.GET("/profile", S.GetItem)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

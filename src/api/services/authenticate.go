package services

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]
	auth, err := base64.StdEncoding.DecodeString(authHeader)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Login declined",
		})
	} else {
		if string(auth) == "rajagusri:password" {
			c.JSON(200, gin.H{
				"message": "Login Approved",
			})
		}
	}
}

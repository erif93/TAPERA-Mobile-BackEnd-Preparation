package services

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	authHeader := c.Request.Header["Authorization"][0]
	runes := []rune(authHeader)
	safeSubstring := string(runes[6:])
	fmt.Println(safeSubstring)
	auth, err := base64.StdEncoding.DecodeString(safeSubstring)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Login declined",
		})
	} else {
		if string(auth) == "rajagusri:password" {
			c.JSON(200, gin.H{
				"message": "Login Approved",
			})
		} else {
			c.JSON(200, gin.H{
				"message": "Login Failed, wrong username/password",
			})
		}
	}
}

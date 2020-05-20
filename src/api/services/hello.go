package services

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"dob":       "YYYY/MM/DD",
			"fullName":  "ANG MIE TJIN ONE",
			"gbCd":      "D",
			"idCardNo":  3322110000000001,
			"maritalCd": 1,
			"motherNm":  "mother name",
			"phoneNo":   6281312341234,
			"place":     "BOGOR",
			"sexCd":     1,
		})
	})

	r.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user added",
		})
	})

	r.DELETE("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user deleted",
		})
	})

	r.PUT("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user updated",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

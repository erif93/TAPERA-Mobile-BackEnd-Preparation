package services

import (
	"net/http"

	models "../models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var items []models.Profile

	db.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})

}

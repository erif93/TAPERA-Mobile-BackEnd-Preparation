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

func CreateItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var input models.Profile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item := models.Profile{First_Name: input.First_Name, Last_Name: input.Last_Name, Description: input.Description}

	db.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})

}

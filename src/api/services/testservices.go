package services

import (
	"net/http"

	models "../models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetItem Function
func GetItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var items []models.Profile

	db.Find(&items)

	c.JSON(http.StatusOK, gin.H{"data": items})

}

// UpdateItem function
func UpdateItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var item models.Profile
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateProfile
	db.Model(&item).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

// DeleteItem function
func DeleteItem(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Get model if exist
	var item models.Profile
	if err := db.Where("id = ?", c.Param("id")).First(&item).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&item)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func CreateItem(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var input models.Profile
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	item := models.Profile{FirstName: input.FirstName, LastName: input.LastName, Description: input.Description}

	db.Create(&item)

	c.JSON(http.StatusOK, gin.H{"data": item})

}

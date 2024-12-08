package handlers

import (
	"net/http"
	"strconv"

	"marketplace-gin/database"
	"marketplace-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&seller).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func GetSeller(c *gin.Context) {
	var seller models.Seller
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := database.DB.First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(http.StatusOK, seller)
}

func UpdateSeller(c *gin.Context) {
	var seller models.Seller
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := database.DB.First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	seller.ID = id
	database.DB.Save(&seller)
	c.JSON(http.StatusOK, seller)
}

func DeleteSeller(c *gin.Context) {
	var seller models.Seller
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := database.DB.First(&seller, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}
	database.DB.Delete(&seller)
	c.JSON(http.StatusOK, gin.H{"message": "Seller deleted"})
}

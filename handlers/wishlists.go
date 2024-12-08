package handlers

import (
	"net/http"
	"strconv"

	"marketplace-gin/database"
	"marketplace-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateWishlist(c *gin.Context) {
	var wishlist models.Wishlist
	if err := c.ShouldBindJSON(&wishlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	if wishlist.UserID == 0 || wishlist.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "UserID and ProductID are required"})
		return
	}

	if err := database.DB.Create(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": wishlist})
}

func GetWishlist(c *gin.Context) {
	var wishlist []models.Wishlist
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid UserID"})
		return
	}

	if err := database.DB.Where("user_id = ?", userID).Find(&wishlist).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Wishlist not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": wishlist})
}

func UpdateWishlist(c *gin.Context) {
	var wishlist models.Wishlist
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ID"})
		return
	}
	if err := database.DB.First(&wishlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Wishlist not found"})
		return
	}

	var input models.Wishlist
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	wishlist.ProductID = input.ProductID // Update hanya kolom tertentu
	if err := database.DB.Save(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": wishlist})
}

func DeleteWishlist(c *gin.Context) {
	var wishlist models.Wishlist
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ID"})
		return
	}
	if err := database.DB.First(&wishlist, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Wishlist not found"})
		return
	}
	if err := database.DB.Delete(&wishlist).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Wishlist deleted"})
}

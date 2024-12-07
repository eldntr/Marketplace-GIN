package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func CreateWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    if err := c.ShouldBindJSON(&wishlist); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&wishlist).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, wishlist)
}

func GetWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    c.JSON(http.StatusOK, wishlist)
}

func UpdateWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    if err := c.ShouldBindJSON(&wishlist); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    wishlist.ID = id
    database.DB.Save(&wishlist)
    c.JSON(http.StatusOK, wishlist)
}

func DeleteWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    database.DB.Delete(&wishlist)
    c.JSON(http.StatusOK, gin.H{"message": "Wishlist deleted"})
}
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func CreateCartItem(c *gin.Context) {
    var cartItem models.CartItem
    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cartItem)
}

func GetCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    c.JSON(http.StatusOK, cartItem)
}

func UpdateCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cartItem.ID = id
    database.DB.Save(&cartItem)
    c.JSON(http.StatusOK, cartItem)
}

func DeleteCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    database.DB.Delete(&cartItem)
    c.JSON(http.StatusOK, gin.H{"message": "CartItem deleted"})
}

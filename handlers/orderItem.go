package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func CreateOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    if err := c.ShouldBindJSON(&orderItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&orderItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orderItem)
}

func GetOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    c.JSON(http.StatusOK, orderItem)
}

func UpdateOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    if err := c.ShouldBindJSON(&orderItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    orderItem.ID = id
    database.DB.Save(&orderItem)
    c.JSON(http.StatusOK, orderItem)
}

func DeleteOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    database.DB.Delete(&orderItem)
    c.JSON(http.StatusOK, gin.H{"message": "OrderItem deleted"})
}
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func CreateReply(c *gin.Context) {
    var reply models.Reply
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&reply).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, reply)
}

func GetReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    c.JSON(http.StatusOK, reply)
}

func UpdateReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    reply.ID = id
    database.DB.Save(&reply)
    c.JSON(http.StatusOK, reply)
}

func DeleteReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    database.DB.Delete(&reply)
    c.JSON(http.StatusOK, gin.H{"message": "Reply deleted"})
}
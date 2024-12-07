package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func CreateDiscussion(c *gin.Context) {
    var discussion models.Discussion
    if err := c.ShouldBindJSON(&discussion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&discussion).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, discussion)
}

func GetDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    c.JSON(http.StatusOK, discussion)
}

func UpdateDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    if err := c.ShouldBindJSON(&discussion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    discussion.ID = id
    database.DB.Save(&discussion)
    c.JSON(http.StatusOK, discussion)
}

func DeleteDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    database.DB.Delete(&discussion)
    c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted"})
}
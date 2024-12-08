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

func ReplyDiscussion(c *gin.Context) {
    var reply models.Discussion
    parentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parent ID"})
        return
    }

    // Verify parent discussion exists
    var parentDiscussion models.Discussion
    if err := database.DB.First(&parentDiscussion, parentID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Parent discussion not found"})
        return
    }

    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Set the product_id same as parent discussion
    reply.ProductID = parentDiscussion.ProductID
    reply.ParentID = &parentID

    if err := database.DB.Create(&reply).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reply)
}

func GetDiscussionsByProduct(c *gin.Context) {
    var discussions []models.Discussion
    productID, err := strconv.ParseUint(c.Param("productId"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    // Get main discussions (where parent_id is NULL)
    if err := database.DB.Where("product_id = ? AND parent_id IS NULL", productID).Find(&discussions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // For each discussion, get its replies
    for i := range discussions {
        var replies []models.Discussion
        if err := database.DB.Where("parent_id = ?", discussions[i].ID).Find(&replies).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        discussions[i].Replies = replies
    }

    c.JSON(http.StatusOK, discussions)
}
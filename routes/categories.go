package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	r.POST("/categories", handlers.CreateCategory)
	r.GET("/categories/:id", handlers.GetCategory)
	r.PUT("/categories/:id", handlers.UpdateCategory)
	r.DELETE("/categories/:id", handlers.DeleteCategory)
}
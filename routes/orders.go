package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterOrderRoutes(r *gin.Engine) {
	r.POST("/orders", handlers.CreateOrder)
	r.GET("/orders/:id", handlers.GetOrder)
	r.PUT("/orders/:id", handlers.UpdateOrder)
	r.DELETE("/orders/:id", handlers.DeleteOrder)
}
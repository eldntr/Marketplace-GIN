package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterOrderItemRoutes(r *gin.Engine) {
	r.POST("/order_items", handlers.CreateOrderItem)
	r.GET("/order_items/:id", handlers.GetOrderItem)
	r.PUT("/order_items/:id", handlers.UpdateOrderItem)
	r.DELETE("/order_items/:id", handlers.DeleteOrderItem)
}
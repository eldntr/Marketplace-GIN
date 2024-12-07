package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterCartItemRoutes(r *gin.Engine) {
	r.POST("/cart_items", handlers.CreateCartItem)
	r.GET("/cart_items/:id", handlers.GetCartItem)
	r.PUT("/cart_items/:id", handlers.UpdateCartItem)
	r.DELETE("/cart_items/:id", handlers.DeleteCartItem)
}
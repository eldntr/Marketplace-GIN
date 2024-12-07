package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterCartRoutes(r *gin.Engine) {
	r.POST("/carts", handlers.CreateCart)
	r.GET("/carts/:id", handlers.GetCart)
	r.PUT("/carts/:id", handlers.UpdateCart)
	r.DELETE("/carts/:id", handlers.DeleteCart)
}
package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterWishlistRoutes(r *gin.Engine) {
	r.POST("/wishlists", handlers.CreateWishlist)
	r.GET("/wishlists/:id", handlers.GetWishlist)
	r.PUT("/wishlists/:id", handlers.UpdateWishlist)
	r.DELETE("/wishlists/:id", handlers.DeleteWishlist)
}
package routes

import (
	"marketplace-gin/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSellerRoutes(r *gin.Engine) {
	r.POST("/sellers", handlers.CreateSeller)
	r.GET("/sellers/:id", handlers.GetSeller)
	r.PUT("/sellers/:id", handlers.UpdateSeller)
	r.DELETE("/sellers/:id", handlers.DeleteSeller)
}

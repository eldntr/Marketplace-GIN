package routes

import (
	"marketplace-gin/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSellerRoutes(r *gin.Engine) {
	r.GET("/seller/dashboard", handlers.GetSellerDashboard)
	r.GET("/seller/summary", handlers.GetSellerSummary)
	r.GET("/seller/products", handlers.GetSellerProducts)
	r.POST("/seller/products", handlers.CreateProduct)
	r.PUT("/seller/products/:id", handlers.UpdateProduct)
	r.DELETE("/seller/products/:id", handlers.DeleteProduct)
	r.GET("/seller/orders", handlers.GetSellerOrders)
	r.GET("/seller/reviews", handlers.GetSellerReviews)
}

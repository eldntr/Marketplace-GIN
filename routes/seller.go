package routes

import (
	"marketplace-gin/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSellerRoutes(r *gin.Engine) {
	r.GET("/seller/dashboard", handlers.GetSellerDashboard)
	r.GET("/api/seller/summary", handlers.GetSellerSummary)
	r.GET("/api/seller/products", handlers.GetSellerProducts)
	r.POST("/api/seller/products", handlers.CreateProduct)
	r.PUT("/api/seller/products/:id", handlers.UpdateProduct)
	r.DELETE("/api/seller/products/:id", handlers.DeleteProduct)
	r.GET("/api/seller/orders", handlers.GetSellerOrders)
	r.GET("/api/seller/reviews", handlers.GetSellerReviews)
}

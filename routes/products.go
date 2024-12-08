package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterProductRoutes(r *gin.Engine) {
	// Products routes
	r.POST("/products", handlers.CreateProduct)
	r.GET("/products/:id", handlers.GetProduct)
	r.GET("/products", handlers.GetAllProducts)
	r.PUT("/products/:id", handlers.UpdateProduct)
	r.DELETE("/products/:id", handlers.DeleteProduct)
}	
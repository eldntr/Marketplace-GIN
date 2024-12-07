package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterTransactionRoutes(r *gin.Engine) {
	r.POST("/transactions", handlers.CreateTransaction)
	r.GET("/transactions/:id", handlers.GetTransaction)
	r.PUT("/transactions/:id", handlers.UpdateTransaction)
	r.DELETE("/transactions/:id", handlers.DeleteTransaction)
}
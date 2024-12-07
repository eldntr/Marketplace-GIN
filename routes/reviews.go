package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterReviewRoutes(r *gin.Engine) {
	r.POST("/reviews", handlers.CreateReview)
	r.GET("/reviews/:id", handlers.GetReview)
	r.PUT("/reviews/:id", handlers.UpdateReview)
	r.DELETE("/reviews/:id", handlers.DeleteReview)
}
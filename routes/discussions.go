package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterDiscussionRoutes(r *gin.Engine) {
	r.POST("/discussions", handlers.CreateDiscussion)
	r.GET("/discussions/:id", handlers.GetDiscussion)
	r.PUT("/discussions/:id", handlers.UpdateDiscussion)
	r.DELETE("/discussions/:id", handlers.DeleteDiscussion)
	r.POST("/discussions/:id/reply", handlers.ReplyDiscussion)
	r.GET("/product/:productId/discussions", handlers.GetDiscussionsByProduct) // Ensure this line is correct
}
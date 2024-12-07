package routes

import (
	"github.com/gin-gonic/gin"
	"marketplace-gin/handlers"
)

func RegisterReplyRoutes(r *gin.Engine) {
	r.POST("/replies", handlers.CreateReply)
	r.GET("/replies/:id", handlers.GetReply)
	r.PUT("/replies/:id", handlers.UpdateReply)
	r.DELETE("/replies/:id", handlers.DeleteReply)
}
package main

import (
	"marketplace-gin/database"
	"marketplace-gin/handlers"
	"marketplace-gin/middlewares"
	"marketplace-gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.InitializeDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Gin-GORM MariaDB example",
		})
	})

	r.POST("/auth/signup", handlers.CreateUser)
	r.POST("/auth/login", handlers.Login)
	r.GET("/user/profile", middlewares.CheckAuth, handlers.GetUserProfile)
	r.POST("/cart/:cart_id/checkout", handlers.CheckoutCart)
	r.POST("/transaction/:transaction_id/pay", handlers.PayTransaction)

	// Setup routes
	routes.RegisterUserRoutes(r)
	routes.RegisterCartRoutes(r)
	routes.RegisterCategoryRoutes(r)
	routes.RegisterProductRoutes(r)
	routes.RegisterCartItemRoutes(r)
	routes.RegisterDiscussionRoutes(r)
	routes.RegisterOrderRoutes(r)
	routes.RegisterOrderItemRoutes(r)
	routes.RegisterReplyRoutes(r)
	routes.RegisterReviewRoutes(r)
	routes.RegisterTransactionRoutes(r)
	routes.RegisterWishlistRoutes(r)
	routes.RegisterSellerRoutes(r)

	r.Run(":8080")
}

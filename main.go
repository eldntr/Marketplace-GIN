package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/routes"
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

    r.Run(":8080")
}
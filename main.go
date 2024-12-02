package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    "marketplace-gin/database"
    "marketplace-gin/models"
)

func initDB() {
    database.ConnectDB()
    database.MigrateDB()
}

func executeSQLFile(filepath string) {
    content, err := ioutil.ReadFile(filepath)
    if err != nil {
        log.Fatalf("Could not read SQL file: %v", err)
    }

    commands := strings.Split(string(content), ";")
    for _, command := range commands {
        command = strings.TrimSpace(command)
        if command != "" {
            if err := database.DB.Exec(command).Error; err != nil {
                log.Printf("Error executing command: %s\nError: %v", command, err)
            }
        }
    }
}

func getTables(c *gin.Context) {
    var tables []string
    database.DB.Raw("SHOW TABLES").Scan(&tables)
    c.JSON(http.StatusOK, gin.H{
        "tables": tables,
    })
}

func main() {
    initDB()

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Welcome to the Gin-GORM MariaDB example",
        })
    })

    r.GET("/tables", getTables)

    r.POST("/users", createUser)
    r.GET("/users/:id", getUser)
    r.PUT("/users/:id", updateUser)
    r.DELETE("/users/:id", deleteUser)

    r.POST("/carts", createCart)
    r.GET("/carts/:id", getCart)
    r.PUT("/carts/:id", updateCart)
    r.DELETE("/carts/:id", deleteCart)

    r.POST("/categories", createCategory)
    r.GET("/categories/:id", getCategory)
    r.PUT("/categories/:id", updateCategory)
    r.DELETE("/categories/:id", deleteCategory)

    r.POST("/products", createProduct)
    r.GET("/products/:id", getProduct)
    r.PUT("/products/:id", updateProduct)
    r.DELETE("/products/:id", deleteProduct)

    r.POST("/cart_items", createCartItem)
    r.GET("/cart_items/:id", getCartItem)
    r.PUT("/cart_items/:id", updateCartItem)
    r.DELETE("/cart_items/:id", deleteCartItem)

    r.POST("/discussions", createDiscussion)
    r.GET("/discussions/:id", getDiscussion)
    r.PUT("/discussions/:id", updateDiscussion)
    r.DELETE("/discussions/:id", deleteDiscussion)

    r.POST("/migrations", createMigration)
    r.GET("/migrations/:id", getMigration)
    r.PUT("/migrations/:id", updateMigration)
    r.DELETE("/migrations/:id", deleteMigration)

    r.POST("/orders", createOrder)
    r.GET("/orders/:id", getOrder)
    r.PUT("/orders/:id", updateOrder)
    r.DELETE("/orders/:id", deleteOrder)

    r.POST("/order_items", createOrderItem)
    r.GET("/order_items/:id", getOrderItem)
    r.PUT("/order_items/:id", updateOrderItem)
    r.DELETE("/order_items/:id", deleteOrderItem)

    r.POST("/replies", createReply)
    r.GET("/replies/:id", getReply)
    r.PUT("/replies/:id", updateReply)
    r.DELETE("/replies/:id", deleteReply)

    r.POST("/reviews", createReview)
    r.GET("/reviews/:id", getReview)
    r.PUT("/reviews/:id", updateReview)
    r.DELETE("/reviews/:id", deleteReview)

    r.POST("/sessions", createSession)
    r.GET("/sessions/:id", getSession)
    r.PUT("/sessions/:id", updateSession)
    r.DELETE("/sessions/:id", deleteSession)

    r.POST("/transactions", createTransaction)
    r.GET("/transactions/:id", getTransaction)
    r.PUT("/transactions/:id", updateTransaction)
    r.DELETE("/transactions/:id", deleteTransaction)

    r.POST("/wishlists", createWishlist)
    r.GET("/wishlists/:id", getWishlist)
    r.PUT("/wishlists/:id", updateWishlist)
    r.DELETE("/wishlists/:id", deleteWishlist)

    r.Run(":8080")
}

func createUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

func getUser(c *gin.Context) {
    var user models.User
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}

func updateUser(c *gin.Context) {
    var user models.User
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    user.ID = id
    database.DB.Save(&user)
    c.JSON(http.StatusOK, user)
}

func deleteUser(c *gin.Context) {
    var user models.User
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func createCart(c *gin.Context) {
    var cart models.Cart
    if err := c.ShouldBindJSON(&cart); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&cart).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cart)
}

func getCart(c *gin.Context) {
    var cart models.Cart
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cart, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }
    c.JSON(http.StatusOK, cart)
}

func updateCart(c *gin.Context) {
    var cart models.Cart
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cart, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }
    if err := c.ShouldBindJSON(&cart); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cart.ID = id
    database.DB.Save(&cart)
    c.JSON(http.StatusOK, cart)
}

func deleteCart(c *gin.Context) {
    var cart models.Cart
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cart, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
        return
    }
    database.DB.Delete(&cart)
    c.JSON(http.StatusOK, gin.H{"message": "Cart deleted"})
}

func createCategory(c *gin.Context) {
    var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&category).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, category)
}

func getCategory(c *gin.Context) {
    var category models.Category
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    c.JSON(http.StatusOK, category)
}

func updateCategory(c *gin.Context) {
    var category models.Category
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    category.ID = id
    database.DB.Save(&category)
    c.JSON(http.StatusOK, category)
}

func deleteCategory(c *gin.Context) {
    var category models.Category
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }
    database.DB.Delete(&category)
    c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

func createProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, product)
}

func getProduct(c *gin.Context) {
    var product models.Product
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

func updateProduct(c *gin.Context) {
    var product models.Product
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    product.ID = id
    database.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func deleteProduct(c *gin.Context) {
    var product models.Product
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    database.DB.Delete(&product)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func createCartItem(c *gin.Context) {
    var cartItem models.CartItem
    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, cartItem)
}

func getCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    c.JSON(http.StatusOK, cartItem)
}

func updateCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    cartItem.ID = id
    database.DB.Save(&cartItem)
    c.JSON(http.StatusOK, cartItem)
}

func deleteCartItem(c *gin.Context) {
    var cartItem models.CartItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&cartItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "CartItem not found"})
        return
    }
    database.DB.Delete(&cartItem)
    c.JSON(http.StatusOK, gin.H{"message": "CartItem deleted"})
}

func createDiscussion(c *gin.Context) {
    var discussion models.Discussion
    if err := c.ShouldBindJSON(&discussion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&discussion).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, discussion)
}

func getDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    c.JSON(http.StatusOK, discussion)
}

func updateDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    if err := c.ShouldBindJSON(&discussion); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    discussion.ID = id
    database.DB.Save(&discussion)
    c.JSON(http.StatusOK, discussion)
}

func deleteDiscussion(c *gin.Context) {
    var discussion models.Discussion
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&discussion, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Discussion not found"})
        return
    }
    database.DB.Delete(&discussion)
    c.JSON(http.StatusOK, gin.H{"message": "Discussion deleted"})
}

func createMigration(c *gin.Context) {
    var migration models.Migration
    if err := c.ShouldBindJSON(&migration); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&migration).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, migration)
}

func getMigration(c *gin.Context) {
    var migration models.Migration
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&migration, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Migration not found"})
        return
    }
    c.JSON(http.StatusOK, migration)
}

func updateMigration(c *gin.Context) {
    var migration models.Migration
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&migration, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Migration not found"})
        return
    }
    if err := c.ShouldBindJSON(&migration); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    migration.ID = id
    database.DB.Save(&migration)
    c.JSON(http.StatusOK, migration)
}

func deleteMigration(c *gin.Context) {
    var migration models.Migration
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&migration, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Migration not found"})
        return
    }
    database.DB.Delete(&migration)
    c.JSON(http.StatusOK, gin.H{"message": "Migration deleted"})
}

func createOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, order)
}

func getOrder(c *gin.Context) {
    var order models.Order
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    c.JSON(http.StatusOK, order)
}

func updateOrder(c *gin.Context) {
    var order models.Order
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    order.ID = id
    database.DB.Save(&order)
    c.JSON(http.StatusOK, order)
}

func deleteOrder(c *gin.Context) {
    var order models.Order
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }
    database.DB.Delete(&order)
    c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func createOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    if err := c.ShouldBindJSON(&orderItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&orderItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, orderItem)
}

func getOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    c.JSON(http.StatusOK, orderItem)
}

func updateOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    if err := c.ShouldBindJSON(&orderItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    orderItem.ID = id
    database.DB.Save(&orderItem)
    c.JSON(http.StatusOK, orderItem)
}

func deleteOrderItem(c *gin.Context) {
    var orderItem models.OrderItem
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&orderItem, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
        return
    }
    database.DB.Delete(&orderItem)
    c.JSON(http.StatusOK, gin.H{"message": "OrderItem deleted"})
}

func createReply(c *gin.Context) {
    var reply models.Reply
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&reply).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, reply)
}

func getReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    c.JSON(http.StatusOK, reply)
}

func updateReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    if err := c.ShouldBindJSON(&reply); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    reply.ID = id
    database.DB.Save(&reply)
    c.JSON(http.StatusOK, reply)
}

func deleteReply(c *gin.Context) {
    var reply models.Reply
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&reply, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Reply not found"})
        return
    }
    database.DB.Delete(&reply)
    c.JSON(http.StatusOK, gin.H{"message": "Reply deleted"})
}

func createReview(c *gin.Context) {
    var review models.Review
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&review).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, review)
}

func getReview(c *gin.Context) {
    var review models.Review
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&review, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
        return
    }
    c.JSON(http.StatusOK, review)
}

func updateReview(c *gin.Context) {
    var review models.Review
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&review, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
        return
    }
    if err := c.ShouldBindJSON(&review); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    review.ID = id
    database.DB.Save(&review)
    c.JSON(http.StatusOK, review)
}

func deleteReview(c *gin.Context) {
    var review models.Review
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&review, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
        return
    }
    database.DB.Delete(&review)
    c.JSON(http.StatusOK, gin.H{"message": "Review deleted"})
}

func createSession(c *gin.Context) {
    var session models.Session
    if err := c.ShouldBindJSON(&session); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&session).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, session)
}

func getSession(c *gin.Context) {
    var session models.Session
    id := c.Param("id")
    if err := database.DB.First(&session, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
        return
    }
    c.JSON(http.StatusOK, session)
}

func updateSession(c *gin.Context) {
    var session models.Session
    id := c.Param("id")
    if err := database.DB.First(&session, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
        return
    }
    if err := c.ShouldBindJSON(&session); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    session.ID = id
    database.DB.Save(&session)
    c.JSON(http.StatusOK, session)
}

func deleteSession(c *gin.Context) {
    var session models.Session
    id := c.Param("id")
    if err := database.DB.First(&session, "id = ?", id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
        return
    }
    database.DB.Delete(&session)
    c.JSON(http.StatusOK, gin.H{"message": "Session deleted"})
}

func createTransaction(c *gin.Context) {
    var transaction models.Transaction
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&transaction).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, transaction)
}

func getTransaction(c *gin.Context) {
    var transaction models.Transaction
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&transaction, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        return
    }
    c.JSON(http.StatusOK, transaction)
}

func updateTransaction(c *gin.Context) {
    var transaction models.Transaction
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&transaction, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        return
    }
    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    transaction.ID = id
    database.DB.Save(&transaction)
    c.JSON(http.StatusOK, transaction)
}

func deleteTransaction(c *gin.Context) {
    var transaction models.Transaction
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&transaction, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
        return
    }
    database.DB.Delete(&transaction)
    c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}

func createWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    if err := c.ShouldBindJSON(&wishlist); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := database.DB.Create(&wishlist).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, wishlist)
}

func getWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    c.JSON(http.StatusOK, wishlist)
}

func updateWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    if err := c.ShouldBindJSON(&wishlist); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    wishlist.ID = id
    database.DB.Save(&wishlist)
    c.JSON(http.StatusOK, wishlist)
}

func deleteWishlist(c *gin.Context) {
    var wishlist models.Wishlist
    id, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    if err := database.DB.First(&wishlist, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Wishlist not found"})
        return
    }
    database.DB.Delete(&wishlist)
    c.JSON(http.StatusOK, gin.H{"message": "Wishlist deleted"})
}
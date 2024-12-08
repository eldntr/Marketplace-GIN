package handlers

import (
	"net/http"
	"strconv"

	"marketplace-gin/database"
	"marketplace-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateCartItem(c *gin.Context) {
	var cartItem models.CartItem
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cartItem.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Quantity must be greater than 0"})
		return
	}
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}
	if product.Stock < cartItem.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Kurangi stok produk
	product.Stock -= cartItem.Quantity
	database.DB.Save(&product)

	// Simpan item ke keranjang
	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItem)
}

func GetCartItem(c *gin.Context) {
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

	// Perbarui jumlah item
	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi jumlah baru terhadap stok
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	stockDelta := input.Quantity - cartItem.Quantity
	if product.Stock < stockDelta {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Perbarui stok produk
	product.Stock -= stockDelta
	database.DB.Save(&product)

	// Perbarui jumlah item dalam keranjang
	cartItem.Quantity = input.Quantity
	if err := database.DB.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cartItem)
}

func UpdateCartItem(c *gin.Context) {
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

	// Perbarui jumlah item
	var input struct {
		Quantity int `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi jumlah baru terhadap stok
	var product models.Product
	if err := database.DB.First(&product, cartItem.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	stockDelta := input.Quantity - cartItem.Quantity
	if product.Stock < stockDelta {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	// Perbarui stok produk
	product.Stock -= stockDelta
	database.DB.Save(&product)

	// Perbarui jumlah item dalam keranjang
	cartItem.Quantity = input.Quantity
	if err := database.DB.Save(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cartItem.ID = id
	database.DB.Save(&cartItem)
	c.JSON(http.StatusOK, cartItem)
}

func DeleteCartItem(c *gin.Context) {
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

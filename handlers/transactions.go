package handlers

import (
	"net/http"
	"strconv"

	"marketplace-gin/database"
	"marketplace-gin/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
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

func GetTransaction(c *gin.Context) {
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

func UpdateTransaction(c *gin.Context) {
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

func DeleteTransaction(c *gin.Context) {
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

func CheckoutCart(c *gin.Context) {
	var cart models.Cart
	cartID, err := strconv.ParseUint(c.Param("cart_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Cart ID"})
		return
	}

	// Ambil Cart beserta item-nya
	if err := database.DB.Preload("CartItems").First(&cart, cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	// Hitung Total Harga
	var totalAmount float64
	for _, item := range cart.CartItems {
		totalAmount += float64(item.Quantity) * item.Price
	}

	// Buat Transaksi
	transaction := models.Transaction{
		UserID:        cart.BuyerID,
		CartID:        cart.ID,
		TotalAmount:   totalAmount,
		PaymentStatus: "Pending",
	}
	if err := database.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func PayTransaction(c *gin.Context) {
	var transaction models.Transaction
	transactionID, err := strconv.ParseUint(c.Param("transaction_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Transaction ID"})
		return
	}

	// Ambil Transaksi
	if err := database.DB.First(&transaction, transactionID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Update Status Pembayaran
	transaction.PaymentStatus = "Paid"
	if err := database.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

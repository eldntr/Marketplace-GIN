package handlers

import (
	"marketplace-gin/database"
	"marketplace-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSellerDashboard - Mengambil data untuk dashboard seller
func GetSellerDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

// GetSellerSummary - Mengambil ringkasan penjualan
func GetSellerSummary(c *gin.Context) {
	// Logika untuk mengambil ringkasan penjualan
	var summary = map[string]interface{}{
		"totalSales":    1000,  // Ganti dengan logika nyata
		"totalProducts": 50,    // Ganti dengan logika nyata
		"totalRevenue":  50000, // Ganti dengan logika nyata
	}
	c.JSON(http.StatusOK, summary)
}

// GetSellerProducts - Mengambil daftar produk seller
func GetSellerProducts(c *gin.Context) {
	var products []models.Product
	if err := database.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetSellerOrders - Mengambil daftar pesanan seller
func GetSellerOrders(c *gin.Context) {
	// Logika untuk mengambil daftar pesanan
	var orders []models.Order
	if err := database.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetSellerReviews - Mengambil daftar ulasan seller
func GetSellerReviews(c *gin.Context) {
	// Logika untuk mengambil daftar ulasan
	var reviews []models.Review
	if err := database.DB.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

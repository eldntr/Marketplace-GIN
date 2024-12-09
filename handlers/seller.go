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
	var totalSales int64
	var totalProducts int64
	var totalRevenue float64

	// Menghitung total penjualan
	if err := database.DB.Model(&models.Order{}).Count(&totalSales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total produk
	if err := database.DB.Model(&models.Product{}).Count(&totalProducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total pendapatan
	if err := database.DB.Model(&models.Order{}).Select("SUM(total) as totalRevenue").Scan(&totalRevenue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menyusun ringkasan penjualan
	var summary = map[string]interface{}{
		"totalSales":    totalSales,
		"totalProducts": totalProducts,
		"totalRevenue":  totalRevenue,
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

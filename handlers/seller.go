package handlers

import (
	"marketplace-gin/database"
	"marketplace-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSellerDashboard - Mengambil data untuk dashboard seller
func GetSellerDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

// GetSellerSummary - Mengambil ringkasan penjualan seller berdasarkan ID seller
func GetSellerSummary(c *gin.Context) {
	sellerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var totalSales int64
	var totalProducts int64
	var totalRevenue float64
	var inProcess int64

	// Menghitung total penjualan
	if err := database.DB.Model(&models.Order{}).Where("seller_id = ?", sellerID).Count(&totalSales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total produk
	if err := database.DB.Model(&models.Product{}).Where("seller_id = ?", sellerID).Count(&totalProducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung total pendapatan
	if err := database.DB.Model(&models.Order{}).Where("seller_id = ?", sellerID).Select("SUM(total) as totalRevenue").Scan(&totalRevenue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menghitung jumlah pesanan yang sedang diproses
	if err := database.DB.Model(&models.Order{}).Where("status = ? AND seller_id = ?", "pending", sellerID).Count(&inProcess).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	summary := gin.H{
		"total_sales":       totalSales,
		"total_products":    totalProducts,
		"total_revenue":     totalRevenue,
		"orders_in_process": inProcess,
	}
	c.JSON(http.StatusOK, summary)
}

// GetSellerProducts - Mengambil daftar produk seller berdasarkan ID seller
func GetSellerProducts(c *gin.Context) {
	sellerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var products []models.Product
	if err := database.DB.Where("seller_id = ?", sellerID).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetSellerOrders - Mengambil daftar pesanan seller berdasarkan ID seller
func GetSellerOrders(c *gin.Context) {
	sellerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var orders []models.Order
	if err := database.DB.Where("seller_id = ?", sellerID).Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetSellerReviews - Mengambil daftar ulasan seller berdasarkan ID seller
func GetSellerReviews(c *gin.Context) {
	sellerID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var reviews []models.Review
	if err := database.DB.Joins("JOIN products ON products.id = reviews.product_id").
		Where("products.seller_id = ?", sellerID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reviews)
}

// UpdateProduct - Mengupdate informasi produk seller
func UpdateProduct(c *gin.Context) {
	var product models.Product
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Mencari produk berdasarkan ID
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Mengikat data JSON yang diterima
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Menyimpan perubahan produk
	product.ID = id
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

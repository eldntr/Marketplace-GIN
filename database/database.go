
package database

import (
    "log"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "marketplace-gin/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "newuser:newpassword@tcp(localhost:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    DB.Exec("CREATE DATABASE IF NOT EXISTS marketplace")

    dsn = "newuser:newpassword@tcp(localhost:3306)/marketplace?charset=utf8mb4&parseTime=True&loc=Local"
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }
}

func MigrateDB() {
    DB.AutoMigrate(&models.User{}, &models.Cart{}, &models.Category{}, &models.Product{}, &models.CartItem{}, &models.Discussion{}, &models.Migration{}, &models.Order{}, &models.OrderItem{}, &models.Reply{}, &models.Review{}, &models.Session{}, &models.Transaction{}, &models.Wishlist{})
}
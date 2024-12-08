
// filepath: /home/debian/Documents/Go/Marketplace/Marketplace-Gin/models/models.go
package models

import "time"

type User struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"password"`
    Role      string    `json:"role"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Cart struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    BuyerID   uint64    `json:"buyer_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Category struct {
    ID          uint64    `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type Product struct {
    ID          uint64    `gorm:"primaryKey" json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Image       string    `json:"image"`
    Price       float64   `json:"price"`
    Stock       int       `json:"stock"`
    SellerID    uint64    `json:"seller_id"`
    CategoryID  uint64    `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CartItem struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    CartID    uint64    `json:"cart_id"`
    ProductID uint64    `json:"product_id"`
    Quantity  int       `json:"quantity"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Discussion struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    ProductID uint64    `json:"product_id"`
    UserID    uint64    `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Order struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    BuyerID   uint64    `json:"buyer_id"`
    Status    string    `json:"status"`
    Total     float64   `json:"total"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    OrderID   uint64    `json:"order_id"`
    ProductID uint64    `json:"product_id"`
    Quantity  int       `json:"quantity"`
    Price     float64   `json:"price"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Reply struct {
    ID           uint64    `gorm:"primaryKey" json:"id"`
    DiscussionID uint64    `json:"discussion_id"`
    UserID       uint64    `json:"user_id"`
    Content      string    `json:"content"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

type Review struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    ProductID uint64    `json:"product_id"`
    UserID    uint64    `json:"user_id"`
    Rating    int       `json:"rating"`
    Comment   string    `json:"comment"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type Transaction struct {
    ID         uint64    `gorm:"primaryKey" json:"id"`
    UserID     uint64    `json:"user_id"`
    ProductID  uint64    `json:"product_id"`
    Quantity   int       `json:"quantity"`
    TotalPrice float64   `json:"total_price"`
    Status     string    `json:"status"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

type Wishlist struct {
    ID        uint64    `gorm:"primaryKey" json:"id"`
    UserID    uint64    `json:"user_id"`
    ProductID uint64    `json:"product_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

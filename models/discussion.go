package models

import "time"

type Discussion struct {
    ID        uint64      `json:"id" gorm:"primaryKey"`
    ProductID uint64      `json:"product_id"`
    UserID    uint64      `json:"user_id"`
    ParentID  *uint64     `json:"parent_id"`
    Content   string      `json:"content"`
    Replies   []Discussion `json:"replies" gorm:"-"`
    CreatedAt time.Time   `json:"created_at"`
    UpdatedAt time.Time   `json:"updated_at"`
}
package models

import "time"

type Discussion struct {
	ID        uint64     `gorm:"primary_key" json:"id"`
	ProductID uint64     `json:"product_id"`
	UserID    uint64     `json:"user_id"`
	ParentID  *uint64    `json:"parent_id,omitempty"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
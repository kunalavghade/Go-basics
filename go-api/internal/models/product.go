package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"not null;uniqueIndex"`
	IsActive bool   `json:"is_active" gorm:"default:true"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	Products []Product `json:"-" gorm:"foreignKey:CategoryID;references:Id"`
}

type Product struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name" gorm:"not null;uniqueIndex"`
	Description string  `json:"description" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int     `json:"stock" gorm:"not null"`
	CategoryID  int     `json:"category_id" gorm:"not null;index"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	Category   Category       `json:"-" gorm:"foreignKey:CategoryID;references:Id"`
	CartItems  []CartItem     `json:"-" gorm:"foreignKey:ProductID;references:Id"`
	Images     []ProductImage `json:"-" gorm:"foreignKey:ProductID;references:Id"`
	OrderItems []OrderItem    `json:"-" gorm:"foreignKey:ProductID;references:Id"`
}

type ProductImage struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"product_id" gorm:"not null;index"`
	ImageURL  string `json:"image_url" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	Product Product `json:"-" gorm:"foreignKey:ProductID;references:Id"`
}

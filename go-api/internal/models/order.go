package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID          int         `json:"id" gorm:"primaryKey"`
	UserID      int         `json:"-" gorm:"not null;index"`
	Status      OrderStatus `json:"status" gorm:"not null"`
	TotalAmount float64     `json:"total_amount" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	User       User        `json:"-" gorm:"foreignKey:UserID;references:Id"`
	OrderItems []OrderItem `json:"-" gorm:"foreignKey:OrderID;references:Id"`
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"-" gorm:"not null;index"`
	ProductID int     `json:"-" gorm:"not null;index"`
	Quantity  int     `json:"quantity" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	Order   Order   `json:"-" gorm:"foreignKey:OrderID;references:Id"`
	Product Product `json:"-" gorm:"foreignKey:ProductID;references:Id"`
}

type Cart struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	UserID    int            `json:"-" gorm:"not null;index"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	User      User       `json:"-" gorm:"foreignKey:UserID;references:Id"`
	CartItems []CartItem `json:"-" gorm:"foreignKey:CartID;references:Id"`
}

type CartItem struct {
	ID        int `json:"id" gorm:"primaryKey"`
	CartID    int `json:"-" gorm:"not null;index"`
	ProductID int `json:"-" gorm:"not null;index"`
	Quantity  int `json:"quantity" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	Cart    Cart    `json:"-" gorm:"foreignKey:CartID;references:Id"`
	Product Product `json:"-" gorm:"foreignKey:ProductID;references:Id"`
}

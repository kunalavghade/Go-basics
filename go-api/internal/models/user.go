package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"-" gorm:"not null"`

	Phone     string   `json:"phone" gorm:"column:phone;uniqueIndex;not null"`
	Address   string   `json:"address" gorm:"-"`
	IsActive  bool     `json:"is_active" gorm:"default:true"`
	FirstName string   `json:"first_name" gorm:"not null"`
	LastName  string   `json:"last_name" gorm:"not null"`
	Role      UserRole `json:"role" gorm:"default:'customer'"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationships
	RefreshTokens []RefreshToken `json:"-" gorm:"foreignKey:UserID;references:Id"`
	Orders        []Order        `json:"-" gorm:"foreignKey:UserID;references:Id"`
	Cart          []Cart         `json:"-" gorm:"foreignKey:UserID;references:Id"`
}

type UserRole string

const (
	UserRoleCustomer UserRole = "customer"
	UserRoleAdmin    UserRole = "admin"
)

type RefreshToken struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"-" gorm:"not null;index"`
	Token     string    `json:"-" gorm:"not null;uniqueIndex;type:varchar(512)"`
	ExpiresAt time.Time `json:"expires_at" gorm:"not null"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-"`

	// Relationship
	User User `json:"-" gorm:"foreignKey:UserID;references:Id"`
}

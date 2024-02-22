package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name" binding:"required"`
	Email    string `gorm:"not null;unique" json:"email" binding:"required"`
	Password string `gorm:"not null" json:"password" binding:"required"`
	Phone    string `gorm:"not null" json:"phone" binding:"required"`
	Address  string `gorm:"not null" json:"address" binding:"required"`
	City     string `gorm:"not null" json:"city" binding:"required"`
	State    string `gorm:"not null" json:"state" binding:"required"`
	Country  string `gorm:"not null" json:"country" binding:"required"`
	Role     string `gorm:"not null;default:basic" json:"role" binding:"required"`
	// Products []product.Product `gorm:"foreignKey:UserID"`
}

type UserUpdateBody struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"not null;unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Phone    string `gorm:"not null" json:"phone"`
	Address  string `gorm:"not null" json:"address"`
	City     string `gorm:"not null" json:"city"`
	State    string `gorm:"not null" json:"state"`
	Country  string `gorm:"not null" json:"country"`
	Role     string `gorm:"not null;default:basic" json:"role"`
}

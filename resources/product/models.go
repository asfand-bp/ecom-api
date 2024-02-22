package product

import (
	"ecom/resources/auth"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name" binding:"required"`
	Description string  `gorm:"not null" json:"description" binding:"required"`
	Price       float64 `gorm:"not null" json:"price" binding:"required"`
	UserID      uint    `gorm:"not null;foreignKey:User.ID" json:"user_id" binding:"required"`
	Category    string  `gorm:"not null" json:"category"`
	Brand       string  `gorm:"not null" json:"brand"`
	Stock       uint    `gorm:"not null" json:"stock"`
	Image       string  `gorm:"not null" json:"image"`
}

type Order struct {
	gorm.Model
	UserID     auth.User `gorm:"not null;foreignKey:User.ID" json:"user_id" binding:"required"`
	TotalPrice float64   `gorm:"not null" json:"total_price" binding:"required"`
	Status     string    `gorm:"not null;default:pending" json:"status" binding:"required"`
	Products   Product   `gorm:"not null;foreignKey:Product.ID" json:"products" binding:"required"`
	Quantity   uint      `gorm:"not null" json:"quantity" binding:"required"`
}

package product

import (
	"ecom/db"

	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func NewService() *Service {
	return &Service{
		DB: db.DB,
	}
}

func (r *Service) GetAllProducts() []Product {
	// Get all users code goes here
}

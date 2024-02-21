package product

import (
	"ecom/db"

	"gorm.io/gorm"
)

type Service struct {
	DB     *gorm.DB
	Access *Service
}

func NewService() *Service {
	return &Service{
		DB:     db.DB,
		Access: NewService(),
	}
}

func (r *Service) GetAllProducts() []Product {
	// Get all users code goes here
}

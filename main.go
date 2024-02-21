package main

import (
	"ecom/db"
	"ecom/router"

	"github.com/gin-gonic/gin"
)

func init() {
	db.StartDB() // Initialize the database connection
	// Migrate the schema(s)
	// db.AutoMigrate(&auth.User{})
	// db.AutoMigrate(&product.Product{})
}

func main() {
	r := gin.Default() // Gin engine with default stuff

	// Initialize API Endpoints
	router.Init(r)

	r.Run() // listens and serves on 0.0.0.0:8080 (on lan and local)
}

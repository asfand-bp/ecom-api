package tests

import (
	"ecom/db"
	"ecom/resources/auth"
	"ecom/resources/product"
	"ecom/router"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func StartDBAndMigrate() {
	db.StartDB() // Initialize the database connection
	// Migrate the schema(s)
	db.DB.AutoMigrate(&auth.User{})
	db.DB.AutoMigrate(&product.Product{})
	db.DB.AutoMigrate(&product.Order{})
}

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	// Routes
	router.Init(r)

	return r
}

func DestroyDB() {
	databaseFilePath := "./app.db"
	err := os.Remove(databaseFilePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}
}

func CreateUser(t *testing.T, R *gin.Engine) auth.User {
	// Define the User data
	userData := auth.User{
		Name:     "Test User",
		Email:    "testuser@gmail.com",
		Password: "test",
		Phone:    "23498285328",
		Address:  "4th Street somewhere in California",
		City:     "Los Angeles",
		State:    "CA",
		Country:  "USA",
		Role:     "basic",
	}

	// Make POST api request
	w := MakePostRequest(t, R, "/users", userData)

	// Parse the response body
	var user auth.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	return user
}

func DeleteUser(t *testing.T, R *gin.Engine, id string) {
	// Make Delete api request
	MakeDeleteRequest(t, R, "/users/"+id)
}

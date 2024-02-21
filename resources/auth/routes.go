package auth

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	controller := New()

	// Get All Users - /v1/api/users
	r.GET("/users", controller.GetAllUsers)
}

package auth

import "github.com/gin-gonic/gin"

func Init(r *gin.RouterGroup) {
	controller := New()

	// User Endpoints - /v1/api/users
	r.GET("/users", controller.GetAllUsers)
	r.POST("/users", controller.PostUser)
	r.GET("/users/:id", controller.GetUserByID)
	r.PUT("/users/:id", controller.PutUser)
	r.DELETE("/users/:id", controller.DeleteUser)
}

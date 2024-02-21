package router

import (
	"ecom/resources/auth"
	"ecom/resources/product"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) { // c.JSON serializes the map given in 2nd argument
	c.JSON(200, gin.H{ // H is shortcut for map[string]any
		"message": "Welcome to homepage",
	})
}

func Init(r *gin.Engine) {
	// All route endpoints initialization goes here
	r.GET("/", index)

	router_group := r.Group("/v1/api") // New Router group for api endpoints

	auth.Init(router_group)
	product.Init(router_group)
}

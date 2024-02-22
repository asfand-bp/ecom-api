package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service ServiceInterface
}

func New() *Controller {
	return &Controller{
		Service: NewService(),
	}
}

func (controller *Controller) GetAllUsers(c *gin.Context) {
	// Get all users
	res, err := controller.Service.GetAllUsers()

	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (controller *Controller) GetUserByID(c *gin.Context) {
	// Get users by ID
	id, ok := c.Params.Get("id")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	res, err := controller.Service.GetUserByID(id)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (controller *Controller) PostUser(c *gin.Context) {
	// Add users
	var input User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := controller.Service.CreateUser(input)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (controller *Controller) PutUser(c *gin.Context) {
	// Update user
	var input UserUpdateBody

	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	res, err := controller.Service.UpdateUser(id, input)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (controller *Controller) DeleteUser(c *gin.Context) {
	// Delete user
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	res, err := controller.Service.DeleteUser(id)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, res)
}

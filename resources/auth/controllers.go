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
	res := controller.Service.GetAllUsers()

	c.JSON(http.StatusOK, res)
}

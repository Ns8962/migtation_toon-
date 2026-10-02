package controllers

import (
	"myapi/models"
	"myapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service *services.UserService
}

// constructor
func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		Service: service,
	}
}

// method
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := ctrl.Service.CreateUser(input)
	c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	users := ctrl.Service.GetUsers()
	c.JSON(http.StatusOK, users)
}

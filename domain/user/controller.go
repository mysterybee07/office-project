package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/office-project-setup/domain/model"
	// "github.com/mysterybee07/office-project-setup/domain/user"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not bind JSON",
		})
		return
	}

	createdUser, err := u.service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user":    createdUser,
	})
}

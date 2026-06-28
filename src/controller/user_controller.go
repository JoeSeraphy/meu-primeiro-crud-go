package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/service"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)

	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}

// UpdateUser implements [UserControllerInterface].

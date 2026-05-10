package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/validation"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"

	"go.uber.org/zap"
)

func (uc *userController) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("Journey", "createUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user request", err,
			zap.String("Journey", "createUser"))

		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.serviceInterface.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	fmt.Println(userRequest)
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zap.String("Journey", "createUser"))
	c.JSON(201, response)
}

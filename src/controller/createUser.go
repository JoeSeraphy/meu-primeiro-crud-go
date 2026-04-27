package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/validation"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/request"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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

package controller

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/validation"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/view"
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

	domainResult, err := uc.serviceInterface.CreateUser(domain)
	if err != nil {
		logger.Error("Error trying to create user in service", err,
			zap.String("Journey", "createUser"))	
		c.JSON(err.Code, err)
		return
	}
	logger.Info("User created successfully", 
	zap.String("userId", domainResult.GetId()),
	zap.String("Journey", "createUser"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,))
}

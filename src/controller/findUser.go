package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init FindUserByID Controller",
		zap.String("Journey", "findUserByID"))

	userId := c.Param("userId")

	if _, err := uuid.Parse(userId); err != nil {
		logger.Error("Error parsing user ID", err,
			zap.String("journey", "findUserByID"))
		errorMessage := rest_err.NewBadRequestError(
			"invalid user id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error trying to find user by ID in service", err,
			zap.String("Journey", "findUserByID"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByID successfully",
		zap.String("journey", "findUserByID"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init FindUserByEmail Controller",
		zap.String("Journey", "findUserByEmail"))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error parsing user email", err,
			zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"invalid user email")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to find user by email in service", err,
			zap.String("Journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("FindUserByEmail successfully",
		zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

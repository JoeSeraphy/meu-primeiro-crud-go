package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/validation"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser Controller",
		zap.String("Journey", "updateUser"))
	var userUpdateRequest request.UserUpdateRequest

	userId := c.Param("id")

	if err := c.ShouldBindJSON(&userUpdateRequest); err != nil ||
		strings.TrimSpace(userId) == "" {
		logger.Error("Error trying to validate user update request", err,
			zap.String("Journey", "updateUser"))
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		userUpdateRequest.Name,
		userUpdateRequest.Age,
	)
	err := uc.service.UpdateUser(userId, domain)
	if err != nil {
		logger.Error("Error trying to update user", err,
			zap.String("Journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}
	logger.Info("updateUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("Journey", "updateUser"))

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

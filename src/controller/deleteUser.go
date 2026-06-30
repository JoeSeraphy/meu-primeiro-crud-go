package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser Controller",
		zap.String("Journey", "DeleteUser"))

	userId := c.Param("userId")

	if strings.TrimSpace(userId) == "" {
		logger.Error("Error: userId is empty", nil,
			zap.String("Journey", "DeleteUser"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	err := uc.service.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to delete user", err,
			zap.String("Journey", "DeleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("DeleteUser controller executed successfully",
		zap.String("userId", userId),
		zap.String("Journey", "DeleteUser"))

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

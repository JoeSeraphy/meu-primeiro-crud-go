package service

import (
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by ID in service...",
		zap.String("journey", "findUserByID"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (
	model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by email in service...",
		zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}

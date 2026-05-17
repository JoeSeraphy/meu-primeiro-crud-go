package service

import (
	"fmt"
	"log"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomain, *rest_err.RestErr) {
	logger.Info("Init createUser modelo", zap.String("journey", "CreateUser"))
	userDomain.EncryptPassword()
	userDomainRepository, err := ud.respository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to create user in database", err,
			zap.String("journey", "CreateUser"))
		return nil, err
	}
	return userDomainRepository, nil
}

package model

import (
	"fmt"

	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (*userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser modelo", zap.String("journey", "CreateUser"))
	userDomain.EncryptPassword()
	fmt.Print(userDomain.GetPassword())
	return nil
}

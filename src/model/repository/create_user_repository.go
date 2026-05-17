package repository

import (
	"context"
	"os"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_errors"
	"github.com/joelsantiago/meu-primeiro-crud-go/src/model/repository/entity/convert"
)

const (
		POSTGRE_USER_DB = "POSTGRESQL_USER_DB"
)

func (ur *userRepository) CreateUser(UserDomain model.UserDomainInterface,
		) (model.UserDomainInterface, *rest_errors.RestErr) {

			logger.Info("Creating user in repository...")
			collection_name := os.Getenv(POSTGRE_USER_DB)

			collection := ur.databaseConnection.Collection(collection_name)

			value := convert.ConvertDomainToEntity(UserDomain)
			if err != nil {
				return nil, rest_errors.NewInternalServerError(err.Error())
			}
			result, err := collection.InsertOne(context.Background(), value)
			if err != nil {
				return nil, rest_errors.NewInternalServerError(err.Error())
			}

			userDomain.SetID(result.InsertedID.(string))

			return userDomain, nil
		}
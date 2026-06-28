package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository/entity/convert"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	userId string,
	UserDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Updating user in repository...",
		zap.String("journey", "updateUser"))
	tableName := os.Getenv(POSTGRE_USER_DB)
	if tableName == "" {
		tableName = "users"
	}
	value := convert.ConvertDomainToEntity(UserDomain)

	query := fmt.Sprintf(`
				UPDATE %s 
				SET email = $1, password = $2, name = $3, age = $4
				WHERE id = $5
				RETURNING id`, tableName)

	var lastInsertID string
	err := ur.databaseConnection.QueryRow(
		context.Background(),
		query,
		value.Email,
		value.Password,
		value.Name,
		value.Age,
		value.ID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.Error("Error updating user in repository", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = lastInsertID

	return convert.ConvertEntityToDomain(value), nil
}

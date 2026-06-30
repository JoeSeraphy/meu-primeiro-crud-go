package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"

	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository/entity/convert"
)

const (
	POSTGRE_USER_DB = "POSTGRESQL_USER_DB"
)

func (ur *userRepository) CreateUser(UserDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Creating user in repository...")
	tableName := os.Getenv(POSTGRE_USER_DB)
	if tableName == "" {
		tableName = "users"
	}
	value := convert.ConvertDomainToEntity(UserDomain)

	query := fmt.Sprintf(`
				INSERT INTO %s (id, email, password, name, age)
				VALUES ($1, $2, $3, $4, $5) 
				RETURNING id`, tableName)

	userId := uuid.NewString()
	var lastInsertID string
	err := ur.databaseConnection.QueryRow(
		context.Background(),
		query,
		userId,
		value.Email,
		value.Password,
		value.Name,
		value.Age,
	).Scan(&lastInsertID)

	if err != nil {
		logger.Error("Error creating user in repository", err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = lastInsertID

	return convert.ConvertEntityToDomain(value), nil
}

package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository/entity"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository/entity/convert"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by email in repository...")
	tableName := os.Getenv(POSTGRE_USER_DB)
	if tableName == "" {
		tableName = "users"
	}

	query := fmt.Sprintf(`
				SELECT id, email, password, name, age
				FROM %s
				WHERE email = $1`, tableName)

	var userEntity entity.UserEntity
	err := ur.databaseConnection.QueryRow(
		context.Background(),
		query,
		email,
	).Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Name,
		&userEntity.Age,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Error("User not found with email:"+email, err)
			return nil, rest_err.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to find user by email", err)
		return nil, rest_err.NewInternalServerError("Database error'")
	}

	return convert.ConvertEntityToDomain(&userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Finding user by ID in repository...")
	tableName := os.Getenv(POSTGRE_USER_DB)
	if tableName == "" {
		tableName = "users"
	}
	query := fmt.Sprintf(`
				SELECT id, email, password, name, age
				FROM %s
				WHERE id = $1`, tableName)

	var userEntity entity.UserEntity
	err := ur.databaseConnection.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&userEntity.ID,
		&userEntity.Email,
		&userEntity.Password,
		&userEntity.Name,
		&userEntity.Age,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			logger.Error("User not found with ID:"+id, err)
			return nil, rest_err.NewNotFoundError("User not found")
		}
		logger.Error("Error trying to find user by ID", err)
		return nil, rest_err.NewInternalServerError("Database error'")
	}

	return convert.ConvertEntityToDomain(&userEntity), nil
}

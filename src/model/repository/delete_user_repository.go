package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *rest_err.RestErr {
	logger.Info("Deleting user from repository...",
		zap.String("journey", "deleteUser"))
	tableName := os.Getenv("POSTGRES_USER_DB")
	if tableName == "" {
		tableName = "users"
	}

	query := fmt.Sprintf(`
				DELETE FROM %s 
				WHERE id = $1`, tableName)

	_, err := ur.databaseConnection.Exec(
		context.Background(),
		query,
		userId,
	)

	if err != nil {
		logger.Error("Error deleting user from repository", err)
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil
}

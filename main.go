package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/database/postgre"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/routes"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/service"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
	}
	ctx := context.Background()
	database, err := postgre.NewPostgreConnection(ctx)
	if err != nil {
		logger.Error("Failed to connect to database", err)
		return
	}
	defer database.Close(ctx)

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		logger.Error("Failed to run server", err)
	}
}

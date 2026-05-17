package main

import (
	"log"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller/routes"
	"github.com/joho/godotenv"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/service"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/database/postgre"
)

func main() {
	logger.Info("Starting server...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err :=postgre.NewPostgreConnection(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v \n", err.Error())
		return
	}

	service := service.NewUserDomainService()
	userController := initDependecies(database)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

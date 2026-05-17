package main

import (
	"github.com/joho/godotenv"
	"github.com/yourusername/yourproject/controller"
	"github.com/yourusername/yourproject/database/postgres"
	"github.com/yourusername/yourproject/model/service"
	"github.com/yourusername/yourproject/repository"
)

func  initDependecies(	database *postgres.Database) (
	controller.UserControllerInterface,
	) {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserController(service)
}
package main

import (
	"github.com/joeseraphy/meu-primeiro-crud-go/src/configuration/postgres"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/repository"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/service"
)

func initDependecies(database *postgres.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserController(service)
}

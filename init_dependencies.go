package main

import (
	"github.com/jackc/pgx/v5"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/controller"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/repository"
	"github.com/joeseraphy/meu-primeiro-crud-go/src/model/service"
)

func initDependecies(database *pgx.Conn) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}

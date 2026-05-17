package convert

import (
	"github.com/joelsantiago/meu-primeiro-crud-go/src/model"
	"github.com/joelsantiago/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		ID:	     domain.GetID(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
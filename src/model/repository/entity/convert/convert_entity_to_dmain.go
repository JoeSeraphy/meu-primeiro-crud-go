package convert

import (
	"github.com/joelsantiago/meu-primeiro-crud-go/src/model"
	"github.com/joelsantiago/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity *entity.UserEntity,
) model.userDomainInterface {
	domain := model.NewUserDomain{
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	}
	domain.SetID(entity.ID)
	return domain
}

package repository


func NewUserRepository(
	databaseConnection *postgre.Database,
) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *postgre.Database
}

type UserRepository interface {
	CreateUser(UserDomain model.UserDomainInterface,
		) (model.UserDomainInterface, *rest_errors.RestErr)
}

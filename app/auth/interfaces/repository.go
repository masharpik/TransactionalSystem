package authinterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

//go:generate mockgen -source=repository.go -destination=mocks_repository.go -package=authinterfaces IAuthRepository
type IAuthRepository interface {
	CreateUser(utils.User) error
	MinusBalance(string, float64) (float64, error)
	PlusBalance(string, float64) (float64, error)
}

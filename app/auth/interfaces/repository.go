package authinterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

//go:generate mockgen -source=repository.go -destination=mocks_repository.go -package=authinterfaces IAuthRepository
type IAuthRepository interface {
	CreateUser(utils.User) error
	GetUser(string) (float64, error)
	UpdateBalance(string, float64) error
}

package authinterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

type IAuthRepository interface {
	CreateUser(utils.User) error
	GetUser(string) (float64, error)
	UpdateBalance(string, float64) error
}

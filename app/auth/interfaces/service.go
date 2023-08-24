package authinterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

type IAuthService interface {
	CreateUser() (utils.User, error)
}

package authinterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

//go:generate mockgen -source=service.go -destination=mocks_service.go -package=authinterfaces IAuthService
type IAuthService interface {
	CreateUser() (utils.User, error)
}

package authservice

import (
	"fmt"

	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

type Service struct {
	repo authinterfaces.IAuthRepository
}

func NewService(repo authinterfaces.IAuthRepository) (*Service, error) {
	service := &Service{
		repo: repo,
	}

	_, ok := interface{}(service).(authinterfaces.IAuthService)
	if !ok {
		return nil, fmt.Errorf(literals.LogStructNotSatisfyInterface, "AuthService")
	}

	return service, nil
}

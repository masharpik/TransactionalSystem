package authservice

import (
	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
)

type Service struct {
	repo authinterfaces.IAuthRepository
}

func NewService(repo authinterfaces.IAuthRepository) (*Service, error) {
	service := &Service{
		repo: repo,
	}

	return service, nil
}

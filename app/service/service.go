package service

import (
	"fmt"

	"github.com/masharpik/TransactionalSystem/app/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

type Service struct {
	repo interfaces.IRepository
}

func NewService(repo interfaces.IRepository) (*Service, error) {
	service := &Service{
		repo: repo,
	}

	_, ok := interface{}(service).(interfaces.IService)
	if !ok {
		return nil, fmt.Errorf(literals.LogStructNotSatisfyInterface, "Service")
	}

	return service, nil
}

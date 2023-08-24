package transactionservice

import (
	"fmt"

	"github.com/masharpik/TransactionalSystem/app/transaction/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

type Service struct {
	repo transactioninterfaces.ITransactionRepository
}

func NewService(repo transactioninterfaces.ITransactionRepository) (*Service, error) {
	service := &Service{
		repo: repo,
	}

	_, ok := interface{}(service).(transactioninterfaces.ITransactionService)
	if !ok {
		return nil, fmt.Errorf(literals.LogStructNotSatisfyInterface, "TransactionService")
	}

	return service, nil
}

package transactionservice

import (
	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	transactioninterfaces "github.com/masharpik/TransactionalSystem/app/transaction/interfaces"
)

type Service struct {
	transactionRepo transactioninterfaces.ITransactionRepository
	authRepo        authinterfaces.IAuthRepository
}

func NewService(transactionRepo transactioninterfaces.ITransactionRepository, authRepo authinterfaces.IAuthRepository) (*Service, error) {
	service := &Service{
		transactionRepo: transactionRepo,
		authRepo:        authRepo,
	}

	return service, nil
}

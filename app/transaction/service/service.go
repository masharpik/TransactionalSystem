package transactionservice

import (
	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	"github.com/masharpik/TransactionalSystem/app/rabbitmq"
)

type Service struct {
	sender   *rabbitmq.Sender
	authRepo authinterfaces.IAuthRepository
}

func NewService(sender *rabbitmq.Sender, authRepo authinterfaces.IAuthRepository) (*Service, error) {
	service := &Service{
		sender:   sender,
		authRepo: authRepo,
	}

	return service, nil
}

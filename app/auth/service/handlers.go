package authservice

import (
	"github.com/google/uuid"

	"github.com/masharpik/TransactionalSystem/app/auth/utils"
)

func (service *Service) CreateUser() (createdUser utils.User, err error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return
	}

	createdUser = utils.User{
		UserID: uuid.String(),
		Balance: 0,
	}

	err = service.repo.CreateUser(createdUser)
	if err != nil {
		return
	}

	return
}

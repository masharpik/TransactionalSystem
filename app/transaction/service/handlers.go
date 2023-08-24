package transactionservice

import (
	"fmt"

	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
)

func (service *Service) InputMoney(userId string, amount float64) (user authUtils.User, err error) {
	var oldAmount float64
	oldAmount, err = service.authRepo.GetUser(userId)
	if err != nil {
		return
	}

	newAmount := oldAmount + amount
	if newAmount < 0 {
		err = fmt.Errorf(utils.LogUnderfundedError)
		return
	}

	err = service.authRepo.UpdateBalance(userId, newAmount)
	if err != nil {
		return
	}

	user = authUtils.User{
		UserID: userId,
		Balance: newAmount,
	}

	return
}

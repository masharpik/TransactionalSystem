package transactionservice

import (
	"fmt"
	"time"

	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func (service *Service) InputMoney(userId string, amount float64) (user authUtils.User, err error) {
	var oldAmount float64
	oldAmount, err = service.authRepo.GetUser(userId)
	if err != nil {
		return
	}

	newAmount := oldAmount + amount

	err = service.authRepo.UpdateBalance(userId, newAmount)
	if err != nil {
		return
	}

	user = authUtils.User{
		UserID:  userId,
		Balance: newAmount,
	}

	return
}

func (service *Service) OutputMoney(userId string, amount float64, link string) (status utils.StatusTransaction, err error) {
	var oldAmount float64
	oldAmount, err = service.authRepo.GetUser(userId)
	if err != nil {
		return
	}

	newAmount := oldAmount - amount
	if newAmount < 0 {
		err = fmt.Errorf(utils.LogUnderfundedError)
		return
	}

	go func() {
		err := service.authRepo.UpdateBalance(userId, newAmount)
		if err != nil {
			logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке записать в бд снятие с баланса: %w", err))
			return
		}

		<-time.After(5 * time.Second)
		err = nil // Гипотетический результат от банка
		if err != nil {
			logger.LogOperationError(fmt.Errorf("Произошла ошибка при запросе к банку: %w\nВозврат средств произойдет в течении нескольких минут.", err))
			err := service.authRepo.UpdateBalance(userId, oldAmount)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке вернуть средства на баланс: %w", err))
				return
			}

			return
		}

		service.sender.PushTask(userId, newAmount, link)
	}()

	status = utils.StatusTransaction{
		UserID: userId,
		Status: fmt.Sprintf("Информация по результату снятия придет по ссылке: %s", link),
	}
	return
}

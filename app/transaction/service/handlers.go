package transactionservice

import (
	"fmt"
	"math"
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
	curr, secondTransaction := math.Inf(1), false
	defer func() {
		if !math.IsInf(curr, 1) && !secondTransaction {
			_, err := service.authRepo.PlusBalance(userId, amount)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке вернуть баланс обратно: %w", err))
				return
			}
		}
	}()
	curr, err = service.authRepo.MinusBalance(userId, amount)
	if err != nil {
		logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке записать в бд снятие с баланса: %w", err))
		return
	}

	if curr < 0 {
		_, err = service.authRepo.PlusBalance(userId, amount)
		if err != nil {
			logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке вернуть баланс обратно: %w", err))
			return
		}
		err = fmt.Errorf(utils.LogUnderfundedError)
		return
	}

	go func() {
		<-time.After(5 * time.Second)
		// Здесь в принципе банк вернет какой-то результат и по нему можно будет смотреть, выполнил ли банк операцию, но пока заглушка с флагом
		secondTransaction = true
		err = nil // Гипотетический результат от банка
		if err != nil {
			logger.LogOperationError(fmt.Errorf("Произошла ошибка при запросе к банку: %w\nВозврат средств произойдет в течении нескольких минут.", err))
			_, err := service.authRepo.PlusBalance(userId, amount)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке вернуть средства на баланс: %w", err))
				return
			}

			return
		}

		service.sender.PushTask(userId, curr, link)
	}()

	status = utils.StatusTransaction{
		UserID: userId,
		Status: fmt.Sprintf("Информация по результату снятия придет по ссылке: %s", link),
	}
	return
}

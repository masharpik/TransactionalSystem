package transactionservice

import (
	"fmt"
	"net/http"
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
		<-time.After(5 * time.Second)

		client := &http.Client{}
		req, _ := http.NewRequest("GET", link, nil)

		callback := func() {
			err := service.authRepo.UpdateBalance(userId, newAmount)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке записать в бд снятие с баланса: %w", err))
				req.Header.Set("Custom-Status", "500")
				return
			}

			logger.LogOperationSuccess(fmt.Sprintf("Снятие с баланса.\nТекущее состояние пользователя: %s\nБаланс: %f\n", userId, newAmount))
			req.Header.Set("Custom-Status", "200")

			_, err = client.Do(req)
			if err != nil {
				logger.LogOperationError(fmt.Errorf("Произошла ошибка при попытке отослать простой GET-запрос: %w", err))
			}
		}
		service.sender.PushTask(callback)
	}()

	status = utils.StatusTransaction{
		UserID: userId,
		Status: fmt.Sprintf("Информация по результату снятия придет по ссылке: %s", link),
	}
	return
}

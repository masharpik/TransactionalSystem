package transactionservice

import (
	"fmt"
	"time"

	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func (service *Service) InputMoney(userId string, amount float64) (user authUtils.User, err error) {
	curr, err := service.authRepo.PlusBalance(userId, amount)
	if err != nil {
		return
	}

	user = authUtils.User{
		UserID:  userId,
		Balance: curr,
	}

	return
}

func (service *Service) OutputMoney(userId string, amount float64, link string) (res utils.OutputTransactionResponse, err error) {
	curr, err := service.authRepo.MinusBalance(userId, amount)
	if err != nil {
		res = utils.OutputTransactionResponse{
			UserID: userId,
			Status: err.Error(),
		}
		return
	}

	go func() {
		// Имитация запроса в банк
		<-time.After(5 * time.Second)
		err = nil // Возврат банка

		var status utils.StatusTransaction
		if err != nil {
			curr, err = service.authRepo.PlusBalance(userId, amount)
			if err != nil {
				logger.LogOperationError(err)
				return
			}

			status = utils.StatusTransaction{
				UserID:      userId,
				Status:      fmt.Sprintf("Произошла ошибка от банка: %s", err.Error()),
				Balance:     curr,
				Destination: link,
			}
		} else {
			status = utils.StatusTransaction{
				UserID:      userId,
				Status:      "Списание успешно",
				Balance:     curr,
				Destination: link,
			}
		}

		if err = service.sender.PushTask(status); err != nil {
			logger.LogOperationError(err)
			return
		}
	}()

	res = utils.OutputTransactionResponse{
		UserID: userId,
		Status: fmt.Sprintf("Запрос на вывод средств взят в работу. По завершении мы отправим уведомление сюда: %s", link),
	}
	return
}

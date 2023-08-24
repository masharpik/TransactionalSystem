package transactioninterfaces

import "github.com/masharpik/TransactionalSystem/app/auth/utils"

type ITransactionService interface {
	InputMoney(string, float64) (utils.User, error)
}

package transactioninterfaces

import (
	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
)

//go:generate mockgen -source=service.go -destination=mocks_service.go -package=transactioninterfaces ITransactionService
type ITransactionService interface {
	InputMoney(string, float64) (authUtils.User, error)
	OutputMoney(string, float64, string) (utils.OutputTransactionResponse, error)
}

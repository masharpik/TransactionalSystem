package transactiondelivery

import (
	"fmt"

	"github.com/gorilla/mux"

	transactioninterfaces "github.com/masharpik/TransactionalSystem/app/transaction/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

type Delivery struct {
	router  *mux.Router
	service transactioninterfaces.ITransactionService
}

func RegisterHandlers(r *mux.Router, service transactioninterfaces.ITransactionService) error {
	router := &Delivery{
		router:  r,
		service: service,
	}

	// Проверка, что структура удовляет требуемому интерфейсу
	_, ok := interface{}(router).(transactioninterfaces.ITransactionDelivery)
	if !ok {
		return fmt.Errorf(literals.LogStructNotSatisfyInterface, "TransactionDelivery")
	}

	router.router.HandleFunc("/input", router.InputHandler).Methods("GET")
	router.router.HandleFunc("/output", router.OutputHandler).Methods("GET")

	return nil
}

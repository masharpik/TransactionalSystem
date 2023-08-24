package transactiondelivery

import (
	"github.com/gorilla/mux"

	transactioninterfaces "github.com/masharpik/TransactionalSystem/app/transaction/interfaces"
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

	router.router.HandleFunc("/input", router.InputHandler).Methods("PUT")
	router.router.HandleFunc("/output", router.OutputHandler).Methods("GET")

	return nil
}

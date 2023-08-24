package delivery

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/masharpik/TransactionalSystem/app/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
)

type Delivery struct {
	router  *mux.Router
	service interfaces.IService
}

func RegisterHandlers(r *mux.Router, service interfaces.IService) error {
	router := &Delivery{
		router:  r,
		service: service,
	}

	// Проверка, что структура удовляет требуемому интерфейсу
	_, ok := interface{}(router).(interfaces.IDelivery)
	if !ok {
		return fmt.Errorf(literals.LogStructNotSatisfyInterface, "Delivery")
	}

	router.router.HandleFunc("/input", router.InputHandler).Methods("GET")
	router.router.HandleFunc("/output", router.OutputHandler).Methods("GET")

	return nil
}

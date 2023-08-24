package authdelivery

import (
	"fmt"

	"github.com/gorilla/mux"

	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

type Delivery struct {
	router  *mux.Router
	service authinterfaces.IAuthService
}

func RegisterHandlers(r *mux.Router, service authinterfaces.IAuthService) error {
	router := &Delivery{
		router:  r,
		service: service,
	}

	// Проверка, что структура удовляет требуемому интерфейсу
	_, ok := interface{}(router).(authinterfaces.IAuthDelivery)
	if !ok {
		return fmt.Errorf(literals.LogStructNotSatisfyInterface, "AuthDelivery")
	}

	router.router.HandleFunc("/", router.AuthHandler).Methods("POST")
	return nil
}

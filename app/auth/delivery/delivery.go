package authdelivery

import (
	"github.com/gorilla/mux"

	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
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

	router.router.HandleFunc("", router.AuthHandler).Methods("POST")
	return nil
}

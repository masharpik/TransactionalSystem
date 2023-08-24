package app

import (
	"github.com/gorilla/mux"

	"github.com/masharpik/TransactionalSystem/app/delivery"
	"github.com/masharpik/TransactionalSystem/app/interfaces"
	"github.com/masharpik/TransactionalSystem/app/middleware"
	"github.com/masharpik/TransactionalSystem/app/repository"
	"github.com/masharpik/TransactionalSystem/app/service"
)

func RegisterUrls() (r *mux.Router, err error) {
	r = mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.JSONMiddleware)

	var repo interfaces.IRepository
	repo, err = repository.NewRepository()
	if err != nil {
		return
	}

	service, err := service.NewService(repo)
	if err != nil {
		return
	}

	transactionSubrouter := apiRouter.PathPrefix("/transaction").Subrouter()
	err = delivery.RegisterHandlers(transactionSubrouter, service)
	if err != nil {
		return
	}

	return
}

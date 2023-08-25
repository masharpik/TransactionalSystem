package app

import (
	"github.com/gorilla/mux"

	authdelivery "github.com/masharpik/TransactionalSystem/app/auth/delivery"
	authinterfaces "github.com/masharpik/TransactionalSystem/app/auth/interfaces"
	authrepository "github.com/masharpik/TransactionalSystem/app/auth/repository"
	authservice "github.com/masharpik/TransactionalSystem/app/auth/service"
	"github.com/masharpik/TransactionalSystem/app/middleware"
	"github.com/masharpik/TransactionalSystem/app/rabbitmq"
	transactiondelivery "github.com/masharpik/TransactionalSystem/app/transaction/delivery"
	transactionservice "github.com/masharpik/TransactionalSystem/app/transaction/service"
)

func RegisterUrls(sender *rabbitmq.Sender) (r *mux.Router, err error) {
	r = mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.JSONMiddleware)

	conn, err := getConnectionDB()
	var authRepo authinterfaces.IAuthRepository
	authRepo = authrepository.NewRepository(conn)
	if err != nil {
		return
	}

	transactionService, err := transactionservice.NewService(sender, authRepo)
	if err != nil {
		return
	}
	authService, err := authservice.NewService(authRepo)
	if err != nil {
		return
	}

	transactionSubrouter := apiRouter.PathPrefix("/transaction").Subrouter()
	err = transactiondelivery.RegisterHandlers(transactionSubrouter, transactionService)
	if err != nil {
		return
	}
	authSubrouter := apiRouter.PathPrefix("/auth").Subrouter()
	err = authdelivery.RegisterHandlers(authSubrouter, authService)
	if err != nil {
		return
	}

	return
}

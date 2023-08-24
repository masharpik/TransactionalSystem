package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/masharpik/TransactionalSystem/utils/literals"
	"github.com/masharpik/TransactionalSystem/utils/logger"
)

func StartServer(r *mux.Router) error {
	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVER_APP_HOST"), os.Getenv("SERVER_APP_PORT"))

	logger.LogOperationSuccess(fmt.Sprintf(literals.LogServerWasStarted, addr))
	return http.ListenAndServe(addr, r)
}

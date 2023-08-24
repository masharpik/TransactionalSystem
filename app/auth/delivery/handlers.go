package authdelivery

import (
	"net/http"

	"github.com/masharpik/TransactionalSystem/utils/writer"
)

func (router *Delivery) AuthHandler(w http.ResponseWriter, r *http.Request) {
	createdUser, err := router.service.CreateUser()
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteSuccessJSONResponse(w, r, http.StatusCreated, createdUser)
}

package transactiondelivery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mailru/easyjson"
	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/writer"
)

func CheckMoney(amount string) error {
	parts := strings.Split(amount, ".")

	if len(parts) > 1 && len(parts[1]) > 2 {
		return fmt.Errorf(utils.LogLengthInputMoneyNotCorrectlyError)
	}

	return nil
}

func (router *Delivery) InputHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transaction utils.Transaction
	err := easyjson.UnmarshalFromReader(r.Body, &transaction)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = CheckMoney(fmt.Sprint(transaction.Amount))
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	user, err := router.service.InputMoney(transaction.UserID, transaction.Amount)
	if err != nil {
		errStr := err.Error()

		switch errStr {
		case authUtils.LogUserNotFoundError:
			writer.WriteErrorMessageRespond(w, r, http.StatusUnauthorized, errStr)
			return
		case utils.LogUnderfundedError:
			writer.WriteErrorMessageRespond(w, r, http.StatusUnprocessableEntity, errStr)
			return
		default:
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
			return
		}
	}

	writer.WriteSuccessJSONResponse(w, r, http.StatusCreated, user)
}

func (router *Delivery) OutputHandler(w http.ResponseWriter, r *http.Request) {
}

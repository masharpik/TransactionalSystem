package transactiondelivery

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mailru/easyjson"
	authUtils "github.com/masharpik/TransactionalSystem/app/auth/utils"
	"github.com/masharpik/TransactionalSystem/app/transaction/utils"
	"github.com/masharpik/TransactionalSystem/utils/logger"
	"github.com/masharpik/TransactionalSystem/utils/writer"
)

func CheckMoneyForInput(amount float64) error {
	if amount < 0 {
		return fmt.Errorf(utils.LogSignInputNotCorrectlyError)
	}

	parts := strings.Split(fmt.Sprint(amount), ".")

	if len(parts) > 1 && len(parts[1]) > 2 {
		return fmt.Errorf(utils.LogLengthInputMoneyNotCorrectlyError)
	}

	return nil
}

func CheckMoneyForOutput(amount float64) error {
	if amount < 0 {
		return fmt.Errorf(utils.LogSignOutputNotCorrectlyError)
	}

	parts := strings.Split(fmt.Sprint(amount), ".")

	if len(parts) > 1 && len(parts[1]) > 2 {
		return fmt.Errorf(utils.LogLengthOutputMoneyNotCorrectlyError)
	}

	return nil
}

func (router *Delivery) InputHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transaction utils.InputTransaction
	err := easyjson.UnmarshalFromReader(r.Body, &transaction)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = CheckMoneyForInput(transaction.Amount)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := router.service.InputMoney(transaction.UserID, transaction.Amount)
	if err != nil {
		errStr := err.Error()

		switch errStr {
		case authUtils.LogUserNotFoundError:
			writer.WriteErrorMessageRespond(w, r, http.StatusUnauthorized, errStr)
			return
		default:
			writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, errStr)
			return
		}
	}

	writer.WriteSuccessJSONResponse(w, r, http.StatusOK, user)
}

func (router *Delivery) OutputHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var transaction utils.OutputTransaction
	err := easyjson.UnmarshalFromReader(r.Body, &transaction)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = CheckMoneyForOutput(transaction.Amount)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	status, err := router.service.OutputMoney(transaction.UserID, transaction.Amount, transaction.Link)
	if err != nil {
		writer.WriteErrorMessageRespond(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	writer.WriteSuccessJSONResponse(w, r, http.StatusOK, status)
}

func (router *Delivery) TestHandler(w http.ResponseWriter, r *http.Request) {
	customCode := r.Header.Get("Custom-Code")
	customStatus := r.Header.Get("Custom-Status")
	customUserId := r.Header.Get("Custom-UserId")
	customBalance := r.Header.Get("Custom-Balance")
	logger.LogOperationSuccess(
		fmt.Sprintf(
			"\n\nПолучено уведомлении о снятии средств:\nКод возврата: %s\nСтатус: %s\nПользователь: %s\nСостояние баланса: %s\n",
			customCode, customStatus, customUserId, customBalance,
		),
	)
}

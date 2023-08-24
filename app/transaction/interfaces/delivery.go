package transactioninterfaces

import "net/http"

type ITransactionDelivery interface {
	InputHandler(http.ResponseWriter, *http.Request)
	OutputHandler(http.ResponseWriter, *http.Request)
}

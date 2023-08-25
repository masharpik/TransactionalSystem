package transactioninterfaces

import "net/http"

//go:generate mockgen -source=delivery.go -destination=mocks_delivery.go -package=transactioninterfaces ITransactionDelivery
type ITransactionDelivery interface {
	InputHandler(http.ResponseWriter, *http.Request)
	OutputHandler(http.ResponseWriter, *http.Request)
}

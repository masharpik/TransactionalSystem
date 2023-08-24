package interfaces

import "net/http"

type IDelivery interface {
	InputHandler(http.ResponseWriter, *http.Request)
	OutputHandler(http.ResponseWriter, *http.Request)
}

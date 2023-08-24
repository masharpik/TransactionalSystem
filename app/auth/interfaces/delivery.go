package authinterfaces

import "net/http"

type IAuthDelivery interface {
	AuthHandler(http.ResponseWriter, *http.Request)
}

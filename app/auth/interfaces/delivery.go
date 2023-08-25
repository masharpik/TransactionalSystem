package authinterfaces

import "net/http"

//go:generate mockgen -source=delivery.go -destination=mocks_delivery.go -package=authinterfaces IAuthDelivery
type IAuthDelivery interface {
	AuthHandler(http.ResponseWriter, *http.Request)
}

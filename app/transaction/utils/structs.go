package utils

//go:generate easyjson -all structs.go

type Transaction struct {
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

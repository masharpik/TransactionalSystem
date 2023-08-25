package utils

//go:generate easyjson -all structs.go

type InputTransaction struct {
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

type OutputTransaction struct {
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
	Link   string  `json:"link"`
}

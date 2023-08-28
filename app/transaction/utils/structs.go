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

type OutputTransactionResponse struct {
	UserID string `json:"userId"`
	Status string `json:"status"`
}

type StatusTransaction struct {
	UserID      string  `json:"userId"`
	Status      string  `json:"status"`
	Balance     float64 `json:"balance"`
	Destination string  `json:"link"`
}

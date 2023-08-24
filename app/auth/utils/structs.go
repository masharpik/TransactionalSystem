package utils

//go:generate easyjson -all structs.go

type User struct {
	UserID  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

package myerrors

//go:generate easyjson -all errors.go

type Error struct {
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

func New(msg string) *Error {
	return &Error{
		Message: msg,
	}
}

package models

type Error struct {
	Msg  string
	Type ErrorType
}

func (e *Error) Error() string {
	return e.Msg
}

type ErrorType uint

const (
	ERR_INTERNAL ErrorType = iota
	ERR_BAD_REQUEST
	ERR_NOT_FOUND
)

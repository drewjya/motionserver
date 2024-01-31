package koderor

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorsKode struct {
	Val  *validator.ValidationErrors
	Kode *ErrorKode
}

func NewErrors(
	Val *validator.ValidationErrors,
	Kode *ErrorKode,
) KodeError {
	return &ErrorsKode{
		Val:  Val,
		Kode: Kode,
	}
}

func (_i *ErrorsKode) Error() string {
	return ""
}

type ErrorKode struct {
	Value string
	Key   string
}

type KodeError interface {
	Error() string
}

func New(Key string,
	Value string) KodeError {
	return &ErrorKode{
		Value: Value,
		Key:   Key,
	}
}

func (_i *ErrorKode) Error() string {
	return fmt.Sprintf("%s\n%s", _i.Key, _i.Value)
}

package main

import (
	"errors"
	"fmt"
)

type WrapErr struct {
	msg string
	err error
}

func (we *WrapErr) Error() string {
	return we.msg
}

func (we *WrapErr) Unwrap() error {
	return we.err
}

func main() {
	err := errors.New("base error\n")
	wrappedErr := &WrapErr{"wrap error 1\n", err}

	unwrappedErr1 := errors.Unwrap(wrappedErr)
	fmt.Printf("%#v\n", unwrappedErr1)
	// &errors.errorString{s:"base error\n"}

	unwrappedErr2 := errors.Unwrap(unwrappedErr1)
	fmt.Printf("%#v\n", unwrappedErr2)
	// <nil>

	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	joinErr := errors.Join(err1, err2, err3)
	unwrappedErr3 := errors.Unwrap(joinErr)
	fmt.Printf("%#v\n", unwrappedErr3)
	// <nil>
}

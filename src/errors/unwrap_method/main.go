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
	err1 := &WrapErr{"wrap error 1\n", err}
	err2 := &WrapErr{"wrap error 2\n", err1}
	fmt.Println(err2.Error())
	fmt.Println(err2.Unwrap().Error())
}

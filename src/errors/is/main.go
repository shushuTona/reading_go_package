package main

import (
	"errors"
	"fmt"
)

type MyErr string

func (e MyErr) Error() string { return string(e) }
func (e MyErr) Unwrap() error {
	return e
}

var ERR1 = MyErr("ERR1")
var ERR2 = MyErr("ERR2")

type PointerErr struct {
	msg string
}

func (pe *PointerErr) Error() string {
	return pe.msg
}

func (pe *PointerErr) Is(err error) bool {
	return pe.msg == err.Error()
}

func main() {
	err := fmt.Errorf("wrap error %w", ERR1)
	fmt.Println(errors.Is(err, ERR1))
	// true

	fmt.Println(errors.Is(err, ERR2))
	// false

	e1 := &PointerErr{"e1"}
	e2 := &PointerErr{"e1"}

	fmt.Println(errors.Is(e1, e2))
}

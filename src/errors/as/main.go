package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	msg string
}

func (me *MyErr) Error() string {
	return me.msg
}

type IHoge interface {
	String() string
}

func main() {
	e1 := &MyErr{"e1"}
	e2 := fmt.Errorf("wrap error %w", e1)

	var me1 *MyErr
	fmt.Println(errors.As(e2, &me1))
	// true

	e3 := errors.New("e3")
	var me2 *MyErr
	fmt.Println(errors.As(e3, &me2))
	// false
}

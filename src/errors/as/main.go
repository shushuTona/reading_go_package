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

func main() {
	var e error
	e1 := errors.New("e1")
	fmt.Println(errors.As(e1, &e))

	var me *MyErr
	e2 := &MyErr{"e2"}
	fmt.Println(errors.As(e2, &me))
}

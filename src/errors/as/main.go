package main

import (
	"errors"
	"fmt"
	"reflect"
)

type MyErr struct {
	msg string
}

func (me *MyErr) Error() string {
	return me.msg
}

func main() {
	e1 := &MyErr{"e1"}
	e2 := fmt.Errorf("wrap error %w", e1)

	var me1 *MyErr
	fmt.Println(errors.As(e2, &me1))

	e3 := errors.New("e3")
	var me2 *MyErr
	fmt.Println(errors.As(e3, &me2))

	val := reflect.ValueOf(me2)
	typ := val.Type()
	fmt.Printf("%#v\n", typ.Kind())
	fmt.Printf("%#v\n", val.IsNil())

	if typ.Kind() != reflect.Ptr || val.IsNil() {
		panic("errors: target must be a non-nil pointer")
	}
}

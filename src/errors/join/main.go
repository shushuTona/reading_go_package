package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	joinErr := errors.Join(err1, err2, err3)
	fmt.Println(joinErr.Error())
	// error1
	// error2
	// error3

	joinErr2 := errors.Join(err1, nil, err3)
	fmt.Println(joinErr2.Error())
	// error1
	// error3

	joinErr3 := errors.Join(nil, nil, nil)
	fmt.Println(joinErr3)
	// <nil>
}

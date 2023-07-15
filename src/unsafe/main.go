package main

import (
	"fmt"
	"unsafe"
)

type StructA struct {
	Name string
	Age  int
}

type StructB struct {
	Msg string
	ID  int
}

func main() {
	var sa = StructA{"sa", 20}
	var sb = *(*StructB)(unsafe.Pointer(&sa))

	fmt.Printf("%#v\n", sa)
	// main.StructA{Name:"sa", Age:20}

	fmt.Printf("%#v\n", sb)
	// main.StructB{Msg:"sa", ID:20}
}

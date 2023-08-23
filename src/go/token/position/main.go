package main

import (
	"fmt"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	fmt.Printf("fset : %#v\n\n", fset)

	src1 := "var a = 100"
	src2 := "func b () int { return 200 }"
	src3 := "var c = []int{0, 1, 2}"

	fmt.Printf("len(src1) : %#v\n", len(src1))
	fmt.Printf("len(src2) : %#v\n", len(src2))
	fmt.Printf("len(src3) : %#v\n", len(src3))
	// len(src1) : 11
	// len(src2) : 28
	// len(src3) : 22

	fset.AddFile("file_1", -1, len(src1))
	fset.AddFile("file_2", -1, len(src2))
	fset.AddFile("file_3", -1, len(src3))

	fmt.Printf("fset.Position(10) : %v\n", fset.Position(token.Pos(10)))
	// fset.Position(10) : file_1:1:10

	fmt.Printf("fset.Position(20) : %v\n", fset.Position(token.Pos(20)))
	// fset.Position(20) : file_2:1:8

	fmt.Printf("fset.Position(52) : %v\n", fset.Position(token.Pos(52)))
	// fset.Position(52) : file_3:1:11
}

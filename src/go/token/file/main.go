package main

import (
	"fmt"
	"go/token"
)

func main() {
	f := token.File{}
	fmt.Printf("f name : %#v\n", f.Name())
	fmt.Printf("f base : %#v\n", f.Base())
	fmt.Printf("f size : %#v\n", f.Size())
	fmt.Printf("f lines : %#v\n", f.Lines())
	fmt.Printf("f LineCount : %#v\n\n", f.LineCount())

	f.AddLine(10)
	fmt.Printf("f lines : %#v\n\n", f.Lines())
	// File構造体は各フィールドはプライベートになっているが、下記のoffset < f.sizeがあるため、Fileから直接インスタンスを生成する際に、AddLineは機能しない。
	// if i := len(f.lines); (i == 0 || f.lines[i-1] < offset) && offset < f.size {
	// 	f.lines = append(f.lines, offset)
	// }

	f.SetLines([]int{0, 2})
	fmt.Printf("f lines : %#v\n\n", f.Lines())

	fileSet := token.NewFileSet()

	src := []byte(`var hoge := 100`)
	file1 := fileSet.AddFile("test1.go", -1, len(src))
	fmt.Printf("file1 : %#v\n\n", file1)
	fmt.Printf("file1 name : %#v\n", file1.Name())
	fmt.Printf("file1 base : %#v\n", file1.Base())
	fmt.Printf("file1 size : %#v\n", file1.Size())
	fmt.Printf("file1 lines : %#v\n", file1.Lines())
	fmt.Printf("file1 LineCount : %#v\n\n", file1.LineCount())

	src2 := []byte(`var piyo := 200`)
	file2 := fileSet.AddFile("test2.go", -1, len(src2))
	fmt.Printf("file2 : %#v\n\n", file2)
	fmt.Printf("file2 name : %#v\n", file2.Name())
	fmt.Printf("file2 base : %#v\n", file2.Base())
	fmt.Printf("file2 size : %#v\n", file2.Size())
	fmt.Printf("file2 lines : %#v\n", file2.Lines())
	fmt.Printf("file2 LineCount : %#v\n", file2.LineCount())
}

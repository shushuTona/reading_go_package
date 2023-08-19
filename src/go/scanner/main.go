package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()
	src := []byte(`var hoge := 100`)
	file := fileSet.AddFile("test.go", -1, len(src))

	s := scanner.Scanner{}
	s.Init(file, src, nil, 0)

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}

		fmt.Printf("fileSet.Position(pos) : %+v\n", fileSet.Position(pos))
		fmt.Printf("tok : %+v\n", tok)
		fmt.Printf("lit : %+v\n\n", lit)
	}
}

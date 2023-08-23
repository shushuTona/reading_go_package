package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	fileSet := token.NewFileSet()
	src := `package example

// stringFunc stringを返す関数です
func stringFunc() string {
	return "ABC"
}

// intFunc intを返す関数です
func intFunc() int {
	return 100
}
`
	file := fileSet.AddFile("example.go", -1, len(src))

	s := scanner.Scanner{}
	s.Init(file, []byte(src), nil, 0)
	// s.Init(file, []byte(src), nil, scanner.ScanComments)

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

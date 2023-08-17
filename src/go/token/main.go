package main

import (
	"fmt"
	"go/token"
)

func main() {
	ident := token.Lookup("100")
	fmt.Printf("ident : %+v\n", ident)
	fmt.Printf("%+v\n", ident == token.IDENT)

	keywordVAR := token.Lookup("var")
	fmt.Printf("keywordVAR : %+v\n", keywordVAR)
	fmt.Printf("%+v\n", keywordVAR == token.VAR)

	keywordGO := token.Lookup("go")
	fmt.Printf("keywordGO : %+v\n", keywordGO)
	fmt.Printf("%+v\n", keywordGO == token.GO)
}

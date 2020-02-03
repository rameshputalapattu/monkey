package main

import (
	"fmt"
	"github.com/rameshputalapattu/monkey/lexer"
)

func main() {

	input := `let five six;`

	l := lexer.New(input)
	fmt.Println("Length:", len(input))

	for i := 1; i <= 6; i++ {

		tok := l.NextToken()
		fmt.Println("nxt=", tok.Type, " literal=", tok.Literal)
	}

}

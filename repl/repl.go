package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rameshputalapattu/monkey/lexer"
	"github.com/rameshputalapattu/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

	}

}

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/K-Saikrishnan/go-monkey/src/lexer"
	"github.com/K-Saikrishnan/go-monkey/src/token"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		if scanned := scanner.Scan(); !scanned {
			return
		}

		line := scanner.Text()
		lex := lexer.New(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

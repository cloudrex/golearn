package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"./parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("// /* */ @attr @attr2() fn pub main ~> int { ret 0+2+2; }")

	// Create the Lexer
	lexer := parser.NewGolearnLexer(is)

	// Read all tokens
	for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}

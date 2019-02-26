package main

import (
	"fmt"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type golearnListener struct {
	*parser.BaseGolearnListener
	mainExists bool
}

func newGolearnListener() golearnListener {
	return golearnListener{mainExists: false}
}

func (s *golearnListener) EnterAssign(ctx *parser.AssignContext) {
	id := ctx.Id()

	fmt.Println("Id:", id)
}

func (s *golearnListener) EnterFn(ctx *parser.FnContext) {
	// Register entry point flag.
	if ctx.Id().GetSymbol().GetText() == "main" {
		if !s.mainExists {
			s.mainExists = true
		} else {
			fmt.Println("Conflicting multiple entry points encountered")
		}
	}
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("space test ; @attrib @attrib() fn main(str *myArg, str mySecond) ~> int64 { myStr = 5; }")

	// Create the Lexer
	lexer := parser.NewGolearnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGolearnParser(stream)

	// Finally parse the expression
	listener := newGolearnListener()

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Ensure entry point exists.
	if !listener.mainExists {
		fmt.Println("No entry point found")
	}
}

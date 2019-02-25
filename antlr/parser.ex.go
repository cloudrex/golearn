package main

import (
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type golearnListener struct {
	*parser.BaseGolearnListener
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("space test ;")

	// Create the Lexer
	lexer := parser.NewGolearnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGolearnParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&golearnListener{}, p.Start())
}

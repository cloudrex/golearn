package main

import (
	"fmt"
)

// Parser : Handles parsing and breaking down code into nodes.
type Parser struct {
	tokens []Token
	pos    int
}

func newParser(tokens []Token) *Parser {
	var parser = Parser{}

	parser.pos = 0
	parser.tokens = tokens

	// Append the end of file token to the end.
	parser.tokens = append(parser.tokens, Token{kind: TokenKindEndOfFile})

	return &parser
}

func (parser *Parser) cycle() *Parser {
	parser.next()

	return parser
}

func (parser *Parser) err(message string) error {
	return fmt.Errorf("[At token position %v] %v", parser.pos, message)
}

func (parser *Parser) peek() Token {
	if parser.pos >= len(parser.tokens) {
		return Token{kind: TokenKindEndOfFile}
	}

	return parser.tokens[parser.pos+1]
}

func (parser *Parser) next() Token {
	// Stop if on the last token.
	if parser.pos+1 < len(parser.tokens) {
		parser.pos++
	}

	return parser.tokens[parser.pos]
}

func (parser *Parser) get() Token {
	return parser.tokens[parser.pos]
}

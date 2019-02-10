package main

import "fmt"

// Parser : Handles parsing and breaking down code into nodes.
type Parser struct {
	tokens     []Token
	pos        int
	syncTarget *Parser
}

// Initializes a new parser.
func newParser(tokens []Token) *Parser {
	var parser = Parser{}

	parser.pos = 0
	parser.tokens = tokens

	// Append the end of file token to the end.
	parser.tokens = append(parser.tokens, Token{kind: TokenKindEndOfFile})

	return &parser
}

// Increment parser position.
func (parser *Parser) consume() *Parser {
	parser.next()

	return parser
}

// Creates an error with parser metadata.
func (parser *Parser) err(message string) error {
	return fmt.Errorf("[At token position %v] %v", parser.pos, message)
}

// Creates and displays a fatal error with parser metadata. Stops the application.
func (parser *Parser) fatal(message string) {
	panic(parser.err(message))
}

// Retrieve the next token in the list without changing the parser's position.
func (parser *Parser) peek() Token {
	if parser.pos >= len(parser.tokens) {
		return Token{kind: TokenKindEndOfFile}
	}

	return parser.tokens[parser.pos+1]
}

// Traverse the token list until the specified token kind is found without changing parser's position.
func (parser *Parser) peekUntil(kind TokenKind) []Token {
	var tokens []Token

	for token := parser.get(); token.kind != kind; token = parser.next() {
		if token.kind == TokenKindEndOfFile {
			parser.fatal(fmt.Sprintf("Unexpected end of input tokens, expecting token kind: %v", kind))
		}

		tokens = append(tokens, token)
	}

	return tokens
}

// Traverse the token list until the specified token kind is found.
func (parser *Parser) until(kind TokenKind) []Token {
	tokens := parser.peekUntil(kind)

	parser.teleport(len(tokens))

	return tokens
}

// Verifies that current position is within bounds, otherwise relocates position to corresponding position.
func (parser *Parser) bounds() *Parser {
	if parser.pos >= len(parser.tokens) {
		parser.pos = len(parser.tokens)
	} else if parser.pos < 0 {
		parser.pos = 0
	}

	return parser
}

// Retrieve the next token in the list. Increments parser position.
func (parser *Parser) next() Token {
	parser.navigate(1)

	return parser.tokens[parser.pos]
}

// Sync local position upon remote parser's position change.
func (parser *Parser) sync(target *Parser) *Parser {
	parser.syncTarget = target

	return parser
}

// Remove attached sync parser to stop tracking position.
func (parser *Parser) unlink() *Parser {
	parser.syncTarget = nil

	return parser
}

// Changes the relative position of the parser.
func (parser *Parser) navigate(deltaPos int) *Parser {
	parser.teleport(parser.pos + deltaPos)

	return parser
}

// Changes the absolute position of the parser.
func (parser *Parser) teleport(pos int) *Parser {
	parser.pos = pos

	// Reset position bounds if applicable.
	parser.bounds()

	// Report change to attached sync target if applicable.
	if parser.syncTarget != nil {
		parser.syncTarget.teleport(parser.pos)
	}

	return parser
}

// Retrieve the token located at the current position.
func (parser *Parser) get() Token {
	return parser.tokens[parser.pos]
}

// Clones the parser along with it's current state.
func (parser Parser) derive() Parser {
	return parser
}

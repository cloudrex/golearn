package parser

import (
	"fmt"
	"golearn/scanner"
	"math"
)

// Parser : Handles parsing and breaking down code into nodes.
type Parser struct {
	tokens     []scanner.Token
	pos        int
	syncTarget *Parser
}

// NewParser : Initializes a new parser.
func NewParser(tokens []scanner.Token) *Parser {
	var parser = Parser{}

	parser.pos = 0
	parser.tokens = tokens

	// Append the end of file token to the end.
	parser.tokens = append(parser.tokens, scanner.Token{Kind: scanner.TokenKindEndOfFile})

	return &parser
}

// GetPos : Retrieve the parser's position.
func (parser *Parser) GetPos() int {
	return parser.pos
}

// Consume : Increment parser position.
func (parser *Parser) Consume() *Parser {
	parser.Next()

	return parser
}

// Err : Creates an error with parser metadata.
func (parser *Parser) Err(message string) error {
	return fmt.Errorf("[At token position %v] %v", parser.pos, message)
}

// Fatal : Creates and displays a fatal error with parser metadata. Stops the application.
func (parser *Parser) Fatal(message string) {
	panic(parser.Err(message))
}

// Peek : Retrieve the next token in the list without changing the parser's position.
func (parser *Parser) Peek() scanner.Token {
	return parser.PeekX(1)
}

// PeekX : Retrieve the X token in the list without changing the parser's position.
func (parser *Parser) PeekX(pos int) scanner.Token {
	absPos := int(math.Abs(float64(pos)))

	if parser.pos+absPos >= len(parser.tokens) {
		return scanner.Token{Kind: scanner.TokenKindEndOfFile}
	}

	return parser.tokens[parser.pos+absPos]
}

// PeekUntil : Traverse the token list until the specified token kind is found without changing parser's position.
func (parser *Parser) PeekUntil(kind scanner.TokenKind) []scanner.Token {
	var tokens []scanner.Token

	for token := parser.Get(); token.Kind != kind; token = parser.Next() {
		if token.Kind == scanner.TokenKindEndOfFile {
			parser.Fatal(fmt.Sprintf("Unexpected end of input tokens, expecting token kind: %v", kind))
		}

		tokens = append(tokens, token)
	}

	return tokens
}

// Until : Traverse the token list until the specified token kind is found.
func (parser *Parser) Until(kind scanner.TokenKind) []scanner.Token {
	tokens := parser.PeekUntil(kind)

	parser.Navigate(len(tokens))

	return tokens
}

// Bounds : Verifies that current position is within bounds, otherwise relocates position to corresponding position.
func (parser *Parser) Bounds() *Parser {
	if parser.pos >= len(parser.tokens) {
		parser.pos = len(parser.tokens) - 1
	} else if parser.pos < 0 {
		parser.pos = 0
	}

	return parser
}

// Next : Retrieve the next token in the list. Increments parser position.
func (parser *Parser) Next() scanner.Token {
	parser.Navigate(1)

	return parser.tokens[parser.pos]
}

// Link : Sync remote target's position upon local position change.
func (parser *Parser) Link(target *Parser) *Parser {
	parser.syncTarget = target

	return parser
}

// Unlink : Remove attached sync parser to stop tracking position.
func (parser *Parser) Unlink() *Parser {
	parser.syncTarget = nil

	return parser
}

// Navigate : Changes the relative position of the parser.
func (parser *Parser) Navigate(deltaPos int) *Parser {
	parser.Teleport(parser.pos + deltaPos)

	return parser
}

// Teleport : Changes the absolute position of the parser.
func (parser *Parser) Teleport(pos int) *Parser {
	parser.pos = pos

	// Reset position bounds if applicable.
	parser.Bounds()

	// Report change to attached sync target if applicable.
	if parser.syncTarget != nil {
		parser.syncTarget.Teleport(parser.pos)
	}

	return parser
}

// Get : Retrieve the token located at the current position.
func (parser *Parser) Get() scanner.Token {
	return parser.tokens[parser.pos]
}

// Derive : Clones the parser along with it's current state.
func (parser *Parser) Derive() *Parser {
	return &Parser{pos: parser.pos, tokens: parser.tokens}
}

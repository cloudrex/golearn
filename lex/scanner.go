package lex

import "regexp"

// Scanner : Performs lexical analysis and breaks up input into tokens.
type Scanner struct {
	input string
	pos   int
}

// NextChar : Retrieve the next character from input and advance the position counter.
func (scanner *Scanner) NextChar() string {
	if scanner.pos+1 >= len(scanner.input) {
		return "EOF"
	}

	scanner.pos++

	return scanner.input[scanner.pos]
}

// Next : Process the next token from input.
func (scanner *Scanner) Next() Token {
	for _, chr := range scanner.input {
		token := Token{Value: chr}

		if IsWhitespace(chr) {
			return Token{Kind: TokenKindWhitespace, Value: chr}
		}
	}
}

// IsWhitespace : Determine if input character is a whitespace character.
func IsWhitespace(input string) bool {
	return regexp.MustCompile("\\s").MatchString(input)
}

// NewScanner : Initialize a new Scanner.
func NewScanner(input string) *Scanner {
	return &Scanner{input: input, pos: 0}
}

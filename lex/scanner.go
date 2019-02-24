package lex

import (
	"regexp"
)

// EOF : Represents the end-of-file of input.
const EOF = "EOF"

// Scanner : Performs lexical analysis and breaks up input into tokens.
type Scanner struct {
	input string
	pos   int
}

// NextChar : Retrieve the next character from input and advance the position counter.
func (scanner *Scanner) NextChar() string {
	if scanner.pos+1 >= len(scanner.input) {
		return EOF
	}

	scanner.pos++

	return string(scanner.input[scanner.pos])
}

// Get : Retrieve the character at the current input position.
func (scanner *Scanner) Get() string {
	return string(scanner.input[scanner.pos])
}

// HasNext : Whether the token iterator can continue.
func (scanner *Scanner) HasNext() bool {
	return scanner.pos < len(scanner.input)-1
}

// Next : Process the next token from input.
func (scanner *Scanner) Next() string {
	var token string

	for char := scanner.Get(); char != EOF; char = scanner.NextChar() {
		if IsWhitespaceChar(char) { // Ignore whitespace.
			continue
		} else if IsIdentifierChar(char) {
			token += char

			// Collect identifier.
			for IsIdentifierChar(scanner.NextChar()) {
				token += scanner.Get()
			}

			return token
		} else { // Unexpected character.
			scanner.Fatal("Unexpected character")
		}
	}

	return token
}

// Scan : Process all input tokens.
func (scanner *Scanner) Scan() []string {
	var tokens []string

	// Append first token.
	tokens = append(tokens, scanner.Get())

	for scanner.HasNext() {
		tokens = append(tokens, scanner.Next())
	}

	return tokens
}

// IsIdentifierChar : Determine if input character is part of an identifier.
func IsIdentifierChar(input string) bool {
	return input != EOF && regexp.MustCompile("[_a-zA-Z]").MatchString(input)

}

// IsWhitespaceChar : Determine if input character is a whitespace character.
func IsWhitespaceChar(input string) bool {
	return input != EOF && regexp.MustCompile("\\s").MatchString(input)
}

// NewScanner : Initialize a new Scanner.
func NewScanner(input string) *Scanner {
	return &Scanner{input: input, pos: 0}
}

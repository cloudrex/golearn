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
func (scn *Scanner) NextChar() string {
	char := scn.PeekChar()

	scn.pos++

	return char
}

// ConsumeX : Consume the next specified amount of characters from the input.
func (scn *Scanner) ConsumeX(amount int) *Scanner {
	for i := 0; i < amount; i++ {
		scn.NextChar()
	}

	return scn
}

// Consume : Consume the next character from the input.
func (scn *Scanner) Consume() *Scanner {
	scn.NextChar()

	return scn
}

// bounds : Ensure position is inside valid bounds.
func (scn *Scanner) bounds() *Scanner {
	if scn.pos >= len(scn.input) {
		scn.pos = len(scn.input) - 1
	} else if scn.pos < 0 {
		scn.pos = 0
	}

	return scn
}

// Navigate : Change the scanner's relative position.
func (scn *Scanner) Navigate(pos int) *Scanner {
	scn.Teleport(scn.pos + pos)

	return scn
}

// Teleport : Set the scanner's absolute position.
func (scn *Scanner) Teleport(pos int) *Scanner {
	scn.pos = pos

	// Ensure bounds are valid.
	scn.bounds()

	return scn
}

// PeekChar : Retrieve the next character from input without changing the position counter.
func (scn *Scanner) PeekChar() string {
	if scn.pos+1 >= len(scn.input) {
		return EOF
	}

	return string(scn.input[scn.pos+1])
}

// Get : Retrieve the character at the current input position.
func (scn *Scanner) Get() string {
	return string(scn.input[scn.pos])
}

// HasNext : Whether the value iterator may continue.
func (scn *Scanner) HasNext() bool {
	return scn.pos < len(scn.input)-1
}

// Next : Process the next value from input.
func (scn *Scanner) Next() string {
	var buffer string

	// TODO: Inner loops (until) need to check for end-of-file. Implement Until() method similar to codegen's.
	for char := scn.Get(); char != EOF; char = scn.NextChar() {
		if IsWhitespaceChar(char) { // Ignore whitespace.
			continue
		} else if char == "\"" { // String literal.
			for next := scn.NextChar(); next != "\"" && next != EOF; next = scn.NextChar() {
				buffer += scn.Get()
			}

			// Consume '"'.
			scn.NextChar()

			return "\"" + buffer + "\""
		} else if char == "-" && scn.PeekChar() == ">" { // Function return type symbol.
			// Consume both tokens.
			scn.ConsumeX(2)

			return "->"
		} else if IsNumericChar(char) || char == "." {
			if char == "." {
				if IsNumericChar(scn.PeekChar()) {
					buffer += "."

					continue
				}

				scn.Fatal("Unexpected character")
			}

			buffer += char

			decimalFlag := false

			for next := scn.NextChar(); IsNumericChar(next) || next == "." && !decimalFlag; next = scn.NextChar() {
				if scn.Get() == "." {
					if !decimalFlag {
						decimalFlag = true
					} else { // Decimal indicator '.' appearing twice
						scn.Fatal("Unexpected character")
					}
				}

				buffer += scn.Get()
			}

			return buffer
		} else if IsIdentifierChar(char) {
			buffer += char

			// Collect identifier.
			for IsIdentifierChar(scn.NextChar()) {
				buffer += scn.Get()
			}

			return buffer
		} else { // Unexpected character.
			scn.Fatal("Unexpected character")
		}
	}

	return buffer
}

// Scan : Process all input tokens.
func (scn *Scanner) Scan() []string {
	var tokens []string

	for scn.HasNext() {
		tokens = append(tokens, scn.Next())
	}

	return tokens
}

// IsIdentifierChar : Determine if input character is part of an identifier.
func IsIdentifierChar(input string) bool {
	return input != EOF && regexp.MustCompile("[_a-zA-Z]").MatchString(input)
}

// IsNumericChar : Determine if input character is a number.
func IsNumericChar(input string) bool {
	return input != EOF && regexp.MustCompile("[0-9]").MatchString(input)
}

// IsWhitespaceChar : Determine if input character is a whitespace character.
func IsWhitespaceChar(input string) bool {
	return input != EOF && regexp.MustCompile("\\s").MatchString(input)
}

// NewScanner : Initialize a new Scanner.
func NewScanner(input string) *Scanner {
	return &Scanner{input: input, pos: 0}
}

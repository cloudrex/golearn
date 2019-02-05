package lexer

import (
	"regexp"
	"strconv"
	"io"
)

// Scanner : Breaks up input code into tokens.
type Scanner struct {
	input            string
	pos              int
	identifierBuffer string
	numberStrBuffer  string
	numberValBuffer  float64
}

func (scanner Scanner) nextChar() string {
	scanner.pos++

	return string(scanner.input[scanner.pos])
}

func (scanner Scanner) getToken() Token {
	lastChar := " "
	whitespacePattern := regexp.MustCompile(`\s`)
	identifierPattern := regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9]*`)
	numberPattern := regexp.MustCompile(`[0-9.]+`)

	// Ignore any whitespace.
	for whitespacePattern.MatchString(lastChar) {
		scanner.pos++
		lastChar = string(scanner.input[scanner.pos])
	}

	if identifierPattern.MatchString(lastChar) {
		scanner.identifierBuffer = lastChar

		for identifierPattern.MatchString(lastChar) {
			scanner.identifierBuffer += lastChar
			lastChar = scanner.nextChar()
		}

		if scanner.identifierBuffer == "fn" {
			return TokenFn
		} else if scanner.identifierBuffer == "extern" {
			return TokenExtern
		}

		return TokenIdentifier
	} else if numberPattern.MatchString(lastChar) {
		// TODO: Should be do-while
		for numberPattern.MatchString(lastChar) {
			scanner.numberStrBuffer += lastChar
			lastChar = scanner.nextChar()
		}

		numberValBuffer, err := strconv.ParseFloat(scanner.numberStrBuffer, 64)

		if err != nil {
			panic(err)
		}

		scanner.numberValBuffer = numberValBuffer

		return TokenNumber
	} else if lastChar == "#" {
		// TODO
	} else if lastChar == io.EOF {
		return TokenEOF
	}

	// Otherwise, return the unknown token
	return TokenUnknown
}

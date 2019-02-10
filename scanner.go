package main

import (
	"regexp"
	"strings"
	"text/scanner"
)

// Scanner : Represents the lexical token scanner utility.
type Scanner struct {
}

func isIdentifier(input string) bool {
	return regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$").MatchString(input)
}

func (sc *Scanner) scan(input string) []Token {
	var scan scanner.Scanner
	var tokens []Token

	scan.Init(strings.NewReader(input))
	scan.Filename = "example"

	// Continue until reaching end of file.
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		// Recognize and append tokens accordingly as they come.
		text := scan.TokenText()
		token := TokenKindUnknown

		if text == "fn" { // Function definition keyword 'fn'.
			token = TokenKindFn
		} else if text == "exit" { // Exit keyword 'exit'.
			token = TokenKindExit
		} else if text == "+" { // Addition operator '+'.
			token = TokenKindAddOp
		} else if isIdentifier(text) { // Identifier.
			token = TokenKindIdentifier
		} else if text == "(" { // Parentheses start '('.
			token = TokenKindParenStart
		} else if text == ")" { // Parentheses end ')'.
			token = TokenKindParenEnd
		} else if text == "{" { // Block start '{'.
			token = TokenKindBlockStart
		} else if text == "}" { // Block end '{'.
			token = TokenKindBlockEnd
		} else if text == ";" { // Semi-colon ';'.
			token = TokenKindSemiColon
		}

		tokens = append(tokens, Token{token, text})
	}

	return tokens
}
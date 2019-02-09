package main

import (
	"regexp"
	"strings"
	"text/scanner"
)

func scan(input string) []Token {
	var scan scanner.Scanner
	var tokens []Token
	var identifierExpr = regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$")

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
		} else if identifierExpr.MatchString(text) { // Identifier.
			token = TokenKindIdentifier
		} else if text == "(" { // Parentheses start '('.
			token = TokenKindParenStart
		} else if text == ")" { // Parentheses end ')'.
			token = TokenKindParenEnd
		}

		tokens = append(tokens, Token{token, text})
	}

	return tokens
}

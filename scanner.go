package main

import (
	"strings"
	"text/scanner"
)

func scan(input string) []Token {
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
		}

		tokens = append(tokens, Token{token, text})
	}

	return tokens
}
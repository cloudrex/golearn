package scanner

import (
	"regexp"
	"strings"
	"text/scanner"
)

// Scanner : Represents the lexical token scanner utility.
type Scanner struct {
}

// IsIdentifier : Determine if input is an identifier token.
func IsIdentifier(input string) bool {
	return regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$").MatchString(input)
}

// IsNumeric : Determine if input is a numeric decimal constant.
func IsNumeric(input string) bool {
	return regexp.MustCompile("^[0-9]*(?:\\.[0-9]+)?$").MatchString(input)
}

// Scan : Scan and break up input into lexical tokens.
func (sc *Scanner) Scan(input string) []Token {
	var scan scanner.Scanner
	var tokens []Token

	scan.Init(strings.NewReader(input))
	scan.Filename = "example"

	// Continue until reaching end of file.
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		// Recognize and append tokens accordingly as they come.
		text := scan.TokenText()
		kind := TokenKindUnknown

		if text == "fn" { // Function definition keyword 'fn'.
			kind = TokenKindFn
		} else if text == "exit" { // Exit keyword 'exit'.
			kind = TokenKindExit
		} else if text == "+" { // Addition operator '+'.
			kind = TokenKindAddOp
		} else if IsIdentifier(text) { // Identifier.
			kind = TokenKindIdentifier
		} else if IsNumeric(text) { // Numeric decimal constant.
			kind = TokenKindNumber
		} else if text == "(" { // Parentheses start '('.
			kind = TokenKindParenStart
		} else if text == ")" { // Parentheses end ')'.
			kind = TokenKindParenEnd
		} else if text == "{" { // Block start '{'.
			kind = TokenKindBlockStart
		} else if text == "}" { // Block end '{'.
			kind = TokenKindBlockEnd
		} else if text == ";" { // Semi-colon ';'.
			kind = TokenKindSemiColon
		} else if text == ":" { // Colon ':'.
			kind = TokenKindColon
		} else if text == "=" { // Equal sign '='.
			kind = TokenKindEqualSign
		}

		tokens = append(tokens, Token{Kind: kind, Value: text})
	}

	return tokens
}

// FindTokenIndex : Find a token's index location in a token array by it's kind. Returns matching first element found or '-1' if not found.
func FindTokenIndex(set []Token, kind TokenKind) int {
	index := -1

	for i := 0; i < len(set); i++ {
		if set[i].Kind == kind {
			index = i

			break
		}
	}

	return index
}

// FindToken : Find a token in a token array by it's kind. Returns matching first element found or EOF token if not found.
func FindToken(set []Token, kind TokenKind) Token {
	index := FindTokenIndex(set, kind)

	if index != -1 {
		return set[index]
	}

	return Token{Kind: TokenKindEndOfFile}
}

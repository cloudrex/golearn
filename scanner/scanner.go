package scanner

import (
	"regexp"
	"strings"
	"text/scanner"
)

// Scanner : Represents the lexical token scanner utility.
type Scanner struct {
	//
}

// IsIdentifier : Determine if input is an identifier token.
func IsIdentifier(input string) bool {
	return regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*$").MatchString(input)
}

// IsFloatLiteral : Determine if input is a numeric float literal.
func IsFloatLiteral(input string) bool {
	return regexp.MustCompile("^[0-9]+\\.[0-9]+$").MatchString(input)
}

// IsIntLiteral : Determine if input is a numeric integer literal.
func IsIntLiteral(input string) bool {
	return regexp.MustCompile("^[0-9]+$").MatchString(input)
}

// IsStringLiteral : Determine if input is a string literal.
func IsStringLiteral(input string) bool {
	return regexp.MustCompile("^\"[^\\\"]*\"$").MatchString(input)
}

// IsCharLiteral : Determine if input is a charater literal.
func IsCharLiteral(input string) bool {
	return regexp.MustCompile("^'[^\\']{0,1}'$").MatchString(input)
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
			kind = TokenKindFnKeyword
		} else if text == "exit" { // Exit keyword 'exit'.
			kind = TokenKindBreak
		} else if text == "extern" { // External definition keyword 'extern'.
			kind = TokenKindExternKeyword
		} else if text == "string" { // String type keyword 'string'.
			kind = TokenKindStringKeyword
		} else if text == "char" {
			kind = TokenKindCharKeyword
		} else if text == "int" { // Integer-32 short-hand type keyword 'int'.
			kind = TokenKindIntKeyword
		} else if text == "float" { // Float type keyword 'float'.
			kind = TokenKindFloatKeyword
		} else if text == "+" { // Addition operator '+'.
			kind = TokenKindAddOp
		} else if text == "-" { // Substraction operator '-'.
			kind = TokenKindSubOp
		} else if text == "*" { // Multiplication operator '*'.
			kind = TokenKindMultOp
		} else if text == "/" { // Division operator '/'.
			kind = TokenKindDivOp
		} else if text == "%" { // Modulos operator '%'.
			kind = TokenKindModulusOp
		} else if IsIdentifier(text) { // Identifier.
			kind = TokenKindIdentifier
		} else if IsStringLiteral(text) { // String literal.
			kind = TokenKindStringLiteral
		} else if IsCharLiteral(text) { // Character literal.
			kind = TokenKindCharLiteral
		} else if IsIntLiteral(text) { // Integer value literal.
			kind = TokenKindIntegerLiteral
		} else if IsFloatLiteral(text) { // Float value literal.
			kind = TokenKindFloatLiteral
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
		} else if text == ">" {
			kind = TokenKindGreaterThanOp
		} else if text == "<" {
			kind = TokenKindLessThanOp
		} else if text == "!" {
			kind = TokenKindNotOp
		} else if text == "&" {
			kind = TokenKindDereferenceOp
		} else if text == "@" {
			kind = TokenKindAttribute
		} else if text == "and" {
			kind = TokenKindAndOp
		} else if text == "or" {
			kind = TokenKindOrOp
		} else if text == "xor" {
			kind = TokenKindXOrOp
		} else if text == "if" {
			kind = TokenKindIfKeyword
		} else if text == "else" {
			kind = TokenKindElseKeyword
		} else if text == "for" {
			kind = TokenKindForKeyword
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

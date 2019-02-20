package lex

import (
	"regexp"
	"strings"
	"text/scanner"
)

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

// Tokenize : Scan and break up input into lexical tokens.
func Tokenize(input string) []Token {
	var scan scanner.Scanner
	var tokens []Token

	scan.Init(strings.NewReader(input))
	scan.Filename = "example"

	// Continue until reaching end of file.
	for tok := scan.Scan(); tok != scanner.EOF; tok = scan.Scan() {
		// Recognize and append tokens accordingly as they come.
		text := scan.TokenText()
		kind := TokenKindUnknown

		if text == KeywordFn { // Function definition keyword 'fn'.
			kind = TokenKindFnKeyword
		} else if text == "exit" { // Exit keyword 'exit'.
			kind = TokenKindBreak
		} else if text == KeywordBreak { // External definition keyword 'extern'.
			kind = TokenKindExternKeyword
		} else if text == KeywordExtern { // String type keyword 'string'.
			kind = TokenKindStringKeyword
		} else if text == KeywordChar {
			kind = TokenKindCharKeyword
		} else if text == KeywordInt { // Integer-32 short-hand type keyword 'int'.
			kind = TokenKindIntKeyword
		} else if text == KeywordFloat { // Float type keyword 'float'.
			kind = TokenKindFloatKeyword
		} else if text == OperatorAdd { // Addition operator '+'.
			kind = TokenKindAddOp
		} else if text == OperatorSub { // Substraction operator '-'.
			kind = TokenKindSubOp
		} else if text == OperatorMult { // Multiplication operator '*'.
			kind = TokenKindMultOp
		} else if text == OperatorDiv { // Division operator '/'.
			kind = TokenKindDivOp
		} else if text == OperatorModulus { // Modulos operator '%'.
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
		} else if text == OperatorGreaterThan { // Greater than operator '>'.
			kind = TokenKindGreaterThanOp
		} else if text == OperatorLessThan { // Less than operator '>'.
			kind = TokenKindLessThanOp
		} else if text == OperatorNot { // NOT operator '!'.
			kind = TokenKindNotOp
		} else if text == KeywordTrue { // True boolean 'true'.
			kind = TokenKindTrueBool
		} else if text == KeywordFalse { // False boolean 'false'.
			kind = TokenKindFalseBool
		} else if text == OperatorDereference { // Deference operator '&'.
			kind = TokenKindDereferenceOp
		} else if text == OperatorAttribute { // Attribute operator '@'.
			kind = TokenKindAttribute
		} else if text == OperatorAnd { // AND operator keyword 'and'.
			kind = TokenKindAndOp
		} else if text == OperatorOr { // OR operator keyword 'or'.
			kind = TokenKindOrOp
		} else if text == OperatorXOr { // XOR Operator keyword 'xor'.
			kind = TokenKindXOrOp
		} else if text == KeywordIf { // If keyword 'if'
			kind = TokenKindIfKeyword
		} else if text == KeywordElse { // Else keyword 'else'
			kind = TokenKindElseKeyword
		} else if text == KeywordFor { // For keyword 'for'
			kind = TokenKindForKeyword
		} else if text == KeywordWhile { // While keyword 'while'
			kind = TokenKindWhileKeyword
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

package main

// TokenKind : Represents a specific token kind or type.
type TokenKind int

const (
	// TokenKindEndOfFile : Represents the end of the input.
	TokenKindEndOfFile TokenKind = -1

	// TokenKindUnknown : Represents the unknown or invalid token.
	TokenKindUnknown TokenKind = 0

	// TokenKindFn : Represents the function definition keyword 'fn'.
	TokenKindFn TokenKind = 1

	// TokenKindExit : Represents the exit keyword 'exit'.
	TokenKindExit TokenKind = 2

	// TokenKindAddOp : Represents the addition operator '+'.
	TokenKindAddOp TokenKind = 3

	// TokenKindIdentifier : Represents an identifier.
	TokenKindIdentifier TokenKind = 4

	// TokenKindParenStart : Represents the opening token of a parentheses expression '('.
	TokenKindParenStart TokenKind = 5

	// TokenKindParenEnd : Represents the closing token of a parentheses expression ')'.
	TokenKindParenEnd TokenKind = 6

	// TokenKindComma : Represents ','.
	TokenKindComma TokenKind = 7
)

// Token : Represents a single token.
type Token struct {
	kind  TokenKind
	value string
}

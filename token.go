package main

// TokenKind : Represents a specific token kind or type.
type TokenKind int

const (
	// TokenKindUnknown : Represents the unknown or invalid token.
	TokenKindUnknown TokenKind = 0

	// TokenKindFn : Represents the function definition keyword 'fn'.
	TokenKindFn TokenKind = 1

	// TokenKindExit : Represents the exit keyword 'exit'.
	TokenKindExit TokenKind = 2

	// TokenKindAddOp : Represents the addition operator '+'.
	TokenKindAddOp TokenKind = 3
)

// Token : Represents a single token.
type Token struct {
	kind TokenKind
	value string
}
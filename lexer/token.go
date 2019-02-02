package lexer

// Token : Represents a single token.
type Token int

const (
	// TokenEOF : Represents the end-of-file token.
	TokenEOF Token = -1

	// TokenFn : Represents the function definition token.
	TokenFn Token = -2

	// TokenExtern : Represents the external definition token.
	TokenExtern Token = -3

	// TokenIdentifier : Represents the generic identifier token.
	TokenIdentifier Token = -4

	// TokenNumber : Represents the number token.
	TokenNumber Token = -5

	// TokenUnknown : Represents an unknown and/or invalid token.
	TokenUnknown Token = -6
)

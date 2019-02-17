package lex

// Token : Represents a single token.
type Token struct {
	Kind  TokenKind
	Value string
}

// TokenKind : Represents a specific token kind or type.
type TokenKind int

const (
	// TokenKindEndOfFile : Represents the end of the input.
	TokenKindEndOfFile TokenKind = -1

	// TokenKindUnknown : Represents the unknown or invalid token.
	TokenKindUnknown TokenKind = 0

	// TokenKindFnKeyword : Represents the function definition keyword 'fn'.
	TokenKindFnKeyword TokenKind = 1

	// TokenKindBreak : Represents the loop-break keyword 'break'.
	TokenKindBreak TokenKind = 2

	// TokenKindAddOp : Represents the binary addition operator '+'.
	TokenKindAddOp TokenKind = 3

	// TokenKindIdentifier : Represents an identifier.
	TokenKindIdentifier TokenKind = 4

	// TokenKindParenStart : Represents the opening token of a parentheses expression '('.
	TokenKindParenStart TokenKind = 5

	// TokenKindParenEnd : Represents the closing token of a parentheses expression ')'.
	TokenKindParenEnd TokenKind = 6

	// TokenKindComma : Represents ','.
	TokenKindComma TokenKind = 7

	// TokenKindBlockStart : Represents the start of a statement block '{'.
	TokenKindBlockStart TokenKind = 8

	// TokenKindBlockEnd : Represents the end of a statement block '{'.
	TokenKindBlockEnd TokenKind = 9

	// TokenKindSemiColon : Represents the semi-colon token ';'.
	TokenKindSemiColon TokenKind = 10

	// TokenKindEqualSign : Represents the equal sign token '='.
	TokenKindEqualSign TokenKind = 11

	// TokenKindColon : Represents the colon token ':'.
	TokenKindColon TokenKind = 12

	// TokenKindIntegerLiteral : Represents the integer literal value token.
	TokenKindIntegerLiteral TokenKind = 13

	// TokenKindFloatLiteral : Represents the float literal value token.
	TokenKindFloatLiteral TokenKind = 14

	// TokenKindStringLiteral : Represents the string literal value token.
	TokenKindStringLiteral TokenKind = 15

	// TokenKindStringKeyword : Represents the string type keyword 'string'.
	TokenKindStringKeyword TokenKind = 16

	// TokenKindIntKeyword : Represents the integer-32 short-hand type keyword 'int'.
	TokenKindIntKeyword TokenKind = 17

	// TokenKindFloatKeyword : Represents the float type keyword 'float'.
	TokenKindFloatKeyword TokenKind = 18

	// TokenKindExternKeyword : Represents the external definition keyword 'extern'.
	TokenKindExternKeyword TokenKind = 19

	// TokenKindCharKeyword : Represents the character type keyword 'char'.
	TokenKindCharKeyword TokenKind = 20

	// TokenKindCharLiteral : Represents the character literal value token.
	TokenKindCharLiteral TokenKind = 21

	// TokenKindSubOp : Represents the binary substraction operator '-'.
	TokenKindSubOp TokenKind = 22

	// TokenKindMultOp : Represents the binary multiplication operator '*'.
	TokenKindMultOp TokenKind = 23

	// TokenKindDivOp : Represents the binary division operator '/'.
	TokenKindDivOp TokenKind = 24

	// TokenKindModulusOp : Represents the binary modulus operator '%'.
	TokenKindModulusOp TokenKind = 25

	// TokenKindGreaterThanOp : Represents the binary logical greater than operator '>'.
	TokenKindGreaterThanOp TokenKind = 26

	// TokenKindLessThanOp : Represents the binary logical less than operator '<'.
	TokenKindLessThanOp TokenKind = 27

	// TokenKindNotOp : Represents the unary logical NOT oerator '!'.
	TokenKindNotOp TokenKind = 28

	// TokenKindDereferenceOp : Represents the unary de-reference operator '&'
	TokenKindDereferenceOp TokenKind = 29

	// TokenKindAttribute : Represents the attribute prefix '@'.
	TokenKindAttribute TokenKind = 30

	// TokenKindAndOp : Represents the logical AND operator 'and'.
	TokenKindAndOp TokenKind = 31

	// TokenKindOrOp : Represents the logical OR operator 'or'.
	TokenKindOrOp TokenKind = 32

	// TokenKindXOrOp : Represents the logical XOR operator 'xor'.
	TokenKindXOrOp TokenKind = 33

	// TokenKindIfKeyword : Represents the if keyword 'if'.
	TokenKindIfKeyword TokenKind = 34

	// TokenKindElseKeyword : Represents the else keyword 'else'.
	TokenKindElseKeyword TokenKind = 35

	// TokenKindForKeyword : Represents the for-loop keyword 'for'.
	TokenKindForKeyword TokenKind = 36

	// TokenKindImportKeyword : Represents the import statement keyword 'import'.
	TokenKindImportKeyword TokenKind = 37
)

package lex

// Entity : Represents common entities such as operators and keywords.
type Entity = string

const (
	// KeywordIf : Represents the if keyword.
	KeywordIf Entity = "if"

	// KeywordElse : Represents the else keyword.
	KeywordElse Entity = "else"

	// KeywordFor : Represents for-loop keyword.
	KeywordFor Entity = "for"

	// KeywordBreak : Represents the loop break keyword.
	KeywordBreak Entity = "break"

	// KeywordContinue : Represents the loop continue keyword.
	KeywordContinue Entity = "continue"

	// KeywordFn : Represents the function definition keyword.
	KeywordFn Entity = "fn"

	// KeywordImport : Represents the import statement keyword.
	KeywordImport Entity = "import"

	// KeywordExtern : Represents the external function definition keyword.
	KeywordExtern Entity = "extern"

	// KeywordChar : Represents the character type keyword.
	KeywordChar Entity = "char"

	// KeywordInt : Represents the integer-32 type keyword.
	KeywordInt Entity = "int"

	// KeywordString : Represents the string type keyword.
	KeywordString Entity = "string"
)

const (
	// OperatorAnd : Represents the logical AND operator.
	OperatorAnd = "and"

	// OperatorOr : Represents the logical OR operator.
	OperatorOr = "or"

	// OperatorXOr : Represents the logical XOR operator.
	OperatorXOr = "xor"

	// OperatorAdd : Represents the addition operator.
	OperatorAdd = "+"

	// OperatorSub : Represents the subtraction operator.
	OperatorSub = "-"

	// OperatorMult : Represents the multiplication operator.
	OperatorMult = "*"

	// OperatorDiv : Represents the division operator.
	OperatorDiv = "/"

	// OperatorModulus : Represents the modoulus operator.
	OperatorModulus = "%"

	// OperatorExpontial : Represents the expontial operator.
	OperatorExpontial = "^"

	// OperatorAttribute : Represents the attribute operator.
	OperatorAttribute = "@"
)

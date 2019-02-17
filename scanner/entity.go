package scanner

// Entity : Represents common entities such as operators and keywords.
type Entity = string

const (
	// KeywordIf : Represents the if keyword.
	KeywordIf = "if"

	// KeywordElse : Represents the else keyword.
	KeywordElse = "else"

	// KeywordFor : Represents for-loop keyword.
	KeywordFor = "for"

	// KeywordBreak : Represents the loop break keyword.
	KeywordBreak = "break"

	// KeywordContinue : Represents the loop continue keyword.
	KeywordContinue = "continue"

	// KeywordFn : Represents the function definition keyword.
	KeywordFn = "fn"
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
)

package ast

// OperatorKind : Represents the type of an operator.
type OperatorKind = int

const (
	// OperatorKindAdd : Represents the addition operator type.
	OperatorKindAdd OperatorKind = 0

	// OperatorKindSubstract : Represents the substraction operator type.
	OperatorKindSubstract OperatorKind = 1

	// OperatorKindMultiply : Represents the multiplication operator type.
	OperatorKindMultiply OperatorKind = 1

	// OperatorKindDivide : Represents the division operator type.
	OperatorKindDivide OperatorKind = 1
)

// OperatorAST : Represents an operator AST node.
type OperatorAST struct {
	value string
	kind  OperatorKind
}

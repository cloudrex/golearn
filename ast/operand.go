package ast

// OperandKind : Represents the type of operand used in an operation.
type OperandKind = int

const (
	// OperandKindIdentifier : Represents an identifier operand.
	OperandKindIdentifier OperandKind = 0

	// OperandKindLiteral : Represents a literal value operand.
	OperandKindLiteral OperandKind = 1
)

// OperandAST : Represents an operand AST node.
type OperandAST struct {
	value string
	kind  OperandKind
}

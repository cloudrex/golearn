package ast

// BinaryOperationAST : Represents a binary operation AST node.
type BinaryOperationAST struct {
	leftSide  OperandAST
	rightSide OperandAST
	operator  OperatorAST
}

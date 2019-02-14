package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// VarAssignmentAST : Represents the variable assignment AST node.
type VarAssignmentAST struct {
	variable value.Value
	value    value.Value
}

// GetVar : Retrieve the variable being assigned.
func (node *VarAssignmentAST) GetVar() value.Value {
	return node.variable
}

// GetValue : Retrieve the value being assigned.
func (node *VarAssignmentAST) GetValue() value.Value {
	return node.value
}

// Emit : Emit the AST representation.
func (node *VarAssignmentAST) Emit(block *ir.Block) {
	block.NewStore(node.value, node.variable)
}

package codegen

import (
	"golearn/util"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// VarAssignmentAST : Represents the variable assignment AST node.
type VarAssignmentAST struct {
	variableName string
	value        value.Value
}

// GetVarName : Retrieve the name of the variable being assigned.
func (node *VarAssignmentAST) GetVarName() string {
	return node.variableName
}

// GetValue : Retrieve the value being assigned.
func (node *VarAssignmentAST) GetValue() value.Value {
	return node.value
}

// Emit : Emit the AST representation.
func (node *VarAssignmentAST) Emit(block *ir.Block) {
	target := util.FindAllocaInBlock(block, node.variableName)

	if target == nil {
		// TODO: Should use Parser's Fatal error report function.
		panic("Undeclared variable in assignment named '" + node.variableName + "'")
	}

	block.NewStore(node.value, target)
}

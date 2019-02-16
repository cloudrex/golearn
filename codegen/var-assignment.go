package codegen

import (
	"fmt"
	"golearn/scanner"
	"golearn/util"

	"github.com/llir/llvm/ir/types"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// VarAssignmentAST : Represents the variable assignment AST node.
type VarAssignmentAST struct {
	variableName string
	variableType scanner.VariableType
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

	fmt.Println("Validation: var type:", node.variableType)

	// Verify that types match.
	if node.value.Type() == types.Float && node.variableType != scanner.VariableTypeFloat {
		// TODO: Use Parser's Fatal method.
		panic("Invalid assignment; Mismatching types; Expecting float value")
	}

	// TODO: Add more checks, else if... (int, string, etc.).

	block.NewStore(node.value, target)
}

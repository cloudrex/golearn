package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// VarDeclarationAST : Represents the variable declaration AST node.
type VarDeclarationAST struct {
	name  string
	value string
}

// GetName : Retrieve the name of the variable being declared.
func (node *VarDeclarationAST) GetName() string {
	return node.name
}

// GetValue : Retrieve the value of the variable being declared.
func (node *VarDeclarationAST) GetValue() string {
	return node.value
}

// Create : Emit the AST representation.
func (node *VarDeclarationAST) Create(block *ir.Block) {
	alloc := block.NewAlloca(types.I32)

	alloc.SetName(node.name)
}

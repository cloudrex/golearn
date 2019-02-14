package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// VarDeclarationAST : Represents the variable declaration AST node.
type VarDeclarationAST struct {
	name  string
	value string
	ref   *ir.InstAlloca
}

// GetName : Retrieve the name of the variable being declared.
func (node *VarDeclarationAST) GetName() string {
	return node.name
}

// GetValue : Retrieve the value of the variable being declared.
func (node *VarDeclarationAST) GetValue() string {
	return node.value
}

// GetRef : Retrieve the LLVM allocation reference value of the variable being declared. Returns nil if has node has not been previously emitted.
func (node *VarDeclarationAST) GetRef() *ir.InstAlloca {
	return node.ref
}

// Emit : Emit the AST representation.
func (node *VarDeclarationAST) Emit(block *ir.Block) {
	ref := block.NewAlloca(types.I32)

	ref.SetName(node.name)

	// Apply the ref to the node to allow future retrieval.
	node.ref = ref
}

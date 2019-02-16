package codegen

import (
	"golearn/scanner"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// VariableAST : Represents the variable AST node.
type VariableAST struct {
	name  string
	kind  scanner.VariableType
	value string
	ref   *ir.InstAlloca
}

// GetName : Retrieve the name of the variable being declared.
func (node *VariableAST) GetName() string {
	return node.name
}

// GetValue : Retrieve the value of the variable being declared.
func (node *VariableAST) GetValue() string {
	return node.value
}

// GetRef : Retrieve the LLVM allocation reference value of the variable being declared. Returns nil if has node has not been previously emitted.
func (node *VariableAST) GetRef() *ir.InstAlloca {
	return node.ref
}

// Emit : Emit the AST representation.
func (node *VariableAST) Emit(block *ir.Block) {
	ref := block.NewAlloca(types.Float)

	// Apply variable name.
	ref.SetName(node.name)

	// Apply the ref to the node to allow future retrieval.
	node.ref = ref
}

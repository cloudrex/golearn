package codegen

import (
	"golearn/lex"
	"golearn/util"

	"github.com/llir/llvm/ir/value"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// VariableAST : Represents the variable AST node.
type VariableAST struct {
	name  string
	kind  lex.VariableType
	Value value.Value
	ref   *ir.InstAlloca
}

// GetName : Retrieve the name of the variable being declared.
func (node *VariableAST) GetName() string {
	return node.name
}

// GetValue : Retrieve the value of the variable being declared.
func (node *VariableAST) GetValue() value.Value {
	return node.Value
}

// GetRef : Retrieve the LLVM allocation reference value of the variable being declared. Returns nil if has node has not been previously emitted.
func (node *VariableAST) GetRef() *ir.InstAlloca {
	return node.ref
}

// Emit : Emit the AST representation.
func (node *VariableAST) Emit(block *ir.Block) {
	// Ensure variable is only declared once.
	existing := util.FindAllocaInBlock(block, node.name)

	if existing != nil {
		// TODO: Use Parser's error reporting.
		panic("Cannot redeclare variable. Variable '" + node.name + "' is already declared.")
	}

	ref := block.NewAlloca(types.Float)

	// Apply variable name.
	ref.SetName(node.name)

	// Apply the ref to the node to allow future retrieval.
	node.ref = ref

	// Apply value if applicable.
	if node.Value != nil {
		block.NewStore(node.Value, node.ref)
	}
}

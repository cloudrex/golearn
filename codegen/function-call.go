package codegen

import (
	"github.com/llir/llvm/ir"
)

// FunctionCallAST : Represents a function call block AST node.
type FunctionCallAST struct {
	name string
	args []FunctionArgAST
}

// Emit : Emit the AST representation.
func (node *FunctionCallAST) Emit(block *ir.Block) {
	//call := block.NewCall()
}

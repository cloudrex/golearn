package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// FunctionBodyAST : Represents a function statement block AST node.
type FunctionBodyAST struct {
	statements  []BlockNode
	returnValue value.Value
}

// Emit : Emit the AST representation.
func (node *FunctionBodyAST) Emit(fn *ir.Func) {
	block := fn.NewBlock("entry")

	for i := 0; i < len(node.statements); i++ {
		statement := node.statements[i]

		statement.Emit(block)
	}

	// Apply the return statement. Will return void if the node has no set returnValue (nil).
	block.NewRet(node.returnValue)
}

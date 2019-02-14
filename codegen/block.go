package codegen

import (
	"github.com/llir/llvm/ir"
)

// BlockAST : Represents a statement block AST node.
type BlockAST struct {
	label      string
	statements []BlockNode
}

// Emit : Emit the AST representation.
func (node *BlockAST) Emit(fn *ir.Func) {
	irBlock := fn.NewBlock(node.label)

	for i := 0; i < len(node.statements); i++ {
		statement := node.statements[i]

		statement.Emit(irBlock)
	}
}

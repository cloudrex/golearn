package codegen

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

// BlockAST : Represents a statement block AST node.
type BlockAST struct {
	label      string
	statements []BlockNode
}

// Create : Emit the AST representation.
func (node *BlockAST) Create(fn *ir.Func) {
	fmt.Println("--- STATEMENTS", node)

	irBlock := fn.NewBlock(node.label)

	for i := 0; i < len(node.statements); i++ {
		statement := node.statements[i]

		statement.Create(irBlock)
	}
}

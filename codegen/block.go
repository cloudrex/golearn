package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// FunctionBodyAST : Represents a function statement block AST node.
type FunctionBodyAST struct {
	label      string
	statements []BlockNode
}

// Emit : Emit the AST representation.
func (node *FunctionBodyAST) Emit(fn *ir.Func) {
	block := fn.NewBlock(node.label)

	for i := 0; i < len(node.statements); i++ {
		statement := node.statements[i]

		statement.Emit(block)
	}

	// TODO: The code below adds a 'dummy' block returned, as it is required for function body blocks. In the future, it should return 'void' if no return type/value specified.

	// Test add addition statement.
	addSt := block.NewAdd(constant.NewInt(types.I32, 5), constant.NewInt(types.I32, 5))

	block.NewRet(addSt)
}

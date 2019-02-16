package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

// FunctionCallAST : Represents a function call block AST node.
type FunctionCallAST struct {
	name string
	args []value.Value
}

// Emit : Emit the AST representation.
func (node *FunctionCallAST) Emit(module *ir.Module, block *ir.Block) {
	var fn *ir.Func

	for i := 0; i < len(module.Funcs); i++ {
		item := module.Funcs[i]

		if item.Name() == node.name {
			fn = item

			break
		}
	}

	if fn == nil {
		// TODO: Should use error reporting methods.
		panic("Cannot invoke nonexistent function named '" + node.name + "'")
	}

	block.NewCall(fn, node.args...)
}

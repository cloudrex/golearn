package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// FunctionArgAST : Represents the function argument/parameter node.
type FunctionArgAST struct {
	name      string
	valueType types.Type
}

func (node FunctionArgAST) get() *ir.Param {
	return ir.NewParam(node.name, node.valueType)
}

package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// FunctionArgAST : Represents the function argument/parameter node.
type FunctionArgAST struct {
	name      string
	valueType types.Type
}

func (arg FunctionArgAST) get() *ir.Param {
	return ir.NewParam(arg.name, arg.valueType)
}

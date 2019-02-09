package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
)

// FunctionAST : Represents the function node.
type FunctionAST struct {
	name       string
	args       []*FunctionArgAST
	returnType types.Type
	module     *ir.Module
}

func (function FunctionAST) create() {
	var args []*ir.Param

	for i := 0; i < len(function.args); i++ {
		args[i] = function.args[i].generate()
	}

	function.module.NewFunc(function.name, function.returnType, args...)
}

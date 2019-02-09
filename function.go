package main

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// FunctionAST : Represents the function node.
type FunctionAST struct {
	name       string
	args       []*FunctionArgAST
	returnType types.Type
	module     *ir.Module
	body       *ir.Block
}

func (function FunctionAST) create() {
	var args []*ir.Param

	for i := 0; i < len(function.args); i++ {
		args[i] = function.args[i].get()
	}

	fn := function.module.NewFunc(function.name, function.returnType, args...)

	// Add the body block.
	block := fn.NewBlock("body")

	// Test add addition statement.
	addSt := block.NewAdd(constant.NewInt(types.I32, 5), constant.NewInt(types.I32, 5))

	block.NewRet(addSt)

	// Apply it to the struct.
	function.body = block
}

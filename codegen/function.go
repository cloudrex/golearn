package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// FunctionAST : Represents the function node.
type FunctionAST struct {
	Name       string
	args       []*FunctionArgAST
	returnType types.Type
	body       *ir.Block
}

// Create : Create and apply the function into an LLVM module.
func (function FunctionAST) Create(module *ir.Module) {
	var args []*ir.Param

	for i := 0; i < len(function.args); i++ {
		args[i] = function.args[i].get()
	}

	fn := module.NewFunc(function.Name, function.returnType, args...)

	// Add the body block.
	block := fn.NewBlock("body")

	// Test add addition statement.
	addSt := block.NewAdd(constant.NewInt(types.I32, 5), constant.NewInt(types.I32, 5))

	block.NewRet(addSt)

	// Apply it to the struct.
	function.body = block
}

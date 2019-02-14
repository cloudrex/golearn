package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// FunctionAST : Represents the function node.
type FunctionAST struct {
	Name       string
	args       []FunctionArgAST
	returnType types.Type
	body       *BlockAST
}

// OldCreate : Create and apply the function into an LLVM module.
func (fn *FunctionAST) OldCreate(module *ir.Module) {
	var args []*ir.Param

	for i := 0; i < len(fn.args); i++ {
		args[i] = fn.args[i].get()
	}

	llvmFn := module.NewFunc(fn.Name, fn.returnType, args...)

	// Add the body block.
	block := llvmFn.NewBlock("body")

	// Test add addition statement.
	addSt := block.NewAdd(constant.NewInt(types.I32, 5), constant.NewInt(types.I32, 5))

	block.NewRet(addSt)
}

// Emit : Emit the AST representation.
func (fn *FunctionAST) Emit(module *ir.Module) {
	var args []*ir.Param

	for i := 0; i < len(fn.args); i++ {
		args[i] = fn.args[i].get()
	}

	llvmFn := module.NewFunc(fn.Name, fn.returnType, args...)

	// Emit the function body.
	fn.body.Emit(llvmFn)
}

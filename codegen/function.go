package codegen

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// FunctionAST : Represents the function node.
type FunctionAST struct {
	name       string
	args       []FunctionArgAST
	returnType types.Type
	body       *FunctionBodyAST
}

// TODO: This is just here for reference purposes, can be safely removed (OldCreate).

// OldCreate : Create and apply the function into an LLVM module.
func (node *FunctionAST) OldCreate(module *ir.Module) {
	var args []*ir.Param

	for i := 0; i < len(node.args); i++ {
		args[i] = node.args[i].get()
	}

	fn := module.NewFunc(node.name, node.returnType, args...)

	// Add the body block.
	block := fn.NewBlock("body")

	// Test add addition statement.
	addSt := block.NewAdd(constant.NewInt(types.I32, 5), constant.NewInt(types.I32, 5))

	block.NewRet(addSt)
}

// Emit : Emit the AST representation.
func (node *FunctionAST) Emit(module *ir.Module) {
	var args []*ir.Param

	for i := 0; i < len(node.args); i++ {
		args[i] = node.args[i].get()
	}

	fn := module.NewFunc(node.name, node.returnType, args...)

	// Emit the function body.
	node.body.Emit(fn)
}

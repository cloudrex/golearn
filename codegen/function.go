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

// GetName : Retrieve the function node's name.
func (node *FunctionAST) GetName() string {
	return node.name
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

	// TODO: Just for debugging.
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

	// TODO: Temporarily set function type to void. It should be 'node.returnType' but it's not assigned anywhere.
	fn := module.NewFunc(node.name, types.Void, args...)

	// Emit the function body.
	node.body.Emit(fn)
}

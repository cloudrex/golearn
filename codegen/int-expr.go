package codegen

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// IntExprAST : Represents the integer-32 value node.
type IntExprAST struct {
	value int
}

// GetValue : Retrieve the node's value.
func (node *IntExprAST) GetValue() *constant.Int {
	return constant.NewInt(types.I32, node.value)
}

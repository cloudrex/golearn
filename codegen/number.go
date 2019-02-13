package codegen

import (
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

// NumberAST : Represents the integer-64 value node.
type NumberAST struct {
	value int64
}

func (node NumberAST) get() *constant.Int {
	return constant.NewInt(types.I32, node.value)
}

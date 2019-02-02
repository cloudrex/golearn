package ast

import "github.com/llir/llvm/ir/value"

// Number : Represents the AST node of a number.
type Number struct {
}

func (number Number) generate() Value {
	return "TODO"
}

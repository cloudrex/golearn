package ast

import "github.com/llir/llvm/ir/value"

// Expression : Represents the AST of a generic expression.
type Expression interface {
	generate() Value
}



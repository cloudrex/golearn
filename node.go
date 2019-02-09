package main

import "github.com/llir/llvm/ir/value"

// ValueNode : Represents an AST node which generates an LLVM Value.
type ValueNode interface {
	// TODO: Should return type LLVM.Value.
	get() value.Value
}

// NamedValueNode : Represents an AST node which generates an LLVM Named Value.
type NamedValueNode interface {
	get() value.Named
}

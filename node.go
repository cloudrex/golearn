package main

import "github.com/llir/llvm/ir/value"

// ValueNode : Represents an AST node which generates an LLVM Value.
type ValueNode interface {
	// TODO: Should return type LLVM.Value.
	generate() value.Value
}

// NamedValueNode : Represents an AST node which generates an LLVM Named Value.
type NamedValueNode interface {
	generate() value.Named
}

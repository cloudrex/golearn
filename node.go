package main

type ExprNode interface {
	// TODO: Should return type LLVM.Value.
	generate()
}
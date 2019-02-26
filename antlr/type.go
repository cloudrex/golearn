package main

import "github.com/llir/llvm/ir/types"

// TypeMap : Map of local types to LLVM types.
var TypeMap = map[string]types.Type{
	"int":    types.I32,
	"int64":  types.I64,
	"long":   types.I128,
	"short":  types.I16,
	"float":  types.Float,
	"double": types.Double,
	"str":    types.I8Ptr,
	"char":   types.I8,
	"bool":   types.I1,
}

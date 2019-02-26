package main

import (
	"fmt"

	"github.com/llir/llvm/ir/types"

	"github.com/llir/llvm/ir"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type golearnListener struct {
	mod *ir.Module
	*parser.BaseGolearnListener
	mainExists bool
}

func newGolearnListener() golearnListener {
	return golearnListener{mod: ir.NewModule(), mainExists: false}
}

func (s *golearnListener) EnterAssign(ctx *parser.AssignContext) {
	id := ctx.Id()

	fmt.Println("Id:", id)
}

func (s *golearnListener) EnterFn(ctx *parser.FnContext) {
	name := ctx.Id().GetSymbol().GetText()

	// Register entry point flag.
	if name == "main" {
		if !s.mainExists {
			s.mainExists = true
		} else {
			fmt.Println("Conflicting multiple entry points encountered")
		}
	}

	var returnType types.Type

	if ctx.Type() != nil {
		returnType = ResolveType(ctx.Type().GetSymbol().GetText())
	}

	s.mod.NewFunc(name, returnType)

	fmt.Println("--- LLVM IR ---\n", s.mod)
}

// ResolveType : Resolve the corresponding LLVM type from a string value.
func ResolveType(value string) types.Type {
	switch value {
	case "short":
		return types.I16

	case "int":
		return types.I32

	case "int64":
		return types.I64

	case "long":
		return types.I128

	case "float":
		return types.Float

	case "double":
		return types.Double

	case "void":
		return nil

	case "bool":
		return types.I1

	default:
		panic(fmt.Errorf("Cannot resolve unknown type value: %v", value))
	}
}

func main() {
	// Setup the input
	is := antlr.NewInputStream("space test ; @attrib @attrib() fn main(*str myArg, str mySecond) ~> float { myStr = 5; }")

	// Create the Lexer
	lexer := parser.NewGolearnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGolearnParser(stream)

	// Finally parse the expression
	listener := newGolearnListener()

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Ensure entry point exists.
	if !listener.mainExists {
		fmt.Println("No entry point found")
	}
}

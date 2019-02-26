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
	//
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

	// Create and register the function signature.
	fn := s.mod.NewFunc(name, returnType)

	// Create and apply the function body block.
	body := fn.NewBlock("entry")

	// Apply the required return instruction.
	body.NewRet(nil)
}

// ResolveType : Resolve the corresponding LLVM type from a string value.
func ResolveType(value string) types.Type {
	return TypeMap[value]
}

func main() {
	// Setup the input.
	is := antlr.NewInputStream("space test ; @attrib @attrib() fn main(str myArg, str mySecond) ~> float { myStr = 5; som.eth.ing = 6; }")

	// Create the Lexer.
	lexer := parser.NewGolearnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser.
	p := parser.NewGolearnParser(stream)

	// Finally parse the expression.
	listener := newGolearnListener()

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Ensure entry point exists.
	if !listener.mainExists {
		panic("No entry point found")
	}

	fmt.Println("--- LLVM IR ---\n", listener.mod)
}

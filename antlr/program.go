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
	// TODO
}

func (s *golearnListener) EnterExtern(ctx *parser.ExternContext) {
	var returnType types.Type

	if ctx.Type() != nil {
		returnType = ResolveType(ctx.Type().GetSymbol().GetText())
	}

	// Emit the function declaration without a body.
	s.mod.NewFunc(ctx.Id().GetSymbol().GetText(), returnType)
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

	// Create and apply block statements.
	/*for _, statement := range ctx.Block().(*parser.BlockContext).AllStatement() {
		// TODO: Debugging
		// fmt.Println("Statement:", statement)
		// fmt.Println("Assign:", statement.(*parser.StatementContext).Assign().(*parser.AssignContext))
	} */

	// Apply the required return instruction.
	body.NewRet(nil)
}

// ResolveType : Resolve the corresponding LLVM type from a string value.
func ResolveType(value string) types.Type {
	return TypeMap[value]
}

func main() {
	const inputFile = "input.golearn"

	if !DoesFilePathExist(inputFile) {
		panic(fmt.Errorf("Input file '%v' does not exist", inputFile))
	}

	input := Read(inputFile)

	// Setup the input.
	is := antlr.NewInputStream(input)

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

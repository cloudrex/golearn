package main

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/llir/llvm/ir"

	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type golearnListener struct {
	mod *ir.Module
	*parser.BaseGolearnListener
	entryFound bool
}

func newGolearnListener() golearnListener {
	return golearnListener{mod: ir.NewModule(), entryFound: false}
}

func (s *golearnListener) EnterAssign(ctx *parser.AssignContext) {
	// TODO
}

func (s *golearnListener) EnterExtern(ctx *parser.ExternContext) {
	// fnSig := ctx.FnSig()

	/*if hasMethod(fnSig, "AnyType") {
		fmt.Println("has anytype!")
	}*/
	// fnSigType := fnSig.AnyType().(*parser.FnSigContext)

	/*if fnSigType != nil {
		fmt.Println("FN Signature not nil!")
	}*/

	// Emit the function declaration without a body.
	//s.mod.NewFunc(ctx.Id().GetSymbol().GetText(), returnType)
}

func (s *golearnListener) EnterFnEntry(ctx *parser.FnEntryContext) {
	if s.entryFound {
		panic("Entry point cannot be declared multiple times")
	}

	s.entryFound = true
}

func (s *golearnListener) EnterFn(ctx *parser.FnContext) {
	name := ctx.Id().GetSymbol().GetText()

	// Emit the function return type if applicable.
	//var returnType types.Type

	// Create and register the function signature.
	fn := s.mod.NewFunc(name, nil /* returnType */)

	// Create and apply the function body block.
	body := fn.NewBlock("begin")

	// Create and apply block statements.
	/*for _, statement := range ctx.Block().(*parser.BlockContext).AllStatement() {
		// TODO: Debugging
		// fmt.Println("Statement:", statement)
		// fmt.Println("Assign:", statement.(*parser.StatementContext).Assign().(*parser.AssignContext))
	} */

	// Apply the required return instruction.
	body.NewRet(nil)
}

func resolveTypeBasic(mayTypeBasic parser.ITypeBasicContext) types.Type {
	// Basic type
	if typeBasic, ok := mayTypeBasic.(*parser.TypeBasicContext); ok {
		typeSimple := typeBasic.TypeSimple()

		// Simple type.
		if typeSimple != nil {
			return ResolveType(typeSimple.GetSymbol().GetText())
		}

		// Otherwise, complex type.
		panic("Complex type support not yet implemented")
	}

	// Otherwise, void type.
	return nil
}

// IsDeclareStatement : Determine if input statement is a variable declaration statement.
func IsDeclareStatement(statement parser.StatementContext) bool {
	if _, ok := statement.Declare().(*parser.DeclareContext); ok {
		return true
	}

	return false
}

// GetBlockStatements : Convert and retrieve all statements within a statement block.
func GetBlockStatements(block parser.IBlockContext) []parser.StatementContext {
	var statements []parser.StatementContext

	for _, statement := range block.(*parser.BlockContext).AllStatement() {
		statements = append(statements, *(statement.(*parser.StatementContext)))
	}

	return statements
}

func handleStatement(ref *ir.Block, statement parser.StatementContext) {
	// TODO: Debugging
	fmt.Println("Statement:", statement)

	if IsDeclareStatement(statement) { // Variable declaration.
		declare := statement.Declare().(*parser.DeclareContext)
		varType := resolveTypeBasic(declare.TypeBasic())
		varName := declare.Id().GetSymbol().GetText()

		// Variable type cannot be void.
		if varType == nil {
			panic("Variable type cannot be void")
		}

		// Allocate type space (declare).
		varRef := ref.NewAlloca(varType)

		// Set the variable name in IR.
		varRef.SetName(varName)

		// Assign value if applicable (assign).
		// TODO: Parse expr and get value.
	}

	// fmt.Println("Assign:", statement.(*parser.StatementContext).Assign().(*parser.AssignContext))
}

func handleBlock(ref *ir.Block, block parser.IBlockContext) {
	var retValue value.Value

	// Create and apply block statements.
	for _, statement := range GetBlockStatements(block) {
		// Invoke the block statement handler.
		handleStatement(ref, statement)
	}

	// Apply the required return instruction.
	ref.NewRet(retValue)
}

func (s *golearnListener) EnterFnTopLevel(ctx *parser.FnTopLevelContext) {
	name := ctx.Id().GetSymbol().GetText()

	// Emit the function return type if applicable.
	var retLLvmType types.Type

	mayRetType := ctx.FnReturnType()

	// Return type specified.
	if retType, ok := mayRetType.(*parser.FnReturnTypeContext); ok {
		retLLvmType = resolveTypeBasic(retType.TypeBasic())
	}

	// Create and register the function signature.
	fn := s.mod.NewFunc(name, retLLvmType)

	// Create and apply the function body block.
	body := fn.NewBlock("begin")

	// Invoke the statement block handler.
	handleBlock(body, ctx.Block())
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

	// Read the input source file.
	input := Read(inputFile)

	// Setup the input.
	is := antlr.NewInputStream(input)

	// Create the Lexer.
	lexer := parser.NewGolearnLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Read all tokens.
	fmt.Println("--- TOKENS ---\n")

	/*for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}*/

	// Create the Parser.
	p := parser.NewGolearnParser(stream)

	// Finally parse the expression.
	listener := newGolearnListener()

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	// Ensure entry point exists.
	if !listener.entryFound {
		panic("No entry point found")
	}

	fmt.Println("\n--- LLVM IR ---\n", listener.mod)
}

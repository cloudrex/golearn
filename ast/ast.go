package ast

import (
	"fmt"
	"golearn/parser"
	"golearn/scanner"

	"github.com/llir/llvm/ir"
)

// Ast : Represents the AST generator.
type Ast struct {
	Parser        *parser.Parser
	stashedParser *parser.Parser
}

// BlockNode : Base interface for an AST block node.
type BlockNode interface {
	create(block *ir.Block)
}

// ModuleNode : Base interface for an AST module node.
type ModuleNode interface {
	create(module *ir.Module)
}

// IdentifierAST : Represents the identifier AST node.
type IdentifierAST struct {
	name string
}

// ExpressionAST : Represents an expression AST node.
type ExpressionAST struct {
	tokens []scanner.Token
}

// BlockAST : Represents a statement block AST node.
type BlockAST struct {
	label      string
	statements []BlockNode
}

func (block *BlockAST) create() *ir.Block {
	irBlock := ir.NewBlock(block.label)

	for i := 0; i < len(block.statements); i++ {
		statement := block.statements[i]

		statement.create(irBlock)
	}

	return irBlock
}

// Function : Process and validate function.
func (gen *Ast) Function() {
	// Must be followed by an identifier.
	if gen.Parser.Peek().Kind != scanner.TokenKindIdentifier {
		gen.Parser.Fatal("Expecting identifier after function definition keyword")

		return
	}

	// Consume identifier.
	token := gen.Parser.Consume().Next()

	if token.Kind != scanner.TokenKindParenStart {
		gen.Parser.Fatal("Expecting argument list after function identifier: '('")

		return
	}

	// Invoke function argument parser.
	gen.functionArgs()

	// Verify block start is present after argument list.
	token = gen.Parser.Next()

	if token.Kind != scanner.TokenKindBlockStart {
		fmt.Println("Type is", token)
		gen.Parser.Fatal("Expecting statement block after function argument list: '{'")
	}

	// Invoke block parser.
	gen.block()
}

// Process and validate function arguments.
func (gen *Ast) functionArgs() {
	derived := gen.Parser.Derive()

	// Sync derived parser's position with original parser.
	derived.Sync(gen.Parser)

	for token := derived.Get(); token.Kind != scanner.TokenKindParenEnd; token = derived.Next() {
		// TODO: Debugging.
		fmt.Println("Parsing args ... Pos", derived.Pos)

		// TODO: Need to process args.
		if derived.Peek().Kind == scanner.TokenKindEndOfFile {
			derived.Fatal("Expecting end of function argument list: ')'")
		}
	}

	// Consume argument list end ')'.
	derived.Consume()
}

// Save the current parser and apply a new parser.
func (gen *Ast) applyParser(parser *parser.Parser) *Ast {
	gen.stashedParser = gen.Parser
	gen.Parser = parser

	return gen
}

// Revert the current parser to a previously stashed parser.
func (gen *Ast) revertParser() *Ast {
	gen.Parser = gen.stashedParser

	return gen
}

// Process and validate a statement block.
func (gen *Ast) block() BlockAST {
	token := gen.Parser.Next()

	if token.Kind != scanner.TokenKindBlockStart {
		gen.Parser.Fatal("Expecting block start: '{'")
	}

	tokens := gen.Parser.Until(scanner.TokenKindBlockEnd)

	// Create and apply a termporal new parser for local use.
	parser := parser.NewParser(tokens)

	gen.applyParser(parser)

	var statements []BlockNode

	// TODO.
	for i := 0; i < len(tokens); i++ {
		statements = append(statements, gen.statement())
	}

	gen.revertParser()

	return BlockAST{label: "anonymous_block", statements: statements}
}

func (gen *Ast) statement() BlockNode {
	tokens := gen.Parser.Until(scanner.TokenKindSemiColon)

	if len(tokens) == 1 || len(tokens) == 1 { // Empty statement.
		// TODO.
	}

	var node BlockNode

	// Validate captured tokens.
	for i := 0; i < len(tokens); i++ {
		// TODO.
	}

	return node
}

func (gen *Ast) expression() ExpressionAST {
	// TODO.
	panic("Not yet implemented")
}

// Process and validate an identifier.
func (gen *Ast) identifier() IdentifierAST {
	token := gen.Parser.Get()

	if !scanner.IsIdentifier(token.Value) {
		gen.Parser.Fatal("Expecting an identifier")
	}

	return IdentifierAST{name: token.Value}
}

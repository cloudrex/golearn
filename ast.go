package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

// Ast : Represents the AST generator.
type Ast struct {
	parser        *Parser
	stashedParser *Parser
}

// AstBlockNode : Base interface for an AST block node.
type AstBlockNode interface {
	create(block *ir.Block)
}

// AstModuleNode : Base interface for an AST module node.
type AstModuleNode interface {
	create(module *ir.Module)
}

// IdentifierAST : Represents the identifier AST node.
type IdentifierAST struct {
	name string
}

// OperandKind : Represents the type of operand used in an operation.
type OperandKind = int

const (
	// OperandKindIdentifier : Represents an identifier operand.
	OperandKindIdentifier OperandKind = 0

	// OperandKindLiteral : Represents a literal value operand.
	OperandKindLiteral OperandKind = 1
)

// OperandAST : Represents an operand AST node.
type OperandAST struct {
	value string
	kind  OperandKind
}

// OperatorKind : Represents the type of an operator.
type OperatorKind = int

const (
	// OperatorKindAdd : Represents the addition operator type.
	OperatorKindAdd OperatorKind = 0

	// OperatorKindSubstract : Represents the substraction operator type.
	OperatorKindSubstract OperatorKind = 1

	// OperatorKindMultiply : Represents the multiplication operator type.
	OperatorKindMultiply OperatorKind = 1

	// OperatorKindDivide : Represents the division operator type.
	OperatorKindDivide OperatorKind = 1
)

// OperatorAST : Represents an operator AST node.
type OperatorAST struct {
	value string
	kind  OperatorKind
}

// BinaryOperationAST : Represents a binary operation AST node.
type BinaryOperationAST struct {
	leftSide  OperandAST
	rightSide OperandAST
	operator  OperatorAST
}

// ExpressionAST : Represents an expression AST node.
type ExpressionAST struct {
	tokens []Token
}

// BlockAST : Represents a statement block AST node.
type BlockAST struct {
	label      string
	statements []AstBlockNode
}

func (block *BlockAST) create() *ir.Block {
	irBlock := ir.NewBlock(block.label)

	for i := 0; i < len(block.statements); i++ {
		statement := block.statements[i]

		statement.create(irBlock)
	}

	return irBlock
}

// Process and validate function.
func (ast *Ast) function() {
	// Must be followed by an identifier.
	if ast.parser.peek().kind != TokenKindIdentifier {
		ast.parser.fatal("Expecting identifier after function definition keyword")

		return
	}

	// Consume identifier.
	token := ast.parser.consume().next()

	if token.kind != TokenKindParenStart {
		ast.parser.fatal("Expecting argument list after function identifier: '('")

		return
	}

	// Invoke function argument parser.
	ast.functionArgs()

	// Verify block start is present after argument list.
	token = ast.parser.next()

	if token.kind != TokenKindBlockStart {
		fmt.Println("Type is", token)
		ast.parser.fatal("Expecting statement block after function argument list: '{'")
	}

	// Invoke block parser.
	ast.block()
}

// Process and validate function arguments.
func (ast *Ast) functionArgs() {
	derived := ast.parser.derive()

	// Sync derived parser's position with original parser.
	derived.sync(ast.parser)

	for token := derived.get(); token.kind != TokenKindParenEnd; token = derived.next() {
		// TODO: Debugging.
		fmt.Println("Parsing args ... Pos", derived.pos)

		// TODO: Need to process args.
		if derived.peek().kind == TokenKindEndOfFile {
			derived.fatal("Expecting end of function argument list: ')'")
		}
	}

	// Consume argument list end ')'.
	derived.consume()
}

// Save the current parser and apply a new parser.
func (ast *Ast) applyParser(parser *Parser) *Ast {
	ast.stashedParser = ast.parser
	ast.parser = parser

	return ast
}

// Revert the current parser to a previously stashed parser.
func (ast *Ast) revertParser() *Ast {
	ast.parser = ast.stashedParser

	return ast
}

// Process and validate a statement block.
func (ast *Ast) block() BlockAST {
	token := ast.parser.next()

	if token.kind != TokenKindBlockStart {
		ast.parser.fatal("Expecting block start: '{'")
	}

	tokens := ast.parser.until(TokenKindBlockEnd)

	// Create and apply a termporal new parser for local use.
	parser := newParser(tokens)

	ast.applyParser(parser)

	var statements []AstBlockNode

	// TODO.
	for i := 0; i < len(tokens); i++ {
		statements = append(statements, ast.statement())
	}

	ast.revertParser()

	return BlockAST{label: "anonymous_block", statements: statements}
}

func (ast *Ast) statement() AstBlockNode {
	tokens := ast.parser.until(TokenKindSemiColon)

	if len(tokens) == 1 || len(tokens) == 1 { // Empty statement.
		// TODO.
	}

	var node AstBlockNode

	// Validate captured tokens.
	for i := 0; i < len(tokens); i++ {
		// TODO.
	}

	return node
}

func (ast *Ast) expression() ExpressionAST {
	// TODO.
	panic("Not yet implemented")
}

// Process and validate an identifier.
func (ast *Ast) identifier() IdentifierAST {
	token := ast.parser.get()

	if !isIdentifier(token.value) {
		ast.parser.fatal("Expecting an identifier")
	}

	return IdentifierAST{name: token.value}
}

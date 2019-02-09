package main

import "fmt"

// AST : Represents the AST generator.
type AST struct {
	parser *Parser
}

// Process and validate function.
func (ast *AST) function() {
	// Must be followed by an identifier.
	if ast.parser.peek().kind != TokenKindIdentifier {
		ast.parser.fatal("Expecting identifier after function definition keyword")

		return
	}

	// Consume identifier.
	token := ast.parser.cycle().next()

	if token.kind != TokenKindParenStart {
		ast.parser.fatal("Expecting argument list after function identifier: '('")

		return
	}

	// Invoke function argument parser.
	ast.functionArgs()
}

// Process and validate function arguments.
func (ast *AST) functionArgs() {
	derived := ast.parser.derive()

	for token := derived.get(); token.kind != TokenKindParenEnd; token = derived.next() {
		fmt.Println("Parsing args ... Pos", derived.pos)

		if derived.peek().kind == TokenKindEndOfFile {
			derived.fatal("Expecting end of function argument list: ')'")
		}
	}
}

// Process and validate a statement block.
func (ast *AST) block() {
	token := ast.parser.next()

	if token.kind != TokenKindBlockStart {
		ast.parser.fatal("Expecting block start: '{'")
	}
}

func (ast *AST) statement() {
	// TODO
}

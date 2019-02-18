package parser

import (
	"golearn/codegen"
	"golearn/lex"
)

// ExprParser : Utility class for parsing various expressions.
type ExprParser struct {
	sequence []lex.Token
	parser   *Parser
}

func (expr *ExprParser) parseIdentifierExpr() *codegen.ExprAST {
	token := expr.parser.Get()
	name := token.Value

	// Consume identifier.
	token = expr.parser.Next()

	if token.Value != "(" { // Variable reference.
		return codegen.VariableAST{name: name}
	}

	// Function call. Consume '('.
	token = expr.parser.Next()

	// TODO: Continue implementing (Kaleidoscope #3).

	panic("Not yet implemented")
}

// parsePrimary : Parse a primary expression.
func (expr *ExprParser) parsePrimaryExpr() *codegen.ExprAST {
	switch expr.parser.Get().Kind {
	case lex.TokenKindIdentifier:
		{
			return expr.parseIdentifierExpr()
		}

		// TODO: Continue implementing (Kaleidoscope #3).
	}

	panic("Not yet implemented")
}

// parseNumber : Parse an integer expression.
func (expr *ExprParser) parseIntExpr() {
	// TODO
}

// Parse : Parse an expression.
func (expr *ExprParser) Parse() {
	leftSide := expr.parsePrimaryExpr()
}

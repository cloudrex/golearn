package codegen

import (
	"fmt"
	"golearn/parser"
	"golearn/scanner"

	"github.com/llir/llvm/ir/constant"

	"github.com/llir/llvm/ir/value"

	"github.com/llir/llvm/ir/types"

	"github.com/llir/llvm/ir"
)

var variableTypes = []string{"string", "float", "int"}

// CodeGenerator : Represents the code generator.
type CodeGenerator struct {
	// TODO: Make 'Parser' private ('parser').
	Parser *parser.Parser

	stashedParser *parser.Parser
	module        *ir.Module
}

// GetParser : Retrieve the code generator's parser.
func (gen *CodeGenerator) GetParser() *parser.Parser {
	return gen.Parser
}

// NewCodeGenerator : Create a new instance of the code generator module.
func NewCodeGenerator(parser *parser.Parser, module *ir.Module) *CodeGenerator {
	return &CodeGenerator{Parser: parser, module: module}
}

// BlockNode : Base interface for an AST block node.
type BlockNode interface {
	Emit(block *ir.Block)
}

// ModuleNode : Base interface for an AST module node.
type ModuleNode interface {
	Emit(module *ir.Module)
}

// FuncNode : Base interface for an AST function node.
type FuncNode interface {
	Emit(fn *ir.Func)
}

// IdentifierAST : Represents the identifier AST node.
type IdentifierAST struct {
	name string
}

// ExpressionAST : Represents an expression AST node.
type ExpressionAST struct {
	tokens []scanner.Token
}

// Function : Process and validate function.
func (gen *CodeGenerator) Function() FunctionAST {
	var fn FunctionAST

	// Must be followed by an identifier.
	if gen.Parser.Peek().Kind != scanner.TokenKindIdentifier {
		gen.Parser.Fatal("Expecting identifier after function definition keyword")

		return fn
	}

	// Function identifier.
	token := gen.Parser.Next()

	// Verify the function doesn't already exist.
	for _, v := range gen.module.Funcs {
		if v.Name() == token.Value {
			gen.Parser.Fatal("Cannot re-define function. A function named '" + token.Value + "' is already defined")
		}
	}

	// Assign function name.
	fn.name = token.Value

	// Continue to next token.
	token = gen.Parser.Next()

	if token.Kind != scanner.TokenKindParenStart {
		gen.Parser.Fatal("Expecting argument list after function identifier: '('")

		return fn
	}

	// Invoke function argument parser and apply.
	fn.args = gen.functionArgs()

	// Verify block start is present after argument list.
	token = gen.Parser.Peek()

	// Override block start error for more specific feedback (regarding a function).
	if token.Kind != scanner.TokenKindBlockStart {
		gen.Parser.Fatal("Expecting statement block after function argument list: '{'")
	}

	// Invoke block parser.
	fn.body = gen.functionBody()

	return fn
}

// Process and validate function arguments.
func (gen *CodeGenerator) functionArgs() []FunctionArgAST {
	derived := gen.Parser.Derive()

	// Sync derived parser's position with original parser.
	derived.Link(gen.Parser)

	for token := derived.Get(); token.Kind != scanner.TokenKindParenEnd; token = derived.Next() {
		// TODO: Debugging.
		fmt.Println("Parsing args ... Pos", derived.GetPos())

		// TODO: Need to process args.
		if derived.Peek().Kind == scanner.TokenKindEndOfFile {
			derived.Fatal("Expecting end of function argument list: ')'")
		}
	}

	// TODO: Does this modify anything/cause any action?

	// TODO: Returning empty for future implementation.
	return []FunctionArgAST{}
}

// Save the current parser and apply a new parser.
func (gen *CodeGenerator) applyParser(parser *parser.Parser) *CodeGenerator {
	gen.stashedParser = gen.Parser
	gen.Parser = parser

	return gen
}

// Revert the current parser to a previously stashed parser.
func (gen *CodeGenerator) revertParser() *CodeGenerator {
	gen.Parser = gen.stashedParser

	return gen
}

// Process and validate a function statement blocks.
func (gen *CodeGenerator) functionBody() *FunctionBodyAST {
	token := gen.Parser.Next()

	if token.Kind != scanner.TokenKindBlockStart {
		gen.Parser.Fatal("Expecting block start: '{'")
	} else if gen.Parser.Peek().Kind == scanner.TokenKindBlockEnd { // Empty block.
		return &FunctionBodyAST{}
	}

	// Consume block start '{'.
	gen.Parser.Consume()

	var statements []BlockNode

	// Statement parse call will consume all statement tokens until the semi-colon; We can safely loop/continue without acquiring next token(s).
	for token = gen.Parser.Get(); token.Kind != scanner.TokenKindBlockEnd; token = gen.Parser.Get() {
		// Ensure parser doesn't reach end of file when expecting end of block token.
		if token.Kind == scanner.TokenKindEndOfFile {
			gen.Parser.Fatal("Expecting block end: '}'")
		}

		statements = append(statements, *gen.statement())
	}

	return &FunctionBodyAST{statements: statements}
}

// TODO: Work on statement().
func (gen *CodeGenerator) statement() *BlockNode {
	tokens := gen.Parser.Until(scanner.TokenKindSemiColon)

	// TODO: Debugging.
	fmt.Println("Statement parser tokens:", tokens)

	if len(tokens) == 1 || len(tokens) == 1 { // Empty statement.
		// TODO
	}

	var node BlockNode

	// Validate captured tokens.
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		// TODO: Each section can be split, ex. parseVariableDeclaration(), parseAssignment(), etc.
		if IsTypeKeyword(token.Value) && i == 0 { // Variable declaration.
			if IsVariableDeclaration(tokens) {
				// TODO: CRITICAL: tokens[i+1] and tokens[i+3] have not been verified as existing and correct by IsVariableDeclaration(). Cannot be verified either because value can be an expression, and expressions can consist of tokens 1+ (IsVariableDeclaration cannot handle).
				name := tokens[i+1].Value
				varNode := &VariableAST{name: name}

				// Declaration sequence consists of 3 tokens.
				if len(tokens) > i+3 { // Declaration is not empty.
					varNode.Value = gen.resolveValue(tokens[i+3])
				}

				node = varNode

				break
			} else {
				gen.Parser.Fatal("Expecting valid variable declaration sequence")
			}
		} else if token.Kind == scanner.TokenKindIdentifier && IsVariableAssignment(tokens) { // Variable assignment.
			// Skip assignment sequence for value (2 tokens).
			// TODO: CRITICAL: variableType is hard-coded to float. Way to determine var type needed.
			node = &VarAssignmentAST{variableName: token.Value, value: gen.resolveValue(tokens[i+2]), variableType: scanner.VariableTypeFloat}

			break
		} else if token.Kind == scanner.TokenKindSemiColon { // Ignore semi-colons.
			continue
		} else {
			gen.Parser.Fatal("Expecting statement containing valid expression(s)")
		}
	}

	return &node
}

func (gen *CodeGenerator) resolveValue(token scanner.Token) value.Value {
	var val value.Value

	// TODO: Implement support for resolving identifiers, function calls and expressions.
	if token.Kind == scanner.TokenKindStringLiteral { // String literal.
		val = constant.NewCharArrayFromString(token.Value)
	} else if token.Kind == scanner.TokenKindIntegerLiteral { // Integer literal.
		result, err := constant.NewIntFromString(types.I32, token.Value)

		if err != nil {
			gen.Parser.Fatal("Failed to convert integer32 literal value from string")
		}

		val = result
	} else if token.Kind == scanner.TokenKindFloatLiteral { // Float literal.
		result, err := constant.NewFloatFromString(types.Float, token.Value)

		if err != nil {
			gen.Parser.Fatal("Failed to convert float literal value to float from string")
		}

		val = result
	} else {
		gen.Parser.Fatal("Expecting a numeric literal value")
	}

	return val
}

func (gen *CodeGenerator) expression() ExpressionAST {
	// TODO
	panic("Not yet implemented")
}

// Process and validate an identifier.
func (gen *CodeGenerator) identifier() IdentifierAST {
	token := gen.Parser.Get()

	if !scanner.IsIdentifier(token.Value) {
		gen.Parser.Fatal("Expecting an identifier")
	}

	return IdentifierAST{name: token.Value}
}

// IsVariableAssignment : Determine if the provided sequence of tokens represent a variable assignment.
func IsVariableAssignment(sequence []scanner.Token) bool {
	parser := parser.NewParser(sequence)

	return parser.Peek().Kind == scanner.TokenKindEqualSign
}

// IsVariableDeclaration : Determine if the provided sequence of tokens represent a variable declaration.
func IsVariableDeclaration(sequence []scanner.Token) bool {
	parser := parser.NewParser(sequence)

	if IsTypeKeyword(parser.Get().Value) && parser.Next().Kind == scanner.TokenKindIdentifier {
		peekKind := parser.Peek().Kind

		if peekKind == scanner.TokenKindSemiColon { // Empty declaration.
			return true
		} else if peekKind == scanner.TokenKindEqualSign { // Declaration + assignment.
			// TODO: Not checking after '='.

			return true
		}
	}

	return false
}

// IsTypeKeyword : Determine if input is a variable type keyword.
func IsTypeKeyword(input string) bool {
	for i := 0; i < len(variableTypes); i++ {
		if variableTypes[i] == input {
			return true
		}
	}

	return false
}

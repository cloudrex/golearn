package main

import (
	"fmt"
	"golearn/parser"
	"golearn/scanner"

	"golearn/codegen"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `fn hello () { hello := ; }`

	var lexer = scanner.Scanner{}
	var tokens = lexer.Scan(src)
	var parser = parser.NewParser(tokens)
	var generator = codegen.CodeGenerator{Parser: parser}

	for token := parser.Get(); parser.Get().Kind != scanner.TokenKindEndOfFile; token = parser.Next() {
		if token.Kind == scanner.TokenKindFn { // Function declaration 'fn'.
			// Invoke the function AST generator.
			generator.Function()
		} else if token.Kind == scanner.TokenKindUnknown { // Unknown token.
			parser.Fatal("Unknown token")

			return
		}

		fmt.Println("[ Token:", parser.Pos, "] ->", token.Value, "(", token.Kind, ")")
	}

	var module = ir.NewModule()

	// TODO: Creating function for debugging/testing.
	fn := codegen.FunctionAST{Name: "very fun func"}

	fn.Create(module)

	// fmt.Println(strings.Join(values, " "))
	// Print the LLVM IR assembly of the module.
	fmt.Println(module)
}

package main

import (
	"fmt"
	"golearn/codegen"
	"golearn/parser"
	"golearn/scanner"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `fn hello () { hello : = ; hello = 3.14 ; } fn hello () { }`

	var lexer = scanner.Scanner{}
	var tokens = lexer.Scan(src)
	var parser = parser.NewParser(tokens)

	// Global LLVM module.
	module := ir.NewModule()

	// Create the code generator and attach parser + module.
	generator := codegen.NewCodeGenerator(parser, module)

	for token := parser.Get(); parser.Get().Kind != scanner.TokenKindEndOfFile; token = parser.Next() {
		if token.Kind == scanner.TokenKindFn { // Function declaration 'fn'.
			// Invoke the function AST generator.
			fn := generator.Function()

			// Emit the function.
			fn.Emit(module)
		} else if token.Kind == scanner.TokenKindUnknown { // Unknown token.
			parser.Fatal("Unknown token")

			return
		}

		fmt.Printf("[Token: %v] -> %v (%v)", parser.GetPos(), token.Value, token.Kind)
	}

	fmt.Println("\n\n--- LLVM IR ---\n")
	fmt.Println(module)
}

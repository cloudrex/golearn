package main

import (
	"fmt"
	"golearn/codegen"
	"golearn/parser"
	"golearn/scanner"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `fn main () { hello : = ; hello = 3.14 ; } fn _hello () { }`
	const mainFnName = "main"

	var lexer = scanner.Scanner{}
	var tokens = lexer.Scan(src)
	var parser = parser.NewParser(tokens)

	// Global LLVM module.
	module := ir.NewModule()

	// Create the code generator and attach parser + module.
	generator := codegen.NewCodeGenerator(parser, module)

	// Flag representing whether the main function was found.
	mainFound := false

	for token := parser.Get(); parser.Get().Kind != scanner.TokenKindEndOfFile; token = parser.Next() {
		if token.Kind == scanner.TokenKindFn { // Function declaration 'fn'.
			// Invoke the function AST generator.
			fn := generator.Function()

			// Emit the function.
			fn.Emit(module)

			// Ensure the main function is declared. No need for re-declaration check, as already handled elsewhere.
			if fn.GetName() == mainFnName {
				mainFound = true
			}
		} else if token.Kind == scanner.TokenKindUnknown { // Unknown token.
			parser.UnknownToken()

			return
		}

		fmt.Printf("[Token: %v] -> %v (%v)", parser.GetPos(), token.Value, token.Kind)
	}

	// Ensure a main function was declared.
	if !mainFound {
		parser.NoMainFnFound()
	}

	fmt.Println("\n\n--- LLVM IR ---\n")
	fmt.Println(module)
}

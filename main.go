package main

import (
	"fmt"
	"golearn/parser"
	"golearn/scanner"

	"github.com/llir/llvm/ir"

	"golearn/codegen"
)

func main() {
	const src = `fn hello () { hello : = ; }`

	var lexer = scanner.Scanner{}
	var tokens = lexer.Scan(src)
	var parser = parser.NewParser(tokens)
	var generator = codegen.CodeGenerator{Parser: parser}

	// Global LLVM module.
	module := ir.NewModule()

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

		fmt.Println("[ Token:", parser.GetPos(), "] ->", token.Value, "(", token.Kind, ")")
	}

	fmt.Println("\n--- LLVM IR ---")
	fmt.Println(module)
}

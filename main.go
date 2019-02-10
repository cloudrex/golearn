package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `fn hello ()`

	var scanner = Scanner{}
	var tokens = scanner.scan(src)
	var parser = newParser(tokens)
	var ast = Ast{parser: parser}

	for token := parser.get(); parser.get().kind != TokenKindEndOfFile; token = parser.next() {
		if token.kind == TokenKindFn { // Function declaration 'fn'.
			// Invoke the function AST generator.
			ast.function()
		} else if token.kind == TokenKindUnknown { // Unknown token.
			parser.fatal("Unknown token")

			return
		}

		fmt.Println("[ Token:", parser.pos, "] ->", token.value, "(", token.kind, ")")
	}

	var module = ir.NewModule()

	fn := FunctionAST{name: "very fun func"}

	fn.create(module)

	// fmt.Println(strings.Join(values, " "))
	// Print the LLVM IR assembly of the module.
	fmt.Println(module)
}

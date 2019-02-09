package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `fn hello (`

	var tokens = scan(src)
	var parser = newParser(tokens)

	for token := parser.get(); parser.get().kind != TokenKindEndOfFile; token = parser.next() {
		if token.kind == TokenKindFn { // Function declaration 'fn'.
			// Must be followed by an identifier.
			if parser.peek().kind != TokenKindIdentifier {
				fmt.Println(parser.err("Expecting identifier after function definition keyword"))

				return
			}

			// Consume identifier.
			token = parser.cycle().next()

			if token.value != "(" {
				fmt.Println(parser.err("Expecting argument list after function identifier"))

				return
			}
		} else if token.kind == TokenKindUnknown { // Unknown token.
			fmt.Println("Unknown token at position", parser.pos)

			return
		}

		fmt.Println("[ Token:", parser.pos, "] ->", token.value, "(", token.kind, ")")
	}

	var module = ir.NewModule()

	fn := FunctionAST{module: module, name: "very fun func"}

	fn.create()

	// fmt.Println(strings.Join(values, " "))
	// Print the LLVM IR assembly of the module.
	fmt.Println(module)
}

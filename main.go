package main

import (
	"golearn/lex"
)

func main() {
	scn := lex.NewScanner("hello world !")

	scn.Scan()
}

/* func tmpOldMain() {
	const src = `fn main () { float pi = "hello world" ; pi = 2.1 ; float pi2 ; pi2 = 3.16 ; float pi3 = 3.17 ; } fn _4hello () { }`
	const mainFnName = "main"

	tokens := lex.Tokenize(src)
	parser := parser.NewParser(tokens)

	// Global LLVM module.
	module := ir.NewModule()

	// Create the code generator and attach parser + module.
	generator := codegen.NewCodeGenerator(parser, module)

	// Flag representing whether the main function was found.
	mainFound := false

	for token := parser.Get(); parser.Get().Kind != lex.TokenKindEndOfFile; token = parser.Next() {
		if token.Kind == lex.TokenKindFnKeyword { // Function declaration 'fn'.
			// Invoke the function AST generator.
			fn := generator.Function()

			// Emit the function.
			fn.Emit(module)

			// Ensure the main function is declared. No need for re-declaration check, as already handled elsewhere.
			if fn.GetName() == mainFnName {
				mainFound = true
			}
		} else if token.Kind == lex.TokenKindImport { // Import statement.

		} else if token.Kind == lex.TokenKindUnknown { // Unknown token.
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
*/

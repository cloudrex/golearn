package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = ("single-token text with space")
}`

	var module = ir.NewModule()
	var tokens = scan(src)
	var values []string

	f := FunctionAST{module: module, name: "my fun"}

	f.create()

	for i := 0; i < len(tokens); i++ {
		values = append(values, tokens[i].value)
	}

	// fmt.Println(strings.Join(values, " "))
	// Print the LLVM IR assembly of the module.
	fmt.Println(module)
}

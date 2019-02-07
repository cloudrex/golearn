package main

import (
	"fmt"
	"strings"
)

func main() {
	const src = `
// This is scanned code.
if a > 10 {
	someParsable = ("single-token text with space")
}`
	var tokens = scan(src)
	var values []string

	for i := 0; i < len(tokens); i++ {
		values = append(values, tokens[i].value)
	}

	fmt.Println(strings.Join(values, " "))
}

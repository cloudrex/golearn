package parser

import "fmt"

// Err : Creates an error with parser metadata.
func (parser *Parser) Err(message string) error {
	return fmt.Errorf("[At token position %v | kind %v | value '%v'] %v", parser.pos, parser.Get().Kind, parser.Get().Value, message)
}

// Fatal : Creates and displays a fatal error with parser metadata. Stops the application.
func (parser *Parser) Fatal(message string) {
	panic(parser.Err(message))
}

// UnknownToken : Report that a scanned token is unknown. Invokes the fatal function.
func (parser *Parser) UnknownToken() {
	parser.Fatal("Unknown token")
}

// NoMainFnFound : Report that the required 'main' function was not found.
func (parser *Parser) NoMainFnFound() {
	parser.Fatal("Expecting an entry-point function declaration named 'main'")
}

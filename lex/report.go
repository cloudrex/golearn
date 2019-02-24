package lex

import "fmt"

// ShouldAccountTracePrefix : Determine if the trace prefix's character amount should be accounted for.
func (scanner *Scanner) ShouldAccountTracePrefix(amount int) bool {
	if scanner.pos == 0 || scanner.pos < amount {
		return false
	}

	return true
}

// TracePointerX : Create the character position pointer to indicate problem at current position.
func (scanner *Scanner) TracePointerX(accountPrefix bool) string {
	pointer := ""

	if accountPrefix {
		pointer += "    " // Account for '... '.
	}

	for i := 0; i < scanner.pos; i++ {
		pointer += " "
	}

	pointer += "^" // Add the pointer character.

	return pointer
}

// TracePointer : Create the character position pointer to indicate problem at current position. Uses 20 as amount.
func (scanner *Scanner) TracePointer() string {
	return scanner.TracePointerX(scanner.ShouldAccountTracePrefix(20))
}

// TraceSequenceX : Provide feedback on where exactly a certain error occurred.
func (scanner *Scanner) TraceSequenceX(amount int) string {
	if !scanner.ShouldAccountTracePrefix(amount) {
		if len(scanner.input) < amount {
			return scanner.input[0:len(scanner.input)]
		}

		return scanner.input[0:amount]
	}

	return "... " + scanner.input[scanner.pos:scanner.pos+amount]
}

// TradeSequence : Provide feedback on where exactly a certain error occurred. Returns last 20 processed characters.
func (scanner *Scanner) TradeSequence() string {
	return scanner.TraceSequenceX(20)
}

// CreateTrace : Creates a trace string with a character pointer.
func (scanner *Scanner) CreateTrace() string {
	return "\t" + scanner.TradeSequence() + "\n\t" + scanner.TracePointer()
}

// Fatal : Report a fatal message. Application will exit.
func (scanner *Scanner) Fatal(message string) {
	panic(fmt.Errorf(message + "\n\n" + scanner.CreateTrace()))
}

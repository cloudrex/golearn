package lex

import "fmt"

// ShouldAccountTracePrefix : Determine if the trace prefix's character amount should be accounted for.
func (scn *Scanner) ShouldAccountTracePrefix(amount int) bool {
	if scn.pos == 0 || scn.pos < amount {
		return false
	}

	return true
}

// TracePointerX : Create the character position pointer to indicate problem at current position.
func (scn *Scanner) TracePointerX(accountPrefix bool) string {
	pointer := ""

	if accountPrefix {
		pointer += "    " // Account for '... '.
	}

	for i := 0; i < scn.pos; i++ {
		pointer += " "
	}

	pointer += "^" // Add the pointer character.

	return pointer
}

// TracePointer : Create the character position pointer to indicate problem at current position. Uses 20 as amount.
func (scn *Scanner) TracePointer() string {
	return scn.TracePointerX(scn.ShouldAccountTracePrefix(20))
}

// TraceSequenceX : Provide feedback on where exactly a certain error occurred.
func (scn *Scanner) TraceSequenceX(amount int) string {
	fmt.Println("Check 1")
	if !scn.ShouldAccountTracePrefix(amount) {
		if len(scn.input) < amount {
			return scn.input[0:len(scn.input)]
		}

		return scn.input[0:amount]
	}

	sequence := scn.input[0:]

	return "... " + sequence
}

// TradeSequence : Provide feedback on where exactly a certain error occurred. Returns last 20 processed characters.
func (scn *Scanner) TradeSequence() string {
	return scn.TraceSequenceX(20)
}

// CreateTrace : Creates a trace string with a character pointer.
func (scn *Scanner) CreateTrace() string {
	return "\t" + scn.TradeSequence() + "\n\t" + scn.TracePointer()
}

// Fatal : Report a fatal message. Application will exit.
func (scn *Scanner) Fatal(message string) {
	panic(fmt.Errorf(message + "\n\n" + scn.CreateTrace()))
}

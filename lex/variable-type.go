package lex

// VariableType : Reresents the various different types of values.
type VariableType int

const (
	// VariableTypeString : Represents a string-type value.
	VariableTypeString VariableType = 1

	// VariableTypeInt : Represents an integer-32-type value.
	VariableTypeInt VariableType = 2

	// VariableTypeFloat : Represents a float-type value.
	VariableTypeFloat VariableType = 3
)

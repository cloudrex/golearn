#!/bin/bash
antlr -Dlanguage=Go -o parser Calc.g4
go run ex.go

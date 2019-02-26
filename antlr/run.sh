#!/bin/bash
antlr -Dlanguage=Go -o parser Golearn.g4
go run program.go type.go

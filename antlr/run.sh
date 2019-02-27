#!/bin/bash
export CLASSPATH="/mnt/c/Users/Atlas/go/src/golearn/antlr/antlr-4.7-complete.jar:$CLASSPATH"
alias antlr4='java -Xmx500M -cp "/mnt/c/Users/Atlas/go/src/golearn/antlr/antlr-4.7-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -o parser Golearn.g4
go run program.go type.go linker.go

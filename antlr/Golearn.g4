grammar Golearn;
import GolearnLexer;

// Entry.
start: namespace EOF;

imprt: KeyImport Path SymEnd;

namespace: KeyNamespace Id SymEnd;

expr:
	Atom
	| Atom OpBin Atom
	| OpUnary Atom
	| SymArgsL expr SymArgsR;

arg: Type Id | Type Id SymEnd;

statement: expr SymEnd;

block: SymBlockL statement* SymBlockR;

fn: KeyFn Id block;

strct: KeyStruct Id block;

class: KeyClass Id block;

grammar Golearn;
import GolearnLexer;

// Entry.
start: imprt* namespace (strct | class | fn)* EOF;

imprt: KeyImport Id SymEnd;

namespace: KeyNamespace Id SymEnd;

expr:
	Atom
	| Atom OpBin Atom
	| OpUnary Atom
	| SymArgsL expr SymArgsR;

args: SymArgsL (Type Id SymComma)* SymArgsR;

statement: expr SymEnd | fnx;

block: SymBlockL statement* SymBlockR;

fn: attrib* KeyFn Id args? (SymFnType Type)? block;

fnx: KeyFnx args? (SymFnType FnReturnType)? block;

attrib: SymAttribute Id args?;

strct: KeyStruct Id block;

class: KeyClass Id block;

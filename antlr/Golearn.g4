grammar Golearn;
import GolearnLexer;

// Entry.
start: imprt* namespace (strct | class | fn)* EOF;

assign: Id '=' expr;

declare: Type Id '=' expr | Type Id;

imprt: KeyImport Id SymEnd;

namespace: KeyNamespace Id SymEnd;

expr:
	Atom
	| Atom OpBin Atom
	| OpUnary Atom
	| SymArgsL expr SymArgsR;

arg: Type '*'? Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

statement: expr SymEnd | fnx;

block: SymBlockL statement* SymBlockR;

fn: attrib* KeyFn Id args? (SymFnType Type)? block;

// Anonymous function.
fnx: KeyFnx args? (SymFnType FnReturnType)? block;

attrib: SymAttribute Id args?;

structEntry: Id ':' Type SymEnd;

strct: KeyStruct Id SymBlockL structEntry* SymBlockR;

class: KeyClass Id block;

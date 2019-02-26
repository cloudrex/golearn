grammar Golearn;
import GolearnLexer;

// Entry.
start: imprt* extern* namespace (strct | class | fn)* EOF;

assign: idPath '=' expr;

declare: Type Id '=' expr | Type Id;

imprt: KeyImport Id SymEnd;

namespace: KeyNamespace Id SymEnd;

expr:
	atom
	| Id args // Function call.
	| KeyNew Id args // Class creation.
	| expr OpBin expr // Binary operation.
	| OpUnary expr // Unary operation.
	| SymArgsL expr SymArgsR;

arg: Type Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

statement:
	expr SymEnd
	| fnx SymEnd // Anonymous function.
	| declare SymEnd // Variable declaration.
	| assign SymEnd; // Variable assignment.

block: SymBlockL statement* SymBlockR;

fn: attrib* KeyFn Id args? (SymFnType Type)? block;

// Anonymous function.
fnx: KeyFnx args? (SymFnType FnReturnType)? block;

attrib: SymAttribute Id args?;

structEntry: Id ':' Type SymEnd;

strct: KeyStruct Id SymBlockL structEntry* SymBlockR;

class: KeyClass Id block;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

extern: KeyExtern Id args (SymFnType FnReturnType)? SymEnd;

atom: idPath | NumLiteral | StrLiteral | CharLiteral;

idPath: Id ('.' Id)*;

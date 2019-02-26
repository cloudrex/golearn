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

fn: attrib* KeyFn Modifier? Id args? (SymFnType Type)? block;

// Anonymous function.
fnx: KeyFnx args? (SymFnType Type)? block;

attrib: SymAttribute Id args?;

structEntry: Id ':' Type SymEnd;

strct: KeyStruct Id SymBlockL structEntry* SymBlockR;

constructor: Modifier? Id args block;

class:
	attrib* KeyClass Generic? Extends? Implements* Id SymBlockL constructor? fn* SymBlockR;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

externArgs:
	SymArgsL (Type SymComma)* Type SymArgsR
	| SymArgsL SymArgsR;

extern: KeyExtern Id externArgs (SymFnType Type)? SymEnd;

atom: idPath | NumLiteral | StrLiteral | CharLiteral;

idPath: Id ('.' Id)*;

if: KeyIf SymArgsL expr SymArgsR block;

elseif: KeyElseIf SymArgsL expr SymArgsR block;

else: (if | elseif) KeyElse SymArgsL expr SymArgsR;

for:
	KeyFor SymArgsL expr SymEnd expr SymEnd expr SymArgsR block;

while: KeyWhile SymArgsL expr SymArgsR block;

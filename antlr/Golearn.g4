grammar Golearn;
import GolearnLexer;

// Entry.
start:
	imprt* extern* namespace (
		strct
		| class
		| topLevelFn
		| topLevelDeclare
	)* EOF;

assign: idPath '=' expr;

declare: Type Id '=' expr | Type Id;

topLevelDeclare: KeyExport? declare;

imprt: KeyImport Id SymEnd;

namespace: KeyNamespace Id SymEnd;

expr:
	atom
	| KeyTypeOf expr // Type extraction.
	| Id args // Function call.
	| KeyNew Id args // Class creation.
	| expr OpBin expr // Binary operation.
	| OpUnary expr // Unary operation.
	| KeyAwait expr // Await async operation.
	| SymArgsL expr SymArgsR; // Encapsulated expression within parenthesis.

arg: Type Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

statement:
	expr SymEnd
	| fnx SymEnd // Anonymous function.
	| declare SymEnd // Variable declaration.
	| assign SymEnd // Variable assignment.
	| goto SymEnd // Goto labeled-block statement.
	| KeyReturn expr? SymEnd; // Function return.

block: (Id ':')? SymBlockL statement* SymBlockR;

fnSigArgs: SymArgsL ((Type SymComma)* Type)? SymArgsR;

fnSig: Id fnSigArgs (SymFnType Type)? SymEnd;

fn:
	attrib* KeyFn ModifierStatic? ModifierAsync? Modifier? Id args? (
		SymFnType Type
	)? block;

topLevelFn:
	KeyExport? attrib* KeyFn ModifierAsync? Id args? (
		SymFnType Type
	)? block;

// Anonymous function.
fnx: KeyFnx args? (SymFnType Type)? block;

attrib: SymAttribute Id args?;

structEntry: Id ':' Type SymEnd;

strct: KeyExport? KeyStruct Id SymBlockL structEntry* SymBlockR;

constructor: Modifier? Id args block;

class:
	attrib* KeyExport? KeyClass Generic? Extends? Implements* Id SymBlockL constructor? fn*
		SymBlockR;

interface:
	attrib* KeyExport? KeyInterface Id Implements SymBlockL SymBlockR;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

extern: KeyExtern fnSig;

atom: idPath | NumLiteral | StrLiteral | CharLiteral;

idPath: Id ('.' Id)*;

if: KeyIf SymArgsL expr SymArgsR block;

elseIf: KeyElseIf SymArgsL expr SymArgsR block;

else: (if | elseIf) KeyElse SymArgsL expr SymArgsR;

loopBlock: block | KeyBreak SymEnd | KeyContinue SymEnd;

for:
	KeyFor SymArgsL expr SymEnd expr SymEnd expr SymArgsR loopBlock;

whileHeader: KeyWhile SymArgsL expr SymArgsR;

while: whileHeader loopBlock;

doWhile: KeyDo SymArgsL expr SymArgsR loopBlock whileHeader;

caseBlock: block | KeyBreak SymEnd;

switch:
	KeySwitch SymArgsL expr SymArgsR SymBlockL (
		KeyCase expr ':' caseBlock
	) (KeyDefault ':' caseBlock)?;

goto: KeyGoto Id;

enumEntry: Id ':' atom;

enum:
	KeyEnum Id (KeyExtends Type)? SymBlockL enumEntry* SymBlockR;

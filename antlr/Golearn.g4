grammar Golearn;
import GolearnLexer;

// Entry.
start:
	imprt* extern* namespace entryFn (
		strct
		| class
		| fnTopLevel
		| topLevelDeclare
	)* EOF;

// Basic.
atom: idPath | NumLiteral | StrLiteral | CharLiteral;

expr:
	atom
	| KeyTypeOf expr // Type extraction.
	| Id args // Function call.
	| KeyNew Id args // Class creation.
	| expr OpBin expr // Binary operation.
	| OpUnary expr // Unary operation.
	| KeyAwait expr // Await async operation.
	| KeyInterpolation StrLiteral // String interpolation.
	| expr KeyAs (Type | ComplexType) // Type casting.
	| SymArgsL (Type | ComplexType) SymArgsR expr // Type casting alternative.
	| SymArgsL expr SymArgsR; // Encapsulated expression within parenthesis.

// Variable.
assign: idPath '=' expr;

declare: Type Id '=' expr | Type Id;

topLevelDeclare: KeyExport? declare;

// Importing.
imprt: KeyImport Id SymEnd;

// Namespace.
namespace: KeyNamespace Id SymEnd;

// Block.
statement:
	expr SymEnd
	| fnAnonymous SymEnd // Anonymous function.
	| declare SymEnd // Variable declaration.
	| assign SymEnd // Variable assignment.
	| goto SymEnd // Goto labeled-block statement.
	| KeyExit expr SymEnd // Exit statement.
	| KeyReturn expr? SymEnd; // Function return.

block: (Id ':')? SymBlockL statement* SymBlockR;

// Function.
arg: Type Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

fnSigArgs: SymArgsL ((Type SymComma)* Type)? SymArgsR;

fnSig: Id fnSigArgs (SymFnType Type)? SymEnd;

fn:
	attrib* KeyFn ModifierStatic? ModifierAsync? Modifier? Id args? (
		SymFnType Type
	)? block;

entryFn: attrib* KeyEntry args? (SymFnType Type)? block;

fnTopLevel:
	KeyExport? attrib* KeyFn ModifierAsync? Id args? (
		SymFnType Type
	)? block;

fnAnonymous: KeyFn args? (SymFnType Type)? block;

// Attribute.
attrib: SymAttribute Id args?;

// Struct.
structEntry: Id ':' Type SymEnd;

strct: KeyExport? KeyStruct Id SymBlockL structEntry* SymBlockR;

// Class.
constructor: Modifier? Id args block;

class:
	attrib* KeyExport? KeyClass Generic? Extends? Implements* Id SymBlockL constructor? fn*
		SymBlockR;

// Interface.
interface:
	attrib* KeyExport? KeyInterface Id Implements SymBlockL fnSig* SymBlockR;

// Object.
idPath: Id ('.' Id)*;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

// External declaration.
extern: KeyExport? KeyExtern fnSig;

// If statement.
if: KeyIf SymArgsL expr SymArgsR block;

// Else-if statement.
elseIf: KeyElseIf SymArgsL expr SymArgsR block;

// Else statement.
else: (if | elseIf) KeyElse SymArgsL expr SymArgsR;

// Generic loop.
loopBlock: block | KeyBreak SymEnd | KeyContinue SymEnd;

// For-loop.
for:
	KeyFor SymArgsL expr SymEnd expr SymEnd expr SymArgsR loopBlock;

// While-loop.
whileHeader: KeyWhile SymArgsL expr SymArgsR;

while: whileHeader loopBlock;

// Do-while loop.
doWhile: KeyDo SymArgsL expr SymArgsR loopBlock whileHeader;

// Switch.
caseBlock: block | KeyBreak SymEnd;

switch:
	KeySwitch SymArgsL expr SymArgsR SymBlockL (
		KeyCase expr ':' caseBlock
	) (KeyDefault ':' caseBlock)?;

// Goto.
goto: KeyGoto Id;

// Enum.
enumEntry: Id ':' atom;

enum:
	KeyEnum Id (KeyExtends Type)? SymBlockL enumEntry* SymBlockR;

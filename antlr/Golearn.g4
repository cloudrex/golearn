grammar Golearn;
import GolearnLexer;

// Entry.
start:
	directive* imprt* extern* namespace fnEntry (
		strct
		| iface
		| class
		| fnTopLevel
		| topLevelDeclare
	)* EOF;

// Basic.
atom: idPath | NumLiteral | StrLiteral | CharLiteral;

expr:
	atom
	| declareInline // Inline declaration.
	| KeyTypeOf expr // Type extraction.
	| Id args // Function call.
	| KeyNew Id args // Class creation.
	| expr OpBin expr // Binary operation.
	| OpUnary expr // Unary operation.
	| KeyAwait expr // Await async operation.
	| KeyInterpolation StrLiteral // String interpolation.
	| SymArgsL typeBasic SymArgsR expr // Type casting.
	| SymArgsL expr SymArgsR; // Encapsulated expression within parenthesis.

// Type.
typeBasic: TypeSimple | TypeComplex;

fnReturnType: typeBasic | TypeVoid;

// Variable.
assign: idPath '=' expr;

declare: typeBasic Id '=' expr | typeBasic Id;

declareInline: SymBracketL declare SymBracketR;

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
	| gotoStatement SymEnd // Goto labeled-block statement.
	| KeyExit expr SymEnd // Exit statement.
	| KeyReturn expr? SymEnd // Function return.
	| KeyAssert expr SymEnd // Assert statement.
	| KeyThrow expr SymEnd; // Throw statement.

block: (Id ':')? SymBlockL statement* SymBlockR;

// Function.
arg: typeBasic Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

fnSigArgs: SymArgsL ((typeBasic SymComma)* typeBasic)? SymArgsR;

fnSig: Id fnSigArgs (SymFnType typeBasic)? SymEnd;

fnEntry: attrib* KeyEntry args? (SymFnType fnReturnType)? block;

fn:
	attrib* KeyFn ModifierStatic? ModifierAsync? Modifier? Id args? (
		SymFnType fnReturnType
	)? block;

fnTopLevel:
	KeyExport? attrib* KeyFn ModifierAsync? Id args? (
		SymFnType fnReturnType
	)? block;

fnAnonymous: KeyFn args? (SymFnType typeBasic)? block;

// Attribute.
attrib: SymAttribute Id args?;

// Struct.
structEntry: Id ':' typeBasic SymEnd;

strct: KeyExport? KeyStruct Id SymBlockL structEntry* SymBlockR;

// Class.
constructor: Modifier? Id args block;

class:
	attrib* KeyExport? KeyClass Generic? Extends? Implements* Id SymBlockL constructor? fn*
		SymBlockR;

// Interface.
iface:
	attrib* KeyExport? KeyInterface Id Implements SymBlockL fnSig* SymBlockR;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

// External declaration.
extern:
	KeyExport? KeyExtern (SymArgsL KeyAs Id SymArgsR)? fnSig;

// If statement.
ifStatement: KeyIf SymArgsL expr SymArgsR block;

// Else-if statement.
elseIf: KeyElseIf SymArgsL expr SymArgsR block;

// Else statement.
elseStatement: (ifStatement | elseIf) KeyElse SymArgsL expr SymArgsR;

// Generic loop.
loopBlock: block | KeyBreak SymEnd | KeyContinue SymEnd;

// For-loop.
forLoop:
	KeyFor SymArgsL expr SymEnd expr SymEnd expr SymArgsR loopBlock;

// While-loop.
whileHeader: KeyWhile SymArgsL expr SymArgsR;

while: whileHeader loopBlock;

// Do-while loop.
doWhile: KeyDo SymArgsL expr SymArgsR loopBlock whileHeader;

// Switch.
caseBlock: block | KeyBreak SymEnd;

switchStatement:
	KeySwitch SymArgsL expr SymArgsR SymBlockL (
		KeyCase expr ':' caseBlock
	) (KeyDefault ':' caseBlock)?;

// Goto.
gotoStatement: KeyGoto Id;

// Enum.
enumEntry: Id ':' atom;

enum: KeyEnum Id (KeyExtends)? SymBlockL enumEntry* SymBlockR;

// Directive.
directive: KeyDirective Id (Id | StrLiteral);

// Definition/alias.
def: KeyDef Id '=' (typeBasic (('|' | '&') typeBasic)*);

// Object.
idPath: Id ('.' Id)*;

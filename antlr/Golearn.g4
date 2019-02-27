grammar Golearn;
import GolearnLexer;

// Entry.
start:
	directive* imprt* extern* namespace entryFn (
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
	| SymArgsL anyType SymArgsR expr // Type casting.
	| SymArgsL expr SymArgsR; // Encapsulated expression within parenthesis.

// Type.
typeSimple:
	TypeInt
	| TypeInt64
	| TypeLong
	| TypeShort
	| TypeFloat
	| TypeDouble
	| TypeString
	| TypeObject
	| TypeChar
	| TypeBool;

anyType: typeSimple | TypeComplex;

// Variable.
assign: idPath '=' expr;

declare: anyType Id '=' expr | anyType Id;

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
arg: anyType Id;

args: SymArgsL (arg SymComma)* arg SymArgsR | SymArgsL SymArgsR;

fnSigArgs: SymArgsL ((anyType SymComma)* anyType)? SymArgsR;

fnSig: Id fnSigArgs (SymFnType anyType)? SymEnd;

fn:
	attrib* KeyFn ModifierStatic? ModifierAsync? Modifier? Id args? (
		SymFnType anyType
	)? block;

entryFn:
	attrib* KeyEntry args? (SymFnType (TypeVoid | TypeInt))? block;

fnTopLevel:
	KeyExport? attrib* KeyFn ModifierAsync? Id args? (
		SymFnType anyType
	)? block;

fnAnonymous: KeyFn args? (SymFnType anyType)? block;

// Attribute.
attrib: SymAttribute Id args?;

// Struct.
structEntry: Id ':' anyType SymEnd;

strct: KeyExport? KeyStruct Id SymBlockL structEntry* SymBlockR;

// Class.
constructor: Modifier? Id args block;

class:
	attrib* KeyExport? KeyClass Generic? Extends? Implements* Id SymBlockL constructor? fn*
		SymBlockR;

// Interface.
iface:
	attrib* KeyExport? KeyInterface Id Implements SymBlockL fnSig* SymBlockR;

// Object.
idPath: Id ('.' Id)*;

objLiteralEntry: Id ':' expr;

objLiteral: SymBlockL objLiteralEntry SymBlockR;

// External declaration.
extern:
	KeyExport? (SymArgsL KeyAs Id SymArgsR)? KeyExtern fnSig;

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
def: KeyDef Id '=' (anyType (('|' | '&') anyType)*);

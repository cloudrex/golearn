lexer grammar Golearn;

// General.
NumLiteral: [0-9]+ (.[0-9]+)?;
Whitespace: [ \r\n\t]+ -> skip;
StrLiteral: '"' [^\\"]* '"';
CharLiteral: '\'' [a-zA-Z]? '\'';

// Keywords.
KeyFn: 'fn';
KeyDelete: 'delete';
KeyReturn: 'ret';
KeyYield: 'yield';
KeyFor: 'for';
KeyNamespace: 'space';
KeyVoid: 'void';
KeyStruct: 'struct';
KeyClass: 'class';
KeyImport: 'import';
KeyTrue: 'true';
KeyFalse: 'false';
KeyNil: 'nil';
KeyStatic: 'stat';
KeyAsync: 'async';
KeyConst: 'const';
KeyFnx: 'fnx';
KeyNew: 'new';
KeyExtern: 'extern';
KeyIf: 'if';
KeyElse: 'else';
KeyElseIf: 'elseif';
KeyWhile: 'while';
KeyExport: 'exp';
KeyDo: 'do';
KeySwitch: 'switch';
KeyCase: 'case';
KeyDefault: 'default';
KeyBreak: 'break';
KeyContinue: 'continue';
KeyAwait: 'await';

// Symbols.
SymAttribute: '@';

SymFnType: '~>';

SymEnd: ';';

SymBlockL: '{';

SymBlockR: '}';

SymArgsL: '(';

SymArgsR: ')';

SymComma: ',';

SymArray: '[]';

SymBracketL: '[';

SymBracketR: ']';

SymSpread: '..';

IdList: (Id ',')* Id;

Generic: '<' IdList '>';

Implements: 'implements' IdList;

Extends: ':' Id;

// Operators.
OpBin:
	'*'
	| '/'
	| '+'
	| '-'
	| '%'
	| '^'
	| '<'
	| '>'
	| '<='
	| '>='
	| '!='
	| '&&'
	| '||';

OpUnary: '-' | '!';

// Other.
Modifier: 'pub' | 'priv' | 'prot';

ModifierStatic: 'static';

ModifierAsync: 'async';

Type:
	'int'
	| 'int64'
	| 'long'
	| 'short'
	| 'float'
	| 'double'
	| 'str'
	| 'obj'
	| 'char'
	| 'dyn'
	| 'bool';
// TODO | '*' Type; Also, missing 'void' type.

Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

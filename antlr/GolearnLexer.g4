lexer grammar Golearn;

// General.
NumLiteral: [0-9]+ (.[0-9]+)?;
Whitespace: [ \r\n\t]+ -> skip;
StrLiteral: '"' [^\\"]* '"';
CharLiteral: '\'' [a-zA-Z]? '\'';

// Comments.
CommentSingle: '//' [^\n]* -> skip;
CommentMulti: '/*' [^(*/)] '*/' -> skip;

// Keywords.
KeyFn: 'fn';
KeyDelete: 'delete';
KeyReturn: 'ret';
KeyYield: 'yield';
KeyFor: 'for ';
KeyNamespace: 'space';
KeyStruct: 'struct';
KeyClass: 'class';
KeyImport: 'import';
KeyTrue: 'true';
KeyFalse: 'false';
KeyNil: 'nil';
KeyStatic: 'stat';
KeyAsync: 'async';
KeyConst: 'const';
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
KeyInterface: 'iface';
KeyGoto: 'goto';
KeyTypeOf: 'typeof';
KeyEnum: 'enum';
KeyExtends: 'ext';
KeyThis: 'this';
KeyFrom: 'from';
KeyEntry: 'entry';
KeyAs: 'as';
KeyExit: 'exit';
KeyInterpolation: '$';
KeyUnsigned: 'unsig';
KeyDirective: '#';
KeyDef: 'def';
KeyThrow: 'throw';
KeyAssert: 'assert';

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
Generic: '<' (Id ',')* Id '>';
Implements: 'impl' (Id ',')* Id;
Extends: KeyExtends Id;

// Types.
TypeSimple:
	'int'
	| 'int64'
	| 'long'
	| 'short'
	| 'float'
	| 'double'
	| 'str'
	| 'obj'
	| 'char'
	| 'bool';

TypeVoid: 'void';

TypeComplex:
	'dyn' // Dynamic (infered by compiler).
	| 'type' // Entity type (ex. str, int, etc.).
	| 'ref'; // Function reference.

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
Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

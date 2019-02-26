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
KeyTest: 'test';
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
Implements: 'impl' IdList;
Extends: KeyExtends Id;

// Types.
TypeInt: 'int' | KeyUnsigned 'int';
TypeInt64: 'int64' | KeyUnsigned 'int64';
TypeLong: 'long' | KeyUnsigned 'long'; // Long integer (int128).

TypeShort:
	'short'
	| KeyUnsigned 'short'; // Short integer (int16).

TypeFloat: 'float' | KeyUnsigned 'float';
TypeDouble: 'double' | KeyUnsigned 'double';
TypeString: 'str';
TypeObject: 'obj';
TypeChar: 'char';
TypeBool: 'bool';
TypeVoid: 'void';

TypeComplex:
	'obj'
	| 'dyn' // Dynamic (infered by compiler).
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

// Function reference. TODO | '*' Type; Also, missing 'void' type.

Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

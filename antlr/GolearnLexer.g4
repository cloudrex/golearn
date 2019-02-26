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
FnModifier: 'pub' | 'prot' | 'priv';

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
	| 'bool'
	| '*' Type;

FnReturnType: Type | 'void';

Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

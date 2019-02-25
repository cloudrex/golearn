lexer grammar Golearn;

// Tokens.
Num: [0-9]+;
Whitespace: [ \r\n\t]+ -> skip;

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

// Symbols.
SymAttribute: '@';

SymFnType: '~>';

SymEnd: ';';

SymBlockL: '{';

SymBlockR: '}';

SymArgsL: '(';

SymArgsR: ')';

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

OpUnary: '-' | '!' | '&' | '*';

// Other.
Path: Id '/'?;

FnModifier: 'pub' | 'prot' | 'priv';

Type:
	'int'
	| 'int64'
	| 'int32'
	| 'float'
	| 'double'
	| 'str'
	| 'obj'
	| 'char'
	| 'dyn'
	| 'bool';

// Define at the end to avoid taking precedence.
Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

Atom: Id | Num;

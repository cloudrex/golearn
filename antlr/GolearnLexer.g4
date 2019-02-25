lexer grammar Golearn;

// Tokens
Num: [0-9]+;
Whitespace: [ \r\n\t]+ -> skip;

// Keywords
KeyFn: 'fn';
KeyDelete: 'delete';
KeyReturn: 'ret';
KeyYield: 'yield';
KeyFor: 'for';
KeyNamespace: 'space';
KeyVoid: 'void';

// Symbols
SymAttribute: '@';

SymFnType: '~>';

SymSemicolon: ';';

SymBlockL: '{';

SymBlockR: '}';

SymArgsL: '(';

SymArgsR: ')';

// Operators
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

// Other
FnModifier: 'pub' | 'prot' | 'priv';

FnStatic: 'stat';

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

Const: 'nil' | 'true' | 'false';

// Define at the end to avoid taking precedence.
Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

Atom: Id | Num;

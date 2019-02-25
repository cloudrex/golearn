lexer grammar Golearn;

// General.
Num: [0-9]+ (.[0-9]+)?;
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
	| 'int32'
	| 'float'
	| 'double'
	| 'str'
	| 'obj'
	| 'char'
	| 'dyn'
	| 'bool';

FnReturnType: Type | 'void';

Id: [a-zA-Z]+ [_a-zA-Z0-9]*;

Atom: Id | Num;

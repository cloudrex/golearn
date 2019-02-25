lexer grammar Golearn;
import GolearnParser;

// Tokens
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Keywords
KEY_FN: 'fn';
KEY_DELETE: 'delete';
KEY_RETURN: 'ret';
KEY_YIELD: 'yield';
KEY_FOR: 'for';
KEY_NAMESPACE: 'space';
KEY_VOID: 'void';

// Symbols
SYM_ATTRIB: '@';

SYM_FN_TYPE: '~>';

// Operators
BIN_OP:
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

UNARY_OP: '-' | '!' | '&' | '*';

SYM_END: ';';

SYM_BLOCK_L: '{';

SYM_BLOCK_R: '}';

SYM_ARGS_L: '(';

SYM_ARGS_R: ')';

// Other
FN_MOD: 'pub' | 'prot' | 'priv';

FN_STATIC: 'stat';

TYPE:
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

CONST: 'nil' | 'true' | 'false';

ARG: TYPE ID | TYPE ID ',';

// Define at the end to avoid taking precedence.
ID: [a-zA-Z]+ [_a-zA-Z0-9]*;

VAL: ID | NUM;

EXPR: VAL | VAL BIN_OP VAL | UNARY_OP VAL | '(' EXPR ')';

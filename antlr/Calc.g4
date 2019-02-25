grammar Calc;

// Tokens
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Keywords
KEY_FN: 'fn';
KEY_DELETE: 'delete';
KEY_RETURN: 'ret';
KEY_YIELD: 'yield';
KEY_FOR: 'for';

// Rules
start: EXPR EOF;

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
	| '!=';

UNARY_OP: '-' | '!' | '&' | '*';

OP: BIN_OP | UNARY_OP;

DECLARE: TYPE ASSIGN | TYPE ID;

ASSIGN: ID '=' EXPR;

STATE: EXPR ';';

BODY: STATE;

BLOCK: '{' BODY* '}';

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

VOID: 'void';

CONST: 'nil' | 'true' | 'false';

ARG: TYPE ID | TYPE ID ',';

ARGS_L: '(';

ARGS_R: ')';

ATTRIB: '@' ID;

FN_TYPE_SYMBOL: '~>';

// Define ID at the end to avoid taking precedence.
ID: [a-zA-Z]+ [_a-zA-Z0-9]*;

VAL: ID | NUM;

EXPR: VAL | VAL BIN_OP VAL | UNARY_OP VAL;

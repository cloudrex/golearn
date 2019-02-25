grammar Calc;

// Tokens
NUM: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start: EXPR EOF;

BIN_OP
: '*'
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
;

UNARY_OP
: '-'
| '!'
| '&'
| '*'
;

OP
: BIN_OP
| UNARY_OP
;

DECLARE
: TYPE ASSIGN
| TYPE ID
;

ASSIGN
: ID '=' EXPR
;

STATE: EXPR ';';

BODY: STATE;

BLOCK
: '{' BODY '}'
| '{' '}'
;

FN_MOD
: 'pub'
| 'prot'
| 'priv'
;

FN_STATIC: 'stat';

TYPE
: 'int'
| 'int64'
| 'int32'
| 'float'
| 'str'
| 'obj'
| 'char'
| 'void'
| 'dyn'
;

CONST
: 'nil'
| 'true'
| 'false'
;

ARGS
: '(' ')'
;

FN_PROTO
: ID FN_TYPE_SYMBOL TYPE
| ID
| ID ARGS FN_TYPE_SYMBOL TYPE
| ID ARGS
;

ATTRIB
: '@' ID
| '@' ID ARGS
;

FN_TYPE_SYMBOL: '~>';

FN_KEY: 'fn';

FN
: FN_KEY FN_PROTO BLOCK
| FN_KEY FN_MOD FN_PROTO BLOCK
| FN_KEY FN_STATIC FN_MOD FN_PROTO BLOCK
| ATTRIB FN
;

// Define ID at the end to avoid taking precedence.
ID: [a-zA-Z]+[_a-zA-Z0-9]*;

VAL
: ID
| NUM
;

EXPR
: VAL
| VAL BIN_OP VAL
| UNARY_OP VAL
;

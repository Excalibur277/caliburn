// $antlr-format alignTrailingComments true, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine true, allowShortBlocksOnASingleLine true, minEmptyLines 0, alignSemicolons ownLine
// $antlr-format alignColons trailing, singleLineOverrulesHangingColon true, alignLexerCommands true, alignLabels true, alignTrailers true

lexer grammar CaliburnLexer;

// Punctuation
PERIOD: '.';

COMMA: ',';

COLON: ':';

QUESTION: '?';

Terminator: ';';

ASSIGN: '=';

ARROW: '->';

// Structural Keywords
IF   : 'if';
ELSE : 'else';

FOR: 'for';

SWITCH  : 'switch';
CASE    : 'case';
DEFAULT : 'default';

FALLTHROUGH : 'fallthrough';
BREAK       : 'break';
CONTINUE    : 'continue';
RETURN      : 'return';

USING: 'using';

AS : 'as';
IN : 'in';

// Keywords
NULL: 'null';

VAR   : 'var';
CONST : 'const';
TYPE  : 'type';

FUNC      : 'func';
STRUCT    : 'struct';
TUPLE     : 'tuple';
CLASS     : 'class';
INTERFACE : 'interface';

// Structural Syntax

L_PAREN : '(';
R_PAREN : ')';

L_S_BRACK : '[';
R_S_BRACK : ']';

L_C_BRACK : '{';
R_C_BRACK : '}';

// Operators

OP_ADD : '+';
OP_SUB : '-';
OP_NOT : '!';

OP_MUL : '*';
OP_DIV : '/';
OP_MOD : '%';

OP_POW  : '^';
OP_ROOT : '~';

OP_OR: '|';
OP_XOR:
    '|!&' // TODO: change to better
;
OP_AND: '&';

OP_LSHIFT : '<<';
OP_RSHIFT : '>>';

// Boolean Operators

OP_EQU : '==';
OP_NEQ : '!=';

OP_GTE : '>=';
OP_GRT : '>';
OP_LTE : '<=';
OP_LST : '<';

// Keywords

// Parse out
WhiteSpace: [\p{White_Space}]+ -> skip;

Comment: '`' .*? '`' | '#' ~['"\r\n\\] '\r'? '\n';

// Identifier

Identifier: [\p{L}_] [\p{L}\p{N}_]*;

// Literals
fragment EOL: '\r'? '\n';

fragment ESCAPE_SEQUENCE: '\\' .;

fragment LITERAL_TEXT: (~['"\r\n\\] | ESCAPE_SEQUENCE)+;

fragment SQUOTE: '\'';

fragment DQUOTE: '"';

Literal: SQUOTE LITERAL_TEXT? SQUOTE | DQUOTE LITERAL_TEXT? DQUOTE;
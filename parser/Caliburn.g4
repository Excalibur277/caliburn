// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, alignColons hanging

grammar Caliburn;

// Operators

OP_ADD
    : '+'
    ;

OP_SUB
    : '-'
    ;

OP_MUL
    : '*'
    ;

OP_DIV
    : '/'
    ;

OP_MOD
    : '%'
    ;

OP_POW
    : '^'
    ;

OP_ROOT
    : '~'
    ;

OP_OR
    : '|'
    ;

OP_XOR
    : '|!&' // TODO
    ;

OP_AND
    : '&'
    ;

OP_NOT
    : '!'
    ;

// Boolean Operators

OP_EQU
    : '=='
    ;

OP_NEQ
    : '!='
    ;

OP_GTE
    : '>='
    ;

OP_GRT
    : '>'
    ;

OP_LTE
    : '<='
    ;

OP_LST
    : '<'
    ;

// Keywords

// Parse out
WhiteSpace
    : [\p{White_Space}]+ -> skip
    ;

Comment
    : '`' .*? '`'
    ;

// Identifier

Identifier
    : [\p{L}_] [\p{L}\p{N}_]*
    ;

// Literals
fragment EOL
    : '\r'? '\n'
    ;

fragment ESCAPE_SEQUENCE
    : '\\' .
    ;

fragment LITERAL_TEXT
    : (~['"\r\n\\] | ESCAPE_SEQUENCE)+
    ;

fragment SQUOTE
    : '\''
    ;

fragment DQUOTE
    : '"'
    ;

Literal
    : SQUOTE LITERAL_TEXT? SQUOTE
    | DQUOTE LITERAL_TEXT? DQUOTE
    ;

// Expression
expression
    : operator_expression
    | expression_atom
    ;

operator_expression
    :
    ;

expression_atom
    : value_atom
    | untyped_literal_atom
    | typed_literal_atom
    ;

// Atoms
value_atom
    : Identifier
    ;

type_atom
    : Identifier
    ;

untyped_literal_atom
    : Literal
    ;

typed_literal_atom
    : type_atom Literal
    ;
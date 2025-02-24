// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, singleLineOverrulesHangingColon true, alignColons hanging
parser grammar CaliburnParser;

options {
    tokenVocab = CaliburnLexer;
}

// Module
module
    : definition*
    ;

// Definitions
definition
    : function_definition
    //| class_definition
    //| type_definition
    ;

// Function Declaration Statement
function_definition
    : type = function_type name = Identifier L_PAREN parameters? R_PAREN (
        return_type = type_expression
    )? block
    ;

parameters
    : parameter (COMMA parameter)*
    ;

parameter
    : typed_assign_declaration       # TypedParameter
    | untyped_assign_declaration     # UntypedParameter
    | L_C_BRACK parameters R_C_BRACK # StructDestrutureParameter
    | L_S_BRACK parameters R_S_BRACK # TupleDestrutureParameter
    ;

// Blocks
block
    : L_C_BRACK statement* R_C_BRACK
    ;

// Statement
statement
    : assign_statement
    | expression_statement
    | control_statement
    | jump_statement
    ;

// Jump Statements
jump_statement
    : return_statement
    | break_statement
    | continue_statement
    ;

return_statement
    : RETURN expression Terminator
    ;

break_statement
    : BREAK Terminator
    ;

continue_statement
    : CONTINUE Terminator
    ;

// Control Statement
control_statement
    : if_statement
    | for_statement
    | switch_statement
    ;

// If Statement
if_statement
    : IF inline_statement* expression block (ELSE (block | if_statement))?
    ;

// For Statement
for_statement
    : FOR inline_statement+ expression (ARROW inline_statement+)? block
    ;

// Switch Statement
switch_statement
    : SWITCH inline_statement* expression L_C_BRACK case_statements R_C_BRACK
    ;

case_statements
    : option_case_statement* default_case_statement?
    ;

option_case_statement
    : CASE expression block
    ;

default_case_statement
    : DEFAULT block
    ;

// Inlinable Statements
inline_statement
    : assign_statement
    | expression_statement
    ;

// Assign Statement
assign_statement
    : assign_declarations ASSIGN expressions Terminator
    | assign_expressions op = (
        OP_ADD
        | OP_SUB
        | OP_MUL
        | OP_DIV
        | OP_MOD
        | OP_POW
        | OP_ROOT
        | OP_OR
        | OP_XOR
        | OP_AND
        | OP_LSHIFT
        | OP_RSHIFT
    ) ASSIGN expressions Terminator
    ;

assign_expressions
    : assign_expression (COMMA assign_expression)*
    ;

assign_expression
    : expression                             # ExpressionAssignExpression
    | L_C_BRACK assign_expressions R_C_BRACK # StructDestrutureAssignExpression
    | L_S_BRACK assign_expressions R_S_BRACK # TupleDestrutureAssignExpression
    ;

assign_declarations
    : assign_declaration (COMMA assign_statement)*
    ;

assign_declaration
    : expression                              # ExpressionAssignDeclaration
    | typed_assign_declaration                # TypedAssignDeclaration
    | VAR untyped_assign_declaration          # UntypedAssignDeclaration
    | L_C_BRACK assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
    | L_S_BRACK assign_declarations R_S_BRACK # TupleDestrutureAssignDeclaration
    ;

typed_assign_declarations
    : typed_assign_declaration (COMMA typed_assign_declaration)*
    ;

typed_assign_declaration
    : type = type_expression untyped_assign_declaration # TypedTypedAssignDeclaration
    ;

untyped_assign_declarations
    : untyped_assign_declaration (COMMA untyped_assign_declaration)*
    ;

untyped_assign_declaration
    : var = Identifier                                # UntypedAtomAssignDeclaration
    | L_C_BRACK untyped_assign_declarations R_C_BRACK # UntypedStructDestrutureAssignDeclaration
    | L_S_BRACK untyped_assign_declarations R_S_BRACK # UntypedTupleDestrutureAssignDeclaration
    ;

// Expression Statement
expression_statement
    : expressions Terminator
    ;

// Expression
expressions
    : expression (COMMA expression)*
    ;

expression
    : L_PAREN type = type_expression R_PAREN exp = expression                                              # CastExpression
    | L_PAREN exp = expression R_PAREN                                                                     # BracketedExpression
    | lhs_exp = expression op = (OP_POW | OP_ROOT) rhs_exp = expression                                    # BinaryExpression
    | op = (OP_NOT | OP_ADD | OP_SUB) exp = expression                                                     # UnaryExpression
    | lhs_exp = expression op = (OP_MUL | OP_DIV | OP_MOD) rhs_exp = expression                            # BinaryExpression
    | lhs_exp = expression op = (OP_ADD | OP_SUB) rhs_exp = expression                                     # BinaryExpression
    | lhs_exp = expression op = (OP_LSHIFT | OP_RSHIFT) rhs_exp = expression                               # BinaryExpression
    | lhs_exp = expression op = (OP_EQU | OP_NEQ | OP_LTE | OP_GTE | OP_LST | OP_GRT) rhs_exp = expression # BooleanBinaryExpression
    | lhs_exp = expression OP_AND rhs_exp = expression                                                     # BinaryExpression
    | lhs_exp = expression OP_XOR rhs_exp = expression                                                     # BinaryExpression
    | lhs_exp = expression OP_OR rhs_exp = expression                                                      # BinaryExpression
    | exp = expression L_PAREN args = expressions? R_PAREN                                                 # CallExpression
    | exp = expression PERIOD identifier = Identifier                                                      # AccessExpression
    | exp = expression L_S_BRACK index = expression R_S_BRACK                                              # IndexExpression
    | exp = expression L_S_BRACK start_index = expression? COLON end_index = expression? R_S_BRACK         # SliceExpression
    | identifier = Identifier                                                                              # IdentifierExpression
    | literal = literal_atom                                                                               # LiteralExpression
    | exp = function_expression                                                                            # FunctionExpression
    | exp = struct_expression                                                                              # StructExpression
    | exp = tuple_expression                                                                               # TupleExpression
    ;

// Composite Expressions
function_expression
    : type = function_type L_PAREN assign_declarations? R_PAREN (return_type = type_expression)? block
    ;

struct_expression
    : type = struct_type L_C_BRACK (Identifier) COLON expression (
        COMMA (Identifier) COLON expression
    )* COMMA? R_C_BRACK
    ;

tuple_expression
    : L_PAREN expression (COMMA | (COMMA expression)+) R_PAREN
    ;

// Type Expressions
function_type
    : FUNC
    | type_expression
    ;

struct_type
    : STRUCT
    | type_expression
    ;

tuple_type
    : TUPLE
    | type_expression
    ;

type_expression
    : Identifier
    | type_expression PERIOD Identifier
    ;

// Atoms
identifiers
    : Identifier (COMMA Identifier)*
    ;

literal_atom
    : untyped_literal_atom # UntypedLiteralAtom
    | typed_literal_atom   # TypedLiteralAtom
    ;

untyped_literal_atom
    : Literal
    ;

typed_literal_atom
    : Identifier Literal
    ;
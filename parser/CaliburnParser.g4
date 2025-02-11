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
    : function_declaration_statement
    ;

// Blocks
scoped_block
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
    : RETURN expressions? Terminator
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
    | function_declaration_statement
    ;

// If Statement
if_statement
    : IF inline_statement* expression scoped_block (ELSE (scoped_block | if_statement))?
    ;

// For Statement
for_statement
    : FOR expression scoped_block
    ;

// Switch Statement
switch_statement
    : SWITCH inline_statement* expression L_C_BRACK (
        CASE (type_expression | expression) scoped_block
    )* (DEFAULT scoped_block)? R_C_BRACK
    ;

// Allow for type and value matching

// Inlinable Statements
inline_statement
    : assign_statement
    | expression_statement
    ;

// Function Declaration Statement
function_declaration_statement
    : type = type_expression name = Identifier L_PAREN assign_declarations? R_PAREN (
        return_type = type_expression
    )? scoped_block
    ;

// Assign Statement
assign_statement
    : assign_declarations ASSIGN expressions Terminator
    | operator_assign_statement
    ;

assign_declarations
    : assign_declaration (COMMA assign_statement)*
    ;

assign_declaration
    : declared_assign_declaration             # DeclaredAssignDeclaration
    | undeclared_assign_declaration           # UndeclaredAssignDeclaration
    | L_C_BRACK assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
    | L_S_BRACK assign_declarations L_S_BRACK # TupleDestrutureAssignDeclaration
    ;

declared_assign_declarations
    : declared_assign_declaration (COMMA declared_assign_declaration)*
    ;

declared_assign_declaration
    : type = type_expression undeclared_assign_declaration # TypedDeclaredAssignDeclaration
    | VAR undeclared_assign_declaration                    # UntypedDeclaredAssignDeclaration
    ;

undeclared_assign_declarations
    : undeclared_assign_declaration (COMMA undeclared_assign_declaration)*
    ;

undeclared_assign_declaration
    : var = type_expression                              # UndeclaredAtomAssignDeclaration
    | L_C_BRACK undeclared_assign_declarations R_C_BRACK # UndeclaredStructDestrutureAssignDeclaration
    | L_PAREN undeclared_assign_declarations R_PAREN     # UndeclaredTupleDestrutureAssignDeclaration
    ;

// Operator Assign Statement
operator_assign_statement
    : var = expression op = (
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
    ) ASSIGN expressions Terminator # OperatorAssignment
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
    : type = type_expression L_PAREN assign_declarations? R_PAREN (return_type = type_expression)? scoped_block
    ;

struct_expression
    : type = type_expression L_C_BRACK (Identifier) COLON expression (
        COMMA (Identifier) COLON expression
    )* COMMA? R_C_BRACK
    ;

tuple_expression
    : L_PAREN expression (COMMA | (COMMA expression)+) R_PAREN
    ;

// Type Expressions

type_expression
    : complex_type_expression
    | FUNC
    | STRUCT
    | TUPLE
    ;

complex_type_expression
    : Identifier
    | complex_type_expression PERIOD Identifier
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
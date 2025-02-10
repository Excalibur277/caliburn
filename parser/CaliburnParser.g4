// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, singleLineOverrulesHangingColon true, alignColons hanging
parser grammar CaliburnParser;

options {
    tokenVocab = CaliburnLexer;
}

// Blocks
scoped_block
    : L_C_BRACK statement* R_C_BRACK
    ;

// Statement
statement
    : assign_statement
    | expression_statement
    | control_statement
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
    : SWITCH inline_statement* expression L_C_BRACK (CASE (Identifier | literal_atom) scoped_block)* (
        DEFAULT scoped_block
    )? R_C_BRACK
    ;

// Allow for type and value matching

// Inlinable Statements
inline_statement
    : assign_statement
    | expression_statement
    ;

// Function Declaration Statement
function_declaration_statement
    : type = (FUNC | Identifier) name = Identifier L_PAREN assign_declarations? R_PAREN (
        return_type = (VAR | Identifier)
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
    : type = Identifier undeclared_assign_declaration # TypedDeclaredAssignDeclaration
    | VAR undeclared_assign_declaration               # UntypedDeclaredAssignDeclaration
    ;

undeclared_assign_declarations
    : undeclared_assign_declaration (COMMA undeclared_assign_declaration)*
    ;

undeclared_assign_declaration
    : var = Identifier                                   # UndeclaredAtomAssignDeclaration
    | L_C_BRACK undeclared_assign_declarations R_C_BRACK # UndeclaredStructDestrutureAssignDeclaration
    | L_PAREN undeclared_assign_declarations R_PAREN     # UndeclaredTupleDestrutureAssignDeclaration
    ;

// Operator Assign Statement
operator_assign_statement
    : var = Identifier op = (
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
    ) expressions Terminator # OperatorAssignment
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
    : L_PAREN expression R_PAREN expression                                            # CastExpression
    | L_PAREN expression R_PAREN                                                       # BracketedExpression
    | expression op = (OP_POW | OP_ROOT) expression                                    # ExponentExpression
    | op = (OP_NOT | OP_ADD | OP_SUB) expression                                       # UnaryExpression
    | expression op = (OP_MUL | OP_DIV | OP_MOD) expression                            # MultiplicativeExpression
    | expression op = (OP_ADD | OP_SUB) expression                                     # AdditionExpression
    | expression op = (OP_LSHIFT | OP_RSHIFT) expression                               # ShiftExpression
    | expression op = (OP_EQU | OP_NEQ | OP_LTE | OP_GTE | OP_LST | OP_GRT) expression # ComparisonExpression
    | expression OP_AND expression                                                     # AndExpression
    | expression OP_XOR expression                                                     # XOrExpression
    | expression OP_OR expression                                                      # OrExpression
    | expression L_PAREN expressions? R_PAREN                                          # CallExpression
    | expression PERIOD Identifier                                                     # AccessExpression
    | expression L_S_BRACK expression R_S_BRACK                                        # IndexExpression
    | expression L_S_BRACK expression? COLON expression? R_S_BRACK                     # SliceExpression
    | expression_atom                                                                  # AtomExpression
    | function_expression                                                              # FunctionExpression
    | struct_expression                                                                # StructExpression
    | tuple_expression                                                                 # TupleExpression
    ;

// Type Expressions
function_expression
    : (type = (FUNC | Identifier)) L_PAREN assign_declarations? R_PAREN (
        return_type = (VAR | Identifier)
    )? scoped_block
    ;

struct_expression
    : type = (STRUCT | Identifier) L_C_BRACK (Identifier) COLON expression (
        COMMA (Identifier) COLON expression
    )* COMMA? R_C_BRACK
    ;

tuple_expression
    : L_PAREN expression (COMMA | (COMMA expression)+) R_PAREN
    ;

expression_atom
    : Identifier
    | literal_atom
    ;

// Atoms
identifiers
    : Identifier (COMMA Identifier)*
    ;

literal_atom
    : untyped_literal_atom
    | typed_literal_atom
    ;

untyped_literal_atom
    : Literal
    ;

typed_literal_atom
    : Identifier Literal
    ;
// $antlr-format alignTrailingComments true, columnLimit 150, minEmptyLines 1, maxEmptyLinesToKeep 1, reflowComments false, useTab false
// $antlr-format allowShortRulesOnASingleLine false, allowShortBlocksOnASingleLine true, alignSemicolons hanging, singleLineOverrulesHangingColon true, alignColons hanging
parser grammar CaliburnParser;

options {
    tokenVocab = CaliburnLexer;
}

// Module
module
    : definitions EOF # ModuleRule
    ;

// Definitions
definitions
    :                        # DefinitionsInitial
    | definitions definition # DefinitionsAppend
    ;

definition
    : function_definition
    //| class_definition
    //| type_definition
    ;

// Function Definition Statement
function_definition
    : type = function_type name = Identifier L_PAREN parameters R_PAREN return_type = type_expression block # FunctionDefinition
    | type = function_type name = Identifier L_PAREN parameters R_PAREN return_type = type_expression block # FunctionDefinitionNoReturnType
    ;

parameters
    :                 # ParametersEmpty
    | parameters_list # ParametersFilled
    ;

parameters_list
    : parameter                       # ParametersListInitial
    | parameters_list COMMA parameter # ParametersListAppend
    ;

parameter
    : typed_assign_declaration       # TypedParameter
    | untyped_assign_declaration     # UntypedParameter
    | L_C_BRACK parameters R_C_BRACK # StructDestrutureParameter
    | L_S_BRACK parameters R_S_BRACK # TupleDestrutureParameter
    ;

// Blocks
block
    : L_C_BRACK statement* R_C_BRACK # BlockRule
    ;

// Statement
statements
    :                      # StatementsInitial
    | statements statement # StatementsAppend
    ;

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
    : RETURN expression Terminator # ReturnStatement
    ;

break_statement
    : BREAK Terminator # BreakStatement
    ;

continue_statement
    : CONTINUE Terminator # ContinueStatement
    ;

// Control Statement
control_statement
    : if_statement
    | for_statement
    | switch_statement
    ;

// If Statement
if_statement
    : IF inline_statements expression block                # IfStatement
    | IF inline_statements expression block else_statement # IfElseStatement
    ;

else_statement
    : ELSE block        # ElseStatement
    | ELSE if_statement # ElseIfStatement
    ;

// For Statement
for_statement
    : FOR inline_statements expression block                         # ForStatement
    | FOR inline_statements expression ARROW inline_statements block # ForStatementWithAfter
    ;

// Switch Statement
switch_statement
    : SWITCH inline_statements expression L_C_BRACK case_blocks R_C_BRACK # SwitchStatement
    ;

case_blocks
    : option_case_block* default_case_block? # CaseBlocks
    ;

option_case_block
    : CASE expression block # OptionCaseBlock
    ;

default_case_block
    : DEFAULT block # DefaultCaseBlock
    ;

// Inlinable Statements
inline_statements
    :                                    # InlineStatementsInitial
    | inline_statements inline_statement # InlineStatementsAppend
    ;

inline_statement
    : assign_statement
    | expression_statement
    ;

// Assign Statement
assign_statement
    : assign_declarations ASSIGN expressions Terminator # AssignStatement
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
    ) ASSIGN expressions Terminator # AssignOperationStatement
    ;

assign_expressions
    : assign_expression (COMMA assign_expression)* # AssignExpressions
    ;

assign_expression
    : expression                             # ExpressionAssignExpression
    | L_C_BRACK assign_expressions R_C_BRACK # StructDestrutureAssignExpression
    | L_S_BRACK assign_expressions R_S_BRACK # TupleDestrutureAssignExpression
    ;

assign_declarations
    : assign_declaration (COMMA assign_statement)* # AssignDeclarations
    ;

assign_declaration
    : expression                              # ExpressionAssignDeclaration
    | typed_assign_declaration                # TypedAssignDeclarationDeclaration
    | VAR untyped_assign_declaration          # UntypedAssignDeclarationDeclaration
    | L_C_BRACK assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
    | L_S_BRACK assign_declarations R_S_BRACK # TupleDestrutureAssignDeclaration
    ;

typed_assign_declarations
    : typed_assign_declaration (COMMA typed_assign_declaration)* # TypedAssignDeclarations
    ;

typed_assign_declaration
    : type = type_expression untyped_assign_declaration # TypedAssignDeclaration
    ;

untyped_assign_declarations
    : untyped_assign_declaration (COMMA untyped_assign_declaration)* # UntypedAssignDeclarations
    ;

untyped_assign_declaration
    : var = Identifier                                # UntypedIdentifierAssignDeclaration
    | L_C_BRACK untyped_assign_declarations R_C_BRACK # UntypedStructDestrutureAssignDeclaration
    | L_S_BRACK untyped_assign_declarations R_S_BRACK # UntypedTupleDestrutureAssignDeclaration
    ;

// Expression Statement
expression_statement
    : expressions Terminator # ExpressionStatement
    ;

// Expression
expressions
    : expression (COMMA expression)* # ExpressionsRule
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
    | type = function_type L_PAREN assign_declarations? R_PAREN (return_type = type_expression)? block     # FunctionExpression
    | type = struct_type L_C_BRACK (Identifier) COLON expression (
        COMMA (Identifier) COLON expression
    )* COMMA? R_C_BRACK                                        # StructExpression
    | L_PAREN expression (COMMA | (COMMA expression)+) R_PAREN # TupleExpression
    ;

// Type Expressions
function_type
    : FUNC            # FunctionTypeFunc
    | type_expression # FunctionTypeExpression
    ;

struct_type
    : STRUCT          # StructTypeStruct
    | type_expression # StructTypeExpression
    ;

tuple_type
    : TUPLE           # TupleTypeTuple
    | type_expression # TupleTypeExpression
    ;

type_expression
    : Identifier                        # IdentifierTypeExpression
    | type_expression PERIOD Identifier # AccessTypeExpression
    ;

// Atoms
identifiers
    : Identifier (COMMA Identifier)* # IdentifiersRule
    ;

literal_atom
    : Literal            # UntypedLiteral
    | Identifier Literal # TypedLiteral
    ;
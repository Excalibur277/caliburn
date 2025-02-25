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
    : function_type identifier L_PAREN parameters R_PAREN type_expression block # FunctionDefinition
    | function_type identifier L_PAREN parameters R_PAREN block                 # FunctionDefinitionNoReturnType
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
    ;

// Blocks
block
    : L_C_BRACK statements R_C_BRACK
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

// Inlinable Statements
inline_statements
    :                                    # InlineStatementsInitial
    | inline_statements inline_statement # InlineStatementsAppend
    ;

inline_statement
    : assign_statement
    | expression_statement
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
    : option_case_blocks                    # CaseBlocks
    | option_case_blocks default_case_block # CaseBlocksDefault
    ;

option_case_blocks
    :                                      # OptionCaseBlocksInitial
    | option_case_blocks option_case_block # OptionCaseBlocksAppend
    ;

option_case_block
    : CASE expression block # OptionCaseBlock
    ;

default_case_block
    : DEFAULT block # DefaultCaseBlock
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
    : assign_expression                          # AssignExpressionsInitial
    | assign_expressions COMMA assign_expression # AssignExpressionsAppend
    ;

aliasable_assign_expressions
    : aliasable_assign_expression                                    # AliasableAssignExpressionsInitial
    | aliasable_assign_expressions COMMA aliasable_assign_expression # AliasableAssignExpressionsAppend
    ;

aliasable_assign_expression
    : assign_expression COLON identifier # AliasedAssignExpression
    | assign_expression                  # UnaliasedAssignExpression
    ;

assign_expression
    : expression                                       # ExpressionAssignExpression
    | L_C_BRACK aliasable_assign_expressions R_C_BRACK # StructDestrutureAssignExpression
    ;

assign_declarations
    : assign_declaration                         # AssignDeclarationsInitial
    | assign_declarations COMMA assign_statement # AssignDeclarationsAppend
    ;

aliasable_assign_declarations
    : aliasable_assign_declaration                                     # AliasableAssignDeclarationsInitial
    | aliasable_assign_declarations COMMA aliasable_assign_declaration # AliasableAssignDeclarationsAppend
    ;

aliasable_assign_declaration
    : assign_declaration COLON identifier # AliasedAssignDeclaration
    | assign_declaration                  # UnaliasedAssignDeclaration
    ;

assign_declaration
    : expression                                        # ExpressionAssignDeclaration
    | typed_assign_declaration                          # TypedAssignDeclarationDeclaration
    | VAR untyped_assign_declaration                    # UntypedAssignDeclarationDeclaration
    | L_C_BRACK aliasable_assign_declarations R_C_BRACK # StructDestrutureAssignDeclaration
    ;

typed_assign_declaration
    : type_expression untyped_assign_declaration # TypedAssignDeclaration
    ;

aliasable_untyped_assign_declarations
    : aliasable_untyped_assign_declaration                                             # AliasableUntypedAssignDeclarationsInitial
    | aliasable_untyped_assign_declarations COMMA aliasable_untyped_assign_declaration # AliasableUntypedAssignDeclarationsAppend
    ;

aliasable_untyped_assign_declaration
    : untyped_assign_declaration COLON identifier # AliasedUntypedAssignDeclaration
    | untyped_assign_declaration                  # UnaliasedUntypedAssignDeclaration
    ;

untyped_assign_declaration
    : var = identifier                                          # UntypedIdentifierAssignDeclaration
    | L_C_BRACK aliasable_untyped_assign_declarations R_C_BRACK # UntypedStructDestrutureAssignDeclaration
    ;

// Expression Statement
expression_statement
    : expressions Terminator # ExpressionStatement
    ;

// Expression
expressions
    : expression                   # ExpressionsInitial
    | expressions COMMA expression # ExpressionsAppend
    ;

expressions_optional
    : expressions # ExpressionsOptional
    |             # ExpressionsOptionalNone
    ;

expression
    : L_PAREN expression R_PAREN                                                       # BracketedExpression
    | expression op = (OP_POW | OP_ROOT) expression                                    # BinaryExpression
    | op = (OP_NOT | OP_ADD | OP_SUB) expression                                       # UnaryExpression
    | expression op = (OP_MUL | OP_DIV | OP_MOD) expression                            # BinaryExpression
    | expression op = (OP_ADD | OP_SUB) expression                                     # BinaryExpression
    | expression op = (OP_LSHIFT | OP_RSHIFT) expression                               # BinaryExpression
    | expression op = (OP_EQU | OP_NEQ | OP_LTE | OP_GTE | OP_LST | OP_GRT) expression # BinaryExpression
    | expression OP_AND expression                                                     # BinaryExpression
    | expression OP_XOR expression                                                     # BinaryExpression
    | expression OP_OR expression                                                      # BinaryExpression
    | expression L_PAREN expressions_optional R_PAREN                                  # CallExpression
    | expression PERIOD identifier                                                     # AccessExpression
    | expression L_S_BRACK expression R_S_BRACK                                        # IndexExpression
    | expression L_S_BRACK expression COLON R_S_BRACK                                  # SliceStartExpression
    | expression L_S_BRACK COLON expression R_S_BRACK                                  # SliceEndExpression
    | expression L_S_BRACK expression COLON expression R_S_BRACK                       # SliceExpression
    | identifier                                                                       # IdentifierExpression
    | literal_atom                                                                     # LiteralExpression
    | function_type L_PAREN parameters R_PAREN type_expression block                   # FunctionExpression
    | function_type L_PAREN parameters R_PAREN block                                   # FunctionExpressionNoReturnType
    | struct_type L_C_BRACK struct_values R_C_BRACK                                    # StructExpression
    | type_expression L_S_BRACK collection_values R_S_BRACK                            # CollectionExpression
    ;

struct_values
    : named_struct_values       # StructValuesNamed
    | named_struct_values COMMA # StructValuesNamed
    | expressions_optional      # StructValuesUnamed
    ;

named_struct_values
    : named_struct_value                           # NamedStructValuesInitial
    | named_struct_values COMMA named_struct_value # NamedStructValuesAppend
    ;

named_struct_value
    : identifier COLON expression # NamedStructValue
    ;

collection_values
    :                                          # CollectionValuesInitial
    | collection_values COMMA collection_value # CollectionValuesAppend
    ;

collection_value
    : expression COLON expression # CollectionValue
    ;

// Type Expressions
function_type
    : FUNC            # FunctionTypeFunc
    | type_expression # FunctionTypeExpression
    ;

struct_type
    : STRUCT          # StructTypeStruct
    | TUPLE           # StructTypeTuple
    | type_expression # StructTypeExpression
    ;

type_expression
    : identifier                        # IdentifierTypeExpression
    | type_expression PERIOD identifier # AccessTypeExpression
    ;

// Atoms
literal_atom
    : literal                 # UntypedLiteral
    | type_expression literal # TypedLiteral
    ;

literal
    : val = LiteralToken
    ;

identifier
    : val = IdentifierToken
    ;
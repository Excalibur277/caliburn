%{
package parser

import (
    "caliburncode/nodes"
)
%}

// Value Tokens
%token IDENTIFIER
%token LITERAL

// Syntax Tokens
%token PERIOD
%token COMMA
%token QUESTION
%token ASSIGN
%token COLON

%token TERMINATOR

// Structural Tokens
%token L_PAREN
%token R_PAREN
%token L_S_BRAC
%token R_S_BRAC
%token L_C_BRAC
%token R_C_BRAC

// Type Keywords
%token FUNC
%token STRUCT
%token TUPLE
%token CLASS

// Control Keywords
%token IF
%token ELSE
%token FOR
%token SWITCH
%token CASE
%token DEFAULT

// Jump Keywords
%token BREAK
%token CONTINUE
%token RETURN


%union{
    identifier_token    Identifier
    literal_token       Literal
    token               nodes.Token
    module              nodes.Module
    definitions[]       []nodes.Definition
    definition          nodes.Definition
    block               nodes.Block
    statements[]        []nodes.Statement
    statement           nodes.Statement
    expressions[]       []nodes.Expression
    expression          nodes.Expression
    parameters          []nodes.Parameter
    parameter           nodes.Parameter
    identifier          nodes.Identifier
    generaltype         nodes.Type
    type_expression     nodes.TypeExpression
}

// Value Tokens
%type<identifier_token> IDENTIFIER
%type<literal_token> LITERAL

// Token
%type<token> TERMINATOR
%type<token> FUNC, STRUCT, TUPLE, CLASS
%type<token> L_PAREN, R_PAREN, L_S_BRAC, R_S_BRAC, L_C_BRAC, R_C_BRAC
%type<token> PERIOD, COMMA, QUESTION, ASSIGN, COLON
%type<token> IF, ELSE, FOR, SWITCH, CASE
%type<token> BREAK, CONTINUE, RETURN

// Module
%type<module> module

// Definiton
%type<definitions> definitions
%type<definition> definition, function_definition //, class_definition, type_definition

// Parameter
%type<parameters> parameters, set_parameters
%type<parameter> parameter

// Block
%type<block> block

// Statement
%type<statements> statements, inline_statements
%type<statement> statement, inline_statement, assign_statement, expression_statement

// Jump Statements
%type<statement> jump_statement, break_statement, continue_statement, return_statement

// Control Statements
%type<statement> control_statement, if_statement, else_statement, for_statement, switch_statement

// Expression
%type<expressions> expressions
%type<expression> expression

// Identifier
%type<identifier> identifier

// Types
%type<generaltype> type, function_type, return_type
%type<type_expression> type_expression


%start module

%%
// Top Level Module
module: 
    definitions {
        $$ = nodes.NewModule()
    };

// Definitions
definitions: 
    definitions definition {
        $$ = append($1, $2)
    } | 
    {
        $$ = []nodes.Definition{}
    };

definition:
    function_definition {
        $$ = $1
    }; //| 
    //class_definition {
    //    $$ = $1
    //} | 
    //type_definition {
    //    $$ = $1
    //};

// Function Definition
function_definition:
    function_type identifier L_PAREN parameters R_PAREN return_type block {
        $$ = $1
    };

function_type:
    FUNC {
        $$ = nodes.NewImplicitFunctionType()
    } |
    type {
        $$ = nodes.NewFunctionType($1)
    };

return_type:
    type {
        $$ = nodes.NewReturnType($1)
    } | 
    {
        $$ = nil
    };

parameters:
    set_parameters {
        $$ = $1
    } | 
    {
        $$ = []nodes.FunctionParameter{}
    };

set_parameters:
    set_parameters COMMA parameter {
        $$ = append($1, $3)
    } |
    parameter {
        $$ = []nodes.FunctionParameter{$1}
    };

parameter:
    identifier {
        $$ = nodes.NewImplicitFunctionParameter($1)
    } |
    type identifier {
        $$ = nodes.NewTypedFunctionParameter($2, $1)
    };

// Blocks
block:
    L_C_BRAC statements R_C_BRAC {
        $$ = nodes.NewBlock($2)
    };

// Statements
statements:
    statements statement {
        $$ = append($1, $2)
    } | 
    {
        $$ = []nodes.Statement{}
    };

statement:
    definition {
        $$ = nodes.NewDefinitionStatement($1)
    } |
    assign_statement {
        $$ = $1
    } |
    expression_statement {
        $$ = $1
    } |
    control_statement {
        $$ = $1
    } |
    jump_statement {
        $$ = $1
    };

expression_statement:
    expressions TERMINATOR {
        $$ = nodes.NewExpressionStatement($1)
    };

assign_statement:
    expressions ASSIGN expressions {
        $$ = nodes.NewExpressionStatement($1)
    };

// Jump Statements
jump_statement:
    break_statement {
        $$ = $1
    } |
    continue_statement {
        $$ = $1
    } | 
    return_statement {
        $$ = $1
    };

break_statement:
    BREAK TERMINATOR {
        $$ = nodes.NewBreakStatement();
    };

continue_statement:
    CONTINUE TERMINATOR {
        $$ = nodes.NewContinueStatement();
    };

return_statement:
    RETURN expression TERMINATOR {
        $$ = nodes.NewReturnStatement($2);
    };

// Control Statements
control_statement:
    if_statement {
        $$ = $1
    } |
    for_statement {
        $$ = $1
    } |
    switch_statement {
        $$ = $1
    };

if_statement:
    IF inline_statements expression block else_statement {
        $$ = nodes.NewIfStatement($2, $3, $4, $5)
    };

else_statement:
    {
        $$ = nodes.NewEmptyStatement()
    } |
    ELSE block {
        $$ = nodes.NewElseStatement()
    } | 
    ELSE if_statement {
        $$ = $2
    };

for_statement:
    FOR expression block {
        $$ = nodes.NewForStatement($2, $3)
    };


// Inline Statements
inline_statements:
    inline_statements inline_statement {
        $$ = append($1, $2)
    } | 
    { 
        $$ = []nodes.Statement{}   
    }; 

inline_statement:
    assign_statement {
        $$ = $1
    } |
    expression_statement {
        $$ = $1
    };

// Expressions
expressions:
    expressions COMMA expression {
        $$ = append($1, $3)
    } |
    expression {
        $$ = []nodes.Expression{$1}
    };

expression:
    identifier {
        $$ = nodes.NewIdentifierExpression($1)
    };

identifier:
    IDENTIFIER {
        $$ = nodes.NewIdentifier($1)
    };

type:
    type_expression {
        $$ = nodes.NewType($1)
    };

type_expression:
    IDENTIFIER {
        $$ = nodes.NewTypeExpressionIdentifier($1)
    } |
    type_expression PERIOD IDENTIFIER {
        $$ = nodes.NewTypeExpressionAccessed($1, $3)
    };

%%
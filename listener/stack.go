package listener

import (
	"caliburncode.com/m/ast"
)

type parseData struct {
	// Base Tokens
	identifier_token ast.Identifier
	literal_token    ast.Literal
	token            ast.Token

	// Main Structure
	module      ast.Module
	definitions [][]ast.Definition
	definition  ast.Definition
	block       ast.Block
	statements  [][]ast.Statement
	statement   ast.Statement
	expressions [][]ast.Expression
	expression  ast.Expression

	// Function Types
	parameters []ast.Parameter
	parameter  ast.Parameter

	// Expressions
	type_expression ast.TypeExpression

	// Base Ast
	identifier  ast.Identifier
	generaltype ast.Type

	// Case Statements
	case_statements        ast.CaseStatements
	default_case_statement ast.DefaultCaseStatement
	option_case_statements []ast.OptionCaseStatement
	option_case_statement  ast.OptionCaseStatement

	// Assignment
	assign_declarations         []ast.AssignDeclaration
	assign_declaration          ast.AssignDeclaration
	typed_assign_declaration    ast.TypedAssignDeclaration
	untyped_assign_declarations []ast.UntypedAssignDeclaration
	untyped_assign_declaration  ast.UntypedAssignDeclaration
}

type parseStack []*parseData

func (s *parseStack) Pop() (item *parseData) {
	if len(*s) == 0 {
		return nil
	}
	item, *s = (*s)[len(*s)-1], (*s)[0:len(*s)-1]
	return item
}

func (s *parseStack) Push(d *parseData) {
	*s = append(*s, d)
}

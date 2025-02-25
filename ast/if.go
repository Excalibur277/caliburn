package ast

import "fmt"

type IfStatement interface {
	Node
	IsIfStatement()
}

type IfStatementFinal struct {
	StatementBase
	priorStatements []InlineStatement
	expression      Expression
	block           Block
}

func NewIfStatement(priorStatements []InlineStatement, expression Expression, block Block) *IfStatementFinal {
	return &IfStatementFinal{priorStatements: priorStatements, expression: expression, block: block}
}

func (is *IfStatementFinal) IsIfStatement() {}
func (is *IfStatementFinal) String() string {
	return fmt.Sprintf("if %s %s %s", SliceToString(is.priorStatements, " "), is.expression, is.block)
}

type IfElseStatement struct {
	IfStatementFinal
	elseStatement ElseStatement
}

func NewIfElseStatement(priorStatements []InlineStatement, expression Expression, block Block, elseStatement ElseStatement) *IfElseStatement {
	return &IfElseStatement{IfStatementFinal: *NewIfStatement(priorStatements, expression, block), elseStatement: elseStatement}
}
func (ies *IfElseStatement) String() string {
	return fmt.Sprintf("if %s %s", ies.IfStatementFinal.String(), ies.elseStatement)
}

type ElseStatement interface {
	Node
	IsElseStatement()
}

type ElseStatementFinal struct {
	StatementBase
	block Block
}

func NewElseStatement(block Block) *ElseStatementFinal {
	return &ElseStatementFinal{block: block}
}

func (es *ElseStatementFinal) IsElseStatement() {}
func (es *ElseStatementFinal) String() string {
	return fmt.Sprintf("else %s", es.block)
}

type ElseIfStatement struct {
	StatementBase
	ifStatement IfStatement
}

func NewElseIfStatement(ifStatement IfStatement) *ElseIfStatement {
	return &ElseIfStatement{ifStatement: ifStatement}
}

func (es *ElseIfStatement) IsElseStatement() {}
func (es *ElseIfStatement) String() string {
	return fmt.Sprintf("else if %s", es.ifStatement)
}

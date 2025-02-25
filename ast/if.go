package ast

type IfStatement interface {
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

type IfElseStatement struct {
	IfStatementFinal
	elseStatement ElseStatement
}

func NewIfElseStatement(priorStatements []InlineStatement, expression Expression, block Block, elseStatement ElseStatement) *IfElseStatement {
	return &IfElseStatement{IfStatementFinal: *NewIfStatement(priorStatements, expression, block), elseStatement: elseStatement}
}

type ElseStatement interface {
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

type ElseIfStatement struct {
	StatementBase
	ifStatement IfStatement
}

func NewElseIfStatement(ifStatement IfStatement) *ElseIfStatement {
	return &ElseIfStatement{ifStatement: ifStatement}
}

func (es *ElseIfStatement) IsElseStatement() {}

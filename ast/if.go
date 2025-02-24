package ast

type IfStatement interface {
	Statement
	IsIfStatement()
}

type IfStatementBase struct{ StatementBase }

func (is *IfStatementBase) IsIfStatement() {}

type ElseStatement interface {
	Statement
	IsElseStatement()
}

type ElseStatementBase struct{ StatementBase }

func (es *ElseStatementBase) IsElseStatement() {}

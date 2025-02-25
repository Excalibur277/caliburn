package ast

type Statement interface {
	IsStatement()
}

type StatementBase struct {
}

func (db *StatementBase) IsStatement() {}

type InlineStatement interface {
	IsInlineStatement()
}
type InlineStatementBase struct {
}

func (db *InlineStatementBase) IsInlineStatement() {}

type ReturnStatement struct {
	StatementBase
	expression Expression
}

func NewReturnStatement(expression Expression) *ReturnStatement {
	return &ReturnStatement{expression: expression}
}

type BreakStatement struct {
	StatementBase
}

func NewBreakStatement() *BreakStatement {
	return &BreakStatement{}
}

type ContinueStatement struct {
	StatementBase
}

func NewContinueStatement() *ContinueStatement {
	return &ContinueStatement{}
}

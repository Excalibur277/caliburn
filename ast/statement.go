package ast

import "fmt"

type Statement interface {
	Node
	IsStatement()
}

type StatementBase struct {
}

func (db *StatementBase) IsStatement() {}

type InlineStatement interface {
	Node
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

func (rs *ReturnStatement) String() string { return fmt.Sprintf("return %s;", rs.expression) }

type BreakStatement struct {
	StatementBase
}

func NewBreakStatement() *BreakStatement {
	return &BreakStatement{}
}

func (rs *BreakStatement) String() string { return "break;" }

type ContinueStatement struct {
	StatementBase
}

func NewContinueStatement() *ContinueStatement {
	return &ContinueStatement{}
}

func (cs *ContinueStatement) String() string { return "continue;" }

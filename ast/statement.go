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

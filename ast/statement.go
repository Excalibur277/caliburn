package ast

type Statement interface {
	IsStatement()
}

type StatementBase struct {
}

func (db *StatementBase) IsStatement() {}

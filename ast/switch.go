package ast

type SwitchStatement struct{ StatementBase }

func NewSwitchStatement() {

}

type CaseStatements interface {
	IsCaseStatements()
}

type CaseStatementsBase struct{}

func (cs *CaseStatementsBase) IsCaseStatements() {}

type DefaultCaseStatement interface {
	IsDefaultCaseStatement()
}

type DefaultCaseStatementBase struct{}

func (dcs *DefaultCaseStatementBase) IsDefaultCaseStatement() {}

type OptionCaseStatement interface {
	IsOptionCaseStatement()
}

type OptionCaseStatementBase struct{}

func (dcs *OptionCaseStatementBase) IsOptionCaseStatement() {}

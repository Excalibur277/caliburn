package ast

type ForStatement interface {
	Statement
	IsForStatement()
}

type ForStatementBase struct {
	StatementBase
	priorStatements []InlineStatement
	expression      Expression
	block           Block
}

func NewForStatement(priorStatements []InlineStatement, expression Expression, block Block) *ForStatementBase {
	return &ForStatementBase{priorStatements: priorStatements, expression: expression, block: block}
}

func (is *ForStatementBase) IsForStatement() {}

type ForStatementWithPost struct {
	ForStatementBase
	postStatements []InlineStatement
}

func NewForStatementWithPost(priorStatements []InlineStatement, expression Expression, postStatements []InlineStatement, block Block) *ForStatementWithPost {
	return &ForStatementWithPost{ForStatementBase: *NewForStatement(priorStatements, expression, block), postStatements: postStatements}
}

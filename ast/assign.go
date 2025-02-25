package ast

type Aliasable interface {
	IsAliasable()
}
type AssignExpression interface {
	IsAssignExpression()
}
type AliasableAssignExpression interface {
	AssignExpression
	Aliasable
}
type AssignDeclaration interface {
	IsAssignDeclaration()
}
type AliasableAssignDeclaration interface {
	AssignDeclaration
	Aliasable
}

type TypedAssignDeclaration interface {
	IsTypedAssignDeclaration()
}
type UntypedAssignDeclaration interface {
	IsUntypedAssignDeclaration()
}

type AliasableUntypedAssignDeclaration interface {
	UntypedAssignDeclaration
	Aliasable
}

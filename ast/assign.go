package ast

type AssignDeclaration interface {
	IsAssignDeclaration()
}

type TypedAssignDeclaration interface {
	IsTypedAssignDeclaration()
}

type UntypedAssignDeclaration interface {
	IsUntypedAssignDeclaration()
}

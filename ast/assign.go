package ast

type AssignStatement struct {
	StatementBase
	assignDeclarations []AssignDeclaration
	expressions        []Expression
}

func NewAssignStatement(assignDeclarations []AssignDeclaration, expressions []Expression) *AssignStatement {
	return &AssignStatement{assignDeclarations: assignDeclarations, expressions: expressions}
}

type AssignOperationStatement struct {
	StatementBase
	assignExpressions []AssignExpression
	binaryOperation   BinaryOperation
	expressions       []Expression
}

func NewAssignOperationStatement(assignExpressions []AssignExpression, binaryOperation BinaryOperation, expressions []Expression) *AssignOperationStatement {
	return &AssignOperationStatement{assignExpressions: assignExpressions, binaryOperation: binaryOperation, expressions: expressions}
}

type Aliasable interface {
	IsAliasable()
}

type AliasableBase struct{}

func (a *AliasableBase) IsAliasable() {}

type AliasableAssignExpression interface {
	AssignExpression
	Aliasable
}

type AliasedAssignExpression struct {
	AliasableBase
	AssignExpression
	fromIdentifier Identifier
}

func NewAliasedAssignExpression(assignExpression AssignExpression, fromIdentifier Identifier) *AliasedAssignExpression {
	return &AliasedAssignExpression{AssignExpression: assignExpression, fromIdentifier: fromIdentifier}
}

type UnaliasedAssignExpression struct {
	AliasableBase
	AssignExpression
}

func NewUnaliasedAssignExpression(assignExpression AssignExpression) *UnaliasedAssignExpression {
	return &UnaliasedAssignExpression{AssignExpression: assignExpression}
}

type AssignExpression interface {
	IsAssignExpression()
}

type AssignExpressionBase struct{}

func (ae *AssignExpressionBase) IsAssignExpression() {}

type ExpressionAssignExpression struct {
	AssignExpressionBase
	expression Expression
}

func NewExpressionAssignExpression(expression Expression) *ExpressionAssignExpression {
	return &ExpressionAssignExpression{expression: expression}
}

type StructDestrutureAssignExpression struct {
	AssignExpressionBase
	expressions []AliasableAssignExpression
}

func NewStructDestrutureAssignExpression(expressions []AliasableAssignExpression) *StructDestrutureAssignExpression {
	return &StructDestrutureAssignExpression{expressions: expressions}
}

type AliasableAssignDeclaration interface {
	AssignDeclaration
	Aliasable
}

type AliasedAssignDeclaration struct {
	AliasableBase
	AssignDeclaration
	fromIdentifier Identifier
}

func NewAliasedAssignDeclaration(assignDeclaration AssignDeclaration, fromIdentifier Identifier) *AliasedAssignDeclaration {
	return &AliasedAssignDeclaration{AssignDeclaration: assignDeclaration, fromIdentifier: fromIdentifier}
}

type UnaliasedAssignDeclaration struct {
	AliasableBase
	AssignDeclaration
}

func NewUnaliasedAssignDeclaration(assignDeclaration AssignDeclaration) *UnaliasedAssignDeclaration {
	return &UnaliasedAssignDeclaration{AssignDeclaration: assignDeclaration}
}

type AssignDeclaration interface {
	IsAssignDeclaration()
}

type AssignDeclarationBase struct{}

func (ae *AssignDeclarationBase) IsAssignExpression() {}

type ExpressionAssignDeclaration struct {
	AssignDeclarationBase
	expression Expression
}

func NewExpressionAssignDeclaration(expression Expression) *ExpressionAssignDeclaration {
	return &ExpressionAssignDeclaration{expression: expression}
}

type TypedAssignDeclarationDeclaration struct {
	AssignDeclarationBase
	typedAssignDeclaration TypedAssignDeclaration
}

func NewTypedAssignDeclarationDeclaration(typedAssignDeclaration TypedAssignDeclaration) *TypedAssignDeclarationDeclaration {
	return &TypedAssignDeclarationDeclaration{typedAssignDeclaration: typedAssignDeclaration}
}

type UntypedAssignDeclarationDeclaration struct {
	AssignDeclarationBase
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewUntypedAssignDeclarationDeclaration(untypedAssignDeclaration UntypedAssignDeclaration) *UntypedAssignDeclarationDeclaration {
	return &UntypedAssignDeclarationDeclaration{untypedAssignDeclaration: untypedAssignDeclaration}
}

type StructDestrutureAssignDeclaration struct {
	AssignDeclarationBase
	declarations []AliasableAssignDeclaration
}

func NewStructDestrutureAssignDeclaration(declarations []AliasableAssignDeclaration) *StructDestrutureAssignDeclaration {
	return &StructDestrutureAssignDeclaration{declarations: declarations}
}

type TypedAssignDeclaration interface {
	IsTypedAssignDeclaration()
}

type TypedAssignDeclarationBase struct {
	typeExpression           TypeExpression
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewTypedAssignDeclaration(typeExpression TypeExpression, untypedAssignDeclaration UntypedAssignDeclaration) *TypedAssignDeclarationBase {
	return &TypedAssignDeclarationBase{typeExpression: typeExpression, untypedAssignDeclaration: untypedAssignDeclaration}
}

type AliasableUntypedAssignDeclaration interface {
	UntypedAssignDeclaration
	Aliasable
}

type AliasedUntypedAssignDeclaration struct {
	AliasableBase
	UntypedAssignDeclaration
	fromIdentifier Identifier
}

func NewAliasedUntypedAssignDeclaration(declaration UntypedAssignDeclaration, fromIdentifier Identifier) *AliasedUntypedAssignDeclaration {
	return &AliasedUntypedAssignDeclaration{UntypedAssignDeclaration: declaration, fromIdentifier: fromIdentifier}
}

type UnaliasedUntypedAssignDeclaration struct {
	AliasableBase
	UntypedAssignDeclaration
}

func NewUnaliasedUntypedAssignDeclaration(declaration UntypedAssignDeclaration) *UnaliasedUntypedAssignDeclaration {
	return &UnaliasedUntypedAssignDeclaration{UntypedAssignDeclaration: declaration}
}

type UntypedAssignDeclaration interface {
	IsUntypedAssignDeclaration()
}

type UntypedIdentifierAssignDeclaration struct {
	AssignDeclarationBase
	identifier Identifier
}

func NewUntypedIdentifierAssignDeclaration(identifier Identifier) *UntypedIdentifierAssignDeclaration {
	return &UntypedIdentifierAssignDeclaration{identifier: identifier}
}

type UntypedStructDestrutureAssignDeclaration struct {
	AssignDeclarationBase
	declarations []AliasableUntypedAssignDeclaration
}

func NewUntypedStructDestrutureAssignDeclaration(declarations []AliasableUntypedAssignDeclaration) *UntypedStructDestrutureAssignDeclaration {
	return &UntypedStructDestrutureAssignDeclaration{declarations: declarations}
}

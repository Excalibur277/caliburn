package ast

import "fmt"

type AssignStatement struct {
	StatementBase
	assignDeclarations []AssignDeclaration
	expressions        []Expression
}

func NewAssignStatement(assignDeclarations []AssignDeclaration, expressions []Expression) *AssignStatement {
	return &AssignStatement{assignDeclarations: assignDeclarations, expressions: expressions}
}
func (as *AssignStatement) String() string {
	return fmt.Sprintf("%s = %s;", SliceToString(as.assignDeclarations, ", "), SliceToString(as.expressions, ", "))
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
func (aos *AssignOperationStatement) String() string {
	return fmt.Sprintf("%s %s= %s;", SliceToString(aos.assignExpressions, ", "), aos.binaryOperation, SliceToString(aos.expressions, ", "))
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
	AssignExpressionBase
	assignExpression AssignExpression
	fromIdentifier   Identifier
}

func NewAliasedAssignExpression(assignExpression AssignExpression, fromIdentifier Identifier) *AliasedAssignExpression {
	return &AliasedAssignExpression{assignExpression: assignExpression, fromIdentifier: fromIdentifier}
}

func (aae *AliasedAssignExpression) String() string {
	return fmt.Sprintf("%s: %s", aae.assignExpression, aae.fromIdentifier)
}

type UnaliasedAssignExpression struct {
	AliasableBase
	AssignExpressionBase
	assignExpression AssignExpression
}

func NewUnaliasedAssignExpression(assignExpression AssignExpression) *UnaliasedAssignExpression {
	return &UnaliasedAssignExpression{assignExpression: assignExpression}
}

func (uae *UnaliasedAssignExpression) String() string {
	return uae.assignExpression.String()
}

type AssignExpression interface {
	Node
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

func (eae *ExpressionAssignExpression) String() string {
	return eae.expression.String()
}

type StructDestrutureAssignExpression struct {
	AssignExpressionBase
	expressions []AliasableAssignExpression
}

func NewStructDestrutureAssignExpression(expressions []AliasableAssignExpression) *StructDestrutureAssignExpression {
	return &StructDestrutureAssignExpression{expressions: expressions}
}

func (sdae *StructDestrutureAssignExpression) String() string {
	return fmt.Sprintf("{%s}", SliceToString(sdae.expressions, ", "))
}

type AliasableAssignDeclaration interface {
	AssignDeclaration
	Aliasable
}

type AliasedAssignDeclaration struct {
	AliasableBase
	AssignDeclarationBase
	assignDeclaration AssignDeclaration
	fromIdentifier    Identifier
}

func NewAliasedAssignDeclaration(assignDeclaration AssignDeclaration, fromIdentifier Identifier) *AliasedAssignDeclaration {
	return &AliasedAssignDeclaration{assignDeclaration: assignDeclaration, fromIdentifier: fromIdentifier}
}

func (aad *AliasedAssignDeclaration) String() string {
	return fmt.Sprintf("%s : %s", aad.assignDeclaration, aad.fromIdentifier)
}

type UnaliasedAssignDeclaration struct {
	AliasableBase
	AssignDeclarationBase
	assignDeclaration AssignDeclaration
}

func NewUnaliasedAssignDeclaration(assignDeclaration AssignDeclaration) *UnaliasedAssignDeclaration {
	return &UnaliasedAssignDeclaration{assignDeclaration: assignDeclaration}
}

func (uad *UnaliasedAssignDeclaration) String() string {
	return uad.assignDeclaration.String()
}

type AssignDeclaration interface {
	Node
	IsAssignDeclaration()
}

type AssignDeclarationBase struct{}

func (ae *AssignDeclarationBase) IsAssignDeclaration() {}

type ExpressionAssignDeclaration struct {
	AssignDeclarationBase
	expression Expression
}

func NewExpressionAssignDeclaration(expression Expression) *ExpressionAssignDeclaration {
	return &ExpressionAssignDeclaration{expression: expression}
}

func (ead *ExpressionAssignDeclaration) String() string {
	return ead.expression.String()
}

type TypedAssignDeclarationDeclaration struct {
	AssignDeclarationBase
	typedAssignDeclaration TypedAssignDeclaration
}

func NewTypedAssignDeclarationDeclaration(typedAssignDeclaration TypedAssignDeclaration) *TypedAssignDeclarationDeclaration {
	return &TypedAssignDeclarationDeclaration{typedAssignDeclaration: typedAssignDeclaration}
}

func (tad *TypedAssignDeclarationDeclaration) String() string {
	return tad.typedAssignDeclaration.String()
}

type UntypedAssignDeclarationDeclaration struct {
	AssignDeclarationBase
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewUntypedAssignDeclarationDeclaration(untypedAssignDeclaration UntypedAssignDeclaration) *UntypedAssignDeclarationDeclaration {
	return &UntypedAssignDeclarationDeclaration{untypedAssignDeclaration: untypedAssignDeclaration}
}

func (uad *UntypedAssignDeclarationDeclaration) String() string {
	return uad.untypedAssignDeclaration.String()
}

type StructDestrutureAssignDeclaration struct {
	AssignDeclarationBase
	declarations []AliasableAssignDeclaration
}

func NewStructDestrutureAssignDeclaration(declarations []AliasableAssignDeclaration) *StructDestrutureAssignDeclaration {
	return &StructDestrutureAssignDeclaration{declarations: declarations}
}

func (sda *StructDestrutureAssignDeclaration) String() string {
	return SliceToString(sda.declarations, ", ")
}

type TypedAssignDeclaration interface {
	Node
	IsTypedAssignDeclaration()
}

type TypedAssignDeclarationBase struct {
	typeExpression           TypeExpression
	untypedAssignDeclaration UntypedAssignDeclaration
}

func NewTypedAssignDeclaration(typeExpression TypeExpression, untypedAssignDeclaration UntypedAssignDeclaration) *TypedAssignDeclarationBase {
	return &TypedAssignDeclarationBase{typeExpression: typeExpression, untypedAssignDeclaration: untypedAssignDeclaration}
}

func (tad *TypedAssignDeclarationBase) String() string {
	return fmt.Sprintf("%s %s", tad.typeExpression, tad.untypedAssignDeclaration)
}

func (tad *TypedAssignDeclarationBase) IsTypedAssignDeclaration() {}

type AliasableUntypedAssignDeclaration interface {
	UntypedAssignDeclaration
	Aliasable
}

type AliasedUntypedAssignDeclaration struct {
	AliasableBase

	untypedAssignDeclaration UntypedAssignDeclaration
	fromIdentifier           Identifier
}

func NewAliasedUntypedAssignDeclaration(declaration UntypedAssignDeclaration, fromIdentifier Identifier) *AliasedUntypedAssignDeclaration {
	return &AliasedUntypedAssignDeclaration{untypedAssignDeclaration: declaration, fromIdentifier: fromIdentifier}
}

func (auad *AliasedUntypedAssignDeclaration) String() string {
	return fmt.Sprintf("%s : %s", auad.untypedAssignDeclaration, auad.fromIdentifier)
}

type UnaliasedUntypedAssignDeclaration struct {
	AliasableBase
	UntypedAssignDeclaration
}

func NewUnaliasedUntypedAssignDeclaration(declaration UntypedAssignDeclaration) *UnaliasedUntypedAssignDeclaration {
	return &UnaliasedUntypedAssignDeclaration{UntypedAssignDeclaration: declaration}
}

func (uuad *UnaliasedUntypedAssignDeclaration) String() string {
	return uuad.String()
}

type UntypedAssignDeclaration interface {
	Node
	IsUntypedAssignDeclaration()
}

type UntypedAssignDeclarationBase struct{}

func (uad *UntypedAssignDeclarationBase) IsUntypedAssignDeclaration() {}

type UntypedIdentifierAssignDeclaration struct {
	UntypedAssignDeclarationBase
	identifier Identifier
}

func NewUntypedIdentifierAssignDeclaration(identifier Identifier) *UntypedIdentifierAssignDeclaration {
	return &UntypedIdentifierAssignDeclaration{identifier: identifier}
}

func (uiad *UntypedIdentifierAssignDeclaration) String() string {
	return uiad.identifier.String()
}

type UntypedStructDestrutureAssignDeclaration struct {
	UntypedAssignDeclarationBase
	declarations []AliasableUntypedAssignDeclaration
}

func NewUntypedStructDestrutureAssignDeclaration(declarations []AliasableUntypedAssignDeclaration) *UntypedStructDestrutureAssignDeclaration {
	return &UntypedStructDestrutureAssignDeclaration{declarations: declarations}
}

func (usdad *UntypedStructDestrutureAssignDeclaration) String() string {
	return fmt.Sprintf("{%s}", SliceToString(usdad.declarations, ", "))
}

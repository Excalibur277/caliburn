// Code generated from Caliburn.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // Caliburn
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CaliburnParser struct {
	*antlr.BaseParser
}

var CaliburnParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func caliburnParserInit() {
	staticData := &CaliburnParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'",
		"'&'", "'!'", "'=='", "'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT",
		"OP_OR", "OP_XOR", "OP_AND", "OP_NOT", "OP_EQU", "OP_NEQ", "OP_GTE",
		"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier",
		"Literal",
	}
	staticData.RuleNames = []string{
		"expression", "operator_expression", "expression_atom", "value_atom",
		"type_atom", "untyped_literal_atom", "typed_literal_atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 35, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 3, 0, 17, 8, 0, 1, 1, 1, 1, 1, 2,
		1, 2, 1, 2, 3, 2, 24, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 30, 0, 16, 1, 0, 0,
		0, 2, 18, 1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 25, 1, 0, 0, 0, 8, 27, 1, 0,
		0, 0, 10, 29, 1, 0, 0, 0, 12, 31, 1, 0, 0, 0, 14, 17, 3, 2, 1, 0, 15, 17,
		3, 4, 2, 0, 16, 14, 1, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17, 1, 1, 0, 0, 0,
		18, 19, 1, 0, 0, 0, 19, 3, 1, 0, 0, 0, 20, 24, 3, 6, 3, 0, 21, 24, 3, 10,
		5, 0, 22, 24, 3, 12, 6, 0, 23, 20, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23,
		22, 1, 0, 0, 0, 24, 5, 1, 0, 0, 0, 25, 26, 5, 20, 0, 0, 26, 7, 1, 0, 0,
		0, 27, 28, 5, 20, 0, 0, 28, 9, 1, 0, 0, 0, 29, 30, 5, 21, 0, 0, 30, 11,
		1, 0, 0, 0, 31, 32, 3, 8, 4, 0, 32, 33, 5, 21, 0, 0, 33, 13, 1, 0, 0, 0,
		2, 16, 23,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CaliburnParserInit initializes any static state used to implement CaliburnParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCaliburnParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CaliburnParserInit() {
	staticData := &CaliburnParserStaticData
	staticData.once.Do(caliburnParserInit)
}

// NewCaliburnParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCaliburnParser(input antlr.TokenStream) *CaliburnParser {
	CaliburnParserInit()
	this := new(CaliburnParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CaliburnParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Caliburn.g4"

	return this
}

// CaliburnParser tokens.
const (
	CaliburnParserEOF        = antlr.TokenEOF
	CaliburnParserOP_ADD     = 1
	CaliburnParserOP_SUB     = 2
	CaliburnParserOP_MUL     = 3
	CaliburnParserOP_DIV     = 4
	CaliburnParserOP_MOD     = 5
	CaliburnParserOP_POW     = 6
	CaliburnParserOP_ROOT    = 7
	CaliburnParserOP_OR      = 8
	CaliburnParserOP_XOR     = 9
	CaliburnParserOP_AND     = 10
	CaliburnParserOP_NOT     = 11
	CaliburnParserOP_EQU     = 12
	CaliburnParserOP_NEQ     = 13
	CaliburnParserOP_GTE     = 14
	CaliburnParserOP_GRT     = 15
	CaliburnParserOP_LTE     = 16
	CaliburnParserOP_LST     = 17
	CaliburnParserWhiteSpace = 18
	CaliburnParserComment    = 19
	CaliburnParserIdentifier = 20
	CaliburnParserLiteral    = 21
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_expression           = 0
	CaliburnParserRULE_operator_expression  = 1
	CaliburnParserRULE_expression_atom      = 2
	CaliburnParserRULE_value_atom           = 3
	CaliburnParserRULE_type_atom            = 4
	CaliburnParserRULE_untyped_literal_atom = 5
	CaliburnParserRULE_typed_literal_atom   = 6
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Operator_expression() IOperator_expressionContext
	Expression_atom() IExpression_atomContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Operator_expression() IOperator_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_expressionContext)
}

func (s *ExpressionContext) Expression_atom() IExpression_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_atomContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *CaliburnParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CaliburnParserRULE_expression)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserEOF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.Operator_expression()
		}

	case CaliburnParserIdentifier, CaliburnParserLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(15)
			p.Expression_atom()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperator_expressionContext is an interface to support dynamic dispatch.
type IOperator_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperator_expressionContext differentiates from other interfaces.
	IsOperator_expressionContext()
}

type Operator_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_expressionContext() *Operator_expressionContext {
	var p = new(Operator_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_operator_expression
	return p
}

func InitEmptyOperator_expressionContext(p *Operator_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_operator_expression
}

func (*Operator_expressionContext) IsOperator_expressionContext() {}

func NewOperator_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_expressionContext {
	var p = new(Operator_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_operator_expression

	return p
}

func (s *Operator_expressionContext) GetParser() antlr.Parser { return s.parser }
func (s *Operator_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operator_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterOperator_expression(s)
	}
}

func (s *Operator_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitOperator_expression(s)
	}
}

func (p *CaliburnParser) Operator_expression() (localctx IOperator_expressionContext) {
	localctx = NewOperator_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CaliburnParserRULE_operator_expression)
	p.EnterOuterAlt(localctx, 1)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpression_atomContext is an interface to support dynamic dispatch.
type IExpression_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value_atom() IValue_atomContext
	Untyped_literal_atom() IUntyped_literal_atomContext
	Typed_literal_atom() ITyped_literal_atomContext

	// IsExpression_atomContext differentiates from other interfaces.
	IsExpression_atomContext()
}

type Expression_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_atomContext() *Expression_atomContext {
	var p = new(Expression_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_atom
	return p
}

func InitEmptyExpression_atomContext(p *Expression_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_atom
}

func (*Expression_atomContext) IsExpression_atomContext() {}

func NewExpression_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_atomContext {
	var p = new(Expression_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expression_atom

	return p
}

func (s *Expression_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_atomContext) Value_atom() IValue_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_atomContext)
}

func (s *Expression_atomContext) Untyped_literal_atom() IUntyped_literal_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_literal_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_literal_atomContext)
}

func (s *Expression_atomContext) Typed_literal_atom() ITyped_literal_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITyped_literal_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITyped_literal_atomContext)
}

func (s *Expression_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterExpression_atom(s)
	}
}

func (s *Expression_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitExpression_atom(s)
	}
}

func (p *CaliburnParser) Expression_atom() (localctx IExpression_atomContext) {
	localctx = NewExpression_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CaliburnParserRULE_expression_atom)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Value_atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.Untyped_literal_atom()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Typed_literal_atom()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValue_atomContext is an interface to support dynamic dispatch.
type IValue_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsValue_atomContext differentiates from other interfaces.
	IsValue_atomContext()
}

type Value_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_atomContext() *Value_atomContext {
	var p = new(Value_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atom
	return p
}

func InitEmptyValue_atomContext(p *Value_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atom
}

func (*Value_atomContext) IsValue_atomContext() {}

func NewValue_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_atomContext {
	var p = new(Value_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_value_atom

	return p
}

func (s *Value_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_atomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Value_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterValue_atom(s)
	}
}

func (s *Value_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitValue_atom(s)
	}
}

func (p *CaliburnParser) Value_atom() (localctx IValue_atomContext) {
	localctx = NewValue_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CaliburnParserRULE_value_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IType_atomContext is an interface to support dynamic dispatch.
type IType_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsType_atomContext differentiates from other interfaces.
	IsType_atomContext()
}

type Type_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_atomContext() *Type_atomContext {
	var p = new(Type_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_atom
	return p
}

func InitEmptyType_atomContext(p *Type_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_atom
}

func (*Type_atomContext) IsType_atomContext() {}

func NewType_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_atomContext {
	var p = new(Type_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_type_atom

	return p
}

func (s *Type_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_atomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Type_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterType_atom(s)
	}
}

func (s *Type_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitType_atom(s)
	}
}

func (p *CaliburnParser) Type_atom() (localctx IType_atomContext) {
	localctx = NewType_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CaliburnParserRULE_type_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUntyped_literal_atomContext is an interface to support dynamic dispatch.
type IUntyped_literal_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() antlr.TerminalNode

	// IsUntyped_literal_atomContext differentiates from other interfaces.
	IsUntyped_literal_atomContext()
}

type Untyped_literal_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntyped_literal_atomContext() *Untyped_literal_atomContext {
	var p = new(Untyped_literal_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_literal_atom
	return p
}

func InitEmptyUntyped_literal_atomContext(p *Untyped_literal_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_literal_atom
}

func (*Untyped_literal_atomContext) IsUntyped_literal_atomContext() {}

func NewUntyped_literal_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Untyped_literal_atomContext {
	var p = new(Untyped_literal_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_untyped_literal_atom

	return p
}

func (s *Untyped_literal_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Untyped_literal_atomContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
}

func (s *Untyped_literal_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Untyped_literal_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Untyped_literal_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterUntyped_literal_atom(s)
	}
}

func (s *Untyped_literal_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitUntyped_literal_atom(s)
	}
}

func (p *CaliburnParser) Untyped_literal_atom() (localctx IUntyped_literal_atomContext) {
	localctx = NewUntyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CaliburnParserRULE_untyped_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)
		p.Match(CaliburnParserLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITyped_literal_atomContext is an interface to support dynamic dispatch.
type ITyped_literal_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_atom() IType_atomContext
	Literal() antlr.TerminalNode

	// IsTyped_literal_atomContext differentiates from other interfaces.
	IsTyped_literal_atomContext()
}

type Typed_literal_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyped_literal_atomContext() *Typed_literal_atomContext {
	var p = new(Typed_literal_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_literal_atom
	return p
}

func InitEmptyTyped_literal_atomContext(p *Typed_literal_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_literal_atom
}

func (*Typed_literal_atomContext) IsTyped_literal_atomContext() {}

func NewTyped_literal_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typed_literal_atomContext {
	var p = new(Typed_literal_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_typed_literal_atom

	return p
}

func (s *Typed_literal_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Typed_literal_atomContext) Type_atom() IType_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_atomContext)
}

func (s *Typed_literal_atomContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
}

func (s *Typed_literal_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typed_literal_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typed_literal_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.EnterTyped_literal_atom(s)
	}
}

func (s *Typed_literal_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnListener); ok {
		listenerT.ExitTyped_literal_atom(s)
	}
}

func (p *CaliburnParser) Typed_literal_atom() (localctx ITyped_literal_atomContext) {
	localctx = NewTyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CaliburnParserRULE_typed_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Type_atom()
	}
	{
		p.SetState(32)
		p.Match(CaliburnParserLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// Generated from /home/excalibur/PersonalProjects/caliburncode/parser/CaliburnParser.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class CaliburnParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PERIOD=1, COMMA=2, COLON=3, QUESTION=4, Terminator=5, ASSIGN=6, ARROW=7, 
		IF=8, ELSE=9, FOR=10, SWITCH=11, CASE=12, DEFAULT=13, FALLTHROUGH=14, 
		BREAK=15, CONTINUE=16, RETURN=17, USING=18, AS=19, IN=20, NULL=21, VAR=22, 
		CONST=23, TYPE=24, FUNC=25, STRUCT=26, TUPLE=27, CLASS=28, INTERFACE=29, 
		L_PAREN=30, R_PAREN=31, L_S_BRACK=32, R_S_BRACK=33, L_C_BRACK=34, R_C_BRACK=35, 
		OP_ADD=36, OP_SUB=37, OP_NOT=38, OP_MUL=39, OP_DIV=40, OP_MOD=41, OP_POW=42, 
		OP_ROOT=43, OP_OR=44, OP_XOR=45, OP_AND=46, OP_LSHIFT=47, OP_RSHIFT=48, 
		OP_EQU=49, OP_NEQ=50, OP_GTE=51, OP_GRT=52, OP_LTE=53, OP_LST=54, WhiteSpace=55, 
		Comment=56, Identifier=57, Literal=58;
	public static final int
		RULE_module = 0, RULE_definition = 1, RULE_function_definition = 2, RULE_parameters = 3, 
		RULE_parameter = 4, RULE_block = 5, RULE_statement = 6, RULE_jump_statement = 7, 
		RULE_return_statement = 8, RULE_break_statement = 9, RULE_continue_statement = 10, 
		RULE_control_statement = 11, RULE_if_statement = 12, RULE_for_statement = 13, 
		RULE_switch_statement = 14, RULE_case_statements = 15, RULE_option_case_statement = 16, 
		RULE_default_case_statement = 17, RULE_inline_statement = 18, RULE_assign_statement = 19, 
		RULE_assign_expressions = 20, RULE_assign_expression = 21, RULE_assign_declarations = 22, 
		RULE_assign_declaration = 23, RULE_typed_assign_declarations = 24, RULE_typed_assign_declaration = 25, 
		RULE_untyped_assign_declarations = 26, RULE_untyped_assign_declaration = 27, 
		RULE_expression_statement = 28, RULE_expressions = 29, RULE_expression = 30, 
		RULE_function_expression = 31, RULE_struct_expression = 32, RULE_tuple_expression = 33, 
		RULE_function_type = 34, RULE_struct_type = 35, RULE_tuple_type = 36, 
		RULE_type_expression = 37, RULE_identifiers = 38, RULE_literal_atom = 39, 
		RULE_untyped_literal_atom = 40, RULE_typed_literal_atom = 41;
	private static String[] makeRuleNames() {
		return new String[] {
			"module", "definition", "function_definition", "parameters", "parameter", 
			"block", "statement", "jump_statement", "return_statement", "break_statement", 
			"continue_statement", "control_statement", "if_statement", "for_statement", 
			"switch_statement", "case_statements", "option_case_statement", "default_case_statement", 
			"inline_statement", "assign_statement", "assign_expressions", "assign_expression", 
			"assign_declarations", "assign_declaration", "typed_assign_declarations", 
			"typed_assign_declaration", "untyped_assign_declarations", "untyped_assign_declaration", 
			"expression_statement", "expressions", "expression", "function_expression", 
			"struct_expression", "tuple_expression", "function_type", "struct_type", 
			"tuple_type", "type_expression", "identifiers", "literal_atom", "untyped_literal_atom", 
			"typed_literal_atom"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "','", "':'", "'?'", "';'", "'='", "'->'", "'if'", "'else'", 
			"'for'", "'switch'", "'case'", "'default'", "'fallthrough'", "'break'", 
			"'continue'", "'return'", "'using'", "'as'", "'in'", "'null'", "'var'", 
			"'const'", "'type'", "'func'", "'struct'", "'tuple'", "'class'", "'interface'", 
			"'('", "')'", "'['", "']'", "'{'", "'}'", "'+'", "'-'", "'!'", "'*'", 
			"'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", "'&'", "'<<'", "'>>'", "'=='", 
			"'!='", "'>='", "'>'", "'<='", "'<'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PERIOD", "COMMA", "COLON", "QUESTION", "Terminator", "ASSIGN", 
			"ARROW", "IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", 
			"BREAK", "CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", 
			"TYPE", "FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN", 
			"L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD", "OP_SUB", 
			"OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", "OP_OR", 
			"OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ", "OP_GTE", 
			"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier", 
			"Literal"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "CaliburnParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CaliburnParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleContext extends ParserRuleContext {
		public List<DefinitionContext> definition() {
			return getRuleContexts(DefinitionContext.class);
		}
		public DefinitionContext definition(int i) {
			return getRuleContext(DefinitionContext.class,i);
		}
		public ModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_module; }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_module);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(87);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC || _la==Identifier) {
				{
				{
				setState(84);
				definition();
				}
				}
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefinitionContext extends ParserRuleContext {
		public Function_definitionContext function_definition() {
			return getRuleContext(Function_definitionContext.class,0);
		}
		public DefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definition; }
	}

	public final DefinitionContext definition() throws RecognitionException {
		DefinitionContext _localctx = new DefinitionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_definition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			function_definition();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_definitionContext extends ParserRuleContext {
		public Function_typeContext type;
		public Token name;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public Function_definitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_definition; }
	}

	public final Function_definitionContext function_definition() throws RecognitionException {
		Function_definitionContext _localctx = new Function_definitionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_function_definition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			((Function_definitionContext)_localctx).type = function_type();
			setState(93);
			((Function_definitionContext)_localctx).name = match(Identifier);
			setState(94);
			match(L_PAREN);
			setState(96);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144115209550692352L) != 0)) {
				{
				setState(95);
				parameters();
				}
			}

			setState(98);
			match(R_PAREN);
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(99);
				((Function_definitionContext)_localctx).return_type = type_expression(0);
				}
			}

			setState(102);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametersContext extends ParserRuleContext {
		public List<ParameterContext> parameter() {
			return getRuleContexts(ParameterContext.class);
		}
		public ParameterContext parameter(int i) {
			return getRuleContext(ParameterContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(104);
			parameter();
			setState(109);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(105);
				match(COMMA);
				setState(106);
				parameter();
				}
				}
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterContext extends ParserRuleContext {
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	 
		public ParameterContext() { }
		public void copyFrom(ParameterContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDestrutureParameterContext extends ParameterContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public StructDestrutureParameterContext(ParameterContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleDestrutureParameterContext extends ParameterContext {
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public TupleDestrutureParameterContext(ParameterContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedParameterContext extends ParameterContext {
		public Typed_assign_declarationContext typed_assign_declaration() {
			return getRuleContext(Typed_assign_declarationContext.class,0);
		}
		public TypedParameterContext(ParameterContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedParameterContext extends ParameterContext {
		public Untyped_assign_declarationContext untyped_assign_declaration() {
			return getRuleContext(Untyped_assign_declarationContext.class,0);
		}
		public UntypedParameterContext(ParameterContext ctx) { copyFrom(ctx); }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_parameter);
		try {
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new TypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				typed_assign_declaration();
				}
				break;
			case 2:
				_localctx = new UntypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(113);
				untyped_assign_declaration();
				}
				break;
			case 3:
				_localctx = new StructDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(114);
				match(L_C_BRACK);
				setState(115);
				parameters();
				setState(116);
				match(R_C_BRACK);
				}
				break;
			case 4:
				_localctx = new TupleDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(118);
				match(L_S_BRACK);
				setState(119);
				parameters();
				setState(120);
				match(R_S_BRACK);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			match(L_C_BRACK);
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346067917573376L) != 0)) {
				{
				{
				setState(125);
				statement();
				}
				}
				setState(130);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(131);
			match(R_C_BRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StatementContext extends ParserRuleContext {
		public Assign_statementContext assign_statement() {
			return getRuleContext(Assign_statementContext.class,0);
		}
		public Expression_statementContext expression_statement() {
			return getRuleContext(Expression_statementContext.class,0);
		}
		public Control_statementContext control_statement() {
			return getRuleContext(Control_statementContext.class,0);
		}
		public Jump_statementContext jump_statement() {
			return getRuleContext(Jump_statementContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_statement);
		try {
			setState(137);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(133);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(134);
				expression_statement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(135);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(136);
				jump_statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Jump_statementContext extends ParserRuleContext {
		public Return_statementContext return_statement() {
			return getRuleContext(Return_statementContext.class,0);
		}
		public Break_statementContext break_statement() {
			return getRuleContext(Break_statementContext.class,0);
		}
		public Continue_statementContext continue_statement() {
			return getRuleContext(Continue_statementContext.class,0);
		}
		public Jump_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_jump_statement; }
	}

	public final Jump_statementContext jump_statement() throws RecognitionException {
		Jump_statementContext _localctx = new Jump_statementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_jump_statement);
		try {
			setState(142);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN:
				enterOuterAlt(_localctx, 1);
				{
				setState(139);
				return_statement();
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 2);
				{
				setState(140);
				break_statement();
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(141);
				continue_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Return_statementContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(CaliburnParser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_return_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(RETURN);
			setState(145);
			expression(0);
			setState(146);
			match(Terminator);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Break_statementContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(CaliburnParser.BREAK, 0); }
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Break_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_statement; }
	}

	public final Break_statementContext break_statement() throws RecognitionException {
		Break_statementContext _localctx = new Break_statementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(BREAK);
			setState(149);
			match(Terminator);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Continue_statementContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(CaliburnParser.CONTINUE, 0); }
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Continue_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_statement; }
	}

	public final Continue_statementContext continue_statement() throws RecognitionException {
		Continue_statementContext _localctx = new Continue_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(151);
			match(CONTINUE);
			setState(152);
			match(Terminator);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Control_statementContext extends ParserRuleContext {
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public For_statementContext for_statement() {
			return getRuleContext(For_statementContext.class,0);
		}
		public Switch_statementContext switch_statement() {
			return getRuleContext(Switch_statementContext.class,0);
		}
		public Control_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_statement; }
	}

	public final Control_statementContext control_statement() throws RecognitionException {
		Control_statementContext _localctx = new Control_statementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_control_statement);
		try {
			setState(157);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(154);
				if_statement();
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(155);
				for_statement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 3);
				{
				setState(156);
				switch_statement();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_statementContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(CaliburnParser.IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public List<Inline_statementContext> inline_statement() {
			return getRuleContexts(Inline_statementContext.class);
		}
		public Inline_statementContext inline_statement(int i) {
			return getRuleContext(Inline_statementContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(CaliburnParser.ELSE, 0); }
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_if_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(IF);
			setState(163);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(160);
					inline_statement();
					}
					} 
				}
				setState(165);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(166);
			expression(0);
			setState(167);
			block();
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(168);
				match(ELSE);
				setState(171);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case L_C_BRACK:
					{
					setState(169);
					block();
					}
					break;
				case IF:
					{
					setState(170);
					if_statement();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_statementContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(CaliburnParser.FOR, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public List<Inline_statementContext> inline_statement() {
			return getRuleContexts(Inline_statementContext.class);
		}
		public Inline_statementContext inline_statement(int i) {
			return getRuleContext(Inline_statementContext.class,i);
		}
		public TerminalNode ARROW() { return getToken(CaliburnParser.ARROW, 0); }
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_for_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(FOR);
			setState(177); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(176);
					inline_statement();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(179); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(181);
			expression(0);
			setState(188);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(182);
				match(ARROW);
				setState(184); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(183);
						inline_statement();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(186); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
			}

			setState(190);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_statementContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(CaliburnParser.SWITCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Case_statementsContext case_statements() {
			return getRuleContext(Case_statementsContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<Inline_statementContext> inline_statement() {
			return getRuleContexts(Inline_statementContext.class);
		}
		public Inline_statementContext inline_statement(int i) {
			return getRuleContext(Inline_statementContext.class,i);
		}
		public Switch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_statement; }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_switch_statement);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(SWITCH);
			setState(196);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(193);
					inline_statement();
					}
					} 
				}
				setState(198);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			setState(199);
			expression(0);
			setState(200);
			match(L_C_BRACK);
			setState(201);
			case_statements();
			setState(202);
			match(R_C_BRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Case_statementsContext extends ParserRuleContext {
		public List<Option_case_statementContext> option_case_statement() {
			return getRuleContexts(Option_case_statementContext.class);
		}
		public Option_case_statementContext option_case_statement(int i) {
			return getRuleContext(Option_case_statementContext.class,i);
		}
		public Default_case_statementContext default_case_statement() {
			return getRuleContext(Default_case_statementContext.class,0);
		}
		public Case_statementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case_statements; }
	}

	public final Case_statementsContext case_statements() throws RecognitionException {
		Case_statementsContext _localctx = new Case_statementsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_case_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(204);
				option_case_statement();
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(211);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(210);
				default_case_statement();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Option_case_statementContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(CaliburnParser.CASE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Option_case_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_option_case_statement; }
	}

	public final Option_case_statementContext option_case_statement() throws RecognitionException {
		Option_case_statementContext _localctx = new Option_case_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_option_case_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(213);
			match(CASE);
			setState(214);
			expression(0);
			setState(215);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Default_case_statementContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(CaliburnParser.DEFAULT, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Default_case_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case_statement; }
	}

	public final Default_case_statementContext default_case_statement() throws RecognitionException {
		Default_case_statementContext _localctx = new Default_case_statementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_default_case_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			match(DEFAULT);
			setState(218);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Inline_statementContext extends ParserRuleContext {
		public Assign_statementContext assign_statement() {
			return getRuleContext(Assign_statementContext.class,0);
		}
		public Expression_statementContext expression_statement() {
			return getRuleContext(Expression_statementContext.class,0);
		}
		public Inline_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inline_statement; }
	}

	public final Inline_statementContext inline_statement() throws RecognitionException {
		Inline_statementContext _localctx = new Inline_statementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_inline_statement);
		try {
			setState(222);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(220);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(221);
				expression_statement();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_statementContext extends ParserRuleContext {
		public Token op;
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(CaliburnParser.ASSIGN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Assign_expressionsContext assign_expressions() {
			return getRuleContext(Assign_expressionsContext.class,0);
		}
		public TerminalNode OP_ADD() { return getToken(CaliburnParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CaliburnParser.OP_SUB, 0); }
		public TerminalNode OP_MUL() { return getToken(CaliburnParser.OP_MUL, 0); }
		public TerminalNode OP_DIV() { return getToken(CaliburnParser.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(CaliburnParser.OP_MOD, 0); }
		public TerminalNode OP_POW() { return getToken(CaliburnParser.OP_POW, 0); }
		public TerminalNode OP_ROOT() { return getToken(CaliburnParser.OP_ROOT, 0); }
		public TerminalNode OP_OR() { return getToken(CaliburnParser.OP_OR, 0); }
		public TerminalNode OP_XOR() { return getToken(CaliburnParser.OP_XOR, 0); }
		public TerminalNode OP_AND() { return getToken(CaliburnParser.OP_AND, 0); }
		public TerminalNode OP_LSHIFT() { return getToken(CaliburnParser.OP_LSHIFT, 0); }
		public TerminalNode OP_RSHIFT() { return getToken(CaliburnParser.OP_RSHIFT, 0); }
		public Assign_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_statement; }
	}

	public final Assign_statementContext assign_statement() throws RecognitionException {
		Assign_statementContext _localctx = new Assign_statementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_assign_statement);
		int _la;
		try {
			setState(235);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(224);
				assign_declarations();
				setState(225);
				match(ASSIGN);
				setState(226);
				expressions();
				setState(227);
				match(Terminator);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(229);
				assign_expressions();
				setState(230);
				((Assign_statementContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 562606356037632L) != 0)) ) {
					((Assign_statementContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(231);
				match(ASSIGN);
				setState(232);
				expressions();
				setState(233);
				match(Terminator);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_expressionsContext extends ParserRuleContext {
		public List<Assign_expressionContext> assign_expression() {
			return getRuleContexts(Assign_expressionContext.class);
		}
		public Assign_expressionContext assign_expression(int i) {
			return getRuleContext(Assign_expressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Assign_expressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_expressions; }
	}

	public final Assign_expressionsContext assign_expressions() throws RecognitionException {
		Assign_expressionsContext _localctx = new Assign_expressionsContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assign_expressions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(237);
			assign_expression();
			setState(242);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(238);
				match(COMMA);
				setState(239);
				assign_expression();
				}
				}
				setState(244);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_expressionContext extends ParserRuleContext {
		public Assign_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_expression; }
	 
		public Assign_expressionContext() { }
		public void copyFrom(Assign_expressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionAssignExpressionContext extends Assign_expressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionAssignExpressionContext(Assign_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleDestrutureAssignExpressionContext extends Assign_expressionContext {
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public Assign_expressionsContext assign_expressions() {
			return getRuleContext(Assign_expressionsContext.class,0);
		}
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public TupleDestrutureAssignExpressionContext(Assign_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDestrutureAssignExpressionContext extends Assign_expressionContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Assign_expressionsContext assign_expressions() {
			return getRuleContext(Assign_expressionsContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public StructDestrutureAssignExpressionContext(Assign_expressionContext ctx) { copyFrom(ctx); }
	}

	public final Assign_expressionContext assign_expression() throws RecognitionException {
		Assign_expressionContext _localctx = new Assign_expressionContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_assign_expression);
		try {
			setState(254);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
			case STRUCT:
			case L_PAREN:
			case OP_ADD:
			case OP_SUB:
			case OP_NOT:
			case Identifier:
			case Literal:
				_localctx = new ExpressionAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(245);
				expression(0);
				}
				break;
			case L_C_BRACK:
				_localctx = new StructDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(246);
				match(L_C_BRACK);
				setState(247);
				assign_expressions();
				setState(248);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new TupleDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(250);
				match(L_S_BRACK);
				setState(251);
				assign_expressions();
				setState(252);
				match(R_S_BRACK);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_declarationsContext extends ParserRuleContext {
		public Assign_declarationContext assign_declaration() {
			return getRuleContext(Assign_declarationContext.class,0);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public List<Assign_statementContext> assign_statement() {
			return getRuleContexts(Assign_statementContext.class);
		}
		public Assign_statementContext assign_statement(int i) {
			return getRuleContext(Assign_statementContext.class,i);
		}
		public Assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_declarations; }
	}

	public final Assign_declarationsContext assign_declarations() throws RecognitionException {
		Assign_declarationsContext _localctx = new Assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(256);
			assign_declaration();
			setState(261);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(257);
				match(COMMA);
				setState(258);
				assign_statement();
				}
				}
				setState(263);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_declarationContext extends ParserRuleContext {
		public Assign_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_declaration; }
	 
		public Assign_declarationContext() { }
		public void copyFrom(Assign_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedAssignDeclarationContext extends Assign_declarationContext {
		public TerminalNode VAR() { return getToken(CaliburnParser.VAR, 0); }
		public Untyped_assign_declarationContext untyped_assign_declaration() {
			return getRuleContext(Untyped_assign_declarationContext.class,0);
		}
		public UntypedAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionAssignDeclarationContext extends Assign_declarationContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDestrutureAssignDeclarationContext extends Assign_declarationContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public StructDestrutureAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleDestrutureAssignDeclarationContext extends Assign_declarationContext {
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public TupleDestrutureAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedAssignDeclarationContext extends Assign_declarationContext {
		public Typed_assign_declarationContext typed_assign_declaration() {
			return getRuleContext(Typed_assign_declarationContext.class,0);
		}
		public TypedAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Assign_declarationContext assign_declaration() throws RecognitionException {
		Assign_declarationContext _localctx = new Assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_assign_declaration);
		try {
			setState(276);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new ExpressionAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(264);
				expression(0);
				}
				break;
			case 2:
				_localctx = new TypedAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(265);
				typed_assign_declaration();
				}
				break;
			case 3:
				_localctx = new UntypedAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(266);
				match(VAR);
				setState(267);
				untyped_assign_declaration();
				}
				break;
			case 4:
				_localctx = new StructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(268);
				match(L_C_BRACK);
				setState(269);
				assign_declarations();
				setState(270);
				match(R_C_BRACK);
				}
				break;
			case 5:
				_localctx = new TupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(272);
				match(L_S_BRACK);
				setState(273);
				assign_declarations();
				setState(274);
				match(R_S_BRACK);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Typed_assign_declarationsContext extends ParserRuleContext {
		public List<Typed_assign_declarationContext> typed_assign_declaration() {
			return getRuleContexts(Typed_assign_declarationContext.class);
		}
		public Typed_assign_declarationContext typed_assign_declaration(int i) {
			return getRuleContext(Typed_assign_declarationContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Typed_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typed_assign_declarations; }
	}

	public final Typed_assign_declarationsContext typed_assign_declarations() throws RecognitionException {
		Typed_assign_declarationsContext _localctx = new Typed_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_typed_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(278);
			typed_assign_declaration();
			setState(283);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(279);
				match(COMMA);
				setState(280);
				typed_assign_declaration();
				}
				}
				setState(285);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Typed_assign_declarationContext extends ParserRuleContext {
		public Typed_assign_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typed_assign_declaration; }
	 
		public Typed_assign_declarationContext() { }
		public void copyFrom(Typed_assign_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedTypedAssignDeclarationContext extends Typed_assign_declarationContext {
		public Type_expressionContext type;
		public Untyped_assign_declarationContext untyped_assign_declaration() {
			return getRuleContext(Untyped_assign_declarationContext.class,0);
		}
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TypedTypedAssignDeclarationContext(Typed_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Typed_assign_declarationContext typed_assign_declaration() throws RecognitionException {
		Typed_assign_declarationContext _localctx = new Typed_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_typed_assign_declaration);
		try {
			_localctx = new TypedTypedAssignDeclarationContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			((TypedTypedAssignDeclarationContext)_localctx).type = type_expression(0);
			setState(287);
			untyped_assign_declaration();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Untyped_assign_declarationsContext extends ParserRuleContext {
		public List<Untyped_assign_declarationContext> untyped_assign_declaration() {
			return getRuleContexts(Untyped_assign_declarationContext.class);
		}
		public Untyped_assign_declarationContext untyped_assign_declaration(int i) {
			return getRuleContext(Untyped_assign_declarationContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Untyped_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untyped_assign_declarations; }
	}

	public final Untyped_assign_declarationsContext untyped_assign_declarations() throws RecognitionException {
		Untyped_assign_declarationsContext _localctx = new Untyped_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_untyped_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			untyped_assign_declaration();
			setState(294);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(290);
				match(COMMA);
				setState(291);
				untyped_assign_declaration();
				}
				}
				setState(296);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Untyped_assign_declarationContext extends ParserRuleContext {
		public Untyped_assign_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untyped_assign_declaration; }
	 
		public Untyped_assign_declarationContext() { }
		public void copyFrom(Untyped_assign_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedAtomAssignDeclarationContext extends Untyped_assign_declarationContext {
		public Token var;
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public UntypedAtomAssignDeclarationContext(Untyped_assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedStructDestrutureAssignDeclarationContext extends Untyped_assign_declarationContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Untyped_assign_declarationsContext untyped_assign_declarations() {
			return getRuleContext(Untyped_assign_declarationsContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public UntypedStructDestrutureAssignDeclarationContext(Untyped_assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedTupleDestrutureAssignDeclarationContext extends Untyped_assign_declarationContext {
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public Untyped_assign_declarationsContext untyped_assign_declarations() {
			return getRuleContext(Untyped_assign_declarationsContext.class,0);
		}
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public UntypedTupleDestrutureAssignDeclarationContext(Untyped_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Untyped_assign_declarationContext untyped_assign_declaration() throws RecognitionException {
		Untyped_assign_declarationContext _localctx = new Untyped_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_untyped_assign_declaration);
		try {
			setState(306);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				_localctx = new UntypedAtomAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(297);
				((UntypedAtomAssignDeclarationContext)_localctx).var = match(Identifier);
				}
				break;
			case L_C_BRACK:
				_localctx = new UntypedStructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(298);
				match(L_C_BRACK);
				setState(299);
				untyped_assign_declarations();
				setState(300);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new UntypedTupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(302);
				match(L_S_BRACK);
				setState(303);
				untyped_assign_declarations();
				setState(304);
				match(R_S_BRACK);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Expression_statementContext extends ParserRuleContext {
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Expression_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_statement; }
	}

	public final Expression_statementContext expression_statement() throws RecognitionException {
		Expression_statementContext _localctx = new Expression_statementContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_expression_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			expressions();
			setState(309);
			match(Terminator);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionsContext extends ParserRuleContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public ExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressions; }
	}

	public final ExpressionsContext expressions() throws RecognitionException {
		ExpressionsContext _localctx = new ExpressionsContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_expressions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(311);
			expression(0);
			setState(316);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(312);
				match(COMMA);
				setState(313);
				expression(0);
				}
				}
				setState(318);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BracketedExpressionContext extends ExpressionContext {
		public ExpressionContext exp;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BracketedExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralExpressionContext extends ExpressionContext {
		public Literal_atomContext literal;
		public Literal_atomContext literal_atom() {
			return getRuleContext(Literal_atomContext.class,0);
		}
		public LiteralExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceExpressionContext extends ExpressionContext {
		public ExpressionContext exp;
		public ExpressionContext start_index;
		public ExpressionContext end_index;
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public TerminalNode COLON() { return getToken(CaliburnParser.COLON, 0); }
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public SliceExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleExpressionContext extends ExpressionContext {
		public Tuple_expressionContext exp;
		public Tuple_expressionContext tuple_expression() {
			return getRuleContext(Tuple_expressionContext.class,0);
		}
		public TupleExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanBinaryExpressionContext extends ExpressionContext {
		public ExpressionContext lhs_exp;
		public Token op;
		public ExpressionContext rhs_exp;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_EQU() { return getToken(CaliburnParser.OP_EQU, 0); }
		public TerminalNode OP_NEQ() { return getToken(CaliburnParser.OP_NEQ, 0); }
		public TerminalNode OP_LTE() { return getToken(CaliburnParser.OP_LTE, 0); }
		public TerminalNode OP_GTE() { return getToken(CaliburnParser.OP_GTE, 0); }
		public TerminalNode OP_LST() { return getToken(CaliburnParser.OP_LST, 0); }
		public TerminalNode OP_GRT() { return getToken(CaliburnParser.OP_GRT, 0); }
		public BooleanBinaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexExpressionContext extends ExpressionContext {
		public ExpressionContext exp;
		public ExpressionContext index;
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public IndexExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnaryExpressionContext extends ExpressionContext {
		public Token op;
		public ExpressionContext exp;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OP_NOT() { return getToken(CaliburnParser.OP_NOT, 0); }
		public TerminalNode OP_ADD() { return getToken(CaliburnParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CaliburnParser.OP_SUB, 0); }
		public UnaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AccessExpressionContext extends ExpressionContext {
		public ExpressionContext exp;
		public Token identifier;
		public TerminalNode PERIOD() { return getToken(CaliburnParser.PERIOD, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public AccessExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierExpressionContext extends ExpressionContext {
		public Token identifier;
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public IdentifierExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionExpressionContext extends ExpressionContext {
		public Function_expressionContext exp;
		public Function_expressionContext function_expression() {
			return getRuleContext(Function_expressionContext.class,0);
		}
		public FunctionExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructExpressionContext extends ExpressionContext {
		public Struct_expressionContext exp;
		public Struct_expressionContext struct_expression() {
			return getRuleContext(Struct_expressionContext.class,0);
		}
		public StructExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BinaryExpressionContext extends ExpressionContext {
		public ExpressionContext lhs_exp;
		public Token op;
		public ExpressionContext rhs_exp;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_POW() { return getToken(CaliburnParser.OP_POW, 0); }
		public TerminalNode OP_ROOT() { return getToken(CaliburnParser.OP_ROOT, 0); }
		public TerminalNode OP_MUL() { return getToken(CaliburnParser.OP_MUL, 0); }
		public TerminalNode OP_DIV() { return getToken(CaliburnParser.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(CaliburnParser.OP_MOD, 0); }
		public TerminalNode OP_ADD() { return getToken(CaliburnParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CaliburnParser.OP_SUB, 0); }
		public TerminalNode OP_LSHIFT() { return getToken(CaliburnParser.OP_LSHIFT, 0); }
		public TerminalNode OP_RSHIFT() { return getToken(CaliburnParser.OP_RSHIFT, 0); }
		public TerminalNode OP_AND() { return getToken(CaliburnParser.OP_AND, 0); }
		public TerminalNode OP_XOR() { return getToken(CaliburnParser.OP_XOR, 0); }
		public TerminalNode OP_OR() { return getToken(CaliburnParser.OP_OR, 0); }
		public BinaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CastExpressionContext extends ExpressionContext {
		public Type_expressionContext type;
		public ExpressionContext exp;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public CastExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallExpressionContext extends ExpressionContext {
		public ExpressionContext exp;
		public ExpressionsContext args;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public CallExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				{
				_localctx = new CastExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(320);
				match(L_PAREN);
				setState(321);
				((CastExpressionContext)_localctx).type = type_expression(0);
				setState(322);
				match(R_PAREN);
				setState(323);
				((CastExpressionContext)_localctx).exp = expression(20);
				}
				break;
			case 2:
				{
				_localctx = new BracketedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(325);
				match(L_PAREN);
				setState(326);
				((BracketedExpressionContext)_localctx).exp = expression(0);
				setState(327);
				match(R_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(329);
				((UnaryExpressionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 481036337152L) != 0)) ) {
					((UnaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(330);
				((UnaryExpressionContext)_localctx).exp = expression(17);
				}
				break;
			case 4:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(331);
				((IdentifierExpressionContext)_localctx).identifier = match(Identifier);
				}
				break;
			case 5:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(332);
				((LiteralExpressionContext)_localctx).literal = literal_atom();
				}
				break;
			case 6:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(333);
				((FunctionExpressionContext)_localctx).exp = function_expression();
				}
				break;
			case 7:
				{
				_localctx = new StructExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(334);
				((StructExpressionContext)_localctx).exp = struct_expression();
				}
				break;
			case 8:
				{
				_localctx = new TupleExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(335);
				((TupleExpressionContext)_localctx).exp = tuple_expression();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(388);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(386);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(338);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(339);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_POW || _la==OP_ROOT) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(340);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(19);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(341);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(342);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3848290697216L) != 0)) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(343);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(17);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(344);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(345);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_ADD || _la==OP_SUB) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(346);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(16);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(347);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(348);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_LSHIFT || _la==OP_RSHIFT) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(349);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(15);
						}
						break;
					case 5:
						{
						_localctx = new BooleanBinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BooleanBinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(350);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(351);
						((BooleanBinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 35465847065542656L) != 0)) ) {
							((BooleanBinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(352);
						((BooleanBinaryExpressionContext)_localctx).rhs_exp = expression(14);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(353);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(354);
						match(OP_AND);
						setState(355);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(13);
						}
						break;
					case 7:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(356);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(357);
						match(OP_XOR);
						setState(358);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(12);
						}
						break;
					case 8:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(359);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(360);
						match(OP_OR);
						setState(361);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(11);
						}
						break;
					case 9:
						{
						_localctx = new CallExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((CallExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(362);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(363);
						match(L_PAREN);
						setState(365);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(364);
							((CallExpressionContext)_localctx).args = expressions();
							}
						}

						setState(367);
						match(R_PAREN);
						}
						break;
					case 10:
						{
						_localctx = new AccessExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((AccessExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(368);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(369);
						match(PERIOD);
						setState(370);
						((AccessExpressionContext)_localctx).identifier = match(Identifier);
						}
						break;
					case 11:
						{
						_localctx = new IndexExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((IndexExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(371);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(372);
						match(L_S_BRACK);
						setState(373);
						((IndexExpressionContext)_localctx).index = expression(0);
						setState(374);
						match(R_S_BRACK);
						}
						break;
					case 12:
						{
						_localctx = new SliceExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((SliceExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(376);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(377);
						match(L_S_BRACK);
						setState(379);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(378);
							((SliceExpressionContext)_localctx).start_index = expression(0);
							}
						}

						setState(381);
						match(COLON);
						setState(383);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(382);
							((SliceExpressionContext)_localctx).end_index = expression(0);
							}
						}

						setState(385);
						match(R_S_BRACK);
						}
						break;
					}
					} 
				}
				setState(390);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,33,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_expressionContext extends ParserRuleContext {
		public Function_typeContext type;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public Function_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_expression; }
	}

	public final Function_expressionContext function_expression() throws RecognitionException {
		Function_expressionContext _localctx = new Function_expressionContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(391);
			((Function_expressionContext)_localctx).type = function_type();
			setState(392);
			match(L_PAREN);
			setState(394);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346067917340672L) != 0)) {
				{
				setState(393);
				assign_declarations();
				}
			}

			setState(396);
			match(R_PAREN);
			setState(398);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(397);
				((Function_expressionContext)_localctx).return_type = type_expression(0);
				}
			}

			setState(400);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_expressionContext extends ParserRuleContext {
		public Struct_typeContext type;
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public List<TerminalNode> COLON() { return getTokens(CaliburnParser.COLON); }
		public TerminalNode COLON(int i) {
			return getToken(CaliburnParser.COLON, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public Struct_typeContext struct_type() {
			return getRuleContext(Struct_typeContext.class,0);
		}
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Struct_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_expression; }
	}

	public final Struct_expressionContext struct_expression() throws RecognitionException {
		Struct_expressionContext _localctx = new Struct_expressionContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_struct_expression);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(402);
			((Struct_expressionContext)_localctx).type = struct_type();
			setState(403);
			match(L_C_BRACK);
			{
			setState(404);
			match(Identifier);
			}
			setState(405);
			match(COLON);
			setState(406);
			expression(0);
			setState(413);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(407);
					match(COMMA);
					{
					setState(408);
					match(Identifier);
					}
					setState(409);
					match(COLON);
					setState(410);
					expression(0);
					}
					} 
				}
				setState(415);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,36,_ctx);
			}
			setState(417);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(416);
				match(COMMA);
				}
			}

			setState(419);
			match(R_C_BRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Tuple_expressionContext extends ParserRuleContext {
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Tuple_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_expression; }
	}

	public final Tuple_expressionContext tuple_expression() throws RecognitionException {
		Tuple_expressionContext _localctx = new Tuple_expressionContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_tuple_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421);
			match(L_PAREN);
			setState(422);
			expression(0);
			setState(430);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				setState(423);
				match(COMMA);
				}
				break;
			case 2:
				{
				setState(426); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(424);
					match(COMMA);
					setState(425);
					expression(0);
					}
					}
					setState(428); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				}
				break;
			}
			setState(432);
			match(R_PAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Function_typeContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(CaliburnParser.FUNC, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public Function_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type; }
	}

	public final Function_typeContext function_type() throws RecognitionException {
		Function_typeContext _localctx = new Function_typeContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_function_type);
		try {
			setState(436);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				enterOuterAlt(_localctx, 1);
				{
				setState(434);
				match(FUNC);
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 2);
				{
				setState(435);
				type_expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_typeContext extends ParserRuleContext {
		public TerminalNode STRUCT() { return getToken(CaliburnParser.STRUCT, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public Struct_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_type; }
	}

	public final Struct_typeContext struct_type() throws RecognitionException {
		Struct_typeContext _localctx = new Struct_typeContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_struct_type);
		try {
			setState(440);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRUCT:
				enterOuterAlt(_localctx, 1);
				{
				setState(438);
				match(STRUCT);
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 2);
				{
				setState(439);
				type_expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Tuple_typeContext extends ParserRuleContext {
		public TerminalNode TUPLE() { return getToken(CaliburnParser.TUPLE, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public Tuple_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type; }
	}

	public final Tuple_typeContext tuple_type() throws RecognitionException {
		Tuple_typeContext _localctx = new Tuple_typeContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_tuple_type);
		try {
			setState(444);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TUPLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(442);
				match(TUPLE);
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 2);
				{
				setState(443);
				type_expression(0);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Type_expressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TerminalNode PERIOD() { return getToken(CaliburnParser.PERIOD, 0); }
		public Type_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_expression; }
	}

	public final Type_expressionContext type_expression() throws RecognitionException {
		return type_expression(0);
	}

	private Type_expressionContext type_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Type_expressionContext _localctx = new Type_expressionContext(_ctx, _parentState);
		Type_expressionContext _prevctx = _localctx;
		int _startState = 74;
		enterRecursionRule(_localctx, 74, RULE_type_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(447);
			match(Identifier);
			}
			_ctx.stop = _input.LT(-1);
			setState(454);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Type_expressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_type_expression);
					setState(449);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(450);
					match(PERIOD);
					setState(451);
					match(Identifier);
					}
					} 
				}
				setState(456);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IdentifiersContext extends ParserRuleContext {
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public IdentifiersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifiers; }
	}

	public final IdentifiersContext identifiers() throws RecognitionException {
		IdentifiersContext _localctx = new IdentifiersContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_identifiers);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(457);
			match(Identifier);
			setState(462);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(458);
				match(COMMA);
				setState(459);
				match(Identifier);
				}
				}
				setState(464);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Literal_atomContext extends ParserRuleContext {
		public Literal_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal_atom; }
	 
		public Literal_atomContext() { }
		public void copyFrom(Literal_atomContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedLiteralAtomContext extends Literal_atomContext {
		public Untyped_literal_atomContext untyped_literal_atom() {
			return getRuleContext(Untyped_literal_atomContext.class,0);
		}
		public UntypedLiteralAtomContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedLiteralAtomContext extends Literal_atomContext {
		public Typed_literal_atomContext typed_literal_atom() {
			return getRuleContext(Typed_literal_atomContext.class,0);
		}
		public TypedLiteralAtomContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}

	public final Literal_atomContext literal_atom() throws RecognitionException {
		Literal_atomContext _localctx = new Literal_atomContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_literal_atom);
		try {
			setState(467);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Literal:
				_localctx = new UntypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(465);
				untyped_literal_atom();
				}
				break;
			case Identifier:
				_localctx = new TypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(466);
				typed_literal_atom();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Untyped_literal_atomContext extends ParserRuleContext {
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public Untyped_literal_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untyped_literal_atom; }
	}

	public final Untyped_literal_atomContext untyped_literal_atom() throws RecognitionException {
		Untyped_literal_atomContext _localctx = new Untyped_literal_atomContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_untyped_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(469);
			match(Literal);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Typed_literal_atomContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public Typed_literal_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typed_literal_atom; }
	}

	public final Typed_literal_atomContext typed_literal_atom() throws RecognitionException {
		Typed_literal_atomContext _localctx = new Typed_literal_atomContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_typed_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(471);
			match(Identifier);
			setState(472);
			match(Literal);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 30:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 37:
			return type_expression_sempred((Type_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 18);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		case 5:
			return precpred(_ctx, 12);
		case 6:
			return precpred(_ctx, 11);
		case 7:
			return precpred(_ctx, 10);
		case 8:
			return precpred(_ctx, 9);
		case 9:
			return precpred(_ctx, 8);
		case 10:
			return precpred(_ctx, 7);
		case 11:
			return precpred(_ctx, 6);
		}
		return true;
	}
	private boolean type_expression_sempred(Type_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001:\u01db\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0001\u0000\u0005\u0000V\b\u0000\n\u0000\f\u0000"+
		"Y\t\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002a\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"e\b\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0005\u0003l\b\u0003\n\u0003\f\u0003o\t\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004{\b\u0004\u0001\u0005\u0001\u0005"+
		"\u0005\u0005\u007f\b\u0005\n\u0005\f\u0005\u0082\t\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u008a"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u008f\b\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u009e\b\u000b"+
		"\u0001\f\u0001\f\u0005\f\u00a2\b\f\n\f\f\f\u00a5\t\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0003\f\u00ac\b\f\u0003\f\u00ae\b\f\u0001\r\u0001\r"+
		"\u0004\r\u00b2\b\r\u000b\r\f\r\u00b3\u0001\r\u0001\r\u0001\r\u0004\r\u00b9"+
		"\b\r\u000b\r\f\r\u00ba\u0003\r\u00bd\b\r\u0001\r\u0001\r\u0001\u000e\u0001"+
		"\u000e\u0005\u000e\u00c3\b\u000e\n\u000e\f\u000e\u00c6\t\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0005\u000f"+
		"\u00ce\b\u000f\n\u000f\f\u000f\u00d1\t\u000f\u0001\u000f\u0003\u000f\u00d4"+
		"\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0003\u0012\u00df\b\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u00ec"+
		"\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00f1\b\u0014"+
		"\n\u0014\f\u0014\u00f4\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003"+
		"\u0015\u00ff\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u0104"+
		"\b\u0016\n\u0016\f\u0016\u0107\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u0115\b\u0017\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0005\u0018\u011a\b\u0018\n\u0018\f\u0018\u011d"+
		"\t\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0005\u001a\u0125\b\u001a\n\u001a\f\u001a\u0128\t\u001a\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0003\u001b\u0133\b\u001b\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u013b\b\u001d"+
		"\n\u001d\f\u001d\u013e\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0003\u001e\u0151\b\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0003\u001e\u016e\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u017c\b\u001e\u0001\u001e\u0001"+
		"\u001e\u0003\u001e\u0180\b\u001e\u0001\u001e\u0005\u001e\u0183\b\u001e"+
		"\n\u001e\f\u001e\u0186\t\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0003"+
		"\u001f\u018b\b\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u018f\b\u001f"+
		"\u0001\u001f\u0001\u001f\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0005 \u019c\b \n \f \u019f\t \u0001 \u0003 \u01a2\b "+
		"\u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001!\u0004!\u01ab\b!\u000b"+
		"!\f!\u01ac\u0003!\u01af\b!\u0001!\u0001!\u0001\"\u0001\"\u0003\"\u01b5"+
		"\b\"\u0001#\u0001#\u0003#\u01b9\b#\u0001$\u0001$\u0003$\u01bd\b$\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0005%\u01c5\b%\n%\f%\u01c8\t%\u0001"+
		"&\u0001&\u0001&\u0005&\u01cd\b&\n&\f&\u01d0\t&\u0001\'\u0001\'\u0003\'"+
		"\u01d4\b\'\u0001(\u0001(\u0001)\u0001)\u0001)\u0001)\u0000\u0002<J*\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.02468:<>@BDFHJLNPR\u0000\u0007\u0002\u0000$%\'0\u0001\u0000"+
		"$&\u0001\u0000*+\u0001\u0000\')\u0001\u0000$%\u0001\u0000/0\u0001\u0000"+
		"16\u01f9\u0000W\u0001\u0000\u0000\u0000\u0002Z\u0001\u0000\u0000\u0000"+
		"\u0004\\\u0001\u0000\u0000\u0000\u0006h\u0001\u0000\u0000\u0000\bz\u0001"+
		"\u0000\u0000\u0000\n|\u0001\u0000\u0000\u0000\f\u0089\u0001\u0000\u0000"+
		"\u0000\u000e\u008e\u0001\u0000\u0000\u0000\u0010\u0090\u0001\u0000\u0000"+
		"\u0000\u0012\u0094\u0001\u0000\u0000\u0000\u0014\u0097\u0001\u0000\u0000"+
		"\u0000\u0016\u009d\u0001\u0000\u0000\u0000\u0018\u009f\u0001\u0000\u0000"+
		"\u0000\u001a\u00af\u0001\u0000\u0000\u0000\u001c\u00c0\u0001\u0000\u0000"+
		"\u0000\u001e\u00cf\u0001\u0000\u0000\u0000 \u00d5\u0001\u0000\u0000\u0000"+
		"\"\u00d9\u0001\u0000\u0000\u0000$\u00de\u0001\u0000\u0000\u0000&\u00eb"+
		"\u0001\u0000\u0000\u0000(\u00ed\u0001\u0000\u0000\u0000*\u00fe\u0001\u0000"+
		"\u0000\u0000,\u0100\u0001\u0000\u0000\u0000.\u0114\u0001\u0000\u0000\u0000"+
		"0\u0116\u0001\u0000\u0000\u00002\u011e\u0001\u0000\u0000\u00004\u0121"+
		"\u0001\u0000\u0000\u00006\u0132\u0001\u0000\u0000\u00008\u0134\u0001\u0000"+
		"\u0000\u0000:\u0137\u0001\u0000\u0000\u0000<\u0150\u0001\u0000\u0000\u0000"+
		">\u0187\u0001\u0000\u0000\u0000@\u0192\u0001\u0000\u0000\u0000B\u01a5"+
		"\u0001\u0000\u0000\u0000D\u01b4\u0001\u0000\u0000\u0000F\u01b8\u0001\u0000"+
		"\u0000\u0000H\u01bc\u0001\u0000\u0000\u0000J\u01be\u0001\u0000\u0000\u0000"+
		"L\u01c9\u0001\u0000\u0000\u0000N\u01d3\u0001\u0000\u0000\u0000P\u01d5"+
		"\u0001\u0000\u0000\u0000R\u01d7\u0001\u0000\u0000\u0000TV\u0003\u0002"+
		"\u0001\u0000UT\u0001\u0000\u0000\u0000VY\u0001\u0000\u0000\u0000WU\u0001"+
		"\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000X\u0001\u0001\u0000\u0000"+
		"\u0000YW\u0001\u0000\u0000\u0000Z[\u0003\u0004\u0002\u0000[\u0003\u0001"+
		"\u0000\u0000\u0000\\]\u0003D\"\u0000]^\u00059\u0000\u0000^`\u0005\u001e"+
		"\u0000\u0000_a\u0003\u0006\u0003\u0000`_\u0001\u0000\u0000\u0000`a\u0001"+
		"\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000bd\u0005\u001f\u0000\u0000"+
		"ce\u0003J%\u0000dc\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000e"+
		"f\u0001\u0000\u0000\u0000fg\u0003\n\u0005\u0000g\u0005\u0001\u0000\u0000"+
		"\u0000hm\u0003\b\u0004\u0000ij\u0005\u0002\u0000\u0000jl\u0003\b\u0004"+
		"\u0000ki\u0001\u0000\u0000\u0000lo\u0001\u0000\u0000\u0000mk\u0001\u0000"+
		"\u0000\u0000mn\u0001\u0000\u0000\u0000n\u0007\u0001\u0000\u0000\u0000"+
		"om\u0001\u0000\u0000\u0000p{\u00032\u0019\u0000q{\u00036\u001b\u0000r"+
		"s\u0005\"\u0000\u0000st\u0003\u0006\u0003\u0000tu\u0005#\u0000\u0000u"+
		"{\u0001\u0000\u0000\u0000vw\u0005 \u0000\u0000wx\u0003\u0006\u0003\u0000"+
		"xy\u0005!\u0000\u0000y{\u0001\u0000\u0000\u0000zp\u0001\u0000\u0000\u0000"+
		"zq\u0001\u0000\u0000\u0000zr\u0001\u0000\u0000\u0000zv\u0001\u0000\u0000"+
		"\u0000{\t\u0001\u0000\u0000\u0000|\u0080\u0005\"\u0000\u0000}\u007f\u0003"+
		"\f\u0006\u0000~}\u0001\u0000\u0000\u0000\u007f\u0082\u0001\u0000\u0000"+
		"\u0000\u0080~\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000\u0000"+
		"\u0081\u0083\u0001\u0000\u0000\u0000\u0082\u0080\u0001\u0000\u0000\u0000"+
		"\u0083\u0084\u0005#\u0000\u0000\u0084\u000b\u0001\u0000\u0000\u0000\u0085"+
		"\u008a\u0003&\u0013\u0000\u0086\u008a\u00038\u001c\u0000\u0087\u008a\u0003"+
		"\u0016\u000b\u0000\u0088\u008a\u0003\u000e\u0007\u0000\u0089\u0085\u0001"+
		"\u0000\u0000\u0000\u0089\u0086\u0001\u0000\u0000\u0000\u0089\u0087\u0001"+
		"\u0000\u0000\u0000\u0089\u0088\u0001\u0000\u0000\u0000\u008a\r\u0001\u0000"+
		"\u0000\u0000\u008b\u008f\u0003\u0010\b\u0000\u008c\u008f\u0003\u0012\t"+
		"\u0000\u008d\u008f\u0003\u0014\n\u0000\u008e\u008b\u0001\u0000\u0000\u0000"+
		"\u008e\u008c\u0001\u0000\u0000\u0000\u008e\u008d\u0001\u0000\u0000\u0000"+
		"\u008f\u000f\u0001\u0000\u0000\u0000\u0090\u0091\u0005\u0011\u0000\u0000"+
		"\u0091\u0092\u0003<\u001e\u0000\u0092\u0093\u0005\u0005\u0000\u0000\u0093"+
		"\u0011\u0001\u0000\u0000\u0000\u0094\u0095\u0005\u000f\u0000\u0000\u0095"+
		"\u0096\u0005\u0005\u0000\u0000\u0096\u0013\u0001\u0000\u0000\u0000\u0097"+
		"\u0098\u0005\u0010\u0000\u0000\u0098\u0099\u0005\u0005\u0000\u0000\u0099"+
		"\u0015\u0001\u0000\u0000\u0000\u009a\u009e\u0003\u0018\f\u0000\u009b\u009e"+
		"\u0003\u001a\r\u0000\u009c\u009e\u0003\u001c\u000e\u0000\u009d\u009a\u0001"+
		"\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009d\u009c\u0001"+
		"\u0000\u0000\u0000\u009e\u0017\u0001\u0000\u0000\u0000\u009f\u00a3\u0005"+
		"\b\u0000\u0000\u00a0\u00a2\u0003$\u0012\u0000\u00a1\u00a0\u0001\u0000"+
		"\u0000\u0000\u00a2\u00a5\u0001\u0000\u0000\u0000\u00a3\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a4\u0001\u0000\u0000\u0000\u00a4\u00a6\u0001\u0000"+
		"\u0000\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a6\u00a7\u0003<\u001e"+
		"\u0000\u00a7\u00ad\u0003\n\u0005\u0000\u00a8\u00ab\u0005\t\u0000\u0000"+
		"\u00a9\u00ac\u0003\n\u0005\u0000\u00aa\u00ac\u0003\u0018\f\u0000\u00ab"+
		"\u00a9\u0001\u0000\u0000\u0000\u00ab\u00aa\u0001\u0000\u0000\u0000\u00ac"+
		"\u00ae\u0001\u0000\u0000\u0000\u00ad\u00a8\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ae\u0001\u0000\u0000\u0000\u00ae\u0019\u0001\u0000\u0000\u0000\u00af"+
		"\u00b1\u0005\n\u0000\u0000\u00b0\u00b2\u0003$\u0012\u0000\u00b1\u00b0"+
		"\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000\u0000\u00b3\u00b1"+
		"\u0001\u0000\u0000\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b5"+
		"\u0001\u0000\u0000\u0000\u00b5\u00bc\u0003<\u001e\u0000\u00b6\u00b8\u0005"+
		"\u0007\u0000\u0000\u00b7\u00b9\u0003$\u0012\u0000\u00b8\u00b7\u0001\u0000"+
		"\u0000\u0000\u00b9\u00ba\u0001\u0000\u0000\u0000\u00ba\u00b8\u0001\u0000"+
		"\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000\u00bb\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bc\u00b6\u0001\u0000\u0000\u0000\u00bc\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bd\u00be\u0001\u0000\u0000\u0000\u00be\u00bf\u0003\n\u0005"+
		"\u0000\u00bf\u001b\u0001\u0000\u0000\u0000\u00c0\u00c4\u0005\u000b\u0000"+
		"\u0000\u00c1\u00c3\u0003$\u0012\u0000\u00c2\u00c1\u0001\u0000\u0000\u0000"+
		"\u00c3\u00c6\u0001\u0000\u0000\u0000\u00c4\u00c2\u0001\u0000\u0000\u0000"+
		"\u00c4\u00c5\u0001\u0000\u0000\u0000\u00c5\u00c7\u0001\u0000\u0000\u0000"+
		"\u00c6\u00c4\u0001\u0000\u0000\u0000\u00c7\u00c8\u0003<\u001e\u0000\u00c8"+
		"\u00c9\u0005\"\u0000\u0000\u00c9\u00ca\u0003\u001e\u000f\u0000\u00ca\u00cb"+
		"\u0005#\u0000\u0000\u00cb\u001d\u0001\u0000\u0000\u0000\u00cc\u00ce\u0003"+
		" \u0010\u0000\u00cd\u00cc\u0001\u0000\u0000\u0000\u00ce\u00d1\u0001\u0000"+
		"\u0000\u0000\u00cf\u00cd\u0001\u0000\u0000\u0000\u00cf\u00d0\u0001\u0000"+
		"\u0000\u0000\u00d0\u00d3\u0001\u0000\u0000\u0000\u00d1\u00cf\u0001\u0000"+
		"\u0000\u0000\u00d2\u00d4\u0003\"\u0011\u0000\u00d3\u00d2\u0001\u0000\u0000"+
		"\u0000\u00d3\u00d4\u0001\u0000\u0000\u0000\u00d4\u001f\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u0005\f\u0000\u0000\u00d6\u00d7\u0003<\u001e\u0000"+
		"\u00d7\u00d8\u0003\n\u0005\u0000\u00d8!\u0001\u0000\u0000\u0000\u00d9"+
		"\u00da\u0005\r\u0000\u0000\u00da\u00db\u0003\n\u0005\u0000\u00db#\u0001"+
		"\u0000\u0000\u0000\u00dc\u00df\u0003&\u0013\u0000\u00dd\u00df\u00038\u001c"+
		"\u0000\u00de\u00dc\u0001\u0000\u0000\u0000\u00de\u00dd\u0001\u0000\u0000"+
		"\u0000\u00df%\u0001\u0000\u0000\u0000\u00e0\u00e1\u0003,\u0016\u0000\u00e1"+
		"\u00e2\u0005\u0006\u0000\u0000\u00e2\u00e3\u0003:\u001d\u0000\u00e3\u00e4"+
		"\u0005\u0005\u0000\u0000\u00e4\u00ec\u0001\u0000\u0000\u0000\u00e5\u00e6"+
		"\u0003(\u0014\u0000\u00e6\u00e7\u0007\u0000\u0000\u0000\u00e7\u00e8\u0005"+
		"\u0006\u0000\u0000\u00e8\u00e9\u0003:\u001d\u0000\u00e9\u00ea\u0005\u0005"+
		"\u0000\u0000\u00ea\u00ec\u0001\u0000\u0000\u0000\u00eb\u00e0\u0001\u0000"+
		"\u0000\u0000\u00eb\u00e5\u0001\u0000\u0000\u0000\u00ec\'\u0001\u0000\u0000"+
		"\u0000\u00ed\u00f2\u0003*\u0015\u0000\u00ee\u00ef\u0005\u0002\u0000\u0000"+
		"\u00ef\u00f1\u0003*\u0015\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1"+
		"\u00f4\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000\u0000\u00f2"+
		"\u00f3\u0001\u0000\u0000\u0000\u00f3)\u0001\u0000\u0000\u0000\u00f4\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f5\u00ff\u0003<\u001e\u0000\u00f6\u00f7\u0005"+
		"\"\u0000\u0000\u00f7\u00f8\u0003(\u0014\u0000\u00f8\u00f9\u0005#\u0000"+
		"\u0000\u00f9\u00ff\u0001\u0000\u0000\u0000\u00fa\u00fb\u0005 \u0000\u0000"+
		"\u00fb\u00fc\u0003(\u0014\u0000\u00fc\u00fd\u0005!\u0000\u0000\u00fd\u00ff"+
		"\u0001\u0000\u0000\u0000\u00fe\u00f5\u0001\u0000\u0000\u0000\u00fe\u00f6"+
		"\u0001\u0000\u0000\u0000\u00fe\u00fa\u0001\u0000\u0000\u0000\u00ff+\u0001"+
		"\u0000\u0000\u0000\u0100\u0105\u0003.\u0017\u0000\u0101\u0102\u0005\u0002"+
		"\u0000\u0000\u0102\u0104\u0003&\u0013\u0000\u0103\u0101\u0001\u0000\u0000"+
		"\u0000\u0104\u0107\u0001\u0000\u0000\u0000\u0105\u0103\u0001\u0000\u0000"+
		"\u0000\u0105\u0106\u0001\u0000\u0000\u0000\u0106-\u0001\u0000\u0000\u0000"+
		"\u0107\u0105\u0001\u0000\u0000\u0000\u0108\u0115\u0003<\u001e\u0000\u0109"+
		"\u0115\u00032\u0019\u0000\u010a\u010b\u0005\u0016\u0000\u0000\u010b\u0115"+
		"\u00036\u001b\u0000\u010c\u010d\u0005\"\u0000\u0000\u010d\u010e\u0003"+
		",\u0016\u0000\u010e\u010f\u0005#\u0000\u0000\u010f\u0115\u0001\u0000\u0000"+
		"\u0000\u0110\u0111\u0005 \u0000\u0000\u0111\u0112\u0003,\u0016\u0000\u0112"+
		"\u0113\u0005!\u0000\u0000\u0113\u0115\u0001\u0000\u0000\u0000\u0114\u0108"+
		"\u0001\u0000\u0000\u0000\u0114\u0109\u0001\u0000\u0000\u0000\u0114\u010a"+
		"\u0001\u0000\u0000\u0000\u0114\u010c\u0001\u0000\u0000\u0000\u0114\u0110"+
		"\u0001\u0000\u0000\u0000\u0115/\u0001\u0000\u0000\u0000\u0116\u011b\u0003"+
		"2\u0019\u0000\u0117\u0118\u0005\u0002\u0000\u0000\u0118\u011a\u00032\u0019"+
		"\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u011d\u0001\u0000\u0000"+
		"\u0000\u011b\u0119\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000\u0000"+
		"\u0000\u011c1\u0001\u0000\u0000\u0000\u011d\u011b\u0001\u0000\u0000\u0000"+
		"\u011e\u011f\u0003J%\u0000\u011f\u0120\u00036\u001b\u0000\u01203\u0001"+
		"\u0000\u0000\u0000\u0121\u0126\u00036\u001b\u0000\u0122\u0123\u0005\u0002"+
		"\u0000\u0000\u0123\u0125\u00036\u001b\u0000\u0124\u0122\u0001\u0000\u0000"+
		"\u0000\u0125\u0128\u0001\u0000\u0000\u0000\u0126\u0124\u0001\u0000\u0000"+
		"\u0000\u0126\u0127\u0001\u0000\u0000\u0000\u01275\u0001\u0000\u0000\u0000"+
		"\u0128\u0126\u0001\u0000\u0000\u0000\u0129\u0133\u00059\u0000\u0000\u012a"+
		"\u012b\u0005\"\u0000\u0000\u012b\u012c\u00034\u001a\u0000\u012c\u012d"+
		"\u0005#\u0000\u0000\u012d\u0133\u0001\u0000\u0000\u0000\u012e\u012f\u0005"+
		" \u0000\u0000\u012f\u0130\u00034\u001a\u0000\u0130\u0131\u0005!\u0000"+
		"\u0000\u0131\u0133\u0001\u0000\u0000\u0000\u0132\u0129\u0001\u0000\u0000"+
		"\u0000\u0132\u012a\u0001\u0000\u0000\u0000\u0132\u012e\u0001\u0000\u0000"+
		"\u0000\u01337\u0001\u0000\u0000\u0000\u0134\u0135\u0003:\u001d\u0000\u0135"+
		"\u0136\u0005\u0005\u0000\u0000\u01369\u0001\u0000\u0000\u0000\u0137\u013c"+
		"\u0003<\u001e\u0000\u0138\u0139\u0005\u0002\u0000\u0000\u0139\u013b\u0003"+
		"<\u001e\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013b\u013e\u0001\u0000"+
		"\u0000\u0000\u013c\u013a\u0001\u0000\u0000\u0000\u013c\u013d\u0001\u0000"+
		"\u0000\u0000\u013d;\u0001\u0000\u0000\u0000\u013e\u013c\u0001\u0000\u0000"+
		"\u0000\u013f\u0140\u0006\u001e\uffff\uffff\u0000\u0140\u0141\u0005\u001e"+
		"\u0000\u0000\u0141\u0142\u0003J%\u0000\u0142\u0143\u0005\u001f\u0000\u0000"+
		"\u0143\u0144\u0003<\u001e\u0014\u0144\u0151\u0001\u0000\u0000\u0000\u0145"+
		"\u0146\u0005\u001e\u0000\u0000\u0146\u0147\u0003<\u001e\u0000\u0147\u0148"+
		"\u0005\u001f\u0000\u0000\u0148\u0151\u0001\u0000\u0000\u0000\u0149\u014a"+
		"\u0007\u0001\u0000\u0000\u014a\u0151\u0003<\u001e\u0011\u014b\u0151\u0005"+
		"9\u0000\u0000\u014c\u0151\u0003N\'\u0000\u014d\u0151\u0003>\u001f\u0000"+
		"\u014e\u0151\u0003@ \u0000\u014f\u0151\u0003B!\u0000\u0150\u013f\u0001"+
		"\u0000\u0000\u0000\u0150\u0145\u0001\u0000\u0000\u0000\u0150\u0149\u0001"+
		"\u0000\u0000\u0000\u0150\u014b\u0001\u0000\u0000\u0000\u0150\u014c\u0001"+
		"\u0000\u0000\u0000\u0150\u014d\u0001\u0000\u0000\u0000\u0150\u014e\u0001"+
		"\u0000\u0000\u0000\u0150\u014f\u0001\u0000\u0000\u0000\u0151\u0184\u0001"+
		"\u0000\u0000\u0000\u0152\u0153\n\u0012\u0000\u0000\u0153\u0154\u0007\u0002"+
		"\u0000\u0000\u0154\u0183\u0003<\u001e\u0013\u0155\u0156\n\u0010\u0000"+
		"\u0000\u0156\u0157\u0007\u0003\u0000\u0000\u0157\u0183\u0003<\u001e\u0011"+
		"\u0158\u0159\n\u000f\u0000\u0000\u0159\u015a\u0007\u0004\u0000\u0000\u015a"+
		"\u0183\u0003<\u001e\u0010\u015b\u015c\n\u000e\u0000\u0000\u015c\u015d"+
		"\u0007\u0005\u0000\u0000\u015d\u0183\u0003<\u001e\u000f\u015e\u015f\n"+
		"\r\u0000\u0000\u015f\u0160\u0007\u0006\u0000\u0000\u0160\u0183\u0003<"+
		"\u001e\u000e\u0161\u0162\n\f\u0000\u0000\u0162\u0163\u0005.\u0000\u0000"+
		"\u0163\u0183\u0003<\u001e\r\u0164\u0165\n\u000b\u0000\u0000\u0165\u0166"+
		"\u0005-\u0000\u0000\u0166\u0183\u0003<\u001e\f\u0167\u0168\n\n\u0000\u0000"+
		"\u0168\u0169\u0005,\u0000\u0000\u0169\u0183\u0003<\u001e\u000b\u016a\u016b"+
		"\n\t\u0000\u0000\u016b\u016d\u0005\u001e\u0000\u0000\u016c\u016e\u0003"+
		":\u001d\u0000\u016d\u016c\u0001\u0000\u0000\u0000\u016d\u016e\u0001\u0000"+
		"\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000\u016f\u0183\u0005\u001f"+
		"\u0000\u0000\u0170\u0171\n\b\u0000\u0000\u0171\u0172\u0005\u0001\u0000"+
		"\u0000\u0172\u0183\u00059\u0000\u0000\u0173\u0174\n\u0007\u0000\u0000"+
		"\u0174\u0175\u0005 \u0000\u0000\u0175\u0176\u0003<\u001e\u0000\u0176\u0177"+
		"\u0005!\u0000\u0000\u0177\u0183\u0001\u0000\u0000\u0000\u0178\u0179\n"+
		"\u0006\u0000\u0000\u0179\u017b\u0005 \u0000\u0000\u017a\u017c\u0003<\u001e"+
		"\u0000\u017b\u017a\u0001\u0000\u0000\u0000\u017b\u017c\u0001\u0000\u0000"+
		"\u0000\u017c\u017d\u0001\u0000\u0000\u0000\u017d\u017f\u0005\u0003\u0000"+
		"\u0000\u017e\u0180\u0003<\u001e\u0000\u017f\u017e\u0001\u0000\u0000\u0000"+
		"\u017f\u0180\u0001\u0000\u0000\u0000\u0180\u0181\u0001\u0000\u0000\u0000"+
		"\u0181\u0183\u0005!\u0000\u0000\u0182\u0152\u0001\u0000\u0000\u0000\u0182"+
		"\u0155\u0001\u0000\u0000\u0000\u0182\u0158\u0001\u0000\u0000\u0000\u0182"+
		"\u015b\u0001\u0000\u0000\u0000\u0182\u015e\u0001\u0000\u0000\u0000\u0182"+
		"\u0161\u0001\u0000\u0000\u0000\u0182\u0164\u0001\u0000\u0000\u0000\u0182"+
		"\u0167\u0001\u0000\u0000\u0000\u0182\u016a\u0001\u0000\u0000\u0000\u0182"+
		"\u0170\u0001\u0000\u0000\u0000\u0182\u0173\u0001\u0000\u0000\u0000\u0182"+
		"\u0178\u0001\u0000\u0000\u0000\u0183\u0186\u0001\u0000\u0000\u0000\u0184"+
		"\u0182\u0001\u0000\u0000\u0000\u0184\u0185\u0001\u0000\u0000\u0000\u0185"+
		"=\u0001\u0000\u0000\u0000\u0186\u0184\u0001\u0000\u0000\u0000\u0187\u0188"+
		"\u0003D\"\u0000\u0188\u018a\u0005\u001e\u0000\u0000\u0189\u018b\u0003"+
		",\u0016\u0000\u018a\u0189\u0001\u0000\u0000\u0000\u018a\u018b\u0001\u0000"+
		"\u0000\u0000\u018b\u018c\u0001\u0000\u0000\u0000\u018c\u018e\u0005\u001f"+
		"\u0000\u0000\u018d\u018f\u0003J%\u0000\u018e\u018d\u0001\u0000\u0000\u0000"+
		"\u018e\u018f\u0001\u0000\u0000\u0000\u018f\u0190\u0001\u0000\u0000\u0000"+
		"\u0190\u0191\u0003\n\u0005\u0000\u0191?\u0001\u0000\u0000\u0000\u0192"+
		"\u0193\u0003F#\u0000\u0193\u0194\u0005\"\u0000\u0000\u0194\u0195\u0005"+
		"9\u0000\u0000\u0195\u0196\u0005\u0003\u0000\u0000\u0196\u019d\u0003<\u001e"+
		"\u0000\u0197\u0198\u0005\u0002\u0000\u0000\u0198\u0199\u00059\u0000\u0000"+
		"\u0199\u019a\u0005\u0003\u0000\u0000\u019a\u019c\u0003<\u001e\u0000\u019b"+
		"\u0197\u0001\u0000\u0000\u0000\u019c\u019f\u0001\u0000\u0000\u0000\u019d"+
		"\u019b\u0001\u0000\u0000\u0000\u019d\u019e\u0001\u0000\u0000\u0000\u019e"+
		"\u01a1\u0001\u0000\u0000\u0000\u019f\u019d\u0001\u0000\u0000\u0000\u01a0"+
		"\u01a2\u0005\u0002\u0000\u0000\u01a1\u01a0\u0001\u0000\u0000\u0000\u01a1"+
		"\u01a2\u0001\u0000\u0000\u0000\u01a2\u01a3\u0001\u0000\u0000\u0000\u01a3"+
		"\u01a4\u0005#\u0000\u0000\u01a4A\u0001\u0000\u0000\u0000\u01a5\u01a6\u0005"+
		"\u001e\u0000\u0000\u01a6\u01ae\u0003<\u001e\u0000\u01a7\u01af\u0005\u0002"+
		"\u0000\u0000\u01a8\u01a9\u0005\u0002\u0000\u0000\u01a9\u01ab\u0003<\u001e"+
		"\u0000\u01aa\u01a8\u0001\u0000\u0000\u0000\u01ab\u01ac\u0001\u0000\u0000"+
		"\u0000\u01ac\u01aa\u0001\u0000\u0000\u0000\u01ac\u01ad\u0001\u0000\u0000"+
		"\u0000\u01ad\u01af\u0001\u0000\u0000\u0000\u01ae\u01a7\u0001\u0000\u0000"+
		"\u0000\u01ae\u01aa\u0001\u0000\u0000\u0000\u01af\u01b0\u0001\u0000\u0000"+
		"\u0000\u01b0\u01b1\u0005\u001f\u0000\u0000\u01b1C\u0001\u0000\u0000\u0000"+
		"\u01b2\u01b5\u0005\u0019\u0000\u0000\u01b3\u01b5\u0003J%\u0000\u01b4\u01b2"+
		"\u0001\u0000\u0000\u0000\u01b4\u01b3\u0001\u0000\u0000\u0000\u01b5E\u0001"+
		"\u0000\u0000\u0000\u01b6\u01b9\u0005\u001a\u0000\u0000\u01b7\u01b9\u0003"+
		"J%\u0000\u01b8\u01b6\u0001\u0000\u0000\u0000\u01b8\u01b7\u0001\u0000\u0000"+
		"\u0000\u01b9G\u0001\u0000\u0000\u0000\u01ba\u01bd\u0005\u001b\u0000\u0000"+
		"\u01bb\u01bd\u0003J%\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bc\u01bb"+
		"\u0001\u0000\u0000\u0000\u01bdI\u0001\u0000\u0000\u0000\u01be\u01bf\u0006"+
		"%\uffff\uffff\u0000\u01bf\u01c0\u00059\u0000\u0000\u01c0\u01c6\u0001\u0000"+
		"\u0000\u0000\u01c1\u01c2\n\u0001\u0000\u0000\u01c2\u01c3\u0005\u0001\u0000"+
		"\u0000\u01c3\u01c5\u00059\u0000\u0000\u01c4\u01c1\u0001\u0000\u0000\u0000"+
		"\u01c5\u01c8\u0001\u0000\u0000\u0000\u01c6\u01c4\u0001\u0000\u0000\u0000"+
		"\u01c6\u01c7\u0001\u0000\u0000\u0000\u01c7K\u0001\u0000\u0000\u0000\u01c8"+
		"\u01c6\u0001\u0000\u0000\u0000\u01c9\u01ce\u00059\u0000\u0000\u01ca\u01cb"+
		"\u0005\u0002\u0000\u0000\u01cb\u01cd\u00059\u0000\u0000\u01cc\u01ca\u0001"+
		"\u0000\u0000\u0000\u01cd\u01d0\u0001\u0000\u0000\u0000\u01ce\u01cc\u0001"+
		"\u0000\u0000\u0000\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cfM\u0001\u0000"+
		"\u0000\u0000\u01d0\u01ce\u0001\u0000\u0000\u0000\u01d1\u01d4\u0003P(\u0000"+
		"\u01d2\u01d4\u0003R)\u0000\u01d3\u01d1\u0001\u0000\u0000\u0000\u01d3\u01d2"+
		"\u0001\u0000\u0000\u0000\u01d4O\u0001\u0000\u0000\u0000\u01d5\u01d6\u0005"+
		":\u0000\u0000\u01d6Q\u0001\u0000\u0000\u0000\u01d7\u01d8\u00059\u0000"+
		"\u0000\u01d8\u01d9\u0005:\u0000\u0000\u01d9S\u0001\u0000\u0000\u0000."+
		"W`dmz\u0080\u0089\u008e\u009d\u00a3\u00ab\u00ad\u00b3\u00ba\u00bc\u00c4"+
		"\u00cf\u00d3\u00de\u00eb\u00f2\u00fe\u0105\u0114\u011b\u0126\u0132\u013c"+
		"\u0150\u016d\u017b\u017f\u0182\u0184\u018a\u018e\u019d\u01a1\u01ac\u01ae"+
		"\u01b4\u01b8\u01bc\u01c6\u01ce\u01d3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
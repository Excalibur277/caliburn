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
		RULE_switch_statement = 14, RULE_case_blocks = 15, RULE_option_case_block = 16, 
		RULE_default_case_block = 17, RULE_inline_statement = 18, RULE_assign_statement = 19, 
		RULE_assign_expressions = 20, RULE_assign_expression = 21, RULE_assign_declarations = 22, 
		RULE_assign_declaration = 23, RULE_typed_assign_declarations = 24, RULE_typed_assign_declaration = 25, 
		RULE_untyped_assign_declarations = 26, RULE_untyped_assign_declaration = 27, 
		RULE_expression_statement = 28, RULE_expressions = 29, RULE_expression = 30, 
		RULE_function_type = 31, RULE_struct_type = 32, RULE_tuple_type = 33, 
		RULE_type_expression = 34, RULE_identifiers = 35, RULE_literal_atom = 36;
	private static String[] makeRuleNames() {
		return new String[] {
			"module", "definition", "function_definition", "parameters", "parameter", 
			"block", "statement", "jump_statement", "return_statement", "break_statement", 
			"continue_statement", "control_statement", "if_statement", "for_statement", 
			"switch_statement", "case_blocks", "option_case_block", "default_case_block", 
			"inline_statement", "assign_statement", "assign_expressions", "assign_expression", 
			"assign_declarations", "assign_declaration", "typed_assign_declarations", 
			"typed_assign_declaration", "untyped_assign_declarations", "untyped_assign_declaration", 
			"expression_statement", "expressions", "expression", "function_type", 
			"struct_type", "tuple_type", "type_expression", "identifiers", "literal_atom"
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
		public ModuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_module; }
	 
		public ModuleContext() { }
		public void copyFrom(ModuleContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ModuleRuleContext extends ModuleContext {
		public List<DefinitionContext> definition() {
			return getRuleContexts(DefinitionContext.class);
		}
		public DefinitionContext definition(int i) {
			return getRuleContext(DefinitionContext.class,i);
		}
		public ModuleRuleContext(ModuleContext ctx) { copyFrom(ctx); }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_module);
		int _la;
		try {
			_localctx = new ModuleRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC || _la==Identifier) {
				{
				{
				setState(74);
				definition();
				}
				}
				setState(79);
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
			setState(80);
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
		public Function_definitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_definition; }
	 
		public Function_definitionContext() { }
		public void copyFrom(Function_definitionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDefinitionContext extends Function_definitionContext {
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
		public FunctionDefinitionContext(Function_definitionContext ctx) { copyFrom(ctx); }
	}

	public final Function_definitionContext function_definition() throws RecognitionException {
		Function_definitionContext _localctx = new Function_definitionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_function_definition);
		int _la;
		try {
			_localctx = new FunctionDefinitionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			((FunctionDefinitionContext)_localctx).type = function_type();
			setState(83);
			((FunctionDefinitionContext)_localctx).name = match(Identifier);
			setState(84);
			match(L_PAREN);
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 144115209550692352L) != 0)) {
				{
				setState(85);
				parameters();
				}
			}

			setState(88);
			match(R_PAREN);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==Identifier) {
				{
				setState(89);
				((FunctionDefinitionContext)_localctx).return_type = type_expression(0);
				}
			}

			setState(92);
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
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	 
		public ParametersContext() { }
		public void copyFrom(ParametersContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametersRuleContext extends ParametersContext {
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
		public ParametersRuleContext(ParametersContext ctx) { copyFrom(ctx); }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_parameters);
		int _la;
		try {
			_localctx = new ParametersRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			parameter();
			setState(99);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(95);
				match(COMMA);
				setState(96);
				parameter();
				}
				}
				setState(101);
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
			setState(112);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new TypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				typed_assign_declaration();
				}
				break;
			case 2:
				_localctx = new UntypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(103);
				untyped_assign_declaration();
				}
				break;
			case 3:
				_localctx = new StructDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(L_C_BRACK);
				setState(105);
				parameters();
				setState(106);
				match(R_C_BRACK);
				}
				break;
			case 4:
				_localctx = new TupleDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(108);
				match(L_S_BRACK);
				setState(109);
				parameters();
				setState(110);
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
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	 
		public BlockContext() { }
		public void copyFrom(BlockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockRuleContext extends BlockContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public BlockRuleContext(BlockContext ctx) { copyFrom(ctx); }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_block);
		int _la;
		try {
			_localctx = new BlockRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(L_C_BRACK);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346067917573376L) != 0)) {
				{
				{
				setState(115);
				statement();
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(121);
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
			setState(127);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(123);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(124);
				expression_statement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(125);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(126);
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
			setState(132);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN:
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				return_statement();
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 2);
				{
				setState(130);
				break_statement();
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(131);
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
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
	 
		public Return_statementContext() { }
		public void copyFrom(Return_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends Return_statementContext {
		public TerminalNode RETURN() { return getToken(CaliburnParser.RETURN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public ReturnStatementContext(Return_statementContext ctx) { copyFrom(ctx); }
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_return_statement);
		try {
			_localctx = new ReturnStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(RETURN);
			setState(135);
			expression(0);
			setState(136);
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
		public Break_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_break_statement; }
	 
		public Break_statementContext() { }
		public void copyFrom(Break_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends Break_statementContext {
		public TerminalNode BREAK() { return getToken(CaliburnParser.BREAK, 0); }
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public BreakStatementContext(Break_statementContext ctx) { copyFrom(ctx); }
	}

	public final Break_statementContext break_statement() throws RecognitionException {
		Break_statementContext _localctx = new Break_statementContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_break_statement);
		try {
			_localctx = new BreakStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(138);
			match(BREAK);
			setState(139);
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
		public Continue_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continue_statement; }
	 
		public Continue_statementContext() { }
		public void copyFrom(Continue_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends Continue_statementContext {
		public TerminalNode CONTINUE() { return getToken(CaliburnParser.CONTINUE, 0); }
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public ContinueStatementContext(Continue_statementContext ctx) { copyFrom(ctx); }
	}

	public final Continue_statementContext continue_statement() throws RecognitionException {
		Continue_statementContext _localctx = new Continue_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_continue_statement);
		try {
			_localctx = new ContinueStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			match(CONTINUE);
			setState(142);
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
			setState(147);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(144);
				if_statement();
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(145);
				for_statement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 3);
				{
				setState(146);
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
		public If_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_statement; }
	 
		public If_statementContext() { }
		public void copyFrom(If_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStatementContext extends If_statementContext {
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
		public IfStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_if_statement);
		int _la;
		try {
			int _alt;
			_localctx = new IfStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(149);
			match(IF);
			setState(153);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(150);
					inline_statement();
					}
					} 
				}
				setState(155);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(156);
			expression(0);
			setState(157);
			block();
			setState(163);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(158);
				match(ELSE);
				setState(161);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case L_C_BRACK:
					{
					setState(159);
					block();
					}
					break;
				case IF:
					{
					setState(160);
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
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	 
		public For_statementContext() { }
		public void copyFrom(For_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementContext extends For_statementContext {
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
		public ForStatementContext(For_statementContext ctx) { copyFrom(ctx); }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_for_statement);
		int _la;
		try {
			int _alt;
			_localctx = new ForStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			match(FOR);
			setState(167); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(166);
					inline_statement();
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(169); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,12,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			setState(171);
			expression(0);
			setState(178);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(172);
				match(ARROW);
				setState(174); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(173);
						inline_statement();
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(176); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
			}

			setState(180);
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
		public Switch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_statement; }
	 
		public Switch_statementContext() { }
		public void copyFrom(Switch_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStatementContext extends Switch_statementContext {
		public TerminalNode SWITCH() { return getToken(CaliburnParser.SWITCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Case_blocksContext case_blocks() {
			return getRuleContext(Case_blocksContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<Inline_statementContext> inline_statement() {
			return getRuleContexts(Inline_statementContext.class);
		}
		public Inline_statementContext inline_statement(int i) {
			return getRuleContext(Inline_statementContext.class,i);
		}
		public SwitchStatementContext(Switch_statementContext ctx) { copyFrom(ctx); }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_switch_statement);
		try {
			int _alt;
			_localctx = new SwitchStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(SWITCH);
			setState(186);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(183);
					inline_statement();
					}
					} 
				}
				setState(188);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			setState(189);
			expression(0);
			setState(190);
			match(L_C_BRACK);
			setState(191);
			case_blocks();
			setState(192);
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
	public static class Case_blocksContext extends ParserRuleContext {
		public Case_blocksContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_case_blocks; }
	 
		public Case_blocksContext() { }
		public void copyFrom(Case_blocksContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CaseBlocksContext extends Case_blocksContext {
		public List<Option_case_blockContext> option_case_block() {
			return getRuleContexts(Option_case_blockContext.class);
		}
		public Option_case_blockContext option_case_block(int i) {
			return getRuleContext(Option_case_blockContext.class,i);
		}
		public Default_case_blockContext default_case_block() {
			return getRuleContext(Default_case_blockContext.class,0);
		}
		public CaseBlocksContext(Case_blocksContext ctx) { copyFrom(ctx); }
	}

	public final Case_blocksContext case_blocks() throws RecognitionException {
		Case_blocksContext _localctx = new Case_blocksContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_case_blocks);
		int _la;
		try {
			_localctx = new CaseBlocksContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(194);
				option_case_block();
				}
				}
				setState(199);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(201);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(200);
				default_case_block();
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
	public static class Option_case_blockContext extends ParserRuleContext {
		public Option_case_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_option_case_block; }
	 
		public Option_case_blockContext() { }
		public void copyFrom(Option_case_blockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OptionCaseBlockContext extends Option_case_blockContext {
		public TerminalNode CASE() { return getToken(CaliburnParser.CASE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public OptionCaseBlockContext(Option_case_blockContext ctx) { copyFrom(ctx); }
	}

	public final Option_case_blockContext option_case_block() throws RecognitionException {
		Option_case_blockContext _localctx = new Option_case_blockContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_option_case_block);
		try {
			_localctx = new OptionCaseBlockContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(203);
			match(CASE);
			setState(204);
			expression(0);
			setState(205);
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
	public static class Default_case_blockContext extends ParserRuleContext {
		public Default_case_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case_block; }
	 
		public Default_case_blockContext() { }
		public void copyFrom(Default_case_blockContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultCaseBlockContext extends Default_case_blockContext {
		public TerminalNode DEFAULT() { return getToken(CaliburnParser.DEFAULT, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public DefaultCaseBlockContext(Default_case_blockContext ctx) { copyFrom(ctx); }
	}

	public final Default_case_blockContext default_case_block() throws RecognitionException {
		Default_case_blockContext _localctx = new Default_case_blockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_default_case_block);
		try {
			_localctx = new DefaultCaseBlockContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(DEFAULT);
			setState(208);
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
			setState(212);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(211);
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
		public Assign_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_statement; }
	 
		public Assign_statementContext() { }
		public void copyFrom(Assign_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignOperationStatementContext extends Assign_statementContext {
		public Token op;
		public Assign_expressionsContext assign_expressions() {
			return getRuleContext(Assign_expressionsContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(CaliburnParser.ASSIGN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
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
		public AssignOperationStatementContext(Assign_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignStatementContext extends Assign_statementContext {
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(CaliburnParser.ASSIGN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public AssignStatementContext(Assign_statementContext ctx) { copyFrom(ctx); }
	}

	public final Assign_statementContext assign_statement() throws RecognitionException {
		Assign_statementContext _localctx = new Assign_statementContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_assign_statement);
		int _la;
		try {
			setState(225);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(214);
				assign_declarations();
				setState(215);
				match(ASSIGN);
				setState(216);
				expressions();
				setState(217);
				match(Terminator);
				}
				break;
			case 2:
				_localctx = new AssignOperationStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(219);
				assign_expressions();
				setState(220);
				((AssignOperationStatementContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 562606356037632L) != 0)) ) {
					((AssignOperationStatementContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(221);
				match(ASSIGN);
				setState(222);
				expressions();
				setState(223);
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
		public Assign_expressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_expressions; }
	 
		public Assign_expressionsContext() { }
		public void copyFrom(Assign_expressionsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignExpressionsContext extends Assign_expressionsContext {
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
		public AssignExpressionsContext(Assign_expressionsContext ctx) { copyFrom(ctx); }
	}

	public final Assign_expressionsContext assign_expressions() throws RecognitionException {
		Assign_expressionsContext _localctx = new Assign_expressionsContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assign_expressions);
		int _la;
		try {
			_localctx = new AssignExpressionsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(227);
			assign_expression();
			setState(232);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(228);
				match(COMMA);
				setState(229);
				assign_expression();
				}
				}
				setState(234);
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
			setState(244);
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
				setState(235);
				expression(0);
				}
				break;
			case L_C_BRACK:
				_localctx = new StructDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(236);
				match(L_C_BRACK);
				setState(237);
				assign_expressions();
				setState(238);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new TupleDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(240);
				match(L_S_BRACK);
				setState(241);
				assign_expressions();
				setState(242);
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
		public Assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_declarations; }
	 
		public Assign_declarationsContext() { }
		public void copyFrom(Assign_declarationsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignDeclarationsContext extends Assign_declarationsContext {
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
		public AssignDeclarationsContext(Assign_declarationsContext ctx) { copyFrom(ctx); }
	}

	public final Assign_declarationsContext assign_declarations() throws RecognitionException {
		Assign_declarationsContext _localctx = new Assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_assign_declarations);
		int _la;
		try {
			_localctx = new AssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			assign_declaration();
			setState(251);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(247);
				match(COMMA);
				setState(248);
				assign_statement();
				}
				}
				setState(253);
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
	public static class ExpressionAssignDeclarationContext extends Assign_declarationContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ExpressionAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedAssignDeclarationDeclarationContext extends Assign_declarationContext {
		public TerminalNode VAR() { return getToken(CaliburnParser.VAR, 0); }
		public Untyped_assign_declarationContext untyped_assign_declaration() {
			return getRuleContext(Untyped_assign_declarationContext.class,0);
		}
		public UntypedAssignDeclarationDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedAssignDeclarationDeclarationContext extends Assign_declarationContext {
		public Typed_assign_declarationContext typed_assign_declaration() {
			return getRuleContext(Typed_assign_declarationContext.class,0);
		}
		public TypedAssignDeclarationDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
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

	public final Assign_declarationContext assign_declaration() throws RecognitionException {
		Assign_declarationContext _localctx = new Assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_assign_declaration);
		try {
			setState(266);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new ExpressionAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(254);
				expression(0);
				}
				break;
			case 2:
				_localctx = new TypedAssignDeclarationDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(255);
				typed_assign_declaration();
				}
				break;
			case 3:
				_localctx = new UntypedAssignDeclarationDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(256);
				match(VAR);
				setState(257);
				untyped_assign_declaration();
				}
				break;
			case 4:
				_localctx = new StructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(258);
				match(L_C_BRACK);
				setState(259);
				assign_declarations();
				setState(260);
				match(R_C_BRACK);
				}
				break;
			case 5:
				_localctx = new TupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(262);
				match(L_S_BRACK);
				setState(263);
				assign_declarations();
				setState(264);
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
		public Typed_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typed_assign_declarations; }
	 
		public Typed_assign_declarationsContext() { }
		public void copyFrom(Typed_assign_declarationsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedAssignDeclarationsContext extends Typed_assign_declarationsContext {
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
		public TypedAssignDeclarationsContext(Typed_assign_declarationsContext ctx) { copyFrom(ctx); }
	}

	public final Typed_assign_declarationsContext typed_assign_declarations() throws RecognitionException {
		Typed_assign_declarationsContext _localctx = new Typed_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_typed_assign_declarations);
		int _la;
		try {
			_localctx = new TypedAssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(268);
			typed_assign_declaration();
			setState(273);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(269);
				match(COMMA);
				setState(270);
				typed_assign_declaration();
				}
				}
				setState(275);
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
	public static class TypedAssignDeclarationContext extends Typed_assign_declarationContext {
		public Type_expressionContext type;
		public Untyped_assign_declarationContext untyped_assign_declaration() {
			return getRuleContext(Untyped_assign_declarationContext.class,0);
		}
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TypedAssignDeclarationContext(Typed_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Typed_assign_declarationContext typed_assign_declaration() throws RecognitionException {
		Typed_assign_declarationContext _localctx = new Typed_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_typed_assign_declaration);
		try {
			_localctx = new TypedAssignDeclarationContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			((TypedAssignDeclarationContext)_localctx).type = type_expression(0);
			setState(277);
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
		public Untyped_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_untyped_assign_declarations; }
	 
		public Untyped_assign_declarationsContext() { }
		public void copyFrom(Untyped_assign_declarationsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedAssignDeclarationsContext extends Untyped_assign_declarationsContext {
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
		public UntypedAssignDeclarationsContext(Untyped_assign_declarationsContext ctx) { copyFrom(ctx); }
	}

	public final Untyped_assign_declarationsContext untyped_assign_declarations() throws RecognitionException {
		Untyped_assign_declarationsContext _localctx = new Untyped_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_untyped_assign_declarations);
		int _la;
		try {
			_localctx = new UntypedAssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			untyped_assign_declaration();
			setState(284);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(280);
				match(COMMA);
				setState(281);
				untyped_assign_declaration();
				}
				}
				setState(286);
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
	public static class UntypedIdentifierAssignDeclarationContext extends Untyped_assign_declarationContext {
		public Token var;
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public UntypedIdentifierAssignDeclarationContext(Untyped_assign_declarationContext ctx) { copyFrom(ctx); }
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
			setState(296);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				_localctx = new UntypedIdentifierAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(287);
				((UntypedIdentifierAssignDeclarationContext)_localctx).var = match(Identifier);
				}
				break;
			case L_C_BRACK:
				_localctx = new UntypedStructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(288);
				match(L_C_BRACK);
				setState(289);
				untyped_assign_declarations();
				setState(290);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new UntypedTupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(292);
				match(L_S_BRACK);
				setState(293);
				untyped_assign_declarations();
				setState(294);
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
		public Expression_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_statement; }
	 
		public Expression_statementContext() { }
		public void copyFrom(Expression_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionStatementContext extends Expression_statementContext {
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public ExpressionStatementContext(Expression_statementContext ctx) { copyFrom(ctx); }
	}

	public final Expression_statementContext expression_statement() throws RecognitionException {
		Expression_statementContext _localctx = new Expression_statementContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_expression_statement);
		try {
			_localctx = new ExpressionStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(298);
			expressions();
			setState(299);
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
		public ExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressions; }
	 
		public ExpressionsContext() { }
		public void copyFrom(ExpressionsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpressionsRuleContext extends ExpressionsContext {
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
		public ExpressionsRuleContext(ExpressionsContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionsContext expressions() throws RecognitionException {
		ExpressionsContext _localctx = new ExpressionsContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_expressions);
		int _la;
		try {
			_localctx = new ExpressionsRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(301);
			expression(0);
			setState(306);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(302);
				match(COMMA);
				setState(303);
				expression(0);
				}
				}
				setState(308);
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
		public FunctionExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructExpressionContext extends ExpressionContext {
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
			setState(366);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				_localctx = new CastExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(310);
				match(L_PAREN);
				setState(311);
				((CastExpressionContext)_localctx).type = type_expression(0);
				setState(312);
				match(R_PAREN);
				setState(313);
				((CastExpressionContext)_localctx).exp = expression(20);
				}
				break;
			case 2:
				{
				_localctx = new BracketedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(315);
				match(L_PAREN);
				setState(316);
				((BracketedExpressionContext)_localctx).exp = expression(0);
				setState(317);
				match(R_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(319);
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
				setState(320);
				((UnaryExpressionContext)_localctx).exp = expression(17);
				}
				break;
			case 4:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(321);
				((IdentifierExpressionContext)_localctx).identifier = match(Identifier);
				}
				break;
			case 5:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(322);
				((LiteralExpressionContext)_localctx).literal = literal_atom();
				}
				break;
			case 6:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(323);
				((FunctionExpressionContext)_localctx).type = function_type();
				setState(324);
				match(L_PAREN);
				setState(326);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346067917340672L) != 0)) {
					{
					setState(325);
					assign_declarations();
					}
				}

				setState(328);
				match(R_PAREN);
				setState(330);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Identifier) {
					{
					setState(329);
					((FunctionExpressionContext)_localctx).return_type = type_expression(0);
					}
				}

				setState(332);
				block();
				}
				break;
			case 7:
				{
				_localctx = new StructExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(334);
				((StructExpressionContext)_localctx).type = struct_type();
				setState(335);
				match(L_C_BRACK);
				{
				setState(336);
				match(Identifier);
				}
				setState(337);
				match(COLON);
				setState(338);
				expression(0);
				setState(345);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(339);
						match(COMMA);
						{
						setState(340);
						match(Identifier);
						}
						setState(341);
						match(COLON);
						setState(342);
						expression(0);
						}
						} 
					}
					setState(347);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,30,_ctx);
				}
				setState(349);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(348);
					match(COMMA);
					}
				}

				setState(351);
				match(R_C_BRACK);
				}
				break;
			case 8:
				{
				_localctx = new TupleExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(353);
				match(L_PAREN);
				setState(354);
				expression(0);
				setState(362);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
				case 1:
					{
					setState(355);
					match(COMMA);
					}
					break;
				case 2:
					{
					setState(358); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(356);
						match(COMMA);
						setState(357);
						expression(0);
						}
						}
						setState(360); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==COMMA );
					}
					break;
				}
				setState(364);
				match(R_PAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(418);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(416);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,38,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(368);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(369);
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
						setState(370);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(19);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(371);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(372);
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
						setState(373);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(17);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(374);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(375);
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
						setState(376);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(16);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(377);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(378);
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
						setState(379);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(15);
						}
						break;
					case 5:
						{
						_localctx = new BooleanBinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BooleanBinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(380);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(381);
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
						setState(382);
						((BooleanBinaryExpressionContext)_localctx).rhs_exp = expression(14);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(383);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(384);
						match(OP_AND);
						setState(385);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(13);
						}
						break;
					case 7:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(386);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(387);
						match(OP_XOR);
						setState(388);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(12);
						}
						break;
					case 8:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(389);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(390);
						match(OP_OR);
						setState(391);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(11);
						}
						break;
					case 9:
						{
						_localctx = new CallExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((CallExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(392);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(393);
						match(L_PAREN);
						setState(395);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(394);
							((CallExpressionContext)_localctx).args = expressions();
							}
						}

						setState(397);
						match(R_PAREN);
						}
						break;
					case 10:
						{
						_localctx = new AccessExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((AccessExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(398);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(399);
						match(PERIOD);
						setState(400);
						((AccessExpressionContext)_localctx).identifier = match(Identifier);
						}
						break;
					case 11:
						{
						_localctx = new IndexExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((IndexExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(401);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(402);
						match(L_S_BRACK);
						setState(403);
						((IndexExpressionContext)_localctx).index = expression(0);
						setState(404);
						match(R_S_BRACK);
						}
						break;
					case 12:
						{
						_localctx = new SliceExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((SliceExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(406);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(407);
						match(L_S_BRACK);
						setState(409);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(408);
							((SliceExpressionContext)_localctx).start_index = expression(0);
							}
						}

						setState(411);
						match(COLON);
						setState(413);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 432346046438309888L) != 0)) {
							{
							setState(412);
							((SliceExpressionContext)_localctx).end_index = expression(0);
							}
						}

						setState(415);
						match(R_S_BRACK);
						}
						break;
					}
					} 
				}
				setState(420);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,39,_ctx);
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
	public static class Function_typeContext extends ParserRuleContext {
		public Function_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_type; }
	 
		public Function_typeContext() { }
		public void copyFrom(Function_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionTypeExpressionContext extends Function_typeContext {
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public FunctionTypeExpressionContext(Function_typeContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionTypeFuncContext extends Function_typeContext {
		public TerminalNode FUNC() { return getToken(CaliburnParser.FUNC, 0); }
		public FunctionTypeFuncContext(Function_typeContext ctx) { copyFrom(ctx); }
	}

	public final Function_typeContext function_type() throws RecognitionException {
		Function_typeContext _localctx = new Function_typeContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_function_type);
		try {
			setState(423);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FunctionTypeFuncContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(421);
				match(FUNC);
				}
				break;
			case Identifier:
				_localctx = new FunctionTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(422);
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
		public Struct_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_type; }
	 
		public Struct_typeContext() { }
		public void copyFrom(Struct_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructTypeStructContext extends Struct_typeContext {
		public TerminalNode STRUCT() { return getToken(CaliburnParser.STRUCT, 0); }
		public StructTypeStructContext(Struct_typeContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructTypeExpressionContext extends Struct_typeContext {
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public StructTypeExpressionContext(Struct_typeContext ctx) { copyFrom(ctx); }
	}

	public final Struct_typeContext struct_type() throws RecognitionException {
		Struct_typeContext _localctx = new Struct_typeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_struct_type);
		try {
			setState(427);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRUCT:
				_localctx = new StructTypeStructContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(425);
				match(STRUCT);
				}
				break;
			case Identifier:
				_localctx = new StructTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(426);
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
		public Tuple_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple_type; }
	 
		public Tuple_typeContext() { }
		public void copyFrom(Tuple_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleTypeTupleContext extends Tuple_typeContext {
		public TerminalNode TUPLE() { return getToken(CaliburnParser.TUPLE, 0); }
		public TupleTypeTupleContext(Tuple_typeContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleTypeExpressionContext extends Tuple_typeContext {
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TupleTypeExpressionContext(Tuple_typeContext ctx) { copyFrom(ctx); }
	}

	public final Tuple_typeContext tuple_type() throws RecognitionException {
		Tuple_typeContext _localctx = new Tuple_typeContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_tuple_type);
		try {
			setState(431);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TUPLE:
				_localctx = new TupleTypeTupleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(429);
				match(TUPLE);
				}
				break;
			case Identifier:
				_localctx = new TupleTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(430);
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
		public Type_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_expression; }
	 
		public Type_expressionContext() { }
		public void copyFrom(Type_expressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AccessTypeExpressionContext extends Type_expressionContext {
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TerminalNode PERIOD() { return getToken(CaliburnParser.PERIOD, 0); }
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public AccessTypeExpressionContext(Type_expressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierTypeExpressionContext extends Type_expressionContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public IdentifierTypeExpressionContext(Type_expressionContext ctx) { copyFrom(ctx); }
	}

	public final Type_expressionContext type_expression() throws RecognitionException {
		return type_expression(0);
	}

	private Type_expressionContext type_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Type_expressionContext _localctx = new Type_expressionContext(_ctx, _parentState);
		Type_expressionContext _prevctx = _localctx;
		int _startState = 68;
		enterRecursionRule(_localctx, 68, RULE_type_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new IdentifierTypeExpressionContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(434);
			match(Identifier);
			}
			_ctx.stop = _input.LT(-1);
			setState(441);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,43,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AccessTypeExpressionContext(new Type_expressionContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_type_expression);
					setState(436);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(437);
					match(PERIOD);
					setState(438);
					match(Identifier);
					}
					} 
				}
				setState(443);
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
		public IdentifiersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_identifiers; }
	 
		public IdentifiersContext() { }
		public void copyFrom(IdentifiersContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifiersRuleContext extends IdentifiersContext {
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public IdentifiersRuleContext(IdentifiersContext ctx) { copyFrom(ctx); }
	}

	public final IdentifiersContext identifiers() throws RecognitionException {
		IdentifiersContext _localctx = new IdentifiersContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_identifiers);
		int _la;
		try {
			_localctx = new IdentifiersRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(444);
			match(Identifier);
			setState(449);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(445);
				match(COMMA);
				setState(446);
				match(Identifier);
				}
				}
				setState(451);
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
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public UntypedLiteralAtomContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedLiteralAtomContext extends Literal_atomContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public TypedLiteralAtomContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}

	public final Literal_atomContext literal_atom() throws RecognitionException {
		Literal_atomContext _localctx = new Literal_atomContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_literal_atom);
		try {
			setState(455);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Literal:
				_localctx = new UntypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(452);
				match(Literal);
				}
				break;
			case Identifier:
				_localctx = new TypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(453);
				match(Identifier);
				setState(454);
				match(Literal);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 30:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 34:
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
		"\u0004\u0001:\u01ca\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"#\u0007#\u0002$\u0007$\u0001\u0000\u0005\u0000L\b\u0000\n\u0000\f\u0000"+
		"O\t\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002W\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"[\b\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0005\u0003b\b\u0003\n\u0003\f\u0003e\t\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004q\b\u0004\u0001\u0005\u0001\u0005"+
		"\u0005\u0005u\b\u0005\n\u0005\f\u0005x\t\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0080\b\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u0085\b\u0007\u0001\b"+
		"\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u0094\b\u000b\u0001"+
		"\f\u0001\f\u0005\f\u0098\b\f\n\f\f\f\u009b\t\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0003\f\u00a2\b\f\u0003\f\u00a4\b\f\u0001\r\u0001\r\u0004"+
		"\r\u00a8\b\r\u000b\r\f\r\u00a9\u0001\r\u0001\r\u0001\r\u0004\r\u00af\b"+
		"\r\u000b\r\f\r\u00b0\u0003\r\u00b3\b\r\u0001\r\u0001\r\u0001\u000e\u0001"+
		"\u000e\u0005\u000e\u00b9\b\u000e\n\u000e\f\u000e\u00bc\t\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0005\u000f"+
		"\u00c4\b\u000f\n\u000f\f\u000f\u00c7\t\u000f\u0001\u000f\u0003\u000f\u00ca"+
		"\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0003\u0012\u00d5\b\u0012\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u00e2"+
		"\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00e7\b\u0014"+
		"\n\u0014\f\u0014\u00ea\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003"+
		"\u0015\u00f5\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u00fa"+
		"\b\u0016\n\u0016\f\u0016\u00fd\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u010b\b\u0017\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0005\u0018\u0110\b\u0018\n\u0018\f\u0018\u0113"+
		"\t\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0005\u001a\u011b\b\u001a\n\u001a\f\u001a\u011e\t\u001a\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0003\u001b\u0129\b\u001b\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u0131\b\u001d"+
		"\n\u001d\f\u001d\u0134\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0003\u001e\u0147\b\u001e\u0001\u001e\u0001\u001e\u0003"+
		"\u001e\u014b\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0005\u001e\u0158\b\u001e\n\u001e\f\u001e\u015b\t\u001e\u0001\u001e"+
		"\u0003\u001e\u015e\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0004\u001e\u0167\b\u001e\u000b\u001e"+
		"\f\u001e\u0168\u0003\u001e\u016b\b\u001e\u0001\u001e\u0001\u001e\u0003"+
		"\u001e\u016f\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u018c"+
		"\b\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0003\u001e\u019a\b\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u019e"+
		"\b\u001e\u0001\u001e\u0005\u001e\u01a1\b\u001e\n\u001e\f\u001e\u01a4\t"+
		"\u001e\u0001\u001f\u0001\u001f\u0003\u001f\u01a8\b\u001f\u0001 \u0001"+
		" \u0003 \u01ac\b \u0001!\u0001!\u0003!\u01b0\b!\u0001\"\u0001\"\u0001"+
		"\"\u0001\"\u0001\"\u0001\"\u0005\"\u01b8\b\"\n\"\f\"\u01bb\t\"\u0001#"+
		"\u0001#\u0001#\u0005#\u01c0\b#\n#\f#\u01c3\t#\u0001$\u0001$\u0001$\u0003"+
		"$\u01c8\b$\u0001$\u0000\u0002<D%\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFH\u0000"+
		"\u0007\u0002\u0000$%\'0\u0001\u0000$&\u0001\u0000*+\u0001\u0000\')\u0001"+
		"\u0000$%\u0001\u0000/0\u0001\u000016\u01ed\u0000M\u0001\u0000\u0000\u0000"+
		"\u0002P\u0001\u0000\u0000\u0000\u0004R\u0001\u0000\u0000\u0000\u0006^"+
		"\u0001\u0000\u0000\u0000\bp\u0001\u0000\u0000\u0000\nr\u0001\u0000\u0000"+
		"\u0000\f\u007f\u0001\u0000\u0000\u0000\u000e\u0084\u0001\u0000\u0000\u0000"+
		"\u0010\u0086\u0001\u0000\u0000\u0000\u0012\u008a\u0001\u0000\u0000\u0000"+
		"\u0014\u008d\u0001\u0000\u0000\u0000\u0016\u0093\u0001\u0000\u0000\u0000"+
		"\u0018\u0095\u0001\u0000\u0000\u0000\u001a\u00a5\u0001\u0000\u0000\u0000"+
		"\u001c\u00b6\u0001\u0000\u0000\u0000\u001e\u00c5\u0001\u0000\u0000\u0000"+
		" \u00cb\u0001\u0000\u0000\u0000\"\u00cf\u0001\u0000\u0000\u0000$\u00d4"+
		"\u0001\u0000\u0000\u0000&\u00e1\u0001\u0000\u0000\u0000(\u00e3\u0001\u0000"+
		"\u0000\u0000*\u00f4\u0001\u0000\u0000\u0000,\u00f6\u0001\u0000\u0000\u0000"+
		".\u010a\u0001\u0000\u0000\u00000\u010c\u0001\u0000\u0000\u00002\u0114"+
		"\u0001\u0000\u0000\u00004\u0117\u0001\u0000\u0000\u00006\u0128\u0001\u0000"+
		"\u0000\u00008\u012a\u0001\u0000\u0000\u0000:\u012d\u0001\u0000\u0000\u0000"+
		"<\u016e\u0001\u0000\u0000\u0000>\u01a7\u0001\u0000\u0000\u0000@\u01ab"+
		"\u0001\u0000\u0000\u0000B\u01af\u0001\u0000\u0000\u0000D\u01b1\u0001\u0000"+
		"\u0000\u0000F\u01bc\u0001\u0000\u0000\u0000H\u01c7\u0001\u0000\u0000\u0000"+
		"JL\u0003\u0002\u0001\u0000KJ\u0001\u0000\u0000\u0000LO\u0001\u0000\u0000"+
		"\u0000MK\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000N\u0001\u0001"+
		"\u0000\u0000\u0000OM\u0001\u0000\u0000\u0000PQ\u0003\u0004\u0002\u0000"+
		"Q\u0003\u0001\u0000\u0000\u0000RS\u0003>\u001f\u0000ST\u00059\u0000\u0000"+
		"TV\u0005\u001e\u0000\u0000UW\u0003\u0006\u0003\u0000VU\u0001\u0000\u0000"+
		"\u0000VW\u0001\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000XZ\u0005\u001f"+
		"\u0000\u0000Y[\u0003D\"\u0000ZY\u0001\u0000\u0000\u0000Z[\u0001\u0000"+
		"\u0000\u0000[\\\u0001\u0000\u0000\u0000\\]\u0003\n\u0005\u0000]\u0005"+
		"\u0001\u0000\u0000\u0000^c\u0003\b\u0004\u0000_`\u0005\u0002\u0000\u0000"+
		"`b\u0003\b\u0004\u0000a_\u0001\u0000\u0000\u0000be\u0001\u0000\u0000\u0000"+
		"ca\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000d\u0007\u0001\u0000"+
		"\u0000\u0000ec\u0001\u0000\u0000\u0000fq\u00032\u0019\u0000gq\u00036\u001b"+
		"\u0000hi\u0005\"\u0000\u0000ij\u0003\u0006\u0003\u0000jk\u0005#\u0000"+
		"\u0000kq\u0001\u0000\u0000\u0000lm\u0005 \u0000\u0000mn\u0003\u0006\u0003"+
		"\u0000no\u0005!\u0000\u0000oq\u0001\u0000\u0000\u0000pf\u0001\u0000\u0000"+
		"\u0000pg\u0001\u0000\u0000\u0000ph\u0001\u0000\u0000\u0000pl\u0001\u0000"+
		"\u0000\u0000q\t\u0001\u0000\u0000\u0000rv\u0005\"\u0000\u0000su\u0003"+
		"\f\u0006\u0000ts\u0001\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001"+
		"\u0000\u0000\u0000vw\u0001\u0000\u0000\u0000wy\u0001\u0000\u0000\u0000"+
		"xv\u0001\u0000\u0000\u0000yz\u0005#\u0000\u0000z\u000b\u0001\u0000\u0000"+
		"\u0000{\u0080\u0003&\u0013\u0000|\u0080\u00038\u001c\u0000}\u0080\u0003"+
		"\u0016\u000b\u0000~\u0080\u0003\u000e\u0007\u0000\u007f{\u0001\u0000\u0000"+
		"\u0000\u007f|\u0001\u0000\u0000\u0000\u007f}\u0001\u0000\u0000\u0000\u007f"+
		"~\u0001\u0000\u0000\u0000\u0080\r\u0001\u0000\u0000\u0000\u0081\u0085"+
		"\u0003\u0010\b\u0000\u0082\u0085\u0003\u0012\t\u0000\u0083\u0085\u0003"+
		"\u0014\n\u0000\u0084\u0081\u0001\u0000\u0000\u0000\u0084\u0082\u0001\u0000"+
		"\u0000\u0000\u0084\u0083\u0001\u0000\u0000\u0000\u0085\u000f\u0001\u0000"+
		"\u0000\u0000\u0086\u0087\u0005\u0011\u0000\u0000\u0087\u0088\u0003<\u001e"+
		"\u0000\u0088\u0089\u0005\u0005\u0000\u0000\u0089\u0011\u0001\u0000\u0000"+
		"\u0000\u008a\u008b\u0005\u000f\u0000\u0000\u008b\u008c\u0005\u0005\u0000"+
		"\u0000\u008c\u0013\u0001\u0000\u0000\u0000\u008d\u008e\u0005\u0010\u0000"+
		"\u0000\u008e\u008f\u0005\u0005\u0000\u0000\u008f\u0015\u0001\u0000\u0000"+
		"\u0000\u0090\u0094\u0003\u0018\f\u0000\u0091\u0094\u0003\u001a\r\u0000"+
		"\u0092\u0094\u0003\u001c\u000e\u0000\u0093\u0090\u0001\u0000\u0000\u0000"+
		"\u0093\u0091\u0001\u0000\u0000\u0000\u0093\u0092\u0001\u0000\u0000\u0000"+
		"\u0094\u0017\u0001\u0000\u0000\u0000\u0095\u0099\u0005\b\u0000\u0000\u0096"+
		"\u0098\u0003$\u0012\u0000\u0097\u0096\u0001\u0000\u0000\u0000\u0098\u009b"+
		"\u0001\u0000\u0000\u0000\u0099\u0097\u0001\u0000\u0000\u0000\u0099\u009a"+
		"\u0001\u0000\u0000\u0000\u009a\u009c\u0001\u0000\u0000\u0000\u009b\u0099"+
		"\u0001\u0000\u0000\u0000\u009c\u009d\u0003<\u001e\u0000\u009d\u00a3\u0003"+
		"\n\u0005\u0000\u009e\u00a1\u0005\t\u0000\u0000\u009f\u00a2\u0003\n\u0005"+
		"\u0000\u00a0\u00a2\u0003\u0018\f\u0000\u00a1\u009f\u0001\u0000\u0000\u0000"+
		"\u00a1\u00a0\u0001\u0000\u0000\u0000\u00a2\u00a4\u0001\u0000\u0000\u0000"+
		"\u00a3\u009e\u0001\u0000\u0000\u0000\u00a3\u00a4\u0001\u0000\u0000\u0000"+
		"\u00a4\u0019\u0001\u0000\u0000\u0000\u00a5\u00a7\u0005\n\u0000\u0000\u00a6"+
		"\u00a8\u0003$\u0012\u0000\u00a7\u00a6\u0001\u0000\u0000\u0000\u00a8\u00a9"+
		"\u0001\u0000\u0000\u0000\u00a9\u00a7\u0001\u0000\u0000\u0000\u00a9\u00aa"+
		"\u0001\u0000\u0000\u0000\u00aa\u00ab\u0001\u0000\u0000\u0000\u00ab\u00b2"+
		"\u0003<\u001e\u0000\u00ac\u00ae\u0005\u0007\u0000\u0000\u00ad\u00af\u0003"+
		"$\u0012\u0000\u00ae\u00ad\u0001\u0000\u0000\u0000\u00af\u00b0\u0001\u0000"+
		"\u0000\u0000\u00b0\u00ae\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001\u0000"+
		"\u0000\u0000\u00b1\u00b3\u0001\u0000\u0000\u0000\u00b2\u00ac\u0001\u0000"+
		"\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000\u0000\u00b3\u00b4\u0001\u0000"+
		"\u0000\u0000\u00b4\u00b5\u0003\n\u0005\u0000\u00b5\u001b\u0001\u0000\u0000"+
		"\u0000\u00b6\u00ba\u0005\u000b\u0000\u0000\u00b7\u00b9\u0003$\u0012\u0000"+
		"\u00b8\u00b7\u0001\u0000\u0000\u0000\u00b9\u00bc\u0001\u0000\u0000\u0000"+
		"\u00ba\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000\u0000"+
		"\u00bb\u00bd\u0001\u0000\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000"+
		"\u00bd\u00be\u0003<\u001e\u0000\u00be\u00bf\u0005\"\u0000\u0000\u00bf"+
		"\u00c0\u0003\u001e\u000f\u0000\u00c0\u00c1\u0005#\u0000\u0000\u00c1\u001d"+
		"\u0001\u0000\u0000\u0000\u00c2\u00c4\u0003 \u0010\u0000\u00c3\u00c2\u0001"+
		"\u0000\u0000\u0000\u00c4\u00c7\u0001\u0000\u0000\u0000\u00c5\u00c3\u0001"+
		"\u0000\u0000\u0000\u00c5\u00c6\u0001\u0000\u0000\u0000\u00c6\u00c9\u0001"+
		"\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c8\u00ca\u0003"+
		"\"\u0011\u0000\u00c9\u00c8\u0001\u0000\u0000\u0000\u00c9\u00ca\u0001\u0000"+
		"\u0000\u0000\u00ca\u001f\u0001\u0000\u0000\u0000\u00cb\u00cc\u0005\f\u0000"+
		"\u0000\u00cc\u00cd\u0003<\u001e\u0000\u00cd\u00ce\u0003\n\u0005\u0000"+
		"\u00ce!\u0001\u0000\u0000\u0000\u00cf\u00d0\u0005\r\u0000\u0000\u00d0"+
		"\u00d1\u0003\n\u0005\u0000\u00d1#\u0001\u0000\u0000\u0000\u00d2\u00d5"+
		"\u0003&\u0013\u0000\u00d3\u00d5\u00038\u001c\u0000\u00d4\u00d2\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d3\u0001\u0000\u0000\u0000\u00d5%\u0001\u0000\u0000"+
		"\u0000\u00d6\u00d7\u0003,\u0016\u0000\u00d7\u00d8\u0005\u0006\u0000\u0000"+
		"\u00d8\u00d9\u0003:\u001d\u0000\u00d9\u00da\u0005\u0005\u0000\u0000\u00da"+
		"\u00e2\u0001\u0000\u0000\u0000\u00db\u00dc\u0003(\u0014\u0000\u00dc\u00dd"+
		"\u0007\u0000\u0000\u0000\u00dd\u00de\u0005\u0006\u0000\u0000\u00de\u00df"+
		"\u0003:\u001d\u0000\u00df\u00e0\u0005\u0005\u0000\u0000\u00e0\u00e2\u0001"+
		"\u0000\u0000\u0000\u00e1\u00d6\u0001\u0000\u0000\u0000\u00e1\u00db\u0001"+
		"\u0000\u0000\u0000\u00e2\'\u0001\u0000\u0000\u0000\u00e3\u00e8\u0003*"+
		"\u0015\u0000\u00e4\u00e5\u0005\u0002\u0000\u0000\u00e5\u00e7\u0003*\u0015"+
		"\u0000\u00e6\u00e4\u0001\u0000\u0000\u0000\u00e7\u00ea\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e6\u0001\u0000\u0000\u0000\u00e8\u00e9\u0001\u0000\u0000"+
		"\u0000\u00e9)\u0001\u0000\u0000\u0000\u00ea\u00e8\u0001\u0000\u0000\u0000"+
		"\u00eb\u00f5\u0003<\u001e\u0000\u00ec\u00ed\u0005\"\u0000\u0000\u00ed"+
		"\u00ee\u0003(\u0014\u0000\u00ee\u00ef\u0005#\u0000\u0000\u00ef\u00f5\u0001"+
		"\u0000\u0000\u0000\u00f0\u00f1\u0005 \u0000\u0000\u00f1\u00f2\u0003(\u0014"+
		"\u0000\u00f2\u00f3\u0005!\u0000\u0000\u00f3\u00f5\u0001\u0000\u0000\u0000"+
		"\u00f4\u00eb\u0001\u0000\u0000\u0000\u00f4\u00ec\u0001\u0000\u0000\u0000"+
		"\u00f4\u00f0\u0001\u0000\u0000\u0000\u00f5+\u0001\u0000\u0000\u0000\u00f6"+
		"\u00fb\u0003.\u0017\u0000\u00f7\u00f8\u0005\u0002\u0000\u0000\u00f8\u00fa"+
		"\u0003&\u0013\u0000\u00f9\u00f7\u0001\u0000\u0000\u0000\u00fa\u00fd\u0001"+
		"\u0000\u0000\u0000\u00fb\u00f9\u0001\u0000\u0000\u0000\u00fb\u00fc\u0001"+
		"\u0000\u0000\u0000\u00fc-\u0001\u0000\u0000\u0000\u00fd\u00fb\u0001\u0000"+
		"\u0000\u0000\u00fe\u010b\u0003<\u001e\u0000\u00ff\u010b\u00032\u0019\u0000"+
		"\u0100\u0101\u0005\u0016\u0000\u0000\u0101\u010b\u00036\u001b\u0000\u0102"+
		"\u0103\u0005\"\u0000\u0000\u0103\u0104\u0003,\u0016\u0000\u0104\u0105"+
		"\u0005#\u0000\u0000\u0105\u010b\u0001\u0000\u0000\u0000\u0106\u0107\u0005"+
		" \u0000\u0000\u0107\u0108\u0003,\u0016\u0000\u0108\u0109\u0005!\u0000"+
		"\u0000\u0109\u010b\u0001\u0000\u0000\u0000\u010a\u00fe\u0001\u0000\u0000"+
		"\u0000\u010a\u00ff\u0001\u0000\u0000\u0000\u010a\u0100\u0001\u0000\u0000"+
		"\u0000\u010a\u0102\u0001\u0000\u0000\u0000\u010a\u0106\u0001\u0000\u0000"+
		"\u0000\u010b/\u0001\u0000\u0000\u0000\u010c\u0111\u00032\u0019\u0000\u010d"+
		"\u010e\u0005\u0002\u0000\u0000\u010e\u0110\u00032\u0019\u0000\u010f\u010d"+
		"\u0001\u0000\u0000\u0000\u0110\u0113\u0001\u0000\u0000\u0000\u0111\u010f"+
		"\u0001\u0000\u0000\u0000\u0111\u0112\u0001\u0000\u0000\u0000\u01121\u0001"+
		"\u0000\u0000\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0114\u0115\u0003"+
		"D\"\u0000\u0115\u0116\u00036\u001b\u0000\u01163\u0001\u0000\u0000\u0000"+
		"\u0117\u011c\u00036\u001b\u0000\u0118\u0119\u0005\u0002\u0000\u0000\u0119"+
		"\u011b\u00036\u001b\u0000\u011a\u0118\u0001\u0000\u0000\u0000\u011b\u011e"+
		"\u0001\u0000\u0000\u0000\u011c\u011a\u0001\u0000\u0000\u0000\u011c\u011d"+
		"\u0001\u0000\u0000\u0000\u011d5\u0001\u0000\u0000\u0000\u011e\u011c\u0001"+
		"\u0000\u0000\u0000\u011f\u0129\u00059\u0000\u0000\u0120\u0121\u0005\""+
		"\u0000\u0000\u0121\u0122\u00034\u001a\u0000\u0122\u0123\u0005#\u0000\u0000"+
		"\u0123\u0129\u0001\u0000\u0000\u0000\u0124\u0125\u0005 \u0000\u0000\u0125"+
		"\u0126\u00034\u001a\u0000\u0126\u0127\u0005!\u0000\u0000\u0127\u0129\u0001"+
		"\u0000\u0000\u0000\u0128\u011f\u0001\u0000\u0000\u0000\u0128\u0120\u0001"+
		"\u0000\u0000\u0000\u0128\u0124\u0001\u0000\u0000\u0000\u01297\u0001\u0000"+
		"\u0000\u0000\u012a\u012b\u0003:\u001d\u0000\u012b\u012c\u0005\u0005\u0000"+
		"\u0000\u012c9\u0001\u0000\u0000\u0000\u012d\u0132\u0003<\u001e\u0000\u012e"+
		"\u012f\u0005\u0002\u0000\u0000\u012f\u0131\u0003<\u001e\u0000\u0130\u012e"+
		"\u0001\u0000\u0000\u0000\u0131\u0134\u0001\u0000\u0000\u0000\u0132\u0130"+
		"\u0001\u0000\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133;\u0001"+
		"\u0000\u0000\u0000\u0134\u0132\u0001\u0000\u0000\u0000\u0135\u0136\u0006"+
		"\u001e\uffff\uffff\u0000\u0136\u0137\u0005\u001e\u0000\u0000\u0137\u0138"+
		"\u0003D\"\u0000\u0138\u0139\u0005\u001f\u0000\u0000\u0139\u013a\u0003"+
		"<\u001e\u0014\u013a\u016f\u0001\u0000\u0000\u0000\u013b\u013c\u0005\u001e"+
		"\u0000\u0000\u013c\u013d\u0003<\u001e\u0000\u013d\u013e\u0005\u001f\u0000"+
		"\u0000\u013e\u016f\u0001\u0000\u0000\u0000\u013f\u0140\u0007\u0001\u0000"+
		"\u0000\u0140\u016f\u0003<\u001e\u0011\u0141\u016f\u00059\u0000\u0000\u0142"+
		"\u016f\u0003H$\u0000\u0143\u0144\u0003>\u001f\u0000\u0144\u0146\u0005"+
		"\u001e\u0000\u0000\u0145\u0147\u0003,\u0016\u0000\u0146\u0145\u0001\u0000"+
		"\u0000\u0000\u0146\u0147\u0001\u0000\u0000\u0000\u0147\u0148\u0001\u0000"+
		"\u0000\u0000\u0148\u014a\u0005\u001f\u0000\u0000\u0149\u014b\u0003D\""+
		"\u0000\u014a\u0149\u0001\u0000\u0000\u0000\u014a\u014b\u0001\u0000\u0000"+
		"\u0000\u014b\u014c\u0001\u0000\u0000\u0000\u014c\u014d\u0003\n\u0005\u0000"+
		"\u014d\u016f\u0001\u0000\u0000\u0000\u014e\u014f\u0003@ \u0000\u014f\u0150"+
		"\u0005\"\u0000\u0000\u0150\u0151\u00059\u0000\u0000\u0151\u0152\u0005"+
		"\u0003\u0000\u0000\u0152\u0159\u0003<\u001e\u0000\u0153\u0154\u0005\u0002"+
		"\u0000\u0000\u0154\u0155\u00059\u0000\u0000\u0155\u0156\u0005\u0003\u0000"+
		"\u0000\u0156\u0158\u0003<\u001e\u0000\u0157\u0153\u0001\u0000\u0000\u0000"+
		"\u0158\u015b\u0001\u0000\u0000\u0000\u0159\u0157\u0001\u0000\u0000\u0000"+
		"\u0159\u015a\u0001\u0000\u0000\u0000\u015a\u015d\u0001\u0000\u0000\u0000"+
		"\u015b\u0159\u0001\u0000\u0000\u0000\u015c\u015e\u0005\u0002\u0000\u0000"+
		"\u015d\u015c\u0001\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000\u0000"+
		"\u015e\u015f\u0001\u0000\u0000\u0000\u015f\u0160\u0005#\u0000\u0000\u0160"+
		"\u016f\u0001\u0000\u0000\u0000\u0161\u0162\u0005\u001e\u0000\u0000\u0162"+
		"\u016a\u0003<\u001e\u0000\u0163\u016b\u0005\u0002\u0000\u0000\u0164\u0165"+
		"\u0005\u0002\u0000\u0000\u0165\u0167\u0003<\u001e\u0000\u0166\u0164\u0001"+
		"\u0000\u0000\u0000\u0167\u0168\u0001\u0000\u0000\u0000\u0168\u0166\u0001"+
		"\u0000\u0000\u0000\u0168\u0169\u0001\u0000\u0000\u0000\u0169\u016b\u0001"+
		"\u0000\u0000\u0000\u016a\u0163\u0001\u0000\u0000\u0000\u016a\u0166\u0001"+
		"\u0000\u0000\u0000\u016b\u016c\u0001\u0000\u0000\u0000\u016c\u016d\u0005"+
		"\u001f\u0000\u0000\u016d\u016f\u0001\u0000\u0000\u0000\u016e\u0135\u0001"+
		"\u0000\u0000\u0000\u016e\u013b\u0001\u0000\u0000\u0000\u016e\u013f\u0001"+
		"\u0000\u0000\u0000\u016e\u0141\u0001\u0000\u0000\u0000\u016e\u0142\u0001"+
		"\u0000\u0000\u0000\u016e\u0143\u0001\u0000\u0000\u0000\u016e\u014e\u0001"+
		"\u0000\u0000\u0000\u016e\u0161\u0001\u0000\u0000\u0000\u016f\u01a2\u0001"+
		"\u0000\u0000\u0000\u0170\u0171\n\u0012\u0000\u0000\u0171\u0172\u0007\u0002"+
		"\u0000\u0000\u0172\u01a1\u0003<\u001e\u0013\u0173\u0174\n\u0010\u0000"+
		"\u0000\u0174\u0175\u0007\u0003\u0000\u0000\u0175\u01a1\u0003<\u001e\u0011"+
		"\u0176\u0177\n\u000f\u0000\u0000\u0177\u0178\u0007\u0004\u0000\u0000\u0178"+
		"\u01a1\u0003<\u001e\u0010\u0179\u017a\n\u000e\u0000\u0000\u017a\u017b"+
		"\u0007\u0005\u0000\u0000\u017b\u01a1\u0003<\u001e\u000f\u017c\u017d\n"+
		"\r\u0000\u0000\u017d\u017e\u0007\u0006\u0000\u0000\u017e\u01a1\u0003<"+
		"\u001e\u000e\u017f\u0180\n\f\u0000\u0000\u0180\u0181\u0005.\u0000\u0000"+
		"\u0181\u01a1\u0003<\u001e\r\u0182\u0183\n\u000b\u0000\u0000\u0183\u0184"+
		"\u0005-\u0000\u0000\u0184\u01a1\u0003<\u001e\f\u0185\u0186\n\n\u0000\u0000"+
		"\u0186\u0187\u0005,\u0000\u0000\u0187\u01a1\u0003<\u001e\u000b\u0188\u0189"+
		"\n\t\u0000\u0000\u0189\u018b\u0005\u001e\u0000\u0000\u018a\u018c\u0003"+
		":\u001d\u0000\u018b\u018a\u0001\u0000\u0000\u0000\u018b\u018c\u0001\u0000"+
		"\u0000\u0000\u018c\u018d\u0001\u0000\u0000\u0000\u018d\u01a1\u0005\u001f"+
		"\u0000\u0000\u018e\u018f\n\b\u0000\u0000\u018f\u0190\u0005\u0001\u0000"+
		"\u0000\u0190\u01a1\u00059\u0000\u0000\u0191\u0192\n\u0007\u0000\u0000"+
		"\u0192\u0193\u0005 \u0000\u0000\u0193\u0194\u0003<\u001e\u0000\u0194\u0195"+
		"\u0005!\u0000\u0000\u0195\u01a1\u0001\u0000\u0000\u0000\u0196\u0197\n"+
		"\u0006\u0000\u0000\u0197\u0199\u0005 \u0000\u0000\u0198\u019a\u0003<\u001e"+
		"\u0000\u0199\u0198\u0001\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000"+
		"\u0000\u019a\u019b\u0001\u0000\u0000\u0000\u019b\u019d\u0005\u0003\u0000"+
		"\u0000\u019c\u019e\u0003<\u001e\u0000\u019d\u019c\u0001\u0000\u0000\u0000"+
		"\u019d\u019e\u0001\u0000\u0000\u0000\u019e\u019f\u0001\u0000\u0000\u0000"+
		"\u019f\u01a1\u0005!\u0000\u0000\u01a0\u0170\u0001\u0000\u0000\u0000\u01a0"+
		"\u0173\u0001\u0000\u0000\u0000\u01a0\u0176\u0001\u0000\u0000\u0000\u01a0"+
		"\u0179\u0001\u0000\u0000\u0000\u01a0\u017c\u0001\u0000\u0000\u0000\u01a0"+
		"\u017f\u0001\u0000\u0000\u0000\u01a0\u0182\u0001\u0000\u0000\u0000\u01a0"+
		"\u0185\u0001\u0000\u0000\u0000\u01a0\u0188\u0001\u0000\u0000\u0000\u01a0"+
		"\u018e\u0001\u0000\u0000\u0000\u01a0\u0191\u0001\u0000\u0000\u0000\u01a0"+
		"\u0196\u0001\u0000\u0000\u0000\u01a1\u01a4\u0001\u0000\u0000\u0000\u01a2"+
		"\u01a0\u0001\u0000\u0000\u0000\u01a2\u01a3\u0001\u0000\u0000\u0000\u01a3"+
		"=\u0001\u0000\u0000\u0000\u01a4\u01a2\u0001\u0000\u0000\u0000\u01a5\u01a8"+
		"\u0005\u0019\u0000\u0000\u01a6\u01a8\u0003D\"\u0000\u01a7\u01a5\u0001"+
		"\u0000\u0000\u0000\u01a7\u01a6\u0001\u0000\u0000\u0000\u01a8?\u0001\u0000"+
		"\u0000\u0000\u01a9\u01ac\u0005\u001a\u0000\u0000\u01aa\u01ac\u0003D\""+
		"\u0000\u01ab\u01a9\u0001\u0000\u0000\u0000\u01ab\u01aa\u0001\u0000\u0000"+
		"\u0000\u01acA\u0001\u0000\u0000\u0000\u01ad\u01b0\u0005\u001b\u0000\u0000"+
		"\u01ae\u01b0\u0003D\"\u0000\u01af\u01ad\u0001\u0000\u0000\u0000\u01af"+
		"\u01ae\u0001\u0000\u0000\u0000\u01b0C\u0001\u0000\u0000\u0000\u01b1\u01b2"+
		"\u0006\"\uffff\uffff\u0000\u01b2\u01b3\u00059\u0000\u0000\u01b3\u01b9"+
		"\u0001\u0000\u0000\u0000\u01b4\u01b5\n\u0001\u0000\u0000\u01b5\u01b6\u0005"+
		"\u0001\u0000\u0000\u01b6\u01b8\u00059\u0000\u0000\u01b7\u01b4\u0001\u0000"+
		"\u0000\u0000\u01b8\u01bb\u0001\u0000\u0000\u0000\u01b9\u01b7\u0001\u0000"+
		"\u0000\u0000\u01b9\u01ba\u0001\u0000\u0000\u0000\u01baE\u0001\u0000\u0000"+
		"\u0000\u01bb\u01b9\u0001\u0000\u0000\u0000\u01bc\u01c1\u00059\u0000\u0000"+
		"\u01bd\u01be\u0005\u0002\u0000\u0000\u01be\u01c0\u00059\u0000\u0000\u01bf"+
		"\u01bd\u0001\u0000\u0000\u0000\u01c0\u01c3\u0001\u0000\u0000\u0000\u01c1"+
		"\u01bf\u0001\u0000\u0000\u0000\u01c1\u01c2\u0001\u0000\u0000\u0000\u01c2"+
		"G\u0001\u0000\u0000\u0000\u01c3\u01c1\u0001\u0000\u0000\u0000\u01c4\u01c8"+
		"\u0005:\u0000\u0000\u01c5\u01c6\u00059\u0000\u0000\u01c6\u01c8\u0005:"+
		"\u0000\u0000\u01c7\u01c4\u0001\u0000\u0000\u0000\u01c7\u01c5\u0001\u0000"+
		"\u0000\u0000\u01c8I\u0001\u0000\u0000\u0000.MVZcpv\u007f\u0084\u0093\u0099"+
		"\u00a1\u00a3\u00a9\u00b0\u00b2\u00ba\u00c5\u00c9\u00d4\u00e1\u00e8\u00f4"+
		"\u00fb\u010a\u0111\u011c\u0128\u0132\u0146\u014a\u0159\u015d\u0168\u016a"+
		"\u016e\u018b\u0199\u019d\u01a0\u01a2\u01a7\u01ab\u01af\u01b9\u01c1\u01c7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
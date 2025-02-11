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
		PERIOD=1, COMMA=2, COLON=3, QUESTION=4, Terminator=5, ASSIGN=6, IF=7, 
		ELSE=8, FOR=9, SWITCH=10, CASE=11, DEFAULT=12, FALLTHROUGH=13, BREAK=14, 
		CONTINUE=15, RETURN=16, USING=17, AS=18, IN=19, NULL=20, VAR=21, CONST=22, 
		TYPE=23, FUNC=24, STRUCT=25, TUPLE=26, CLASS=27, INTERFACE=28, L_PAREN=29, 
		R_PAREN=30, L_S_BRACK=31, R_S_BRACK=32, L_C_BRACK=33, R_C_BRACK=34, OP_ADD=35, 
		OP_SUB=36, OP_NOT=37, OP_MUL=38, OP_DIV=39, OP_MOD=40, OP_POW=41, OP_ROOT=42, 
		OP_OR=43, OP_XOR=44, OP_AND=45, OP_LSHIFT=46, OP_RSHIFT=47, OP_EQU=48, 
		OP_NEQ=49, OP_GTE=50, OP_GRT=51, OP_LTE=52, OP_LST=53, WhiteSpace=54, 
		Comment=55, Identifier=56, Literal=57;
	public static final int
		RULE_module = 0, RULE_definition = 1, RULE_scoped_block = 2, RULE_statement = 3, 
		RULE_jump_statement = 4, RULE_return_statement = 5, RULE_break_statement = 6, 
		RULE_continue_statement = 7, RULE_control_statement = 8, RULE_if_statement = 9, 
		RULE_for_statement = 10, RULE_switch_statement = 11, RULE_inline_statement = 12, 
		RULE_function_declaration_statement = 13, RULE_assign_statement = 14, 
		RULE_assign_declarations = 15, RULE_assign_declaration = 16, RULE_declared_assign_declarations = 17, 
		RULE_declared_assign_declaration = 18, RULE_undeclared_assign_declarations = 19, 
		RULE_undeclared_assign_declaration = 20, RULE_operator_assign_statement = 21, 
		RULE_expression_statement = 22, RULE_expressions = 23, RULE_expression = 24, 
		RULE_function_expression = 25, RULE_struct_expression = 26, RULE_tuple_expression = 27, 
		RULE_type_expression = 28, RULE_complex_type_expression = 29, RULE_identifiers = 30, 
		RULE_literal_atom = 31, RULE_untyped_literal_atom = 32, RULE_typed_literal_atom = 33;
	private static String[] makeRuleNames() {
		return new String[] {
			"module", "definition", "scoped_block", "statement", "jump_statement", 
			"return_statement", "break_statement", "continue_statement", "control_statement", 
			"if_statement", "for_statement", "switch_statement", "inline_statement", 
			"function_declaration_statement", "assign_statement", "assign_declarations", 
			"assign_declaration", "declared_assign_declarations", "declared_assign_declaration", 
			"undeclared_assign_declarations", "undeclared_assign_declaration", "operator_assign_statement", 
			"expression_statement", "expressions", "expression", "function_expression", 
			"struct_expression", "tuple_expression", "type_expression", "complex_type_expression", 
			"identifiers", "literal_atom", "untyped_literal_atom", "typed_literal_atom"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "','", "':'", "'?'", "';'", "'='", "'if'", "'else'", "'for'", 
			"'switch'", "'case'", "'default'", "'fallthrough'", "'break'", "'continue'", 
			"'return'", "'using'", "'as'", "'in'", "'null'", "'var'", "'const'", 
			"'type'", "'func'", "'struct'", "'tuple'", "'class'", "'interface'", 
			"'('", "')'", "'['", "']'", "'{'", "'}'", "'+'", "'-'", "'!'", "'*'", 
			"'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", "'&'", "'<<'", "'>>'", "'=='", 
			"'!='", "'>='", "'>'", "'<='", "'<'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PERIOD", "COMMA", "COLON", "QUESTION", "Terminator", "ASSIGN", 
			"IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", "BREAK", 
			"CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", "TYPE", 
			"FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN", 
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
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057594155368448L) != 0)) {
				{
				{
				setState(68);
				definition();
				}
				}
				setState(73);
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
		public Function_declaration_statementContext function_declaration_statement() {
			return getRuleContext(Function_declaration_statementContext.class,0);
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
			setState(74);
			function_declaration_statement();
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
	public static class Scoped_blockContext extends ParserRuleContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public Scoped_blockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_scoped_block; }
	}

	public final Scoped_blockContext scoped_block() throws RecognitionException {
		Scoped_blockContext _localctx = new Scoped_blockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_scoped_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			match(L_C_BRACK);
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173034025895552L) != 0)) {
				{
				{
				setState(77);
				statement();
				}
				}
				setState(82);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(83);
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
		enterRule(_localctx, 6, RULE_statement);
		try {
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				expression_statement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(88);
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
		enterRule(_localctx, 8, RULE_jump_statement);
		try {
			setState(94);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN:
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				return_statement();
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				break_statement();
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
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
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public Return_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_return_statement; }
	}

	public final Return_statementContext return_statement() throws RecognitionException {
		Return_statementContext _localctx = new Return_statementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_return_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(RETURN);
			setState(98);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023286263808L) != 0)) {
				{
				setState(97);
				expressions();
				}
			}

			setState(100);
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
		enterRule(_localctx, 12, RULE_break_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			match(BREAK);
			setState(103);
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
		enterRule(_localctx, 14, RULE_continue_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(CONTINUE);
			setState(106);
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
		public Function_declaration_statementContext function_declaration_statement() {
			return getRuleContext(Function_declaration_statementContext.class,0);
		}
		public Control_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_control_statement; }
	}

	public final Control_statementContext control_statement() throws RecognitionException {
		Control_statementContext _localctx = new Control_statementContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_control_statement);
		try {
			setState(112);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(108);
				if_statement();
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(109);
				for_statement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 3);
				{
				setState(110);
				switch_statement();
				}
				break;
			case FUNC:
			case STRUCT:
			case TUPLE:
			case Identifier:
				enterOuterAlt(_localctx, 4);
				{
				setState(111);
				function_declaration_statement();
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
		public List<Scoped_blockContext> scoped_block() {
			return getRuleContexts(Scoped_blockContext.class);
		}
		public Scoped_blockContext scoped_block(int i) {
			return getRuleContext(Scoped_blockContext.class,i);
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
		enterRule(_localctx, 18, RULE_if_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			match(IF);
			setState(118);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(115);
					inline_statement();
					}
					} 
				}
				setState(120);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			setState(121);
			expression(0);
			setState(122);
			scoped_block();
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(123);
				match(ELSE);
				setState(126);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case L_C_BRACK:
					{
					setState(124);
					scoped_block();
					}
					break;
				case IF:
					{
					setState(125);
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
		public Scoped_blockContext scoped_block() {
			return getRuleContext(Scoped_blockContext.class,0);
		}
		public For_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_statement; }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_for_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			match(FOR);
			setState(131);
			expression(0);
			setState(132);
			scoped_block();
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
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public List<Inline_statementContext> inline_statement() {
			return getRuleContexts(Inline_statementContext.class);
		}
		public Inline_statementContext inline_statement(int i) {
			return getRuleContext(Inline_statementContext.class,i);
		}
		public List<TerminalNode> CASE() { return getTokens(CaliburnParser.CASE); }
		public TerminalNode CASE(int i) {
			return getToken(CaliburnParser.CASE, i);
		}
		public List<Scoped_blockContext> scoped_block() {
			return getRuleContexts(Scoped_blockContext.class);
		}
		public Scoped_blockContext scoped_block(int i) {
			return getRuleContext(Scoped_blockContext.class,i);
		}
		public TerminalNode DEFAULT() { return getToken(CaliburnParser.DEFAULT, 0); }
		public List<Type_expressionContext> type_expression() {
			return getRuleContexts(Type_expressionContext.class);
		}
		public Type_expressionContext type_expression(int i) {
			return getRuleContext(Type_expressionContext.class,i);
		}
		public Switch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_statement; }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_switch_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(134);
			match(SWITCH);
			setState(138);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(135);
					inline_statement();
					}
					} 
				}
				setState(140);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			}
			setState(141);
			expression(0);
			setState(142);
			match(L_C_BRACK);
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(143);
				match(CASE);
				setState(146);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
				case 1:
					{
					setState(144);
					type_expression();
					}
					break;
				case 2:
					{
					setState(145);
					expression(0);
					}
					break;
				}
				setState(148);
				scoped_block();
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(155);
				match(DEFAULT);
				setState(156);
				scoped_block();
				}
			}

			setState(159);
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
		enterRule(_localctx, 24, RULE_inline_statement);
		try {
			setState(163);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(161);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(162);
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
	public static class Function_declaration_statementContext extends ParserRuleContext {
		public Type_expressionContext type;
		public Token name;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public Scoped_blockContext scoped_block() {
			return getRuleContext(Scoped_blockContext.class,0);
		}
		public List<Type_expressionContext> type_expression() {
			return getRuleContexts(Type_expressionContext.class);
		}
		public Type_expressionContext type_expression(int i) {
			return getRuleContext(Type_expressionContext.class,i);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public Function_declaration_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration_statement; }
	}

	public final Function_declaration_statementContext function_declaration_statement() throws RecognitionException {
		Function_declaration_statementContext _localctx = new Function_declaration_statementContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_function_declaration_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(165);
			((Function_declaration_statementContext)_localctx).type = type_expression();
			setState(166);
			((Function_declaration_statementContext)_localctx).name = match(Identifier);
			setState(167);
			match(L_PAREN);
			setState(169);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057605431754752L) != 0)) {
				{
				setState(168);
				assign_declarations();
				}
			}

			setState(171);
			match(R_PAREN);
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057594155368448L) != 0)) {
				{
				setState(172);
				((Function_declaration_statementContext)_localctx).return_type = type_expression();
				}
			}

			setState(175);
			scoped_block();
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
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(CaliburnParser.ASSIGN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public Operator_assign_statementContext operator_assign_statement() {
			return getRuleContext(Operator_assign_statementContext.class,0);
		}
		public Assign_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_statement; }
	}

	public final Assign_statementContext assign_statement() throws RecognitionException {
		Assign_statementContext _localctx = new Assign_statementContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_assign_statement);
		try {
			setState(183);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				assign_declarations();
				setState(178);
				match(ASSIGN);
				setState(179);
				expressions();
				setState(180);
				match(Terminator);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(182);
				operator_assign_statement();
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
		enterRule(_localctx, 30, RULE_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(185);
			assign_declaration();
			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(186);
				match(COMMA);
				setState(187);
				assign_statement();
				}
				}
				setState(192);
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
	public static class UndeclaredAssignDeclarationContext extends Assign_declarationContext {
		public Undeclared_assign_declarationContext undeclared_assign_declaration() {
			return getRuleContext(Undeclared_assign_declarationContext.class,0);
		}
		public UndeclaredAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclaredAssignDeclarationContext extends Assign_declarationContext {
		public Declared_assign_declarationContext declared_assign_declaration() {
			return getRuleContext(Declared_assign_declarationContext.class,0);
		}
		public DeclaredAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
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
		public List<TerminalNode> L_S_BRACK() { return getTokens(CaliburnParser.L_S_BRACK); }
		public TerminalNode L_S_BRACK(int i) {
			return getToken(CaliburnParser.L_S_BRACK, i);
		}
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TupleDestrutureAssignDeclarationContext(Assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Assign_declarationContext assign_declaration() throws RecognitionException {
		Assign_declarationContext _localctx = new Assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_assign_declaration);
		try {
			setState(203);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				_localctx = new DeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				declared_assign_declaration();
				}
				break;
			case 2:
				_localctx = new UndeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				undeclared_assign_declaration();
				}
				break;
			case 3:
				_localctx = new StructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				match(L_C_BRACK);
				setState(196);
				assign_declarations();
				setState(197);
				match(R_C_BRACK);
				}
				break;
			case 4:
				_localctx = new TupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(199);
				match(L_S_BRACK);
				setState(200);
				assign_declarations();
				setState(201);
				match(L_S_BRACK);
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
	public static class Declared_assign_declarationsContext extends ParserRuleContext {
		public List<Declared_assign_declarationContext> declared_assign_declaration() {
			return getRuleContexts(Declared_assign_declarationContext.class);
		}
		public Declared_assign_declarationContext declared_assign_declaration(int i) {
			return getRuleContext(Declared_assign_declarationContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Declared_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declared_assign_declarations; }
	}

	public final Declared_assign_declarationsContext declared_assign_declarations() throws RecognitionException {
		Declared_assign_declarationsContext _localctx = new Declared_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_declared_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			declared_assign_declaration();
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(206);
				match(COMMA);
				setState(207);
				declared_assign_declaration();
				}
				}
				setState(212);
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
	public static class Declared_assign_declarationContext extends ParserRuleContext {
		public Declared_assign_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declared_assign_declaration; }
	 
		public Declared_assign_declarationContext() { }
		public void copyFrom(Declared_assign_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedDeclaredAssignDeclarationContext extends Declared_assign_declarationContext {
		public TerminalNode VAR() { return getToken(CaliburnParser.VAR, 0); }
		public Undeclared_assign_declarationContext undeclared_assign_declaration() {
			return getRuleContext(Undeclared_assign_declarationContext.class,0);
		}
		public UntypedDeclaredAssignDeclarationContext(Declared_assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedDeclaredAssignDeclarationContext extends Declared_assign_declarationContext {
		public Type_expressionContext type;
		public Undeclared_assign_declarationContext undeclared_assign_declaration() {
			return getRuleContext(Undeclared_assign_declarationContext.class,0);
		}
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public TypedDeclaredAssignDeclarationContext(Declared_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Declared_assign_declarationContext declared_assign_declaration() throws RecognitionException {
		Declared_assign_declarationContext _localctx = new Declared_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_declared_assign_declaration);
		try {
			setState(218);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
			case STRUCT:
			case TUPLE:
			case Identifier:
				_localctx = new TypedDeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(213);
				((TypedDeclaredAssignDeclarationContext)_localctx).type = type_expression();
				setState(214);
				undeclared_assign_declaration();
				}
				break;
			case VAR:
				_localctx = new UntypedDeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				match(VAR);
				setState(217);
				undeclared_assign_declaration();
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
	public static class Undeclared_assign_declarationsContext extends ParserRuleContext {
		public List<Undeclared_assign_declarationContext> undeclared_assign_declaration() {
			return getRuleContexts(Undeclared_assign_declarationContext.class);
		}
		public Undeclared_assign_declarationContext undeclared_assign_declaration(int i) {
			return getRuleContext(Undeclared_assign_declarationContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(CaliburnParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(CaliburnParser.COMMA, i);
		}
		public Undeclared_assign_declarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_undeclared_assign_declarations; }
	}

	public final Undeclared_assign_declarationsContext undeclared_assign_declarations() throws RecognitionException {
		Undeclared_assign_declarationsContext _localctx = new Undeclared_assign_declarationsContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_undeclared_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(220);
			undeclared_assign_declaration();
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(221);
				match(COMMA);
				setState(222);
				undeclared_assign_declaration();
				}
				}
				setState(227);
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
	public static class Undeclared_assign_declarationContext extends ParserRuleContext {
		public Undeclared_assign_declarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_undeclared_assign_declaration; }
	 
		public Undeclared_assign_declarationContext() { }
		public void copyFrom(Undeclared_assign_declarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UndeclaredAtomAssignDeclarationContext extends Undeclared_assign_declarationContext {
		public Type_expressionContext var;
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public UndeclaredAtomAssignDeclarationContext(Undeclared_assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UndeclaredTupleDestrutureAssignDeclarationContext extends Undeclared_assign_declarationContext {
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public Undeclared_assign_declarationsContext undeclared_assign_declarations() {
			return getRuleContext(Undeclared_assign_declarationsContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public UndeclaredTupleDestrutureAssignDeclarationContext(Undeclared_assign_declarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UndeclaredStructDestrutureAssignDeclarationContext extends Undeclared_assign_declarationContext {
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Undeclared_assign_declarationsContext undeclared_assign_declarations() {
			return getRuleContext(Undeclared_assign_declarationsContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public UndeclaredStructDestrutureAssignDeclarationContext(Undeclared_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Undeclared_assign_declarationContext undeclared_assign_declaration() throws RecognitionException {
		Undeclared_assign_declarationContext _localctx = new Undeclared_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_undeclared_assign_declaration);
		try {
			setState(237);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
			case STRUCT:
			case TUPLE:
			case Identifier:
				_localctx = new UndeclaredAtomAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(228);
				((UndeclaredAtomAssignDeclarationContext)_localctx).var = type_expression();
				}
				break;
			case L_C_BRACK:
				_localctx = new UndeclaredStructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(229);
				match(L_C_BRACK);
				setState(230);
				undeclared_assign_declarations();
				setState(231);
				match(R_C_BRACK);
				}
				break;
			case L_PAREN:
				_localctx = new UndeclaredTupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(233);
				match(L_PAREN);
				setState(234);
				undeclared_assign_declarations();
				setState(235);
				match(R_PAREN);
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
	public static class Operator_assign_statementContext extends ParserRuleContext {
		public Operator_assign_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_assign_statement; }
	 
		public Operator_assign_statementContext() { }
		public void copyFrom(Operator_assign_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OperatorAssignmentContext extends Operator_assign_statementContext {
		public ExpressionContext var;
		public Token op;
		public TerminalNode ASSIGN() { return getToken(CaliburnParser.ASSIGN, 0); }
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		public OperatorAssignmentContext(Operator_assign_statementContext ctx) { copyFrom(ctx); }
	}

	public final Operator_assign_statementContext operator_assign_statement() throws RecognitionException {
		Operator_assign_statementContext _localctx = new Operator_assign_statementContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_operator_assign_statement);
		int _la;
		try {
			_localctx = new OperatorAssignmentContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			((OperatorAssignmentContext)_localctx).var = expression(0);
			setState(240);
			((OperatorAssignmentContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 281303178018816L) != 0)) ) {
				((OperatorAssignmentContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(241);
			match(ASSIGN);
			setState(242);
			expressions();
			setState(243);
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
		enterRule(_localctx, 44, RULE_expression_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			expressions();
			setState(246);
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
		enterRule(_localctx, 46, RULE_expressions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			expression(0);
			setState(253);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(249);
				match(COMMA);
				setState(250);
				expression(0);
				}
				}
				setState(255);
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
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				{
				_localctx = new CastExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(257);
				match(L_PAREN);
				setState(258);
				((CastExpressionContext)_localctx).type = type_expression();
				setState(259);
				match(R_PAREN);
				setState(260);
				((CastExpressionContext)_localctx).exp = expression(20);
				}
				break;
			case 2:
				{
				_localctx = new BracketedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(262);
				match(L_PAREN);
				setState(263);
				((BracketedExpressionContext)_localctx).exp = expression(0);
				setState(264);
				match(R_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(266);
				((UnaryExpressionContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 240518168576L) != 0)) ) {
					((UnaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(267);
				((UnaryExpressionContext)_localctx).exp = expression(17);
				}
				break;
			case 4:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(268);
				((IdentifierExpressionContext)_localctx).identifier = match(Identifier);
				}
				break;
			case 5:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(269);
				((LiteralExpressionContext)_localctx).literal = literal_atom();
				}
				break;
			case 6:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(270);
				((FunctionExpressionContext)_localctx).exp = function_expression();
				}
				break;
			case 7:
				{
				_localctx = new StructExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(271);
				((StructExpressionContext)_localctx).exp = struct_expression();
				}
				break;
			case 8:
				{
				_localctx = new TupleExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(272);
				((TupleExpressionContext)_localctx).exp = tuple_expression();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(325);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(323);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(275);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(276);
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
						setState(277);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(19);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(278);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(279);
						((BinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1924145348608L) != 0)) ) {
							((BinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(280);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(17);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(281);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(282);
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
						setState(283);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(16);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(284);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(285);
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
						setState(286);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(15);
						}
						break;
					case 5:
						{
						_localctx = new BooleanBinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BooleanBinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(287);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(288);
						((BooleanBinaryExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 17732923532771328L) != 0)) ) {
							((BooleanBinaryExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(289);
						((BooleanBinaryExpressionContext)_localctx).rhs_exp = expression(14);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(290);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(291);
						match(OP_AND);
						setState(292);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(13);
						}
						break;
					case 7:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(293);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(294);
						match(OP_XOR);
						setState(295);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(12);
						}
						break;
					case 8:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(296);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(297);
						match(OP_OR);
						setState(298);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(11);
						}
						break;
					case 9:
						{
						_localctx = new CallExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((CallExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(299);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(300);
						match(L_PAREN);
						setState(302);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023286263808L) != 0)) {
							{
							setState(301);
							((CallExpressionContext)_localctx).args = expressions();
							}
						}

						setState(304);
						match(R_PAREN);
						}
						break;
					case 10:
						{
						_localctx = new AccessExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((AccessExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(305);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(306);
						match(PERIOD);
						setState(307);
						((AccessExpressionContext)_localctx).identifier = match(Identifier);
						}
						break;
					case 11:
						{
						_localctx = new IndexExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((IndexExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(308);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(309);
						match(L_S_BRACK);
						setState(310);
						((IndexExpressionContext)_localctx).index = expression(0);
						setState(311);
						match(R_S_BRACK);
						}
						break;
					case 12:
						{
						_localctx = new SliceExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((SliceExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(313);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(314);
						match(L_S_BRACK);
						setState(316);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023286263808L) != 0)) {
							{
							setState(315);
							((SliceExpressionContext)_localctx).start_index = expression(0);
							}
						}

						setState(318);
						match(COLON);
						setState(320);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023286263808L) != 0)) {
							{
							setState(319);
							((SliceExpressionContext)_localctx).end_index = expression(0);
							}
						}

						setState(322);
						match(R_S_BRACK);
						}
						break;
					}
					} 
				}
				setState(327);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
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
		public Type_expressionContext type;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public Scoped_blockContext scoped_block() {
			return getRuleContext(Scoped_blockContext.class,0);
		}
		public List<Type_expressionContext> type_expression() {
			return getRuleContexts(Type_expressionContext.class);
		}
		public Type_expressionContext type_expression(int i) {
			return getRuleContext(Type_expressionContext.class,i);
		}
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public Function_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_expression; }
	}

	public final Function_expressionContext function_expression() throws RecognitionException {
		Function_expressionContext _localctx = new Function_expressionContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(328);
			((Function_expressionContext)_localctx).type = type_expression();
			setState(329);
			match(L_PAREN);
			setState(331);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057605431754752L) != 0)) {
				{
				setState(330);
				assign_declarations();
				}
			}

			setState(333);
			match(R_PAREN);
			setState(335);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057594155368448L) != 0)) {
				{
				setState(334);
				((Function_expressionContext)_localctx).return_type = type_expression();
				}
			}

			setState(337);
			scoped_block();
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
		public Type_expressionContext type;
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
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
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
		enterRule(_localctx, 52, RULE_struct_expression);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			((Struct_expressionContext)_localctx).type = type_expression();
			setState(340);
			match(L_C_BRACK);
			{
			setState(341);
			match(Identifier);
			}
			setState(342);
			match(COLON);
			setState(343);
			expression(0);
			setState(350);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(344);
					match(COMMA);
					{
					setState(345);
					match(Identifier);
					}
					setState(346);
					match(COLON);
					setState(347);
					expression(0);
					}
					} 
				}
				setState(352);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
			}
			setState(354);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(353);
				match(COMMA);
				}
			}

			setState(356);
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
		enterRule(_localctx, 54, RULE_tuple_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(358);
			match(L_PAREN);
			setState(359);
			expression(0);
			setState(367);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				{
				setState(360);
				match(COMMA);
				}
				break;
			case 2:
				{
				setState(363); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(361);
					match(COMMA);
					setState(362);
					expression(0);
					}
					}
					setState(365); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				}
				break;
			}
			setState(369);
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
	public static class Type_expressionContext extends ParserRuleContext {
		public Complex_type_expressionContext complex_type_expression() {
			return getRuleContext(Complex_type_expressionContext.class,0);
		}
		public TerminalNode FUNC() { return getToken(CaliburnParser.FUNC, 0); }
		public TerminalNode STRUCT() { return getToken(CaliburnParser.STRUCT, 0); }
		public TerminalNode TUPLE() { return getToken(CaliburnParser.TUPLE, 0); }
		public Type_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_expression; }
	}

	public final Type_expressionContext type_expression() throws RecognitionException {
		Type_expressionContext _localctx = new Type_expressionContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_type_expression);
		try {
			setState(375);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				enterOuterAlt(_localctx, 1);
				{
				setState(371);
				complex_type_expression(0);
				}
				break;
			case FUNC:
				enterOuterAlt(_localctx, 2);
				{
				setState(372);
				match(FUNC);
				}
				break;
			case STRUCT:
				enterOuterAlt(_localctx, 3);
				{
				setState(373);
				match(STRUCT);
				}
				break;
			case TUPLE:
				enterOuterAlt(_localctx, 4);
				{
				setState(374);
				match(TUPLE);
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
	public static class Complex_type_expressionContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Complex_type_expressionContext complex_type_expression() {
			return getRuleContext(Complex_type_expressionContext.class,0);
		}
		public TerminalNode PERIOD() { return getToken(CaliburnParser.PERIOD, 0); }
		public Complex_type_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_complex_type_expression; }
	}

	public final Complex_type_expressionContext complex_type_expression() throws RecognitionException {
		return complex_type_expression(0);
	}

	private Complex_type_expressionContext complex_type_expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Complex_type_expressionContext _localctx = new Complex_type_expressionContext(_ctx, _parentState);
		Complex_type_expressionContext _prevctx = _localctx;
		int _startState = 58;
		enterRecursionRule(_localctx, 58, RULE_complex_type_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(378);
			match(Identifier);
			}
			_ctx.stop = _input.LT(-1);
			setState(385);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Complex_type_expressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_complex_type_expression);
					setState(380);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(381);
					match(PERIOD);
					setState(382);
					match(Identifier);
					}
					} 
				}
				setState(387);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
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
		enterRule(_localctx, 60, RULE_identifiers);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(388);
			match(Identifier);
			setState(393);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(389);
				match(COMMA);
				setState(390);
				match(Identifier);
				}
				}
				setState(395);
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
		enterRule(_localctx, 62, RULE_literal_atom);
		try {
			setState(398);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Literal:
				_localctx = new UntypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(396);
				untyped_literal_atom();
				}
				break;
			case Identifier:
				_localctx = new TypedLiteralAtomContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(397);
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
		enterRule(_localctx, 64, RULE_untyped_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(400);
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
		enterRule(_localctx, 66, RULE_typed_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(402);
			match(Identifier);
			setState(403);
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
		case 24:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 29:
			return complex_type_expression_sempred((Complex_type_expressionContext)_localctx, predIndex);
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
	private boolean complex_type_expression_sempred(Complex_type_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u00019\u0196\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0001\u0000\u0005"+
		"\u0000F\b\u0000\n\u0000\f\u0000I\t\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0002\u0001\u0002\u0005\u0002O\b\u0002\n\u0002\f\u0002R\t\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003"+
		"\u0003Z\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004_\b\u0004"+
		"\u0001\u0005\u0001\u0005\u0003\u0005c\b\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0003\bq\b\b\u0001\t\u0001\t\u0005\t"+
		"u\b\t\n\t\f\tx\t\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u007f"+
		"\b\t\u0003\t\u0081\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001"+
		"\u000b\u0005\u000b\u0089\b\u000b\n\u000b\f\u000b\u008c\t\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u0093\b\u000b"+
		"\u0001\u000b\u0001\u000b\u0005\u000b\u0097\b\u000b\n\u000b\f\u000b\u009a"+
		"\t\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u009e\b\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0003\f\u00a4\b\f\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0003\r\u00aa\b\r\u0001\r\u0001\r\u0003\r\u00ae\b\r\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u00b8\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u00bd\b\u000f\n\u000f\f\u000f\u00c0\t\u000f\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u00cc\b\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0005\u0011\u00d1\b\u0011\n\u0011\f\u0011\u00d4\t\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00db\b\u0012"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u00e0\b\u0013\n\u0013"+
		"\f\u0013\u00e3\t\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u00ee\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0005\u0017\u00fc\b\u0017\n\u0017\f\u0017\u00ff\t\u0017\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0112"+
		"\b\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u012f\b\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003"+
		"\u0018\u013d\b\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0141\b\u0018"+
		"\u0001\u0018\u0005\u0018\u0144\b\u0018\n\u0018\f\u0018\u0147\t\u0018\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u014c\b\u0019\u0001\u0019\u0001"+
		"\u0019\u0003\u0019\u0150\b\u0019\u0001\u0019\u0001\u0019\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0005\u001a\u015d\b\u001a\n\u001a\f\u001a\u0160\t\u001a"+
		"\u0001\u001a\u0003\u001a\u0163\b\u001a\u0001\u001a\u0001\u001a\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0004\u001b\u016c\b\u001b"+
		"\u000b\u001b\f\u001b\u016d\u0003\u001b\u0170\b\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u0178"+
		"\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0005\u001d\u0180\b\u001d\n\u001d\f\u001d\u0183\t\u001d\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0005\u001e\u0188\b\u001e\n\u001e\f\u001e\u018b"+
		"\t\u001e\u0001\u001f\u0001\u001f\u0003\u001f\u018f\b\u001f\u0001 \u0001"+
		" \u0001!\u0001!\u0001!\u0001!\u0000\u00020:\"\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,."+
		"02468:<>@B\u0000\u0007\u0002\u0000#$&/\u0001\u0000#%\u0001\u0000)*\u0001"+
		"\u0000&(\u0001\u0000#$\u0001\u0000./\u0001\u000005\u01b5\u0000G\u0001"+
		"\u0000\u0000\u0000\u0002J\u0001\u0000\u0000\u0000\u0004L\u0001\u0000\u0000"+
		"\u0000\u0006Y\u0001\u0000\u0000\u0000\b^\u0001\u0000\u0000\u0000\n`\u0001"+
		"\u0000\u0000\u0000\ff\u0001\u0000\u0000\u0000\u000ei\u0001\u0000\u0000"+
		"\u0000\u0010p\u0001\u0000\u0000\u0000\u0012r\u0001\u0000\u0000\u0000\u0014"+
		"\u0082\u0001\u0000\u0000\u0000\u0016\u0086\u0001\u0000\u0000\u0000\u0018"+
		"\u00a3\u0001\u0000\u0000\u0000\u001a\u00a5\u0001\u0000\u0000\u0000\u001c"+
		"\u00b7\u0001\u0000\u0000\u0000\u001e\u00b9\u0001\u0000\u0000\u0000 \u00cb"+
		"\u0001\u0000\u0000\u0000\"\u00cd\u0001\u0000\u0000\u0000$\u00da\u0001"+
		"\u0000\u0000\u0000&\u00dc\u0001\u0000\u0000\u0000(\u00ed\u0001\u0000\u0000"+
		"\u0000*\u00ef\u0001\u0000\u0000\u0000,\u00f5\u0001\u0000\u0000\u0000."+
		"\u00f8\u0001\u0000\u0000\u00000\u0111\u0001\u0000\u0000\u00002\u0148\u0001"+
		"\u0000\u0000\u00004\u0153\u0001\u0000\u0000\u00006\u0166\u0001\u0000\u0000"+
		"\u00008\u0177\u0001\u0000\u0000\u0000:\u0179\u0001\u0000\u0000\u0000<"+
		"\u0184\u0001\u0000\u0000\u0000>\u018e\u0001\u0000\u0000\u0000@\u0190\u0001"+
		"\u0000\u0000\u0000B\u0192\u0001\u0000\u0000\u0000DF\u0003\u0002\u0001"+
		"\u0000ED\u0001\u0000\u0000\u0000FI\u0001\u0000\u0000\u0000GE\u0001\u0000"+
		"\u0000\u0000GH\u0001\u0000\u0000\u0000H\u0001\u0001\u0000\u0000\u0000"+
		"IG\u0001\u0000\u0000\u0000JK\u0003\u001a\r\u0000K\u0003\u0001\u0000\u0000"+
		"\u0000LP\u0005!\u0000\u0000MO\u0003\u0006\u0003\u0000NM\u0001\u0000\u0000"+
		"\u0000OR\u0001\u0000\u0000\u0000PN\u0001\u0000\u0000\u0000PQ\u0001\u0000"+
		"\u0000\u0000QS\u0001\u0000\u0000\u0000RP\u0001\u0000\u0000\u0000ST\u0005"+
		"\"\u0000\u0000T\u0005\u0001\u0000\u0000\u0000UZ\u0003\u001c\u000e\u0000"+
		"VZ\u0003,\u0016\u0000WZ\u0003\u0010\b\u0000XZ\u0003\b\u0004\u0000YU\u0001"+
		"\u0000\u0000\u0000YV\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000"+
		"YX\u0001\u0000\u0000\u0000Z\u0007\u0001\u0000\u0000\u0000[_\u0003\n\u0005"+
		"\u0000\\_\u0003\f\u0006\u0000]_\u0003\u000e\u0007\u0000^[\u0001\u0000"+
		"\u0000\u0000^\\\u0001\u0000\u0000\u0000^]\u0001\u0000\u0000\u0000_\t\u0001"+
		"\u0000\u0000\u0000`b\u0005\u0010\u0000\u0000ac\u0003.\u0017\u0000ba\u0001"+
		"\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000"+
		"de\u0005\u0005\u0000\u0000e\u000b\u0001\u0000\u0000\u0000fg\u0005\u000e"+
		"\u0000\u0000gh\u0005\u0005\u0000\u0000h\r\u0001\u0000\u0000\u0000ij\u0005"+
		"\u000f\u0000\u0000jk\u0005\u0005\u0000\u0000k\u000f\u0001\u0000\u0000"+
		"\u0000lq\u0003\u0012\t\u0000mq\u0003\u0014\n\u0000nq\u0003\u0016\u000b"+
		"\u0000oq\u0003\u001a\r\u0000pl\u0001\u0000\u0000\u0000pm\u0001\u0000\u0000"+
		"\u0000pn\u0001\u0000\u0000\u0000po\u0001\u0000\u0000\u0000q\u0011\u0001"+
		"\u0000\u0000\u0000rv\u0005\u0007\u0000\u0000su\u0003\u0018\f\u0000ts\u0001"+
		"\u0000\u0000\u0000ux\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000\u0000"+
		"vw\u0001\u0000\u0000\u0000wy\u0001\u0000\u0000\u0000xv\u0001\u0000\u0000"+
		"\u0000yz\u00030\u0018\u0000z\u0080\u0003\u0004\u0002\u0000{~\u0005\b\u0000"+
		"\u0000|\u007f\u0003\u0004\u0002\u0000}\u007f\u0003\u0012\t\u0000~|\u0001"+
		"\u0000\u0000\u0000~}\u0001\u0000\u0000\u0000\u007f\u0081\u0001\u0000\u0000"+
		"\u0000\u0080{\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000\u0000"+
		"\u0081\u0013\u0001\u0000\u0000\u0000\u0082\u0083\u0005\t\u0000\u0000\u0083"+
		"\u0084\u00030\u0018\u0000\u0084\u0085\u0003\u0004\u0002\u0000\u0085\u0015"+
		"\u0001\u0000\u0000\u0000\u0086\u008a\u0005\n\u0000\u0000\u0087\u0089\u0003"+
		"\u0018\f\u0000\u0088\u0087\u0001\u0000\u0000\u0000\u0089\u008c\u0001\u0000"+
		"\u0000\u0000\u008a\u0088\u0001\u0000\u0000\u0000\u008a\u008b\u0001\u0000"+
		"\u0000\u0000\u008b\u008d\u0001\u0000\u0000\u0000\u008c\u008a\u0001\u0000"+
		"\u0000\u0000\u008d\u008e\u00030\u0018\u0000\u008e\u0098\u0005!\u0000\u0000"+
		"\u008f\u0092\u0005\u000b\u0000\u0000\u0090\u0093\u00038\u001c\u0000\u0091"+
		"\u0093\u00030\u0018\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0091"+
		"\u0001\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\u0095"+
		"\u0003\u0004\u0002\u0000\u0095\u0097\u0001\u0000\u0000\u0000\u0096\u008f"+
		"\u0001\u0000\u0000\u0000\u0097\u009a\u0001\u0000\u0000\u0000\u0098\u0096"+
		"\u0001\u0000\u0000\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u009d"+
		"\u0001\u0000\u0000\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009b\u009c"+
		"\u0005\f\u0000\u0000\u009c\u009e\u0003\u0004\u0002\u0000\u009d\u009b\u0001"+
		"\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000\u0000\u009e\u009f\u0001"+
		"\u0000\u0000\u0000\u009f\u00a0\u0005\"\u0000\u0000\u00a0\u0017\u0001\u0000"+
		"\u0000\u0000\u00a1\u00a4\u0003\u001c\u000e\u0000\u00a2\u00a4\u0003,\u0016"+
		"\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a2\u0001\u0000\u0000"+
		"\u0000\u00a4\u0019\u0001\u0000\u0000\u0000\u00a5\u00a6\u00038\u001c\u0000"+
		"\u00a6\u00a7\u00058\u0000\u0000\u00a7\u00a9\u0005\u001d\u0000\u0000\u00a8"+
		"\u00aa\u0003\u001e\u000f\u0000\u00a9\u00a8\u0001\u0000\u0000\u0000\u00a9"+
		"\u00aa\u0001\u0000\u0000\u0000\u00aa\u00ab\u0001\u0000\u0000\u0000\u00ab"+
		"\u00ad\u0005\u001e\u0000\u0000\u00ac\u00ae\u00038\u001c\u0000\u00ad\u00ac"+
		"\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae\u00af"+
		"\u0001\u0000\u0000\u0000\u00af\u00b0\u0003\u0004\u0002\u0000\u00b0\u001b"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b2\u0003\u001e\u000f\u0000\u00b2\u00b3"+
		"\u0005\u0006\u0000\u0000\u00b3\u00b4\u0003.\u0017\u0000\u00b4\u00b5\u0005"+
		"\u0005\u0000\u0000\u00b5\u00b8\u0001\u0000\u0000\u0000\u00b6\u00b8\u0003"+
		"*\u0015\u0000\u00b7\u00b1\u0001\u0000\u0000\u0000\u00b7\u00b6\u0001\u0000"+
		"\u0000\u0000\u00b8\u001d\u0001\u0000\u0000\u0000\u00b9\u00be\u0003 \u0010"+
		"\u0000\u00ba\u00bb\u0005\u0002\u0000\u0000\u00bb\u00bd\u0003\u001c\u000e"+
		"\u0000\u00bc\u00ba\u0001\u0000\u0000\u0000\u00bd\u00c0\u0001\u0000\u0000"+
		"\u0000\u00be\u00bc\u0001\u0000\u0000\u0000\u00be\u00bf\u0001\u0000\u0000"+
		"\u0000\u00bf\u001f\u0001\u0000\u0000\u0000\u00c0\u00be\u0001\u0000\u0000"+
		"\u0000\u00c1\u00cc\u0003$\u0012\u0000\u00c2\u00cc\u0003(\u0014\u0000\u00c3"+
		"\u00c4\u0005!\u0000\u0000\u00c4\u00c5\u0003\u001e\u000f\u0000\u00c5\u00c6"+
		"\u0005\"\u0000\u0000\u00c6\u00cc\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005"+
		"\u001f\u0000\u0000\u00c8\u00c9\u0003\u001e\u000f\u0000\u00c9\u00ca\u0005"+
		"\u001f\u0000\u0000\u00ca\u00cc\u0001\u0000\u0000\u0000\u00cb\u00c1\u0001"+
		"\u0000\u0000\u0000\u00cb\u00c2\u0001\u0000\u0000\u0000\u00cb\u00c3\u0001"+
		"\u0000\u0000\u0000\u00cb\u00c7\u0001\u0000\u0000\u0000\u00cc!\u0001\u0000"+
		"\u0000\u0000\u00cd\u00d2\u0003$\u0012\u0000\u00ce\u00cf\u0005\u0002\u0000"+
		"\u0000\u00cf\u00d1\u0003$\u0012\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000"+
		"\u00d1\u00d4\u0001\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000"+
		"\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3#\u0001\u0000\u0000\u0000\u00d4"+
		"\u00d2\u0001\u0000\u0000\u0000\u00d5\u00d6\u00038\u001c\u0000\u00d6\u00d7"+
		"\u0003(\u0014\u0000\u00d7\u00db\u0001\u0000\u0000\u0000\u00d8\u00d9\u0005"+
		"\u0015\u0000\u0000\u00d9\u00db\u0003(\u0014\u0000\u00da\u00d5\u0001\u0000"+
		"\u0000\u0000\u00da\u00d8\u0001\u0000\u0000\u0000\u00db%\u0001\u0000\u0000"+
		"\u0000\u00dc\u00e1\u0003(\u0014\u0000\u00dd\u00de\u0005\u0002\u0000\u0000"+
		"\u00de\u00e0\u0003(\u0014\u0000\u00df\u00dd\u0001\u0000\u0000\u0000\u00e0"+
		"\u00e3\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000\u00e1"+
		"\u00e2\u0001\u0000\u0000\u0000\u00e2\'\u0001\u0000\u0000\u0000\u00e3\u00e1"+
		"\u0001\u0000\u0000\u0000\u00e4\u00ee\u00038\u001c\u0000\u00e5\u00e6\u0005"+
		"!\u0000\u0000\u00e6\u00e7\u0003&\u0013\u0000\u00e7\u00e8\u0005\"\u0000"+
		"\u0000\u00e8\u00ee\u0001\u0000\u0000\u0000\u00e9\u00ea\u0005\u001d\u0000"+
		"\u0000\u00ea\u00eb\u0003&\u0013\u0000\u00eb\u00ec\u0005\u001e\u0000\u0000"+
		"\u00ec\u00ee\u0001\u0000\u0000\u0000\u00ed\u00e4\u0001\u0000\u0000\u0000"+
		"\u00ed\u00e5\u0001\u0000\u0000\u0000\u00ed\u00e9\u0001\u0000\u0000\u0000"+
		"\u00ee)\u0001\u0000\u0000\u0000\u00ef\u00f0\u00030\u0018\u0000\u00f0\u00f1"+
		"\u0007\u0000\u0000\u0000\u00f1\u00f2\u0005\u0006\u0000\u0000\u00f2\u00f3"+
		"\u0003.\u0017\u0000\u00f3\u00f4\u0005\u0005\u0000\u0000\u00f4+\u0001\u0000"+
		"\u0000\u0000\u00f5\u00f6\u0003.\u0017\u0000\u00f6\u00f7\u0005\u0005\u0000"+
		"\u0000\u00f7-\u0001\u0000\u0000\u0000\u00f8\u00fd\u00030\u0018\u0000\u00f9"+
		"\u00fa\u0005\u0002\u0000\u0000\u00fa\u00fc\u00030\u0018\u0000\u00fb\u00f9"+
		"\u0001\u0000\u0000\u0000\u00fc\u00ff\u0001\u0000\u0000\u0000\u00fd\u00fb"+
		"\u0001\u0000\u0000\u0000\u00fd\u00fe\u0001\u0000\u0000\u0000\u00fe/\u0001"+
		"\u0000\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000\u0100\u0101\u0006"+
		"\u0018\uffff\uffff\u0000\u0101\u0102\u0005\u001d\u0000\u0000\u0102\u0103"+
		"\u00038\u001c\u0000\u0103\u0104\u0005\u001e\u0000\u0000\u0104\u0105\u0003"+
		"0\u0018\u0014\u0105\u0112\u0001\u0000\u0000\u0000\u0106\u0107\u0005\u001d"+
		"\u0000\u0000\u0107\u0108\u00030\u0018\u0000\u0108\u0109\u0005\u001e\u0000"+
		"\u0000\u0109\u0112\u0001\u0000\u0000\u0000\u010a\u010b\u0007\u0001\u0000"+
		"\u0000\u010b\u0112\u00030\u0018\u0011\u010c\u0112\u00058\u0000\u0000\u010d"+
		"\u0112\u0003>\u001f\u0000\u010e\u0112\u00032\u0019\u0000\u010f\u0112\u0003"+
		"4\u001a\u0000\u0110\u0112\u00036\u001b\u0000\u0111\u0100\u0001\u0000\u0000"+
		"\u0000\u0111\u0106\u0001\u0000\u0000\u0000\u0111\u010a\u0001\u0000\u0000"+
		"\u0000\u0111\u010c\u0001\u0000\u0000\u0000\u0111\u010d\u0001\u0000\u0000"+
		"\u0000\u0111\u010e\u0001\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000"+
		"\u0000\u0111\u0110\u0001\u0000\u0000\u0000\u0112\u0145\u0001\u0000\u0000"+
		"\u0000\u0113\u0114\n\u0012\u0000\u0000\u0114\u0115\u0007\u0002\u0000\u0000"+
		"\u0115\u0144\u00030\u0018\u0013\u0116\u0117\n\u0010\u0000\u0000\u0117"+
		"\u0118\u0007\u0003\u0000\u0000\u0118\u0144\u00030\u0018\u0011\u0119\u011a"+
		"\n\u000f\u0000\u0000\u011a\u011b\u0007\u0004\u0000\u0000\u011b\u0144\u0003"+
		"0\u0018\u0010\u011c\u011d\n\u000e\u0000\u0000\u011d\u011e\u0007\u0005"+
		"\u0000\u0000\u011e\u0144\u00030\u0018\u000f\u011f\u0120\n\r\u0000\u0000"+
		"\u0120\u0121\u0007\u0006\u0000\u0000\u0121\u0144\u00030\u0018\u000e\u0122"+
		"\u0123\n\f\u0000\u0000\u0123\u0124\u0005-\u0000\u0000\u0124\u0144\u0003"+
		"0\u0018\r\u0125\u0126\n\u000b\u0000\u0000\u0126\u0127\u0005,\u0000\u0000"+
		"\u0127\u0144\u00030\u0018\f\u0128\u0129\n\n\u0000\u0000\u0129\u012a\u0005"+
		"+\u0000\u0000\u012a\u0144\u00030\u0018\u000b\u012b\u012c\n\t\u0000\u0000"+
		"\u012c\u012e\u0005\u001d\u0000\u0000\u012d\u012f\u0003.\u0017\u0000\u012e"+
		"\u012d\u0001\u0000\u0000\u0000\u012e\u012f\u0001\u0000\u0000\u0000\u012f"+
		"\u0130\u0001\u0000\u0000\u0000\u0130\u0144\u0005\u001e\u0000\u0000\u0131"+
		"\u0132\n\b\u0000\u0000\u0132\u0133\u0005\u0001\u0000\u0000\u0133\u0144"+
		"\u00058\u0000\u0000\u0134\u0135\n\u0007\u0000\u0000\u0135\u0136\u0005"+
		"\u001f\u0000\u0000\u0136\u0137\u00030\u0018\u0000\u0137\u0138\u0005 \u0000"+
		"\u0000\u0138\u0144\u0001\u0000\u0000\u0000\u0139\u013a\n\u0006\u0000\u0000"+
		"\u013a\u013c\u0005\u001f\u0000\u0000\u013b\u013d\u00030\u0018\u0000\u013c"+
		"\u013b\u0001\u0000\u0000\u0000\u013c\u013d\u0001\u0000\u0000\u0000\u013d"+
		"\u013e\u0001\u0000\u0000\u0000\u013e\u0140\u0005\u0003\u0000\u0000\u013f"+
		"\u0141\u00030\u0018\u0000\u0140\u013f\u0001\u0000\u0000\u0000\u0140\u0141"+
		"\u0001\u0000\u0000\u0000\u0141\u0142\u0001\u0000\u0000\u0000\u0142\u0144"+
		"\u0005 \u0000\u0000\u0143\u0113\u0001\u0000\u0000\u0000\u0143\u0116\u0001"+
		"\u0000\u0000\u0000\u0143\u0119\u0001\u0000\u0000\u0000\u0143\u011c\u0001"+
		"\u0000\u0000\u0000\u0143\u011f\u0001\u0000\u0000\u0000\u0143\u0122\u0001"+
		"\u0000\u0000\u0000\u0143\u0125\u0001\u0000\u0000\u0000\u0143\u0128\u0001"+
		"\u0000\u0000\u0000\u0143\u012b\u0001\u0000\u0000\u0000\u0143\u0131\u0001"+
		"\u0000\u0000\u0000\u0143\u0134\u0001\u0000\u0000\u0000\u0143\u0139\u0001"+
		"\u0000\u0000\u0000\u0144\u0147\u0001\u0000\u0000\u0000\u0145\u0143\u0001"+
		"\u0000\u0000\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u01461\u0001\u0000"+
		"\u0000\u0000\u0147\u0145\u0001\u0000\u0000\u0000\u0148\u0149\u00038\u001c"+
		"\u0000\u0149\u014b\u0005\u001d\u0000\u0000\u014a\u014c\u0003\u001e\u000f"+
		"\u0000\u014b\u014a\u0001\u0000\u0000\u0000\u014b\u014c\u0001\u0000\u0000"+
		"\u0000\u014c\u014d\u0001\u0000\u0000\u0000\u014d\u014f\u0005\u001e\u0000"+
		"\u0000\u014e\u0150\u00038\u001c\u0000\u014f\u014e\u0001\u0000\u0000\u0000"+
		"\u014f\u0150\u0001\u0000\u0000\u0000\u0150\u0151\u0001\u0000\u0000\u0000"+
		"\u0151\u0152\u0003\u0004\u0002\u0000\u01523\u0001\u0000\u0000\u0000\u0153"+
		"\u0154\u00038\u001c\u0000\u0154\u0155\u0005!\u0000\u0000\u0155\u0156\u0005"+
		"8\u0000\u0000\u0156\u0157\u0005\u0003\u0000\u0000\u0157\u015e\u00030\u0018"+
		"\u0000\u0158\u0159\u0005\u0002\u0000\u0000\u0159\u015a\u00058\u0000\u0000"+
		"\u015a\u015b\u0005\u0003\u0000\u0000\u015b\u015d\u00030\u0018\u0000\u015c"+
		"\u0158\u0001\u0000\u0000\u0000\u015d\u0160\u0001\u0000\u0000\u0000\u015e"+
		"\u015c\u0001\u0000\u0000\u0000\u015e\u015f\u0001\u0000\u0000\u0000\u015f"+
		"\u0162\u0001\u0000\u0000\u0000\u0160\u015e\u0001\u0000\u0000\u0000\u0161"+
		"\u0163\u0005\u0002\u0000\u0000\u0162\u0161\u0001\u0000\u0000\u0000\u0162"+
		"\u0163\u0001\u0000\u0000\u0000\u0163\u0164\u0001\u0000\u0000\u0000\u0164"+
		"\u0165\u0005\"\u0000\u0000\u01655\u0001\u0000\u0000\u0000\u0166\u0167"+
		"\u0005\u001d\u0000\u0000\u0167\u016f\u00030\u0018\u0000\u0168\u0170\u0005"+
		"\u0002\u0000\u0000\u0169\u016a\u0005\u0002\u0000\u0000\u016a\u016c\u0003"+
		"0\u0018\u0000\u016b\u0169\u0001\u0000\u0000\u0000\u016c\u016d\u0001\u0000"+
		"\u0000\u0000\u016d\u016b\u0001\u0000\u0000\u0000\u016d\u016e\u0001\u0000"+
		"\u0000\u0000\u016e\u0170\u0001\u0000\u0000\u0000\u016f\u0168\u0001\u0000"+
		"\u0000\u0000\u016f\u016b\u0001\u0000\u0000\u0000\u0170\u0171\u0001\u0000"+
		"\u0000\u0000\u0171\u0172\u0005\u001e\u0000\u0000\u01727\u0001\u0000\u0000"+
		"\u0000\u0173\u0178\u0003:\u001d\u0000\u0174\u0178\u0005\u0018\u0000\u0000"+
		"\u0175\u0178\u0005\u0019\u0000\u0000\u0176\u0178\u0005\u001a\u0000\u0000"+
		"\u0177\u0173\u0001\u0000\u0000\u0000\u0177\u0174\u0001\u0000\u0000\u0000"+
		"\u0177\u0175\u0001\u0000\u0000\u0000\u0177\u0176\u0001\u0000\u0000\u0000"+
		"\u01789\u0001\u0000\u0000\u0000\u0179\u017a\u0006\u001d\uffff\uffff\u0000"+
		"\u017a\u017b\u00058\u0000\u0000\u017b\u0181\u0001\u0000\u0000\u0000\u017c"+
		"\u017d\n\u0001\u0000\u0000\u017d\u017e\u0005\u0001\u0000\u0000\u017e\u0180"+
		"\u00058\u0000\u0000\u017f\u017c\u0001\u0000\u0000\u0000\u0180\u0183\u0001"+
		"\u0000\u0000\u0000\u0181\u017f\u0001\u0000\u0000\u0000\u0181\u0182\u0001"+
		"\u0000\u0000\u0000\u0182;\u0001\u0000\u0000\u0000\u0183\u0181\u0001\u0000"+
		"\u0000\u0000\u0184\u0189\u00058\u0000\u0000\u0185\u0186\u0005\u0002\u0000"+
		"\u0000\u0186\u0188\u00058\u0000\u0000\u0187\u0185\u0001\u0000\u0000\u0000"+
		"\u0188\u018b\u0001\u0000\u0000\u0000\u0189\u0187\u0001\u0000\u0000\u0000"+
		"\u0189\u018a\u0001\u0000\u0000\u0000\u018a=\u0001\u0000\u0000\u0000\u018b"+
		"\u0189\u0001\u0000\u0000\u0000\u018c\u018f\u0003@ \u0000\u018d\u018f\u0003"+
		"B!\u0000\u018e\u018c\u0001\u0000\u0000\u0000\u018e\u018d\u0001\u0000\u0000"+
		"\u0000\u018f?\u0001\u0000\u0000\u0000\u0190\u0191\u00059\u0000\u0000\u0191"+
		"A\u0001\u0000\u0000\u0000\u0192\u0193\u00058\u0000\u0000\u0193\u0194\u0005"+
		"9\u0000\u0000\u0194C\u0001\u0000\u0000\u0000(GPY^bpv~\u0080\u008a\u0092"+
		"\u0098\u009d\u00a3\u00a9\u00ad\u00b7\u00be\u00cb\u00d2\u00da\u00e1\u00ed"+
		"\u00fd\u0111\u012e\u013c\u0140\u0143\u0145\u014b\u014f\u015e\u0162\u016d"+
		"\u016f\u0177\u0181\u0189\u018e";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
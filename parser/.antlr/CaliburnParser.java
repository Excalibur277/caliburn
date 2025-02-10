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
		PERIOD=1, COMMA=2, COLON=3, SEMICOLON=4, QUESTION=5, ASSIGN=6, IF=7, ELSE=8, 
		FOR=9, SWITCH=10, CASE=11, DEFAULT=12, FALLTHROUGH=13, BREAK=14, CONTINUE=15, 
		RETURN=16, USING=17, AS=18, IN=19, NULL=20, VAR=21, CONST=22, TYPE=23, 
		FUNC=24, STRUCT=25, TUPLE=26, CLASS=27, INTERFACE=28, L_PAREN=29, R_PAREN=30, 
		L_S_BRACK=31, R_S_BRACK=32, L_C_BRACK=33, R_C_BRACK=34, OP_ADD=35, OP_SUB=36, 
		OP_NOT=37, OP_MUL=38, OP_DIV=39, OP_MOD=40, OP_POW=41, OP_ROOT=42, OP_OR=43, 
		OP_XOR=44, OP_AND=45, OP_LSHIFT=46, OP_RSHIFT=47, OP_EQU=48, OP_NEQ=49, 
		OP_GTE=50, OP_GRT=51, OP_LTE=52, OP_LST=53, WhiteSpace=54, Comment=55, 
		Identifier=56, Literal=57, Terminator=58;
	public static final int
		RULE_scoped_block = 0, RULE_statement = 1, RULE_control_statement = 2, 
		RULE_if_statement = 3, RULE_for_statement = 4, RULE_switch_statement = 5, 
		RULE_inline_statement = 6, RULE_function_declaration_statement = 7, RULE_assign_statement = 8, 
		RULE_assign_declarations = 9, RULE_assign_declaration = 10, RULE_declared_assign_declarations = 11, 
		RULE_declared_assign_declaration = 12, RULE_undeclared_assign_declarations = 13, 
		RULE_undeclared_assign_declaration = 14, RULE_operator_assign_statement = 15, 
		RULE_expression_statement = 16, RULE_expressions = 17, RULE_expression = 18, 
		RULE_function_expression = 19, RULE_struct_expression = 20, RULE_tuple_expression = 21, 
		RULE_expression_atom = 22, RULE_identifiers = 23, RULE_literal_atom = 24, 
		RULE_untyped_literal_atom = 25, RULE_typed_literal_atom = 26;
	private static String[] makeRuleNames() {
		return new String[] {
			"scoped_block", "statement", "control_statement", "if_statement", "for_statement", 
			"switch_statement", "inline_statement", "function_declaration_statement", 
			"assign_statement", "assign_declarations", "assign_declaration", "declared_assign_declarations", 
			"declared_assign_declaration", "undeclared_assign_declarations", "undeclared_assign_declaration", 
			"operator_assign_statement", "expression_statement", "expressions", "expression", 
			"function_expression", "struct_expression", "tuple_expression", "expression_atom", 
			"identifiers", "literal_atom", "untyped_literal_atom", "typed_literal_atom"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "','", "':'", "';'", "'?'", "'='", "'if'", "'else'", "'for'", 
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
			null, "PERIOD", "COMMA", "COLON", "SEMICOLON", "QUESTION", "ASSIGN", 
			"IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", "BREAK", 
			"CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", "TYPE", 
			"FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN", 
			"L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD", "OP_SUB", 
			"OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", "OP_OR", 
			"OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ", "OP_GTE", 
			"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier", 
			"Literal", "Terminator"
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
		enterRule(_localctx, 0, RULE_scoped_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(L_C_BRACK);
			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173033958672000L) != 0)) {
				{
				{
				setState(55);
				statement();
				}
				}
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(61);
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
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_statement);
		try {
			setState(66);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(63);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(64);
				expression_statement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(65);
				control_statement();
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
		enterRule(_localctx, 4, RULE_control_statement);
		try {
			setState(72);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(68);
				if_statement();
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				for_statement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 3);
				{
				setState(70);
				switch_statement();
				}
				break;
			case FUNC:
			case Identifier:
				enterOuterAlt(_localctx, 4);
				{
				setState(71);
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
		enterRule(_localctx, 6, RULE_if_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(74);
			match(IF);
			setState(78);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(75);
					inline_statement();
					}
					} 
				}
				setState(80);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			}
			setState(81);
			expression(0);
			setState(82);
			scoped_block();
			setState(88);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(83);
				match(ELSE);
				setState(86);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case L_C_BRACK:
					{
					setState(84);
					scoped_block();
					}
					break;
				case IF:
					{
					setState(85);
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
		enterRule(_localctx, 8, RULE_for_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			match(FOR);
			setState(91);
			expression(0);
			setState(92);
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
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
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
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public List<Literal_atomContext> literal_atom() {
			return getRuleContexts(Literal_atomContext.class);
		}
		public Literal_atomContext literal_atom(int i) {
			return getRuleContext(Literal_atomContext.class,i);
		}
		public Switch_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_statement; }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_switch_statement);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(94);
			match(SWITCH);
			setState(98);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(95);
					inline_statement();
					}
					} 
				}
				setState(100);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			setState(101);
			expression(0);
			setState(102);
			match(L_C_BRACK);
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(103);
				match(CASE);
				setState(106);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
				case 1:
					{
					setState(104);
					match(Identifier);
					}
					break;
				case 2:
					{
					setState(105);
					literal_atom();
					}
					break;
				}
				setState(108);
				scoped_block();
				}
				}
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(116);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(114);
				match(DEFAULT);
				setState(115);
				scoped_block();
				}
			}

			setState(118);
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
		enterRule(_localctx, 12, RULE_inline_statement);
		try {
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(120);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(121);
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
		public Token type;
		public Token name;
		public Token return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public Scoped_blockContext scoped_block() {
			return getRuleContext(Scoped_blockContext.class,0);
		}
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public TerminalNode FUNC() { return getToken(CaliburnParser.FUNC, 0); }
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode VAR() { return getToken(CaliburnParser.VAR, 0); }
		public Function_declaration_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_declaration_statement; }
	}

	public final Function_declaration_statementContext function_declaration_statement() throws RecognitionException {
		Function_declaration_statementContext _localctx = new Function_declaration_statementContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_function_declaration_statement);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			((Function_declaration_statementContext)_localctx).type = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==FUNC || _la==Identifier) ) {
				((Function_declaration_statementContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(125);
			((Function_declaration_statementContext)_localctx).name = match(Identifier);
			setState(126);
			match(L_PAREN);
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057605314314240L) != 0)) {
				{
				setState(127);
				assign_declarations();
				}
			}

			setState(130);
			match(R_PAREN);
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VAR || _la==Identifier) {
				{
				setState(131);
				((Function_declaration_statementContext)_localctx).return_type = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==Identifier) ) {
					((Function_declaration_statementContext)_localctx).return_type = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(134);
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
		enterRule(_localctx, 16, RULE_assign_statement);
		try {
			setState(142);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(136);
				assign_declarations();
				setState(137);
				match(ASSIGN);
				setState(138);
				expressions();
				setState(139);
				match(Terminator);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(141);
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
		enterRule(_localctx, 18, RULE_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			assign_declaration();
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(145);
				match(COMMA);
				setState(146);
				assign_statement();
				}
				}
				setState(151);
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
		enterRule(_localctx, 20, RULE_assign_declaration);
		try {
			setState(162);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				_localctx = new DeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(152);
				declared_assign_declaration();
				}
				break;
			case 2:
				_localctx = new UndeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(153);
				undeclared_assign_declaration();
				}
				break;
			case 3:
				_localctx = new StructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(154);
				match(L_C_BRACK);
				setState(155);
				assign_declarations();
				setState(156);
				match(R_C_BRACK);
				}
				break;
			case 4:
				_localctx = new TupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(158);
				match(L_S_BRACK);
				setState(159);
				assign_declarations();
				setState(160);
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
		enterRule(_localctx, 22, RULE_declared_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			declared_assign_declaration();
			setState(169);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(165);
				match(COMMA);
				setState(166);
				declared_assign_declaration();
				}
				}
				setState(171);
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
		public Token type;
		public Undeclared_assign_declarationContext undeclared_assign_declaration() {
			return getRuleContext(Undeclared_assign_declarationContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public TypedDeclaredAssignDeclarationContext(Declared_assign_declarationContext ctx) { copyFrom(ctx); }
	}

	public final Declared_assign_declarationContext declared_assign_declaration() throws RecognitionException {
		Declared_assign_declarationContext _localctx = new Declared_assign_declarationContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_declared_assign_declaration);
		try {
			setState(176);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				_localctx = new TypedDeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(172);
				((TypedDeclaredAssignDeclarationContext)_localctx).type = match(Identifier);
				setState(173);
				undeclared_assign_declaration();
				}
				break;
			case VAR:
				_localctx = new UntypedDeclaredAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				match(VAR);
				setState(175);
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
		enterRule(_localctx, 26, RULE_undeclared_assign_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(178);
			undeclared_assign_declaration();
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(179);
				match(COMMA);
				setState(180);
				undeclared_assign_declaration();
				}
				}
				setState(185);
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
		public Token var;
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
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
		enterRule(_localctx, 28, RULE_undeclared_assign_declaration);
		try {
			setState(195);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				_localctx = new UndeclaredAtomAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(186);
				((UndeclaredAtomAssignDeclarationContext)_localctx).var = match(Identifier);
				}
				break;
			case L_C_BRACK:
				_localctx = new UndeclaredStructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(187);
				match(L_C_BRACK);
				setState(188);
				undeclared_assign_declarations();
				setState(189);
				match(R_C_BRACK);
				}
				break;
			case L_PAREN:
				_localctx = new UndeclaredTupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(191);
				match(L_PAREN);
				setState(192);
				undeclared_assign_declarations();
				setState(193);
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
		public Token var;
		public Token op;
		public ExpressionsContext expressions() {
			return getRuleContext(ExpressionsContext.class,0);
		}
		public TerminalNode Terminator() { return getToken(CaliburnParser.Terminator, 0); }
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
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
		enterRule(_localctx, 30, RULE_operator_assign_statement);
		int _la;
		try {
			_localctx = new OperatorAssignmentContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			((OperatorAssignmentContext)_localctx).var = match(Identifier);
			setState(198);
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
			setState(199);
			expressions();
			setState(200);
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
		enterRule(_localctx, 32, RULE_expression_statement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			expressions();
			setState(203);
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
		enterRule(_localctx, 34, RULE_expressions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			expression(0);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(206);
				match(COMMA);
				setState(207);
				expression(0);
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
	public static class ExponentExpressionContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_POW() { return getToken(CaliburnParser.OP_POW, 0); }
		public TerminalNode OP_ROOT() { return getToken(CaliburnParser.OP_ROOT, 0); }
		public ExponentExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AdditionExpressionContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_ADD() { return getToken(CaliburnParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CaliburnParser.OP_SUB, 0); }
		public AdditionExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BracketedExpressionContext extends ExpressionContext {
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public BracketedExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceExpressionContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public TerminalNode COLON() { return getToken(CaliburnParser.COLON, 0); }
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public SliceExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TupleExpressionContext extends ExpressionContext {
		public Tuple_expressionContext tuple_expression() {
			return getRuleContext(Tuple_expressionContext.class,0);
		}
		public TupleExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexExpressionContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode L_S_BRACK() { return getToken(CaliburnParser.L_S_BRACK, 0); }
		public TerminalNode R_S_BRACK() { return getToken(CaliburnParser.R_S_BRACK, 0); }
		public IndexExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnaryExpressionContext extends ExpressionContext {
		public Token op;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OP_NOT() { return getToken(CaliburnParser.OP_NOT, 0); }
		public TerminalNode OP_ADD() { return getToken(CaliburnParser.OP_ADD, 0); }
		public TerminalNode OP_SUB() { return getToken(CaliburnParser.OP_SUB, 0); }
		public UnaryExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AtomExpressionContext extends ExpressionContext {
		public Expression_atomContext expression_atom() {
			return getRuleContext(Expression_atomContext.class,0);
		}
		public AtomExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ShiftExpressionContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_LSHIFT() { return getToken(CaliburnParser.OP_LSHIFT, 0); }
		public TerminalNode OP_RSHIFT() { return getToken(CaliburnParser.OP_RSHIFT, 0); }
		public ShiftExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrExpressionContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_OR() { return getToken(CaliburnParser.OP_OR, 0); }
		public OrExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AccessExpressionContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PERIOD() { return getToken(CaliburnParser.PERIOD, 0); }
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public AccessExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ComparisonExpressionContext extends ExpressionContext {
		public Token op;
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
		public ComparisonExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class XOrExpressionContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_XOR() { return getToken(CaliburnParser.OP_XOR, 0); }
		public XOrExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionExpressionContext extends ExpressionContext {
		public Function_expressionContext function_expression() {
			return getRuleContext(Function_expressionContext.class,0);
		}
		public FunctionExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructExpressionContext extends ExpressionContext {
		public Struct_expressionContext struct_expression() {
			return getRuleContext(Struct_expressionContext.class,0);
		}
		public StructExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AndExpressionContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_AND() { return getToken(CaliburnParser.OP_AND, 0); }
		public AndExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CastExpressionContext extends ExpressionContext {
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public CastExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MultiplicativeExpressionContext extends ExpressionContext {
		public Token op;
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode OP_MUL() { return getToken(CaliburnParser.OP_MUL, 0); }
		public TerminalNode OP_DIV() { return getToken(CaliburnParser.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(CaliburnParser.OP_MOD, 0); }
		public MultiplicativeExpressionContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallExpressionContext extends ExpressionContext {
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				{
				_localctx = new CastExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(214);
				match(L_PAREN);
				setState(215);
				expression(0);
				setState(216);
				match(R_PAREN);
				setState(217);
				expression(19);
				}
				break;
			case 2:
				{
				_localctx = new BracketedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(219);
				match(L_PAREN);
				setState(220);
				expression(0);
				setState(221);
				match(R_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(223);
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
				setState(224);
				expression(16);
				}
				break;
			case 4:
				{
				_localctx = new AtomExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(225);
				expression_atom();
				}
				break;
			case 5:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(226);
				function_expression();
				}
				break;
			case 6:
				{
				_localctx = new StructExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(227);
				struct_expression();
				}
				break;
			case 7:
				{
				_localctx = new TupleExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(228);
				tuple_expression();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(281);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(279);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
					case 1:
						{
						_localctx = new ExponentExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(231);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(232);
						((ExponentExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_POW || _la==OP_ROOT) ) {
							((ExponentExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(233);
						expression(18);
						}
						break;
					case 2:
						{
						_localctx = new MultiplicativeExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(234);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(235);
						((MultiplicativeExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1924145348608L) != 0)) ) {
							((MultiplicativeExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(236);
						expression(16);
						}
						break;
					case 3:
						{
						_localctx = new AdditionExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(237);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(238);
						((AdditionExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_ADD || _la==OP_SUB) ) {
							((AdditionExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(239);
						expression(15);
						}
						break;
					case 4:
						{
						_localctx = new ShiftExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(240);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(241);
						((ShiftExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_LSHIFT || _la==OP_RSHIFT) ) {
							((ShiftExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(242);
						expression(14);
						}
						break;
					case 5:
						{
						_localctx = new ComparisonExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(243);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(244);
						((ComparisonExpressionContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 17732923532771328L) != 0)) ) {
							((ComparisonExpressionContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(245);
						expression(13);
						}
						break;
					case 6:
						{
						_localctx = new AndExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(246);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(247);
						match(OP_AND);
						setState(248);
						expression(12);
						}
						break;
					case 7:
						{
						_localctx = new XOrExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(249);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(250);
						match(OP_XOR);
						setState(251);
						expression(11);
						}
						break;
					case 8:
						{
						_localctx = new OrExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(252);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(253);
						match(OP_OR);
						setState(254);
						expression(10);
						}
						break;
					case 9:
						{
						_localctx = new CallExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(255);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(256);
						match(L_PAREN);
						setState(258);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023219154944L) != 0)) {
							{
							setState(257);
							expressions();
							}
						}

						setState(260);
						match(R_PAREN);
						}
						break;
					case 10:
						{
						_localctx = new AccessExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(261);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(262);
						match(PERIOD);
						setState(263);
						match(Identifier);
						}
						break;
					case 11:
						{
						_localctx = new IndexExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(264);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(265);
						match(L_S_BRACK);
						setState(266);
						expression(0);
						setState(267);
						match(R_S_BRACK);
						}
						break;
					case 12:
						{
						_localctx = new SliceExpressionContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(269);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(270);
						match(L_S_BRACK);
						setState(272);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023219154944L) != 0)) {
							{
							setState(271);
							expression(0);
							}
						}

						setState(274);
						match(COLON);
						setState(276);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 216173023219154944L) != 0)) {
							{
							setState(275);
							expression(0);
							}
						}

						setState(278);
						match(R_S_BRACK);
						}
						break;
					}
					} 
				}
				setState(283);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
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
		public Token type;
		public Token return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public Scoped_blockContext scoped_block() {
			return getRuleContext(Scoped_blockContext.class,0);
		}
		public Assign_declarationsContext assign_declarations() {
			return getRuleContext(Assign_declarationsContext.class,0);
		}
		public TerminalNode FUNC() { return getToken(CaliburnParser.FUNC, 0); }
		public List<TerminalNode> Identifier() { return getTokens(CaliburnParser.Identifier); }
		public TerminalNode Identifier(int i) {
			return getToken(CaliburnParser.Identifier, i);
		}
		public TerminalNode VAR() { return getToken(CaliburnParser.VAR, 0); }
		public Function_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function_expression; }
	}

	public final Function_expressionContext function_expression() throws RecognitionException {
		Function_expressionContext _localctx = new Function_expressionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_function_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(284);
			((Function_expressionContext)_localctx).type = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==FUNC || _la==Identifier) ) {
				((Function_expressionContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
			setState(285);
			match(L_PAREN);
			setState(287);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 72057605314314240L) != 0)) {
				{
				setState(286);
				assign_declarations();
				}
			}

			setState(289);
			match(R_PAREN);
			setState(291);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VAR || _la==Identifier) {
				{
				setState(290);
				((Function_expressionContext)_localctx).return_type = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==Identifier) ) {
					((Function_expressionContext)_localctx).return_type = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(293);
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
		public Token type;
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
		public TerminalNode STRUCT() { return getToken(CaliburnParser.STRUCT, 0); }
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
		enterRule(_localctx, 40, RULE_struct_expression);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			((Struct_expressionContext)_localctx).type = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==STRUCT || _la==Identifier) ) {
				((Struct_expressionContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(296);
			match(L_C_BRACK);
			{
			setState(297);
			match(Identifier);
			}
			setState(298);
			match(COLON);
			setState(299);
			expression(0);
			setState(306);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(300);
					match(COMMA);
					{
					setState(301);
					match(Identifier);
					}
					setState(302);
					match(COLON);
					setState(303);
					expression(0);
					}
					} 
				}
				setState(308);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			}
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMMA) {
				{
				setState(309);
				match(COMMA);
				}
			}

			setState(312);
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
		enterRule(_localctx, 42, RULE_tuple_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(314);
			match(L_PAREN);
			setState(315);
			expression(0);
			setState(323);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				setState(316);
				match(COMMA);
				}
				break;
			case 2:
				{
				setState(319); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(317);
					match(COMMA);
					setState(318);
					expression(0);
					}
					}
					setState(321); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==COMMA );
				}
				break;
			}
			setState(325);
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
	public static class Expression_atomContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Literal_atomContext literal_atom() {
			return getRuleContext(Literal_atomContext.class,0);
		}
		public Expression_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_atom; }
	}

	public final Expression_atomContext expression_atom() throws RecognitionException {
		Expression_atomContext _localctx = new Expression_atomContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_expression_atom);
		try {
			setState(329);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(327);
				match(Identifier);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(328);
				literal_atom();
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
		enterRule(_localctx, 46, RULE_identifiers);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(331);
			match(Identifier);
			setState(336);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(332);
				match(COMMA);
				setState(333);
				match(Identifier);
				}
				}
				setState(338);
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
		public Untyped_literal_atomContext untyped_literal_atom() {
			return getRuleContext(Untyped_literal_atomContext.class,0);
		}
		public Typed_literal_atomContext typed_literal_atom() {
			return getRuleContext(Typed_literal_atomContext.class,0);
		}
		public Literal_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal_atom; }
	}

	public final Literal_atomContext literal_atom() throws RecognitionException {
		Literal_atomContext _localctx = new Literal_atomContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_literal_atom);
		try {
			setState(341);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Literal:
				enterOuterAlt(_localctx, 1);
				{
				setState(339);
				untyped_literal_atom();
				}
				break;
			case Identifier:
				enterOuterAlt(_localctx, 2);
				{
				setState(340);
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
		enterRule(_localctx, 50, RULE_untyped_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(343);
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
		enterRule(_localctx, 52, RULE_typed_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(Identifier);
			setState(346);
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
		case 18:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 15);
		case 2:
			return precpred(_ctx, 14);
		case 3:
			return precpred(_ctx, 13);
		case 4:
			return precpred(_ctx, 12);
		case 5:
			return precpred(_ctx, 11);
		case 6:
			return precpred(_ctx, 10);
		case 7:
			return precpred(_ctx, 9);
		case 8:
			return precpred(_ctx, 8);
		case 9:
			return precpred(_ctx, 7);
		case 10:
			return precpred(_ctx, 6);
		case 11:
			return precpred(_ctx, 5);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001:\u015d\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0001\u0000\u0001\u0000"+
		"\u0005\u00009\b\u0000\n\u0000\f\u0000<\t\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001C\b\u0001\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002I\b\u0002\u0001\u0003"+
		"\u0001\u0003\u0005\u0003M\b\u0003\n\u0003\f\u0003P\t\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003W\b\u0003"+
		"\u0003\u0003Y\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0005\u0005a\b\u0005\n\u0005\f\u0005d\t\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005"+
		"k\b\u0005\u0001\u0005\u0005\u0005n\b\u0005\n\u0005\f\u0005q\t\u0005\u0001"+
		"\u0005\u0001\u0005\u0003\u0005u\b\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0003\u0006{\b\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0003\u0007\u0081\b\u0007\u0001\u0007\u0001\u0007\u0003"+
		"\u0007\u0085\b\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0003\b\u008f\b\b\u0001\t\u0001\t\u0001\t\u0005\t\u0094"+
		"\b\t\n\t\f\t\u0097\t\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n"+
		"\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u00a3\b\n\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0005\u000b\u00a8\b\u000b\n\u000b\f\u000b\u00ab\t\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0003\f\u00b1\b\f\u0001\r\u0001\r\u0001\r\u0005"+
		"\r\u00b6\b\r\n\r\f\r\u00b9\t\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003"+
		"\u000e\u00c4\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0005\u0011\u00d1\b\u0011\n\u0011\f\u0011\u00d4\t\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u00e6\b\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0003\u0012\u0103\b\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u0111\b\u0012"+
		"\u0001\u0012\u0001\u0012\u0003\u0012\u0115\b\u0012\u0001\u0012\u0005\u0012"+
		"\u0118\b\u0012\n\u0012\f\u0012\u011b\t\u0012\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0003\u0013\u0120\b\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u0124"+
		"\b\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005"+
		"\u0014\u0131\b\u0014\n\u0014\f\u0014\u0134\t\u0014\u0001\u0014\u0003\u0014"+
		"\u0137\b\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0001\u0015\u0004\u0015\u0140\b\u0015\u000b\u0015\f\u0015"+
		"\u0141\u0003\u0015\u0144\b\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u014a\b\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0005"+
		"\u0017\u014f\b\u0017\n\u0017\f\u0017\u0152\t\u0017\u0001\u0018\u0001\u0018"+
		"\u0003\u0018\u0156\b\u0018\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0000\u0001$\u001b\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02"+
		"4\u0000\n\u0002\u0000\u0018\u001888\u0002\u0000\u0015\u001588\u0002\u0000"+
		"#$&/\u0001\u0000#%\u0001\u0000)*\u0001\u0000&(\u0001\u0000#$\u0001\u0000"+
		"./\u0001\u000005\u0002\u0000\u0019\u001988\u017a\u00006\u0001\u0000\u0000"+
		"\u0000\u0002B\u0001\u0000\u0000\u0000\u0004H\u0001\u0000\u0000\u0000\u0006"+
		"J\u0001\u0000\u0000\u0000\bZ\u0001\u0000\u0000\u0000\n^\u0001\u0000\u0000"+
		"\u0000\fz\u0001\u0000\u0000\u0000\u000e|\u0001\u0000\u0000\u0000\u0010"+
		"\u008e\u0001\u0000\u0000\u0000\u0012\u0090\u0001\u0000\u0000\u0000\u0014"+
		"\u00a2\u0001\u0000\u0000\u0000\u0016\u00a4\u0001\u0000\u0000\u0000\u0018"+
		"\u00b0\u0001\u0000\u0000\u0000\u001a\u00b2\u0001\u0000\u0000\u0000\u001c"+
		"\u00c3\u0001\u0000\u0000\u0000\u001e\u00c5\u0001\u0000\u0000\u0000 \u00ca"+
		"\u0001\u0000\u0000\u0000\"\u00cd\u0001\u0000\u0000\u0000$\u00e5\u0001"+
		"\u0000\u0000\u0000&\u011c\u0001\u0000\u0000\u0000(\u0127\u0001\u0000\u0000"+
		"\u0000*\u013a\u0001\u0000\u0000\u0000,\u0149\u0001\u0000\u0000\u0000."+
		"\u014b\u0001\u0000\u0000\u00000\u0155\u0001\u0000\u0000\u00002\u0157\u0001"+
		"\u0000\u0000\u00004\u0159\u0001\u0000\u0000\u00006:\u0005!\u0000\u0000"+
		"79\u0003\u0002\u0001\u000087\u0001\u0000\u0000\u00009<\u0001\u0000\u0000"+
		"\u0000:8\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000;=\u0001\u0000"+
		"\u0000\u0000<:\u0001\u0000\u0000\u0000=>\u0005\"\u0000\u0000>\u0001\u0001"+
		"\u0000\u0000\u0000?C\u0003\u0010\b\u0000@C\u0003 \u0010\u0000AC\u0003"+
		"\u0004\u0002\u0000B?\u0001\u0000\u0000\u0000B@\u0001\u0000\u0000\u0000"+
		"BA\u0001\u0000\u0000\u0000C\u0003\u0001\u0000\u0000\u0000DI\u0003\u0006"+
		"\u0003\u0000EI\u0003\b\u0004\u0000FI\u0003\n\u0005\u0000GI\u0003\u000e"+
		"\u0007\u0000HD\u0001\u0000\u0000\u0000HE\u0001\u0000\u0000\u0000HF\u0001"+
		"\u0000\u0000\u0000HG\u0001\u0000\u0000\u0000I\u0005\u0001\u0000\u0000"+
		"\u0000JN\u0005\u0007\u0000\u0000KM\u0003\f\u0006\u0000LK\u0001\u0000\u0000"+
		"\u0000MP\u0001\u0000\u0000\u0000NL\u0001\u0000\u0000\u0000NO\u0001\u0000"+
		"\u0000\u0000OQ\u0001\u0000\u0000\u0000PN\u0001\u0000\u0000\u0000QR\u0003"+
		"$\u0012\u0000RX\u0003\u0000\u0000\u0000SV\u0005\b\u0000\u0000TW\u0003"+
		"\u0000\u0000\u0000UW\u0003\u0006\u0003\u0000VT\u0001\u0000\u0000\u0000"+
		"VU\u0001\u0000\u0000\u0000WY\u0001\u0000\u0000\u0000XS\u0001\u0000\u0000"+
		"\u0000XY\u0001\u0000\u0000\u0000Y\u0007\u0001\u0000\u0000\u0000Z[\u0005"+
		"\t\u0000\u0000[\\\u0003$\u0012\u0000\\]\u0003\u0000\u0000\u0000]\t\u0001"+
		"\u0000\u0000\u0000^b\u0005\n\u0000\u0000_a\u0003\f\u0006\u0000`_\u0001"+
		"\u0000\u0000\u0000ad\u0001\u0000\u0000\u0000b`\u0001\u0000\u0000\u0000"+
		"bc\u0001\u0000\u0000\u0000ce\u0001\u0000\u0000\u0000db\u0001\u0000\u0000"+
		"\u0000ef\u0003$\u0012\u0000fo\u0005!\u0000\u0000gj\u0005\u000b\u0000\u0000"+
		"hk\u00058\u0000\u0000ik\u00030\u0018\u0000jh\u0001\u0000\u0000\u0000j"+
		"i\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000ln\u0003\u0000\u0000"+
		"\u0000mg\u0001\u0000\u0000\u0000nq\u0001\u0000\u0000\u0000om\u0001\u0000"+
		"\u0000\u0000op\u0001\u0000\u0000\u0000pt\u0001\u0000\u0000\u0000qo\u0001"+
		"\u0000\u0000\u0000rs\u0005\f\u0000\u0000su\u0003\u0000\u0000\u0000tr\u0001"+
		"\u0000\u0000\u0000tu\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000"+
		"vw\u0005\"\u0000\u0000w\u000b\u0001\u0000\u0000\u0000x{\u0003\u0010\b"+
		"\u0000y{\u0003 \u0010\u0000zx\u0001\u0000\u0000\u0000zy\u0001\u0000\u0000"+
		"\u0000{\r\u0001\u0000\u0000\u0000|}\u0007\u0000\u0000\u0000}~\u00058\u0000"+
		"\u0000~\u0080\u0005\u001d\u0000\u0000\u007f\u0081\u0003\u0012\t\u0000"+
		"\u0080\u007f\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000\u0000"+
		"\u0081\u0082\u0001\u0000\u0000\u0000\u0082\u0084\u0005\u001e\u0000\u0000"+
		"\u0083\u0085\u0007\u0001\u0000\u0000\u0084\u0083\u0001\u0000\u0000\u0000"+
		"\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000"+
		"\u0086\u0087\u0003\u0000\u0000\u0000\u0087\u000f\u0001\u0000\u0000\u0000"+
		"\u0088\u0089\u0003\u0012\t\u0000\u0089\u008a\u0005\u0006\u0000\u0000\u008a"+
		"\u008b\u0003\"\u0011\u0000\u008b\u008c\u0005:\u0000\u0000\u008c\u008f"+
		"\u0001\u0000\u0000\u0000\u008d\u008f\u0003\u001e\u000f\u0000\u008e\u0088"+
		"\u0001\u0000\u0000\u0000\u008e\u008d\u0001\u0000\u0000\u0000\u008f\u0011"+
		"\u0001\u0000\u0000\u0000\u0090\u0095\u0003\u0014\n\u0000\u0091\u0092\u0005"+
		"\u0002\u0000\u0000\u0092\u0094\u0003\u0010\b\u0000\u0093\u0091\u0001\u0000"+
		"\u0000\u0000\u0094\u0097\u0001\u0000\u0000\u0000\u0095\u0093\u0001\u0000"+
		"\u0000\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u0013\u0001\u0000"+
		"\u0000\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0098\u00a3\u0003\u0018"+
		"\f\u0000\u0099\u00a3\u0003\u001c\u000e\u0000\u009a\u009b\u0005!\u0000"+
		"\u0000\u009b\u009c\u0003\u0012\t\u0000\u009c\u009d\u0005\"\u0000\u0000"+
		"\u009d\u00a3\u0001\u0000\u0000\u0000\u009e\u009f\u0005\u001f\u0000\u0000"+
		"\u009f\u00a0\u0003\u0012\t\u0000\u00a0\u00a1\u0005\u001f\u0000\u0000\u00a1"+
		"\u00a3\u0001\u0000\u0000\u0000\u00a2\u0098\u0001\u0000\u0000\u0000\u00a2"+
		"\u0099\u0001\u0000\u0000\u0000\u00a2\u009a\u0001\u0000\u0000\u0000\u00a2"+
		"\u009e\u0001\u0000\u0000\u0000\u00a3\u0015\u0001\u0000\u0000\u0000\u00a4"+
		"\u00a9\u0003\u0018\f\u0000\u00a5\u00a6\u0005\u0002\u0000\u0000\u00a6\u00a8"+
		"\u0003\u0018\f\u0000\u00a7\u00a5\u0001\u0000\u0000\u0000\u00a8\u00ab\u0001"+
		"\u0000\u0000\u0000\u00a9\u00a7\u0001\u0000\u0000\u0000\u00a9\u00aa\u0001"+
		"\u0000\u0000\u0000\u00aa\u0017\u0001\u0000\u0000\u0000\u00ab\u00a9\u0001"+
		"\u0000\u0000\u0000\u00ac\u00ad\u00058\u0000\u0000\u00ad\u00b1\u0003\u001c"+
		"\u000e\u0000\u00ae\u00af\u0005\u0015\u0000\u0000\u00af\u00b1\u0003\u001c"+
		"\u000e\u0000\u00b0\u00ac\u0001\u0000\u0000\u0000\u00b0\u00ae\u0001\u0000"+
		"\u0000\u0000\u00b1\u0019\u0001\u0000\u0000\u0000\u00b2\u00b7\u0003\u001c"+
		"\u000e\u0000\u00b3\u00b4\u0005\u0002\u0000\u0000\u00b4\u00b6\u0003\u001c"+
		"\u000e\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u00b9\u0001\u0000"+
		"\u0000\u0000\u00b7\u00b5\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000"+
		"\u0000\u0000\u00b8\u001b\u0001\u0000\u0000\u0000\u00b9\u00b7\u0001\u0000"+
		"\u0000\u0000\u00ba\u00c4\u00058\u0000\u0000\u00bb\u00bc\u0005!\u0000\u0000"+
		"\u00bc\u00bd\u0003\u001a\r\u0000\u00bd\u00be\u0005\"\u0000\u0000\u00be"+
		"\u00c4\u0001\u0000\u0000\u0000\u00bf\u00c0\u0005\u001d\u0000\u0000\u00c0"+
		"\u00c1\u0003\u001a\r\u0000\u00c1\u00c2\u0005\u001e\u0000\u0000\u00c2\u00c4"+
		"\u0001\u0000\u0000\u0000\u00c3\u00ba\u0001\u0000\u0000\u0000\u00c3\u00bb"+
		"\u0001\u0000\u0000\u0000\u00c3\u00bf\u0001\u0000\u0000\u0000\u00c4\u001d"+
		"\u0001\u0000\u0000\u0000\u00c5\u00c6\u00058\u0000\u0000\u00c6\u00c7\u0007"+
		"\u0002\u0000\u0000\u00c7\u00c8\u0003\"\u0011\u0000\u00c8\u00c9\u0005:"+
		"\u0000\u0000\u00c9\u001f\u0001\u0000\u0000\u0000\u00ca\u00cb\u0003\"\u0011"+
		"\u0000\u00cb\u00cc\u0005:\u0000\u0000\u00cc!\u0001\u0000\u0000\u0000\u00cd"+
		"\u00d2\u0003$\u0012\u0000\u00ce\u00cf\u0005\u0002\u0000\u0000\u00cf\u00d1"+
		"\u0003$\u0012\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d1\u00d4\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d3#\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000"+
		"\u0000\u0000\u00d5\u00d6\u0006\u0012\uffff\uffff\u0000\u00d6\u00d7\u0005"+
		"\u001d\u0000\u0000\u00d7\u00d8\u0003$\u0012\u0000\u00d8\u00d9\u0005\u001e"+
		"\u0000\u0000\u00d9\u00da\u0003$\u0012\u0013\u00da\u00e6\u0001\u0000\u0000"+
		"\u0000\u00db\u00dc\u0005\u001d\u0000\u0000\u00dc\u00dd\u0003$\u0012\u0000"+
		"\u00dd\u00de\u0005\u001e\u0000\u0000\u00de\u00e6\u0001\u0000\u0000\u0000"+
		"\u00df\u00e0\u0007\u0003\u0000\u0000\u00e0\u00e6\u0003$\u0012\u0010\u00e1"+
		"\u00e6\u0003,\u0016\u0000\u00e2\u00e6\u0003&\u0013\u0000\u00e3\u00e6\u0003"+
		"(\u0014\u0000\u00e4\u00e6\u0003*\u0015\u0000\u00e5\u00d5\u0001\u0000\u0000"+
		"\u0000\u00e5\u00db\u0001\u0000\u0000\u0000\u00e5\u00df\u0001\u0000\u0000"+
		"\u0000\u00e5\u00e1\u0001\u0000\u0000\u0000\u00e5\u00e2\u0001\u0000\u0000"+
		"\u0000\u00e5\u00e3\u0001\u0000\u0000\u0000\u00e5\u00e4\u0001\u0000\u0000"+
		"\u0000\u00e6\u0119\u0001\u0000\u0000\u0000\u00e7\u00e8\n\u0011\u0000\u0000"+
		"\u00e8\u00e9\u0007\u0004\u0000\u0000\u00e9\u0118\u0003$\u0012\u0012\u00ea"+
		"\u00eb\n\u000f\u0000\u0000\u00eb\u00ec\u0007\u0005\u0000\u0000\u00ec\u0118"+
		"\u0003$\u0012\u0010\u00ed\u00ee\n\u000e\u0000\u0000\u00ee\u00ef\u0007"+
		"\u0006\u0000\u0000\u00ef\u0118\u0003$\u0012\u000f\u00f0\u00f1\n\r\u0000"+
		"\u0000\u00f1\u00f2\u0007\u0007\u0000\u0000\u00f2\u0118\u0003$\u0012\u000e"+
		"\u00f3\u00f4\n\f\u0000\u0000\u00f4\u00f5\u0007\b\u0000\u0000\u00f5\u0118"+
		"\u0003$\u0012\r\u00f6\u00f7\n\u000b\u0000\u0000\u00f7\u00f8\u0005-\u0000"+
		"\u0000\u00f8\u0118\u0003$\u0012\f\u00f9\u00fa\n\n\u0000\u0000\u00fa\u00fb"+
		"\u0005,\u0000\u0000\u00fb\u0118\u0003$\u0012\u000b\u00fc\u00fd\n\t\u0000"+
		"\u0000\u00fd\u00fe\u0005+\u0000\u0000\u00fe\u0118\u0003$\u0012\n\u00ff"+
		"\u0100\n\b\u0000\u0000\u0100\u0102\u0005\u001d\u0000\u0000\u0101\u0103"+
		"\u0003\"\u0011\u0000\u0102\u0101\u0001\u0000\u0000\u0000\u0102\u0103\u0001"+
		"\u0000\u0000\u0000\u0103\u0104\u0001\u0000\u0000\u0000\u0104\u0118\u0005"+
		"\u001e\u0000\u0000\u0105\u0106\n\u0007\u0000\u0000\u0106\u0107\u0005\u0001"+
		"\u0000\u0000\u0107\u0118\u00058\u0000\u0000\u0108\u0109\n\u0006\u0000"+
		"\u0000\u0109\u010a\u0005\u001f\u0000\u0000\u010a\u010b\u0003$\u0012\u0000"+
		"\u010b\u010c\u0005 \u0000\u0000\u010c\u0118\u0001\u0000\u0000\u0000\u010d"+
		"\u010e\n\u0005\u0000\u0000\u010e\u0110\u0005\u001f\u0000\u0000\u010f\u0111"+
		"\u0003$\u0012\u0000\u0110\u010f\u0001\u0000\u0000\u0000\u0110\u0111\u0001"+
		"\u0000\u0000\u0000\u0111\u0112\u0001\u0000\u0000\u0000\u0112\u0114\u0005"+
		"\u0003\u0000\u0000\u0113\u0115\u0003$\u0012\u0000\u0114\u0113\u0001\u0000"+
		"\u0000\u0000\u0114\u0115\u0001\u0000\u0000\u0000\u0115\u0116\u0001\u0000"+
		"\u0000\u0000\u0116\u0118\u0005 \u0000\u0000\u0117\u00e7\u0001\u0000\u0000"+
		"\u0000\u0117\u00ea\u0001\u0000\u0000\u0000\u0117\u00ed\u0001\u0000\u0000"+
		"\u0000\u0117\u00f0\u0001\u0000\u0000\u0000\u0117\u00f3\u0001\u0000\u0000"+
		"\u0000\u0117\u00f6\u0001\u0000\u0000\u0000\u0117\u00f9\u0001\u0000\u0000"+
		"\u0000\u0117\u00fc\u0001\u0000\u0000\u0000\u0117\u00ff\u0001\u0000\u0000"+
		"\u0000\u0117\u0105\u0001\u0000\u0000\u0000\u0117\u0108\u0001\u0000\u0000"+
		"\u0000\u0117\u010d\u0001\u0000\u0000\u0000\u0118\u011b\u0001\u0000\u0000"+
		"\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u0119\u011a\u0001\u0000\u0000"+
		"\u0000\u011a%\u0001\u0000\u0000\u0000\u011b\u0119\u0001\u0000\u0000\u0000"+
		"\u011c\u011d\u0007\u0000\u0000\u0000\u011d\u011f\u0005\u001d\u0000\u0000"+
		"\u011e\u0120\u0003\u0012\t\u0000\u011f\u011e\u0001\u0000\u0000\u0000\u011f"+
		"\u0120\u0001\u0000\u0000\u0000\u0120\u0121\u0001\u0000\u0000\u0000\u0121"+
		"\u0123\u0005\u001e\u0000\u0000\u0122\u0124\u0007\u0001\u0000\u0000\u0123"+
		"\u0122\u0001\u0000\u0000\u0000\u0123\u0124\u0001\u0000\u0000\u0000\u0124"+
		"\u0125\u0001\u0000\u0000\u0000\u0125\u0126\u0003\u0000\u0000\u0000\u0126"+
		"\'\u0001\u0000\u0000\u0000\u0127\u0128\u0007\t\u0000\u0000\u0128\u0129"+
		"\u0005!\u0000\u0000\u0129\u012a\u00058\u0000\u0000\u012a\u012b\u0005\u0003"+
		"\u0000\u0000\u012b\u0132\u0003$\u0012\u0000\u012c\u012d\u0005\u0002\u0000"+
		"\u0000\u012d\u012e\u00058\u0000\u0000\u012e\u012f\u0005\u0003\u0000\u0000"+
		"\u012f\u0131\u0003$\u0012\u0000\u0130\u012c\u0001\u0000\u0000\u0000\u0131"+
		"\u0134\u0001\u0000\u0000\u0000\u0132\u0130\u0001\u0000\u0000\u0000\u0132"+
		"\u0133\u0001\u0000\u0000\u0000\u0133\u0136\u0001\u0000\u0000\u0000\u0134"+
		"\u0132\u0001\u0000\u0000\u0000\u0135\u0137\u0005\u0002\u0000\u0000\u0136"+
		"\u0135\u0001\u0000\u0000\u0000\u0136\u0137\u0001\u0000\u0000\u0000\u0137"+
		"\u0138\u0001\u0000\u0000\u0000\u0138\u0139\u0005\"\u0000\u0000\u0139)"+
		"\u0001\u0000\u0000\u0000\u013a\u013b\u0005\u001d\u0000\u0000\u013b\u0143"+
		"\u0003$\u0012\u0000\u013c\u0144\u0005\u0002\u0000\u0000\u013d\u013e\u0005"+
		"\u0002\u0000\u0000\u013e\u0140\u0003$\u0012\u0000\u013f\u013d\u0001\u0000"+
		"\u0000\u0000\u0140\u0141\u0001\u0000\u0000\u0000\u0141\u013f\u0001\u0000"+
		"\u0000\u0000\u0141\u0142\u0001\u0000\u0000\u0000\u0142\u0144\u0001\u0000"+
		"\u0000\u0000\u0143\u013c\u0001\u0000\u0000\u0000\u0143\u013f\u0001\u0000"+
		"\u0000\u0000\u0144\u0145\u0001\u0000\u0000\u0000\u0145\u0146\u0005\u001e"+
		"\u0000\u0000\u0146+\u0001\u0000\u0000\u0000\u0147\u014a\u00058\u0000\u0000"+
		"\u0148\u014a\u00030\u0018\u0000\u0149\u0147\u0001\u0000\u0000\u0000\u0149"+
		"\u0148\u0001\u0000\u0000\u0000\u014a-\u0001\u0000\u0000\u0000\u014b\u0150"+
		"\u00058\u0000\u0000\u014c\u014d\u0005\u0002\u0000\u0000\u014d\u014f\u0005"+
		"8\u0000\u0000\u014e\u014c\u0001\u0000\u0000\u0000\u014f\u0152\u0001\u0000"+
		"\u0000\u0000\u0150\u014e\u0001\u0000\u0000\u0000\u0150\u0151\u0001\u0000"+
		"\u0000\u0000\u0151/\u0001\u0000\u0000\u0000\u0152\u0150\u0001\u0000\u0000"+
		"\u0000\u0153\u0156\u00032\u0019\u0000\u0154\u0156\u00034\u001a\u0000\u0155"+
		"\u0153\u0001\u0000\u0000\u0000\u0155\u0154\u0001\u0000\u0000\u0000\u0156"+
		"1\u0001\u0000\u0000\u0000\u0157\u0158\u00059\u0000\u0000\u01583\u0001"+
		"\u0000\u0000\u0000\u0159\u015a\u00058\u0000\u0000\u015a\u015b\u00059\u0000"+
		"\u0000\u015b5\u0001\u0000\u0000\u0000$:BHNVXbjotz\u0080\u0084\u008e\u0095"+
		"\u00a2\u00a9\u00b0\u00b7\u00c3\u00d2\u00e5\u0102\u0110\u0114\u0117\u0119"+
		"\u011f\u0123\u0132\u0136\u0141\u0143\u0149\u0150\u0155";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
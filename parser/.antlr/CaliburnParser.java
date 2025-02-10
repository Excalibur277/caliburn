// Generated from /home/excalibur/PersonalProjects/caliburncode/parser/Caliburn.g4 by ANTLR 4.13.1
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
		OP_ADD=1, OP_SUB=2, OP_MUL=3, OP_DIV=4, OP_MOD=5, OP_POW=6, OP_ROOT=7, 
		OP_OR=8, OP_XOR=9, OP_AND=10, OP_NOT=11, OP_EQU=12, OP_NEQ=13, OP_GTE=14, 
		OP_GRT=15, OP_LTE=16, OP_LST=17, WhiteSpace=18, Comment=19, Identifier=20, 
		Literal=21;
	public static final int
		RULE_expression = 0, RULE_operator_expression = 1, RULE_expression_atom = 2, 
		RULE_value_atom = 3, RULE_type_atom = 4, RULE_untyped_literal_atom = 5, 
		RULE_typed_literal_atom = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"expression", "operator_expression", "expression_atom", "value_atom", 
			"type_atom", "untyped_literal_atom", "typed_literal_atom"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", 
			"'&'", "'!'", "'=='", "'!='", "'>='", "'>'", "'<='", "'<'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "OP_ADD", "OP_SUB", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", 
			"OP_OR", "OP_XOR", "OP_AND", "OP_NOT", "OP_EQU", "OP_NEQ", "OP_GTE", 
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
	public String getGrammarFileName() { return "Caliburn.g4"; }

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
	public static class ExpressionContext extends ParserRuleContext {
		public Operator_expressionContext operator_expression() {
			return getRuleContext(Operator_expressionContext.class,0);
		}
		public Expression_atomContext expression_atom() {
			return getRuleContext(Expression_atomContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_expression);
		try {
			setState(16);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EOF:
				enterOuterAlt(_localctx, 1);
				{
				setState(14);
				operator_expression();
				}
				break;
			case Identifier:
			case Literal:
				enterOuterAlt(_localctx, 2);
				{
				setState(15);
				expression_atom();
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
	public static class Operator_expressionContext extends ParserRuleContext {
		public Operator_expressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operator_expression; }
	}

	public final Operator_expressionContext operator_expression() throws RecognitionException {
		Operator_expressionContext _localctx = new Operator_expressionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_operator_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
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
		public Value_atomContext value_atom() {
			return getRuleContext(Value_atomContext.class,0);
		}
		public Untyped_literal_atomContext untyped_literal_atom() {
			return getRuleContext(Untyped_literal_atomContext.class,0);
		}
		public Typed_literal_atomContext typed_literal_atom() {
			return getRuleContext(Typed_literal_atomContext.class,0);
		}
		public Expression_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression_atom; }
	}

	public final Expression_atomContext expression_atom() throws RecognitionException {
		Expression_atomContext _localctx = new Expression_atomContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_expression_atom);
		try {
			setState(23);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(20);
				value_atom();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(21);
				untyped_literal_atom();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(22);
				typed_literal_atom();
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
	public static class Value_atomContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Value_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value_atom; }
	}

	public final Value_atomContext value_atom() throws RecognitionException {
		Value_atomContext _localctx = new Value_atomContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_value_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(Identifier);
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
	public static class Type_atomContext extends ParserRuleContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Type_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type_atom; }
	}

	public final Type_atomContext type_atom() throws RecognitionException {
		Type_atomContext _localctx = new Type_atomContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_type_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(27);
			match(Identifier);
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
		enterRule(_localctx, 10, RULE_untyped_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(29);
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
		public Type_atomContext type_atom() {
			return getRuleContext(Type_atomContext.class,0);
		}
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public Typed_literal_atomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typed_literal_atom; }
	}

	public final Typed_literal_atomContext typed_literal_atom() throws RecognitionException {
		Typed_literal_atomContext _localctx = new Typed_literal_atomContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_typed_literal_atom);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(31);
			type_atom();
			setState(32);
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

	public static final String _serializedATN =
		"\u0004\u0001\u0015#\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0001\u0000\u0001\u0000\u0003"+
		"\u0000\u0011\b\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002\u0018\b\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0000\u0000\u0007\u0000\u0002\u0004\u0006\b\n\f\u0000\u0000\u001e"+
		"\u0000\u0010\u0001\u0000\u0000\u0000\u0002\u0012\u0001\u0000\u0000\u0000"+
		"\u0004\u0017\u0001\u0000\u0000\u0000\u0006\u0019\u0001\u0000\u0000\u0000"+
		"\b\u001b\u0001\u0000\u0000\u0000\n\u001d\u0001\u0000\u0000\u0000\f\u001f"+
		"\u0001\u0000\u0000\u0000\u000e\u0011\u0003\u0002\u0001\u0000\u000f\u0011"+
		"\u0003\u0004\u0002\u0000\u0010\u000e\u0001\u0000\u0000\u0000\u0010\u000f"+
		"\u0001\u0000\u0000\u0000\u0011\u0001\u0001\u0000\u0000\u0000\u0012\u0013"+
		"\u0001\u0000\u0000\u0000\u0013\u0003\u0001\u0000\u0000\u0000\u0014\u0018"+
		"\u0003\u0006\u0003\u0000\u0015\u0018\u0003\n\u0005\u0000\u0016\u0018\u0003"+
		"\f\u0006\u0000\u0017\u0014\u0001\u0000\u0000\u0000\u0017\u0015\u0001\u0000"+
		"\u0000\u0000\u0017\u0016\u0001\u0000\u0000\u0000\u0018\u0005\u0001\u0000"+
		"\u0000\u0000\u0019\u001a\u0005\u0014\u0000\u0000\u001a\u0007\u0001\u0000"+
		"\u0000\u0000\u001b\u001c\u0005\u0014\u0000\u0000\u001c\t\u0001\u0000\u0000"+
		"\u0000\u001d\u001e\u0005\u0015\u0000\u0000\u001e\u000b\u0001\u0000\u0000"+
		"\u0000\u001f \u0003\b\u0004\u0000 !\u0005\u0015\u0000\u0000!\r\u0001\u0000"+
		"\u0000\u0000\u0002\u0010\u0017";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
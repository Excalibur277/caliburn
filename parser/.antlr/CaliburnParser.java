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
		CommentMultiLine=56, CommentSingleLine=57, Identifier=58, Literal=59;
	public static final int
		RULE_module = 0, RULE_definitions = 1, RULE_definition = 2, RULE_function_definition = 3, 
		RULE_parameters = 4, RULE_parameters_list = 5, RULE_parameter = 6, RULE_block = 7, 
		RULE_statements = 8, RULE_statement = 9, RULE_jump_statement = 10, RULE_return_statement = 11, 
		RULE_break_statement = 12, RULE_continue_statement = 13, RULE_control_statement = 14, 
		RULE_if_statement = 15, RULE_else_statement = 16, RULE_for_statement = 17, 
		RULE_switch_statement = 18, RULE_case_blocks = 19, RULE_option_case_block = 20, 
		RULE_default_case_block = 21, RULE_inline_statements = 22, RULE_inline_statement = 23, 
		RULE_assign_statement = 24, RULE_assign_expressions = 25, RULE_assign_expression = 26, 
		RULE_assign_declarations = 27, RULE_assign_declaration = 28, RULE_typed_assign_declarations = 29, 
		RULE_typed_assign_declaration = 30, RULE_untyped_assign_declarations = 31, 
		RULE_untyped_assign_declaration = 32, RULE_expression_statement = 33, 
		RULE_expressions = 34, RULE_expression = 35, RULE_function_type = 36, 
		RULE_struct_type = 37, RULE_tuple_type = 38, RULE_type_expression = 39, 
		RULE_identifiers = 40, RULE_literal_atom = 41;
	private static String[] makeRuleNames() {
		return new String[] {
			"module", "definitions", "definition", "function_definition", "parameters", 
			"parameters_list", "parameter", "block", "statements", "statement", "jump_statement", 
			"return_statement", "break_statement", "continue_statement", "control_statement", 
			"if_statement", "else_statement", "for_statement", "switch_statement", 
			"case_blocks", "option_case_block", "default_case_block", "inline_statements", 
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
			"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "CommentMultiLine", "CommentSingleLine", 
			"Identifier", "Literal"
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
		public DefinitionsContext definitions() {
			return getRuleContext(DefinitionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(CaliburnParser.EOF, 0); }
		public ModuleRuleContext(ModuleContext ctx) { copyFrom(ctx); }
	}

	public final ModuleContext module() throws RecognitionException {
		ModuleContext _localctx = new ModuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_module);
		try {
			_localctx = new ModuleRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			definitions(0);
			setState(85);
			match(EOF);
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
	public static class DefinitionsContext extends ParserRuleContext {
		public DefinitionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_definitions; }
	 
		public DefinitionsContext() { }
		public void copyFrom(DefinitionsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefinitionsAppendContext extends DefinitionsContext {
		public DefinitionsContext definitions() {
			return getRuleContext(DefinitionsContext.class,0);
		}
		public DefinitionContext definition() {
			return getRuleContext(DefinitionContext.class,0);
		}
		public DefinitionsAppendContext(DefinitionsContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefinitionsInitialContext extends DefinitionsContext {
		public DefinitionsInitialContext(DefinitionsContext ctx) { copyFrom(ctx); }
	}

	public final DefinitionsContext definitions() throws RecognitionException {
		return definitions(0);
	}

	private DefinitionsContext definitions(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		DefinitionsContext _localctx = new DefinitionsContext(_ctx, _parentState);
		DefinitionsContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_definitions, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new DefinitionsInitialContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			}
			_ctx.stop = _input.LT(-1);
			setState(92);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new DefinitionsAppendContext(new DefinitionsContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_definitions);
					setState(88);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(89);
					definition();
					}
					} 
				}
				setState(94);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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
		enterRule(_localctx, 4, RULE_definition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
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
	public static class FunctionDefinitionNoReturnTypeContext extends Function_definitionContext {
		public Function_typeContext type;
		public Token name;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public FunctionDefinitionNoReturnTypeContext(Function_definitionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionDefinitionContext extends Function_definitionContext {
		public Function_typeContext type;
		public Token name;
		public Type_expressionContext return_type;
		public TerminalNode L_PAREN() { return getToken(CaliburnParser.L_PAREN, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(CaliburnParser.R_PAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Function_typeContext function_type() {
			return getRuleContext(Function_typeContext.class,0);
		}
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public Type_expressionContext type_expression() {
			return getRuleContext(Type_expressionContext.class,0);
		}
		public FunctionDefinitionContext(Function_definitionContext ctx) { copyFrom(ctx); }
	}

	public final Function_definitionContext function_definition() throws RecognitionException {
		Function_definitionContext _localctx = new Function_definitionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_function_definition);
		try {
			setState(113);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				_localctx = new FunctionDefinitionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				((FunctionDefinitionContext)_localctx).type = function_type();
				setState(98);
				((FunctionDefinitionContext)_localctx).name = match(Identifier);
				setState(99);
				match(L_PAREN);
				setState(100);
				parameters();
				setState(101);
				match(R_PAREN);
				setState(102);
				((FunctionDefinitionContext)_localctx).return_type = type_expression(0);
				setState(103);
				block();
				}
				break;
			case 2:
				_localctx = new FunctionDefinitionNoReturnTypeContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(105);
				((FunctionDefinitionNoReturnTypeContext)_localctx).type = function_type();
				setState(106);
				((FunctionDefinitionNoReturnTypeContext)_localctx).name = match(Identifier);
				setState(107);
				match(L_PAREN);
				setState(108);
				parameters();
				setState(109);
				match(R_PAREN);
				setState(110);
				((FunctionDefinitionNoReturnTypeContext)_localctx).return_type = type_expression(0);
				setState(111);
				block();
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
	public static class ParametersEmptyContext extends ParametersContext {
		public ParametersEmptyContext(ParametersContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametersFilledContext extends ParametersContext {
		public Parameters_listContext parameters_list() {
			return getRuleContext(Parameters_listContext.class,0);
		}
		public ParametersFilledContext(ParametersContext ctx) { copyFrom(ctx); }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_parameters);
		try {
			setState(117);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case R_PAREN:
			case R_S_BRACK:
			case R_C_BRACK:
				_localctx = new ParametersEmptyContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				}
				break;
			case L_S_BRACK:
			case L_C_BRACK:
			case Identifier:
				_localctx = new ParametersFilledContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				parameters_list(0);
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
	public static class Parameters_listContext extends ParserRuleContext {
		public Parameters_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters_list; }
	 
		public Parameters_listContext() { }
		public void copyFrom(Parameters_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametersListInitialContext extends Parameters_listContext {
		public ParameterContext parameter() {
			return getRuleContext(ParameterContext.class,0);
		}
		public ParametersListInitialContext(Parameters_listContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParametersListAppendContext extends Parameters_listContext {
		public Parameters_listContext parameters_list() {
			return getRuleContext(Parameters_listContext.class,0);
		}
		public TerminalNode COMMA() { return getToken(CaliburnParser.COMMA, 0); }
		public ParameterContext parameter() {
			return getRuleContext(ParameterContext.class,0);
		}
		public ParametersListAppendContext(Parameters_listContext ctx) { copyFrom(ctx); }
	}

	public final Parameters_listContext parameters_list() throws RecognitionException {
		return parameters_list(0);
	}

	private Parameters_listContext parameters_list(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Parameters_listContext _localctx = new Parameters_listContext(_ctx, _parentState);
		Parameters_listContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_parameters_list, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new ParametersListInitialContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(120);
			parameter();
			}
			_ctx.stop = _input.LT(-1);
			setState(127);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ParametersListAppendContext(new Parameters_listContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_parameters_list);
					setState(122);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(123);
					match(COMMA);
					setState(124);
					parameter();
					}
					} 
				}
				setState(129);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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
		enterRule(_localctx, 12, RULE_parameter);
		try {
			setState(140);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new TypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(130);
				typed_assign_declaration();
				}
				break;
			case 2:
				_localctx = new UntypedParameterContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(131);
				untyped_assign_declaration();
				}
				break;
			case 3:
				_localctx = new StructDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(132);
				match(L_C_BRACK);
				setState(133);
				parameters();
				setState(134);
				match(R_C_BRACK);
				}
				break;
			case 4:
				_localctx = new TupleDestrutureParameterContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(136);
				match(L_S_BRACK);
				setState(137);
				parameters();
				setState(138);
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
		enterRule(_localctx, 14, RULE_block);
		int _la;
		try {
			_localctx = new BlockRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(142);
			match(L_C_BRACK);
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 864691632145140992L) != 0)) {
				{
				{
				setState(143);
				statement();
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
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
	public static class StatementsContext extends ParserRuleContext {
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	 
		public StatementsContext() { }
		public void copyFrom(StatementsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementsAppendContext extends StatementsContext {
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public StatementsAppendContext(StatementsContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StatementsInitialContext extends StatementsContext {
		public StatementsInitialContext(StatementsContext ctx) { copyFrom(ctx); }
	}

	public final StatementsContext statements() throws RecognitionException {
		return statements(0);
	}

	private StatementsContext statements(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		StatementsContext _localctx = new StatementsContext(_ctx, _parentState);
		StatementsContext _prevctx = _localctx;
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_statements, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new StatementsInitialContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			}
			_ctx.stop = _input.LT(-1);
			setState(156);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new StatementsAppendContext(new StatementsContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_statements);
					setState(152);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(153);
					statement();
					}
					} 
				}
				setState(158);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
		enterRule(_localctx, 18, RULE_statement);
		try {
			setState(163);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(159);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(160);
				expression_statement();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(161);
				control_statement();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(162);
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
		enterRule(_localctx, 20, RULE_jump_statement);
		try {
			setState(168);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				return_statement();
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				break_statement();
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(167);
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
		enterRule(_localctx, 22, RULE_return_statement);
		try {
			_localctx = new ReturnStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(RETURN);
			setState(171);
			expression(0);
			setState(172);
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
		enterRule(_localctx, 24, RULE_break_statement);
		try {
			_localctx = new BreakStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			match(BREAK);
			setState(175);
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
		enterRule(_localctx, 26, RULE_continue_statement);
		try {
			_localctx = new ContinueStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(CONTINUE);
			setState(178);
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
		enterRule(_localctx, 28, RULE_control_statement);
		try {
			setState(183);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				if_statement();
				}
				break;
			case FOR:
				enterOuterAlt(_localctx, 2);
				{
				setState(181);
				for_statement();
				}
				break;
			case SWITCH:
				enterOuterAlt(_localctx, 3);
				{
				setState(182);
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
		public Inline_statementsContext inline_statements() {
			return getRuleContext(Inline_statementsContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public IfStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfElseStatementContext extends If_statementContext {
		public TerminalNode IF() { return getToken(CaliburnParser.IF, 0); }
		public Inline_statementsContext inline_statements() {
			return getRuleContext(Inline_statementsContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public Else_statementContext else_statement() {
			return getRuleContext(Else_statementContext.class,0);
		}
		public IfElseStatementContext(If_statementContext ctx) { copyFrom(ctx); }
	}

	public final If_statementContext if_statement() throws RecognitionException {
		If_statementContext _localctx = new If_statementContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_if_statement);
		try {
			setState(196);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new IfStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(185);
				match(IF);
				setState(186);
				inline_statements(0);
				setState(187);
				expression(0);
				setState(188);
				block();
				}
				break;
			case 2:
				_localctx = new IfElseStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(190);
				match(IF);
				setState(191);
				inline_statements(0);
				setState(192);
				expression(0);
				setState(193);
				block();
				setState(194);
				else_statement();
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
	public static class Else_statementContext extends ParserRuleContext {
		public Else_statementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_statement; }
	 
		public Else_statementContext() { }
		public void copyFrom(Else_statementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseIfStatementContext extends Else_statementContext {
		public TerminalNode ELSE() { return getToken(CaliburnParser.ELSE, 0); }
		public If_statementContext if_statement() {
			return getRuleContext(If_statementContext.class,0);
		}
		public ElseIfStatementContext(Else_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseStatementContext extends Else_statementContext {
		public TerminalNode ELSE() { return getToken(CaliburnParser.ELSE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ElseStatementContext(Else_statementContext ctx) { copyFrom(ctx); }
	}

	public final Else_statementContext else_statement() throws RecognitionException {
		Else_statementContext _localctx = new Else_statementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_else_statement);
		try {
			setState(202);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				_localctx = new ElseStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(198);
				match(ELSE);
				setState(199);
				block();
				}
				break;
			case 2:
				_localctx = new ElseIfStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(200);
				match(ELSE);
				setState(201);
				if_statement();
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
		public Inline_statementsContext inline_statements() {
			return getRuleContext(Inline_statementsContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForStatementContext(For_statementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStatementWithAfterContext extends For_statementContext {
		public TerminalNode FOR() { return getToken(CaliburnParser.FOR, 0); }
		public List<Inline_statementsContext> inline_statements() {
			return getRuleContexts(Inline_statementsContext.class);
		}
		public Inline_statementsContext inline_statements(int i) {
			return getRuleContext(Inline_statementsContext.class,i);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(CaliburnParser.ARROW, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForStatementWithAfterContext(For_statementContext ctx) { copyFrom(ctx); }
	}

	public final For_statementContext for_statement() throws RecognitionException {
		For_statementContext _localctx = new For_statementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_for_statement);
		try {
			setState(216);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new ForStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(204);
				match(FOR);
				setState(205);
				inline_statements(0);
				setState(206);
				expression(0);
				setState(207);
				block();
				}
				break;
			case 2:
				_localctx = new ForStatementWithAfterContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(209);
				match(FOR);
				setState(210);
				inline_statements(0);
				setState(211);
				expression(0);
				setState(212);
				match(ARROW);
				setState(213);
				inline_statements(0);
				setState(214);
				block();
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
		public Inline_statementsContext inline_statements() {
			return getRuleContext(Inline_statementsContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode L_C_BRACK() { return getToken(CaliburnParser.L_C_BRACK, 0); }
		public Case_blocksContext case_blocks() {
			return getRuleContext(Case_blocksContext.class,0);
		}
		public TerminalNode R_C_BRACK() { return getToken(CaliburnParser.R_C_BRACK, 0); }
		public SwitchStatementContext(Switch_statementContext ctx) { copyFrom(ctx); }
	}

	public final Switch_statementContext switch_statement() throws RecognitionException {
		Switch_statementContext _localctx = new Switch_statementContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_switch_statement);
		try {
			_localctx = new SwitchStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(218);
			match(SWITCH);
			setState(219);
			inline_statements(0);
			setState(220);
			expression(0);
			setState(221);
			match(L_C_BRACK);
			setState(222);
			case_blocks();
			setState(223);
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
		enterRule(_localctx, 38, RULE_case_blocks);
		int _la;
		try {
			_localctx = new CaseBlocksContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(225);
				option_case_block();
				}
				}
				setState(230);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(232);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(231);
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
		enterRule(_localctx, 40, RULE_option_case_block);
		try {
			_localctx = new OptionCaseBlockContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			match(CASE);
			setState(235);
			expression(0);
			setState(236);
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
		enterRule(_localctx, 42, RULE_default_case_block);
		try {
			_localctx = new DefaultCaseBlockContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(238);
			match(DEFAULT);
			setState(239);
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
	public static class Inline_statementsContext extends ParserRuleContext {
		public Inline_statementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inline_statements; }
	 
		public Inline_statementsContext() { }
		public void copyFrom(Inline_statementsContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InlineStatementsInitialContext extends Inline_statementsContext {
		public InlineStatementsInitialContext(Inline_statementsContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InlineStatementsAppendContext extends Inline_statementsContext {
		public Inline_statementsContext inline_statements() {
			return getRuleContext(Inline_statementsContext.class,0);
		}
		public Inline_statementContext inline_statement() {
			return getRuleContext(Inline_statementContext.class,0);
		}
		public InlineStatementsAppendContext(Inline_statementsContext ctx) { copyFrom(ctx); }
	}

	public final Inline_statementsContext inline_statements() throws RecognitionException {
		return inline_statements(0);
	}

	private Inline_statementsContext inline_statements(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Inline_statementsContext _localctx = new Inline_statementsContext(_ctx, _parentState);
		Inline_statementsContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_inline_statements, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new InlineStatementsInitialContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			}
			_ctx.stop = _input.LT(-1);
			setState(246);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new InlineStatementsAppendContext(new Inline_statementsContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_inline_statements);
					setState(242);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(243);
					inline_statement();
					}
					} 
				}
				setState(248);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
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
		enterRule(_localctx, 46, RULE_inline_statement);
		try {
			setState(251);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(249);
				assign_statement();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(250);
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
		enterRule(_localctx, 48, RULE_assign_statement);
		int _la;
		try {
			setState(264);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				_localctx = new AssignStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(253);
				assign_declarations();
				setState(254);
				match(ASSIGN);
				setState(255);
				expressions();
				setState(256);
				match(Terminator);
				}
				break;
			case 2:
				_localctx = new AssignOperationStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(258);
				assign_expressions();
				setState(259);
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
				setState(260);
				match(ASSIGN);
				setState(261);
				expressions();
				setState(262);
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
		enterRule(_localctx, 50, RULE_assign_expressions);
		int _la;
		try {
			_localctx = new AssignExpressionsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(266);
			assign_expression();
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(267);
				match(COMMA);
				setState(268);
				assign_expression();
				}
				}
				setState(273);
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
		enterRule(_localctx, 52, RULE_assign_expression);
		try {
			setState(283);
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
				setState(274);
				expression(0);
				}
				break;
			case L_C_BRACK:
				_localctx = new StructDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(275);
				match(L_C_BRACK);
				setState(276);
				assign_expressions();
				setState(277);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new TupleDestrutureAssignExpressionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(279);
				match(L_S_BRACK);
				setState(280);
				assign_expressions();
				setState(281);
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
		enterRule(_localctx, 54, RULE_assign_declarations);
		int _la;
		try {
			_localctx = new AssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			assign_declaration();
			setState(290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(286);
				match(COMMA);
				setState(287);
				assign_statement();
				}
				}
				setState(292);
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
		enterRule(_localctx, 56, RULE_assign_declaration);
		try {
			setState(305);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				_localctx = new ExpressionAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(293);
				expression(0);
				}
				break;
			case 2:
				_localctx = new TypedAssignDeclarationDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(294);
				typed_assign_declaration();
				}
				break;
			case 3:
				_localctx = new UntypedAssignDeclarationDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(295);
				match(VAR);
				setState(296);
				untyped_assign_declaration();
				}
				break;
			case 4:
				_localctx = new StructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(297);
				match(L_C_BRACK);
				setState(298);
				assign_declarations();
				setState(299);
				match(R_C_BRACK);
				}
				break;
			case 5:
				_localctx = new TupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(301);
				match(L_S_BRACK);
				setState(302);
				assign_declarations();
				setState(303);
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
		enterRule(_localctx, 58, RULE_typed_assign_declarations);
		int _la;
		try {
			_localctx = new TypedAssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(307);
			typed_assign_declaration();
			setState(312);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(308);
				match(COMMA);
				setState(309);
				typed_assign_declaration();
				}
				}
				setState(314);
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
		enterRule(_localctx, 60, RULE_typed_assign_declaration);
		try {
			_localctx = new TypedAssignDeclarationContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			((TypedAssignDeclarationContext)_localctx).type = type_expression(0);
			setState(316);
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
		enterRule(_localctx, 62, RULE_untyped_assign_declarations);
		int _la;
		try {
			_localctx = new UntypedAssignDeclarationsContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(318);
			untyped_assign_declaration();
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(319);
				match(COMMA);
				setState(320);
				untyped_assign_declaration();
				}
				}
				setState(325);
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
		enterRule(_localctx, 64, RULE_untyped_assign_declaration);
		try {
			setState(335);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Identifier:
				_localctx = new UntypedIdentifierAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(326);
				((UntypedIdentifierAssignDeclarationContext)_localctx).var = match(Identifier);
				}
				break;
			case L_C_BRACK:
				_localctx = new UntypedStructDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(327);
				match(L_C_BRACK);
				setState(328);
				untyped_assign_declarations();
				setState(329);
				match(R_C_BRACK);
				}
				break;
			case L_S_BRACK:
				_localctx = new UntypedTupleDestrutureAssignDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(331);
				match(L_S_BRACK);
				setState(332);
				untyped_assign_declarations();
				setState(333);
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
		enterRule(_localctx, 66, RULE_expression_statement);
		try {
			_localctx = new ExpressionStatementContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			expressions();
			setState(338);
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
		enterRule(_localctx, 68, RULE_expressions);
		int _la;
		try {
			_localctx = new ExpressionsRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			expression(0);
			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(341);
				match(COMMA);
				setState(342);
				expression(0);
				}
				}
				setState(347);
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
		int _startState = 70;
		enterRecursionRule(_localctx, 70, RULE_expression, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				_localctx = new CastExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(349);
				match(L_PAREN);
				setState(350);
				((CastExpressionContext)_localctx).type = type_expression(0);
				setState(351);
				match(R_PAREN);
				setState(352);
				((CastExpressionContext)_localctx).exp = expression(20);
				}
				break;
			case 2:
				{
				_localctx = new BracketedExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(354);
				match(L_PAREN);
				setState(355);
				((BracketedExpressionContext)_localctx).exp = expression(0);
				setState(356);
				match(R_PAREN);
				}
				break;
			case 3:
				{
				_localctx = new UnaryExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(358);
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
				setState(359);
				((UnaryExpressionContext)_localctx).exp = expression(17);
				}
				break;
			case 4:
				{
				_localctx = new IdentifierExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(360);
				((IdentifierExpressionContext)_localctx).identifier = match(Identifier);
				}
				break;
			case 5:
				{
				_localctx = new LiteralExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(361);
				((LiteralExpressionContext)_localctx).literal = literal_atom();
				}
				break;
			case 6:
				{
				_localctx = new FunctionExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(362);
				((FunctionExpressionContext)_localctx).type = function_type();
				setState(363);
				match(L_PAREN);
				setState(365);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 864691632144908288L) != 0)) {
					{
					setState(364);
					assign_declarations();
					}
				}

				setState(367);
				match(R_PAREN);
				setState(369);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==Identifier) {
					{
					setState(368);
					((FunctionExpressionContext)_localctx).return_type = type_expression(0);
					}
				}

				setState(371);
				block();
				}
				break;
			case 7:
				{
				_localctx = new StructExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(373);
				((StructExpressionContext)_localctx).type = struct_type();
				setState(374);
				match(L_C_BRACK);
				{
				setState(375);
				match(Identifier);
				}
				setState(376);
				match(COLON);
				setState(377);
				expression(0);
				setState(384);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(378);
						match(COMMA);
						{
						setState(379);
						match(Identifier);
						}
						setState(380);
						match(COLON);
						setState(381);
						expression(0);
						}
						} 
					}
					setState(386);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
				}
				setState(388);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COMMA) {
					{
					setState(387);
					match(COMMA);
					}
				}

				setState(390);
				match(R_C_BRACK);
				}
				break;
			case 8:
				{
				_localctx = new TupleExpressionContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(392);
				match(L_PAREN);
				setState(393);
				expression(0);
				setState(401);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
				case 1:
					{
					setState(394);
					match(COMMA);
					}
					break;
				case 2:
					{
					setState(397); 
					_errHandler.sync(this);
					_la = _input.LA(1);
					do {
						{
						{
						setState(395);
						match(COMMA);
						setState(396);
						expression(0);
						}
						}
						setState(399); 
						_errHandler.sync(this);
						_la = _input.LA(1);
					} while ( _la==COMMA );
					}
					break;
				}
				setState(403);
				match(R_PAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(457);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(455);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
					case 1:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(407);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(408);
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
						setState(409);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(19);
						}
						break;
					case 2:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(410);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(411);
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
						setState(412);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(17);
						}
						break;
					case 3:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(413);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(414);
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
						setState(415);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(16);
						}
						break;
					case 4:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(416);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(417);
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
						setState(418);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(15);
						}
						break;
					case 5:
						{
						_localctx = new BooleanBinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BooleanBinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(419);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(420);
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
						setState(421);
						((BooleanBinaryExpressionContext)_localctx).rhs_exp = expression(14);
						}
						break;
					case 6:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(422);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(423);
						match(OP_AND);
						setState(424);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(13);
						}
						break;
					case 7:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(425);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(426);
						match(OP_XOR);
						setState(427);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(12);
						}
						break;
					case 8:
						{
						_localctx = new BinaryExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((BinaryExpressionContext)_localctx).lhs_exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(428);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(429);
						match(OP_OR);
						setState(430);
						((BinaryExpressionContext)_localctx).rhs_exp = expression(11);
						}
						break;
					case 9:
						{
						_localctx = new CallExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((CallExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(431);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(432);
						match(L_PAREN);
						setState(434);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 864691610665877504L) != 0)) {
							{
							setState(433);
							((CallExpressionContext)_localctx).args = expressions();
							}
						}

						setState(436);
						match(R_PAREN);
						}
						break;
					case 10:
						{
						_localctx = new AccessExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((AccessExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(437);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(438);
						match(PERIOD);
						setState(439);
						((AccessExpressionContext)_localctx).identifier = match(Identifier);
						}
						break;
					case 11:
						{
						_localctx = new IndexExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((IndexExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(440);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(441);
						match(L_S_BRACK);
						setState(442);
						((IndexExpressionContext)_localctx).index = expression(0);
						setState(443);
						match(R_S_BRACK);
						}
						break;
					case 12:
						{
						_localctx = new SliceExpressionContext(new ExpressionContext(_parentctx, _parentState));
						((SliceExpressionContext)_localctx).exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(445);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(446);
						match(L_S_BRACK);
						setState(448);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 864691610665877504L) != 0)) {
							{
							setState(447);
							((SliceExpressionContext)_localctx).start_index = expression(0);
							}
						}

						setState(450);
						match(COLON);
						setState(452);
						_errHandler.sync(this);
						_la = _input.LA(1);
						if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 864691610665877504L) != 0)) {
							{
							setState(451);
							((SliceExpressionContext)_localctx).end_index = expression(0);
							}
						}

						setState(454);
						match(R_S_BRACK);
						}
						break;
					}
					} 
				}
				setState(459);
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
		enterRule(_localctx, 72, RULE_function_type);
		try {
			setState(462);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case FUNC:
				_localctx = new FunctionTypeFuncContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(460);
				match(FUNC);
				}
				break;
			case Identifier:
				_localctx = new FunctionTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(461);
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
		enterRule(_localctx, 74, RULE_struct_type);
		try {
			setState(466);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRUCT:
				_localctx = new StructTypeStructContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(464);
				match(STRUCT);
				}
				break;
			case Identifier:
				_localctx = new StructTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(465);
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
		enterRule(_localctx, 76, RULE_tuple_type);
		try {
			setState(470);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TUPLE:
				_localctx = new TupleTypeTupleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(468);
				match(TUPLE);
				}
				break;
			case Identifier:
				_localctx = new TupleTypeExpressionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(469);
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
		int _startState = 78;
		enterRecursionRule(_localctx, 78, RULE_type_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new IdentifierTypeExpressionContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(473);
			match(Identifier);
			}
			_ctx.stop = _input.LT(-1);
			setState(480);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new AccessTypeExpressionContext(new Type_expressionContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_type_expression);
					setState(475);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(476);
					match(PERIOD);
					setState(477);
					match(Identifier);
					}
					} 
				}
				setState(482);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,41,_ctx);
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
		enterRule(_localctx, 80, RULE_identifiers);
		int _la;
		try {
			_localctx = new IdentifiersRuleContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(483);
			match(Identifier);
			setState(488);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(484);
				match(COMMA);
				setState(485);
				match(Identifier);
				}
				}
				setState(490);
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
	public static class TypedLiteralContext extends Literal_atomContext {
		public TerminalNode Identifier() { return getToken(CaliburnParser.Identifier, 0); }
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public TypedLiteralContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UntypedLiteralContext extends Literal_atomContext {
		public TerminalNode Literal() { return getToken(CaliburnParser.Literal, 0); }
		public UntypedLiteralContext(Literal_atomContext ctx) { copyFrom(ctx); }
	}

	public final Literal_atomContext literal_atom() throws RecognitionException {
		Literal_atomContext _localctx = new Literal_atomContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_literal_atom);
		try {
			setState(494);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case Literal:
				_localctx = new UntypedLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(491);
				match(Literal);
				}
				break;
			case Identifier:
				_localctx = new TypedLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(492);
				match(Identifier);
				setState(493);
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
		case 1:
			return definitions_sempred((DefinitionsContext)_localctx, predIndex);
		case 5:
			return parameters_list_sempred((Parameters_listContext)_localctx, predIndex);
		case 8:
			return statements_sempred((StatementsContext)_localctx, predIndex);
		case 22:
			return inline_statements_sempred((Inline_statementsContext)_localctx, predIndex);
		case 35:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 39:
			return type_expression_sempred((Type_expressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean definitions_sempred(DefinitionsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean parameters_list_sempred(Parameters_listContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean statements_sempred(StatementsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean inline_statements_sempred(Inline_statementsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 3:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 18);
		case 5:
			return precpred(_ctx, 16);
		case 6:
			return precpred(_ctx, 15);
		case 7:
			return precpred(_ctx, 14);
		case 8:
			return precpred(_ctx, 13);
		case 9:
			return precpred(_ctx, 12);
		case 10:
			return precpred(_ctx, 11);
		case 11:
			return precpred(_ctx, 10);
		case 12:
			return precpred(_ctx, 9);
		case 13:
			return precpred(_ctx, 8);
		case 14:
			return precpred(_ctx, 7);
		case 15:
			return precpred(_ctx, 6);
		}
		return true;
	}
	private boolean type_expression_sempred(Type_expressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 16:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001;\u01f1\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"(\u0007(\u0002)\u0007)\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0005\u0001[\b\u0001\n\u0001\f\u0001^\t\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003r\b\u0003\u0001\u0004\u0001\u0004\u0003\u0004v\b\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005"+
		"\u0005~\b\u0005\n\u0005\f\u0005\u0081\t\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0003\u0006\u008d\b\u0006\u0001\u0007\u0001\u0007"+
		"\u0005\u0007\u0091\b\u0007\n\u0007\f\u0007\u0094\t\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\b\u0001\b\u0001\b\u0005\b\u009b\b\b\n\b\f\b\u009e\t\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0003\t\u00a4\b\t\u0001\n\u0001\n\u0001\n\u0003"+
		"\n\u00a9\b\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u00b8\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0003\u000f\u00c5\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0003\u0010\u00cb\b\u0010\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u00d9\b\u0011\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0013\u0005\u0013\u00e3\b\u0013\n\u0013\f\u0013\u00e6\t\u0013\u0001"+
		"\u0013\u0003\u0013\u00e9\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0005\u0016\u00f5\b\u0016\n\u0016\f\u0016\u00f8\t\u0016\u0001\u0017"+
		"\u0001\u0017\u0003\u0017\u00fc\b\u0017\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0003\u0018\u0109\b\u0018\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0005\u0019\u010e\b\u0019\n\u0019\f\u0019\u0111\t\u0019\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u011c\b\u001a\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0005\u001b\u0121\b\u001b\n\u001b\f\u001b\u0124\t\u001b"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0003\u001c\u0132\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d"+
		"\u0137\b\u001d\n\u001d\f\u001d\u013a\t\u001d\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0005\u001f\u0142\b\u001f\n"+
		"\u001f\f\u001f\u0145\t\u001f\u0001 \u0001 \u0001 \u0001 \u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0003 \u0150\b \u0001!\u0001!\u0001!\u0001\"\u0001"+
		"\"\u0001\"\u0005\"\u0158\b\"\n\"\f\"\u015b\t\"\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0003#\u016e\b#\u0001#\u0001#\u0003#\u0172\b#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0005#\u017f\b#\n#\f#\u0182\t#\u0001#\u0003#\u0185\b#\u0001#\u0001#"+
		"\u0001#\u0001#\u0001#\u0001#\u0001#\u0004#\u018e\b#\u000b#\f#\u018f\u0003"+
		"#\u0192\b#\u0001#\u0001#\u0003#\u0196\b#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0003#\u01b3\b#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001#\u0001#\u0001#\u0001#\u0001#\u0003#\u01c1\b#\u0001#\u0001"+
		"#\u0003#\u01c5\b#\u0001#\u0005#\u01c8\b#\n#\f#\u01cb\t#\u0001$\u0001$"+
		"\u0003$\u01cf\b$\u0001%\u0001%\u0003%\u01d3\b%\u0001&\u0001&\u0003&\u01d7"+
		"\b&\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0001\'\u0005\'\u01df\b\'"+
		"\n\'\f\'\u01e2\t\'\u0001(\u0001(\u0001(\u0005(\u01e7\b(\n(\f(\u01ea\t"+
		"(\u0001)\u0001)\u0001)\u0003)\u01ef\b)\u0001)\u0000\u0006\u0002\n\u0010"+
		",FN*\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018"+
		"\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPR\u0000\u0007\u0002\u0000"+
		"$%\'0\u0001\u0000$&\u0001\u0000*+\u0001\u0000\')\u0001\u0000$%\u0001\u0000"+
		"/0\u0001\u000016\u020d\u0000T\u0001\u0000\u0000\u0000\u0002W\u0001\u0000"+
		"\u0000\u0000\u0004_\u0001\u0000\u0000\u0000\u0006q\u0001\u0000\u0000\u0000"+
		"\bu\u0001\u0000\u0000\u0000\nw\u0001\u0000\u0000\u0000\f\u008c\u0001\u0000"+
		"\u0000\u0000\u000e\u008e\u0001\u0000\u0000\u0000\u0010\u0097\u0001\u0000"+
		"\u0000\u0000\u0012\u00a3\u0001\u0000\u0000\u0000\u0014\u00a8\u0001\u0000"+
		"\u0000\u0000\u0016\u00aa\u0001\u0000\u0000\u0000\u0018\u00ae\u0001\u0000"+
		"\u0000\u0000\u001a\u00b1\u0001\u0000\u0000\u0000\u001c\u00b7\u0001\u0000"+
		"\u0000\u0000\u001e\u00c4\u0001\u0000\u0000\u0000 \u00ca\u0001\u0000\u0000"+
		"\u0000\"\u00d8\u0001\u0000\u0000\u0000$\u00da\u0001\u0000\u0000\u0000"+
		"&\u00e4\u0001\u0000\u0000\u0000(\u00ea\u0001\u0000\u0000\u0000*\u00ee"+
		"\u0001\u0000\u0000\u0000,\u00f1\u0001\u0000\u0000\u0000.\u00fb\u0001\u0000"+
		"\u0000\u00000\u0108\u0001\u0000\u0000\u00002\u010a\u0001\u0000\u0000\u0000"+
		"4\u011b\u0001\u0000\u0000\u00006\u011d\u0001\u0000\u0000\u00008\u0131"+
		"\u0001\u0000\u0000\u0000:\u0133\u0001\u0000\u0000\u0000<\u013b\u0001\u0000"+
		"\u0000\u0000>\u013e\u0001\u0000\u0000\u0000@\u014f\u0001\u0000\u0000\u0000"+
		"B\u0151\u0001\u0000\u0000\u0000D\u0154\u0001\u0000\u0000\u0000F\u0195"+
		"\u0001\u0000\u0000\u0000H\u01ce\u0001\u0000\u0000\u0000J\u01d2\u0001\u0000"+
		"\u0000\u0000L\u01d6\u0001\u0000\u0000\u0000N\u01d8\u0001\u0000\u0000\u0000"+
		"P\u01e3\u0001\u0000\u0000\u0000R\u01ee\u0001\u0000\u0000\u0000TU\u0003"+
		"\u0002\u0001\u0000UV\u0005\u0000\u0000\u0001V\u0001\u0001\u0000\u0000"+
		"\u0000W\\\u0006\u0001\uffff\uffff\u0000XY\n\u0001\u0000\u0000Y[\u0003"+
		"\u0004\u0002\u0000ZX\u0001\u0000\u0000\u0000[^\u0001\u0000\u0000\u0000"+
		"\\Z\u0001\u0000\u0000\u0000\\]\u0001\u0000\u0000\u0000]\u0003\u0001\u0000"+
		"\u0000\u0000^\\\u0001\u0000\u0000\u0000_`\u0003\u0006\u0003\u0000`\u0005"+
		"\u0001\u0000\u0000\u0000ab\u0003H$\u0000bc\u0005:\u0000\u0000cd\u0005"+
		"\u001e\u0000\u0000de\u0003\b\u0004\u0000ef\u0005\u001f\u0000\u0000fg\u0003"+
		"N\'\u0000gh\u0003\u000e\u0007\u0000hr\u0001\u0000\u0000\u0000ij\u0003"+
		"H$\u0000jk\u0005:\u0000\u0000kl\u0005\u001e\u0000\u0000lm\u0003\b\u0004"+
		"\u0000mn\u0005\u001f\u0000\u0000no\u0003N\'\u0000op\u0003\u000e\u0007"+
		"\u0000pr\u0001\u0000\u0000\u0000qa\u0001\u0000\u0000\u0000qi\u0001\u0000"+
		"\u0000\u0000r\u0007\u0001\u0000\u0000\u0000sv\u0001\u0000\u0000\u0000"+
		"tv\u0003\n\u0005\u0000us\u0001\u0000\u0000\u0000ut\u0001\u0000\u0000\u0000"+
		"v\t\u0001\u0000\u0000\u0000wx\u0006\u0005\uffff\uffff\u0000xy\u0003\f"+
		"\u0006\u0000y\u007f\u0001\u0000\u0000\u0000z{\n\u0001\u0000\u0000{|\u0005"+
		"\u0002\u0000\u0000|~\u0003\f\u0006\u0000}z\u0001\u0000\u0000\u0000~\u0081"+
		"\u0001\u0000\u0000\u0000\u007f}\u0001\u0000\u0000\u0000\u007f\u0080\u0001"+
		"\u0000\u0000\u0000\u0080\u000b\u0001\u0000\u0000\u0000\u0081\u007f\u0001"+
		"\u0000\u0000\u0000\u0082\u008d\u0003<\u001e\u0000\u0083\u008d\u0003@ "+
		"\u0000\u0084\u0085\u0005\"\u0000\u0000\u0085\u0086\u0003\b\u0004\u0000"+
		"\u0086\u0087\u0005#\u0000\u0000\u0087\u008d\u0001\u0000\u0000\u0000\u0088"+
		"\u0089\u0005 \u0000\u0000\u0089\u008a\u0003\b\u0004\u0000\u008a\u008b"+
		"\u0005!\u0000\u0000\u008b\u008d\u0001\u0000\u0000\u0000\u008c\u0082\u0001"+
		"\u0000\u0000\u0000\u008c\u0083\u0001\u0000\u0000\u0000\u008c\u0084\u0001"+
		"\u0000\u0000\u0000\u008c\u0088\u0001\u0000\u0000\u0000\u008d\r\u0001\u0000"+
		"\u0000\u0000\u008e\u0092\u0005\"\u0000\u0000\u008f\u0091\u0003\u0012\t"+
		"\u0000\u0090\u008f\u0001\u0000\u0000\u0000\u0091\u0094\u0001\u0000\u0000"+
		"\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0093\u0001\u0000\u0000"+
		"\u0000\u0093\u0095\u0001\u0000\u0000\u0000\u0094\u0092\u0001\u0000\u0000"+
		"\u0000\u0095\u0096\u0005#\u0000\u0000\u0096\u000f\u0001\u0000\u0000\u0000"+
		"\u0097\u009c\u0006\b\uffff\uffff\u0000\u0098\u0099\n\u0001\u0000\u0000"+
		"\u0099\u009b\u0003\u0012\t\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009b"+
		"\u009e\u0001\u0000\u0000\u0000\u009c\u009a\u0001\u0000\u0000\u0000\u009c"+
		"\u009d\u0001\u0000\u0000\u0000\u009d\u0011\u0001\u0000\u0000\u0000\u009e"+
		"\u009c\u0001\u0000\u0000\u0000\u009f\u00a4\u00030\u0018\u0000\u00a0\u00a4"+
		"\u0003B!\u0000\u00a1\u00a4\u0003\u001c\u000e\u0000\u00a2\u00a4\u0003\u0014"+
		"\n\u0000\u00a3\u009f\u0001\u0000\u0000\u0000\u00a3\u00a0\u0001\u0000\u0000"+
		"\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a2\u0001\u0000\u0000"+
		"\u0000\u00a4\u0013\u0001\u0000\u0000\u0000\u00a5\u00a9\u0003\u0016\u000b"+
		"\u0000\u00a6\u00a9\u0003\u0018\f\u0000\u00a7\u00a9\u0003\u001a\r\u0000"+
		"\u00a8\u00a5\u0001\u0000\u0000\u0000\u00a8\u00a6\u0001\u0000\u0000\u0000"+
		"\u00a8\u00a7\u0001\u0000\u0000\u0000\u00a9\u0015\u0001\u0000\u0000\u0000"+
		"\u00aa\u00ab\u0005\u0011\u0000\u0000\u00ab\u00ac\u0003F#\u0000\u00ac\u00ad"+
		"\u0005\u0005\u0000\u0000\u00ad\u0017\u0001\u0000\u0000\u0000\u00ae\u00af"+
		"\u0005\u000f\u0000\u0000\u00af\u00b0\u0005\u0005\u0000\u0000\u00b0\u0019"+
		"\u0001\u0000\u0000\u0000\u00b1\u00b2\u0005\u0010\u0000\u0000\u00b2\u00b3"+
		"\u0005\u0005\u0000\u0000\u00b3\u001b\u0001\u0000\u0000\u0000\u00b4\u00b8"+
		"\u0003\u001e\u000f\u0000\u00b5\u00b8\u0003\"\u0011\u0000\u00b6\u00b8\u0003"+
		"$\u0012\u0000\u00b7\u00b4\u0001\u0000\u0000\u0000\u00b7\u00b5\u0001\u0000"+
		"\u0000\u0000\u00b7\u00b6\u0001\u0000\u0000\u0000\u00b8\u001d\u0001\u0000"+
		"\u0000\u0000\u00b9\u00ba\u0005\b\u0000\u0000\u00ba\u00bb\u0003,\u0016"+
		"\u0000\u00bb\u00bc\u0003F#\u0000\u00bc\u00bd\u0003\u000e\u0007\u0000\u00bd"+
		"\u00c5\u0001\u0000\u0000\u0000\u00be\u00bf\u0005\b\u0000\u0000\u00bf\u00c0"+
		"\u0003,\u0016\u0000\u00c0\u00c1\u0003F#\u0000\u00c1\u00c2\u0003\u000e"+
		"\u0007\u0000\u00c2\u00c3\u0003 \u0010\u0000\u00c3\u00c5\u0001\u0000\u0000"+
		"\u0000\u00c4\u00b9\u0001\u0000\u0000\u0000\u00c4\u00be\u0001\u0000\u0000"+
		"\u0000\u00c5\u001f\u0001\u0000\u0000\u0000\u00c6\u00c7\u0005\t\u0000\u0000"+
		"\u00c7\u00cb\u0003\u000e\u0007\u0000\u00c8\u00c9\u0005\t\u0000\u0000\u00c9"+
		"\u00cb\u0003\u001e\u000f\u0000\u00ca\u00c6\u0001\u0000\u0000\u0000\u00ca"+
		"\u00c8\u0001\u0000\u0000\u0000\u00cb!\u0001\u0000\u0000\u0000\u00cc\u00cd"+
		"\u0005\n\u0000\u0000\u00cd\u00ce\u0003,\u0016\u0000\u00ce\u00cf\u0003"+
		"F#\u0000\u00cf\u00d0\u0003\u000e\u0007\u0000\u00d0\u00d9\u0001\u0000\u0000"+
		"\u0000\u00d1\u00d2\u0005\n\u0000\u0000\u00d2\u00d3\u0003,\u0016\u0000"+
		"\u00d3\u00d4\u0003F#\u0000\u00d4\u00d5\u0005\u0007\u0000\u0000\u00d5\u00d6"+
		"\u0003,\u0016\u0000\u00d6\u00d7\u0003\u000e\u0007\u0000\u00d7\u00d9\u0001"+
		"\u0000\u0000\u0000\u00d8\u00cc\u0001\u0000\u0000\u0000\u00d8\u00d1\u0001"+
		"\u0000\u0000\u0000\u00d9#\u0001\u0000\u0000\u0000\u00da\u00db\u0005\u000b"+
		"\u0000\u0000\u00db\u00dc\u0003,\u0016\u0000\u00dc\u00dd\u0003F#\u0000"+
		"\u00dd\u00de\u0005\"\u0000\u0000\u00de\u00df\u0003&\u0013\u0000\u00df"+
		"\u00e0\u0005#\u0000\u0000\u00e0%\u0001\u0000\u0000\u0000\u00e1\u00e3\u0003"+
		"(\u0014\u0000\u00e2\u00e1\u0001\u0000\u0000\u0000\u00e3\u00e6\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e2\u0001\u0000\u0000\u0000\u00e4\u00e5\u0001\u0000"+
		"\u0000\u0000\u00e5\u00e8\u0001\u0000\u0000\u0000\u00e6\u00e4\u0001\u0000"+
		"\u0000\u0000\u00e7\u00e9\u0003*\u0015\u0000\u00e8\u00e7\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9\'\u0001\u0000\u0000\u0000"+
		"\u00ea\u00eb\u0005\f\u0000\u0000\u00eb\u00ec\u0003F#\u0000\u00ec\u00ed"+
		"\u0003\u000e\u0007\u0000\u00ed)\u0001\u0000\u0000\u0000\u00ee\u00ef\u0005"+
		"\r\u0000\u0000\u00ef\u00f0\u0003\u000e\u0007\u0000\u00f0+\u0001\u0000"+
		"\u0000\u0000\u00f1\u00f6\u0006\u0016\uffff\uffff\u0000\u00f2\u00f3\n\u0001"+
		"\u0000\u0000\u00f3\u00f5\u0003.\u0017\u0000\u00f4\u00f2\u0001\u0000\u0000"+
		"\u0000\u00f5\u00f8\u0001\u0000\u0000\u0000\u00f6\u00f4\u0001\u0000\u0000"+
		"\u0000\u00f6\u00f7\u0001\u0000\u0000\u0000\u00f7-\u0001\u0000\u0000\u0000"+
		"\u00f8\u00f6\u0001\u0000\u0000\u0000\u00f9\u00fc\u00030\u0018\u0000\u00fa"+
		"\u00fc\u0003B!\u0000\u00fb\u00f9\u0001\u0000\u0000\u0000\u00fb\u00fa\u0001"+
		"\u0000\u0000\u0000\u00fc/\u0001\u0000\u0000\u0000\u00fd\u00fe\u00036\u001b"+
		"\u0000\u00fe\u00ff\u0005\u0006\u0000\u0000\u00ff\u0100\u0003D\"\u0000"+
		"\u0100\u0101\u0005\u0005\u0000\u0000\u0101\u0109\u0001\u0000\u0000\u0000"+
		"\u0102\u0103\u00032\u0019\u0000\u0103\u0104\u0007\u0000\u0000\u0000\u0104"+
		"\u0105\u0005\u0006\u0000\u0000\u0105\u0106\u0003D\"\u0000\u0106\u0107"+
		"\u0005\u0005\u0000\u0000\u0107\u0109\u0001\u0000\u0000\u0000\u0108\u00fd"+
		"\u0001\u0000\u0000\u0000\u0108\u0102\u0001\u0000\u0000\u0000\u01091\u0001"+
		"\u0000\u0000\u0000\u010a\u010f\u00034\u001a\u0000\u010b\u010c\u0005\u0002"+
		"\u0000\u0000\u010c\u010e\u00034\u001a\u0000\u010d\u010b\u0001\u0000\u0000"+
		"\u0000\u010e\u0111\u0001\u0000\u0000\u0000\u010f\u010d\u0001\u0000\u0000"+
		"\u0000\u010f\u0110\u0001\u0000\u0000\u0000\u01103\u0001\u0000\u0000\u0000"+
		"\u0111\u010f\u0001\u0000\u0000\u0000\u0112\u011c\u0003F#\u0000\u0113\u0114"+
		"\u0005\"\u0000\u0000\u0114\u0115\u00032\u0019\u0000\u0115\u0116\u0005"+
		"#\u0000\u0000\u0116\u011c\u0001\u0000\u0000\u0000\u0117\u0118\u0005 \u0000"+
		"\u0000\u0118\u0119\u00032\u0019\u0000\u0119\u011a\u0005!\u0000\u0000\u011a"+
		"\u011c\u0001\u0000\u0000\u0000\u011b\u0112\u0001\u0000\u0000\u0000\u011b"+
		"\u0113\u0001\u0000\u0000\u0000\u011b\u0117\u0001\u0000\u0000\u0000\u011c"+
		"5\u0001\u0000\u0000\u0000\u011d\u0122\u00038\u001c\u0000\u011e\u011f\u0005"+
		"\u0002\u0000\u0000\u011f\u0121\u00030\u0018\u0000\u0120\u011e\u0001\u0000"+
		"\u0000\u0000\u0121\u0124\u0001\u0000\u0000\u0000\u0122\u0120\u0001\u0000"+
		"\u0000\u0000\u0122\u0123\u0001\u0000\u0000\u0000\u01237\u0001\u0000\u0000"+
		"\u0000\u0124\u0122\u0001\u0000\u0000\u0000\u0125\u0132\u0003F#\u0000\u0126"+
		"\u0132\u0003<\u001e\u0000\u0127\u0128\u0005\u0016\u0000\u0000\u0128\u0132"+
		"\u0003@ \u0000\u0129\u012a\u0005\"\u0000\u0000\u012a\u012b\u00036\u001b"+
		"\u0000\u012b\u012c\u0005#\u0000\u0000\u012c\u0132\u0001\u0000\u0000\u0000"+
		"\u012d\u012e\u0005 \u0000\u0000\u012e\u012f\u00036\u001b\u0000\u012f\u0130"+
		"\u0005!\u0000\u0000\u0130\u0132\u0001\u0000\u0000\u0000\u0131\u0125\u0001"+
		"\u0000\u0000\u0000\u0131\u0126\u0001\u0000\u0000\u0000\u0131\u0127\u0001"+
		"\u0000\u0000\u0000\u0131\u0129\u0001\u0000\u0000\u0000\u0131\u012d\u0001"+
		"\u0000\u0000\u0000\u01329\u0001\u0000\u0000\u0000\u0133\u0138\u0003<\u001e"+
		"\u0000\u0134\u0135\u0005\u0002\u0000\u0000\u0135\u0137\u0003<\u001e\u0000"+
		"\u0136\u0134\u0001\u0000\u0000\u0000\u0137\u013a\u0001\u0000\u0000\u0000"+
		"\u0138\u0136\u0001\u0000\u0000\u0000\u0138\u0139\u0001\u0000\u0000\u0000"+
		"\u0139;\u0001\u0000\u0000\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013b"+
		"\u013c\u0003N\'\u0000\u013c\u013d\u0003@ \u0000\u013d=\u0001\u0000\u0000"+
		"\u0000\u013e\u0143\u0003@ \u0000\u013f\u0140\u0005\u0002\u0000\u0000\u0140"+
		"\u0142\u0003@ \u0000\u0141\u013f\u0001\u0000\u0000\u0000\u0142\u0145\u0001"+
		"\u0000\u0000\u0000\u0143\u0141\u0001\u0000\u0000\u0000\u0143\u0144\u0001"+
		"\u0000\u0000\u0000\u0144?\u0001\u0000\u0000\u0000\u0145\u0143\u0001\u0000"+
		"\u0000\u0000\u0146\u0150\u0005:\u0000\u0000\u0147\u0148\u0005\"\u0000"+
		"\u0000\u0148\u0149\u0003>\u001f\u0000\u0149\u014a\u0005#\u0000\u0000\u014a"+
		"\u0150\u0001\u0000\u0000\u0000\u014b\u014c\u0005 \u0000\u0000\u014c\u014d"+
		"\u0003>\u001f\u0000\u014d\u014e\u0005!\u0000\u0000\u014e\u0150\u0001\u0000"+
		"\u0000\u0000\u014f\u0146\u0001\u0000\u0000\u0000\u014f\u0147\u0001\u0000"+
		"\u0000\u0000\u014f\u014b\u0001\u0000\u0000\u0000\u0150A\u0001\u0000\u0000"+
		"\u0000\u0151\u0152\u0003D\"\u0000\u0152\u0153\u0005\u0005\u0000\u0000"+
		"\u0153C\u0001\u0000\u0000\u0000\u0154\u0159\u0003F#\u0000\u0155\u0156"+
		"\u0005\u0002\u0000\u0000\u0156\u0158\u0003F#\u0000\u0157\u0155\u0001\u0000"+
		"\u0000\u0000\u0158\u015b\u0001\u0000\u0000\u0000\u0159\u0157\u0001\u0000"+
		"\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015aE\u0001\u0000\u0000"+
		"\u0000\u015b\u0159\u0001\u0000\u0000\u0000\u015c\u015d\u0006#\uffff\uffff"+
		"\u0000\u015d\u015e\u0005\u001e\u0000\u0000\u015e\u015f\u0003N\'\u0000"+
		"\u015f\u0160\u0005\u001f\u0000\u0000\u0160\u0161\u0003F#\u0014\u0161\u0196"+
		"\u0001\u0000\u0000\u0000\u0162\u0163\u0005\u001e\u0000\u0000\u0163\u0164"+
		"\u0003F#\u0000\u0164\u0165\u0005\u001f\u0000\u0000\u0165\u0196\u0001\u0000"+
		"\u0000\u0000\u0166\u0167\u0007\u0001\u0000\u0000\u0167\u0196\u0003F#\u0011"+
		"\u0168\u0196\u0005:\u0000\u0000\u0169\u0196\u0003R)\u0000\u016a\u016b"+
		"\u0003H$\u0000\u016b\u016d\u0005\u001e\u0000\u0000\u016c\u016e\u00036"+
		"\u001b\u0000\u016d\u016c\u0001\u0000\u0000\u0000\u016d\u016e\u0001\u0000"+
		"\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000\u016f\u0171\u0005\u001f"+
		"\u0000\u0000\u0170\u0172\u0003N\'\u0000\u0171\u0170\u0001\u0000\u0000"+
		"\u0000\u0171\u0172\u0001\u0000\u0000\u0000\u0172\u0173\u0001\u0000\u0000"+
		"\u0000\u0173\u0174\u0003\u000e\u0007\u0000\u0174\u0196\u0001\u0000\u0000"+
		"\u0000\u0175\u0176\u0003J%\u0000\u0176\u0177\u0005\"\u0000\u0000\u0177"+
		"\u0178\u0005:\u0000\u0000\u0178\u0179\u0005\u0003\u0000\u0000\u0179\u0180"+
		"\u0003F#\u0000\u017a\u017b\u0005\u0002\u0000\u0000\u017b\u017c\u0005:"+
		"\u0000\u0000\u017c\u017d\u0005\u0003\u0000\u0000\u017d\u017f\u0003F#\u0000"+
		"\u017e\u017a\u0001\u0000\u0000\u0000\u017f\u0182\u0001\u0000\u0000\u0000"+
		"\u0180\u017e\u0001\u0000\u0000\u0000\u0180\u0181\u0001\u0000\u0000\u0000"+
		"\u0181\u0184\u0001\u0000\u0000\u0000\u0182\u0180\u0001\u0000\u0000\u0000"+
		"\u0183\u0185\u0005\u0002\u0000\u0000\u0184\u0183\u0001\u0000\u0000\u0000"+
		"\u0184\u0185\u0001\u0000\u0000\u0000\u0185\u0186\u0001\u0000\u0000\u0000"+
		"\u0186\u0187\u0005#\u0000\u0000\u0187\u0196\u0001\u0000\u0000\u0000\u0188"+
		"\u0189\u0005\u001e\u0000\u0000\u0189\u0191\u0003F#\u0000\u018a\u0192\u0005"+
		"\u0002\u0000\u0000\u018b\u018c\u0005\u0002\u0000\u0000\u018c\u018e\u0003"+
		"F#\u0000\u018d\u018b\u0001\u0000\u0000\u0000\u018e\u018f\u0001\u0000\u0000"+
		"\u0000\u018f\u018d\u0001\u0000\u0000\u0000\u018f\u0190\u0001\u0000\u0000"+
		"\u0000\u0190\u0192\u0001\u0000\u0000\u0000\u0191\u018a\u0001\u0000\u0000"+
		"\u0000\u0191\u018d\u0001\u0000\u0000\u0000\u0192\u0193\u0001\u0000\u0000"+
		"\u0000\u0193\u0194\u0005\u001f\u0000\u0000\u0194\u0196\u0001\u0000\u0000"+
		"\u0000\u0195\u015c\u0001\u0000\u0000\u0000\u0195\u0162\u0001\u0000\u0000"+
		"\u0000\u0195\u0166\u0001\u0000\u0000\u0000\u0195\u0168\u0001\u0000\u0000"+
		"\u0000\u0195\u0169\u0001\u0000\u0000\u0000\u0195\u016a\u0001\u0000\u0000"+
		"\u0000\u0195\u0175\u0001\u0000\u0000\u0000\u0195\u0188\u0001\u0000\u0000"+
		"\u0000\u0196\u01c9\u0001\u0000\u0000\u0000\u0197\u0198\n\u0012\u0000\u0000"+
		"\u0198\u0199\u0007\u0002\u0000\u0000\u0199\u01c8\u0003F#\u0013\u019a\u019b"+
		"\n\u0010\u0000\u0000\u019b\u019c\u0007\u0003\u0000\u0000\u019c\u01c8\u0003"+
		"F#\u0011\u019d\u019e\n\u000f\u0000\u0000\u019e\u019f\u0007\u0004\u0000"+
		"\u0000\u019f\u01c8\u0003F#\u0010\u01a0\u01a1\n\u000e\u0000\u0000\u01a1"+
		"\u01a2\u0007\u0005\u0000\u0000\u01a2\u01c8\u0003F#\u000f\u01a3\u01a4\n"+
		"\r\u0000\u0000\u01a4\u01a5\u0007\u0006\u0000\u0000\u01a5\u01c8\u0003F"+
		"#\u000e\u01a6\u01a7\n\f\u0000\u0000\u01a7\u01a8\u0005.\u0000\u0000\u01a8"+
		"\u01c8\u0003F#\r\u01a9\u01aa\n\u000b\u0000\u0000\u01aa\u01ab\u0005-\u0000"+
		"\u0000\u01ab\u01c8\u0003F#\f\u01ac\u01ad\n\n\u0000\u0000\u01ad\u01ae\u0005"+
		",\u0000\u0000\u01ae\u01c8\u0003F#\u000b\u01af\u01b0\n\t\u0000\u0000\u01b0"+
		"\u01b2\u0005\u001e\u0000\u0000\u01b1\u01b3\u0003D\"\u0000\u01b2\u01b1"+
		"\u0001\u0000\u0000\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b3\u01b4"+
		"\u0001\u0000\u0000\u0000\u01b4\u01c8\u0005\u001f\u0000\u0000\u01b5\u01b6"+
		"\n\b\u0000\u0000\u01b6\u01b7\u0005\u0001\u0000\u0000\u01b7\u01c8\u0005"+
		":\u0000\u0000\u01b8\u01b9\n\u0007\u0000\u0000\u01b9\u01ba\u0005 \u0000"+
		"\u0000\u01ba\u01bb\u0003F#\u0000\u01bb\u01bc\u0005!\u0000\u0000\u01bc"+
		"\u01c8\u0001\u0000\u0000\u0000\u01bd\u01be\n\u0006\u0000\u0000\u01be\u01c0"+
		"\u0005 \u0000\u0000\u01bf\u01c1\u0003F#\u0000\u01c0\u01bf\u0001\u0000"+
		"\u0000\u0000\u01c0\u01c1\u0001\u0000\u0000\u0000\u01c1\u01c2\u0001\u0000"+
		"\u0000\u0000\u01c2\u01c4\u0005\u0003\u0000\u0000\u01c3\u01c5\u0003F#\u0000"+
		"\u01c4\u01c3\u0001\u0000\u0000\u0000\u01c4\u01c5\u0001\u0000\u0000\u0000"+
		"\u01c5\u01c6\u0001\u0000\u0000\u0000\u01c6\u01c8\u0005!\u0000\u0000\u01c7"+
		"\u0197\u0001\u0000\u0000\u0000\u01c7\u019a\u0001\u0000\u0000\u0000\u01c7"+
		"\u019d\u0001\u0000\u0000\u0000\u01c7\u01a0\u0001\u0000\u0000\u0000\u01c7"+
		"\u01a3\u0001\u0000\u0000\u0000\u01c7\u01a6\u0001\u0000\u0000\u0000\u01c7"+
		"\u01a9\u0001\u0000\u0000\u0000\u01c7\u01ac\u0001\u0000\u0000\u0000\u01c7"+
		"\u01af\u0001\u0000\u0000\u0000\u01c7\u01b5\u0001\u0000\u0000\u0000\u01c7"+
		"\u01b8\u0001\u0000\u0000\u0000\u01c7\u01bd\u0001\u0000\u0000\u0000\u01c8"+
		"\u01cb\u0001\u0000\u0000\u0000\u01c9\u01c7\u0001\u0000\u0000\u0000\u01c9"+
		"\u01ca\u0001\u0000\u0000\u0000\u01caG\u0001\u0000\u0000\u0000\u01cb\u01c9"+
		"\u0001\u0000\u0000\u0000\u01cc\u01cf\u0005\u0019\u0000\u0000\u01cd\u01cf"+
		"\u0003N\'\u0000\u01ce\u01cc\u0001\u0000\u0000\u0000\u01ce\u01cd\u0001"+
		"\u0000\u0000\u0000\u01cfI\u0001\u0000\u0000\u0000\u01d0\u01d3\u0005\u001a"+
		"\u0000\u0000\u01d1\u01d3\u0003N\'\u0000\u01d2\u01d0\u0001\u0000\u0000"+
		"\u0000\u01d2\u01d1\u0001\u0000\u0000\u0000\u01d3K\u0001\u0000\u0000\u0000"+
		"\u01d4\u01d7\u0005\u001b\u0000\u0000\u01d5\u01d7\u0003N\'\u0000\u01d6"+
		"\u01d4\u0001\u0000\u0000\u0000\u01d6\u01d5\u0001\u0000\u0000\u0000\u01d7"+
		"M\u0001\u0000\u0000\u0000\u01d8\u01d9\u0006\'\uffff\uffff\u0000\u01d9"+
		"\u01da\u0005:\u0000\u0000\u01da\u01e0\u0001\u0000\u0000\u0000\u01db\u01dc"+
		"\n\u0001\u0000\u0000\u01dc\u01dd\u0005\u0001\u0000\u0000\u01dd\u01df\u0005"+
		":\u0000\u0000\u01de\u01db\u0001\u0000\u0000\u0000\u01df\u01e2\u0001\u0000"+
		"\u0000\u0000\u01e0\u01de\u0001\u0000\u0000\u0000\u01e0\u01e1\u0001\u0000"+
		"\u0000\u0000\u01e1O\u0001\u0000\u0000\u0000\u01e2\u01e0\u0001\u0000\u0000"+
		"\u0000\u01e3\u01e8\u0005:\u0000\u0000\u01e4\u01e5\u0005\u0002\u0000\u0000"+
		"\u01e5\u01e7\u0005:\u0000\u0000\u01e6\u01e4\u0001\u0000\u0000\u0000\u01e7"+
		"\u01ea\u0001\u0000\u0000\u0000\u01e8\u01e6\u0001\u0000\u0000\u0000\u01e8"+
		"\u01e9\u0001\u0000\u0000\u0000\u01e9Q\u0001\u0000\u0000\u0000\u01ea\u01e8"+
		"\u0001\u0000\u0000\u0000\u01eb\u01ef\u0005;\u0000\u0000\u01ec\u01ed\u0005"+
		":\u0000\u0000\u01ed\u01ef\u0005;\u0000\u0000\u01ee\u01eb\u0001\u0000\u0000"+
		"\u0000\u01ee\u01ec\u0001\u0000\u0000\u0000\u01efS\u0001\u0000\u0000\u0000"+
		",\\qu\u007f\u008c\u0092\u009c\u00a3\u00a8\u00b7\u00c4\u00ca\u00d8\u00e4"+
		"\u00e8\u00f6\u00fb\u0108\u010f\u011b\u0122\u0131\u0138\u0143\u014f\u0159"+
		"\u016d\u0171\u0180\u0184\u018f\u0191\u0195\u01b2\u01c0\u01c4\u01c7\u01c9"+
		"\u01ce\u01d2\u01d6\u01e0\u01e8\u01ee";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
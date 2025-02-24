// Code generated from CaliburnParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing // CaliburnParser
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

var CaliburnParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func caliburnparserParserInit() {
	staticData := &CaliburnParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "','", "':'", "'?'", "';'", "'='", "'->'", "'if'", "'else'",
		"'for'", "'switch'", "'case'", "'default'", "'fallthrough'", "'break'",
		"'continue'", "'return'", "'using'", "'as'", "'in'", "'null'", "'var'",
		"'const'", "'type'", "'func'", "'struct'", "'tuple'", "'class'", "'interface'",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'+'", "'-'", "'!'", "'*'",
		"'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", "'&'", "'<<'", "'>>'", "'=='",
		"'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "PERIOD", "COMMA", "COLON", "QUESTION", "Terminator", "ASSIGN",
		"ARROW", "IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH",
		"BREAK", "CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST",
		"TYPE", "FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN",
		"R_PAREN", "L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD",
		"OP_SUB", "OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT",
		"OP_OR", "OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ",
		"OP_GTE", "OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "CommentMultiLine",
		"CommentSingleLine", "Identifier", "Literal",
	}
	staticData.RuleNames = []string{
		"module", "definitions", "definition", "function_definition", "parameters",
		"parameters_list", "parameter", "block", "statements", "statement",
		"jump_statement", "return_statement", "break_statement", "continue_statement",
		"control_statement", "if_statement", "else_statement", "for_statement",
		"switch_statement", "case_blocks", "option_case_block", "default_case_block",
		"inline_statements", "inline_statement", "assign_statement", "assign_expressions",
		"assign_expression", "assign_declarations", "assign_declaration", "typed_assign_declarations",
		"typed_assign_declaration", "untyped_assign_declarations", "untyped_assign_declaration",
		"expression_statement", "expressions", "expression", "function_type",
		"struct_type", "tuple_type", "type_expression", "identifiers", "literal_atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 497, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 1,
		0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 91, 8, 1, 10, 1, 12, 1, 94, 9, 1,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 114, 8, 3, 1, 4, 1, 4, 3, 4,
		118, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 126, 8, 5, 10, 5,
		12, 5, 129, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 141, 8, 6, 1, 7, 1, 7, 5, 7, 145, 8, 7, 10, 7, 12, 7, 148,
		9, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 155, 8, 8, 10, 8, 12, 8, 158,
		9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 164, 8, 9, 1, 10, 1, 10, 1, 10, 3,
		10, 169, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 3, 14, 184, 8, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 197,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 203, 8, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17,
		217, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 5,
		19, 227, 8, 19, 10, 19, 12, 19, 230, 9, 19, 1, 19, 3, 19, 233, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 5, 22,
		245, 8, 22, 10, 22, 12, 22, 248, 9, 22, 1, 23, 1, 23, 3, 23, 252, 8, 23,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 3, 24, 265, 8, 24, 1, 25, 1, 25, 1, 25, 5, 25, 270, 8, 25, 10, 25,
		12, 25, 273, 9, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 3, 26, 284, 8, 26, 1, 27, 1, 27, 1, 27, 5, 27, 289, 8, 27, 10,
		27, 12, 27, 292, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 306, 8, 28, 1, 29, 1, 29, 1,
		29, 5, 29, 311, 8, 29, 10, 29, 12, 29, 314, 9, 29, 1, 30, 1, 30, 1, 30,
		1, 31, 1, 31, 1, 31, 5, 31, 322, 8, 31, 10, 31, 12, 31, 325, 9, 31, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 336,
		8, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 5, 34, 344, 8, 34, 10,
		34, 12, 34, 347, 9, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3,
		35, 366, 8, 35, 1, 35, 1, 35, 3, 35, 370, 8, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5, 35, 383, 8, 35,
		10, 35, 12, 35, 386, 9, 35, 1, 35, 3, 35, 389, 8, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 4, 35, 398, 8, 35, 11, 35, 12, 35, 399,
		3, 35, 402, 8, 35, 1, 35, 1, 35, 3, 35, 406, 8, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 3, 35, 435, 8, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 449, 8, 35,
		1, 35, 1, 35, 3, 35, 453, 8, 35, 1, 35, 5, 35, 456, 8, 35, 10, 35, 12,
		35, 459, 9, 35, 1, 36, 1, 36, 3, 36, 463, 8, 36, 1, 37, 1, 37, 3, 37, 467,
		8, 37, 1, 38, 1, 38, 3, 38, 471, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 5, 39, 479, 8, 39, 10, 39, 12, 39, 482, 9, 39, 1, 40, 1, 40,
		1, 40, 5, 40, 487, 8, 40, 10, 40, 12, 40, 490, 9, 40, 1, 41, 1, 41, 1,
		41, 3, 41, 495, 8, 41, 1, 41, 0, 6, 2, 10, 16, 44, 70, 78, 42, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
		80, 82, 0, 7, 2, 0, 36, 37, 39, 48, 1, 0, 36, 38, 1, 0, 42, 43, 1, 0, 39,
		41, 1, 0, 36, 37, 1, 0, 47, 48, 1, 0, 49, 54, 525, 0, 84, 1, 0, 0, 0, 2,
		87, 1, 0, 0, 0, 4, 95, 1, 0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 117, 1, 0, 0,
		0, 10, 119, 1, 0, 0, 0, 12, 140, 1, 0, 0, 0, 14, 142, 1, 0, 0, 0, 16, 151,
		1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 168, 1, 0, 0, 0, 22, 170, 1, 0, 0,
		0, 24, 174, 1, 0, 0, 0, 26, 177, 1, 0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 196,
		1, 0, 0, 0, 32, 202, 1, 0, 0, 0, 34, 216, 1, 0, 0, 0, 36, 218, 1, 0, 0,
		0, 38, 228, 1, 0, 0, 0, 40, 234, 1, 0, 0, 0, 42, 238, 1, 0, 0, 0, 44, 241,
		1, 0, 0, 0, 46, 251, 1, 0, 0, 0, 48, 264, 1, 0, 0, 0, 50, 266, 1, 0, 0,
		0, 52, 283, 1, 0, 0, 0, 54, 285, 1, 0, 0, 0, 56, 305, 1, 0, 0, 0, 58, 307,
		1, 0, 0, 0, 60, 315, 1, 0, 0, 0, 62, 318, 1, 0, 0, 0, 64, 335, 1, 0, 0,
		0, 66, 337, 1, 0, 0, 0, 68, 340, 1, 0, 0, 0, 70, 405, 1, 0, 0, 0, 72, 462,
		1, 0, 0, 0, 74, 466, 1, 0, 0, 0, 76, 470, 1, 0, 0, 0, 78, 472, 1, 0, 0,
		0, 80, 483, 1, 0, 0, 0, 82, 494, 1, 0, 0, 0, 84, 85, 3, 2, 1, 0, 85, 86,
		5, 0, 0, 1, 86, 1, 1, 0, 0, 0, 87, 92, 6, 1, -1, 0, 88, 89, 10, 1, 0, 0,
		89, 91, 3, 4, 2, 0, 90, 88, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90, 1,
		0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 3, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95,
		96, 3, 6, 3, 0, 96, 5, 1, 0, 0, 0, 97, 98, 3, 72, 36, 0, 98, 99, 5, 58,
		0, 0, 99, 100, 5, 30, 0, 0, 100, 101, 3, 8, 4, 0, 101, 102, 5, 31, 0, 0,
		102, 103, 3, 78, 39, 0, 103, 104, 3, 14, 7, 0, 104, 114, 1, 0, 0, 0, 105,
		106, 3, 72, 36, 0, 106, 107, 5, 58, 0, 0, 107, 108, 5, 30, 0, 0, 108, 109,
		3, 8, 4, 0, 109, 110, 5, 31, 0, 0, 110, 111, 3, 78, 39, 0, 111, 112, 3,
		14, 7, 0, 112, 114, 1, 0, 0, 0, 113, 97, 1, 0, 0, 0, 113, 105, 1, 0, 0,
		0, 114, 7, 1, 0, 0, 0, 115, 118, 1, 0, 0, 0, 116, 118, 3, 10, 5, 0, 117,
		115, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 9, 1, 0, 0, 0, 119, 120, 6,
		5, -1, 0, 120, 121, 3, 12, 6, 0, 121, 127, 1, 0, 0, 0, 122, 123, 10, 1,
		0, 0, 123, 124, 5, 2, 0, 0, 124, 126, 3, 12, 6, 0, 125, 122, 1, 0, 0, 0,
		126, 129, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		11, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0, 130, 141, 3, 60, 30, 0, 131, 141,
		3, 64, 32, 0, 132, 133, 5, 34, 0, 0, 133, 134, 3, 8, 4, 0, 134, 135, 5,
		35, 0, 0, 135, 141, 1, 0, 0, 0, 136, 137, 5, 32, 0, 0, 137, 138, 3, 8,
		4, 0, 138, 139, 5, 33, 0, 0, 139, 141, 1, 0, 0, 0, 140, 130, 1, 0, 0, 0,
		140, 131, 1, 0, 0, 0, 140, 132, 1, 0, 0, 0, 140, 136, 1, 0, 0, 0, 141,
		13, 1, 0, 0, 0, 142, 146, 5, 34, 0, 0, 143, 145, 3, 18, 9, 0, 144, 143,
		1, 0, 0, 0, 145, 148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 149, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 5, 35, 0, 0,
		150, 15, 1, 0, 0, 0, 151, 156, 6, 8, -1, 0, 152, 153, 10, 1, 0, 0, 153,
		155, 3, 18, 9, 0, 154, 152, 1, 0, 0, 0, 155, 158, 1, 0, 0, 0, 156, 154,
		1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 17, 1, 0, 0, 0, 158, 156, 1, 0,
		0, 0, 159, 164, 3, 48, 24, 0, 160, 164, 3, 66, 33, 0, 161, 164, 3, 28,
		14, 0, 162, 164, 3, 20, 10, 0, 163, 159, 1, 0, 0, 0, 163, 160, 1, 0, 0,
		0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 19, 1, 0, 0, 0, 165,
		169, 3, 22, 11, 0, 166, 169, 3, 24, 12, 0, 167, 169, 3, 26, 13, 0, 168,
		165, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 21, 1,
		0, 0, 0, 170, 171, 5, 17, 0, 0, 171, 172, 3, 70, 35, 0, 172, 173, 5, 5,
		0, 0, 173, 23, 1, 0, 0, 0, 174, 175, 5, 15, 0, 0, 175, 176, 5, 5, 0, 0,
		176, 25, 1, 0, 0, 0, 177, 178, 5, 16, 0, 0, 178, 179, 5, 5, 0, 0, 179,
		27, 1, 0, 0, 0, 180, 184, 3, 30, 15, 0, 181, 184, 3, 34, 17, 0, 182, 184,
		3, 36, 18, 0, 183, 180, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 183, 182, 1,
		0, 0, 0, 184, 29, 1, 0, 0, 0, 185, 186, 5, 8, 0, 0, 186, 187, 3, 44, 22,
		0, 187, 188, 3, 70, 35, 0, 188, 189, 3, 14, 7, 0, 189, 197, 1, 0, 0, 0,
		190, 191, 5, 8, 0, 0, 191, 192, 3, 44, 22, 0, 192, 193, 3, 70, 35, 0, 193,
		194, 3, 14, 7, 0, 194, 195, 3, 32, 16, 0, 195, 197, 1, 0, 0, 0, 196, 185,
		1, 0, 0, 0, 196, 190, 1, 0, 0, 0, 197, 31, 1, 0, 0, 0, 198, 199, 5, 9,
		0, 0, 199, 203, 3, 14, 7, 0, 200, 201, 5, 9, 0, 0, 201, 203, 3, 30, 15,
		0, 202, 198, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 33, 1, 0, 0, 0, 204,
		205, 5, 10, 0, 0, 205, 206, 3, 44, 22, 0, 206, 207, 3, 70, 35, 0, 207,
		208, 3, 14, 7, 0, 208, 217, 1, 0, 0, 0, 209, 210, 5, 10, 0, 0, 210, 211,
		3, 44, 22, 0, 211, 212, 3, 70, 35, 0, 212, 213, 5, 7, 0, 0, 213, 214, 3,
		44, 22, 0, 214, 215, 3, 14, 7, 0, 215, 217, 1, 0, 0, 0, 216, 204, 1, 0,
		0, 0, 216, 209, 1, 0, 0, 0, 217, 35, 1, 0, 0, 0, 218, 219, 5, 11, 0, 0,
		219, 220, 3, 44, 22, 0, 220, 221, 3, 70, 35, 0, 221, 222, 5, 34, 0, 0,
		222, 223, 3, 38, 19, 0, 223, 224, 5, 35, 0, 0, 224, 37, 1, 0, 0, 0, 225,
		227, 3, 40, 20, 0, 226, 225, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226,
		1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0,
		0, 0, 231, 233, 3, 42, 21, 0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0,
		0, 233, 39, 1, 0, 0, 0, 234, 235, 5, 12, 0, 0, 235, 236, 3, 70, 35, 0,
		236, 237, 3, 14, 7, 0, 237, 41, 1, 0, 0, 0, 238, 239, 5, 13, 0, 0, 239,
		240, 3, 14, 7, 0, 240, 43, 1, 0, 0, 0, 241, 246, 6, 22, -1, 0, 242, 243,
		10, 1, 0, 0, 243, 245, 3, 46, 23, 0, 244, 242, 1, 0, 0, 0, 245, 248, 1,
		0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 45, 1, 0, 0,
		0, 248, 246, 1, 0, 0, 0, 249, 252, 3, 48, 24, 0, 250, 252, 3, 66, 33, 0,
		251, 249, 1, 0, 0, 0, 251, 250, 1, 0, 0, 0, 252, 47, 1, 0, 0, 0, 253, 254,
		3, 54, 27, 0, 254, 255, 5, 6, 0, 0, 255, 256, 3, 68, 34, 0, 256, 257, 5,
		5, 0, 0, 257, 265, 1, 0, 0, 0, 258, 259, 3, 50, 25, 0, 259, 260, 7, 0,
		0, 0, 260, 261, 5, 6, 0, 0, 261, 262, 3, 68, 34, 0, 262, 263, 5, 5, 0,
		0, 263, 265, 1, 0, 0, 0, 264, 253, 1, 0, 0, 0, 264, 258, 1, 0, 0, 0, 265,
		49, 1, 0, 0, 0, 266, 271, 3, 52, 26, 0, 267, 268, 5, 2, 0, 0, 268, 270,
		3, 52, 26, 0, 269, 267, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1,
		0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 51, 1, 0, 0, 0, 273, 271, 1, 0, 0,
		0, 274, 284, 3, 70, 35, 0, 275, 276, 5, 34, 0, 0, 276, 277, 3, 50, 25,
		0, 277, 278, 5, 35, 0, 0, 278, 284, 1, 0, 0, 0, 279, 280, 5, 32, 0, 0,
		280, 281, 3, 50, 25, 0, 281, 282, 5, 33, 0, 0, 282, 284, 1, 0, 0, 0, 283,
		274, 1, 0, 0, 0, 283, 275, 1, 0, 0, 0, 283, 279, 1, 0, 0, 0, 284, 53, 1,
		0, 0, 0, 285, 290, 3, 56, 28, 0, 286, 287, 5, 2, 0, 0, 287, 289, 3, 48,
		24, 0, 288, 286, 1, 0, 0, 0, 289, 292, 1, 0, 0, 0, 290, 288, 1, 0, 0, 0,
		290, 291, 1, 0, 0, 0, 291, 55, 1, 0, 0, 0, 292, 290, 1, 0, 0, 0, 293, 306,
		3, 70, 35, 0, 294, 306, 3, 60, 30, 0, 295, 296, 5, 22, 0, 0, 296, 306,
		3, 64, 32, 0, 297, 298, 5, 34, 0, 0, 298, 299, 3, 54, 27, 0, 299, 300,
		5, 35, 0, 0, 300, 306, 1, 0, 0, 0, 301, 302, 5, 32, 0, 0, 302, 303, 3,
		54, 27, 0, 303, 304, 5, 33, 0, 0, 304, 306, 1, 0, 0, 0, 305, 293, 1, 0,
		0, 0, 305, 294, 1, 0, 0, 0, 305, 295, 1, 0, 0, 0, 305, 297, 1, 0, 0, 0,
		305, 301, 1, 0, 0, 0, 306, 57, 1, 0, 0, 0, 307, 312, 3, 60, 30, 0, 308,
		309, 5, 2, 0, 0, 309, 311, 3, 60, 30, 0, 310, 308, 1, 0, 0, 0, 311, 314,
		1, 0, 0, 0, 312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 59, 1, 0,
		0, 0, 314, 312, 1, 0, 0, 0, 315, 316, 3, 78, 39, 0, 316, 317, 3, 64, 32,
		0, 317, 61, 1, 0, 0, 0, 318, 323, 3, 64, 32, 0, 319, 320, 5, 2, 0, 0, 320,
		322, 3, 64, 32, 0, 321, 319, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321,
		1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 63, 1, 0, 0, 0, 325, 323, 1, 0,
		0, 0, 326, 336, 5, 58, 0, 0, 327, 328, 5, 34, 0, 0, 328, 329, 3, 62, 31,
		0, 329, 330, 5, 35, 0, 0, 330, 336, 1, 0, 0, 0, 331, 332, 5, 32, 0, 0,
		332, 333, 3, 62, 31, 0, 333, 334, 5, 33, 0, 0, 334, 336, 1, 0, 0, 0, 335,
		326, 1, 0, 0, 0, 335, 327, 1, 0, 0, 0, 335, 331, 1, 0, 0, 0, 336, 65, 1,
		0, 0, 0, 337, 338, 3, 68, 34, 0, 338, 339, 5, 5, 0, 0, 339, 67, 1, 0, 0,
		0, 340, 345, 3, 70, 35, 0, 341, 342, 5, 2, 0, 0, 342, 344, 3, 70, 35, 0,
		343, 341, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345,
		346, 1, 0, 0, 0, 346, 69, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 349, 6,
		35, -1, 0, 349, 350, 5, 30, 0, 0, 350, 351, 3, 78, 39, 0, 351, 352, 5,
		31, 0, 0, 352, 353, 3, 70, 35, 20, 353, 406, 1, 0, 0, 0, 354, 355, 5, 30,
		0, 0, 355, 356, 3, 70, 35, 0, 356, 357, 5, 31, 0, 0, 357, 406, 1, 0, 0,
		0, 358, 359, 7, 1, 0, 0, 359, 406, 3, 70, 35, 17, 360, 406, 5, 58, 0, 0,
		361, 406, 3, 82, 41, 0, 362, 363, 3, 72, 36, 0, 363, 365, 5, 30, 0, 0,
		364, 366, 3, 54, 27, 0, 365, 364, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366,
		367, 1, 0, 0, 0, 367, 369, 5, 31, 0, 0, 368, 370, 3, 78, 39, 0, 369, 368,
		1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 372, 3, 14,
		7, 0, 372, 406, 1, 0, 0, 0, 373, 374, 3, 74, 37, 0, 374, 375, 5, 34, 0,
		0, 375, 376, 5, 58, 0, 0, 376, 377, 5, 3, 0, 0, 377, 384, 3, 70, 35, 0,
		378, 379, 5, 2, 0, 0, 379, 380, 5, 58, 0, 0, 380, 381, 5, 3, 0, 0, 381,
		383, 3, 70, 35, 0, 382, 378, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382,
		1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 384, 1, 0,
		0, 0, 387, 389, 5, 2, 0, 0, 388, 387, 1, 0, 0, 0, 388, 389, 1, 0, 0, 0,
		389, 390, 1, 0, 0, 0, 390, 391, 5, 35, 0, 0, 391, 406, 1, 0, 0, 0, 392,
		393, 5, 30, 0, 0, 393, 401, 3, 70, 35, 0, 394, 402, 5, 2, 0, 0, 395, 396,
		5, 2, 0, 0, 396, 398, 3, 70, 35, 0, 397, 395, 1, 0, 0, 0, 398, 399, 1,
		0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 402, 1, 0, 0,
		0, 401, 394, 1, 0, 0, 0, 401, 397, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403,
		404, 5, 31, 0, 0, 404, 406, 1, 0, 0, 0, 405, 348, 1, 0, 0, 0, 405, 354,
		1, 0, 0, 0, 405, 358, 1, 0, 0, 0, 405, 360, 1, 0, 0, 0, 405, 361, 1, 0,
		0, 0, 405, 362, 1, 0, 0, 0, 405, 373, 1, 0, 0, 0, 405, 392, 1, 0, 0, 0,
		406, 457, 1, 0, 0, 0, 407, 408, 10, 18, 0, 0, 408, 409, 7, 2, 0, 0, 409,
		456, 3, 70, 35, 19, 410, 411, 10, 16, 0, 0, 411, 412, 7, 3, 0, 0, 412,
		456, 3, 70, 35, 17, 413, 414, 10, 15, 0, 0, 414, 415, 7, 4, 0, 0, 415,
		456, 3, 70, 35, 16, 416, 417, 10, 14, 0, 0, 417, 418, 7, 5, 0, 0, 418,
		456, 3, 70, 35, 15, 419, 420, 10, 13, 0, 0, 420, 421, 7, 6, 0, 0, 421,
		456, 3, 70, 35, 14, 422, 423, 10, 12, 0, 0, 423, 424, 5, 46, 0, 0, 424,
		456, 3, 70, 35, 13, 425, 426, 10, 11, 0, 0, 426, 427, 5, 45, 0, 0, 427,
		456, 3, 70, 35, 12, 428, 429, 10, 10, 0, 0, 429, 430, 5, 44, 0, 0, 430,
		456, 3, 70, 35, 11, 431, 432, 10, 9, 0, 0, 432, 434, 5, 30, 0, 0, 433,
		435, 3, 68, 34, 0, 434, 433, 1, 0, 0, 0, 434, 435, 1, 0, 0, 0, 435, 436,
		1, 0, 0, 0, 436, 456, 5, 31, 0, 0, 437, 438, 10, 8, 0, 0, 438, 439, 5,
		1, 0, 0, 439, 456, 5, 58, 0, 0, 440, 441, 10, 7, 0, 0, 441, 442, 5, 32,
		0, 0, 442, 443, 3, 70, 35, 0, 443, 444, 5, 33, 0, 0, 444, 456, 1, 0, 0,
		0, 445, 446, 10, 6, 0, 0, 446, 448, 5, 32, 0, 0, 447, 449, 3, 70, 35, 0,
		448, 447, 1, 0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450,
		452, 5, 3, 0, 0, 451, 453, 3, 70, 35, 0, 452, 451, 1, 0, 0, 0, 452, 453,
		1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 456, 5, 33, 0, 0, 455, 407, 1, 0,
		0, 0, 455, 410, 1, 0, 0, 0, 455, 413, 1, 0, 0, 0, 455, 416, 1, 0, 0, 0,
		455, 419, 1, 0, 0, 0, 455, 422, 1, 0, 0, 0, 455, 425, 1, 0, 0, 0, 455,
		428, 1, 0, 0, 0, 455, 431, 1, 0, 0, 0, 455, 437, 1, 0, 0, 0, 455, 440,
		1, 0, 0, 0, 455, 445, 1, 0, 0, 0, 456, 459, 1, 0, 0, 0, 457, 455, 1, 0,
		0, 0, 457, 458, 1, 0, 0, 0, 458, 71, 1, 0, 0, 0, 459, 457, 1, 0, 0, 0,
		460, 463, 5, 25, 0, 0, 461, 463, 3, 78, 39, 0, 462, 460, 1, 0, 0, 0, 462,
		461, 1, 0, 0, 0, 463, 73, 1, 0, 0, 0, 464, 467, 5, 26, 0, 0, 465, 467,
		3, 78, 39, 0, 466, 464, 1, 0, 0, 0, 466, 465, 1, 0, 0, 0, 467, 75, 1, 0,
		0, 0, 468, 471, 5, 27, 0, 0, 469, 471, 3, 78, 39, 0, 470, 468, 1, 0, 0,
		0, 470, 469, 1, 0, 0, 0, 471, 77, 1, 0, 0, 0, 472, 473, 6, 39, -1, 0, 473,
		474, 5, 58, 0, 0, 474, 480, 1, 0, 0, 0, 475, 476, 10, 1, 0, 0, 476, 477,
		5, 1, 0, 0, 477, 479, 5, 58, 0, 0, 478, 475, 1, 0, 0, 0, 479, 482, 1, 0,
		0, 0, 480, 478, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 79, 1, 0, 0, 0,
		482, 480, 1, 0, 0, 0, 483, 488, 5, 58, 0, 0, 484, 485, 5, 2, 0, 0, 485,
		487, 5, 58, 0, 0, 486, 484, 1, 0, 0, 0, 487, 490, 1, 0, 0, 0, 488, 486,
		1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489, 81, 1, 0, 0, 0, 490, 488, 1, 0,
		0, 0, 491, 495, 5, 59, 0, 0, 492, 493, 5, 58, 0, 0, 493, 495, 5, 59, 0,
		0, 494, 491, 1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 495, 83, 1, 0, 0, 0, 44,
		92, 113, 117, 127, 140, 146, 156, 163, 168, 183, 196, 202, 216, 228, 232,
		246, 251, 264, 271, 283, 290, 305, 312, 323, 335, 345, 365, 369, 384, 388,
		399, 401, 405, 434, 448, 452, 455, 457, 462, 466, 470, 480, 488, 494,
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
	staticData := &CaliburnParserParserStaticData
	staticData.once.Do(caliburnparserParserInit)
}

// NewCaliburnParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCaliburnParser(input antlr.TokenStream) *CaliburnParser {
	CaliburnParserInit()
	this := new(CaliburnParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CaliburnParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "CaliburnParser.g4"

	return this
}

// CaliburnParser tokens.
const (
	CaliburnParserEOF               = antlr.TokenEOF
	CaliburnParserPERIOD            = 1
	CaliburnParserCOMMA             = 2
	CaliburnParserCOLON             = 3
	CaliburnParserQUESTION          = 4
	CaliburnParserTerminator        = 5
	CaliburnParserASSIGN            = 6
	CaliburnParserARROW             = 7
	CaliburnParserIF                = 8
	CaliburnParserELSE              = 9
	CaliburnParserFOR               = 10
	CaliburnParserSWITCH            = 11
	CaliburnParserCASE              = 12
	CaliburnParserDEFAULT           = 13
	CaliburnParserFALLTHROUGH       = 14
	CaliburnParserBREAK             = 15
	CaliburnParserCONTINUE          = 16
	CaliburnParserRETURN            = 17
	CaliburnParserUSING             = 18
	CaliburnParserAS                = 19
	CaliburnParserIN                = 20
	CaliburnParserNULL              = 21
	CaliburnParserVAR               = 22
	CaliburnParserCONST             = 23
	CaliburnParserTYPE              = 24
	CaliburnParserFUNC              = 25
	CaliburnParserSTRUCT            = 26
	CaliburnParserTUPLE             = 27
	CaliburnParserCLASS             = 28
	CaliburnParserINTERFACE         = 29
	CaliburnParserL_PAREN           = 30
	CaliburnParserR_PAREN           = 31
	CaliburnParserL_S_BRACK         = 32
	CaliburnParserR_S_BRACK         = 33
	CaliburnParserL_C_BRACK         = 34
	CaliburnParserR_C_BRACK         = 35
	CaliburnParserOP_ADD            = 36
	CaliburnParserOP_SUB            = 37
	CaliburnParserOP_NOT            = 38
	CaliburnParserOP_MUL            = 39
	CaliburnParserOP_DIV            = 40
	CaliburnParserOP_MOD            = 41
	CaliburnParserOP_POW            = 42
	CaliburnParserOP_ROOT           = 43
	CaliburnParserOP_OR             = 44
	CaliburnParserOP_XOR            = 45
	CaliburnParserOP_AND            = 46
	CaliburnParserOP_LSHIFT         = 47
	CaliburnParserOP_RSHIFT         = 48
	CaliburnParserOP_EQU            = 49
	CaliburnParserOP_NEQ            = 50
	CaliburnParserOP_GTE            = 51
	CaliburnParserOP_GRT            = 52
	CaliburnParserOP_LTE            = 53
	CaliburnParserOP_LST            = 54
	CaliburnParserWhiteSpace        = 55
	CaliburnParserCommentMultiLine  = 56
	CaliburnParserCommentSingleLine = 57
	CaliburnParserIdentifier        = 58
	CaliburnParserLiteral           = 59
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_module                      = 0
	CaliburnParserRULE_definitions                 = 1
	CaliburnParserRULE_definition                  = 2
	CaliburnParserRULE_function_definition         = 3
	CaliburnParserRULE_parameters                  = 4
	CaliburnParserRULE_parameters_list             = 5
	CaliburnParserRULE_parameter                   = 6
	CaliburnParserRULE_block                       = 7
	CaliburnParserRULE_statements                  = 8
	CaliburnParserRULE_statement                   = 9
	CaliburnParserRULE_jump_statement              = 10
	CaliburnParserRULE_return_statement            = 11
	CaliburnParserRULE_break_statement             = 12
	CaliburnParserRULE_continue_statement          = 13
	CaliburnParserRULE_control_statement           = 14
	CaliburnParserRULE_if_statement                = 15
	CaliburnParserRULE_else_statement              = 16
	CaliburnParserRULE_for_statement               = 17
	CaliburnParserRULE_switch_statement            = 18
	CaliburnParserRULE_case_blocks                 = 19
	CaliburnParserRULE_option_case_block           = 20
	CaliburnParserRULE_default_case_block          = 21
	CaliburnParserRULE_inline_statements           = 22
	CaliburnParserRULE_inline_statement            = 23
	CaliburnParserRULE_assign_statement            = 24
	CaliburnParserRULE_assign_expressions          = 25
	CaliburnParserRULE_assign_expression           = 26
	CaliburnParserRULE_assign_declarations         = 27
	CaliburnParserRULE_assign_declaration          = 28
	CaliburnParserRULE_typed_assign_declarations   = 29
	CaliburnParserRULE_typed_assign_declaration    = 30
	CaliburnParserRULE_untyped_assign_declarations = 31
	CaliburnParserRULE_untyped_assign_declaration  = 32
	CaliburnParserRULE_expression_statement        = 33
	CaliburnParserRULE_expressions                 = 34
	CaliburnParserRULE_expression                  = 35
	CaliburnParserRULE_function_type               = 36
	CaliburnParserRULE_struct_type                 = 37
	CaliburnParserRULE_tuple_type                  = 38
	CaliburnParserRULE_type_expression             = 39
	CaliburnParserRULE_identifiers                 = 40
	CaliburnParserRULE_literal_atom                = 41
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsModuleContext differentiates from other interfaces.
	IsModuleContext()
}

type ModuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_module
	return p
}

func InitEmptyModuleContext(p *ModuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_module
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleContext) CopyAll(ctx *ModuleContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ModuleRuleContext struct {
	ModuleContext
}

func NewModuleRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleRuleContext {
	var p = new(ModuleRuleContext)

	InitEmptyModuleContext(&p.ModuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*ModuleContext))

	return p
}

func (s *ModuleRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleRuleContext) Definitions() IDefinitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionsContext)
}

func (s *ModuleRuleContext) EOF() antlr.TerminalNode {
	return s.GetToken(CaliburnParserEOF, 0)
}

func (s *ModuleRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterModuleRule(s)
	}
}

func (s *ModuleRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitModuleRule(s)
	}
}

func (s *ModuleRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitModuleRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CaliburnParserRULE_module)
	localctx = NewModuleRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(84)
		p.definitions(0)
	}
	{
		p.SetState(85)
		p.Match(CaliburnParserEOF)
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

// IDefinitionsContext is an interface to support dynamic dispatch.
type IDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefinitionsContext differentiates from other interfaces.
	IsDefinitionsContext()
}

type DefinitionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionsContext() *DefinitionsContext {
	var p = new(DefinitionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_definitions
	return p
}

func InitEmptyDefinitionsContext(p *DefinitionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_definitions
}

func (*DefinitionsContext) IsDefinitionsContext() {}

func NewDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionsContext {
	var p = new(DefinitionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_definitions

	return p
}

func (s *DefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionsContext) CopyAll(ctx *DefinitionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefinitionsAppendContext struct {
	DefinitionsContext
}

func NewDefinitionsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefinitionsAppendContext {
	var p = new(DefinitionsAppendContext)

	InitEmptyDefinitionsContext(&p.DefinitionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DefinitionsContext))

	return p
}

func (s *DefinitionsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsAppendContext) Definitions() IDefinitionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionsContext)
}

func (s *DefinitionsAppendContext) Definition() IDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefinitionsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDefinitionsAppend(s)
	}
}

func (s *DefinitionsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDefinitionsAppend(s)
	}
}

func (s *DefinitionsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDefinitionsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefinitionsInitialContext struct {
	DefinitionsContext
}

func NewDefinitionsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefinitionsInitialContext {
	var p = new(DefinitionsInitialContext)

	InitEmptyDefinitionsContext(&p.DefinitionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*DefinitionsContext))

	return p
}

func (s *DefinitionsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDefinitionsInitial(s)
	}
}

func (s *DefinitionsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDefinitionsInitial(s)
	}
}

func (s *DefinitionsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDefinitionsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Definitions() (localctx IDefinitionsContext) {
	return p.definitions(0)
}

func (p *CaliburnParser) definitions(_p int) (localctx IDefinitionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewDefinitionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDefinitionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CaliburnParserRULE_definitions, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewDefinitionsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDefinitionsAppendContext(p, NewDefinitionsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_definitions)
			p.SetState(88)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(89)
				p.Definition()
			}

		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_definition() IFunction_definitionContext

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_definition
	return p
}

func InitEmptyDefinitionContext(p *DefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_definition
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Function_definition() IFunction_definitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_definitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_definitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CaliburnParserRULE_definition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Function_definition()
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

// IFunction_definitionContext is an interface to support dynamic dispatch.
type IFunction_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunction_definitionContext differentiates from other interfaces.
	IsFunction_definitionContext()
}

type Function_definitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_definitionContext() *Function_definitionContext {
	var p = new(Function_definitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_definition
	return p
}

func InitEmptyFunction_definitionContext(p *Function_definitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_definition
}

func (*Function_definitionContext) IsFunction_definitionContext() {}

func NewFunction_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_definitionContext {
	var p = new(Function_definitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_function_definition

	return p
}

func (s *Function_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_definitionContext) CopyAll(ctx *Function_definitionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Function_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionDefinitionNoReturnTypeContext struct {
	Function_definitionContext
	type_       IFunction_typeContext
	name        antlr.Token
	return_type IType_expressionContext
}

func NewFunctionDefinitionNoReturnTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefinitionNoReturnTypeContext {
	var p = new(FunctionDefinitionNoReturnTypeContext)

	InitEmptyFunction_definitionContext(&p.Function_definitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_definitionContext))

	return p
}

func (s *FunctionDefinitionNoReturnTypeContext) GetName() antlr.Token { return s.name }

func (s *FunctionDefinitionNoReturnTypeContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDefinitionNoReturnTypeContext) GetType_() IFunction_typeContext { return s.type_ }

func (s *FunctionDefinitionNoReturnTypeContext) GetReturn_type() IType_expressionContext {
	return s.return_type
}

func (s *FunctionDefinitionNoReturnTypeContext) SetType_(v IFunction_typeContext) { s.type_ = v }

func (s *FunctionDefinitionNoReturnTypeContext) SetReturn_type(v IType_expressionContext) {
	s.return_type = v
}

func (s *FunctionDefinitionNoReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionNoReturnTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *FunctionDefinitionNoReturnTypeContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDefinitionNoReturnTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *FunctionDefinitionNoReturnTypeContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefinitionNoReturnTypeContext) Function_type() IFunction_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *FunctionDefinitionNoReturnTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *FunctionDefinitionNoReturnTypeContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *FunctionDefinitionNoReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionDefinitionNoReturnType(s)
	}
}

func (s *FunctionDefinitionNoReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionDefinitionNoReturnType(s)
	}
}

func (s *FunctionDefinitionNoReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionDefinitionNoReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionDefinitionContext struct {
	Function_definitionContext
	type_       IFunction_typeContext
	name        antlr.Token
	return_type IType_expressionContext
}

func NewFunctionDefinitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	InitEmptyFunction_definitionContext(&p.Function_definitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_definitionContext))

	return p
}

func (s *FunctionDefinitionContext) GetName() antlr.Token { return s.name }

func (s *FunctionDefinitionContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDefinitionContext) GetType_() IFunction_typeContext { return s.type_ }

func (s *FunctionDefinitionContext) GetReturn_type() IType_expressionContext { return s.return_type }

func (s *FunctionDefinitionContext) SetType_(v IFunction_typeContext) { s.type_ = v }

func (s *FunctionDefinitionContext) SetReturn_type(v IType_expressionContext) { s.return_type = v }

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *FunctionDefinitionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDefinitionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *FunctionDefinitionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefinitionContext) Function_type() IFunction_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *FunctionDefinitionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *FunctionDefinitionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Function_definition() (localctx IFunction_definitionContext) {
	localctx = NewFunction_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CaliburnParserRULE_function_definition)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionDefinitionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(97)

			var _x = p.Function_type()

			localctx.(*FunctionDefinitionContext).type_ = _x
		}
		{
			p.SetState(98)

			var _m = p.Match(CaliburnParserIdentifier)

			localctx.(*FunctionDefinitionContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Parameters()
		}
		{
			p.SetState(101)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)

			var _x = p.type_expression(0)

			localctx.(*FunctionDefinitionContext).return_type = _x
		}
		{
			p.SetState(103)
			p.Block()
		}

	case 2:
		localctx = NewFunctionDefinitionNoReturnTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)

			var _x = p.Function_type()

			localctx.(*FunctionDefinitionNoReturnTypeContext).type_ = _x
		}
		{
			p.SetState(106)

			var _m = p.Match(CaliburnParserIdentifier)

			localctx.(*FunctionDefinitionNoReturnTypeContext).name = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Parameters()
		}
		{
			p.SetState(109)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)

			var _x = p.type_expression(0)

			localctx.(*FunctionDefinitionNoReturnTypeContext).return_type = _x
		}
		{
			p.SetState(111)
			p.Block()
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

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) CopyAll(ctx *ParametersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametersEmptyContext struct {
	ParametersContext
}

func NewParametersEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersEmptyContext {
	var p = new(ParametersEmptyContext)

	InitEmptyParametersContext(&p.ParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametersContext))

	return p
}

func (s *ParametersEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterParametersEmpty(s)
	}
}

func (s *ParametersEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitParametersEmpty(s)
	}
}

func (s *ParametersEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitParametersEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParametersFilledContext struct {
	ParametersContext
}

func NewParametersFilledContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersFilledContext {
	var p = new(ParametersFilledContext)

	InitEmptyParametersContext(&p.ParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametersContext))

	return p
}

func (s *ParametersFilledContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersFilledContext) Parameters_list() IParameters_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameters_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameters_listContext)
}

func (s *ParametersFilledContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterParametersFilled(s)
	}
}

func (s *ParametersFilledContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitParametersFilled(s)
	}
}

func (s *ParametersFilledContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitParametersFilled(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CaliburnParserRULE_parameters)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserR_PAREN, CaliburnParserR_S_BRACK, CaliburnParserR_C_BRACK:
		localctx = NewParametersEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)

	case CaliburnParserL_S_BRACK, CaliburnParserL_C_BRACK, CaliburnParserIdentifier:
		localctx = NewParametersFilledContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.parameters_list(0)
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

// IParameters_listContext is an interface to support dynamic dispatch.
type IParameters_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParameters_listContext differentiates from other interfaces.
	IsParameters_listContext()
}

type Parameters_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameters_listContext() *Parameters_listContext {
	var p = new(Parameters_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameters_list
	return p
}

func InitEmptyParameters_listContext(p *Parameters_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameters_list
}

func (*Parameters_listContext) IsParameters_listContext() {}

func NewParameters_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameters_listContext {
	var p = new(Parameters_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_parameters_list

	return p
}

func (s *Parameters_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameters_listContext) CopyAll(ctx *Parameters_listContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Parameters_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameters_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParametersListInitialContext struct {
	Parameters_listContext
}

func NewParametersListInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersListInitialContext {
	var p = new(ParametersListInitialContext)

	InitEmptyParameters_listContext(&p.Parameters_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parameters_listContext))

	return p
}

func (s *ParametersListInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersListInitialContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersListInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterParametersListInitial(s)
	}
}

func (s *ParametersListInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitParametersListInitial(s)
	}
}

func (s *ParametersListInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitParametersListInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParametersListAppendContext struct {
	Parameters_listContext
}

func NewParametersListAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersListAppendContext {
	var p = new(ParametersListAppendContext)

	InitEmptyParameters_listContext(&p.Parameters_listContext)
	p.parser = parser
	p.CopyAll(ctx.(*Parameters_listContext))

	return p
}

func (s *ParametersListAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersListAppendContext) Parameters_list() IParameters_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameters_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameters_listContext)
}

func (s *ParametersListAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *ParametersListAppendContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersListAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterParametersListAppend(s)
	}
}

func (s *ParametersListAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitParametersListAppend(s)
	}
}

func (s *ParametersListAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitParametersListAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Parameters_list() (localctx IParameters_listContext) {
	return p.parameters_list(0)
}

func (p *CaliburnParser) parameters_list(_p int) (localctx IParameters_listContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewParameters_listContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParameters_listContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, CaliburnParserRULE_parameters_list, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewParametersListInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(120)
		p.Parameter()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParametersListAppendContext(p, NewParameters_listContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_parameters_list)
			p.SetState(122)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(123)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(124)
				p.Parameter()
			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) CopyAll(ctx *ParameterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructDestrutureParameterContext struct {
	ParameterContext
}

func NewStructDestrutureParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDestrutureParameterContext {
	var p = new(StructDestrutureParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *StructDestrutureParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDestrutureParameterContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *StructDestrutureParameterContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *StructDestrutureParameterContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *StructDestrutureParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructDestrutureParameter(s)
	}
}

func (s *StructDestrutureParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructDestrutureParameter(s)
	}
}

func (s *StructDestrutureParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructDestrutureParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleDestrutureParameterContext struct {
	ParameterContext
}

func NewTupleDestrutureParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleDestrutureParameterContext {
	var p = new(TupleDestrutureParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *TupleDestrutureParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleDestrutureParameterContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *TupleDestrutureParameterContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *TupleDestrutureParameterContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *TupleDestrutureParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleDestrutureParameter(s)
	}
}

func (s *TupleDestrutureParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleDestrutureParameter(s)
	}
}

func (s *TupleDestrutureParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleDestrutureParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedParameterContext struct {
	ParameterContext
}

func NewTypedParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedParameterContext {
	var p = new(TypedParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *TypedParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedParameterContext) Typed_assign_declaration() ITyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITyped_assign_declarationContext)
}

func (s *TypedParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedParameter(s)
	}
}

func (s *TypedParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedParameter(s)
	}
}

func (s *TypedParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedParameterContext struct {
	ParameterContext
}

func NewUntypedParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedParameterContext {
	var p = new(UntypedParameterContext)

	InitEmptyParameterContext(&p.ParameterContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParameterContext))

	return p
}

func (s *UntypedParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedParameterContext) Untyped_assign_declaration() IUntyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationContext)
}

func (s *UntypedParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedParameter(s)
	}
}

func (s *UntypedParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedParameter(s)
	}
}

func (s *UntypedParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CaliburnParserRULE_parameter)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Typed_assign_declaration()
		}

	case 2:
		localctx = NewUntypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Untyped_assign_declaration()
		}

	case 3:
		localctx = NewStructDestrutureParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Parameters()
		}
		{
			p.SetState(134)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTupleDestrutureParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.Parameters()
		}
		{
			p.SetState(138)
			p.Match(CaliburnParserR_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) CopyAll(ctx *BlockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BlockRuleContext struct {
	BlockContext
}

func NewBlockRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockRuleContext {
	var p = new(BlockRuleContext)

	InitEmptyBlockContext(&p.BlockContext)
	p.parser = parser
	p.CopyAll(ctx.(*BlockContext))

	return p
}

func (s *BlockRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockRuleContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *BlockRuleContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *BlockRuleContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockRuleContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBlockRule(s)
	}
}

func (s *BlockRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBlockRule(s)
	}
}

func (s *BlockRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBlockRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CaliburnParserRULE_block)
	var _la int

	localctx = NewBlockRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(142)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864691632145140992) != 0 {
		{
			p.SetState(143)
			p.Statement()
		}

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)
		p.Match(CaliburnParserR_C_BRACK)
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

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) CopyAll(ctx *StatementsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatementsAppendContext struct {
	StatementsContext
}

func NewStatementsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementsAppendContext {
	var p = new(StatementsAppendContext)

	InitEmptyStatementsContext(&p.StatementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementsContext))

	return p
}

func (s *StatementsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsAppendContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *StatementsAppendContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStatementsAppend(s)
	}
}

func (s *StatementsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStatementsAppend(s)
	}
}

func (s *StatementsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStatementsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type StatementsInitialContext struct {
	StatementsContext
}

func NewStatementsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatementsInitialContext {
	var p = new(StatementsInitialContext)

	InitEmptyStatementsContext(&p.StatementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementsContext))

	return p
}

func (s *StatementsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStatementsInitial(s)
	}
}

func (s *StatementsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStatementsInitial(s)
	}
}

func (s *StatementsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStatementsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Statements() (localctx IStatementsContext) {
	return p.statements(0)
}

func (p *CaliburnParser) statements(_p int) (localctx IStatementsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, CaliburnParserRULE_statements, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewStatementsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewStatementsAppendContext(p, NewStatementsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_statements)
			p.SetState(152)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(153)
				p.Statement()
			}

		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign_statement() IAssign_statementContext
	Expression_statement() IExpression_statementContext
	Control_statement() IControl_statementContext
	Jump_statement() IJump_statementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assign_statement() IAssign_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_statementContext)
}

func (s *StatementContext) Expression_statement() IExpression_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *StatementContext) Control_statement() IControl_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IControl_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IControl_statementContext)
}

func (s *StatementContext) Jump_statement() IJump_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJump_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJump_statementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CaliburnParserRULE_statement)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Expression_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Jump_statement()
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

// IJump_statementContext is an interface to support dynamic dispatch.
type IJump_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Return_statement() IReturn_statementContext
	Break_statement() IBreak_statementContext
	Continue_statement() IContinue_statementContext

	// IsJump_statementContext differentiates from other interfaces.
	IsJump_statementContext()
}

type Jump_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJump_statementContext() *Jump_statementContext {
	var p = new(Jump_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_jump_statement
	return p
}

func InitEmptyJump_statementContext(p *Jump_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_jump_statement
}

func (*Jump_statementContext) IsJump_statementContext() {}

func NewJump_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Jump_statementContext {
	var p = new(Jump_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_jump_statement

	return p
}

func (s *Jump_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Jump_statementContext) Return_statement() IReturn_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturn_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturn_statementContext)
}

func (s *Jump_statementContext) Break_statement() IBreak_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreak_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreak_statementContext)
}

func (s *Jump_statementContext) Continue_statement() IContinue_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinue_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinue_statementContext)
}

func (s *Jump_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Jump_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Jump_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterJump_statement(s)
	}
}

func (s *Jump_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitJump_statement(s)
	}
}

func (s *Jump_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitJump_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Jump_statement() (localctx IJump_statementContext) {
	localctx = NewJump_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CaliburnParserRULE_jump_statement)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserRETURN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Return_statement()
		}

	case CaliburnParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Break_statement()
		}

	case CaliburnParserCONTINUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(167)
			p.Continue_statement()
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

// IReturn_statementContext is an interface to support dynamic dispatch.
type IReturn_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReturn_statementContext differentiates from other interfaces.
	IsReturn_statementContext()
}

type Return_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_statementContext() *Return_statementContext {
	var p = new(Return_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_return_statement
	return p
}

func InitEmptyReturn_statementContext(p *Return_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_return_statement
}

func (*Return_statementContext) IsReturn_statementContext() {}

func NewReturn_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_statementContext {
	var p = new(Return_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_return_statement

	return p
}

func (s *Return_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_statementContext) CopyAll(ctx *Return_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnStatementContext struct {
	Return_statementContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	InitEmptyReturn_statementContext(&p.Return_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Return_statementContext))

	return p
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CaliburnParserRULE_return_statement)
	localctx = NewReturnStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(CaliburnParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.expression(0)
	}
	{
		p.SetState(172)
		p.Match(CaliburnParserTerminator)
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

// IBreak_statementContext is an interface to support dynamic dispatch.
type IBreak_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBreak_statementContext differentiates from other interfaces.
	IsBreak_statementContext()
}

type Break_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_statementContext() *Break_statementContext {
	var p = new(Break_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_break_statement
	return p
}

func InitEmptyBreak_statementContext(p *Break_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_break_statement
}

func (*Break_statementContext) IsBreak_statementContext() {}

func NewBreak_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_statementContext {
	var p = new(Break_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_break_statement

	return p
}

func (s *Break_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Break_statementContext) CopyAll(ctx *Break_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	Break_statementContext
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	InitEmptyBreak_statementContext(&p.Break_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Break_statementContext))

	return p
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserBREAK, 0)
}

func (s *BreakStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CaliburnParserRULE_break_statement)
	localctx = NewBreakStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(CaliburnParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(CaliburnParserTerminator)
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

// IContinue_statementContext is an interface to support dynamic dispatch.
type IContinue_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsContinue_statementContext differentiates from other interfaces.
	IsContinue_statementContext()
}

type Continue_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_statementContext() *Continue_statementContext {
	var p = new(Continue_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_continue_statement
	return p
}

func InitEmptyContinue_statementContext(p *Continue_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_continue_statement
}

func (*Continue_statementContext) IsContinue_statementContext() {}

func NewContinue_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_statementContext {
	var p = new(Continue_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_continue_statement

	return p
}

func (s *Continue_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Continue_statementContext) CopyAll(ctx *Continue_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStatementContext struct {
	Continue_statementContext
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	InitEmptyContinue_statementContext(&p.Continue_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Continue_statementContext))

	return p
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCONTINUE, 0)
}

func (s *ContinueStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CaliburnParserRULE_continue_statement)
	localctx = NewContinueStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.Match(CaliburnParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(CaliburnParserTerminator)
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

// IControl_statementContext is an interface to support dynamic dispatch.
type IControl_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If_statement() IIf_statementContext
	For_statement() IFor_statementContext
	Switch_statement() ISwitch_statementContext

	// IsControl_statementContext differentiates from other interfaces.
	IsControl_statementContext()
}

type Control_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControl_statementContext() *Control_statementContext {
	var p = new(Control_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_control_statement
	return p
}

func InitEmptyControl_statementContext(p *Control_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_control_statement
}

func (*Control_statementContext) IsControl_statementContext() {}

func NewControl_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_statementContext {
	var p = new(Control_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_control_statement

	return p
}

func (s *Control_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_statementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *Control_statementContext) For_statement() IFor_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_statementContext)
}

func (s *Control_statementContext) Switch_statement() ISwitch_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_statementContext)
}

func (s *Control_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterControl_statement(s)
	}
}

func (s *Control_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitControl_statement(s)
	}
}

func (s *Control_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitControl_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Control_statement() (localctx IControl_statementContext) {
	localctx = NewControl_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CaliburnParserRULE_control_statement)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.If_statement()
		}

	case CaliburnParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.For_statement()
		}

	case CaliburnParserSWITCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Switch_statement()
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

// IIf_statementContext is an interface to support dynamic dispatch.
type IIf_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_statementContext differentiates from other interfaces.
	IsIf_statementContext()
}

type If_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_statementContext() *If_statementContext {
	var p = new(If_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_if_statement
	return p
}

func InitEmptyIf_statementContext(p *If_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_if_statement
}

func (*If_statementContext) IsIf_statementContext() {}

func NewIf_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_statementContext {
	var p = new(If_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_if_statement

	return p
}

func (s *If_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *If_statementContext) CopyAll(ctx *If_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStatementContext struct {
	If_statementContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIF, 0)
}

func (s *IfStatementContext) Inline_statements() IInline_statementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseStatementContext struct {
	If_statementContext
}

func NewIfElseStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseStatementContext {
	var p = new(IfElseStatementContext)

	InitEmptyIf_statementContext(&p.If_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_statementContext))

	return p
}

func (s *IfElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIF, 0)
}

func (s *IfElseStatementContext) Inline_statements() IInline_statementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *IfElseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseStatementContext) Else_statement() IElse_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElse_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElse_statementContext)
}

func (s *IfElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIfElseStatement(s)
	}
}

func (s *IfElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIfElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CaliburnParserRULE_if_statement)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Match(CaliburnParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.inline_statements(0)
		}
		{
			p.SetState(187)
			p.expression(0)
		}
		{
			p.SetState(188)
			p.Block()
		}

	case 2:
		localctx = NewIfElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Match(CaliburnParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.inline_statements(0)
		}
		{
			p.SetState(192)
			p.expression(0)
		}
		{
			p.SetState(193)
			p.Block()
		}
		{
			p.SetState(194)
			p.Else_statement()
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

// IElse_statementContext is an interface to support dynamic dispatch.
type IElse_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsElse_statementContext differentiates from other interfaces.
	IsElse_statementContext()
}

type Else_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_statementContext() *Else_statementContext {
	var p = new(Else_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_else_statement
	return p
}

func InitEmptyElse_statementContext(p *Else_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_else_statement
}

func (*Else_statementContext) IsElse_statementContext() {}

func NewElse_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_statementContext {
	var p = new(Else_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_else_statement

	return p
}

func (s *Else_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_statementContext) CopyAll(ctx *Else_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Else_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ElseIfStatementContext struct {
	Else_statementContext
}

func NewElseIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseIfStatementContext {
	var p = new(ElseIfStatementContext)

	InitEmptyElse_statementContext(&p.Else_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_statementContext))

	return p
}

func (s *ElseIfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserELSE, 0)
}

func (s *ElseIfStatementContext) If_statement() IIf_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_statementContext)
}

func (s *ElseIfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitElseIfStatement(s)
	}
}

func (s *ElseIfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitElseIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ElseStatementContext struct {
	Else_statementContext
}

func NewElseStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ElseStatementContext {
	var p = new(ElseStatementContext)

	InitEmptyElse_statementContext(&p.Else_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Else_statementContext))

	return p
}

func (s *ElseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserELSE, 0)
}

func (s *ElseStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterElseStatement(s)
	}
}

func (s *ElseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitElseStatement(s)
	}
}

func (s *ElseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitElseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Else_statement() (localctx IElse_statementContext) {
	localctx = NewElse_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CaliburnParserRULE_else_statement)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(198)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Block()
		}

	case 2:
		localctx = NewElseIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.If_statement()
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

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_statementContext differentiates from other interfaces.
	IsFor_statementContext()
}

type For_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_statementContext() *For_statementContext {
	var p = new(For_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_for_statement
	return p
}

func InitEmptyFor_statementContext(p *For_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_for_statement
}

func (*For_statementContext) IsFor_statementContext() {}

func NewFor_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_statementContext {
	var p = new(For_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_for_statement

	return p
}

func (s *For_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *For_statementContext) CopyAll(ctx *For_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForStatementContext struct {
	For_statementContext
}

func NewForStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementContext {
	var p = new(ForStatementContext)

	InitEmptyFor_statementContext(&p.For_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_statementContext))

	return p
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) FOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserFOR, 0)
}

func (s *ForStatementContext) Inline_statements() IInline_statementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *ForStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (s *ForStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitForStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStatementWithAfterContext struct {
	For_statementContext
}

func NewForStatementWithAfterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementWithAfterContext {
	var p = new(ForStatementWithAfterContext)

	InitEmptyFor_statementContext(&p.For_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_statementContext))

	return p
}

func (s *ForStatementWithAfterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementWithAfterContext) FOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserFOR, 0)
}

func (s *ForStatementWithAfterContext) AllInline_statements() []IInline_statementsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInline_statementsContext); ok {
			len++
		}
	}

	tst := make([]IInline_statementsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInline_statementsContext); ok {
			tst[i] = t.(IInline_statementsContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementWithAfterContext) Inline_statements(i int) IInline_statementsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *ForStatementWithAfterContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementWithAfterContext) ARROW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserARROW, 0)
}

func (s *ForStatementWithAfterContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementWithAfterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterForStatementWithAfter(s)
	}
}

func (s *ForStatementWithAfterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitForStatementWithAfter(s)
	}
}

func (s *ForStatementWithAfterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitForStatementWithAfter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CaliburnParserRULE_for_statement)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(CaliburnParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.inline_statements(0)
		}
		{
			p.SetState(206)
			p.expression(0)
		}
		{
			p.SetState(207)
			p.Block()
		}

	case 2:
		localctx = NewForStatementWithAfterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(CaliburnParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.inline_statements(0)
		}
		{
			p.SetState(211)
			p.expression(0)
		}
		{
			p.SetState(212)
			p.Match(CaliburnParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.inline_statements(0)
		}
		{
			p.SetState(214)
			p.Block()
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

// ISwitch_statementContext is an interface to support dynamic dispatch.
type ISwitch_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_statementContext differentiates from other interfaces.
	IsSwitch_statementContext()
}

type Switch_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_statementContext() *Switch_statementContext {
	var p = new(Switch_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_switch_statement
	return p
}

func InitEmptySwitch_statementContext(p *Switch_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_switch_statement
}

func (*Switch_statementContext) IsSwitch_statementContext() {}

func NewSwitch_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_statementContext {
	var p = new(Switch_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_switch_statement

	return p
}

func (s *Switch_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_statementContext) CopyAll(ctx *Switch_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchStatementContext struct {
	Switch_statementContext
}

func NewSwitchStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	InitEmptySwitch_statementContext(&p.Switch_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_statementContext))

	return p
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(CaliburnParserSWITCH, 0)
}

func (s *SwitchStatementContext) Inline_statements() IInline_statementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *SwitchStatementContext) Case_blocks() ICase_blocksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICase_blocksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICase_blocksContext)
}

func (s *SwitchStatementContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CaliburnParserRULE_switch_statement)
	localctx = NewSwitchStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(CaliburnParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.inline_statements(0)
	}
	{
		p.SetState(220)
		p.expression(0)
	}
	{
		p.SetState(221)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.Case_blocks()
	}
	{
		p.SetState(223)
		p.Match(CaliburnParserR_C_BRACK)
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

// ICase_blocksContext is an interface to support dynamic dispatch.
type ICase_blocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCase_blocksContext differentiates from other interfaces.
	IsCase_blocksContext()
}

type Case_blocksContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_blocksContext() *Case_blocksContext {
	var p = new(Case_blocksContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_case_blocks
	return p
}

func InitEmptyCase_blocksContext(p *Case_blocksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_case_blocks
}

func (*Case_blocksContext) IsCase_blocksContext() {}

func NewCase_blocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_blocksContext {
	var p = new(Case_blocksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_case_blocks

	return p
}

func (s *Case_blocksContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_blocksContext) CopyAll(ctx *Case_blocksContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Case_blocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_blocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CaseBlocksContext struct {
	Case_blocksContext
}

func NewCaseBlocksContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseBlocksContext {
	var p = new(CaseBlocksContext)

	InitEmptyCase_blocksContext(&p.Case_blocksContext)
	p.parser = parser
	p.CopyAll(ctx.(*Case_blocksContext))

	return p
}

func (s *CaseBlocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlocksContext) AllOption_case_block() []IOption_case_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOption_case_blockContext); ok {
			len++
		}
	}

	tst := make([]IOption_case_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOption_case_blockContext); ok {
			tst[i] = t.(IOption_case_blockContext)
			i++
		}
	}

	return tst
}

func (s *CaseBlocksContext) Option_case_block(i int) IOption_case_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_case_blockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_case_blockContext)
}

func (s *CaseBlocksContext) Default_case_block() IDefault_case_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefault_case_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefault_case_blockContext)
}

func (s *CaseBlocksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCaseBlocks(s)
	}
}

func (s *CaseBlocksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCaseBlocks(s)
	}
}

func (s *CaseBlocksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCaseBlocks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Case_blocks() (localctx ICase_blocksContext) {
	localctx = NewCase_blocksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CaliburnParserRULE_case_blocks)
	var _la int

	localctx = NewCaseBlocksContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCASE {
		{
			p.SetState(225)
			p.Option_case_block()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserDEFAULT {
		{
			p.SetState(231)
			p.Default_case_block()
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

// IOption_case_blockContext is an interface to support dynamic dispatch.
type IOption_case_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOption_case_blockContext differentiates from other interfaces.
	IsOption_case_blockContext()
}

type Option_case_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_case_blockContext() *Option_case_blockContext {
	var p = new(Option_case_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_option_case_block
	return p
}

func InitEmptyOption_case_blockContext(p *Option_case_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_option_case_block
}

func (*Option_case_blockContext) IsOption_case_blockContext() {}

func NewOption_case_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_case_blockContext {
	var p = new(Option_case_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_option_case_block

	return p
}

func (s *Option_case_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_case_blockContext) CopyAll(ctx *Option_case_blockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Option_case_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_case_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OptionCaseBlockContext struct {
	Option_case_blockContext
}

func NewOptionCaseBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionCaseBlockContext {
	var p = new(OptionCaseBlockContext)

	InitEmptyOption_case_blockContext(&p.Option_case_blockContext)
	p.parser = parser
	p.CopyAll(ctx.(*Option_case_blockContext))

	return p
}

func (s *OptionCaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionCaseBlockContext) CASE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCASE, 0)
}

func (s *OptionCaseBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OptionCaseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *OptionCaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterOptionCaseBlock(s)
	}
}

func (s *OptionCaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitOptionCaseBlock(s)
	}
}

func (s *OptionCaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitOptionCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Option_case_block() (localctx IOption_case_blockContext) {
	localctx = NewOption_case_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CaliburnParserRULE_option_case_block)
	localctx = NewOptionCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(CaliburnParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.expression(0)
	}
	{
		p.SetState(236)
		p.Block()
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

// IDefault_case_blockContext is an interface to support dynamic dispatch.
type IDefault_case_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDefault_case_blockContext differentiates from other interfaces.
	IsDefault_case_blockContext()
}

type Default_case_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_case_blockContext() *Default_case_blockContext {
	var p = new(Default_case_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_default_case_block
	return p
}

func InitEmptyDefault_case_blockContext(p *Default_case_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_default_case_block
}

func (*Default_case_blockContext) IsDefault_case_blockContext() {}

func NewDefault_case_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_case_blockContext {
	var p = new(Default_case_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_default_case_block

	return p
}

func (s *Default_case_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_case_blockContext) CopyAll(ctx *Default_case_blockContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Default_case_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_case_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultCaseBlockContext struct {
	Default_case_blockContext
}

func NewDefaultCaseBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultCaseBlockContext {
	var p = new(DefaultCaseBlockContext)

	InitEmptyDefault_case_blockContext(&p.Default_case_blockContext)
	p.parser = parser
	p.CopyAll(ctx.(*Default_case_blockContext))

	return p
}

func (s *DefaultCaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultCaseBlockContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserDEFAULT, 0)
}

func (s *DefaultCaseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *DefaultCaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDefaultCaseBlock(s)
	}
}

func (s *DefaultCaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDefaultCaseBlock(s)
	}
}

func (s *DefaultCaseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDefaultCaseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Default_case_block() (localctx IDefault_case_blockContext) {
	localctx = NewDefault_case_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CaliburnParserRULE_default_case_block)
	localctx = NewDefaultCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(CaliburnParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Block()
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

// IInline_statementsContext is an interface to support dynamic dispatch.
type IInline_statementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInline_statementsContext differentiates from other interfaces.
	IsInline_statementsContext()
}

type Inline_statementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_statementsContext() *Inline_statementsContext {
	var p = new(Inline_statementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_inline_statements
	return p
}

func InitEmptyInline_statementsContext(p *Inline_statementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_inline_statements
}

func (*Inline_statementsContext) IsInline_statementsContext() {}

func NewInline_statementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_statementsContext {
	var p = new(Inline_statementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_inline_statements

	return p
}

func (s *Inline_statementsContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_statementsContext) CopyAll(ctx *Inline_statementsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Inline_statementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_statementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InlineStatementsInitialContext struct {
	Inline_statementsContext
}

func NewInlineStatementsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineStatementsInitialContext {
	var p = new(InlineStatementsInitialContext)

	InitEmptyInline_statementsContext(&p.Inline_statementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Inline_statementsContext))

	return p
}

func (s *InlineStatementsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineStatementsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterInlineStatementsInitial(s)
	}
}

func (s *InlineStatementsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitInlineStatementsInitial(s)
	}
}

func (s *InlineStatementsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitInlineStatementsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type InlineStatementsAppendContext struct {
	Inline_statementsContext
}

func NewInlineStatementsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InlineStatementsAppendContext {
	var p = new(InlineStatementsAppendContext)

	InitEmptyInline_statementsContext(&p.Inline_statementsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Inline_statementsContext))

	return p
}

func (s *InlineStatementsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineStatementsAppendContext) Inline_statements() IInline_statementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementsContext)
}

func (s *InlineStatementsAppendContext) Inline_statement() IInline_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInline_statementContext)
}

func (s *InlineStatementsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterInlineStatementsAppend(s)
	}
}

func (s *InlineStatementsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitInlineStatementsAppend(s)
	}
}

func (s *InlineStatementsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitInlineStatementsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Inline_statements() (localctx IInline_statementsContext) {
	return p.inline_statements(0)
}

func (p *CaliburnParser) inline_statements(_p int) (localctx IInline_statementsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewInline_statementsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInline_statementsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, CaliburnParserRULE_inline_statements, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewInlineStatementsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewInlineStatementsAppendContext(p, NewInline_statementsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_inline_statements)
			p.SetState(242)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(243)
				p.Inline_statement()
			}

		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInline_statementContext is an interface to support dynamic dispatch.
type IInline_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign_statement() IAssign_statementContext
	Expression_statement() IExpression_statementContext

	// IsInline_statementContext differentiates from other interfaces.
	IsInline_statementContext()
}

type Inline_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInline_statementContext() *Inline_statementContext {
	var p = new(Inline_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_inline_statement
	return p
}

func InitEmptyInline_statementContext(p *Inline_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_inline_statement
}

func (*Inline_statementContext) IsInline_statementContext() {}

func NewInline_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inline_statementContext {
	var p = new(Inline_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_inline_statement

	return p
}

func (s *Inline_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Inline_statementContext) Assign_statement() IAssign_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_statementContext)
}

func (s *Inline_statementContext) Expression_statement() IExpression_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_statementContext)
}

func (s *Inline_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inline_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inline_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterInline_statement(s)
	}
}

func (s *Inline_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitInline_statement(s)
	}
}

func (s *Inline_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitInline_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Inline_statement() (localctx IInline_statementContext) {
	localctx = NewInline_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CaliburnParserRULE_inline_statement)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Expression_statement()
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

// IAssign_statementContext is an interface to support dynamic dispatch.
type IAssign_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_statementContext differentiates from other interfaces.
	IsAssign_statementContext()
}

type Assign_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_statementContext() *Assign_statementContext {
	var p = new(Assign_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_statement
	return p
}

func InitEmptyAssign_statementContext(p *Assign_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_statement
}

func (*Assign_statementContext) IsAssign_statementContext() {}

func NewAssign_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_statementContext {
	var p = new(Assign_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_assign_statement

	return p
}

func (s *Assign_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_statementContext) CopyAll(ctx *Assign_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignOperationStatementContext struct {
	Assign_statementContext
	op antlr.Token
}

func NewAssignOperationStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignOperationStatementContext {
	var p = new(AssignOperationStatementContext)

	InitEmptyAssign_statementContext(&p.Assign_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_statementContext))

	return p
}

func (s *AssignOperationStatementContext) GetOp() antlr.Token { return s.op }

func (s *AssignOperationStatementContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignOperationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOperationStatementContext) Assign_expressions() IAssign_expressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionsContext)
}

func (s *AssignOperationStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserASSIGN, 0)
}

func (s *AssignOperationStatementContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *AssignOperationStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *AssignOperationStatementContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ADD, 0)
}

func (s *AssignOperationStatementContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_SUB, 0)
}

func (s *AssignOperationStatementContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MUL, 0)
}

func (s *AssignOperationStatementContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_DIV, 0)
}

func (s *AssignOperationStatementContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MOD, 0)
}

func (s *AssignOperationStatementContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_POW, 0)
}

func (s *AssignOperationStatementContext) OP_ROOT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ROOT, 0)
}

func (s *AssignOperationStatementContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_OR, 0)
}

func (s *AssignOperationStatementContext) OP_XOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_XOR, 0)
}

func (s *AssignOperationStatementContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_AND, 0)
}

func (s *AssignOperationStatementContext) OP_LSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LSHIFT, 0)
}

func (s *AssignOperationStatementContext) OP_RSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_RSHIFT, 0)
}

func (s *AssignOperationStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignOperationStatement(s)
	}
}

func (s *AssignOperationStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignOperationStatement(s)
	}
}

func (s *AssignOperationStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignOperationStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignStatementContext struct {
	Assign_statementContext
}

func NewAssignStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatementContext {
	var p = new(AssignStatementContext)

	InitEmptyAssign_statementContext(&p.Assign_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_statementContext))

	return p
}

func (s *AssignStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatementContext) Assign_declarations() IAssign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_declarationsContext)
}

func (s *AssignStatementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserASSIGN, 0)
}

func (s *AssignStatementContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *AssignStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *AssignStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignStatement(s)
	}
}

func (s *AssignStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignStatement(s)
	}
}

func (s *AssignStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_statement() (localctx IAssign_statementContext) {
	localctx = NewAssign_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CaliburnParserRULE_assign_statement)
	var _la int

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Assign_declarations()
		}
		{
			p.SetState(254)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(255)
			p.Expressions()
		}
		{
			p.SetState(256)
			p.Match(CaliburnParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewAssignOperationStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Assign_expressions()
		}
		{
			p.SetState(259)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*AssignOperationStatementContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&562606356037632) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*AssignOperationStatementContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(260)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Expressions()
		}
		{
			p.SetState(262)
			p.Match(CaliburnParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IAssign_expressionsContext is an interface to support dynamic dispatch.
type IAssign_expressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_expressionsContext differentiates from other interfaces.
	IsAssign_expressionsContext()
}

type Assign_expressionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_expressionsContext() *Assign_expressionsContext {
	var p = new(Assign_expressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_expressions
	return p
}

func InitEmptyAssign_expressionsContext(p *Assign_expressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_expressions
}

func (*Assign_expressionsContext) IsAssign_expressionsContext() {}

func NewAssign_expressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_expressionsContext {
	var p = new(Assign_expressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_assign_expressions

	return p
}

func (s *Assign_expressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_expressionsContext) CopyAll(ctx *Assign_expressionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_expressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_expressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignExpressionsContext struct {
	Assign_expressionsContext
}

func NewAssignExpressionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExpressionsContext {
	var p = new(AssignExpressionsContext)

	InitEmptyAssign_expressionsContext(&p.Assign_expressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionsContext))

	return p
}

func (s *AssignExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionsContext) AllAssign_expression() []IAssign_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			len++
		}
	}

	tst := make([]IAssign_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssign_expressionContext); ok {
			tst[i] = t.(IAssign_expressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignExpressionsContext) Assign_expression(i int) IAssign_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *AssignExpressionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *AssignExpressionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *AssignExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignExpressions(s)
	}
}

func (s *AssignExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignExpressions(s)
	}
}

func (s *AssignExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_expressions() (localctx IAssign_expressionsContext) {
	localctx = NewAssign_expressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CaliburnParserRULE_assign_expressions)
	var _la int

	localctx = NewAssignExpressionsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Assign_expression()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(267)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Assign_expression()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAssign_expressionContext is an interface to support dynamic dispatch.
type IAssign_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_expressionContext differentiates from other interfaces.
	IsAssign_expressionContext()
}

type Assign_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_expressionContext() *Assign_expressionContext {
	var p = new(Assign_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_expression
	return p
}

func InitEmptyAssign_expressionContext(p *Assign_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_expression
}

func (*Assign_expressionContext) IsAssign_expressionContext() {}

func NewAssign_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_expressionContext {
	var p = new(Assign_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_assign_expression

	return p
}

func (s *Assign_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_expressionContext) CopyAll(ctx *Assign_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionAssignExpressionContext struct {
	Assign_expressionContext
}

func NewExpressionAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAssignExpressionContext {
	var p = new(ExpressionAssignExpressionContext)

	InitEmptyAssign_expressionContext(&p.Assign_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionContext))

	return p
}

func (s *ExpressionAssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAssignExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionAssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionAssignExpression(s)
	}
}

func (s *ExpressionAssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionAssignExpression(s)
	}
}

func (s *ExpressionAssignExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionAssignExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleDestrutureAssignExpressionContext struct {
	Assign_expressionContext
}

func NewTupleDestrutureAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleDestrutureAssignExpressionContext {
	var p = new(TupleDestrutureAssignExpressionContext)

	InitEmptyAssign_expressionContext(&p.Assign_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionContext))

	return p
}

func (s *TupleDestrutureAssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleDestrutureAssignExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *TupleDestrutureAssignExpressionContext) Assign_expressions() IAssign_expressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionsContext)
}

func (s *TupleDestrutureAssignExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *TupleDestrutureAssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleDestrutureAssignExpression(s)
	}
}

func (s *TupleDestrutureAssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleDestrutureAssignExpression(s)
	}
}

func (s *TupleDestrutureAssignExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleDestrutureAssignExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDestrutureAssignExpressionContext struct {
	Assign_expressionContext
}

func NewStructDestrutureAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDestrutureAssignExpressionContext {
	var p = new(StructDestrutureAssignExpressionContext)

	InitEmptyAssign_expressionContext(&p.Assign_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionContext))

	return p
}

func (s *StructDestrutureAssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDestrutureAssignExpressionContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *StructDestrutureAssignExpressionContext) Assign_expressions() IAssign_expressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionsContext)
}

func (s *StructDestrutureAssignExpressionContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *StructDestrutureAssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructDestrutureAssignExpression(s)
	}
}

func (s *StructDestrutureAssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructDestrutureAssignExpression(s)
	}
}

func (s *StructDestrutureAssignExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructDestrutureAssignExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_expression() (localctx IAssign_expressionContext) {
	localctx = NewAssign_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CaliburnParserRULE_assign_expression)
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserL_PAREN, CaliburnParserOP_ADD, CaliburnParserOP_SUB, CaliburnParserOP_NOT, CaliburnParserIdentifier, CaliburnParserLiteral:
		localctx = NewExpressionAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.expression(0)
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewStructDestrutureAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.Assign_expressions()
		}
		{
			p.SetState(277)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserL_S_BRACK:
		localctx = NewTupleDestrutureAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(280)
			p.Assign_expressions()
		}
		{
			p.SetState(281)
			p.Match(CaliburnParserR_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IAssign_declarationsContext is an interface to support dynamic dispatch.
type IAssign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_declarationsContext differentiates from other interfaces.
	IsAssign_declarationsContext()
}

type Assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_declarationsContext() *Assign_declarationsContext {
	var p = new(Assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_declarations
	return p
}

func InitEmptyAssign_declarationsContext(p *Assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_declarations
}

func (*Assign_declarationsContext) IsAssign_declarationsContext() {}

func NewAssign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_declarationsContext {
	var p = new(Assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_assign_declarations

	return p
}

func (s *Assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_declarationsContext) CopyAll(ctx *Assign_declarationsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignDeclarationsContext struct {
	Assign_declarationsContext
}

func NewAssignDeclarationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDeclarationsContext {
	var p = new(AssignDeclarationsContext)

	InitEmptyAssign_declarationsContext(&p.Assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationsContext))

	return p
}

func (s *AssignDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDeclarationsContext) Assign_declaration() IAssign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_declarationContext)
}

func (s *AssignDeclarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *AssignDeclarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *AssignDeclarationsContext) AllAssign_statement() []IAssign_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAssign_statementContext); ok {
			len++
		}
	}

	tst := make([]IAssign_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAssign_statementContext); ok {
			tst[i] = t.(IAssign_statementContext)
			i++
		}
	}

	return tst
}

func (s *AssignDeclarationsContext) Assign_statement(i int) IAssign_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_statementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_statementContext)
}

func (s *AssignDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignDeclarations(s)
	}
}

func (s *AssignDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignDeclarations(s)
	}
}

func (s *AssignDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_declarations() (localctx IAssign_declarationsContext) {
	localctx = NewAssign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CaliburnParserRULE_assign_declarations)
	var _la int

	localctx = NewAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Assign_declaration()
	}
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(286)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(287)
			p.Assign_statement()
		}

		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IAssign_declarationContext is an interface to support dynamic dispatch.
type IAssign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssign_declarationContext differentiates from other interfaces.
	IsAssign_declarationContext()
}

type Assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_declarationContext() *Assign_declarationContext {
	var p = new(Assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_declaration
	return p
}

func InitEmptyAssign_declarationContext(p *Assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_assign_declaration
}

func (*Assign_declarationContext) IsAssign_declarationContext() {}

func NewAssign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_declarationContext {
	var p = new(Assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_assign_declaration

	return p
}

func (s *Assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_declarationContext) CopyAll(ctx *Assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionAssignDeclarationContext struct {
	Assign_declarationContext
}

func NewExpressionAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAssignDeclarationContext {
	var p = new(ExpressionAssignDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *ExpressionAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAssignDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionAssignDeclaration(s)
	}
}

func (s *ExpressionAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionAssignDeclaration(s)
	}
}

func (s *ExpressionAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedAssignDeclarationDeclarationContext struct {
	Assign_declarationContext
}

func NewUntypedAssignDeclarationDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedAssignDeclarationDeclarationContext {
	var p = new(UntypedAssignDeclarationDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *UntypedAssignDeclarationDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedAssignDeclarationDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserVAR, 0)
}

func (s *UntypedAssignDeclarationDeclarationContext) Untyped_assign_declaration() IUntyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationContext)
}

func (s *UntypedAssignDeclarationDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedAssignDeclarationDeclaration(s)
	}
}

func (s *UntypedAssignDeclarationDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedAssignDeclarationDeclaration(s)
	}
}

func (s *UntypedAssignDeclarationDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedAssignDeclarationDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedAssignDeclarationDeclarationContext struct {
	Assign_declarationContext
}

func NewTypedAssignDeclarationDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedAssignDeclarationDeclarationContext {
	var p = new(TypedAssignDeclarationDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *TypedAssignDeclarationDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedAssignDeclarationDeclarationContext) Typed_assign_declaration() ITyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITyped_assign_declarationContext)
}

func (s *TypedAssignDeclarationDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedAssignDeclarationDeclaration(s)
	}
}

func (s *TypedAssignDeclarationDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedAssignDeclarationDeclaration(s)
	}
}

func (s *TypedAssignDeclarationDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedAssignDeclarationDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructDestrutureAssignDeclarationContext struct {
	Assign_declarationContext
}

func NewStructDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructDestrutureAssignDeclarationContext {
	var p = new(StructDestrutureAssignDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *StructDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDestrutureAssignDeclarationContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *StructDestrutureAssignDeclarationContext) Assign_declarations() IAssign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_declarationsContext)
}

func (s *StructDestrutureAssignDeclarationContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *StructDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructDestrutureAssignDeclaration(s)
	}
}

func (s *StructDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructDestrutureAssignDeclaration(s)
	}
}

func (s *StructDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleDestrutureAssignDeclarationContext struct {
	Assign_declarationContext
}

func NewTupleDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleDestrutureAssignDeclarationContext {
	var p = new(TupleDestrutureAssignDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *TupleDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleDestrutureAssignDeclarationContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *TupleDestrutureAssignDeclarationContext) Assign_declarations() IAssign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_declarationsContext)
}

func (s *TupleDestrutureAssignDeclarationContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *TupleDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleDestrutureAssignDeclaration(s)
	}
}

func (s *TupleDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleDestrutureAssignDeclaration(s)
	}
}

func (s *TupleDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_declaration() (localctx IAssign_declarationContext) {
	localctx = NewAssign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CaliburnParserRULE_assign_declaration)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.expression(0)
		}

	case 2:
		localctx = NewTypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Typed_assign_declaration()
		}

	case 3:
		localctx = NewUntypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(295)
			p.Match(CaliburnParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(296)
			p.Untyped_assign_declaration()
		}

	case 4:
		localctx = NewStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Assign_declarations()
		}
		{
			p.SetState(299)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewTupleDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(301)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(302)
			p.Assign_declarations()
		}
		{
			p.SetState(303)
			p.Match(CaliburnParserR_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// ITyped_assign_declarationsContext is an interface to support dynamic dispatch.
type ITyped_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTyped_assign_declarationsContext differentiates from other interfaces.
	IsTyped_assign_declarationsContext()
}

type Typed_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyped_assign_declarationsContext() *Typed_assign_declarationsContext {
	var p = new(Typed_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_assign_declarations
	return p
}

func InitEmptyTyped_assign_declarationsContext(p *Typed_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_assign_declarations
}

func (*Typed_assign_declarationsContext) IsTyped_assign_declarationsContext() {}

func NewTyped_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typed_assign_declarationsContext {
	var p = new(Typed_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_typed_assign_declarations

	return p
}

func (s *Typed_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Typed_assign_declarationsContext) CopyAll(ctx *Typed_assign_declarationsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Typed_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typed_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedAssignDeclarationsContext struct {
	Typed_assign_declarationsContext
}

func NewTypedAssignDeclarationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedAssignDeclarationsContext {
	var p = new(TypedAssignDeclarationsContext)

	InitEmptyTyped_assign_declarationsContext(&p.Typed_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Typed_assign_declarationsContext))

	return p
}

func (s *TypedAssignDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedAssignDeclarationsContext) AllTyped_assign_declaration() []ITyped_assign_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITyped_assign_declarationContext); ok {
			len++
		}
	}

	tst := make([]ITyped_assign_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITyped_assign_declarationContext); ok {
			tst[i] = t.(ITyped_assign_declarationContext)
			i++
		}
	}

	return tst
}

func (s *TypedAssignDeclarationsContext) Typed_assign_declaration(i int) ITyped_assign_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITyped_assign_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITyped_assign_declarationContext)
}

func (s *TypedAssignDeclarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *TypedAssignDeclarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *TypedAssignDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedAssignDeclarations(s)
	}
}

func (s *TypedAssignDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedAssignDeclarations(s)
	}
}

func (s *TypedAssignDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedAssignDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Typed_assign_declarations() (localctx ITyped_assign_declarationsContext) {
	localctx = NewTyped_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CaliburnParserRULE_typed_assign_declarations)
	var _la int

	localctx = NewTypedAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Typed_assign_declaration()
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(308)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Typed_assign_declaration()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ITyped_assign_declarationContext is an interface to support dynamic dispatch.
type ITyped_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTyped_assign_declarationContext differentiates from other interfaces.
	IsTyped_assign_declarationContext()
}

type Typed_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyped_assign_declarationContext() *Typed_assign_declarationContext {
	var p = new(Typed_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_assign_declaration
	return p
}

func InitEmptyTyped_assign_declarationContext(p *Typed_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_typed_assign_declaration
}

func (*Typed_assign_declarationContext) IsTyped_assign_declarationContext() {}

func NewTyped_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typed_assign_declarationContext {
	var p = new(Typed_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_typed_assign_declaration

	return p
}

func (s *Typed_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Typed_assign_declarationContext) CopyAll(ctx *Typed_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Typed_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typed_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedAssignDeclarationContext struct {
	Typed_assign_declarationContext
	type_ IType_expressionContext
}

func NewTypedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedAssignDeclarationContext {
	var p = new(TypedAssignDeclarationContext)

	InitEmptyTyped_assign_declarationContext(&p.Typed_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Typed_assign_declarationContext))

	return p
}

func (s *TypedAssignDeclarationContext) GetType_() IType_expressionContext { return s.type_ }

func (s *TypedAssignDeclarationContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *TypedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedAssignDeclarationContext) Untyped_assign_declaration() IUntyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationContext)
}

func (s *TypedAssignDeclarationContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *TypedAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedAssignDeclaration(s)
	}
}

func (s *TypedAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedAssignDeclaration(s)
	}
}

func (s *TypedAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Typed_assign_declaration() (localctx ITyped_assign_declarationContext) {
	localctx = NewTyped_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CaliburnParserRULE_typed_assign_declaration)
	localctx = NewTypedAssignDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)

		var _x = p.type_expression(0)

		localctx.(*TypedAssignDeclarationContext).type_ = _x
	}
	{
		p.SetState(316)
		p.Untyped_assign_declaration()
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

// IUntyped_assign_declarationsContext is an interface to support dynamic dispatch.
type IUntyped_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUntyped_assign_declarationsContext differentiates from other interfaces.
	IsUntyped_assign_declarationsContext()
}

type Untyped_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntyped_assign_declarationsContext() *Untyped_assign_declarationsContext {
	var p = new(Untyped_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declarations
	return p
}

func InitEmptyUntyped_assign_declarationsContext(p *Untyped_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declarations
}

func (*Untyped_assign_declarationsContext) IsUntyped_assign_declarationsContext() {}

func NewUntyped_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Untyped_assign_declarationsContext {
	var p = new(Untyped_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declarations

	return p
}

func (s *Untyped_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Untyped_assign_declarationsContext) CopyAll(ctx *Untyped_assign_declarationsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Untyped_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Untyped_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UntypedAssignDeclarationsContext struct {
	Untyped_assign_declarationsContext
}

func NewUntypedAssignDeclarationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedAssignDeclarationsContext {
	var p = new(UntypedAssignDeclarationsContext)

	InitEmptyUntyped_assign_declarationsContext(&p.Untyped_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Untyped_assign_declarationsContext))

	return p
}

func (s *UntypedAssignDeclarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedAssignDeclarationsContext) AllUntyped_assign_declaration() []IUntyped_assign_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUntyped_assign_declarationContext); ok {
			len++
		}
	}

	tst := make([]IUntyped_assign_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUntyped_assign_declarationContext); ok {
			tst[i] = t.(IUntyped_assign_declarationContext)
			i++
		}
	}

	return tst
}

func (s *UntypedAssignDeclarationsContext) Untyped_assign_declaration(i int) IUntyped_assign_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationContext)
}

func (s *UntypedAssignDeclarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *UntypedAssignDeclarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *UntypedAssignDeclarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedAssignDeclarations(s)
	}
}

func (s *UntypedAssignDeclarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedAssignDeclarations(s)
	}
}

func (s *UntypedAssignDeclarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedAssignDeclarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Untyped_assign_declarations() (localctx IUntyped_assign_declarationsContext) {
	localctx = NewUntyped_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CaliburnParserRULE_untyped_assign_declarations)
	var _la int

	localctx = NewUntypedAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)
		p.Untyped_assign_declaration()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(319)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)
			p.Untyped_assign_declaration()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IUntyped_assign_declarationContext is an interface to support dynamic dispatch.
type IUntyped_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUntyped_assign_declarationContext differentiates from other interfaces.
	IsUntyped_assign_declarationContext()
}

type Untyped_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntyped_assign_declarationContext() *Untyped_assign_declarationContext {
	var p = new(Untyped_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declaration
	return p
}

func InitEmptyUntyped_assign_declarationContext(p *Untyped_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declaration
}

func (*Untyped_assign_declarationContext) IsUntyped_assign_declarationContext() {}

func NewUntyped_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Untyped_assign_declarationContext {
	var p = new(Untyped_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_untyped_assign_declaration

	return p
}

func (s *Untyped_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Untyped_assign_declarationContext) CopyAll(ctx *Untyped_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Untyped_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Untyped_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UntypedIdentifierAssignDeclarationContext struct {
	Untyped_assign_declarationContext
	var_ antlr.Token
}

func NewUntypedIdentifierAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedIdentifierAssignDeclarationContext {
	var p = new(UntypedIdentifierAssignDeclarationContext)

	InitEmptyUntyped_assign_declarationContext(&p.Untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Untyped_assign_declarationContext))

	return p
}

func (s *UntypedIdentifierAssignDeclarationContext) GetVar_() antlr.Token { return s.var_ }

func (s *UntypedIdentifierAssignDeclarationContext) SetVar_(v antlr.Token) { s.var_ = v }

func (s *UntypedIdentifierAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedIdentifierAssignDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *UntypedIdentifierAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedIdentifierAssignDeclaration(s)
	}
}

func (s *UntypedIdentifierAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedIdentifierAssignDeclaration(s)
	}
}

func (s *UntypedIdentifierAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedIdentifierAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedStructDestrutureAssignDeclarationContext struct {
	Untyped_assign_declarationContext
}

func NewUntypedStructDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedStructDestrutureAssignDeclarationContext {
	var p = new(UntypedStructDestrutureAssignDeclarationContext)

	InitEmptyUntyped_assign_declarationContext(&p.Untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Untyped_assign_declarationContext))

	return p
}

func (s *UntypedStructDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedStructDestrutureAssignDeclarationContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *UntypedStructDestrutureAssignDeclarationContext) Untyped_assign_declarations() IUntyped_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationsContext)
}

func (s *UntypedStructDestrutureAssignDeclarationContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *UntypedStructDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedStructDestrutureAssignDeclaration(s)
	}
}

func (s *UntypedStructDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedStructDestrutureAssignDeclaration(s)
	}
}

func (s *UntypedStructDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedStructDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedTupleDestrutureAssignDeclarationContext struct {
	Untyped_assign_declarationContext
}

func NewUntypedTupleDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedTupleDestrutureAssignDeclarationContext {
	var p = new(UntypedTupleDestrutureAssignDeclarationContext)

	InitEmptyUntyped_assign_declarationContext(&p.Untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Untyped_assign_declarationContext))

	return p
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) Untyped_assign_declarations() IUntyped_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUntyped_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUntyped_assign_declarationsContext)
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedTupleDestrutureAssignDeclaration(s)
	}
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedTupleDestrutureAssignDeclaration(s)
	}
}

func (s *UntypedTupleDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedTupleDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Untyped_assign_declaration() (localctx IUntyped_assign_declarationContext) {
	localctx = NewUntyped_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CaliburnParserRULE_untyped_assign_declaration)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifier:
		localctx = NewUntypedIdentifierAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)

			var _m = p.Match(CaliburnParserIdentifier)

			localctx.(*UntypedIdentifierAssignDeclarationContext).var_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewUntypedStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Untyped_assign_declarations()
		}
		{
			p.SetState(329)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserL_S_BRACK:
		localctx = NewUntypedTupleDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(331)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Untyped_assign_declarations()
		}
		{
			p.SetState(333)
			p.Match(CaliburnParserR_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IExpression_statementContext is an interface to support dynamic dispatch.
type IExpression_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpression_statementContext differentiates from other interfaces.
	IsExpression_statementContext()
}

type Expression_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_statementContext() *Expression_statementContext {
	var p = new(Expression_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_statement
	return p
}

func InitEmptyExpression_statementContext(p *Expression_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_statement
}

func (*Expression_statementContext) IsExpression_statementContext() {}

func NewExpression_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_statementContext {
	var p = new(Expression_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expression_statement

	return p
}

func (s *Expression_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_statementContext) CopyAll(ctx *Expression_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Expression_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionStatementContext struct {
	Expression_statementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	InitEmptyExpression_statementContext(&p.Expression_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Expression_statementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *ExpressionStatementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expression_statement() (localctx IExpression_statementContext) {
	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CaliburnParserRULE_expression_statement)
	localctx = NewExpressionStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Expressions()
	}
	{
		p.SetState(338)
		p.Match(CaliburnParserTerminator)
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

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expressions
	return p
}

func InitEmptyExpressionsContext(p *ExpressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expressions
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) CopyAll(ctx *ExpressionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionsRuleContext struct {
	ExpressionsContext
}

func NewExpressionsRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionsRuleContext {
	var p = new(ExpressionsRuleContext)

	InitEmptyExpressionsContext(&p.ExpressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionsContext))

	return p
}

func (s *ExpressionsRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsRuleContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionsRuleContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionsRuleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *ExpressionsRuleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *ExpressionsRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionsRule(s)
	}
}

func (s *ExpressionsRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionsRule(s)
	}
}

func (s *ExpressionsRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionsRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CaliburnParserRULE_expressions)
	var _la int

	localctx = NewExpressionsRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.expression(0)
	}
	p.SetState(345)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(341)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.expression(0)
		}

		p.SetState(347)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BracketedExpressionContext struct {
	ExpressionContext
	exp IExpressionContext
}

func NewBracketedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketedExpressionContext {
	var p = new(BracketedExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BracketedExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *BracketedExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *BracketedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketedExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *BracketedExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *BracketedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBracketedExpression(s)
	}
}

func (s *BracketedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBracketedExpression(s)
	}
}

func (s *BracketedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBracketedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	ExpressionContext
	literal ILiteral_atomContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetLiteral() ILiteral_atomContext { return s.literal }

func (s *LiteralExpressionContext) SetLiteral(v ILiteral_atomContext) { s.literal = v }

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal_atom() ILiteral_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_atomContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type SliceExpressionContext struct {
	ExpressionContext
	exp         IExpressionContext
	start_index IExpressionContext
	end_index   IExpressionContext
}

func NewSliceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceExpressionContext {
	var p = new(SliceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SliceExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *SliceExpressionContext) GetStart_index() IExpressionContext { return s.start_index }

func (s *SliceExpressionContext) GetEnd_index() IExpressionContext { return s.end_index }

func (s *SliceExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *SliceExpressionContext) SetStart_index(v IExpressionContext) { s.start_index = v }

func (s *SliceExpressionContext) SetEnd_index(v IExpressionContext) { s.end_index = v }

func (s *SliceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *SliceExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *SliceExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *SliceExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *SliceExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SliceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterSliceExpression(s)
	}
}

func (s *SliceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitSliceExpression(s)
	}
}

func (s *SliceExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitSliceExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleExpressionContext struct {
	ExpressionContext
}

func NewTupleExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleExpressionContext {
	var p = new(TupleExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TupleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *TupleExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *TupleExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TupleExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *TupleExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *TupleExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *TupleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleExpression(s)
	}
}

func (s *TupleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleExpression(s)
	}
}

func (s *TupleExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanBinaryExpressionContext struct {
	ExpressionContext
	lhs_exp IExpressionContext
	op      antlr.Token
	rhs_exp IExpressionContext
}

func NewBooleanBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanBinaryExpressionContext {
	var p = new(BooleanBinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BooleanBinaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BooleanBinaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BooleanBinaryExpressionContext) GetLhs_exp() IExpressionContext { return s.lhs_exp }

func (s *BooleanBinaryExpressionContext) GetRhs_exp() IExpressionContext { return s.rhs_exp }

func (s *BooleanBinaryExpressionContext) SetLhs_exp(v IExpressionContext) { s.lhs_exp = v }

func (s *BooleanBinaryExpressionContext) SetRhs_exp(v IExpressionContext) { s.rhs_exp = v }

func (s *BooleanBinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanBinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BooleanBinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BooleanBinaryExpressionContext) OP_EQU() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_EQU, 0)
}

func (s *BooleanBinaryExpressionContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_NEQ, 0)
}

func (s *BooleanBinaryExpressionContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LTE, 0)
}

func (s *BooleanBinaryExpressionContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GTE, 0)
}

func (s *BooleanBinaryExpressionContext) OP_LST() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LST, 0)
}

func (s *BooleanBinaryExpressionContext) OP_GRT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GRT, 0)
}

func (s *BooleanBinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBooleanBinaryExpression(s)
	}
}

func (s *BooleanBinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBooleanBinaryExpression(s)
	}
}

func (s *BooleanBinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBooleanBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExpressionContext struct {
	ExpressionContext
	exp   IExpressionContext
	index IExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *IndexExpressionContext) GetIndex() IExpressionContext { return s.index }

func (s *IndexExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *IndexExpressionContext) SetIndex(v IExpressionContext) { s.index = v }

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *IndexExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *IndexExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *IndexExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (s *IndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	ExpressionContext
	op  antlr.Token
	exp IExpressionContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *UnaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *UnaryExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) OP_NOT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_NOT, 0)
}

func (s *UnaryExpressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ADD, 0)
}

func (s *UnaryExpressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_SUB, 0)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccessExpressionContext struct {
	ExpressionContext
	exp        IExpressionContext
	identifier antlr.Token
}

func NewAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessExpressionContext {
	var p = new(AccessExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AccessExpressionContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *AccessExpressionContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *AccessExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *AccessExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *AccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessExpressionContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserPERIOD, 0)
}

func (s *AccessExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *AccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAccessExpression(s)
	}
}

func (s *AccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAccessExpression(s)
	}
}

func (s *AccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAccessExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
	identifier antlr.Token
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetIdentifier() antlr.Token { return s.identifier }

func (s *IdentifierExpressionContext) SetIdentifier(v antlr.Token) { s.identifier = v }

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionExpressionContext struct {
	ExpressionContext
	type_       IFunction_typeContext
	return_type IType_expressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetType_() IFunction_typeContext { return s.type_ }

func (s *FunctionExpressionContext) GetReturn_type() IType_expressionContext { return s.return_type }

func (s *FunctionExpressionContext) SetType_(v IFunction_typeContext) { s.type_ = v }

func (s *FunctionExpressionContext) SetReturn_type(v IType_expressionContext) { s.return_type = v }

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *FunctionExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *FunctionExpressionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionExpressionContext) Function_type() IFunction_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *FunctionExpressionContext) Assign_declarations() IAssign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_declarationsContext)
}

func (s *FunctionExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructExpressionContext struct {
	ExpressionContext
	type_ IStruct_typeContext
}

func NewStructExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExpressionContext {
	var p = new(StructExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StructExpressionContext) GetType_() IStruct_typeContext { return s.type_ }

func (s *StructExpressionContext) SetType_(v IStruct_typeContext) { s.type_ = v }

func (s *StructExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExpressionContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *StructExpressionContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOLON)
}

func (s *StructExpressionContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, i)
}

func (s *StructExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StructExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StructExpressionContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *StructExpressionContext) Struct_type() IStruct_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_typeContext)
}

func (s *StructExpressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserIdentifier)
}

func (s *StructExpressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, i)
}

func (s *StructExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *StructExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *StructExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructExpression(s)
	}
}

func (s *StructExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructExpression(s)
	}
}

func (s *StructExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpressionContext struct {
	ExpressionContext
	lhs_exp IExpressionContext
	op      antlr.Token
	rhs_exp IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinaryExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinaryExpressionContext) GetLhs_exp() IExpressionContext { return s.lhs_exp }

func (s *BinaryExpressionContext) GetRhs_exp() IExpressionContext { return s.rhs_exp }

func (s *BinaryExpressionContext) SetLhs_exp(v IExpressionContext) { s.lhs_exp = v }

func (s *BinaryExpressionContext) SetRhs_exp(v IExpressionContext) { s.rhs_exp = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_POW, 0)
}

func (s *BinaryExpressionContext) OP_ROOT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ROOT, 0)
}

func (s *BinaryExpressionContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MUL, 0)
}

func (s *BinaryExpressionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_DIV, 0)
}

func (s *BinaryExpressionContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MOD, 0)
}

func (s *BinaryExpressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ADD, 0)
}

func (s *BinaryExpressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_SUB, 0)
}

func (s *BinaryExpressionContext) OP_LSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LSHIFT, 0)
}

func (s *BinaryExpressionContext) OP_RSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_RSHIFT, 0)
}

func (s *BinaryExpressionContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_AND, 0)
}

func (s *BinaryExpressionContext) OP_XOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_XOR, 0)
}

func (s *BinaryExpressionContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_OR, 0)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastExpressionContext struct {
	ExpressionContext
	type_ IType_expressionContext
	exp   IExpressionContext
}

func NewCastExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExpressionContext {
	var p = new(CastExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CastExpressionContext) GetType_() IType_expressionContext { return s.type_ }

func (s *CastExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *CastExpressionContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *CastExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *CastExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *CastExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *CastExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *CastExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CastExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCastExpression(s)
	}
}

func (s *CastExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCastExpression(s)
	}
}

func (s *CastExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCastExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpressionContext struct {
	ExpressionContext
	exp  IExpressionContext
	args IExpressionsContext
}

func NewCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpressionContext {
	var p = new(CallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *CallExpressionContext) GetArgs() IExpressionsContext { return s.args }

func (s *CallExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *CallExpressionContext) SetArgs(v IExpressionsContext) { s.args = v }

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *CallExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *CallExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CallExpressionContext) Expressions() IExpressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *CallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCallExpression(s)
	}
}

func (s *CallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCallExpression(s)
	}
}

func (s *CallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CaliburnParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, CaliburnParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(349)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(350)

			var _x = p.type_expression(0)

			localctx.(*CastExpressionContext).type_ = _x
		}
		{
			p.SetState(351)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(352)

			var _x = p.expression(20)

			localctx.(*CastExpressionContext).exp = _x
		}

	case 2:
		localctx = NewBracketedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(354)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _x = p.expression(0)

			localctx.(*BracketedExpressionContext).exp = _x
		}
		{
			p.SetState(356)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(358)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&481036337152) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(359)

			var _x = p.expression(17)

			localctx.(*UnaryExpressionContext).exp = _x
		}

	case 4:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(360)

			var _m = p.Match(CaliburnParserIdentifier)

			localctx.(*IdentifierExpressionContext).identifier = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(361)

			var _x = p.Literal_atom()

			localctx.(*LiteralExpressionContext).literal = _x
		}

	case 6:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(362)

			var _x = p.Function_type()

			localctx.(*FunctionExpressionContext).type_ = _x
		}
		{
			p.SetState(363)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864691632144908288) != 0 {
			{
				p.SetState(364)
				p.Assign_declarations()
			}

		}
		{
			p.SetState(367)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CaliburnParserIdentifier {
			{
				p.SetState(368)

				var _x = p.type_expression(0)

				localctx.(*FunctionExpressionContext).return_type = _x
			}

		}
		{
			p.SetState(371)
			p.Block()
		}

	case 7:
		localctx = NewStructExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(373)

			var _x = p.Struct_type()

			localctx.(*StructExpressionContext).type_ = _x
		}
		{
			p.SetState(374)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(375)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(376)
			p.Match(CaliburnParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(377)
			p.expression(0)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(378)
					p.Match(CaliburnParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(379)
					p.Match(CaliburnParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(380)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(381)
					p.expression(0)
				}

			}
			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CaliburnParserCOMMA {
			{
				p.SetState(387)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(390)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewTupleExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(392)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.expression(0)
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(394)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == CaliburnParserCOMMA {
				{
					p.SetState(395)
					p.Match(CaliburnParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(396)
					p.expression(0)
				}

				p.SetState(399)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(403)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(407)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(408)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_POW || _la == CaliburnParserOP_ROOT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(409)

					var _x = p.expression(19)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(410)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(411)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3848290697216) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(412)

					var _x = p.expression(17)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(414)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_ADD || _la == CaliburnParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(415)

					var _x = p.expression(16)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(417)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_LSHIFT || _la == CaliburnParserOP_RSHIFT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(418)

					var _x = p.expression(15)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 5:
				localctx = NewBooleanBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BooleanBinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(420)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BooleanBinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35465847065542656) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BooleanBinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(421)

					var _x = p.expression(14)

					localctx.(*BooleanBinaryExpressionContext).rhs_exp = _x
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(422)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(423)
					p.Match(CaliburnParserOP_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(424)

					var _x = p.expression(13)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(425)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(426)
					p.Match(CaliburnParserOP_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(427)

					var _x = p.expression(12)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(428)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(429)
					p.Match(CaliburnParserOP_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(430)

					var _x = p.expression(11)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 9:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CallExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(431)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(432)
					p.Match(CaliburnParserL_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(434)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864691610665877504) != 0 {
					{
						p.SetState(433)

						var _x = p.Expressions()

						localctx.(*CallExpressionContext).args = _x
					}

				}
				{
					p.SetState(436)
					p.Match(CaliburnParserR_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*AccessExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(438)
					p.Match(CaliburnParserPERIOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(439)

					var _m = p.Match(CaliburnParserIdentifier)

					localctx.(*AccessExpressionContext).identifier = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 11:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*IndexExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(440)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(441)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(442)

					var _x = p.expression(0)

					localctx.(*IndexExpressionContext).index = _x
				}
				{
					p.SetState(443)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewSliceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*SliceExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(446)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(448)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864691610665877504) != 0 {
					{
						p.SetState(447)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).start_index = _x
					}

				}
				{
					p.SetState(450)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(452)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&864691610665877504) != 0 {
					{
						p.SetState(451)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).end_index = _x
					}

				}
				{
					p.SetState(454)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 37, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_typeContext is an interface to support dynamic dispatch.
type IFunction_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunction_typeContext differentiates from other interfaces.
	IsFunction_typeContext()
}

type Function_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_typeContext() *Function_typeContext {
	var p = new(Function_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_type
	return p
}

func InitEmptyFunction_typeContext(p *Function_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_type
}

func (*Function_typeContext) IsFunction_typeContext() {}

func NewFunction_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_typeContext {
	var p = new(Function_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_function_type

	return p
}

func (s *Function_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_typeContext) CopyAll(ctx *Function_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Function_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionTypeExpressionContext struct {
	Function_typeContext
}

func NewFunctionTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionTypeExpressionContext {
	var p = new(FunctionTypeExpressionContext)

	InitEmptyFunction_typeContext(&p.Function_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_typeContext))

	return p
}

func (s *FunctionTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *FunctionTypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionTypeExpression(s)
	}
}

func (s *FunctionTypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionTypeExpression(s)
	}
}

func (s *FunctionTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionTypeFuncContext struct {
	Function_typeContext
}

func NewFunctionTypeFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionTypeFuncContext {
	var p = new(FunctionTypeFuncContext)

	InitEmptyFunction_typeContext(&p.Function_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_typeContext))

	return p
}

func (s *FunctionTypeFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeFuncContext) FUNC() antlr.TerminalNode {
	return s.GetToken(CaliburnParserFUNC, 0)
}

func (s *FunctionTypeFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionTypeFunc(s)
	}
}

func (s *FunctionTypeFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionTypeFunc(s)
	}
}

func (s *FunctionTypeFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionTypeFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Function_type() (localctx IFunction_typeContext) {
	localctx = NewFunction_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CaliburnParserRULE_function_type)
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC:
		localctx = NewFunctionTypeFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(460)
			p.Match(CaliburnParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifier:
		localctx = NewFunctionTypeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.type_expression(0)
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

// IStruct_typeContext is an interface to support dynamic dispatch.
type IStruct_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_typeContext differentiates from other interfaces.
	IsStruct_typeContext()
}

type Struct_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_typeContext() *Struct_typeContext {
	var p = new(Struct_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_type
	return p
}

func InitEmptyStruct_typeContext(p *Struct_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_type
}

func (*Struct_typeContext) IsStruct_typeContext() {}

func NewStruct_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_typeContext {
	var p = new(Struct_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_struct_type

	return p
}

func (s *Struct_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_typeContext) CopyAll(ctx *Struct_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructTypeStructContext struct {
	Struct_typeContext
}

func NewStructTypeStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeStructContext {
	var p = new(StructTypeStructContext)

	InitEmptyStruct_typeContext(&p.Struct_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_typeContext))

	return p
}

func (s *StructTypeStructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeStructContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserSTRUCT, 0)
}

func (s *StructTypeStructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructTypeStruct(s)
	}
}

func (s *StructTypeStructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructTypeStruct(s)
	}
}

func (s *StructTypeStructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructTypeStruct(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructTypeExpressionContext struct {
	Struct_typeContext
}

func NewStructTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeExpressionContext {
	var p = new(StructTypeExpressionContext)

	InitEmptyStruct_typeContext(&p.Struct_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_typeContext))

	return p
}

func (s *StructTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *StructTypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructTypeExpression(s)
	}
}

func (s *StructTypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructTypeExpression(s)
	}
}

func (s *StructTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Struct_type() (localctx IStruct_typeContext) {
	localctx = NewStruct_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CaliburnParserRULE_struct_type)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserSTRUCT:
		localctx = NewStructTypeStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.Match(CaliburnParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifier:
		localctx = NewStructTypeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.type_expression(0)
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

// ITuple_typeContext is an interface to support dynamic dispatch.
type ITuple_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTuple_typeContext differentiates from other interfaces.
	IsTuple_typeContext()
}

type Tuple_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuple_typeContext() *Tuple_typeContext {
	var p = new(Tuple_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_tuple_type
	return p
}

func InitEmptyTuple_typeContext(p *Tuple_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_tuple_type
}

func (*Tuple_typeContext) IsTuple_typeContext() {}

func NewTuple_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tuple_typeContext {
	var p = new(Tuple_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_tuple_type

	return p
}

func (s *Tuple_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Tuple_typeContext) CopyAll(ctx *Tuple_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Tuple_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tuple_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TupleTypeTupleContext struct {
	Tuple_typeContext
}

func NewTupleTypeTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeTupleContext {
	var p = new(TupleTypeTupleContext)

	InitEmptyTuple_typeContext(&p.Tuple_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tuple_typeContext))

	return p
}

func (s *TupleTypeTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeTupleContext) TUPLE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTUPLE, 0)
}

func (s *TupleTypeTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleTypeTuple(s)
	}
}

func (s *TupleTypeTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleTypeTuple(s)
	}
}

func (s *TupleTypeTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleTypeTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

type TupleTypeExpressionContext struct {
	Tuple_typeContext
}

func NewTupleTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleTypeExpressionContext {
	var p = new(TupleTypeExpressionContext)

	InitEmptyTuple_typeContext(&p.Tuple_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Tuple_typeContext))

	return p
}

func (s *TupleTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleTypeExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *TupleTypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTupleTypeExpression(s)
	}
}

func (s *TupleTypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTupleTypeExpression(s)
	}
}

func (s *TupleTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTupleTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Tuple_type() (localctx ITuple_typeContext) {
	localctx = NewTuple_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CaliburnParserRULE_tuple_type)
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserTUPLE:
		localctx = NewTupleTypeTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Match(CaliburnParserTUPLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifier:
		localctx = NewTupleTypeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
			p.type_expression(0)
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

// IType_expressionContext is an interface to support dynamic dispatch.
type IType_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsType_expressionContext differentiates from other interfaces.
	IsType_expressionContext()
}

type Type_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_expressionContext() *Type_expressionContext {
	var p = new(Type_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_expression
	return p
}

func InitEmptyType_expressionContext(p *Type_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_expression
}

func (*Type_expressionContext) IsType_expressionContext() {}

func NewType_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_expressionContext {
	var p = new(Type_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_type_expression

	return p
}

func (s *Type_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_expressionContext) CopyAll(ctx *Type_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Type_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccessTypeExpressionContext struct {
	Type_expressionContext
}

func NewAccessTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessTypeExpressionContext {
	var p = new(AccessTypeExpressionContext)

	InitEmptyType_expressionContext(&p.Type_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Type_expressionContext))

	return p
}

func (s *AccessTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessTypeExpressionContext) Type_expression() IType_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *AccessTypeExpressionContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserPERIOD, 0)
}

func (s *AccessTypeExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *AccessTypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAccessTypeExpression(s)
	}
}

func (s *AccessTypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAccessTypeExpression(s)
	}
}

func (s *AccessTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAccessTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierTypeExpressionContext struct {
	Type_expressionContext
}

func NewIdentifierTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierTypeExpressionContext {
	var p = new(IdentifierTypeExpressionContext)

	InitEmptyType_expressionContext(&p.Type_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Type_expressionContext))

	return p
}

func (s *IdentifierTypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierTypeExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *IdentifierTypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIdentifierTypeExpression(s)
	}
}

func (s *IdentifierTypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIdentifierTypeExpression(s)
	}
}

func (s *IdentifierTypeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIdentifierTypeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Type_expression() (localctx IType_expressionContext) {
	return p.type_expression(0)
}

func (p *CaliburnParser) type_expression(_p int) (localctx IType_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewType_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IType_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, CaliburnParserRULE_type_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewIdentifierTypeExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(473)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(480)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAccessTypeExpressionContext(p, NewType_expressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_type_expression)
			p.SetState(475)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(476)
				p.Match(CaliburnParserPERIOD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(477)
				p.Match(CaliburnParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(482)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIdentifiersContext differentiates from other interfaces.
	IsIdentifiersContext()
}

type IdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiersContext() *IdentifiersContext {
	var p = new(IdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_identifiers
	return p
}

func InitEmptyIdentifiersContext(p *IdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_identifiers
}

func (*IdentifiersContext) IsIdentifiersContext() {}

func NewIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiersContext {
	var p = new(IdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_identifiers

	return p
}

func (s *IdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiersContext) CopyAll(ctx *IdentifiersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IdentifiersRuleContext struct {
	IdentifiersContext
}

func NewIdentifiersRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifiersRuleContext {
	var p = new(IdentifiersRuleContext)

	InitEmptyIdentifiersContext(&p.IdentifiersContext)
	p.parser = parser
	p.CopyAll(ctx.(*IdentifiersContext))

	return p
}

func (s *IdentifiersRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersRuleContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserIdentifier)
}

func (s *IdentifiersRuleContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, i)
}

func (s *IdentifiersRuleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *IdentifiersRuleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *IdentifiersRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIdentifiersRule(s)
	}
}

func (s *IdentifiersRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIdentifiersRule(s)
	}
}

func (s *IdentifiersRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIdentifiersRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CaliburnParserRULE_identifiers)
	var _la int

	localctx = NewIdentifiersRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(483)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(484)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ILiteral_atomContext is an interface to support dynamic dispatch.
type ILiteral_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteral_atomContext differentiates from other interfaces.
	IsLiteral_atomContext()
}

type Literal_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_atomContext() *Literal_atomContext {
	var p = new(Literal_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_literal_atom
	return p
}

func InitEmptyLiteral_atomContext(p *Literal_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_literal_atom
}

func (*Literal_atomContext) IsLiteral_atomContext() {}

func NewLiteral_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_atomContext {
	var p = new(Literal_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_literal_atom

	return p
}

func (s *Literal_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_atomContext) CopyAll(ctx *Literal_atomContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Literal_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedLiteralContext struct {
	Literal_atomContext
}

func NewTypedLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedLiteralContext {
	var p = new(TypedLiteralContext)

	InitEmptyLiteral_atomContext(&p.Literal_atomContext)
	p.parser = parser
	p.CopyAll(ctx.(*Literal_atomContext))

	return p
}

func (s *TypedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedLiteralContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *TypedLiteralContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
}

func (s *TypedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedLiteral(s)
	}
}

func (s *TypedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedLiteral(s)
	}
}

func (s *TypedLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type UntypedLiteralContext struct {
	Literal_atomContext
}

func NewUntypedLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedLiteralContext {
	var p = new(UntypedLiteralContext)

	InitEmptyLiteral_atomContext(&p.Literal_atomContext)
	p.parser = parser
	p.CopyAll(ctx.(*Literal_atomContext))

	return p
}

func (s *UntypedLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedLiteralContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
}

func (s *UntypedLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedLiteral(s)
	}
}

func (s *UntypedLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedLiteral(s)
	}
}

func (s *UntypedLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Literal_atom() (localctx ILiteral_atomContext) {
	localctx = NewLiteral_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, CaliburnParserRULE_literal_atom)
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserLiteral:
		localctx = NewUntypedLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(491)
			p.Match(CaliburnParserLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifier:
		localctx = NewTypedLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(492)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Match(CaliburnParserLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

func (p *CaliburnParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *DefinitionsContext = nil
		if localctx != nil {
			t = localctx.(*DefinitionsContext)
		}
		return p.Definitions_Sempred(t, predIndex)

	case 5:
		var t *Parameters_listContext = nil
		if localctx != nil {
			t = localctx.(*Parameters_listContext)
		}
		return p.Parameters_list_Sempred(t, predIndex)

	case 8:
		var t *StatementsContext = nil
		if localctx != nil {
			t = localctx.(*StatementsContext)
		}
		return p.Statements_Sempred(t, predIndex)

	case 22:
		var t *Inline_statementsContext = nil
		if localctx != nil {
			t = localctx.(*Inline_statementsContext)
		}
		return p.Inline_statements_Sempred(t, predIndex)

	case 35:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 39:
		var t *Type_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Type_expressionContext)
		}
		return p.Type_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CaliburnParser) Definitions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Parameters_list_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Statements_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Inline_statements_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Type_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 16:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

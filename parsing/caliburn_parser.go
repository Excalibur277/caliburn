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
		"CommentSingleLine", "IdentifierToken", "LiteralToken",
	}
	staticData.RuleNames = []string{
		"module", "definitions", "definition", "function_definition", "parameters",
		"parameters_list", "parameter", "block", "statements", "statement",
		"inline_statements", "inline_statement", "jump_statement", "return_statement",
		"break_statement", "continue_statement", "control_statement", "if_statement",
		"else_statement", "for_statement", "switch_statement", "case_blocks",
		"option_case_blocks", "option_case_block", "default_case_block", "assign_statement",
		"assign_expressions", "aliasable_assign_expressions", "aliasable_assign_expression",
		"assign_expression", "assign_declarations", "aliasable_assign_declarations",
		"aliasable_assign_declaration", "assign_declaration", "typed_assign_declaration",
		"aliasable_untyped_assign_declarations", "aliasable_untyped_assign_declaration",
		"untyped_assign_declaration", "expression_statement", "expressions",
		"expressions_optional", "expression", "struct_values", "named_struct_values",
		"named_struct_value", "collection_values", "collection_value", "function_type",
		"struct_type", "type_expression", "literal_atom", "literal", "identifier",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 565, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 113, 8, 1, 10, 1, 12, 1,
		116, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 135, 8, 3, 1, 4, 1, 4,
		3, 4, 139, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 5, 5, 147, 8, 5, 10,
		5, 12, 5, 150, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 158, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 167, 8, 8, 10, 8, 12,
		8, 170, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 176, 8, 9, 1, 10, 1, 10, 1,
		10, 5, 10, 181, 8, 10, 10, 10, 12, 10, 184, 9, 10, 1, 11, 1, 11, 3, 11,
		188, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 193, 8, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3,
		16, 208, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 221, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 227, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 241, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 254, 8, 21,
		1, 22, 1, 22, 1, 22, 5, 22, 259, 8, 22, 10, 22, 12, 22, 262, 9, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 282, 8, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 290, 8, 26, 10, 26, 12, 26,
		293, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 301, 8, 27,
		10, 27, 12, 27, 304, 9, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 311,
		8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 3, 29, 318, 8, 29, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 326, 8, 30, 10, 30, 12, 30, 329,
		9, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 337, 8, 31, 10,
		31, 12, 31, 340, 9, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 347,
		8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 3, 33, 357,
		8, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5,
		35, 368, 8, 35, 10, 35, 12, 35, 371, 9, 35, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 3, 36, 378, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 3, 37, 385,
		8, 37, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 5,
		39, 396, 8, 39, 10, 39, 12, 39, 399, 9, 39, 1, 40, 1, 40, 3, 40, 403, 8,
		40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 3, 41, 437, 8, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41,
		1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 5, 41, 495, 8, 41,
		10, 41, 12, 41, 498, 9, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 505,
		8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 513, 8, 43, 10,
		43, 12, 43, 516, 9, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45,
		1, 45, 5, 45, 526, 8, 45, 10, 45, 12, 45, 529, 9, 45, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 47, 1, 47, 3, 47, 537, 8, 47, 1, 48, 1, 48, 1, 48, 3, 48,
		542, 8, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 5, 49, 550, 8, 49,
		10, 49, 12, 49, 553, 9, 49, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 559, 8,
		50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 0, 15, 2, 10, 16, 20, 44, 52, 54,
		60, 62, 70, 78, 82, 86, 90, 98, 53, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90,
		92, 94, 96, 98, 100, 102, 104, 0, 7, 2, 0, 36, 37, 39, 48, 1, 0, 36, 38,
		1, 0, 42, 43, 1, 0, 39, 41, 1, 0, 36, 37, 1, 0, 47, 48, 1, 0, 49, 54, 578,
		0, 106, 1, 0, 0, 0, 2, 109, 1, 0, 0, 0, 4, 117, 1, 0, 0, 0, 6, 134, 1,
		0, 0, 0, 8, 138, 1, 0, 0, 0, 10, 140, 1, 0, 0, 0, 12, 157, 1, 0, 0, 0,
		14, 159, 1, 0, 0, 0, 16, 163, 1, 0, 0, 0, 18, 175, 1, 0, 0, 0, 20, 177,
		1, 0, 0, 0, 22, 187, 1, 0, 0, 0, 24, 192, 1, 0, 0, 0, 26, 194, 1, 0, 0,
		0, 28, 198, 1, 0, 0, 0, 30, 201, 1, 0, 0, 0, 32, 207, 1, 0, 0, 0, 34, 220,
		1, 0, 0, 0, 36, 226, 1, 0, 0, 0, 38, 240, 1, 0, 0, 0, 40, 242, 1, 0, 0,
		0, 42, 253, 1, 0, 0, 0, 44, 255, 1, 0, 0, 0, 46, 263, 1, 0, 0, 0, 48, 267,
		1, 0, 0, 0, 50, 281, 1, 0, 0, 0, 52, 283, 1, 0, 0, 0, 54, 294, 1, 0, 0,
		0, 56, 310, 1, 0, 0, 0, 58, 317, 1, 0, 0, 0, 60, 319, 1, 0, 0, 0, 62, 330,
		1, 0, 0, 0, 64, 346, 1, 0, 0, 0, 66, 356, 1, 0, 0, 0, 68, 358, 1, 0, 0,
		0, 70, 361, 1, 0, 0, 0, 72, 377, 1, 0, 0, 0, 74, 384, 1, 0, 0, 0, 76, 386,
		1, 0, 0, 0, 78, 389, 1, 0, 0, 0, 80, 402, 1, 0, 0, 0, 82, 436, 1, 0, 0,
		0, 84, 504, 1, 0, 0, 0, 86, 506, 1, 0, 0, 0, 88, 517, 1, 0, 0, 0, 90, 521,
		1, 0, 0, 0, 92, 530, 1, 0, 0, 0, 94, 536, 1, 0, 0, 0, 96, 541, 1, 0, 0,
		0, 98, 543, 1, 0, 0, 0, 100, 558, 1, 0, 0, 0, 102, 560, 1, 0, 0, 0, 104,
		562, 1, 0, 0, 0, 106, 107, 3, 2, 1, 0, 107, 108, 5, 0, 0, 1, 108, 1, 1,
		0, 0, 0, 109, 114, 6, 1, -1, 0, 110, 111, 10, 1, 0, 0, 111, 113, 3, 4,
		2, 0, 112, 110, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0,
		114, 115, 1, 0, 0, 0, 115, 3, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118,
		3, 6, 3, 0, 118, 5, 1, 0, 0, 0, 119, 120, 3, 94, 47, 0, 120, 121, 3, 104,
		52, 0, 121, 122, 5, 30, 0, 0, 122, 123, 3, 8, 4, 0, 123, 124, 5, 31, 0,
		0, 124, 125, 3, 98, 49, 0, 125, 126, 3, 14, 7, 0, 126, 135, 1, 0, 0, 0,
		127, 128, 3, 94, 47, 0, 128, 129, 3, 104, 52, 0, 129, 130, 5, 30, 0, 0,
		130, 131, 3, 8, 4, 0, 131, 132, 5, 31, 0, 0, 132, 133, 3, 14, 7, 0, 133,
		135, 1, 0, 0, 0, 134, 119, 1, 0, 0, 0, 134, 127, 1, 0, 0, 0, 135, 7, 1,
		0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 139, 3, 10, 5, 0, 138, 136, 1, 0, 0,
		0, 138, 137, 1, 0, 0, 0, 139, 9, 1, 0, 0, 0, 140, 141, 6, 5, -1, 0, 141,
		142, 3, 12, 6, 0, 142, 148, 1, 0, 0, 0, 143, 144, 10, 1, 0, 0, 144, 145,
		5, 2, 0, 0, 145, 147, 3, 12, 6, 0, 146, 143, 1, 0, 0, 0, 147, 150, 1, 0,
		0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 11, 1, 0, 0, 0,
		150, 148, 1, 0, 0, 0, 151, 158, 3, 68, 34, 0, 152, 158, 3, 74, 37, 0, 153,
		154, 5, 34, 0, 0, 154, 155, 3, 8, 4, 0, 155, 156, 5, 35, 0, 0, 156, 158,
		1, 0, 0, 0, 157, 151, 1, 0, 0, 0, 157, 152, 1, 0, 0, 0, 157, 153, 1, 0,
		0, 0, 158, 13, 1, 0, 0, 0, 159, 160, 5, 34, 0, 0, 160, 161, 3, 16, 8, 0,
		161, 162, 5, 35, 0, 0, 162, 15, 1, 0, 0, 0, 163, 168, 6, 8, -1, 0, 164,
		165, 10, 1, 0, 0, 165, 167, 3, 18, 9, 0, 166, 164, 1, 0, 0, 0, 167, 170,
		1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 17, 1, 0,
		0, 0, 170, 168, 1, 0, 0, 0, 171, 176, 3, 50, 25, 0, 172, 176, 3, 76, 38,
		0, 173, 176, 3, 32, 16, 0, 174, 176, 3, 24, 12, 0, 175, 171, 1, 0, 0, 0,
		175, 172, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176,
		19, 1, 0, 0, 0, 177, 182, 6, 10, -1, 0, 178, 179, 10, 1, 0, 0, 179, 181,
		3, 22, 11, 0, 180, 178, 1, 0, 0, 0, 181, 184, 1, 0, 0, 0, 182, 180, 1,
		0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 21, 1, 0, 0, 0, 184, 182, 1, 0, 0,
		0, 185, 188, 3, 50, 25, 0, 186, 188, 3, 76, 38, 0, 187, 185, 1, 0, 0, 0,
		187, 186, 1, 0, 0, 0, 188, 23, 1, 0, 0, 0, 189, 193, 3, 26, 13, 0, 190,
		193, 3, 28, 14, 0, 191, 193, 3, 30, 15, 0, 192, 189, 1, 0, 0, 0, 192, 190,
		1, 0, 0, 0, 192, 191, 1, 0, 0, 0, 193, 25, 1, 0, 0, 0, 194, 195, 5, 17,
		0, 0, 195, 196, 3, 82, 41, 0, 196, 197, 5, 5, 0, 0, 197, 27, 1, 0, 0, 0,
		198, 199, 5, 15, 0, 0, 199, 200, 5, 5, 0, 0, 200, 29, 1, 0, 0, 0, 201,
		202, 5, 16, 0, 0, 202, 203, 5, 5, 0, 0, 203, 31, 1, 0, 0, 0, 204, 208,
		3, 34, 17, 0, 205, 208, 3, 38, 19, 0, 206, 208, 3, 40, 20, 0, 207, 204,
		1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0, 208, 33, 1, 0,
		0, 0, 209, 210, 5, 8, 0, 0, 210, 211, 3, 20, 10, 0, 211, 212, 3, 82, 41,
		0, 212, 213, 3, 14, 7, 0, 213, 221, 1, 0, 0, 0, 214, 215, 5, 8, 0, 0, 215,
		216, 3, 20, 10, 0, 216, 217, 3, 82, 41, 0, 217, 218, 3, 14, 7, 0, 218,
		219, 3, 36, 18, 0, 219, 221, 1, 0, 0, 0, 220, 209, 1, 0, 0, 0, 220, 214,
		1, 0, 0, 0, 221, 35, 1, 0, 0, 0, 222, 223, 5, 9, 0, 0, 223, 227, 3, 14,
		7, 0, 224, 225, 5, 9, 0, 0, 225, 227, 3, 34, 17, 0, 226, 222, 1, 0, 0,
		0, 226, 224, 1, 0, 0, 0, 227, 37, 1, 0, 0, 0, 228, 229, 5, 10, 0, 0, 229,
		230, 3, 20, 10, 0, 230, 231, 3, 82, 41, 0, 231, 232, 3, 14, 7, 0, 232,
		241, 1, 0, 0, 0, 233, 234, 5, 10, 0, 0, 234, 235, 3, 20, 10, 0, 235, 236,
		3, 82, 41, 0, 236, 237, 5, 7, 0, 0, 237, 238, 3, 20, 10, 0, 238, 239, 3,
		14, 7, 0, 239, 241, 1, 0, 0, 0, 240, 228, 1, 0, 0, 0, 240, 233, 1, 0, 0,
		0, 241, 39, 1, 0, 0, 0, 242, 243, 5, 11, 0, 0, 243, 244, 3, 20, 10, 0,
		244, 245, 3, 82, 41, 0, 245, 246, 5, 34, 0, 0, 246, 247, 3, 42, 21, 0,
		247, 248, 5, 35, 0, 0, 248, 41, 1, 0, 0, 0, 249, 254, 3, 44, 22, 0, 250,
		251, 3, 44, 22, 0, 251, 252, 3, 48, 24, 0, 252, 254, 1, 0, 0, 0, 253, 249,
		1, 0, 0, 0, 253, 250, 1, 0, 0, 0, 254, 43, 1, 0, 0, 0, 255, 260, 6, 22,
		-1, 0, 256, 257, 10, 1, 0, 0, 257, 259, 3, 46, 23, 0, 258, 256, 1, 0, 0,
		0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		45, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 264, 5, 12, 0, 0, 264, 265,
		3, 82, 41, 0, 265, 266, 3, 14, 7, 0, 266, 47, 1, 0, 0, 0, 267, 268, 5,
		13, 0, 0, 268, 269, 3, 14, 7, 0, 269, 49, 1, 0, 0, 0, 270, 271, 3, 60,
		30, 0, 271, 272, 5, 6, 0, 0, 272, 273, 3, 78, 39, 0, 273, 274, 5, 5, 0,
		0, 274, 282, 1, 0, 0, 0, 275, 276, 3, 52, 26, 0, 276, 277, 7, 0, 0, 0,
		277, 278, 5, 6, 0, 0, 278, 279, 3, 78, 39, 0, 279, 280, 5, 5, 0, 0, 280,
		282, 1, 0, 0, 0, 281, 270, 1, 0, 0, 0, 281, 275, 1, 0, 0, 0, 282, 51, 1,
		0, 0, 0, 283, 284, 6, 26, -1, 0, 284, 285, 3, 58, 29, 0, 285, 291, 1, 0,
		0, 0, 286, 287, 10, 1, 0, 0, 287, 288, 5, 2, 0, 0, 288, 290, 3, 58, 29,
		0, 289, 286, 1, 0, 0, 0, 290, 293, 1, 0, 0, 0, 291, 289, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 53, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 294, 295, 6,
		27, -1, 0, 295, 296, 3, 56, 28, 0, 296, 302, 1, 0, 0, 0, 297, 298, 10,
		1, 0, 0, 298, 299, 5, 2, 0, 0, 299, 301, 3, 56, 28, 0, 300, 297, 1, 0,
		0, 0, 301, 304, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0,
		303, 55, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 305, 306, 3, 58, 29, 0, 306,
		307, 5, 3, 0, 0, 307, 308, 3, 104, 52, 0, 308, 311, 1, 0, 0, 0, 309, 311,
		3, 58, 29, 0, 310, 305, 1, 0, 0, 0, 310, 309, 1, 0, 0, 0, 311, 57, 1, 0,
		0, 0, 312, 318, 3, 82, 41, 0, 313, 314, 5, 34, 0, 0, 314, 315, 3, 54, 27,
		0, 315, 316, 5, 35, 0, 0, 316, 318, 1, 0, 0, 0, 317, 312, 1, 0, 0, 0, 317,
		313, 1, 0, 0, 0, 318, 59, 1, 0, 0, 0, 319, 320, 6, 30, -1, 0, 320, 321,
		3, 66, 33, 0, 321, 327, 1, 0, 0, 0, 322, 323, 10, 1, 0, 0, 323, 324, 5,
		2, 0, 0, 324, 326, 3, 50, 25, 0, 325, 322, 1, 0, 0, 0, 326, 329, 1, 0,
		0, 0, 327, 325, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0, 328, 61, 1, 0, 0, 0,
		329, 327, 1, 0, 0, 0, 330, 331, 6, 31, -1, 0, 331, 332, 3, 64, 32, 0, 332,
		338, 1, 0, 0, 0, 333, 334, 10, 1, 0, 0, 334, 335, 5, 2, 0, 0, 335, 337,
		3, 64, 32, 0, 336, 333, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0, 338, 336, 1,
		0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 63, 1, 0, 0, 0, 340, 338, 1, 0, 0,
		0, 341, 342, 3, 66, 33, 0, 342, 343, 5, 3, 0, 0, 343, 344, 3, 104, 52,
		0, 344, 347, 1, 0, 0, 0, 345, 347, 3, 66, 33, 0, 346, 341, 1, 0, 0, 0,
		346, 345, 1, 0, 0, 0, 347, 65, 1, 0, 0, 0, 348, 357, 3, 82, 41, 0, 349,
		357, 3, 68, 34, 0, 350, 351, 5, 22, 0, 0, 351, 357, 3, 74, 37, 0, 352,
		353, 5, 34, 0, 0, 353, 354, 3, 62, 31, 0, 354, 355, 5, 35, 0, 0, 355, 357,
		1, 0, 0, 0, 356, 348, 1, 0, 0, 0, 356, 349, 1, 0, 0, 0, 356, 350, 1, 0,
		0, 0, 356, 352, 1, 0, 0, 0, 357, 67, 1, 0, 0, 0, 358, 359, 3, 98, 49, 0,
		359, 360, 3, 74, 37, 0, 360, 69, 1, 0, 0, 0, 361, 362, 6, 35, -1, 0, 362,
		363, 3, 72, 36, 0, 363, 369, 1, 0, 0, 0, 364, 365, 10, 1, 0, 0, 365, 366,
		5, 2, 0, 0, 366, 368, 3, 72, 36, 0, 367, 364, 1, 0, 0, 0, 368, 371, 1,
		0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 71, 1, 0, 0,
		0, 371, 369, 1, 0, 0, 0, 372, 373, 3, 74, 37, 0, 373, 374, 5, 3, 0, 0,
		374, 375, 3, 104, 52, 0, 375, 378, 1, 0, 0, 0, 376, 378, 3, 74, 37, 0,
		377, 372, 1, 0, 0, 0, 377, 376, 1, 0, 0, 0, 378, 73, 1, 0, 0, 0, 379, 385,
		3, 104, 52, 0, 380, 381, 5, 34, 0, 0, 381, 382, 3, 70, 35, 0, 382, 383,
		5, 35, 0, 0, 383, 385, 1, 0, 0, 0, 384, 379, 1, 0, 0, 0, 384, 380, 1, 0,
		0, 0, 385, 75, 1, 0, 0, 0, 386, 387, 3, 78, 39, 0, 387, 388, 5, 5, 0, 0,
		388, 77, 1, 0, 0, 0, 389, 390, 6, 39, -1, 0, 390, 391, 3, 82, 41, 0, 391,
		397, 1, 0, 0, 0, 392, 393, 10, 1, 0, 0, 393, 394, 5, 2, 0, 0, 394, 396,
		3, 82, 41, 0, 395, 392, 1, 0, 0, 0, 396, 399, 1, 0, 0, 0, 397, 395, 1,
		0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 79, 1, 0, 0, 0, 399, 397, 1, 0, 0,
		0, 400, 403, 3, 78, 39, 0, 401, 403, 1, 0, 0, 0, 402, 400, 1, 0, 0, 0,
		402, 401, 1, 0, 0, 0, 403, 81, 1, 0, 0, 0, 404, 405, 6, 41, -1, 0, 405,
		406, 5, 30, 0, 0, 406, 407, 3, 82, 41, 0, 407, 408, 5, 31, 0, 0, 408, 437,
		1, 0, 0, 0, 409, 410, 7, 1, 0, 0, 410, 437, 3, 82, 41, 20, 411, 437, 3,
		104, 52, 0, 412, 437, 3, 100, 50, 0, 413, 414, 3, 94, 47, 0, 414, 415,
		5, 30, 0, 0, 415, 416, 3, 8, 4, 0, 416, 417, 5, 31, 0, 0, 417, 418, 3,
		98, 49, 0, 418, 419, 3, 14, 7, 0, 419, 437, 1, 0, 0, 0, 420, 421, 3, 94,
		47, 0, 421, 422, 5, 30, 0, 0, 422, 423, 3, 8, 4, 0, 423, 424, 5, 31, 0,
		0, 424, 425, 3, 14, 7, 0, 425, 437, 1, 0, 0, 0, 426, 427, 3, 96, 48, 0,
		427, 428, 5, 34, 0, 0, 428, 429, 3, 84, 42, 0, 429, 430, 5, 35, 0, 0, 430,
		437, 1, 0, 0, 0, 431, 432, 3, 98, 49, 0, 432, 433, 5, 32, 0, 0, 433, 434,
		3, 90, 45, 0, 434, 435, 5, 33, 0, 0, 435, 437, 1, 0, 0, 0, 436, 404, 1,
		0, 0, 0, 436, 409, 1, 0, 0, 0, 436, 411, 1, 0, 0, 0, 436, 412, 1, 0, 0,
		0, 436, 413, 1, 0, 0, 0, 436, 420, 1, 0, 0, 0, 436, 426, 1, 0, 0, 0, 436,
		431, 1, 0, 0, 0, 437, 496, 1, 0, 0, 0, 438, 439, 10, 21, 0, 0, 439, 440,
		7, 2, 0, 0, 440, 495, 3, 82, 41, 22, 441, 442, 10, 19, 0, 0, 442, 443,
		7, 3, 0, 0, 443, 495, 3, 82, 41, 20, 444, 445, 10, 18, 0, 0, 445, 446,
		7, 4, 0, 0, 446, 495, 3, 82, 41, 19, 447, 448, 10, 17, 0, 0, 448, 449,
		7, 5, 0, 0, 449, 495, 3, 82, 41, 18, 450, 451, 10, 16, 0, 0, 451, 452,
		7, 6, 0, 0, 452, 495, 3, 82, 41, 17, 453, 454, 10, 15, 0, 0, 454, 455,
		5, 46, 0, 0, 455, 495, 3, 82, 41, 16, 456, 457, 10, 14, 0, 0, 457, 458,
		5, 45, 0, 0, 458, 495, 3, 82, 41, 15, 459, 460, 10, 13, 0, 0, 460, 461,
		5, 44, 0, 0, 461, 495, 3, 82, 41, 14, 462, 463, 10, 12, 0, 0, 463, 464,
		5, 30, 0, 0, 464, 465, 3, 80, 40, 0, 465, 466, 5, 31, 0, 0, 466, 495, 1,
		0, 0, 0, 467, 468, 10, 11, 0, 0, 468, 469, 5, 1, 0, 0, 469, 495, 3, 104,
		52, 0, 470, 471, 10, 10, 0, 0, 471, 472, 5, 32, 0, 0, 472, 473, 3, 82,
		41, 0, 473, 474, 5, 33, 0, 0, 474, 495, 1, 0, 0, 0, 475, 476, 10, 9, 0,
		0, 476, 477, 5, 32, 0, 0, 477, 478, 3, 82, 41, 0, 478, 479, 5, 3, 0, 0,
		479, 480, 5, 33, 0, 0, 480, 495, 1, 0, 0, 0, 481, 482, 10, 8, 0, 0, 482,
		483, 5, 32, 0, 0, 483, 484, 5, 3, 0, 0, 484, 485, 3, 82, 41, 0, 485, 486,
		5, 33, 0, 0, 486, 495, 1, 0, 0, 0, 487, 488, 10, 7, 0, 0, 488, 489, 5,
		32, 0, 0, 489, 490, 3, 82, 41, 0, 490, 491, 5, 3, 0, 0, 491, 492, 3, 82,
		41, 0, 492, 493, 5, 33, 0, 0, 493, 495, 1, 0, 0, 0, 494, 438, 1, 0, 0,
		0, 494, 441, 1, 0, 0, 0, 494, 444, 1, 0, 0, 0, 494, 447, 1, 0, 0, 0, 494,
		450, 1, 0, 0, 0, 494, 453, 1, 0, 0, 0, 494, 456, 1, 0, 0, 0, 494, 459,
		1, 0, 0, 0, 494, 462, 1, 0, 0, 0, 494, 467, 1, 0, 0, 0, 494, 470, 1, 0,
		0, 0, 494, 475, 1, 0, 0, 0, 494, 481, 1, 0, 0, 0, 494, 487, 1, 0, 0, 0,
		495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0, 497,
		83, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 499, 505, 3, 86, 43, 0, 500, 501,
		3, 86, 43, 0, 501, 502, 5, 2, 0, 0, 502, 505, 1, 0, 0, 0, 503, 505, 3,
		80, 40, 0, 504, 499, 1, 0, 0, 0, 504, 500, 1, 0, 0, 0, 504, 503, 1, 0,
		0, 0, 505, 85, 1, 0, 0, 0, 506, 507, 6, 43, -1, 0, 507, 508, 3, 88, 44,
		0, 508, 514, 1, 0, 0, 0, 509, 510, 10, 1, 0, 0, 510, 511, 5, 2, 0, 0, 511,
		513, 3, 88, 44, 0, 512, 509, 1, 0, 0, 0, 513, 516, 1, 0, 0, 0, 514, 512,
		1, 0, 0, 0, 514, 515, 1, 0, 0, 0, 515, 87, 1, 0, 0, 0, 516, 514, 1, 0,
		0, 0, 517, 518, 3, 104, 52, 0, 518, 519, 5, 3, 0, 0, 519, 520, 3, 82, 41,
		0, 520, 89, 1, 0, 0, 0, 521, 527, 6, 45, -1, 0, 522, 523, 10, 1, 0, 0,
		523, 524, 5, 2, 0, 0, 524, 526, 3, 92, 46, 0, 525, 522, 1, 0, 0, 0, 526,
		529, 1, 0, 0, 0, 527, 525, 1, 0, 0, 0, 527, 528, 1, 0, 0, 0, 528, 91, 1,
		0, 0, 0, 529, 527, 1, 0, 0, 0, 530, 531, 3, 82, 41, 0, 531, 532, 5, 3,
		0, 0, 532, 533, 3, 82, 41, 0, 533, 93, 1, 0, 0, 0, 534, 537, 5, 25, 0,
		0, 535, 537, 3, 98, 49, 0, 536, 534, 1, 0, 0, 0, 536, 535, 1, 0, 0, 0,
		537, 95, 1, 0, 0, 0, 538, 542, 5, 26, 0, 0, 539, 542, 5, 27, 0, 0, 540,
		542, 3, 98, 49, 0, 541, 538, 1, 0, 0, 0, 541, 539, 1, 0, 0, 0, 541, 540,
		1, 0, 0, 0, 542, 97, 1, 0, 0, 0, 543, 544, 6, 49, -1, 0, 544, 545, 3, 104,
		52, 0, 545, 551, 1, 0, 0, 0, 546, 547, 10, 1, 0, 0, 547, 548, 5, 1, 0,
		0, 548, 550, 3, 104, 52, 0, 549, 546, 1, 0, 0, 0, 550, 553, 1, 0, 0, 0,
		551, 549, 1, 0, 0, 0, 551, 552, 1, 0, 0, 0, 552, 99, 1, 0, 0, 0, 553, 551,
		1, 0, 0, 0, 554, 559, 3, 102, 51, 0, 555, 556, 3, 98, 49, 0, 556, 557,
		3, 102, 51, 0, 557, 559, 1, 0, 0, 0, 558, 554, 1, 0, 0, 0, 558, 555, 1,
		0, 0, 0, 559, 101, 1, 0, 0, 0, 560, 561, 5, 59, 0, 0, 561, 103, 1, 0, 0,
		0, 562, 563, 5, 58, 0, 0, 563, 105, 1, 0, 0, 0, 40, 114, 134, 138, 148,
		157, 168, 175, 182, 187, 192, 207, 220, 226, 240, 253, 260, 281, 291, 302,
		310, 317, 327, 338, 346, 356, 369, 377, 384, 397, 402, 436, 494, 496, 504,
		514, 527, 536, 541, 551, 558,
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
	CaliburnParserIdentifierToken   = 58
	CaliburnParserLiteralToken      = 59
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_module                                = 0
	CaliburnParserRULE_definitions                           = 1
	CaliburnParserRULE_definition                            = 2
	CaliburnParserRULE_function_definition                   = 3
	CaliburnParserRULE_parameters                            = 4
	CaliburnParserRULE_parameters_list                       = 5
	CaliburnParserRULE_parameter                             = 6
	CaliburnParserRULE_block                                 = 7
	CaliburnParserRULE_statements                            = 8
	CaliburnParserRULE_statement                             = 9
	CaliburnParserRULE_inline_statements                     = 10
	CaliburnParserRULE_inline_statement                      = 11
	CaliburnParserRULE_jump_statement                        = 12
	CaliburnParserRULE_return_statement                      = 13
	CaliburnParserRULE_break_statement                       = 14
	CaliburnParserRULE_continue_statement                    = 15
	CaliburnParserRULE_control_statement                     = 16
	CaliburnParserRULE_if_statement                          = 17
	CaliburnParserRULE_else_statement                        = 18
	CaliburnParserRULE_for_statement                         = 19
	CaliburnParserRULE_switch_statement                      = 20
	CaliburnParserRULE_case_blocks                           = 21
	CaliburnParserRULE_option_case_blocks                    = 22
	CaliburnParserRULE_option_case_block                     = 23
	CaliburnParserRULE_default_case_block                    = 24
	CaliburnParserRULE_assign_statement                      = 25
	CaliburnParserRULE_assign_expressions                    = 26
	CaliburnParserRULE_aliasable_assign_expressions          = 27
	CaliburnParserRULE_aliasable_assign_expression           = 28
	CaliburnParserRULE_assign_expression                     = 29
	CaliburnParserRULE_assign_declarations                   = 30
	CaliburnParserRULE_aliasable_assign_declarations         = 31
	CaliburnParserRULE_aliasable_assign_declaration          = 32
	CaliburnParserRULE_assign_declaration                    = 33
	CaliburnParserRULE_typed_assign_declaration              = 34
	CaliburnParserRULE_aliasable_untyped_assign_declarations = 35
	CaliburnParserRULE_aliasable_untyped_assign_declaration  = 36
	CaliburnParserRULE_untyped_assign_declaration            = 37
	CaliburnParserRULE_expression_statement                  = 38
	CaliburnParserRULE_expressions                           = 39
	CaliburnParserRULE_expressions_optional                  = 40
	CaliburnParserRULE_expression                            = 41
	CaliburnParserRULE_struct_values                         = 42
	CaliburnParserRULE_named_struct_values                   = 43
	CaliburnParserRULE_named_struct_value                    = 44
	CaliburnParserRULE_collection_values                     = 45
	CaliburnParserRULE_collection_value                      = 46
	CaliburnParserRULE_function_type                         = 47
	CaliburnParserRULE_struct_type                           = 48
	CaliburnParserRULE_type_expression                       = 49
	CaliburnParserRULE_literal_atom                          = 50
	CaliburnParserRULE_literal                               = 51
	CaliburnParserRULE_identifier                            = 52
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
		p.SetState(106)
		p.definitions(0)
	}
	{
		p.SetState(107)
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
	p.SetState(114)
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
			p.SetState(110)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(111)
				p.Definition()
			}

		}
		p.SetState(116)
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
		p.SetState(117)
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
}

func NewFunctionDefinitionNoReturnTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefinitionNoReturnTypeContext {
	var p = new(FunctionDefinitionNoReturnTypeContext)

	InitEmptyFunction_definitionContext(&p.Function_definitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_definitionContext))

	return p
}

func (s *FunctionDefinitionNoReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *FunctionDefinitionNoReturnTypeContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
}

func NewFunctionDefinitionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	InitEmptyFunction_definitionContext(&p.Function_definitionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Function_definitionContext))

	return p
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *FunctionDefinitionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunctionDefinitionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Function_type()
		}
		{
			p.SetState(120)
			p.Identifier()
		}
		{
			p.SetState(121)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Parameters()
		}
		{
			p.SetState(123)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.type_expression(0)
		}
		{
			p.SetState(125)
			p.Block()
		}

	case 2:
		localctx = NewFunctionDefinitionNoReturnTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.Function_type()
		}
		{
			p.SetState(128)
			p.Identifier()
		}
		{
			p.SetState(129)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(130)
			p.Parameters()
		}
		{
			p.SetState(131)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserR_PAREN, CaliburnParserR_C_BRACK:
		localctx = NewParametersEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)

	case CaliburnParserL_C_BRACK, CaliburnParserIdentifierToken:
		localctx = NewParametersFilledContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
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
		p.SetState(141)
		p.Parameter()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(148)
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
			p.SetState(143)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(144)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(145)
				p.Parameter()
			}

		}
		p.SetState(150)
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
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.Typed_assign_declaration()
		}

	case 2:
		localctx = NewUntypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(152)
			p.Untyped_assign_declaration()
		}

	case 3:
		localctx = NewStructDestrutureParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(153)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(154)
			p.Parameters()
		}
		{
			p.SetState(155)
			p.Match(CaliburnParserR_C_BRACK)
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

	// Getter signatures
	L_C_BRACK() antlr.TerminalNode
	Statements() IStatementsContext
	R_C_BRACK() antlr.TerminalNode

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

func (s *BlockContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
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

func (s *BlockContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CaliburnParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.statements(0)
	}
	{
		p.SetState(161)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
			p.SetState(164)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(165)
				p.Statement()
			}

		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Expression_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
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
	_startState := 20
	p.EnterRecursionRule(localctx, 20, CaliburnParserRULE_inline_statements, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewInlineStatementsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
			p.SetState(178)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(179)
				p.Inline_statement()
			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 22, CaliburnParserRULE_inline_statement)
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(185)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(186)
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
	p.EnterRule(localctx, 24, CaliburnParserRULE_jump_statement)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserRETURN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Return_statement()
		}

	case CaliburnParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Break_statement()
		}

	case CaliburnParserCONTINUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
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
	p.EnterRule(localctx, 26, CaliburnParserRULE_return_statement)
	localctx = NewReturnStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(CaliburnParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.expression(0)
	}
	{
		p.SetState(196)
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
	p.EnterRule(localctx, 28, CaliburnParserRULE_break_statement)
	localctx = NewBreakStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(CaliburnParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
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
	p.EnterRule(localctx, 30, CaliburnParserRULE_continue_statement)
	localctx = NewContinueStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(CaliburnParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
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
	p.EnterRule(localctx, 32, CaliburnParserRULE_control_statement)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.If_statement()
		}

	case CaliburnParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(205)
			p.For_statement()
		}

	case CaliburnParserSWITCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
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
	p.EnterRule(localctx, 34, CaliburnParserRULE_if_statement)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(CaliburnParserIF)
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
			p.Block()
		}

	case 2:
		localctx = NewIfElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(CaliburnParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.inline_statements(0)
		}
		{
			p.SetState(216)
			p.expression(0)
		}
		{
			p.SetState(217)
			p.Block()
		}
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 36, CaliburnParserRULE_else_statement)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewElseStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Block()
		}

	case 2:
		localctx = NewElseIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
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
	p.EnterRule(localctx, 38, CaliburnParserRULE_for_statement)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(CaliburnParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.inline_statements(0)
		}
		{
			p.SetState(230)
			p.expression(0)
		}
		{
			p.SetState(231)
			p.Block()
		}

	case 2:
		localctx = NewForStatementWithAfterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.Match(CaliburnParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.inline_statements(0)
		}
		{
			p.SetState(235)
			p.expression(0)
		}
		{
			p.SetState(236)
			p.Match(CaliburnParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.inline_statements(0)
		}
		{
			p.SetState(238)
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
	p.EnterRule(localctx, 40, CaliburnParserRULE_switch_statement)
	localctx = NewSwitchStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(CaliburnParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.inline_statements(0)
	}
	{
		p.SetState(244)
		p.expression(0)
	}
	{
		p.SetState(245)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Case_blocks()
	}
	{
		p.SetState(247)
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

type CaseBlocksDefaultContext struct {
	Case_blocksContext
}

func NewCaseBlocksDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseBlocksDefaultContext {
	var p = new(CaseBlocksDefaultContext)

	InitEmptyCase_blocksContext(&p.Case_blocksContext)
	p.parser = parser
	p.CopyAll(ctx.(*Case_blocksContext))

	return p
}

func (s *CaseBlocksDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlocksDefaultContext) Option_case_blocks() IOption_case_blocksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_case_blocksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_case_blocksContext)
}

func (s *CaseBlocksDefaultContext) Default_case_block() IDefault_case_blockContext {
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

func (s *CaseBlocksDefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCaseBlocksDefault(s)
	}
}

func (s *CaseBlocksDefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCaseBlocksDefault(s)
	}
}

func (s *CaseBlocksDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCaseBlocksDefault(s)

	default:
		return t.VisitChildren(s)
	}
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

func (s *CaseBlocksContext) Option_case_blocks() IOption_case_blocksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_case_blocksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_case_blocksContext)
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
	p.EnterRule(localctx, 42, CaliburnParserRULE_case_blocks)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCaseBlocksContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.option_case_blocks(0)
		}

	case 2:
		localctx = NewCaseBlocksDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.option_case_blocks(0)
		}
		{
			p.SetState(251)
			p.Default_case_block()
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

// IOption_case_blocksContext is an interface to support dynamic dispatch.
type IOption_case_blocksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOption_case_blocksContext differentiates from other interfaces.
	IsOption_case_blocksContext()
}

type Option_case_blocksContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOption_case_blocksContext() *Option_case_blocksContext {
	var p = new(Option_case_blocksContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_option_case_blocks
	return p
}

func InitEmptyOption_case_blocksContext(p *Option_case_blocksContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_option_case_blocks
}

func (*Option_case_blocksContext) IsOption_case_blocksContext() {}

func NewOption_case_blocksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Option_case_blocksContext {
	var p = new(Option_case_blocksContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_option_case_blocks

	return p
}

func (s *Option_case_blocksContext) GetParser() antlr.Parser { return s.parser }

func (s *Option_case_blocksContext) CopyAll(ctx *Option_case_blocksContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Option_case_blocksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Option_case_blocksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OptionCaseBlocksAppendContext struct {
	Option_case_blocksContext
}

func NewOptionCaseBlocksAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionCaseBlocksAppendContext {
	var p = new(OptionCaseBlocksAppendContext)

	InitEmptyOption_case_blocksContext(&p.Option_case_blocksContext)
	p.parser = parser
	p.CopyAll(ctx.(*Option_case_blocksContext))

	return p
}

func (s *OptionCaseBlocksAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionCaseBlocksAppendContext) Option_case_blocks() IOption_case_blocksContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_case_blocksContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_case_blocksContext)
}

func (s *OptionCaseBlocksAppendContext) Option_case_block() IOption_case_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOption_case_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOption_case_blockContext)
}

func (s *OptionCaseBlocksAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterOptionCaseBlocksAppend(s)
	}
}

func (s *OptionCaseBlocksAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitOptionCaseBlocksAppend(s)
	}
}

func (s *OptionCaseBlocksAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitOptionCaseBlocksAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type OptionCaseBlocksInitialContext struct {
	Option_case_blocksContext
}

func NewOptionCaseBlocksInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionCaseBlocksInitialContext {
	var p = new(OptionCaseBlocksInitialContext)

	InitEmptyOption_case_blocksContext(&p.Option_case_blocksContext)
	p.parser = parser
	p.CopyAll(ctx.(*Option_case_blocksContext))

	return p
}

func (s *OptionCaseBlocksInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionCaseBlocksInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterOptionCaseBlocksInitial(s)
	}
}

func (s *OptionCaseBlocksInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitOptionCaseBlocksInitial(s)
	}
}

func (s *OptionCaseBlocksInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitOptionCaseBlocksInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Option_case_blocks() (localctx IOption_case_blocksContext) {
	return p.option_case_blocks(0)
}

func (p *CaliburnParser) option_case_blocks(_p int) (localctx IOption_case_blocksContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewOption_case_blocksContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IOption_case_blocksContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, CaliburnParserRULE_option_case_blocks, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewOptionCaseBlocksInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
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
			localctx = NewOptionCaseBlocksAppendContext(p, NewOption_case_blocksContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_option_case_blocks)
			p.SetState(256)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(257)
				p.Option_case_block()
			}

		}
		p.SetState(262)
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
	p.EnterRule(localctx, 46, CaliburnParserRULE_option_case_block)
	localctx = NewOptionCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(CaliburnParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(264)
		p.expression(0)
	}
	{
		p.SetState(265)
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
	p.EnterRule(localctx, 48, CaliburnParserRULE_default_case_block)
	localctx = NewDefaultCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(267)
		p.Match(CaliburnParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
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
	p.EnterRule(localctx, 50, CaliburnParserRULE_assign_statement)
	var _la int

	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.assign_declarations(0)
		}
		{
			p.SetState(271)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.expressions(0)
		}
		{
			p.SetState(273)
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
			p.SetState(275)
			p.assign_expressions(0)
		}
		{
			p.SetState(276)

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
			p.SetState(277)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.expressions(0)
		}
		{
			p.SetState(279)
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

type AssignExpressionsAppendContext struct {
	Assign_expressionsContext
}

func NewAssignExpressionsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExpressionsAppendContext {
	var p = new(AssignExpressionsAppendContext)

	InitEmptyAssign_expressionsContext(&p.Assign_expressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionsContext))

	return p
}

func (s *AssignExpressionsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionsAppendContext) Assign_expressions() IAssign_expressionsContext {
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

func (s *AssignExpressionsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *AssignExpressionsAppendContext) Assign_expression() IAssign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *AssignExpressionsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignExpressionsAppend(s)
	}
}

func (s *AssignExpressionsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignExpressionsAppend(s)
	}
}

func (s *AssignExpressionsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignExpressionsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignExpressionsInitialContext struct {
	Assign_expressionsContext
}

func NewAssignExpressionsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignExpressionsInitialContext {
	var p = new(AssignExpressionsInitialContext)

	InitEmptyAssign_expressionsContext(&p.Assign_expressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_expressionsContext))

	return p
}

func (s *AssignExpressionsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExpressionsInitialContext) Assign_expression() IAssign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *AssignExpressionsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignExpressionsInitial(s)
	}
}

func (s *AssignExpressionsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignExpressionsInitial(s)
	}
}

func (s *AssignExpressionsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignExpressionsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_expressions() (localctx IAssign_expressionsContext) {
	return p.assign_expressions(0)
}

func (p *CaliburnParser) assign_expressions(_p int) (localctx IAssign_expressionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAssign_expressionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAssign_expressionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, CaliburnParserRULE_assign_expressions, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAssignExpressionsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(284)
		p.Assign_expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAssignExpressionsAppendContext(p, NewAssign_expressionsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_assign_expressions)
			p.SetState(286)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(287)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(288)
				p.Assign_expression()
			}

		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
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

// IAliasable_assign_expressionsContext is an interface to support dynamic dispatch.
type IAliasable_assign_expressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_assign_expressionsContext differentiates from other interfaces.
	IsAliasable_assign_expressionsContext()
}

type Aliasable_assign_expressionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_assign_expressionsContext() *Aliasable_assign_expressionsContext {
	var p = new(Aliasable_assign_expressionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expressions
	return p
}

func InitEmptyAliasable_assign_expressionsContext(p *Aliasable_assign_expressionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expressions
}

func (*Aliasable_assign_expressionsContext) IsAliasable_assign_expressionsContext() {}

func NewAliasable_assign_expressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_assign_expressionsContext {
	var p = new(Aliasable_assign_expressionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expressions

	return p
}

func (s *Aliasable_assign_expressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_assign_expressionsContext) CopyAll(ctx *Aliasable_assign_expressionsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_assign_expressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_assign_expressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AliasableAssignExpressionsInitialContext struct {
	Aliasable_assign_expressionsContext
}

func NewAliasableAssignExpressionsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableAssignExpressionsInitialContext {
	var p = new(AliasableAssignExpressionsInitialContext)

	InitEmptyAliasable_assign_expressionsContext(&p.Aliasable_assign_expressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_expressionsContext))

	return p
}

func (s *AliasableAssignExpressionsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableAssignExpressionsInitialContext) Aliasable_assign_expression() IAliasable_assign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_expressionContext)
}

func (s *AliasableAssignExpressionsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableAssignExpressionsInitial(s)
	}
}

func (s *AliasableAssignExpressionsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableAssignExpressionsInitial(s)
	}
}

func (s *AliasableAssignExpressionsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableAssignExpressionsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasableAssignExpressionsAppendContext struct {
	Aliasable_assign_expressionsContext
}

func NewAliasableAssignExpressionsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableAssignExpressionsAppendContext {
	var p = new(AliasableAssignExpressionsAppendContext)

	InitEmptyAliasable_assign_expressionsContext(&p.Aliasable_assign_expressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_expressionsContext))

	return p
}

func (s *AliasableAssignExpressionsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableAssignExpressionsAppendContext) Aliasable_assign_expressions() IAliasable_assign_expressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_expressionsContext)
}

func (s *AliasableAssignExpressionsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *AliasableAssignExpressionsAppendContext) Aliasable_assign_expression() IAliasable_assign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_expressionContext)
}

func (s *AliasableAssignExpressionsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableAssignExpressionsAppend(s)
	}
}

func (s *AliasableAssignExpressionsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableAssignExpressionsAppend(s)
	}
}

func (s *AliasableAssignExpressionsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableAssignExpressionsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_assign_expressions() (localctx IAliasable_assign_expressionsContext) {
	return p.aliasable_assign_expressions(0)
}

func (p *CaliburnParser) aliasable_assign_expressions(_p int) (localctx IAliasable_assign_expressionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAliasable_assign_expressionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAliasable_assign_expressionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, CaliburnParserRULE_aliasable_assign_expressions, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAliasableAssignExpressionsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(295)
		p.Aliasable_assign_expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAliasableAssignExpressionsAppendContext(p, NewAliasable_assign_expressionsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_aliasable_assign_expressions)
			p.SetState(297)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(298)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(299)
				p.Aliasable_assign_expression()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
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

// IAliasable_assign_expressionContext is an interface to support dynamic dispatch.
type IAliasable_assign_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_assign_expressionContext differentiates from other interfaces.
	IsAliasable_assign_expressionContext()
}

type Aliasable_assign_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_assign_expressionContext() *Aliasable_assign_expressionContext {
	var p = new(Aliasable_assign_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expression
	return p
}

func InitEmptyAliasable_assign_expressionContext(p *Aliasable_assign_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expression
}

func (*Aliasable_assign_expressionContext) IsAliasable_assign_expressionContext() {}

func NewAliasable_assign_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_assign_expressionContext {
	var p = new(Aliasable_assign_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_expression

	return p
}

func (s *Aliasable_assign_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_assign_expressionContext) CopyAll(ctx *Aliasable_assign_expressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_assign_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_assign_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaliasedAssignExpressionContext struct {
	Aliasable_assign_expressionContext
}

func NewUnaliasedAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaliasedAssignExpressionContext {
	var p = new(UnaliasedAssignExpressionContext)

	InitEmptyAliasable_assign_expressionContext(&p.Aliasable_assign_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_expressionContext))

	return p
}

func (s *UnaliasedAssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaliasedAssignExpressionContext) Assign_expression() IAssign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *UnaliasedAssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUnaliasedAssignExpression(s)
	}
}

func (s *UnaliasedAssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUnaliasedAssignExpression(s)
	}
}

func (s *UnaliasedAssignExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUnaliasedAssignExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasedAssignExpressionContext struct {
	Aliasable_assign_expressionContext
}

func NewAliasedAssignExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedAssignExpressionContext {
	var p = new(AliasedAssignExpressionContext)

	InitEmptyAliasable_assign_expressionContext(&p.Aliasable_assign_expressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_expressionContext))

	return p
}

func (s *AliasedAssignExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedAssignExpressionContext) Assign_expression() IAssign_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_expressionContext)
}

func (s *AliasedAssignExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *AliasedAssignExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AliasedAssignExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasedAssignExpression(s)
	}
}

func (s *AliasedAssignExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasedAssignExpression(s)
	}
}

func (s *AliasedAssignExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasedAssignExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_assign_expression() (localctx IAliasable_assign_expressionContext) {
	localctx = NewAliasable_assign_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CaliburnParserRULE_aliasable_assign_expression)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.Assign_expression()
		}
		{
			p.SetState(306)
			p.Match(CaliburnParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Identifier()
		}

	case 2:
		localctx = NewUnaliasedAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Assign_expression()
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

func (s *StructDestrutureAssignExpressionContext) Aliasable_assign_expressions() IAliasable_assign_expressionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_expressionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_expressionsContext)
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
	p.EnterRule(localctx, 58, CaliburnParserRULE_assign_expression)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserTUPLE, CaliburnParserL_PAREN, CaliburnParserOP_ADD, CaliburnParserOP_SUB, CaliburnParserOP_NOT, CaliburnParserIdentifierToken, CaliburnParserLiteralToken:
		localctx = NewExpressionAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(312)
			p.expression(0)
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewStructDestrutureAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(313)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.aliasable_assign_expressions(0)
		}
		{
			p.SetState(315)
			p.Match(CaliburnParserR_C_BRACK)
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

type AssignDeclarationsAppendContext struct {
	Assign_declarationsContext
}

func NewAssignDeclarationsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDeclarationsAppendContext {
	var p = new(AssignDeclarationsAppendContext)

	InitEmptyAssign_declarationsContext(&p.Assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationsContext))

	return p
}

func (s *AssignDeclarationsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDeclarationsAppendContext) Assign_declarations() IAssign_declarationsContext {
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

func (s *AssignDeclarationsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *AssignDeclarationsAppendContext) Assign_statement() IAssign_statementContext {
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

func (s *AssignDeclarationsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignDeclarationsAppend(s)
	}
}

func (s *AssignDeclarationsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignDeclarationsAppend(s)
	}
}

func (s *AssignDeclarationsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignDeclarationsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignDeclarationsInitialContext struct {
	Assign_declarationsContext
}

func NewAssignDeclarationsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignDeclarationsInitialContext {
	var p = new(AssignDeclarationsInitialContext)

	InitEmptyAssign_declarationsContext(&p.Assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationsContext))

	return p
}

func (s *AssignDeclarationsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignDeclarationsInitialContext) Assign_declaration() IAssign_declarationContext {
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

func (s *AssignDeclarationsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssignDeclarationsInitial(s)
	}
}

func (s *AssignDeclarationsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssignDeclarationsInitial(s)
	}
}

func (s *AssignDeclarationsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssignDeclarationsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_declarations() (localctx IAssign_declarationsContext) {
	return p.assign_declarations(0)
}

func (p *CaliburnParser) assign_declarations(_p int) (localctx IAssign_declarationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAssign_declarationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAssign_declarationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, CaliburnParserRULE_assign_declarations, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAssignDeclarationsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(320)
		p.Assign_declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAssignDeclarationsAppendContext(p, NewAssign_declarationsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_assign_declarations)
			p.SetState(322)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(323)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(324)
				p.Assign_statement()
			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext())
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

// IAliasable_assign_declarationsContext is an interface to support dynamic dispatch.
type IAliasable_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_assign_declarationsContext differentiates from other interfaces.
	IsAliasable_assign_declarationsContext()
}

type Aliasable_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_assign_declarationsContext() *Aliasable_assign_declarationsContext {
	var p = new(Aliasable_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declarations
	return p
}

func InitEmptyAliasable_assign_declarationsContext(p *Aliasable_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declarations
}

func (*Aliasable_assign_declarationsContext) IsAliasable_assign_declarationsContext() {}

func NewAliasable_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_assign_declarationsContext {
	var p = new(Aliasable_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declarations

	return p
}

func (s *Aliasable_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_assign_declarationsContext) CopyAll(ctx *Aliasable_assign_declarationsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AliasableAssignDeclarationsAppendContext struct {
	Aliasable_assign_declarationsContext
}

func NewAliasableAssignDeclarationsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableAssignDeclarationsAppendContext {
	var p = new(AliasableAssignDeclarationsAppendContext)

	InitEmptyAliasable_assign_declarationsContext(&p.Aliasable_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_declarationsContext))

	return p
}

func (s *AliasableAssignDeclarationsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableAssignDeclarationsAppendContext) Aliasable_assign_declarations() IAliasable_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_declarationsContext)
}

func (s *AliasableAssignDeclarationsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *AliasableAssignDeclarationsAppendContext) Aliasable_assign_declaration() IAliasable_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_declarationContext)
}

func (s *AliasableAssignDeclarationsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableAssignDeclarationsAppend(s)
	}
}

func (s *AliasableAssignDeclarationsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableAssignDeclarationsAppend(s)
	}
}

func (s *AliasableAssignDeclarationsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableAssignDeclarationsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasableAssignDeclarationsInitialContext struct {
	Aliasable_assign_declarationsContext
}

func NewAliasableAssignDeclarationsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableAssignDeclarationsInitialContext {
	var p = new(AliasableAssignDeclarationsInitialContext)

	InitEmptyAliasable_assign_declarationsContext(&p.Aliasable_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_declarationsContext))

	return p
}

func (s *AliasableAssignDeclarationsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableAssignDeclarationsInitialContext) Aliasable_assign_declaration() IAliasable_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_declarationContext)
}

func (s *AliasableAssignDeclarationsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableAssignDeclarationsInitial(s)
	}
}

func (s *AliasableAssignDeclarationsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableAssignDeclarationsInitial(s)
	}
}

func (s *AliasableAssignDeclarationsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableAssignDeclarationsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_assign_declarations() (localctx IAliasable_assign_declarationsContext) {
	return p.aliasable_assign_declarations(0)
}

func (p *CaliburnParser) aliasable_assign_declarations(_p int) (localctx IAliasable_assign_declarationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAliasable_assign_declarationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAliasable_assign_declarationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, CaliburnParserRULE_aliasable_assign_declarations, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAliasableAssignDeclarationsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(331)
		p.Aliasable_assign_declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAliasableAssignDeclarationsAppendContext(p, NewAliasable_assign_declarationsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_aliasable_assign_declarations)
			p.SetState(333)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(334)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(335)
				p.Aliasable_assign_declaration()
			}

		}
		p.SetState(340)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
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

// IAliasable_assign_declarationContext is an interface to support dynamic dispatch.
type IAliasable_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_assign_declarationContext differentiates from other interfaces.
	IsAliasable_assign_declarationContext()
}

type Aliasable_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_assign_declarationContext() *Aliasable_assign_declarationContext {
	var p = new(Aliasable_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declaration
	return p
}

func InitEmptyAliasable_assign_declarationContext(p *Aliasable_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declaration
}

func (*Aliasable_assign_declarationContext) IsAliasable_assign_declarationContext() {}

func NewAliasable_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_assign_declarationContext {
	var p = new(Aliasable_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_assign_declaration

	return p
}

func (s *Aliasable_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_assign_declarationContext) CopyAll(ctx *Aliasable_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaliasedAssignDeclarationContext struct {
	Aliasable_assign_declarationContext
}

func NewUnaliasedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaliasedAssignDeclarationContext {
	var p = new(UnaliasedAssignDeclarationContext)

	InitEmptyAliasable_assign_declarationContext(&p.Aliasable_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_declarationContext))

	return p
}

func (s *UnaliasedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaliasedAssignDeclarationContext) Assign_declaration() IAssign_declarationContext {
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

func (s *UnaliasedAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUnaliasedAssignDeclaration(s)
	}
}

func (s *UnaliasedAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUnaliasedAssignDeclaration(s)
	}
}

func (s *UnaliasedAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUnaliasedAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasedAssignDeclarationContext struct {
	Aliasable_assign_declarationContext
}

func NewAliasedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedAssignDeclarationContext {
	var p = new(AliasedAssignDeclarationContext)

	InitEmptyAliasable_assign_declarationContext(&p.Aliasable_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_assign_declarationContext))

	return p
}

func (s *AliasedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedAssignDeclarationContext) Assign_declaration() IAssign_declarationContext {
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

func (s *AliasedAssignDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *AliasedAssignDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AliasedAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasedAssignDeclaration(s)
	}
}

func (s *AliasedAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasedAssignDeclaration(s)
	}
}

func (s *AliasedAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasedAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_assign_declaration() (localctx IAliasable_assign_declarationContext) {
	localctx = NewAliasable_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CaliburnParserRULE_aliasable_assign_declaration)
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.Assign_declaration()
		}
		{
			p.SetState(342)
			p.Match(CaliburnParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Identifier()
		}

	case 2:
		localctx = NewUnaliasedAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.Assign_declaration()
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

func (s *StructDestrutureAssignDeclarationContext) Aliasable_assign_declarations() IAliasable_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_assign_declarationsContext)
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

func (p *CaliburnParser) Assign_declaration() (localctx IAssign_declarationContext) {
	localctx = NewAssign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CaliburnParserRULE_assign_declaration)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.expression(0)
		}

	case 2:
		localctx = NewTypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Typed_assign_declaration()
		}

	case 3:
		localctx = NewUntypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.Match(CaliburnParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.Untyped_assign_declaration()
		}

	case 4:
		localctx = NewStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(352)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.aliasable_assign_declarations(0)
		}
		{
			p.SetState(354)
			p.Match(CaliburnParserR_C_BRACK)
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
}

func NewTypedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedAssignDeclarationContext {
	var p = new(TypedAssignDeclarationContext)

	InitEmptyTyped_assign_declarationContext(&p.Typed_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Typed_assign_declarationContext))

	return p
}

func (s *TypedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
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
	p.EnterRule(localctx, 68, CaliburnParserRULE_typed_assign_declaration)
	localctx = NewTypedAssignDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.type_expression(0)
	}
	{
		p.SetState(359)
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

// IAliasable_untyped_assign_declarationsContext is an interface to support dynamic dispatch.
type IAliasable_untyped_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_untyped_assign_declarationsContext differentiates from other interfaces.
	IsAliasable_untyped_assign_declarationsContext()
}

type Aliasable_untyped_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_untyped_assign_declarationsContext() *Aliasable_untyped_assign_declarationsContext {
	var p = new(Aliasable_untyped_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declarations
	return p
}

func InitEmptyAliasable_untyped_assign_declarationsContext(p *Aliasable_untyped_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declarations
}

func (*Aliasable_untyped_assign_declarationsContext) IsAliasable_untyped_assign_declarationsContext() {
}

func NewAliasable_untyped_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_untyped_assign_declarationsContext {
	var p = new(Aliasable_untyped_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declarations

	return p
}

func (s *Aliasable_untyped_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_untyped_assign_declarationsContext) CopyAll(ctx *Aliasable_untyped_assign_declarationsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_untyped_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_untyped_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AliasableUntypedAssignDeclarationsInitialContext struct {
	Aliasable_untyped_assign_declarationsContext
}

func NewAliasableUntypedAssignDeclarationsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableUntypedAssignDeclarationsInitialContext {
	var p = new(AliasableUntypedAssignDeclarationsInitialContext)

	InitEmptyAliasable_untyped_assign_declarationsContext(&p.Aliasable_untyped_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_untyped_assign_declarationsContext))

	return p
}

func (s *AliasableUntypedAssignDeclarationsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableUntypedAssignDeclarationsInitialContext) Aliasable_untyped_assign_declaration() IAliasable_untyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_untyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_untyped_assign_declarationContext)
}

func (s *AliasableUntypedAssignDeclarationsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableUntypedAssignDeclarationsInitial(s)
	}
}

func (s *AliasableUntypedAssignDeclarationsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableUntypedAssignDeclarationsInitial(s)
	}
}

func (s *AliasableUntypedAssignDeclarationsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableUntypedAssignDeclarationsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasableUntypedAssignDeclarationsAppendContext struct {
	Aliasable_untyped_assign_declarationsContext
}

func NewAliasableUntypedAssignDeclarationsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasableUntypedAssignDeclarationsAppendContext {
	var p = new(AliasableUntypedAssignDeclarationsAppendContext)

	InitEmptyAliasable_untyped_assign_declarationsContext(&p.Aliasable_untyped_assign_declarationsContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_untyped_assign_declarationsContext))

	return p
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) Aliasable_untyped_assign_declarations() IAliasable_untyped_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_untyped_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_untyped_assign_declarationsContext)
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) Aliasable_untyped_assign_declaration() IAliasable_untyped_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_untyped_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_untyped_assign_declarationContext)
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasableUntypedAssignDeclarationsAppend(s)
	}
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasableUntypedAssignDeclarationsAppend(s)
	}
}

func (s *AliasableUntypedAssignDeclarationsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasableUntypedAssignDeclarationsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_untyped_assign_declarations() (localctx IAliasable_untyped_assign_declarationsContext) {
	return p.aliasable_untyped_assign_declarations(0)
}

func (p *CaliburnParser) aliasable_untyped_assign_declarations(_p int) (localctx IAliasable_untyped_assign_declarationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewAliasable_untyped_assign_declarationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAliasable_untyped_assign_declarationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, CaliburnParserRULE_aliasable_untyped_assign_declarations, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewAliasableUntypedAssignDeclarationsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(362)
		p.Aliasable_untyped_assign_declaration()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAliasableUntypedAssignDeclarationsAppendContext(p, NewAliasable_untyped_assign_declarationsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_aliasable_untyped_assign_declarations)
			p.SetState(364)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(365)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(366)
				p.Aliasable_untyped_assign_declaration()
			}

		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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

// IAliasable_untyped_assign_declarationContext is an interface to support dynamic dispatch.
type IAliasable_untyped_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAliasable_untyped_assign_declarationContext differentiates from other interfaces.
	IsAliasable_untyped_assign_declarationContext()
}

type Aliasable_untyped_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasable_untyped_assign_declarationContext() *Aliasable_untyped_assign_declarationContext {
	var p = new(Aliasable_untyped_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declaration
	return p
}

func InitEmptyAliasable_untyped_assign_declarationContext(p *Aliasable_untyped_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declaration
}

func (*Aliasable_untyped_assign_declarationContext) IsAliasable_untyped_assign_declarationContext() {}

func NewAliasable_untyped_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliasable_untyped_assign_declarationContext {
	var p = new(Aliasable_untyped_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_aliasable_untyped_assign_declaration

	return p
}

func (s *Aliasable_untyped_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliasable_untyped_assign_declarationContext) CopyAll(ctx *Aliasable_untyped_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Aliasable_untyped_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliasable_untyped_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaliasedUntypedAssignDeclarationContext struct {
	Aliasable_untyped_assign_declarationContext
}

func NewUnaliasedUntypedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaliasedUntypedAssignDeclarationContext {
	var p = new(UnaliasedUntypedAssignDeclarationContext)

	InitEmptyAliasable_untyped_assign_declarationContext(&p.Aliasable_untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_untyped_assign_declarationContext))

	return p
}

func (s *UnaliasedUntypedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaliasedUntypedAssignDeclarationContext) Untyped_assign_declaration() IUntyped_assign_declarationContext {
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

func (s *UnaliasedUntypedAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUnaliasedUntypedAssignDeclaration(s)
	}
}

func (s *UnaliasedUntypedAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUnaliasedUntypedAssignDeclaration(s)
	}
}

func (s *UnaliasedUntypedAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUnaliasedUntypedAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type AliasedUntypedAssignDeclarationContext struct {
	Aliasable_untyped_assign_declarationContext
}

func NewAliasedUntypedAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AliasedUntypedAssignDeclarationContext {
	var p = new(AliasedUntypedAssignDeclarationContext)

	InitEmptyAliasable_untyped_assign_declarationContext(&p.Aliasable_untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Aliasable_untyped_assign_declarationContext))

	return p
}

func (s *AliasedUntypedAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasedUntypedAssignDeclarationContext) Untyped_assign_declaration() IUntyped_assign_declarationContext {
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

func (s *AliasedUntypedAssignDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *AliasedUntypedAssignDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AliasedUntypedAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAliasedUntypedAssignDeclaration(s)
	}
}

func (s *AliasedUntypedAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAliasedUntypedAssignDeclaration(s)
	}
}

func (s *AliasedUntypedAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAliasedUntypedAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Aliasable_untyped_assign_declaration() (localctx IAliasable_untyped_assign_declarationContext) {
	localctx = NewAliasable_untyped_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CaliburnParserRULE_aliasable_untyped_assign_declaration)
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAliasedUntypedAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.Untyped_assign_declaration()
		}
		{
			p.SetState(373)
			p.Match(CaliburnParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Identifier()
		}

	case 2:
		localctx = NewUnaliasedUntypedAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(376)
			p.Untyped_assign_declaration()
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
	var_ IIdentifierContext
}

func NewUntypedIdentifierAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedIdentifierAssignDeclarationContext {
	var p = new(UntypedIdentifierAssignDeclarationContext)

	InitEmptyUntyped_assign_declarationContext(&p.Untyped_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Untyped_assign_declarationContext))

	return p
}

func (s *UntypedIdentifierAssignDeclarationContext) GetVar_() IIdentifierContext { return s.var_ }

func (s *UntypedIdentifierAssignDeclarationContext) SetVar_(v IIdentifierContext) { s.var_ = v }

func (s *UntypedIdentifierAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedIdentifierAssignDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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

func (s *UntypedStructDestrutureAssignDeclarationContext) Aliasable_untyped_assign_declarations() IAliasable_untyped_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAliasable_untyped_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAliasable_untyped_assign_declarationsContext)
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

func (p *CaliburnParser) Untyped_assign_declaration() (localctx IUntyped_assign_declarationContext) {
	localctx = NewUntyped_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CaliburnParserRULE_untyped_assign_declaration)
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifierToken:
		localctx = NewUntypedIdentifierAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)

			var _x = p.Identifier()

			localctx.(*UntypedIdentifierAssignDeclarationContext).var_ = _x
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewUntypedStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.aliasable_untyped_assign_declarations(0)
		}
		{
			p.SetState(382)
			p.Match(CaliburnParserR_C_BRACK)
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
	p.EnterRule(localctx, 76, CaliburnParserRULE_expression_statement)
	localctx = NewExpressionStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(386)
		p.expressions(0)
	}
	{
		p.SetState(387)
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

type ExpressionsInitialContext struct {
	ExpressionsContext
}

func NewExpressionsInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionsInitialContext {
	var p = new(ExpressionsInitialContext)

	InitEmptyExpressionsContext(&p.ExpressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionsContext))

	return p
}

func (s *ExpressionsInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsInitialContext) Expression() IExpressionContext {
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

func (s *ExpressionsInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionsInitial(s)
	}
}

func (s *ExpressionsInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionsInitial(s)
	}
}

func (s *ExpressionsInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionsInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionsAppendContext struct {
	ExpressionsContext
}

func NewExpressionsAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionsAppendContext {
	var p = new(ExpressionsAppendContext)

	InitEmptyExpressionsContext(&p.ExpressionsContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionsContext))

	return p
}

func (s *ExpressionsAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsAppendContext) Expressions() IExpressionsContext {
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

func (s *ExpressionsAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *ExpressionsAppendContext) Expression() IExpressionContext {
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

func (s *ExpressionsAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionsAppend(s)
	}
}

func (s *ExpressionsAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionsAppend(s)
	}
}

func (s *ExpressionsAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionsAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expressions() (localctx IExpressionsContext) {
	return p.expressions(0)
}

func (p *CaliburnParser) expressions(_p int) (localctx IExpressionsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 78
	p.EnterRecursionRule(localctx, 78, CaliburnParserRULE_expressions, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewExpressionsInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(390)
		p.expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(397)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionsAppendContext(p, NewExpressionsContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expressions)
			p.SetState(392)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(393)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(394)
				p.expression(0)
			}

		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext())
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

// IExpressions_optionalContext is an interface to support dynamic dispatch.
type IExpressions_optionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressions_optionalContext differentiates from other interfaces.
	IsExpressions_optionalContext()
}

type Expressions_optionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressions_optionalContext() *Expressions_optionalContext {
	var p = new(Expressions_optionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expressions_optional
	return p
}

func InitEmptyExpressions_optionalContext(p *Expressions_optionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expressions_optional
}

func (*Expressions_optionalContext) IsExpressions_optionalContext() {}

func NewExpressions_optionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expressions_optionalContext {
	var p = new(Expressions_optionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expressions_optional

	return p
}

func (s *Expressions_optionalContext) GetParser() antlr.Parser { return s.parser }

func (s *Expressions_optionalContext) CopyAll(ctx *Expressions_optionalContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Expressions_optionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expressions_optionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionsOptionalContext struct {
	Expressions_optionalContext
}

func NewExpressionsOptionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionsOptionalContext {
	var p = new(ExpressionsOptionalContext)

	InitEmptyExpressions_optionalContext(&p.Expressions_optionalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Expressions_optionalContext))

	return p
}

func (s *ExpressionsOptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsOptionalContext) Expressions() IExpressionsContext {
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

func (s *ExpressionsOptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionsOptional(s)
	}
}

func (s *ExpressionsOptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionsOptional(s)
	}
}

func (s *ExpressionsOptionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionsOptional(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionsOptionalNoneContext struct {
	Expressions_optionalContext
}

func NewExpressionsOptionalNoneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionsOptionalNoneContext {
	var p = new(ExpressionsOptionalNoneContext)

	InitEmptyExpressions_optionalContext(&p.Expressions_optionalContext)
	p.parser = parser
	p.CopyAll(ctx.(*Expressions_optionalContext))

	return p
}

func (s *ExpressionsOptionalNoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsOptionalNoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressionsOptionalNone(s)
	}
}

func (s *ExpressionsOptionalNoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressionsOptionalNone(s)
	}
}

func (s *ExpressionsOptionalNoneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressionsOptionalNone(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expressions_optional() (localctx IExpressions_optionalContext) {
	localctx = NewExpressions_optionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CaliburnParserRULE_expressions_optional)
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserTUPLE, CaliburnParserL_PAREN, CaliburnParserOP_ADD, CaliburnParserOP_SUB, CaliburnParserOP_NOT, CaliburnParserIdentifierToken, CaliburnParserLiteralToken:
		localctx = NewExpressionsOptionalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			p.expressions(0)
		}

	case CaliburnParserR_PAREN, CaliburnParserR_C_BRACK:
		localctx = NewExpressionsOptionalNoneContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

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
}

func NewBracketedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketedExpressionContext {
	var p = new(BracketedExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BracketedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketedExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
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

func (s *BracketedExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
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
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

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
}

func NewSliceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceExpressionContext {
	var p = new(SliceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SliceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *SliceExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *SliceExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *SliceExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
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

type IndexExpressionContext struct {
	ExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *IndexExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *IndexExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
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
	op antlr.Token
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
}

func NewAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessExpressionContext {
	var p = new(AccessExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *AccessExpressionContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserPERIOD, 0)
}

func (s *AccessExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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

type SliceEndExpressionContext struct {
	ExpressionContext
}

func NewSliceEndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceEndExpressionContext {
	var p = new(SliceEndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SliceEndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceEndExpressionContext) AllExpression() []IExpressionContext {
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

func (s *SliceEndExpressionContext) Expression(i int) IExpressionContext {
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

func (s *SliceEndExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *SliceEndExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *SliceEndExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *SliceEndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterSliceEndExpression(s)
	}
}

func (s *SliceEndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitSliceEndExpression(s)
	}
}

func (s *SliceEndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitSliceEndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	ExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *FunctionExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *FunctionExpressionContext) Parameters() IParametersContext {
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

func (s *FunctionExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
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
}

func NewStructExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExpressionContext {
	var p = new(StructExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StructExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *StructExpressionContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *StructExpressionContext) Struct_values() IStruct_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_valuesContext)
}

func (s *StructExpressionContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
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
	op antlr.Token
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

func (s *BinaryExpressionContext) OP_EQU() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_EQU, 0)
}

func (s *BinaryExpressionContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_NEQ, 0)
}

func (s *BinaryExpressionContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LTE, 0)
}

func (s *BinaryExpressionContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GTE, 0)
}

func (s *BinaryExpressionContext) OP_LST() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LST, 0)
}

func (s *BinaryExpressionContext) OP_GRT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GRT, 0)
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

type SliceStartExpressionContext struct {
	ExpressionContext
}

func NewSliceStartExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceStartExpressionContext {
	var p = new(SliceStartExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *SliceStartExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceStartExpressionContext) AllExpression() []IExpressionContext {
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

func (s *SliceStartExpressionContext) Expression(i int) IExpressionContext {
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

func (s *SliceStartExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *SliceStartExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *SliceStartExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *SliceStartExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterSliceStartExpression(s)
	}
}

func (s *SliceStartExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitSliceStartExpression(s)
	}
}

func (s *SliceStartExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitSliceStartExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionExpressionNoReturnTypeContext struct {
	ExpressionContext
}

func NewFunctionExpressionNoReturnTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionNoReturnTypeContext {
	var p = new(FunctionExpressionNoReturnTypeContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionNoReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionNoReturnTypeContext) Function_type() IFunction_typeContext {
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

func (s *FunctionExpressionNoReturnTypeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *FunctionExpressionNoReturnTypeContext) Parameters() IParametersContext {
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

func (s *FunctionExpressionNoReturnTypeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *FunctionExpressionNoReturnTypeContext) Block() IBlockContext {
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

func (s *FunctionExpressionNoReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunctionExpressionNoReturnType(s)
	}
}

func (s *FunctionExpressionNoReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunctionExpressionNoReturnType(s)
	}
}

func (s *FunctionExpressionNoReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunctionExpressionNoReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

type CollectionExpressionContext struct {
	ExpressionContext
}

func NewCollectionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CollectionExpressionContext {
	var p = new(CollectionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CollectionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionExpressionContext) Type_expression() IType_expressionContext {
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

func (s *CollectionExpressionContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
}

func (s *CollectionExpressionContext) Collection_values() ICollection_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_valuesContext)
}

func (s *CollectionExpressionContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
}

func (s *CollectionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCollectionExpression(s)
	}
}

func (s *CollectionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCollectionExpression(s)
	}
}

func (s *CollectionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCollectionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpressionContext struct {
	ExpressionContext
}

func NewCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpressionContext {
	var p = new(CallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
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

func (s *CallExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *CallExpressionContext) Expressions_optional() IExpressions_optionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressions_optionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressions_optionalContext)
}

func (s *CallExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
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
	_startState := 82
	p.EnterRecursionRule(localctx, 82, CaliburnParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(405)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(406)
			p.expression(0)
		}
		{
			p.SetState(407)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(409)

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
			p.SetState(410)
			p.expression(20)
		}

	case 3:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(411)
			p.Identifier()
		}

	case 4:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(412)
			p.Literal_atom()
		}

	case 5:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(413)
			p.Function_type()
		}
		{
			p.SetState(414)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(415)
			p.Parameters()
		}
		{
			p.SetState(416)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.type_expression(0)
		}
		{
			p.SetState(418)
			p.Block()
		}

	case 6:
		localctx = NewFunctionExpressionNoReturnTypeContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.Function_type()
		}
		{
			p.SetState(421)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Parameters()
		}
		{
			p.SetState(423)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(424)
			p.Block()
		}

	case 7:
		localctx = NewStructExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(426)
			p.Struct_type()
		}
		{
			p.SetState(427)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(428)
			p.Struct_values()
		}
		{
			p.SetState(429)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewCollectionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(431)
			p.type_expression(0)
		}
		{
			p.SetState(432)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(433)
			p.collection_values(0)
		}
		{
			p.SetState(434)
			p.Match(CaliburnParserR_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(494)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(438)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(439)

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
					p.SetState(440)
					p.expression(22)
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(441)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(442)

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
					p.SetState(443)
					p.expression(20)
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(444)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(445)

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
					p.SetState(446)
					p.expression(19)
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(447)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(448)

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
					p.SetState(449)
					p.expression(18)
				}

			case 5:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(450)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(451)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35465847065542656) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(452)
					p.expression(17)
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(453)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(454)

					var _m = p.Match(CaliburnParserOP_AND)

					localctx.(*BinaryExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(455)
					p.expression(16)
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(456)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(457)

					var _m = p.Match(CaliburnParserOP_XOR)

					localctx.(*BinaryExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(458)
					p.expression(15)
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(459)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(460)

					var _m = p.Match(CaliburnParserOP_OR)

					localctx.(*BinaryExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(461)
					p.expression(14)
				}

			case 9:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(462)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(463)
					p.Match(CaliburnParserL_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(464)
					p.Expressions_optional()
				}
				{
					p.SetState(465)
					p.Match(CaliburnParserR_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(467)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(468)
					p.Match(CaliburnParserPERIOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(469)
					p.Identifier()
				}

			case 11:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(470)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(471)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(472)
					p.expression(0)
				}
				{
					p.SetState(473)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewSliceStartExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(475)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(476)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(477)
					p.expression(0)
				}
				{
					p.SetState(478)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(479)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 13:
				localctx = NewSliceEndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(481)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(482)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(483)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(484)
					p.expression(0)
				}
				{
					p.SetState(485)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 14:
				localctx = NewSliceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(487)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(488)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(489)
					p.expression(0)
				}
				{
					p.SetState(490)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(491)
					p.expression(0)
				}
				{
					p.SetState(492)
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
		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
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

// IStruct_valuesContext is an interface to support dynamic dispatch.
type IStruct_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStruct_valuesContext differentiates from other interfaces.
	IsStruct_valuesContext()
}

type Struct_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_valuesContext() *Struct_valuesContext {
	var p = new(Struct_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_values
	return p
}

func InitEmptyStruct_valuesContext(p *Struct_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_values
}

func (*Struct_valuesContext) IsStruct_valuesContext() {}

func NewStruct_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_valuesContext {
	var p = new(Struct_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_struct_values

	return p
}

func (s *Struct_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_valuesContext) CopyAll(ctx *Struct_valuesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Struct_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructValuesUnamedContext struct {
	Struct_valuesContext
}

func NewStructValuesUnamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructValuesUnamedContext {
	var p = new(StructValuesUnamedContext)

	InitEmptyStruct_valuesContext(&p.Struct_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_valuesContext))

	return p
}

func (s *StructValuesUnamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructValuesUnamedContext) Expressions_optional() IExpressions_optionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressions_optionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressions_optionalContext)
}

func (s *StructValuesUnamedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructValuesUnamed(s)
	}
}

func (s *StructValuesUnamedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructValuesUnamed(s)
	}
}

func (s *StructValuesUnamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructValuesUnamed(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructValuesNamedContext struct {
	Struct_valuesContext
}

func NewStructValuesNamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructValuesNamedContext {
	var p = new(StructValuesNamedContext)

	InitEmptyStruct_valuesContext(&p.Struct_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_valuesContext))

	return p
}

func (s *StructValuesNamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructValuesNamedContext) Named_struct_values() INamed_struct_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamed_struct_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamed_struct_valuesContext)
}

func (s *StructValuesNamedContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *StructValuesNamedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructValuesNamed(s)
	}
}

func (s *StructValuesNamedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructValuesNamed(s)
	}
}

func (s *StructValuesNamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructValuesNamed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Struct_values() (localctx IStruct_valuesContext) {
	localctx = NewStruct_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, CaliburnParserRULE_struct_values)
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructValuesNamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.named_struct_values(0)
		}

	case 2:
		localctx = NewStructValuesNamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(500)
			p.named_struct_values(0)
		}
		{
			p.SetState(501)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewStructValuesUnamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(503)
			p.Expressions_optional()
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

// INamed_struct_valuesContext is an interface to support dynamic dispatch.
type INamed_struct_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNamed_struct_valuesContext differentiates from other interfaces.
	IsNamed_struct_valuesContext()
}

type Named_struct_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamed_struct_valuesContext() *Named_struct_valuesContext {
	var p = new(Named_struct_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_named_struct_values
	return p
}

func InitEmptyNamed_struct_valuesContext(p *Named_struct_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_named_struct_values
}

func (*Named_struct_valuesContext) IsNamed_struct_valuesContext() {}

func NewNamed_struct_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Named_struct_valuesContext {
	var p = new(Named_struct_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_named_struct_values

	return p
}

func (s *Named_struct_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Named_struct_valuesContext) CopyAll(ctx *Named_struct_valuesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Named_struct_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Named_struct_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedStructValuesAppendContext struct {
	Named_struct_valuesContext
}

func NewNamedStructValuesAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedStructValuesAppendContext {
	var p = new(NamedStructValuesAppendContext)

	InitEmptyNamed_struct_valuesContext(&p.Named_struct_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Named_struct_valuesContext))

	return p
}

func (s *NamedStructValuesAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedStructValuesAppendContext) Named_struct_values() INamed_struct_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamed_struct_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamed_struct_valuesContext)
}

func (s *NamedStructValuesAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *NamedStructValuesAppendContext) Named_struct_value() INamed_struct_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamed_struct_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamed_struct_valueContext)
}

func (s *NamedStructValuesAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterNamedStructValuesAppend(s)
	}
}

func (s *NamedStructValuesAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitNamedStructValuesAppend(s)
	}
}

func (s *NamedStructValuesAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitNamedStructValuesAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

type NamedStructValuesInitialContext struct {
	Named_struct_valuesContext
}

func NewNamedStructValuesInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedStructValuesInitialContext {
	var p = new(NamedStructValuesInitialContext)

	InitEmptyNamed_struct_valuesContext(&p.Named_struct_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Named_struct_valuesContext))

	return p
}

func (s *NamedStructValuesInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedStructValuesInitialContext) Named_struct_value() INamed_struct_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamed_struct_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamed_struct_valueContext)
}

func (s *NamedStructValuesInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterNamedStructValuesInitial(s)
	}
}

func (s *NamedStructValuesInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitNamedStructValuesInitial(s)
	}
}

func (s *NamedStructValuesInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitNamedStructValuesInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Named_struct_values() (localctx INamed_struct_valuesContext) {
	return p.named_struct_values(0)
}

func (p *CaliburnParser) named_struct_values(_p int) (localctx INamed_struct_valuesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewNamed_struct_valuesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INamed_struct_valuesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 86
	p.EnterRecursionRule(localctx, 86, CaliburnParserRULE_named_struct_values, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewNamedStructValuesInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(507)
		p.Named_struct_value()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewNamedStructValuesAppendContext(p, NewNamed_struct_valuesContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_named_struct_values)
			p.SetState(509)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(510)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(511)
				p.Named_struct_value()
			}

		}
		p.SetState(516)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext())
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

// INamed_struct_valueContext is an interface to support dynamic dispatch.
type INamed_struct_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsNamed_struct_valueContext differentiates from other interfaces.
	IsNamed_struct_valueContext()
}

type Named_struct_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamed_struct_valueContext() *Named_struct_valueContext {
	var p = new(Named_struct_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_named_struct_value
	return p
}

func InitEmptyNamed_struct_valueContext(p *Named_struct_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_named_struct_value
}

func (*Named_struct_valueContext) IsNamed_struct_valueContext() {}

func NewNamed_struct_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Named_struct_valueContext {
	var p = new(Named_struct_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_named_struct_value

	return p
}

func (s *Named_struct_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Named_struct_valueContext) CopyAll(ctx *Named_struct_valueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Named_struct_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Named_struct_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NamedStructValueContext struct {
	Named_struct_valueContext
}

func NewNamedStructValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NamedStructValueContext {
	var p = new(NamedStructValueContext)

	InitEmptyNamed_struct_valueContext(&p.Named_struct_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Named_struct_valueContext))

	return p
}

func (s *NamedStructValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedStructValueContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *NamedStructValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *NamedStructValueContext) Expression() IExpressionContext {
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

func (s *NamedStructValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterNamedStructValue(s)
	}
}

func (s *NamedStructValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitNamedStructValue(s)
	}
}

func (s *NamedStructValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitNamedStructValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Named_struct_value() (localctx INamed_struct_valueContext) {
	localctx = NewNamed_struct_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, CaliburnParserRULE_named_struct_value)
	localctx = NewNamedStructValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Identifier()
	}
	{
		p.SetState(518)
		p.Match(CaliburnParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(519)
		p.expression(0)
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

// ICollection_valuesContext is an interface to support dynamic dispatch.
type ICollection_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCollection_valuesContext differentiates from other interfaces.
	IsCollection_valuesContext()
}

type Collection_valuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_valuesContext() *Collection_valuesContext {
	var p = new(Collection_valuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_collection_values
	return p
}

func InitEmptyCollection_valuesContext(p *Collection_valuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_collection_values
}

func (*Collection_valuesContext) IsCollection_valuesContext() {}

func NewCollection_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_valuesContext {
	var p = new(Collection_valuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_collection_values

	return p
}

func (s *Collection_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_valuesContext) CopyAll(ctx *Collection_valuesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Collection_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CollectionValuesInitialContext struct {
	Collection_valuesContext
}

func NewCollectionValuesInitialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CollectionValuesInitialContext {
	var p = new(CollectionValuesInitialContext)

	InitEmptyCollection_valuesContext(&p.Collection_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Collection_valuesContext))

	return p
}

func (s *CollectionValuesInitialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionValuesInitialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCollectionValuesInitial(s)
	}
}

func (s *CollectionValuesInitialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCollectionValuesInitial(s)
	}
}

func (s *CollectionValuesInitialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCollectionValuesInitial(s)

	default:
		return t.VisitChildren(s)
	}
}

type CollectionValuesAppendContext struct {
	Collection_valuesContext
}

func NewCollectionValuesAppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CollectionValuesAppendContext {
	var p = new(CollectionValuesAppendContext)

	InitEmptyCollection_valuesContext(&p.Collection_valuesContext)
	p.parser = parser
	p.CopyAll(ctx.(*Collection_valuesContext))

	return p
}

func (s *CollectionValuesAppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionValuesAppendContext) Collection_values() ICollection_valuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_valuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_valuesContext)
}

func (s *CollectionValuesAppendContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, 0)
}

func (s *CollectionValuesAppendContext) Collection_value() ICollection_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollection_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollection_valueContext)
}

func (s *CollectionValuesAppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCollectionValuesAppend(s)
	}
}

func (s *CollectionValuesAppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCollectionValuesAppend(s)
	}
}

func (s *CollectionValuesAppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCollectionValuesAppend(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Collection_values() (localctx ICollection_valuesContext) {
	return p.collection_values(0)
}

func (p *CaliburnParser) collection_values(_p int) (localctx ICollection_valuesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCollection_valuesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICollection_valuesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 90
	p.EnterRecursionRule(localctx, 90, CaliburnParserRULE_collection_values, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewCollectionValuesInitialContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCollectionValuesAppendContext(p, NewCollection_valuesContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_collection_values)
			p.SetState(522)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(523)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(524)
				p.Collection_value()
			}

		}
		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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

// ICollection_valueContext is an interface to support dynamic dispatch.
type ICollection_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCollection_valueContext differentiates from other interfaces.
	IsCollection_valueContext()
}

type Collection_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollection_valueContext() *Collection_valueContext {
	var p = new(Collection_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_collection_value
	return p
}

func InitEmptyCollection_valueContext(p *Collection_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_collection_value
}

func (*Collection_valueContext) IsCollection_valueContext() {}

func NewCollection_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collection_valueContext {
	var p = new(Collection_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_collection_value

	return p
}

func (s *Collection_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Collection_valueContext) CopyAll(ctx *Collection_valueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Collection_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collection_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CollectionValueContext struct {
	Collection_valueContext
}

func NewCollectionValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CollectionValueContext {
	var p = new(CollectionValueContext)

	InitEmptyCollection_valueContext(&p.Collection_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Collection_valueContext))

	return p
}

func (s *CollectionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionValueContext) AllExpression() []IExpressionContext {
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

func (s *CollectionValueContext) Expression(i int) IExpressionContext {
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

func (s *CollectionValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, 0)
}

func (s *CollectionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterCollectionValue(s)
	}
}

func (s *CollectionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitCollectionValue(s)
	}
}

func (s *CollectionValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitCollectionValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Collection_value() (localctx ICollection_valueContext) {
	localctx = NewCollection_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, CaliburnParserRULE_collection_value)
	localctx = NewCollectionValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.expression(0)
	}
	{
		p.SetState(531)
		p.Match(CaliburnParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(532)
		p.expression(0)
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
	p.EnterRule(localctx, 94, CaliburnParserRULE_function_type)
	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC:
		localctx = NewFunctionTypeFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Match(CaliburnParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifierToken:
		localctx = NewFunctionTypeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)
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

type StructTypeTupleContext struct {
	Struct_typeContext
}

func NewStructTypeTupleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeTupleContext {
	var p = new(StructTypeTupleContext)

	InitEmptyStruct_typeContext(&p.Struct_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Struct_typeContext))

	return p
}

func (s *StructTypeTupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeTupleContext) TUPLE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTUPLE, 0)
}

func (s *StructTypeTupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStructTypeTuple(s)
	}
}

func (s *StructTypeTupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStructTypeTuple(s)
	}
}

func (s *StructTypeTupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStructTypeTuple(s)

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
	p.EnterRule(localctx, 96, CaliburnParserRULE_struct_type)
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserSTRUCT:
		localctx = NewStructTypeStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(538)
			p.Match(CaliburnParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserTUPLE:
		localctx = NewStructTypeTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(539)
			p.Match(CaliburnParserTUPLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifierToken:
		localctx = NewStructTypeExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(540)
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

func (s *AccessTypeExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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

func (s *IdentifierTypeExpressionContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
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
	_startState := 98
	p.EnterRecursionRule(localctx, 98, CaliburnParserRULE_type_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewIdentifierTypeExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(544)
		p.Identifier()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(551)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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
			p.SetState(546)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(547)
				p.Match(CaliburnParserPERIOD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(548)
				p.Identifier()
			}

		}
		p.SetState(553)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext())
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

func (s *TypedLiteralContext) Type_expression() IType_expressionContext {
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

func (s *TypedLiteralContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
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

func (s *UntypedLiteralContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
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
	p.EnterRule(localctx, 100, CaliburnParserRULE_literal_atom)
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserLiteralToken:
		localctx = NewUntypedLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(554)
			p.Literal()
		}

	case CaliburnParserIdentifierToken:
		localctx = NewTypedLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(555)
			p.type_expression(0)
		}
		{
			p.SetState(556)
			p.Literal()
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

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// Getter signatures
	LiteralToken() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    antlr.Token
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) GetVal() antlr.Token { return s.val }

func (s *LiteralContext) SetVal(v antlr.Token) { s.val = v }

func (s *LiteralContext) LiteralToken() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteralToken, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, CaliburnParserRULE_literal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(560)

		var _m = p.Match(CaliburnParserLiteralToken)

		localctx.(*LiteralContext).val = _m
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVal returns the val token.
	GetVal() antlr.Token

	// SetVal sets the val token.
	SetVal(antlr.Token)

	// Getter signatures
	IdentifierToken() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	val    antlr.Token
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) GetVal() antlr.Token { return s.val }

func (s *IdentifierContext) SetVal(v antlr.Token) { s.val = v }

func (s *IdentifierContext) IdentifierToken() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifierToken, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, CaliburnParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(562)

		var _m = p.Match(CaliburnParserIdentifierToken)

		localctx.(*IdentifierContext).val = _m
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

	case 10:
		var t *Inline_statementsContext = nil
		if localctx != nil {
			t = localctx.(*Inline_statementsContext)
		}
		return p.Inline_statements_Sempred(t, predIndex)

	case 22:
		var t *Option_case_blocksContext = nil
		if localctx != nil {
			t = localctx.(*Option_case_blocksContext)
		}
		return p.Option_case_blocks_Sempred(t, predIndex)

	case 26:
		var t *Assign_expressionsContext = nil
		if localctx != nil {
			t = localctx.(*Assign_expressionsContext)
		}
		return p.Assign_expressions_Sempred(t, predIndex)

	case 27:
		var t *Aliasable_assign_expressionsContext = nil
		if localctx != nil {
			t = localctx.(*Aliasable_assign_expressionsContext)
		}
		return p.Aliasable_assign_expressions_Sempred(t, predIndex)

	case 30:
		var t *Assign_declarationsContext = nil
		if localctx != nil {
			t = localctx.(*Assign_declarationsContext)
		}
		return p.Assign_declarations_Sempred(t, predIndex)

	case 31:
		var t *Aliasable_assign_declarationsContext = nil
		if localctx != nil {
			t = localctx.(*Aliasable_assign_declarationsContext)
		}
		return p.Aliasable_assign_declarations_Sempred(t, predIndex)

	case 35:
		var t *Aliasable_untyped_assign_declarationsContext = nil
		if localctx != nil {
			t = localctx.(*Aliasable_untyped_assign_declarationsContext)
		}
		return p.Aliasable_untyped_assign_declarations_Sempred(t, predIndex)

	case 39:
		var t *ExpressionsContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionsContext)
		}
		return p.Expressions_Sempred(t, predIndex)

	case 41:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 43:
		var t *Named_struct_valuesContext = nil
		if localctx != nil {
			t = localctx.(*Named_struct_valuesContext)
		}
		return p.Named_struct_values_Sempred(t, predIndex)

	case 45:
		var t *Collection_valuesContext = nil
		if localctx != nil {
			t = localctx.(*Collection_valuesContext)
		}
		return p.Collection_values_Sempred(t, predIndex)

	case 49:
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

func (p *CaliburnParser) Option_case_blocks_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Assign_expressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Aliasable_assign_expressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Assign_declarations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Aliasable_assign_declarations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Aliasable_untyped_assign_declarations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Expressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 10:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 23:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 24:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Named_struct_values_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 25:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Collection_values_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 26:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Type_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 27:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

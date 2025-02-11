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
		"", "'.'", "','", "':'", "'?'", "';'", "'='", "'if'", "'else'", "'for'",
		"'switch'", "'case'", "'default'", "'fallthrough'", "'break'", "'continue'",
		"'return'", "'using'", "'as'", "'in'", "'null'", "'var'", "'const'",
		"'type'", "'func'", "'struct'", "'tuple'", "'class'", "'interface'",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'+'", "'-'", "'!'", "'*'",
		"'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", "'&'", "'<<'", "'>>'", "'=='",
		"'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "PERIOD", "COMMA", "COLON", "QUESTION", "Terminator", "ASSIGN",
		"IF", "ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", "BREAK",
		"CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", "TYPE",
		"FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN",
		"L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD", "OP_SUB",
		"OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", "OP_OR",
		"OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ", "OP_GTE",
		"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier",
		"Literal",
	}
	staticData.RuleNames = []string{
		"module", "definition", "scoped_block", "statement", "jump_statement",
		"return_statement", "break_statement", "continue_statement", "control_statement",
		"if_statement", "for_statement", "switch_statement", "inline_statement",
		"function_declaration_statement", "assign_statement", "assign_declarations",
		"assign_declaration", "declared_assign_declarations", "declared_assign_declaration",
		"undeclared_assign_declarations", "undeclared_assign_declaration", "operator_assign_statement",
		"expression_statement", "expressions", "expression", "function_expression",
		"struct_expression", "tuple_expression", "type_expression", "complex_type_expression",
		"identifiers", "literal_atom", "untyped_literal_atom", "typed_literal_atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 57, 406, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 5, 0, 70, 8, 0, 10, 0, 12, 0, 73,
		9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 79, 8, 2, 10, 2, 12, 2, 82, 9, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 90, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4,
		95, 8, 4, 1, 5, 1, 5, 3, 5, 99, 8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 113, 8, 8, 1, 9, 1, 9, 5,
		9, 117, 8, 9, 10, 9, 12, 9, 120, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3,
		9, 127, 8, 9, 3, 9, 129, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		5, 11, 137, 8, 11, 10, 11, 12, 11, 140, 9, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 147, 8, 11, 1, 11, 1, 11, 5, 11, 151, 8, 11, 10, 11,
		12, 11, 154, 9, 11, 1, 11, 1, 11, 3, 11, 158, 8, 11, 1, 11, 1, 11, 1, 12,
		1, 12, 3, 12, 164, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 170, 8, 13,
		1, 13, 1, 13, 3, 13, 174, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 3, 14, 184, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 189, 8,
		15, 10, 15, 12, 15, 192, 9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 204, 8, 16, 1, 17, 1, 17, 1, 17, 5,
		17, 209, 8, 17, 10, 17, 12, 17, 212, 9, 17, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 3, 18, 219, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 224, 8, 19, 10, 19,
		12, 19, 227, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 238, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 5, 23, 252, 8, 23, 10, 23, 12,
		23, 255, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 274,
		8, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 303, 8, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 3, 24, 317, 8, 24, 1, 24, 1, 24, 3, 24, 321, 8, 24, 1, 24, 5, 24,
		324, 8, 24, 10, 24, 12, 24, 327, 9, 24, 1, 25, 1, 25, 1, 25, 3, 25, 332,
		8, 25, 1, 25, 1, 25, 3, 25, 336, 8, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 349, 8, 26, 10, 26,
		12, 26, 352, 9, 26, 1, 26, 3, 26, 355, 8, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 4, 27, 364, 8, 27, 11, 27, 12, 27, 365, 3, 27, 368,
		8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 376, 8, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 384, 8, 29, 10, 29, 12, 29,
		387, 9, 29, 1, 30, 1, 30, 1, 30, 5, 30, 392, 8, 30, 10, 30, 12, 30, 395,
		9, 30, 1, 31, 1, 31, 3, 31, 399, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		33, 1, 33, 0, 2, 48, 58, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 0, 7, 2, 0, 35, 36, 38, 47, 1, 0, 35, 37, 1, 0, 41, 42,
		1, 0, 38, 40, 1, 0, 35, 36, 1, 0, 46, 47, 1, 0, 48, 53, 437, 0, 71, 1,
		0, 0, 0, 2, 74, 1, 0, 0, 0, 4, 76, 1, 0, 0, 0, 6, 89, 1, 0, 0, 0, 8, 94,
		1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12, 102, 1, 0, 0, 0, 14, 105, 1, 0, 0,
		0, 16, 112, 1, 0, 0, 0, 18, 114, 1, 0, 0, 0, 20, 130, 1, 0, 0, 0, 22, 134,
		1, 0, 0, 0, 24, 163, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0, 28, 183, 1, 0, 0,
		0, 30, 185, 1, 0, 0, 0, 32, 203, 1, 0, 0, 0, 34, 205, 1, 0, 0, 0, 36, 218,
		1, 0, 0, 0, 38, 220, 1, 0, 0, 0, 40, 237, 1, 0, 0, 0, 42, 239, 1, 0, 0,
		0, 44, 245, 1, 0, 0, 0, 46, 248, 1, 0, 0, 0, 48, 273, 1, 0, 0, 0, 50, 328,
		1, 0, 0, 0, 52, 339, 1, 0, 0, 0, 54, 358, 1, 0, 0, 0, 56, 375, 1, 0, 0,
		0, 58, 377, 1, 0, 0, 0, 60, 388, 1, 0, 0, 0, 62, 398, 1, 0, 0, 0, 64, 400,
		1, 0, 0, 0, 66, 402, 1, 0, 0, 0, 68, 70, 3, 2, 1, 0, 69, 68, 1, 0, 0, 0,
		70, 73, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 1, 1, 0,
		0, 0, 73, 71, 1, 0, 0, 0, 74, 75, 3, 26, 13, 0, 75, 3, 1, 0, 0, 0, 76,
		80, 5, 33, 0, 0, 77, 79, 3, 6, 3, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1, 0,
		0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 83, 1, 0, 0, 0, 82, 80,
		1, 0, 0, 0, 83, 84, 5, 34, 0, 0, 84, 5, 1, 0, 0, 0, 85, 90, 3, 28, 14,
		0, 86, 90, 3, 44, 22, 0, 87, 90, 3, 16, 8, 0, 88, 90, 3, 8, 4, 0, 89, 85,
		1, 0, 0, 0, 89, 86, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 88, 1, 0, 0, 0,
		90, 7, 1, 0, 0, 0, 91, 95, 3, 10, 5, 0, 92, 95, 3, 12, 6, 0, 93, 95, 3,
		14, 7, 0, 94, 91, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95,
		9, 1, 0, 0, 0, 96, 98, 5, 16, 0, 0, 97, 99, 3, 46, 23, 0, 98, 97, 1, 0,
		0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 5, 0, 0, 101,
		11, 1, 0, 0, 0, 102, 103, 5, 14, 0, 0, 103, 104, 5, 5, 0, 0, 104, 13, 1,
		0, 0, 0, 105, 106, 5, 15, 0, 0, 106, 107, 5, 5, 0, 0, 107, 15, 1, 0, 0,
		0, 108, 113, 3, 18, 9, 0, 109, 113, 3, 20, 10, 0, 110, 113, 3, 22, 11,
		0, 111, 113, 3, 26, 13, 0, 112, 108, 1, 0, 0, 0, 112, 109, 1, 0, 0, 0,
		112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113, 17, 1, 0, 0, 0, 114, 118,
		5, 7, 0, 0, 115, 117, 3, 24, 12, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1,
		0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 121, 1, 0, 0,
		0, 120, 118, 1, 0, 0, 0, 121, 122, 3, 48, 24, 0, 122, 128, 3, 4, 2, 0,
		123, 126, 5, 8, 0, 0, 124, 127, 3, 4, 2, 0, 125, 127, 3, 18, 9, 0, 126,
		124, 1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 128, 123,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 19, 1, 0, 0, 0, 130, 131, 5, 9,
		0, 0, 131, 132, 3, 48, 24, 0, 132, 133, 3, 4, 2, 0, 133, 21, 1, 0, 0, 0,
		134, 138, 5, 10, 0, 0, 135, 137, 3, 24, 12, 0, 136, 135, 1, 0, 0, 0, 137,
		140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139, 141,
		1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 3, 48, 24, 0, 142, 152, 5,
		33, 0, 0, 143, 146, 5, 11, 0, 0, 144, 147, 3, 56, 28, 0, 145, 147, 3, 48,
		24, 0, 146, 144, 1, 0, 0, 0, 146, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 149, 3, 4, 2, 0, 149, 151, 1, 0, 0, 0, 150, 143, 1, 0, 0, 0, 151,
		154, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 157,
		1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 155, 156, 5, 12, 0, 0, 156, 158, 3, 4,
		2, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0,
		159, 160, 5, 34, 0, 0, 160, 23, 1, 0, 0, 0, 161, 164, 3, 28, 14, 0, 162,
		164, 3, 44, 22, 0, 163, 161, 1, 0, 0, 0, 163, 162, 1, 0, 0, 0, 164, 25,
		1, 0, 0, 0, 165, 166, 3, 56, 28, 0, 166, 167, 5, 56, 0, 0, 167, 169, 5,
		29, 0, 0, 168, 170, 3, 30, 15, 0, 169, 168, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 171, 1, 0, 0, 0, 171, 173, 5, 30, 0, 0, 172, 174, 3, 56, 28,
		0, 173, 172, 1, 0, 0, 0, 173, 174, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175,
		176, 3, 4, 2, 0, 176, 27, 1, 0, 0, 0, 177, 178, 3, 30, 15, 0, 178, 179,
		5, 6, 0, 0, 179, 180, 3, 46, 23, 0, 180, 181, 5, 5, 0, 0, 181, 184, 1,
		0, 0, 0, 182, 184, 3, 42, 21, 0, 183, 177, 1, 0, 0, 0, 183, 182, 1, 0,
		0, 0, 184, 29, 1, 0, 0, 0, 185, 190, 3, 32, 16, 0, 186, 187, 5, 2, 0, 0,
		187, 189, 3, 28, 14, 0, 188, 186, 1, 0, 0, 0, 189, 192, 1, 0, 0, 0, 190,
		188, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 31, 1, 0, 0, 0, 192, 190, 1,
		0, 0, 0, 193, 204, 3, 36, 18, 0, 194, 204, 3, 40, 20, 0, 195, 196, 5, 33,
		0, 0, 196, 197, 3, 30, 15, 0, 197, 198, 5, 34, 0, 0, 198, 204, 1, 0, 0,
		0, 199, 200, 5, 31, 0, 0, 200, 201, 3, 30, 15, 0, 201, 202, 5, 31, 0, 0,
		202, 204, 1, 0, 0, 0, 203, 193, 1, 0, 0, 0, 203, 194, 1, 0, 0, 0, 203,
		195, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 204, 33, 1, 0, 0, 0, 205, 210, 3,
		36, 18, 0, 206, 207, 5, 2, 0, 0, 207, 209, 3, 36, 18, 0, 208, 206, 1, 0,
		0, 0, 209, 212, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 211, 1, 0, 0, 0,
		211, 35, 1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 213, 214, 3, 56, 28, 0, 214,
		215, 3, 40, 20, 0, 215, 219, 1, 0, 0, 0, 216, 217, 5, 21, 0, 0, 217, 219,
		3, 40, 20, 0, 218, 213, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 37, 1, 0,
		0, 0, 220, 225, 3, 40, 20, 0, 221, 222, 5, 2, 0, 0, 222, 224, 3, 40, 20,
		0, 223, 221, 1, 0, 0, 0, 224, 227, 1, 0, 0, 0, 225, 223, 1, 0, 0, 0, 225,
		226, 1, 0, 0, 0, 226, 39, 1, 0, 0, 0, 227, 225, 1, 0, 0, 0, 228, 238, 3,
		56, 28, 0, 229, 230, 5, 33, 0, 0, 230, 231, 3, 38, 19, 0, 231, 232, 5,
		34, 0, 0, 232, 238, 1, 0, 0, 0, 233, 234, 5, 29, 0, 0, 234, 235, 3, 38,
		19, 0, 235, 236, 5, 30, 0, 0, 236, 238, 1, 0, 0, 0, 237, 228, 1, 0, 0,
		0, 237, 229, 1, 0, 0, 0, 237, 233, 1, 0, 0, 0, 238, 41, 1, 0, 0, 0, 239,
		240, 3, 48, 24, 0, 240, 241, 7, 0, 0, 0, 241, 242, 5, 6, 0, 0, 242, 243,
		3, 46, 23, 0, 243, 244, 5, 5, 0, 0, 244, 43, 1, 0, 0, 0, 245, 246, 3, 46,
		23, 0, 246, 247, 5, 5, 0, 0, 247, 45, 1, 0, 0, 0, 248, 253, 3, 48, 24,
		0, 249, 250, 5, 2, 0, 0, 250, 252, 3, 48, 24, 0, 251, 249, 1, 0, 0, 0,
		252, 255, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254,
		47, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 256, 257, 6, 24, -1, 0, 257, 258,
		5, 29, 0, 0, 258, 259, 3, 56, 28, 0, 259, 260, 5, 30, 0, 0, 260, 261, 3,
		48, 24, 20, 261, 274, 1, 0, 0, 0, 262, 263, 5, 29, 0, 0, 263, 264, 3, 48,
		24, 0, 264, 265, 5, 30, 0, 0, 265, 274, 1, 0, 0, 0, 266, 267, 7, 1, 0,
		0, 267, 274, 3, 48, 24, 17, 268, 274, 5, 56, 0, 0, 269, 274, 3, 62, 31,
		0, 270, 274, 3, 50, 25, 0, 271, 274, 3, 52, 26, 0, 272, 274, 3, 54, 27,
		0, 273, 256, 1, 0, 0, 0, 273, 262, 1, 0, 0, 0, 273, 266, 1, 0, 0, 0, 273,
		268, 1, 0, 0, 0, 273, 269, 1, 0, 0, 0, 273, 270, 1, 0, 0, 0, 273, 271,
		1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 325, 1, 0, 0, 0, 275, 276, 10, 18,
		0, 0, 276, 277, 7, 2, 0, 0, 277, 324, 3, 48, 24, 19, 278, 279, 10, 16,
		0, 0, 279, 280, 7, 3, 0, 0, 280, 324, 3, 48, 24, 17, 281, 282, 10, 15,
		0, 0, 282, 283, 7, 4, 0, 0, 283, 324, 3, 48, 24, 16, 284, 285, 10, 14,
		0, 0, 285, 286, 7, 5, 0, 0, 286, 324, 3, 48, 24, 15, 287, 288, 10, 13,
		0, 0, 288, 289, 7, 6, 0, 0, 289, 324, 3, 48, 24, 14, 290, 291, 10, 12,
		0, 0, 291, 292, 5, 45, 0, 0, 292, 324, 3, 48, 24, 13, 293, 294, 10, 11,
		0, 0, 294, 295, 5, 44, 0, 0, 295, 324, 3, 48, 24, 12, 296, 297, 10, 10,
		0, 0, 297, 298, 5, 43, 0, 0, 298, 324, 3, 48, 24, 11, 299, 300, 10, 9,
		0, 0, 300, 302, 5, 29, 0, 0, 301, 303, 3, 46, 23, 0, 302, 301, 1, 0, 0,
		0, 302, 303, 1, 0, 0, 0, 303, 304, 1, 0, 0, 0, 304, 324, 5, 30, 0, 0, 305,
		306, 10, 8, 0, 0, 306, 307, 5, 1, 0, 0, 307, 324, 5, 56, 0, 0, 308, 309,
		10, 7, 0, 0, 309, 310, 5, 31, 0, 0, 310, 311, 3, 48, 24, 0, 311, 312, 5,
		32, 0, 0, 312, 324, 1, 0, 0, 0, 313, 314, 10, 6, 0, 0, 314, 316, 5, 31,
		0, 0, 315, 317, 3, 48, 24, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0, 0,
		0, 317, 318, 1, 0, 0, 0, 318, 320, 5, 3, 0, 0, 319, 321, 3, 48, 24, 0,
		320, 319, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0, 321, 322, 1, 0, 0, 0, 322,
		324, 5, 32, 0, 0, 323, 275, 1, 0, 0, 0, 323, 278, 1, 0, 0, 0, 323, 281,
		1, 0, 0, 0, 323, 284, 1, 0, 0, 0, 323, 287, 1, 0, 0, 0, 323, 290, 1, 0,
		0, 0, 323, 293, 1, 0, 0, 0, 323, 296, 1, 0, 0, 0, 323, 299, 1, 0, 0, 0,
		323, 305, 1, 0, 0, 0, 323, 308, 1, 0, 0, 0, 323, 313, 1, 0, 0, 0, 324,
		327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 49, 1,
		0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 329, 3, 56, 28, 0, 329, 331, 5, 29,
		0, 0, 330, 332, 3, 30, 15, 0, 331, 330, 1, 0, 0, 0, 331, 332, 1, 0, 0,
		0, 332, 333, 1, 0, 0, 0, 333, 335, 5, 30, 0, 0, 334, 336, 3, 56, 28, 0,
		335, 334, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337,
		338, 3, 4, 2, 0, 338, 51, 1, 0, 0, 0, 339, 340, 3, 56, 28, 0, 340, 341,
		5, 33, 0, 0, 341, 342, 5, 56, 0, 0, 342, 343, 5, 3, 0, 0, 343, 350, 3,
		48, 24, 0, 344, 345, 5, 2, 0, 0, 345, 346, 5, 56, 0, 0, 346, 347, 5, 3,
		0, 0, 347, 349, 3, 48, 24, 0, 348, 344, 1, 0, 0, 0, 349, 352, 1, 0, 0,
		0, 350, 348, 1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 354, 1, 0, 0, 0, 352,
		350, 1, 0, 0, 0, 353, 355, 5, 2, 0, 0, 354, 353, 1, 0, 0, 0, 354, 355,
		1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357, 5, 34, 0, 0, 357, 53, 1, 0,
		0, 0, 358, 359, 5, 29, 0, 0, 359, 367, 3, 48, 24, 0, 360, 368, 5, 2, 0,
		0, 361, 362, 5, 2, 0, 0, 362, 364, 3, 48, 24, 0, 363, 361, 1, 0, 0, 0,
		364, 365, 1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366,
		368, 1, 0, 0, 0, 367, 360, 1, 0, 0, 0, 367, 363, 1, 0, 0, 0, 368, 369,
		1, 0, 0, 0, 369, 370, 5, 30, 0, 0, 370, 55, 1, 0, 0, 0, 371, 376, 3, 58,
		29, 0, 372, 376, 5, 24, 0, 0, 373, 376, 5, 25, 0, 0, 374, 376, 5, 26, 0,
		0, 375, 371, 1, 0, 0, 0, 375, 372, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375,
		374, 1, 0, 0, 0, 376, 57, 1, 0, 0, 0, 377, 378, 6, 29, -1, 0, 378, 379,
		5, 56, 0, 0, 379, 385, 1, 0, 0, 0, 380, 381, 10, 1, 0, 0, 381, 382, 5,
		1, 0, 0, 382, 384, 5, 56, 0, 0, 383, 380, 1, 0, 0, 0, 384, 387, 1, 0, 0,
		0, 385, 383, 1, 0, 0, 0, 385, 386, 1, 0, 0, 0, 386, 59, 1, 0, 0, 0, 387,
		385, 1, 0, 0, 0, 388, 393, 5, 56, 0, 0, 389, 390, 5, 2, 0, 0, 390, 392,
		5, 56, 0, 0, 391, 389, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391, 1, 0,
		0, 0, 393, 394, 1, 0, 0, 0, 394, 61, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0,
		396, 399, 3, 64, 32, 0, 397, 399, 3, 66, 33, 0, 398, 396, 1, 0, 0, 0, 398,
		397, 1, 0, 0, 0, 399, 63, 1, 0, 0, 0, 400, 401, 5, 57, 0, 0, 401, 65, 1,
		0, 0, 0, 402, 403, 5, 56, 0, 0, 403, 404, 5, 57, 0, 0, 404, 67, 1, 0, 0,
		0, 40, 71, 80, 89, 94, 98, 112, 118, 126, 128, 138, 146, 152, 157, 163,
		169, 173, 183, 190, 203, 210, 218, 225, 237, 253, 273, 302, 316, 320, 323,
		325, 331, 335, 350, 354, 365, 367, 375, 385, 393, 398,
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
	CaliburnParserEOF         = antlr.TokenEOF
	CaliburnParserPERIOD      = 1
	CaliburnParserCOMMA       = 2
	CaliburnParserCOLON       = 3
	CaliburnParserQUESTION    = 4
	CaliburnParserTerminator  = 5
	CaliburnParserASSIGN      = 6
	CaliburnParserIF          = 7
	CaliburnParserELSE        = 8
	CaliburnParserFOR         = 9
	CaliburnParserSWITCH      = 10
	CaliburnParserCASE        = 11
	CaliburnParserDEFAULT     = 12
	CaliburnParserFALLTHROUGH = 13
	CaliburnParserBREAK       = 14
	CaliburnParserCONTINUE    = 15
	CaliburnParserRETURN      = 16
	CaliburnParserUSING       = 17
	CaliburnParserAS          = 18
	CaliburnParserIN          = 19
	CaliburnParserNULL        = 20
	CaliburnParserVAR         = 21
	CaliburnParserCONST       = 22
	CaliburnParserTYPE        = 23
	CaliburnParserFUNC        = 24
	CaliburnParserSTRUCT      = 25
	CaliburnParserTUPLE       = 26
	CaliburnParserCLASS       = 27
	CaliburnParserINTERFACE   = 28
	CaliburnParserL_PAREN     = 29
	CaliburnParserR_PAREN     = 30
	CaliburnParserL_S_BRACK   = 31
	CaliburnParserR_S_BRACK   = 32
	CaliburnParserL_C_BRACK   = 33
	CaliburnParserR_C_BRACK   = 34
	CaliburnParserOP_ADD      = 35
	CaliburnParserOP_SUB      = 36
	CaliburnParserOP_NOT      = 37
	CaliburnParserOP_MUL      = 38
	CaliburnParserOP_DIV      = 39
	CaliburnParserOP_MOD      = 40
	CaliburnParserOP_POW      = 41
	CaliburnParserOP_ROOT     = 42
	CaliburnParserOP_OR       = 43
	CaliburnParserOP_XOR      = 44
	CaliburnParserOP_AND      = 45
	CaliburnParserOP_LSHIFT   = 46
	CaliburnParserOP_RSHIFT   = 47
	CaliburnParserOP_EQU      = 48
	CaliburnParserOP_NEQ      = 49
	CaliburnParserOP_GTE      = 50
	CaliburnParserOP_GRT      = 51
	CaliburnParserOP_LTE      = 52
	CaliburnParserOP_LST      = 53
	CaliburnParserWhiteSpace  = 54
	CaliburnParserComment     = 55
	CaliburnParserIdentifier  = 56
	CaliburnParserLiteral     = 57
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_module                         = 0
	CaliburnParserRULE_definition                     = 1
	CaliburnParserRULE_scoped_block                   = 2
	CaliburnParserRULE_statement                      = 3
	CaliburnParserRULE_jump_statement                 = 4
	CaliburnParserRULE_return_statement               = 5
	CaliburnParserRULE_break_statement                = 6
	CaliburnParserRULE_continue_statement             = 7
	CaliburnParserRULE_control_statement              = 8
	CaliburnParserRULE_if_statement                   = 9
	CaliburnParserRULE_for_statement                  = 10
	CaliburnParserRULE_switch_statement               = 11
	CaliburnParserRULE_inline_statement               = 12
	CaliburnParserRULE_function_declaration_statement = 13
	CaliburnParserRULE_assign_statement               = 14
	CaliburnParserRULE_assign_declarations            = 15
	CaliburnParserRULE_assign_declaration             = 16
	CaliburnParserRULE_declared_assign_declarations   = 17
	CaliburnParserRULE_declared_assign_declaration    = 18
	CaliburnParserRULE_undeclared_assign_declarations = 19
	CaliburnParserRULE_undeclared_assign_declaration  = 20
	CaliburnParserRULE_operator_assign_statement      = 21
	CaliburnParserRULE_expression_statement           = 22
	CaliburnParserRULE_expressions                    = 23
	CaliburnParserRULE_expression                     = 24
	CaliburnParserRULE_function_expression            = 25
	CaliburnParserRULE_struct_expression              = 26
	CaliburnParserRULE_tuple_expression               = 27
	CaliburnParserRULE_type_expression                = 28
	CaliburnParserRULE_complex_type_expression        = 29
	CaliburnParserRULE_identifiers                    = 30
	CaliburnParserRULE_literal_atom                   = 31
	CaliburnParserRULE_untyped_literal_atom           = 32
	CaliburnParserRULE_typed_literal_atom             = 33
)

// IModuleContext is an interface to support dynamic dispatch.
type IModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDefinition() []IDefinitionContext
	Definition(i int) IDefinitionContext

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

func (s *ModuleContext) AllDefinition() []IDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDefinitionContext); ok {
			len++
		}
	}

	tst := make([]IDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDefinitionContext); ok {
			tst[i] = t.(IDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *ModuleContext) Definition(i int) IDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefinitionContext); ok {
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

	return t.(IDefinitionContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitModule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CaliburnParserRULE_module)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057594155368448) != 0 {
		{
			p.SetState(68)
			p.Definition()
		}

		p.SetState(73)
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

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_declaration_statement() IFunction_declaration_statementContext

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

func (s *DefinitionContext) Function_declaration_statement() IFunction_declaration_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declaration_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_statementContext)
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
	p.EnterRule(localctx, 2, CaliburnParserRULE_definition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(74)
		p.Function_declaration_statement()
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

// IScoped_blockContext is an interface to support dynamic dispatch.
type IScoped_blockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_C_BRACK() antlr.TerminalNode
	R_C_BRACK() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsScoped_blockContext differentiates from other interfaces.
	IsScoped_blockContext()
}

type Scoped_blockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScoped_blockContext() *Scoped_blockContext {
	var p = new(Scoped_blockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_scoped_block
	return p
}

func InitEmptyScoped_blockContext(p *Scoped_blockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_scoped_block
}

func (*Scoped_blockContext) IsScoped_blockContext() {}

func NewScoped_blockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Scoped_blockContext {
	var p = new(Scoped_blockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_scoped_block

	return p
}

func (s *Scoped_blockContext) GetParser() antlr.Parser { return s.parser }

func (s *Scoped_blockContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *Scoped_blockContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *Scoped_blockContext) AllStatement() []IStatementContext {
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

func (s *Scoped_blockContext) Statement(i int) IStatementContext {
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

func (s *Scoped_blockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Scoped_blockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Scoped_blockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterScoped_block(s)
	}
}

func (s *Scoped_blockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitScoped_block(s)
	}
}

func (s *Scoped_blockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitScoped_block(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Scoped_block() (localctx IScoped_blockContext) {
	localctx = NewScoped_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CaliburnParserRULE_scoped_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173034025895552) != 0 {
		{
			p.SetState(77)
			p.Statement()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(83)
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
	p.EnterRule(localctx, 6, CaliburnParserRULE_statement)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Expression_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(87)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(88)
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
	p.EnterRule(localctx, 8, CaliburnParserRULE_jump_statement)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserRETURN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Return_statement()
		}

	case CaliburnParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Break_statement()
		}

	case CaliburnParserCONTINUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
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

	// Getter signatures
	RETURN() antlr.TerminalNode
	Terminator() antlr.TerminalNode
	Expressions() IExpressionsContext

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

func (s *Return_statementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserRETURN, 0)
}

func (s *Return_statementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *Return_statementContext) Expressions() IExpressionsContext {
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

func (s *Return_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterReturn_statement(s)
	}
}

func (s *Return_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitReturn_statement(s)
	}
}

func (s *Return_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitReturn_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Return_statement() (localctx IReturn_statementContext) {
	localctx = NewReturn_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CaliburnParserRULE_return_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(CaliburnParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023286263808) != 0 {
		{
			p.SetState(97)
			p.Expressions()
		}

	}
	{
		p.SetState(100)
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

	// Getter signatures
	BREAK() antlr.TerminalNode
	Terminator() antlr.TerminalNode

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

func (s *Break_statementContext) BREAK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserBREAK, 0)
}

func (s *Break_statementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *Break_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterBreak_statement(s)
	}
}

func (s *Break_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitBreak_statement(s)
	}
}

func (s *Break_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitBreak_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Break_statement() (localctx IBreak_statementContext) {
	localctx = NewBreak_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CaliburnParserRULE_break_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(CaliburnParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
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

	// Getter signatures
	CONTINUE() antlr.TerminalNode
	Terminator() antlr.TerminalNode

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

func (s *Continue_statementContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserCONTINUE, 0)
}

func (s *Continue_statementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *Continue_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterContinue_statement(s)
	}
}

func (s *Continue_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitContinue_statement(s)
	}
}

func (s *Continue_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitContinue_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Continue_statement() (localctx IContinue_statementContext) {
	localctx = NewContinue_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CaliburnParserRULE_continue_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		p.Match(CaliburnParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
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
	Function_declaration_statement() IFunction_declaration_statementContext

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

func (s *Control_statementContext) Function_declaration_statement() IFunction_declaration_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_declaration_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_declaration_statementContext)
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
	p.EnterRule(localctx, 16, CaliburnParserRULE_control_statement)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.If_statement()
		}

	case CaliburnParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.For_statement()
		}

	case CaliburnParserSWITCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Switch_statement()
		}

	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserTUPLE, CaliburnParserIdentifier:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Function_declaration_statement()
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

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	AllScoped_block() []IScoped_blockContext
	Scoped_block(i int) IScoped_blockContext
	AllInline_statement() []IInline_statementContext
	Inline_statement(i int) IInline_statementContext
	ELSE() antlr.TerminalNode
	If_statement() IIf_statementContext

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

func (s *If_statementContext) IF() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIF, 0)
}

func (s *If_statementContext) Expression() IExpressionContext {
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

func (s *If_statementContext) AllScoped_block() []IScoped_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScoped_blockContext); ok {
			len++
		}
	}

	tst := make([]IScoped_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScoped_blockContext); ok {
			tst[i] = t.(IScoped_blockContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Scoped_block(i int) IScoped_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScoped_blockContext); ok {
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

	return t.(IScoped_blockContext)
}

func (s *If_statementContext) AllInline_statement() []IInline_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInline_statementContext); ok {
			len++
		}
	}

	tst := make([]IInline_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInline_statementContext); ok {
			tst[i] = t.(IInline_statementContext)
			i++
		}
	}

	return tst
}

func (s *If_statementContext) Inline_statement(i int) IInline_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementContext); ok {
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

	return t.(IInline_statementContext)
}

func (s *If_statementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserELSE, 0)
}

func (s *If_statementContext) If_statement() IIf_statementContext {
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

func (s *If_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIf_statement(s)
	}
}

func (s *If_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIf_statement(s)
	}
}

func (s *If_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIf_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CaliburnParserRULE_if_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(CaliburnParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(118)
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
			{
				p.SetState(115)
				p.Inline_statement()
			}

		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.expression(0)
	}
	{
		p.SetState(122)
		p.Scoped_block()
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserELSE {
		{
			p.SetState(123)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CaliburnParserL_C_BRACK:
			{
				p.SetState(124)
				p.Scoped_block()
			}

		case CaliburnParserIF:
			{
				p.SetState(125)
				p.If_statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IFor_statementContext is an interface to support dynamic dispatch.
type IFor_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Expression() IExpressionContext
	Scoped_block() IScoped_blockContext

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

func (s *For_statementContext) FOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserFOR, 0)
}

func (s *For_statementContext) Expression() IExpressionContext {
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

func (s *For_statementContext) Scoped_block() IScoped_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScoped_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScoped_blockContext)
}

func (s *For_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFor_statement(s)
	}
}

func (s *For_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFor_statement(s)
	}
}

func (s *For_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFor_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CaliburnParserRULE_for_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(CaliburnParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.expression(0)
	}
	{
		p.SetState(132)
		p.Scoped_block()
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

	// Getter signatures
	SWITCH() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	L_C_BRACK() antlr.TerminalNode
	R_C_BRACK() antlr.TerminalNode
	AllInline_statement() []IInline_statementContext
	Inline_statement(i int) IInline_statementContext
	AllCASE() []antlr.TerminalNode
	CASE(i int) antlr.TerminalNode
	AllScoped_block() []IScoped_blockContext
	Scoped_block(i int) IScoped_blockContext
	DEFAULT() antlr.TerminalNode
	AllType_expression() []IType_expressionContext
	Type_expression(i int) IType_expressionContext

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

func (s *Switch_statementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(CaliburnParserSWITCH, 0)
}

func (s *Switch_statementContext) AllExpression() []IExpressionContext {
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

func (s *Switch_statementContext) Expression(i int) IExpressionContext {
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

func (s *Switch_statementContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *Switch_statementContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *Switch_statementContext) AllInline_statement() []IInline_statementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInline_statementContext); ok {
			len++
		}
	}

	tst := make([]IInline_statementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInline_statementContext); ok {
			tst[i] = t.(IInline_statementContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Inline_statement(i int) IInline_statementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInline_statementContext); ok {
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

	return t.(IInline_statementContext)
}

func (s *Switch_statementContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCASE)
}

func (s *Switch_statementContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCASE, i)
}

func (s *Switch_statementContext) AllScoped_block() []IScoped_blockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IScoped_blockContext); ok {
			len++
		}
	}

	tst := make([]IScoped_blockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IScoped_blockContext); ok {
			tst[i] = t.(IScoped_blockContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Scoped_block(i int) IScoped_blockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScoped_blockContext); ok {
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

	return t.(IScoped_blockContext)
}

func (s *Switch_statementContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserDEFAULT, 0)
}

func (s *Switch_statementContext) AllType_expression() []IType_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_expressionContext); ok {
			len++
		}
	}

	tst := make([]IType_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_expressionContext); ok {
			tst[i] = t.(IType_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Type_expression(i int) IType_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
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

	return t.(IType_expressionContext)
}

func (s *Switch_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterSwitch_statement(s)
	}
}

func (s *Switch_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitSwitch_statement(s)
	}
}

func (s *Switch_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitSwitch_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CaliburnParserRULE_switch_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(CaliburnParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(135)
				p.Inline_statement()
			}

		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.expression(0)
	}
	{
		p.SetState(142)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCASE {
		{
			p.SetState(143)
			p.Match(CaliburnParserCASE)
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

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(144)
				p.Type_expression()
			}

		case 2:
			{
				p.SetState(145)
				p.expression(0)
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(148)
			p.Scoped_block()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserDEFAULT {
		{
			p.SetState(155)
			p.Match(CaliburnParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Scoped_block()
		}

	}
	{
		p.SetState(159)
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
	p.EnterRule(localctx, 24, CaliburnParserRULE_inline_statement)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
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

// IFunction_declaration_statementContext is an interface to support dynamic dispatch.
type IFunction_declaration_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetType_ returns the type_ rule contexts.
	GetType_() IType_expressionContext

	// GetReturn_type returns the return_type rule contexts.
	GetReturn_type() IType_expressionContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IType_expressionContext)

	// SetReturn_type sets the return_type rule contexts.
	SetReturn_type(IType_expressionContext)

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Scoped_block() IScoped_blockContext
	AllType_expression() []IType_expressionContext
	Type_expression(i int) IType_expressionContext
	Identifier() antlr.TerminalNode
	Assign_declarations() IAssign_declarationsContext

	// IsFunction_declaration_statementContext differentiates from other interfaces.
	IsFunction_declaration_statementContext()
}

type Function_declaration_statementContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	type_       IType_expressionContext
	name        antlr.Token
	return_type IType_expressionContext
}

func NewEmptyFunction_declaration_statementContext() *Function_declaration_statementContext {
	var p = new(Function_declaration_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_declaration_statement
	return p
}

func InitEmptyFunction_declaration_statementContext(p *Function_declaration_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_declaration_statement
}

func (*Function_declaration_statementContext) IsFunction_declaration_statementContext() {}

func NewFunction_declaration_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declaration_statementContext {
	var p = new(Function_declaration_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_function_declaration_statement

	return p
}

func (s *Function_declaration_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declaration_statementContext) GetName() antlr.Token { return s.name }

func (s *Function_declaration_statementContext) SetName(v antlr.Token) { s.name = v }

func (s *Function_declaration_statementContext) GetType_() IType_expressionContext { return s.type_ }

func (s *Function_declaration_statementContext) GetReturn_type() IType_expressionContext {
	return s.return_type
}

func (s *Function_declaration_statementContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *Function_declaration_statementContext) SetReturn_type(v IType_expressionContext) {
	s.return_type = v
}

func (s *Function_declaration_statementContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *Function_declaration_statementContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *Function_declaration_statementContext) Scoped_block() IScoped_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScoped_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScoped_blockContext)
}

func (s *Function_declaration_statementContext) AllType_expression() []IType_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_expressionContext); ok {
			len++
		}
	}

	tst := make([]IType_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_expressionContext); ok {
			tst[i] = t.(IType_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_declaration_statementContext) Type_expression(i int) IType_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
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

	return t.(IType_expressionContext)
}

func (s *Function_declaration_statementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Function_declaration_statementContext) Assign_declarations() IAssign_declarationsContext {
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

func (s *Function_declaration_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declaration_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declaration_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunction_declaration_statement(s)
	}
}

func (s *Function_declaration_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunction_declaration_statement(s)
	}
}

func (s *Function_declaration_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunction_declaration_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Function_declaration_statement() (localctx IFunction_declaration_statementContext) {
	localctx = NewFunction_declaration_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CaliburnParserRULE_function_declaration_statement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)

		var _x = p.Type_expression()

		localctx.(*Function_declaration_statementContext).type_ = _x
	}
	{
		p.SetState(166)

		var _m = p.Match(CaliburnParserIdentifier)

		localctx.(*Function_declaration_statementContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(CaliburnParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057605431754752) != 0 {
		{
			p.SetState(168)
			p.Assign_declarations()
		}

	}
	{
		p.SetState(171)
		p.Match(CaliburnParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057594155368448) != 0 {
		{
			p.SetState(172)

			var _x = p.Type_expression()

			localctx.(*Function_declaration_statementContext).return_type = _x
		}

	}
	{
		p.SetState(175)
		p.Scoped_block()
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

	// Getter signatures
	Assign_declarations() IAssign_declarationsContext
	ASSIGN() antlr.TerminalNode
	Expressions() IExpressionsContext
	Terminator() antlr.TerminalNode
	Operator_assign_statement() IOperator_assign_statementContext

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

func (s *Assign_statementContext) Assign_declarations() IAssign_declarationsContext {
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

func (s *Assign_statementContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserASSIGN, 0)
}

func (s *Assign_statementContext) Expressions() IExpressionsContext {
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

func (s *Assign_statementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *Assign_statementContext) Operator_assign_statement() IOperator_assign_statementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperator_assign_statementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperator_assign_statementContext)
}

func (s *Assign_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssign_statement(s)
	}
}

func (s *Assign_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssign_statement(s)
	}
}

func (s *Assign_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssign_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_statement() (localctx IAssign_statementContext) {
	localctx = NewAssign_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CaliburnParserRULE_assign_statement)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Assign_declarations()
		}
		{
			p.SetState(178)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Expressions()
		}
		{
			p.SetState(180)
			p.Match(CaliburnParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Operator_assign_statement()
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

// IAssign_declarationsContext is an interface to support dynamic dispatch.
type IAssign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign_declaration() IAssign_declarationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllAssign_statement() []IAssign_statementContext
	Assign_statement(i int) IAssign_statementContext

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

func (s *Assign_declarationsContext) Assign_declaration() IAssign_declarationContext {
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

func (s *Assign_declarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Assign_declarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Assign_declarationsContext) AllAssign_statement() []IAssign_statementContext {
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

func (s *Assign_declarationsContext) Assign_statement(i int) IAssign_statementContext {
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

func (s *Assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_declarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAssign_declarations(s)
	}
}

func (s *Assign_declarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAssign_declarations(s)
	}
}

func (s *Assign_declarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitAssign_declarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Assign_declarations() (localctx IAssign_declarationsContext) {
	localctx = NewAssign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CaliburnParserRULE_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(185)
		p.Assign_declaration()
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(186)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Assign_statement()
		}

		p.SetState(192)
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

type UndeclaredAssignDeclarationContext struct {
	Assign_declarationContext
}

func NewUndeclaredAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UndeclaredAssignDeclarationContext {
	var p = new(UndeclaredAssignDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *UndeclaredAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndeclaredAssignDeclarationContext) Undeclared_assign_declaration() IUndeclared_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndeclared_assign_declarationContext)
}

func (s *UndeclaredAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUndeclaredAssignDeclaration(s)
	}
}

func (s *UndeclaredAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUndeclaredAssignDeclaration(s)
	}
}

func (s *UndeclaredAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUndeclaredAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclaredAssignDeclarationContext struct {
	Assign_declarationContext
}

func NewDeclaredAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclaredAssignDeclarationContext {
	var p = new(DeclaredAssignDeclarationContext)

	InitEmptyAssign_declarationContext(&p.Assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Assign_declarationContext))

	return p
}

func (s *DeclaredAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaredAssignDeclarationContext) Declared_assign_declaration() IDeclared_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclared_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclared_assign_declarationContext)
}

func (s *DeclaredAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDeclaredAssignDeclaration(s)
	}
}

func (s *DeclaredAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDeclaredAssignDeclaration(s)
	}
}

func (s *DeclaredAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDeclaredAssignDeclaration(s)

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

func (s *TupleDestrutureAssignDeclarationContext) AllL_S_BRACK() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserL_S_BRACK)
}

func (s *TupleDestrutureAssignDeclarationContext) L_S_BRACK(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, i)
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
	p.EnterRule(localctx, 32, CaliburnParserRULE_assign_declaration)
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			p.Declared_assign_declaration()
		}

	case 2:
		localctx = NewUndeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Undeclared_assign_declaration()
		}

	case 3:
		localctx = NewStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Assign_declarations()
		}
		{
			p.SetState(197)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewTupleDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Assign_declarations()
		}
		{
			p.SetState(201)
			p.Match(CaliburnParserL_S_BRACK)
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

// IDeclared_assign_declarationsContext is an interface to support dynamic dispatch.
type IDeclared_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllDeclared_assign_declaration() []IDeclared_assign_declarationContext
	Declared_assign_declaration(i int) IDeclared_assign_declarationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDeclared_assign_declarationsContext differentiates from other interfaces.
	IsDeclared_assign_declarationsContext()
}

type Declared_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclared_assign_declarationsContext() *Declared_assign_declarationsContext {
	var p = new(Declared_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_declared_assign_declarations
	return p
}

func InitEmptyDeclared_assign_declarationsContext(p *Declared_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_declared_assign_declarations
}

func (*Declared_assign_declarationsContext) IsDeclared_assign_declarationsContext() {}

func NewDeclared_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declared_assign_declarationsContext {
	var p = new(Declared_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_declared_assign_declarations

	return p
}

func (s *Declared_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Declared_assign_declarationsContext) AllDeclared_assign_declaration() []IDeclared_assign_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclared_assign_declarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclared_assign_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclared_assign_declarationContext); ok {
			tst[i] = t.(IDeclared_assign_declarationContext)
			i++
		}
	}

	return tst
}

func (s *Declared_assign_declarationsContext) Declared_assign_declaration(i int) IDeclared_assign_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclared_assign_declarationContext); ok {
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

	return t.(IDeclared_assign_declarationContext)
}

func (s *Declared_assign_declarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Declared_assign_declarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Declared_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declared_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declared_assign_declarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterDeclared_assign_declarations(s)
	}
}

func (s *Declared_assign_declarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitDeclared_assign_declarations(s)
	}
}

func (s *Declared_assign_declarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitDeclared_assign_declarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Declared_assign_declarations() (localctx IDeclared_assign_declarationsContext) {
	localctx = NewDeclared_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CaliburnParserRULE_declared_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Declared_assign_declaration()
	}
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(206)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.Declared_assign_declaration()
		}

		p.SetState(212)
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

// IDeclared_assign_declarationContext is an interface to support dynamic dispatch.
type IDeclared_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclared_assign_declarationContext differentiates from other interfaces.
	IsDeclared_assign_declarationContext()
}

type Declared_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclared_assign_declarationContext() *Declared_assign_declarationContext {
	var p = new(Declared_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_declared_assign_declaration
	return p
}

func InitEmptyDeclared_assign_declarationContext(p *Declared_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_declared_assign_declaration
}

func (*Declared_assign_declarationContext) IsDeclared_assign_declarationContext() {}

func NewDeclared_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declared_assign_declarationContext {
	var p = new(Declared_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_declared_assign_declaration

	return p
}

func (s *Declared_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Declared_assign_declarationContext) CopyAll(ctx *Declared_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declared_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declared_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UntypedDeclaredAssignDeclarationContext struct {
	Declared_assign_declarationContext
}

func NewUntypedDeclaredAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedDeclaredAssignDeclarationContext {
	var p = new(UntypedDeclaredAssignDeclarationContext)

	InitEmptyDeclared_assign_declarationContext(&p.Declared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declared_assign_declarationContext))

	return p
}

func (s *UntypedDeclaredAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedDeclaredAssignDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserVAR, 0)
}

func (s *UntypedDeclaredAssignDeclarationContext) Undeclared_assign_declaration() IUndeclared_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndeclared_assign_declarationContext)
}

func (s *UntypedDeclaredAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedDeclaredAssignDeclaration(s)
	}
}

func (s *UntypedDeclaredAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedDeclaredAssignDeclaration(s)
	}
}

func (s *UntypedDeclaredAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedDeclaredAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedDeclaredAssignDeclarationContext struct {
	Declared_assign_declarationContext
	type_ IType_expressionContext
}

func NewTypedDeclaredAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclaredAssignDeclarationContext {
	var p = new(TypedDeclaredAssignDeclarationContext)

	InitEmptyDeclared_assign_declarationContext(&p.Declared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declared_assign_declarationContext))

	return p
}

func (s *TypedDeclaredAssignDeclarationContext) GetType_() IType_expressionContext { return s.type_ }

func (s *TypedDeclaredAssignDeclarationContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *TypedDeclaredAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclaredAssignDeclarationContext) Undeclared_assign_declaration() IUndeclared_assign_declarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndeclared_assign_declarationContext)
}

func (s *TypedDeclaredAssignDeclarationContext) Type_expression() IType_expressionContext {
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

func (s *TypedDeclaredAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedDeclaredAssignDeclaration(s)
	}
}

func (s *TypedDeclaredAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedDeclaredAssignDeclaration(s)
	}
}

func (s *TypedDeclaredAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedDeclaredAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Declared_assign_declaration() (localctx IDeclared_assign_declarationContext) {
	localctx = NewDeclared_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CaliburnParserRULE_declared_assign_declaration)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserTUPLE, CaliburnParserIdentifier:
		localctx = NewTypedDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)

			var _x = p.Type_expression()

			localctx.(*TypedDeclaredAssignDeclarationContext).type_ = _x
		}
		{
			p.SetState(214)
			p.Undeclared_assign_declaration()
		}

	case CaliburnParserVAR:
		localctx = NewUntypedDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(216)
			p.Match(CaliburnParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.Undeclared_assign_declaration()
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

// IUndeclared_assign_declarationsContext is an interface to support dynamic dispatch.
type IUndeclared_assign_declarationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUndeclared_assign_declaration() []IUndeclared_assign_declarationContext
	Undeclared_assign_declaration(i int) IUndeclared_assign_declarationContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsUndeclared_assign_declarationsContext differentiates from other interfaces.
	IsUndeclared_assign_declarationsContext()
}

type Undeclared_assign_declarationsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndeclared_assign_declarationsContext() *Undeclared_assign_declarationsContext {
	var p = new(Undeclared_assign_declarationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declarations
	return p
}

func InitEmptyUndeclared_assign_declarationsContext(p *Undeclared_assign_declarationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declarations
}

func (*Undeclared_assign_declarationsContext) IsUndeclared_assign_declarationsContext() {}

func NewUndeclared_assign_declarationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Undeclared_assign_declarationsContext {
	var p = new(Undeclared_assign_declarationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declarations

	return p
}

func (s *Undeclared_assign_declarationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Undeclared_assign_declarationsContext) AllUndeclared_assign_declaration() []IUndeclared_assign_declarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUndeclared_assign_declarationContext); ok {
			len++
		}
	}

	tst := make([]IUndeclared_assign_declarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUndeclared_assign_declarationContext); ok {
			tst[i] = t.(IUndeclared_assign_declarationContext)
			i++
		}
	}

	return tst
}

func (s *Undeclared_assign_declarationsContext) Undeclared_assign_declaration(i int) IUndeclared_assign_declarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationContext); ok {
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

	return t.(IUndeclared_assign_declarationContext)
}

func (s *Undeclared_assign_declarationsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Undeclared_assign_declarationsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Undeclared_assign_declarationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Undeclared_assign_declarationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Undeclared_assign_declarationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUndeclared_assign_declarations(s)
	}
}

func (s *Undeclared_assign_declarationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUndeclared_assign_declarations(s)
	}
}

func (s *Undeclared_assign_declarationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUndeclared_assign_declarations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Undeclared_assign_declarations() (localctx IUndeclared_assign_declarationsContext) {
	localctx = NewUndeclared_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CaliburnParserRULE_undeclared_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Undeclared_assign_declaration()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(221)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Undeclared_assign_declaration()
		}

		p.SetState(227)
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

// IUndeclared_assign_declarationContext is an interface to support dynamic dispatch.
type IUndeclared_assign_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsUndeclared_assign_declarationContext differentiates from other interfaces.
	IsUndeclared_assign_declarationContext()
}

type Undeclared_assign_declarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUndeclared_assign_declarationContext() *Undeclared_assign_declarationContext {
	var p = new(Undeclared_assign_declarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declaration
	return p
}

func InitEmptyUndeclared_assign_declarationContext(p *Undeclared_assign_declarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declaration
}

func (*Undeclared_assign_declarationContext) IsUndeclared_assign_declarationContext() {}

func NewUndeclared_assign_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Undeclared_assign_declarationContext {
	var p = new(Undeclared_assign_declarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_undeclared_assign_declaration

	return p
}

func (s *Undeclared_assign_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Undeclared_assign_declarationContext) CopyAll(ctx *Undeclared_assign_declarationContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Undeclared_assign_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Undeclared_assign_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UndeclaredAtomAssignDeclarationContext struct {
	Undeclared_assign_declarationContext
	var_ IType_expressionContext
}

func NewUndeclaredAtomAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UndeclaredAtomAssignDeclarationContext {
	var p = new(UndeclaredAtomAssignDeclarationContext)

	InitEmptyUndeclared_assign_declarationContext(&p.Undeclared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Undeclared_assign_declarationContext))

	return p
}

func (s *UndeclaredAtomAssignDeclarationContext) GetVar_() IType_expressionContext { return s.var_ }

func (s *UndeclaredAtomAssignDeclarationContext) SetVar_(v IType_expressionContext) { s.var_ = v }

func (s *UndeclaredAtomAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndeclaredAtomAssignDeclarationContext) Type_expression() IType_expressionContext {
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

func (s *UndeclaredAtomAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUndeclaredAtomAssignDeclaration(s)
	}
}

func (s *UndeclaredAtomAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUndeclaredAtomAssignDeclaration(s)
	}
}

func (s *UndeclaredAtomAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUndeclaredAtomAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type UndeclaredTupleDestrutureAssignDeclarationContext struct {
	Undeclared_assign_declarationContext
}

func NewUndeclaredTupleDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UndeclaredTupleDestrutureAssignDeclarationContext {
	var p = new(UndeclaredTupleDestrutureAssignDeclarationContext)

	InitEmptyUndeclared_assign_declarationContext(&p.Undeclared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Undeclared_assign_declarationContext))

	return p
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) Undeclared_assign_declarations() IUndeclared_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndeclared_assign_declarationsContext)
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUndeclaredTupleDestrutureAssignDeclaration(s)
	}
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUndeclaredTupleDestrutureAssignDeclaration(s)
	}
}

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUndeclaredTupleDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

type UndeclaredStructDestrutureAssignDeclarationContext struct {
	Undeclared_assign_declarationContext
}

func NewUndeclaredStructDestrutureAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UndeclaredStructDestrutureAssignDeclarationContext {
	var p = new(UndeclaredStructDestrutureAssignDeclarationContext)

	InitEmptyUndeclared_assign_declarationContext(&p.Undeclared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Undeclared_assign_declarationContext))

	return p
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) Undeclared_assign_declarations() IUndeclared_assign_declarationsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUndeclared_assign_declarationsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUndeclared_assign_declarationsContext)
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUndeclaredStructDestrutureAssignDeclaration(s)
	}
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUndeclaredStructDestrutureAssignDeclaration(s)
	}
}

func (s *UndeclaredStructDestrutureAssignDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUndeclaredStructDestrutureAssignDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Undeclared_assign_declaration() (localctx IUndeclared_assign_declarationContext) {
	localctx = NewUndeclared_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CaliburnParserRULE_undeclared_assign_declaration)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserTUPLE, CaliburnParserIdentifier:
		localctx = NewUndeclaredAtomAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)

			var _x = p.Type_expression()

			localctx.(*UndeclaredAtomAssignDeclarationContext).var_ = _x
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewUndeclaredStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(230)
			p.Undeclared_assign_declarations()
		}
		{
			p.SetState(231)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserL_PAREN:
		localctx = NewUndeclaredTupleDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(233)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Undeclared_assign_declarations()
		}
		{
			p.SetState(235)
			p.Match(CaliburnParserR_PAREN)
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

// IOperator_assign_statementContext is an interface to support dynamic dispatch.
type IOperator_assign_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsOperator_assign_statementContext differentiates from other interfaces.
	IsOperator_assign_statementContext()
}

type Operator_assign_statementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperator_assign_statementContext() *Operator_assign_statementContext {
	var p = new(Operator_assign_statementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_operator_assign_statement
	return p
}

func InitEmptyOperator_assign_statementContext(p *Operator_assign_statementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_operator_assign_statement
}

func (*Operator_assign_statementContext) IsOperator_assign_statementContext() {}

func NewOperator_assign_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operator_assign_statementContext {
	var p = new(Operator_assign_statementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_operator_assign_statement

	return p
}

func (s *Operator_assign_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Operator_assign_statementContext) CopyAll(ctx *Operator_assign_statementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Operator_assign_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operator_assign_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OperatorAssignmentContext struct {
	Operator_assign_statementContext
	var_ IExpressionContext
	op   antlr.Token
}

func NewOperatorAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OperatorAssignmentContext {
	var p = new(OperatorAssignmentContext)

	InitEmptyOperator_assign_statementContext(&p.Operator_assign_statementContext)
	p.parser = parser
	p.CopyAll(ctx.(*Operator_assign_statementContext))

	return p
}

func (s *OperatorAssignmentContext) GetOp() antlr.Token { return s.op }

func (s *OperatorAssignmentContext) SetOp(v antlr.Token) { s.op = v }

func (s *OperatorAssignmentContext) GetVar_() IExpressionContext { return s.var_ }

func (s *OperatorAssignmentContext) SetVar_(v IExpressionContext) { s.var_ = v }

func (s *OperatorAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserASSIGN, 0)
}

func (s *OperatorAssignmentContext) Expressions() IExpressionsContext {
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

func (s *OperatorAssignmentContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *OperatorAssignmentContext) Expression() IExpressionContext {
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

func (s *OperatorAssignmentContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ADD, 0)
}

func (s *OperatorAssignmentContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_SUB, 0)
}

func (s *OperatorAssignmentContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MUL, 0)
}

func (s *OperatorAssignmentContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_DIV, 0)
}

func (s *OperatorAssignmentContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MOD, 0)
}

func (s *OperatorAssignmentContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_POW, 0)
}

func (s *OperatorAssignmentContext) OP_ROOT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ROOT, 0)
}

func (s *OperatorAssignmentContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_OR, 0)
}

func (s *OperatorAssignmentContext) OP_XOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_XOR, 0)
}

func (s *OperatorAssignmentContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_AND, 0)
}

func (s *OperatorAssignmentContext) OP_LSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LSHIFT, 0)
}

func (s *OperatorAssignmentContext) OP_RSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_RSHIFT, 0)
}

func (s *OperatorAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterOperatorAssignment(s)
	}
}

func (s *OperatorAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitOperatorAssignment(s)
	}
}

func (s *OperatorAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitOperatorAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Operator_assign_statement() (localctx IOperator_assign_statementContext) {
	localctx = NewOperator_assign_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CaliburnParserRULE_operator_assign_statement)
	var _la int

	localctx = NewOperatorAssignmentContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)

		var _x = p.expression(0)

		localctx.(*OperatorAssignmentContext).var_ = _x
	}
	{
		p.SetState(240)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OperatorAssignmentContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281303178018816) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OperatorAssignmentContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(241)
		p.Match(CaliburnParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Expressions()
	}
	{
		p.SetState(243)
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

// IExpression_statementContext is an interface to support dynamic dispatch.
type IExpression_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expressions() IExpressionsContext
	Terminator() antlr.TerminalNode

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

func (s *Expression_statementContext) Expressions() IExpressionsContext {
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

func (s *Expression_statementContext) Terminator() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTerminator, 0)
}

func (s *Expression_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpression_statement(s)
	}
}

func (s *Expression_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpression_statement(s)
	}
}

func (s *Expression_statementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpression_statement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expression_statement() (localctx IExpression_statementContext) {
	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CaliburnParserRULE_expression_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.Expressions()
	}
	{
		p.SetState(246)
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

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *ExpressionsContext) AllExpression() []IExpressionContext {
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

func (s *ExpressionsContext) Expression(i int) IExpressionContext {
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

func (s *ExpressionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *ExpressionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (s *ExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CaliburnParserRULE_expressions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.expression(0)
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(249)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(250)
			p.expression(0)
		}

		p.SetState(255)
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
	exp ITuple_expressionContext
}

func NewTupleExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleExpressionContext {
	var p = new(TupleExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *TupleExpressionContext) GetExp() ITuple_expressionContext { return s.exp }

func (s *TupleExpressionContext) SetExp(v ITuple_expressionContext) { s.exp = v }

func (s *TupleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleExpressionContext) Tuple_expression() ITuple_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITuple_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITuple_expressionContext)
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
	exp IFunction_expressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetExp() IFunction_expressionContext { return s.exp }

func (s *FunctionExpressionContext) SetExp(v IFunction_expressionContext) { s.exp = v }

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Function_expression() IFunction_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_expressionContext)
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
	exp IStruct_expressionContext
}

func NewStructExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExpressionContext {
	var p = new(StructExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StructExpressionContext) GetExp() IStruct_expressionContext { return s.exp }

func (s *StructExpressionContext) SetExp(v IStruct_expressionContext) { s.exp = v }

func (s *StructExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExpressionContext) Struct_expression() IStruct_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_expressionContext)
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
	_startState := 48
	p.EnterRecursionRule(localctx, 48, CaliburnParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(257)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)

			var _x = p.Type_expression()

			localctx.(*CastExpressionContext).type_ = _x
		}
		{
			p.SetState(259)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)

			var _x = p.expression(20)

			localctx.(*CastExpressionContext).exp = _x
		}

	case 2:
		localctx = NewBracketedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(262)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)

			var _x = p.expression(0)

			localctx.(*BracketedExpressionContext).exp = _x
		}
		{
			p.SetState(264)
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
			p.SetState(266)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryExpressionContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&240518168576) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryExpressionContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(267)

			var _x = p.expression(17)

			localctx.(*UnaryExpressionContext).exp = _x
		}

	case 4:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(268)

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
			p.SetState(269)

			var _x = p.Literal_atom()

			localctx.(*LiteralExpressionContext).literal = _x
		}

	case 6:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(270)

			var _x = p.Function_expression()

			localctx.(*FunctionExpressionContext).exp = _x
		}

	case 7:
		localctx = NewStructExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(271)

			var _x = p.Struct_expression()

			localctx.(*StructExpressionContext).exp = _x
		}

	case 8:
		localctx = NewTupleExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)

			var _x = p.Tuple_expression()

			localctx.(*TupleExpressionContext).exp = _x
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(323)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(276)

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
					p.SetState(277)

					var _x = p.expression(19)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(279)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(280)

					var _x = p.expression(17)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(282)

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
					p.SetState(283)

					var _x = p.expression(16)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(285)

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
					p.SetState(286)

					var _x = p.expression(15)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 5:
				localctx = NewBooleanBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BooleanBinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(288)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BooleanBinaryExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BooleanBinaryExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(289)

					var _x = p.expression(14)

					localctx.(*BooleanBinaryExpressionContext).rhs_exp = _x
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(291)
					p.Match(CaliburnParserOP_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(292)

					var _x = p.expression(13)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(294)
					p.Match(CaliburnParserOP_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(295)

					var _x = p.expression(12)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(297)
					p.Match(CaliburnParserOP_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(298)

					var _x = p.expression(11)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 9:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CallExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(300)
					p.Match(CaliburnParserL_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(302)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023286263808) != 0 {
					{
						p.SetState(301)

						var _x = p.Expressions()

						localctx.(*CallExpressionContext).args = _x
					}

				}
				{
					p.SetState(304)
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
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(306)
					p.Match(CaliburnParserPERIOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(307)

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
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(309)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(310)

					var _x = p.expression(0)

					localctx.(*IndexExpressionContext).index = _x
				}
				{
					p.SetState(311)
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
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(314)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(316)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023286263808) != 0 {
					{
						p.SetState(315)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).start_index = _x
					}

				}
				{
					p.SetState(318)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(320)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023286263808) != 0 {
					{
						p.SetState(319)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).end_index = _x
					}

				}
				{
					p.SetState(322)
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
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext())
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

// IFunction_expressionContext is an interface to support dynamic dispatch.
type IFunction_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ rule contexts.
	GetType_() IType_expressionContext

	// GetReturn_type returns the return_type rule contexts.
	GetReturn_type() IType_expressionContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IType_expressionContext)

	// SetReturn_type sets the return_type rule contexts.
	SetReturn_type(IType_expressionContext)

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Scoped_block() IScoped_blockContext
	AllType_expression() []IType_expressionContext
	Type_expression(i int) IType_expressionContext
	Assign_declarations() IAssign_declarationsContext

	// IsFunction_expressionContext differentiates from other interfaces.
	IsFunction_expressionContext()
}

type Function_expressionContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	type_       IType_expressionContext
	return_type IType_expressionContext
}

func NewEmptyFunction_expressionContext() *Function_expressionContext {
	var p = new(Function_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_expression
	return p
}

func InitEmptyFunction_expressionContext(p *Function_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_function_expression
}

func (*Function_expressionContext) IsFunction_expressionContext() {}

func NewFunction_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_expressionContext {
	var p = new(Function_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_function_expression

	return p
}

func (s *Function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_expressionContext) GetType_() IType_expressionContext { return s.type_ }

func (s *Function_expressionContext) GetReturn_type() IType_expressionContext { return s.return_type }

func (s *Function_expressionContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *Function_expressionContext) SetReturn_type(v IType_expressionContext) { s.return_type = v }

func (s *Function_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *Function_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *Function_expressionContext) Scoped_block() IScoped_blockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScoped_blockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScoped_blockContext)
}

func (s *Function_expressionContext) AllType_expression() []IType_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_expressionContext); ok {
			len++
		}
	}

	tst := make([]IType_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_expressionContext); ok {
			tst[i] = t.(IType_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Function_expressionContext) Type_expression(i int) IType_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_expressionContext); ok {
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

	return t.(IType_expressionContext)
}

func (s *Function_expressionContext) Assign_declarations() IAssign_declarationsContext {
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

func (s *Function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterFunction_expression(s)
	}
}

func (s *Function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitFunction_expression(s)
	}
}

func (s *Function_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitFunction_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Function_expression() (localctx IFunction_expressionContext) {
	localctx = NewFunction_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CaliburnParserRULE_function_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)

		var _x = p.Type_expression()

		localctx.(*Function_expressionContext).type_ = _x
	}
	{
		p.SetState(329)
		p.Match(CaliburnParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057605431754752) != 0 {
		{
			p.SetState(330)
			p.Assign_declarations()
		}

	}
	{
		p.SetState(333)
		p.Match(CaliburnParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057594155368448) != 0 {
		{
			p.SetState(334)

			var _x = p.Type_expression()

			localctx.(*Function_expressionContext).return_type = _x
		}

	}
	{
		p.SetState(337)
		p.Scoped_block()
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

// IStruct_expressionContext is an interface to support dynamic dispatch.
type IStruct_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ rule contexts.
	GetType_() IType_expressionContext

	// SetType_ sets the type_ rule contexts.
	SetType_(IType_expressionContext)

	// Getter signatures
	L_C_BRACK() antlr.TerminalNode
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_C_BRACK() antlr.TerminalNode
	Type_expression() IType_expressionContext
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsStruct_expressionContext differentiates from other interfaces.
	IsStruct_expressionContext()
}

type Struct_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  IType_expressionContext
}

func NewEmptyStruct_expressionContext() *Struct_expressionContext {
	var p = new(Struct_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_expression
	return p
}

func InitEmptyStruct_expressionContext(p *Struct_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_struct_expression
}

func (*Struct_expressionContext) IsStruct_expressionContext() {}

func NewStruct_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_expressionContext {
	var p = new(Struct_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_struct_expression

	return p
}

func (s *Struct_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_expressionContext) GetType_() IType_expressionContext { return s.type_ }

func (s *Struct_expressionContext) SetType_(v IType_expressionContext) { s.type_ = v }

func (s *Struct_expressionContext) L_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_C_BRACK, 0)
}

func (s *Struct_expressionContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOLON)
}

func (s *Struct_expressionContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOLON, i)
}

func (s *Struct_expressionContext) AllExpression() []IExpressionContext {
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

func (s *Struct_expressionContext) Expression(i int) IExpressionContext {
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

func (s *Struct_expressionContext) R_C_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_C_BRACK, 0)
}

func (s *Struct_expressionContext) Type_expression() IType_expressionContext {
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

func (s *Struct_expressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserIdentifier)
}

func (s *Struct_expressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, i)
}

func (s *Struct_expressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Struct_expressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Struct_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterStruct_expression(s)
	}
}

func (s *Struct_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitStruct_expression(s)
	}
}

func (s *Struct_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitStruct_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Struct_expression() (localctx IStruct_expressionContext) {
	localctx = NewStruct_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CaliburnParserRULE_struct_expression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)

		var _x = p.Type_expression()

		localctx.(*Struct_expressionContext).type_ = _x
	}
	{
		p.SetState(340)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(341)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
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
		p.expression(0)
	}
	p.SetState(350)
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
			{
				p.SetState(344)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(345)
				p.Match(CaliburnParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			{
				p.SetState(346)
				p.Match(CaliburnParserCOLON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(347)
				p.expression(0)
			}

		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserCOMMA {
		{
			p.SetState(353)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(356)
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

// ITuple_expressionContext is an interface to support dynamic dispatch.
type ITuple_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_PAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTuple_expressionContext differentiates from other interfaces.
	IsTuple_expressionContext()
}

type Tuple_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuple_expressionContext() *Tuple_expressionContext {
	var p = new(Tuple_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_tuple_expression
	return p
}

func InitEmptyTuple_expressionContext(p *Tuple_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_tuple_expression
}

func (*Tuple_expressionContext) IsTuple_expressionContext() {}

func NewTuple_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tuple_expressionContext {
	var p = new(Tuple_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_tuple_expression

	return p
}

func (s *Tuple_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Tuple_expressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_PAREN, 0)
}

func (s *Tuple_expressionContext) AllExpression() []IExpressionContext {
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

func (s *Tuple_expressionContext) Expression(i int) IExpressionContext {
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

func (s *Tuple_expressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
}

func (s *Tuple_expressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Tuple_expressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Tuple_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tuple_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tuple_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTuple_expression(s)
	}
}

func (s *Tuple_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTuple_expression(s)
	}
}

func (s *Tuple_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTuple_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Tuple_expression() (localctx ITuple_expressionContext) {
	localctx = NewTuple_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CaliburnParserRULE_tuple_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(CaliburnParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(359)
		p.expression(0)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(360)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == CaliburnParserCOMMA {
			{
				p.SetState(361)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(362)
				p.expression(0)
			}

			p.SetState(365)
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
		p.SetState(369)
		p.Match(CaliburnParserR_PAREN)
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

// IType_expressionContext is an interface to support dynamic dispatch.
type IType_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Complex_type_expression() IComplex_type_expressionContext
	FUNC() antlr.TerminalNode
	STRUCT() antlr.TerminalNode
	TUPLE() antlr.TerminalNode

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

func (s *Type_expressionContext) Complex_type_expression() IComplex_type_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComplex_type_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComplex_type_expressionContext)
}

func (s *Type_expressionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(CaliburnParserFUNC, 0)
}

func (s *Type_expressionContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserSTRUCT, 0)
}

func (s *Type_expressionContext) TUPLE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserTUPLE, 0)
}

func (s *Type_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterType_expression(s)
	}
}

func (s *Type_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitType_expression(s)
	}
}

func (s *Type_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitType_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Type_expression() (localctx IType_expressionContext) {
	localctx = NewType_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CaliburnParserRULE_type_expression)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.complex_type_expression(0)
		}

	case CaliburnParserFUNC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Match(CaliburnParserFUNC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserSTRUCT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Match(CaliburnParserSTRUCT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserTUPLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Match(CaliburnParserTUPLE)
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

// IComplex_type_expressionContext is an interface to support dynamic dispatch.
type IComplex_type_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Complex_type_expression() IComplex_type_expressionContext
	PERIOD() antlr.TerminalNode

	// IsComplex_type_expressionContext differentiates from other interfaces.
	IsComplex_type_expressionContext()
}

type Complex_type_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComplex_type_expressionContext() *Complex_type_expressionContext {
	var p = new(Complex_type_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_complex_type_expression
	return p
}

func InitEmptyComplex_type_expressionContext(p *Complex_type_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_complex_type_expression
}

func (*Complex_type_expressionContext) IsComplex_type_expressionContext() {}

func NewComplex_type_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Complex_type_expressionContext {
	var p = new(Complex_type_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_complex_type_expression

	return p
}

func (s *Complex_type_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Complex_type_expressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Complex_type_expressionContext) Complex_type_expression() IComplex_type_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComplex_type_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComplex_type_expressionContext)
}

func (s *Complex_type_expressionContext) PERIOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserPERIOD, 0)
}

func (s *Complex_type_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Complex_type_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Complex_type_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterComplex_type_expression(s)
	}
}

func (s *Complex_type_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitComplex_type_expression(s)
	}
}

func (s *Complex_type_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitComplex_type_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Complex_type_expression() (localctx IComplex_type_expressionContext) {
	return p.complex_type_expression(0)
}

func (p *CaliburnParser) complex_type_expression(_p int) (localctx IComplex_type_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewComplex_type_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IComplex_type_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, CaliburnParserRULE_complex_type_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(385)
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
			localctx = NewComplex_type_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_complex_type_expression)
			p.SetState(380)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(381)
				p.Match(CaliburnParserPERIOD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(382)
				p.Match(CaliburnParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(387)
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

// IIdentifiersContext is an interface to support dynamic dispatch.
type IIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdentifier() []antlr.TerminalNode
	Identifier(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

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

func (s *IdentifiersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserIdentifier)
}

func (s *IdentifiersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, i)
}

func (s *IdentifiersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *IdentifiersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *IdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterIdentifiers(s)
	}
}

func (s *IdentifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitIdentifiers(s)
	}
}

func (s *IdentifiersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitIdentifiers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Identifiers() (localctx IIdentifiersContext) {
	localctx = NewIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CaliburnParserRULE_identifiers)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(389)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(395)
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

type UntypedLiteralAtomContext struct {
	Literal_atomContext
}

func NewUntypedLiteralAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UntypedLiteralAtomContext {
	var p = new(UntypedLiteralAtomContext)

	InitEmptyLiteral_atomContext(&p.Literal_atomContext)
	p.parser = parser
	p.CopyAll(ctx.(*Literal_atomContext))

	return p
}

func (s *UntypedLiteralAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntypedLiteralAtomContext) Untyped_literal_atom() IUntyped_literal_atomContext {
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

func (s *UntypedLiteralAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntypedLiteralAtom(s)
	}
}

func (s *UntypedLiteralAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntypedLiteralAtom(s)
	}
}

func (s *UntypedLiteralAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntypedLiteralAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypedLiteralAtomContext struct {
	Literal_atomContext
}

func NewTypedLiteralAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedLiteralAtomContext {
	var p = new(TypedLiteralAtomContext)

	InitEmptyLiteral_atomContext(&p.Literal_atomContext)
	p.parser = parser
	p.CopyAll(ctx.(*Literal_atomContext))

	return p
}

func (s *TypedLiteralAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedLiteralAtomContext) Typed_literal_atom() ITyped_literal_atomContext {
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

func (s *TypedLiteralAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTypedLiteralAtom(s)
	}
}

func (s *TypedLiteralAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTypedLiteralAtom(s)
	}
}

func (s *TypedLiteralAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTypedLiteralAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Literal_atom() (localctx ILiteral_atomContext) {
	localctx = NewLiteral_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CaliburnParserRULE_literal_atom)
	p.SetState(398)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserLiteral:
		localctx = NewUntypedLiteralAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(396)
			p.Untyped_literal_atom()
		}

	case CaliburnParserIdentifier:
		localctx = NewTypedLiteralAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Typed_literal_atom()
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
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterUntyped_literal_atom(s)
	}
}

func (s *Untyped_literal_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitUntyped_literal_atom(s)
	}
}

func (s *Untyped_literal_atomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitUntyped_literal_atom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Untyped_literal_atom() (localctx IUntyped_literal_atomContext) {
	localctx = NewUntyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CaliburnParserRULE_untyped_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(400)
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
	Identifier() antlr.TerminalNode
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

func (s *Typed_literal_atomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
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
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterTyped_literal_atom(s)
	}
}

func (s *Typed_literal_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitTyped_literal_atom(s)
	}
}

func (s *Typed_literal_atomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitTyped_literal_atom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Typed_literal_atom() (localctx ITyped_literal_atomContext) {
	localctx = NewTyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CaliburnParserRULE_typed_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(402)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(403)
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

func (p *CaliburnParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 24:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 29:
		var t *Complex_type_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Complex_type_expressionContext)
		}
		return p.Complex_type_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CaliburnParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CaliburnParser) Complex_type_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
		"OP_GTE", "OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier",
		"Literal",
	}
	staticData.RuleNames = []string{
		"module", "definition", "function_definition", "parameters", "parameter",
		"block", "statement", "jump_statement", "return_statement", "break_statement",
		"continue_statement", "control_statement", "if_statement", "for_statement",
		"switch_statement", "case_blocks", "option_case_block", "default_case_block",
		"inline_statement", "assign_statement", "assign_expressions", "assign_expression",
		"assign_declarations", "assign_declaration", "typed_assign_declarations",
		"typed_assign_declaration", "untyped_assign_declarations", "untyped_assign_declaration",
		"expression_statement", "expressions", "expression", "function_type",
		"struct_type", "tuple_type", "type_expression", "identifiers", "literal_atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 458, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 3, 2, 87, 8, 2, 1, 2, 1, 2, 3, 2, 91, 8, 2, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 5, 3, 98, 8, 3, 10, 3, 12, 3, 101, 9, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 113, 8, 4, 1, 5, 1, 5,
		5, 5, 117, 8, 5, 10, 5, 12, 5, 120, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 6, 3, 6, 128, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 133, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3,
		11, 148, 8, 11, 1, 12, 1, 12, 5, 12, 152, 8, 12, 10, 12, 12, 12, 155, 9,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 162, 8, 12, 3, 12, 164, 8,
		12, 1, 13, 1, 13, 4, 13, 168, 8, 13, 11, 13, 12, 13, 169, 1, 13, 1, 13,
		1, 13, 4, 13, 175, 8, 13, 11, 13, 12, 13, 176, 3, 13, 179, 8, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 5, 14, 185, 8, 14, 10, 14, 12, 14, 188, 9, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 5, 15, 196, 8, 15, 10, 15, 12, 15,
		199, 9, 15, 1, 15, 3, 15, 202, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17,
		1, 17, 1, 17, 1, 18, 1, 18, 3, 18, 213, 8, 18, 1, 19, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 226, 8, 19,
		1, 20, 1, 20, 1, 20, 5, 20, 231, 8, 20, 10, 20, 12, 20, 234, 9, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 245,
		8, 21, 1, 22, 1, 22, 1, 22, 5, 22, 250, 8, 22, 10, 22, 12, 22, 253, 9,
		22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 267, 8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 272, 8, 24,
		10, 24, 12, 24, 275, 9, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 5,
		26, 283, 8, 26, 10, 26, 12, 26, 286, 9, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 297, 8, 27, 1, 28, 1, 28, 1,
		28, 1, 29, 1, 29, 1, 29, 5, 29, 305, 8, 29, 10, 29, 12, 29, 308, 9, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 327, 8, 30, 1, 30,
		1, 30, 3, 30, 331, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 5, 30, 344, 8, 30, 10, 30, 12, 30, 347,
		9, 30, 1, 30, 3, 30, 350, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 4, 30, 359, 8, 30, 11, 30, 12, 30, 360, 3, 30, 363, 8, 30, 1,
		30, 1, 30, 3, 30, 367, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		3, 30, 396, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 410, 8, 30, 1, 30, 1, 30, 3, 30,
		414, 8, 30, 1, 30, 5, 30, 417, 8, 30, 10, 30, 12, 30, 420, 9, 30, 1, 31,
		1, 31, 3, 31, 424, 8, 31, 1, 32, 1, 32, 3, 32, 428, 8, 32, 1, 33, 1, 33,
		3, 33, 432, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 440,
		8, 34, 10, 34, 12, 34, 443, 9, 34, 1, 35, 1, 35, 1, 35, 5, 35, 448, 8,
		35, 10, 35, 12, 35, 451, 9, 35, 1, 36, 1, 36, 1, 36, 3, 36, 456, 8, 36,
		1, 36, 0, 2, 60, 68, 37, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 0, 7, 2, 0, 36, 37, 39, 48, 1, 0, 36, 38, 1, 0,
		42, 43, 1, 0, 39, 41, 1, 0, 36, 37, 1, 0, 47, 48, 1, 0, 49, 54, 493, 0,
		77, 1, 0, 0, 0, 2, 80, 1, 0, 0, 0, 4, 82, 1, 0, 0, 0, 6, 94, 1, 0, 0, 0,
		8, 112, 1, 0, 0, 0, 10, 114, 1, 0, 0, 0, 12, 127, 1, 0, 0, 0, 14, 132,
		1, 0, 0, 0, 16, 134, 1, 0, 0, 0, 18, 138, 1, 0, 0, 0, 20, 141, 1, 0, 0,
		0, 22, 147, 1, 0, 0, 0, 24, 149, 1, 0, 0, 0, 26, 165, 1, 0, 0, 0, 28, 182,
		1, 0, 0, 0, 30, 197, 1, 0, 0, 0, 32, 203, 1, 0, 0, 0, 34, 207, 1, 0, 0,
		0, 36, 212, 1, 0, 0, 0, 38, 225, 1, 0, 0, 0, 40, 227, 1, 0, 0, 0, 42, 244,
		1, 0, 0, 0, 44, 246, 1, 0, 0, 0, 46, 266, 1, 0, 0, 0, 48, 268, 1, 0, 0,
		0, 50, 276, 1, 0, 0, 0, 52, 279, 1, 0, 0, 0, 54, 296, 1, 0, 0, 0, 56, 298,
		1, 0, 0, 0, 58, 301, 1, 0, 0, 0, 60, 366, 1, 0, 0, 0, 62, 423, 1, 0, 0,
		0, 64, 427, 1, 0, 0, 0, 66, 431, 1, 0, 0, 0, 68, 433, 1, 0, 0, 0, 70, 444,
		1, 0, 0, 0, 72, 455, 1, 0, 0, 0, 74, 76, 3, 2, 1, 0, 75, 74, 1, 0, 0, 0,
		76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 1, 1, 0,
		0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 3, 4, 2, 0, 81, 3, 1, 0, 0, 0, 82, 83,
		3, 62, 31, 0, 83, 84, 5, 57, 0, 0, 84, 86, 5, 30, 0, 0, 85, 87, 3, 6, 3,
		0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 90,
		5, 31, 0, 0, 89, 91, 3, 68, 34, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0,
		0, 91, 92, 1, 0, 0, 0, 92, 93, 3, 10, 5, 0, 93, 5, 1, 0, 0, 0, 94, 99,
		3, 8, 4, 0, 95, 96, 5, 2, 0, 0, 96, 98, 3, 8, 4, 0, 97, 95, 1, 0, 0, 0,
		98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 7, 1,
		0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 113, 3, 50, 25, 0, 103, 113, 3, 54,
		27, 0, 104, 105, 5, 34, 0, 0, 105, 106, 3, 6, 3, 0, 106, 107, 5, 35, 0,
		0, 107, 113, 1, 0, 0, 0, 108, 109, 5, 32, 0, 0, 109, 110, 3, 6, 3, 0, 110,
		111, 5, 33, 0, 0, 111, 113, 1, 0, 0, 0, 112, 102, 1, 0, 0, 0, 112, 103,
		1, 0, 0, 0, 112, 104, 1, 0, 0, 0, 112, 108, 1, 0, 0, 0, 113, 9, 1, 0, 0,
		0, 114, 118, 5, 34, 0, 0, 115, 117, 3, 12, 6, 0, 116, 115, 1, 0, 0, 0,
		117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119,
		121, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 122, 5, 35, 0, 0, 122, 11,
		1, 0, 0, 0, 123, 128, 3, 38, 19, 0, 124, 128, 3, 56, 28, 0, 125, 128, 3,
		22, 11, 0, 126, 128, 3, 14, 7, 0, 127, 123, 1, 0, 0, 0, 127, 124, 1, 0,
		0, 0, 127, 125, 1, 0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 13, 1, 0, 0, 0,
		129, 133, 3, 16, 8, 0, 130, 133, 3, 18, 9, 0, 131, 133, 3, 20, 10, 0, 132,
		129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 15, 1,
		0, 0, 0, 134, 135, 5, 17, 0, 0, 135, 136, 3, 60, 30, 0, 136, 137, 5, 5,
		0, 0, 137, 17, 1, 0, 0, 0, 138, 139, 5, 15, 0, 0, 139, 140, 5, 5, 0, 0,
		140, 19, 1, 0, 0, 0, 141, 142, 5, 16, 0, 0, 142, 143, 5, 5, 0, 0, 143,
		21, 1, 0, 0, 0, 144, 148, 3, 24, 12, 0, 145, 148, 3, 26, 13, 0, 146, 148,
		3, 28, 14, 0, 147, 144, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 146, 1,
		0, 0, 0, 148, 23, 1, 0, 0, 0, 149, 153, 5, 8, 0, 0, 150, 152, 3, 36, 18,
		0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157,
		3, 60, 30, 0, 157, 163, 3, 10, 5, 0, 158, 161, 5, 9, 0, 0, 159, 162, 3,
		10, 5, 0, 160, 162, 3, 24, 12, 0, 161, 159, 1, 0, 0, 0, 161, 160, 1, 0,
		0, 0, 162, 164, 1, 0, 0, 0, 163, 158, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0,
		164, 25, 1, 0, 0, 0, 165, 167, 5, 10, 0, 0, 166, 168, 3, 36, 18, 0, 167,
		166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170,
		1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 178, 3, 60, 30, 0, 172, 174, 5,
		7, 0, 0, 173, 175, 3, 36, 18, 0, 174, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 179, 1, 0, 0, 0,
		178, 172, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180,
		181, 3, 10, 5, 0, 181, 27, 1, 0, 0, 0, 182, 186, 5, 11, 0, 0, 183, 185,
		3, 36, 18, 0, 184, 183, 1, 0, 0, 0, 185, 188, 1, 0, 0, 0, 186, 184, 1,
		0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 189, 1, 0, 0, 0, 188, 186, 1, 0, 0,
		0, 189, 190, 3, 60, 30, 0, 190, 191, 5, 34, 0, 0, 191, 192, 3, 30, 15,
		0, 192, 193, 5, 35, 0, 0, 193, 29, 1, 0, 0, 0, 194, 196, 3, 32, 16, 0,
		195, 194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 197,
		198, 1, 0, 0, 0, 198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 202,
		3, 34, 17, 0, 201, 200, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 31, 1, 0,
		0, 0, 203, 204, 5, 12, 0, 0, 204, 205, 3, 60, 30, 0, 205, 206, 3, 10, 5,
		0, 206, 33, 1, 0, 0, 0, 207, 208, 5, 13, 0, 0, 208, 209, 3, 10, 5, 0, 209,
		35, 1, 0, 0, 0, 210, 213, 3, 38, 19, 0, 211, 213, 3, 56, 28, 0, 212, 210,
		1, 0, 0, 0, 212, 211, 1, 0, 0, 0, 213, 37, 1, 0, 0, 0, 214, 215, 3, 44,
		22, 0, 215, 216, 5, 6, 0, 0, 216, 217, 3, 58, 29, 0, 217, 218, 5, 5, 0,
		0, 218, 226, 1, 0, 0, 0, 219, 220, 3, 40, 20, 0, 220, 221, 7, 0, 0, 0,
		221, 222, 5, 6, 0, 0, 222, 223, 3, 58, 29, 0, 223, 224, 5, 5, 0, 0, 224,
		226, 1, 0, 0, 0, 225, 214, 1, 0, 0, 0, 225, 219, 1, 0, 0, 0, 226, 39, 1,
		0, 0, 0, 227, 232, 3, 42, 21, 0, 228, 229, 5, 2, 0, 0, 229, 231, 3, 42,
		21, 0, 230, 228, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0,
		232, 233, 1, 0, 0, 0, 233, 41, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 245,
		3, 60, 30, 0, 236, 237, 5, 34, 0, 0, 237, 238, 3, 40, 20, 0, 238, 239,
		5, 35, 0, 0, 239, 245, 1, 0, 0, 0, 240, 241, 5, 32, 0, 0, 241, 242, 3,
		40, 20, 0, 242, 243, 5, 33, 0, 0, 243, 245, 1, 0, 0, 0, 244, 235, 1, 0,
		0, 0, 244, 236, 1, 0, 0, 0, 244, 240, 1, 0, 0, 0, 245, 43, 1, 0, 0, 0,
		246, 251, 3, 46, 23, 0, 247, 248, 5, 2, 0, 0, 248, 250, 3, 38, 19, 0, 249,
		247, 1, 0, 0, 0, 250, 253, 1, 0, 0, 0, 251, 249, 1, 0, 0, 0, 251, 252,
		1, 0, 0, 0, 252, 45, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 267, 3, 60,
		30, 0, 255, 267, 3, 50, 25, 0, 256, 257, 5, 22, 0, 0, 257, 267, 3, 54,
		27, 0, 258, 259, 5, 34, 0, 0, 259, 260, 3, 44, 22, 0, 260, 261, 5, 35,
		0, 0, 261, 267, 1, 0, 0, 0, 262, 263, 5, 32, 0, 0, 263, 264, 3, 44, 22,
		0, 264, 265, 5, 33, 0, 0, 265, 267, 1, 0, 0, 0, 266, 254, 1, 0, 0, 0, 266,
		255, 1, 0, 0, 0, 266, 256, 1, 0, 0, 0, 266, 258, 1, 0, 0, 0, 266, 262,
		1, 0, 0, 0, 267, 47, 1, 0, 0, 0, 268, 273, 3, 50, 25, 0, 269, 270, 5, 2,
		0, 0, 270, 272, 3, 50, 25, 0, 271, 269, 1, 0, 0, 0, 272, 275, 1, 0, 0,
		0, 273, 271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 49, 1, 0, 0, 0, 275,
		273, 1, 0, 0, 0, 276, 277, 3, 68, 34, 0, 277, 278, 3, 54, 27, 0, 278, 51,
		1, 0, 0, 0, 279, 284, 3, 54, 27, 0, 280, 281, 5, 2, 0, 0, 281, 283, 3,
		54, 27, 0, 282, 280, 1, 0, 0, 0, 283, 286, 1, 0, 0, 0, 284, 282, 1, 0,
		0, 0, 284, 285, 1, 0, 0, 0, 285, 53, 1, 0, 0, 0, 286, 284, 1, 0, 0, 0,
		287, 297, 5, 57, 0, 0, 288, 289, 5, 34, 0, 0, 289, 290, 3, 52, 26, 0, 290,
		291, 5, 35, 0, 0, 291, 297, 1, 0, 0, 0, 292, 293, 5, 32, 0, 0, 293, 294,
		3, 52, 26, 0, 294, 295, 5, 33, 0, 0, 295, 297, 1, 0, 0, 0, 296, 287, 1,
		0, 0, 0, 296, 288, 1, 0, 0, 0, 296, 292, 1, 0, 0, 0, 297, 55, 1, 0, 0,
		0, 298, 299, 3, 58, 29, 0, 299, 300, 5, 5, 0, 0, 300, 57, 1, 0, 0, 0, 301,
		306, 3, 60, 30, 0, 302, 303, 5, 2, 0, 0, 303, 305, 3, 60, 30, 0, 304, 302,
		1, 0, 0, 0, 305, 308, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 306, 307, 1, 0,
		0, 0, 307, 59, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 309, 310, 6, 30, -1, 0,
		310, 311, 5, 30, 0, 0, 311, 312, 3, 68, 34, 0, 312, 313, 5, 31, 0, 0, 313,
		314, 3, 60, 30, 20, 314, 367, 1, 0, 0, 0, 315, 316, 5, 30, 0, 0, 316, 317,
		3, 60, 30, 0, 317, 318, 5, 31, 0, 0, 318, 367, 1, 0, 0, 0, 319, 320, 7,
		1, 0, 0, 320, 367, 3, 60, 30, 17, 321, 367, 5, 57, 0, 0, 322, 367, 3, 72,
		36, 0, 323, 324, 3, 62, 31, 0, 324, 326, 5, 30, 0, 0, 325, 327, 3, 44,
		22, 0, 326, 325, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 328, 1, 0, 0, 0,
		328, 330, 5, 31, 0, 0, 329, 331, 3, 68, 34, 0, 330, 329, 1, 0, 0, 0, 330,
		331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333, 3, 10, 5, 0, 333, 367,
		1, 0, 0, 0, 334, 335, 3, 64, 32, 0, 335, 336, 5, 34, 0, 0, 336, 337, 5,
		57, 0, 0, 337, 338, 5, 3, 0, 0, 338, 345, 3, 60, 30, 0, 339, 340, 5, 2,
		0, 0, 340, 341, 5, 57, 0, 0, 341, 342, 5, 3, 0, 0, 342, 344, 3, 60, 30,
		0, 343, 339, 1, 0, 0, 0, 344, 347, 1, 0, 0, 0, 345, 343, 1, 0, 0, 0, 345,
		346, 1, 0, 0, 0, 346, 349, 1, 0, 0, 0, 347, 345, 1, 0, 0, 0, 348, 350,
		5, 2, 0, 0, 349, 348, 1, 0, 0, 0, 349, 350, 1, 0, 0, 0, 350, 351, 1, 0,
		0, 0, 351, 352, 5, 35, 0, 0, 352, 367, 1, 0, 0, 0, 353, 354, 5, 30, 0,
		0, 354, 362, 3, 60, 30, 0, 355, 363, 5, 2, 0, 0, 356, 357, 5, 2, 0, 0,
		357, 359, 3, 60, 30, 0, 358, 356, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360,
		358, 1, 0, 0, 0, 360, 361, 1, 0, 0, 0, 361, 363, 1, 0, 0, 0, 362, 355,
		1, 0, 0, 0, 362, 358, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 5, 31,
		0, 0, 365, 367, 1, 0, 0, 0, 366, 309, 1, 0, 0, 0, 366, 315, 1, 0, 0, 0,
		366, 319, 1, 0, 0, 0, 366, 321, 1, 0, 0, 0, 366, 322, 1, 0, 0, 0, 366,
		323, 1, 0, 0, 0, 366, 334, 1, 0, 0, 0, 366, 353, 1, 0, 0, 0, 367, 418,
		1, 0, 0, 0, 368, 369, 10, 18, 0, 0, 369, 370, 7, 2, 0, 0, 370, 417, 3,
		60, 30, 19, 371, 372, 10, 16, 0, 0, 372, 373, 7, 3, 0, 0, 373, 417, 3,
		60, 30, 17, 374, 375, 10, 15, 0, 0, 375, 376, 7, 4, 0, 0, 376, 417, 3,
		60, 30, 16, 377, 378, 10, 14, 0, 0, 378, 379, 7, 5, 0, 0, 379, 417, 3,
		60, 30, 15, 380, 381, 10, 13, 0, 0, 381, 382, 7, 6, 0, 0, 382, 417, 3,
		60, 30, 14, 383, 384, 10, 12, 0, 0, 384, 385, 5, 46, 0, 0, 385, 417, 3,
		60, 30, 13, 386, 387, 10, 11, 0, 0, 387, 388, 5, 45, 0, 0, 388, 417, 3,
		60, 30, 12, 389, 390, 10, 10, 0, 0, 390, 391, 5, 44, 0, 0, 391, 417, 3,
		60, 30, 11, 392, 393, 10, 9, 0, 0, 393, 395, 5, 30, 0, 0, 394, 396, 3,
		58, 29, 0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0,
		0, 0, 397, 417, 5, 31, 0, 0, 398, 399, 10, 8, 0, 0, 399, 400, 5, 1, 0,
		0, 400, 417, 5, 57, 0, 0, 401, 402, 10, 7, 0, 0, 402, 403, 5, 32, 0, 0,
		403, 404, 3, 60, 30, 0, 404, 405, 5, 33, 0, 0, 405, 417, 1, 0, 0, 0, 406,
		407, 10, 6, 0, 0, 407, 409, 5, 32, 0, 0, 408, 410, 3, 60, 30, 0, 409, 408,
		1, 0, 0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 1, 0, 0, 0, 411, 413, 5, 3,
		0, 0, 412, 414, 3, 60, 30, 0, 413, 412, 1, 0, 0, 0, 413, 414, 1, 0, 0,
		0, 414, 415, 1, 0, 0, 0, 415, 417, 5, 33, 0, 0, 416, 368, 1, 0, 0, 0, 416,
		371, 1, 0, 0, 0, 416, 374, 1, 0, 0, 0, 416, 377, 1, 0, 0, 0, 416, 380,
		1, 0, 0, 0, 416, 383, 1, 0, 0, 0, 416, 386, 1, 0, 0, 0, 416, 389, 1, 0,
		0, 0, 416, 392, 1, 0, 0, 0, 416, 398, 1, 0, 0, 0, 416, 401, 1, 0, 0, 0,
		416, 406, 1, 0, 0, 0, 417, 420, 1, 0, 0, 0, 418, 416, 1, 0, 0, 0, 418,
		419, 1, 0, 0, 0, 419, 61, 1, 0, 0, 0, 420, 418, 1, 0, 0, 0, 421, 424, 5,
		25, 0, 0, 422, 424, 3, 68, 34, 0, 423, 421, 1, 0, 0, 0, 423, 422, 1, 0,
		0, 0, 424, 63, 1, 0, 0, 0, 425, 428, 5, 26, 0, 0, 426, 428, 3, 68, 34,
		0, 427, 425, 1, 0, 0, 0, 427, 426, 1, 0, 0, 0, 428, 65, 1, 0, 0, 0, 429,
		432, 5, 27, 0, 0, 430, 432, 3, 68, 34, 0, 431, 429, 1, 0, 0, 0, 431, 430,
		1, 0, 0, 0, 432, 67, 1, 0, 0, 0, 433, 434, 6, 34, -1, 0, 434, 435, 5, 57,
		0, 0, 435, 441, 1, 0, 0, 0, 436, 437, 10, 1, 0, 0, 437, 438, 5, 1, 0, 0,
		438, 440, 5, 57, 0, 0, 439, 436, 1, 0, 0, 0, 440, 443, 1, 0, 0, 0, 441,
		439, 1, 0, 0, 0, 441, 442, 1, 0, 0, 0, 442, 69, 1, 0, 0, 0, 443, 441, 1,
		0, 0, 0, 444, 449, 5, 57, 0, 0, 445, 446, 5, 2, 0, 0, 446, 448, 5, 57,
		0, 0, 447, 445, 1, 0, 0, 0, 448, 451, 1, 0, 0, 0, 449, 447, 1, 0, 0, 0,
		449, 450, 1, 0, 0, 0, 450, 71, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 452, 456,
		5, 58, 0, 0, 453, 454, 5, 57, 0, 0, 454, 456, 5, 58, 0, 0, 455, 452, 1,
		0, 0, 0, 455, 453, 1, 0, 0, 0, 456, 73, 1, 0, 0, 0, 46, 77, 86, 90, 99,
		112, 118, 127, 132, 147, 153, 161, 163, 169, 176, 178, 186, 197, 201, 212,
		225, 232, 244, 251, 266, 273, 284, 296, 306, 326, 330, 345, 349, 360, 362,
		366, 395, 409, 413, 416, 418, 423, 427, 431, 441, 449, 455,
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
	CaliburnParserARROW       = 7
	CaliburnParserIF          = 8
	CaliburnParserELSE        = 9
	CaliburnParserFOR         = 10
	CaliburnParserSWITCH      = 11
	CaliburnParserCASE        = 12
	CaliburnParserDEFAULT     = 13
	CaliburnParserFALLTHROUGH = 14
	CaliburnParserBREAK       = 15
	CaliburnParserCONTINUE    = 16
	CaliburnParserRETURN      = 17
	CaliburnParserUSING       = 18
	CaliburnParserAS          = 19
	CaliburnParserIN          = 20
	CaliburnParserNULL        = 21
	CaliburnParserVAR         = 22
	CaliburnParserCONST       = 23
	CaliburnParserTYPE        = 24
	CaliburnParserFUNC        = 25
	CaliburnParserSTRUCT      = 26
	CaliburnParserTUPLE       = 27
	CaliburnParserCLASS       = 28
	CaliburnParserINTERFACE   = 29
	CaliburnParserL_PAREN     = 30
	CaliburnParserR_PAREN     = 31
	CaliburnParserL_S_BRACK   = 32
	CaliburnParserR_S_BRACK   = 33
	CaliburnParserL_C_BRACK   = 34
	CaliburnParserR_C_BRACK   = 35
	CaliburnParserOP_ADD      = 36
	CaliburnParserOP_SUB      = 37
	CaliburnParserOP_NOT      = 38
	CaliburnParserOP_MUL      = 39
	CaliburnParserOP_DIV      = 40
	CaliburnParserOP_MOD      = 41
	CaliburnParserOP_POW      = 42
	CaliburnParserOP_ROOT     = 43
	CaliburnParserOP_OR       = 44
	CaliburnParserOP_XOR      = 45
	CaliburnParserOP_AND      = 46
	CaliburnParserOP_LSHIFT   = 47
	CaliburnParserOP_RSHIFT   = 48
	CaliburnParserOP_EQU      = 49
	CaliburnParserOP_NEQ      = 50
	CaliburnParserOP_GTE      = 51
	CaliburnParserOP_GRT      = 52
	CaliburnParserOP_LTE      = 53
	CaliburnParserOP_LST      = 54
	CaliburnParserWhiteSpace  = 55
	CaliburnParserComment     = 56
	CaliburnParserIdentifier  = 57
	CaliburnParserLiteral     = 58
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_module                      = 0
	CaliburnParserRULE_definition                  = 1
	CaliburnParserRULE_function_definition         = 2
	CaliburnParserRULE_parameters                  = 3
	CaliburnParserRULE_parameter                   = 4
	CaliburnParserRULE_block                       = 5
	CaliburnParserRULE_statement                   = 6
	CaliburnParserRULE_jump_statement              = 7
	CaliburnParserRULE_return_statement            = 8
	CaliburnParserRULE_break_statement             = 9
	CaliburnParserRULE_continue_statement          = 10
	CaliburnParserRULE_control_statement           = 11
	CaliburnParserRULE_if_statement                = 12
	CaliburnParserRULE_for_statement               = 13
	CaliburnParserRULE_switch_statement            = 14
	CaliburnParserRULE_case_blocks                 = 15
	CaliburnParserRULE_option_case_block           = 16
	CaliburnParserRULE_default_case_block          = 17
	CaliburnParserRULE_inline_statement            = 18
	CaliburnParserRULE_assign_statement            = 19
	CaliburnParserRULE_assign_expressions          = 20
	CaliburnParserRULE_assign_expression           = 21
	CaliburnParserRULE_assign_declarations         = 22
	CaliburnParserRULE_assign_declaration          = 23
	CaliburnParserRULE_typed_assign_declarations   = 24
	CaliburnParserRULE_typed_assign_declaration    = 25
	CaliburnParserRULE_untyped_assign_declarations = 26
	CaliburnParserRULE_untyped_assign_declaration  = 27
	CaliburnParserRULE_expression_statement        = 28
	CaliburnParserRULE_expressions                 = 29
	CaliburnParserRULE_expression                  = 30
	CaliburnParserRULE_function_type               = 31
	CaliburnParserRULE_struct_type                 = 32
	CaliburnParserRULE_tuple_type                  = 33
	CaliburnParserRULE_type_expression             = 34
	CaliburnParserRULE_identifiers                 = 35
	CaliburnParserRULE_literal_atom                = 36
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

func (s *ModuleRuleContext) AllDefinition() []IDefinitionContext {
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

func (s *ModuleRuleContext) Definition(i int) IDefinitionContext {
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
	var _la int

	localctx = NewModuleRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserFUNC || _la == CaliburnParserIdentifier {
		{
			p.SetState(74)
			p.Definition()
		}

		p.SetState(79)
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
	p.EnterRule(localctx, 2, CaliburnParserRULE_definition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
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
	p.EnterRule(localctx, 4, CaliburnParserRULE_function_definition)
	var _la int

	localctx = NewFunctionDefinitionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)

		var _x = p.Function_type()

		localctx.(*FunctionDefinitionContext).type_ = _x
	}
	{
		p.SetState(83)

		var _m = p.Match(CaliburnParserIdentifier)

		localctx.(*FunctionDefinitionContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(CaliburnParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115209550692352) != 0 {
		{
			p.SetState(85)
			p.Parameters()
		}

	}
	{
		p.SetState(88)
		p.Match(CaliburnParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserIdentifier {
		{
			p.SetState(89)

			var _x = p.type_expression(0)

			localctx.(*FunctionDefinitionContext).return_type = _x
		}

	}
	{
		p.SetState(92)
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

type ParametersRuleContext struct {
	ParametersContext
}

func NewParametersRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParametersRuleContext {
	var p = new(ParametersRuleContext)

	InitEmptyParametersContext(&p.ParametersContext)
	p.parser = parser
	p.CopyAll(ctx.(*ParametersContext))

	return p
}

func (s *ParametersRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersRuleContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParametersRuleContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
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

	return t.(IParameterContext)
}

func (s *ParametersRuleContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *ParametersRuleContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *ParametersRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterParametersRule(s)
	}
}

func (s *ParametersRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitParametersRule(s)
	}
}

func (s *ParametersRuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CaliburnParserVisitor:
		return t.VisitParametersRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CaliburnParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CaliburnParserRULE_parameters)
	var _la int

	localctx = NewParametersRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Parameter()
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(95)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.Parameter()
		}

		p.SetState(101)
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
	p.EnterRule(localctx, 8, CaliburnParserRULE_parameter)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(102)
			p.Typed_assign_declaration()
		}

	case 2:
		localctx = NewUntypedParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Untyped_assign_declaration()
		}

	case 3:
		localctx = NewStructDestrutureParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(104)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Parameters()
		}
		{
			p.SetState(106)
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
			p.SetState(108)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Parameters()
		}
		{
			p.SetState(110)
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
	p.EnterRule(localctx, 10, CaliburnParserRULE_block)
	var _la int

	localctx = NewBlockRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(CaliburnParserL_C_BRACK)
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
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432346067917573376) != 0 {
		{
			p.SetState(115)
			p.Statement()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
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
	p.EnterRule(localctx, 12, CaliburnParserRULE_statement)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Expression_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Control_statement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
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
	p.EnterRule(localctx, 14, CaliburnParserRULE_jump_statement)
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserRETURN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Return_statement()
		}

	case CaliburnParserBREAK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Break_statement()
		}

	case CaliburnParserCONTINUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(131)
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
	p.EnterRule(localctx, 16, CaliburnParserRULE_return_statement)
	localctx = NewReturnStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(CaliburnParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.expression(0)
	}
	{
		p.SetState(136)
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
	p.EnterRule(localctx, 18, CaliburnParserRULE_break_statement)
	localctx = NewBreakStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(CaliburnParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
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
	p.EnterRule(localctx, 20, CaliburnParserRULE_continue_statement)
	localctx = NewContinueStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(141)
		p.Match(CaliburnParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
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
	p.EnterRule(localctx, 22, CaliburnParserRULE_control_statement)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.If_statement()
		}

	case CaliburnParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
			p.For_statement()
		}

	case CaliburnParserSWITCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(146)
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

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
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

	return t.(IBlockContext)
}

func (s *IfStatementContext) AllInline_statement() []IInline_statementContext {
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

func (s *IfStatementContext) Inline_statement(i int) IInline_statementContext {
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

func (s *IfStatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserELSE, 0)
}

func (s *IfStatementContext) If_statement() IIf_statementContext {
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

func (p *CaliburnParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CaliburnParserRULE_if_statement)
	var _la int

	var _alt int

	localctx = NewIfStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(CaliburnParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
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
				p.SetState(150)
				p.Inline_statement()
			}

		}
		p.SetState(155)
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
		p.SetState(156)
		p.expression(0)
	}
	{
		p.SetState(157)
		p.Block()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserELSE {
		{
			p.SetState(158)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CaliburnParserL_C_BRACK:
			{
				p.SetState(159)
				p.Block()
			}

		case CaliburnParserIF:
			{
				p.SetState(160)
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

func (s *ForStatementContext) AllInline_statement() []IInline_statementContext {
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

func (s *ForStatementContext) Inline_statement(i int) IInline_statementContext {
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

func (s *ForStatementContext) ARROW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserARROW, 0)
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

func (p *CaliburnParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CaliburnParserRULE_for_statement)
	var _la int

	var _alt int

	localctx = NewForStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(CaliburnParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(166)
				p.Inline_statement()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.expression(0)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserARROW {
		{
			p.SetState(172)
			p.Match(CaliburnParserARROW)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(173)
					p.Inline_statement()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	{
		p.SetState(180)
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

func (s *SwitchStatementContext) AllInline_statement() []IInline_statementContext {
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

func (s *SwitchStatementContext) Inline_statement(i int) IInline_statementContext {
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
	p.EnterRule(localctx, 28, CaliburnParserRULE_switch_statement)
	var _alt int

	localctx = NewSwitchStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.Match(CaliburnParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(186)
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
			{
				p.SetState(183)
				p.Inline_statement()
			}

		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(189)
		p.expression(0)
	}
	{
		p.SetState(190)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Case_blocks()
	}
	{
		p.SetState(192)
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
	p.EnterRule(localctx, 30, CaliburnParserRULE_case_blocks)
	var _la int

	localctx = NewCaseBlocksContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCASE {
		{
			p.SetState(194)
			p.Option_case_block()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserDEFAULT {
		{
			p.SetState(200)
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
	p.EnterRule(localctx, 32, CaliburnParserRULE_option_case_block)
	localctx = NewOptionCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(CaliburnParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.expression(0)
	}
	{
		p.SetState(205)
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
	p.EnterRule(localctx, 34, CaliburnParserRULE_default_case_block)
	localctx = NewDefaultCaseBlockContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(CaliburnParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
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
	p.EnterRule(localctx, 36, CaliburnParserRULE_inline_statement)
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(211)
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
	p.EnterRule(localctx, 38, CaliburnParserRULE_assign_statement)
	var _la int

	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(214)
			p.Assign_declarations()
		}
		{
			p.SetState(215)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.Expressions()
		}
		{
			p.SetState(217)
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
			p.SetState(219)
			p.Assign_expressions()
		}
		{
			p.SetState(220)

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
			p.SetState(221)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Expressions()
		}
		{
			p.SetState(223)
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
	p.EnterRule(localctx, 40, CaliburnParserRULE_assign_expressions)
	var _la int

	localctx = NewAssignExpressionsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Assign_expression()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(228)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Assign_expression()
		}

		p.SetState(234)
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
	p.EnterRule(localctx, 42, CaliburnParserRULE_assign_expression)
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC, CaliburnParserSTRUCT, CaliburnParserL_PAREN, CaliburnParserOP_ADD, CaliburnParserOP_SUB, CaliburnParserOP_NOT, CaliburnParserIdentifier, CaliburnParserLiteral:
		localctx = NewExpressionAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.expression(0)
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewStructDestrutureAssignExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Assign_expressions()
		}
		{
			p.SetState(238)
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
			p.SetState(240)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(241)
			p.Assign_expressions()
		}
		{
			p.SetState(242)
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
	p.EnterRule(localctx, 44, CaliburnParserRULE_assign_declarations)
	var _la int

	localctx = NewAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Assign_declaration()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(247)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(248)
			p.Assign_statement()
		}

		p.SetState(253)
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
	p.EnterRule(localctx, 46, CaliburnParserRULE_assign_declaration)
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpressionAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(254)
			p.expression(0)
		}

	case 2:
		localctx = NewTypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(255)
			p.Typed_assign_declaration()
		}

	case 3:
		localctx = NewUntypedAssignDeclarationDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
			p.Match(CaliburnParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(257)
			p.Untyped_assign_declaration()
		}

	case 4:
		localctx = NewStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(258)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Assign_declarations()
		}
		{
			p.SetState(260)
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
			p.SetState(262)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Assign_declarations()
		}
		{
			p.SetState(264)
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
	p.EnterRule(localctx, 48, CaliburnParserRULE_typed_assign_declarations)
	var _la int

	localctx = NewTypedAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Typed_assign_declaration()
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(269)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(270)
			p.Typed_assign_declaration()
		}

		p.SetState(275)
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
	p.EnterRule(localctx, 50, CaliburnParserRULE_typed_assign_declaration)
	localctx = NewTypedAssignDeclarationContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)

		var _x = p.type_expression(0)

		localctx.(*TypedAssignDeclarationContext).type_ = _x
	}
	{
		p.SetState(277)
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
	p.EnterRule(localctx, 52, CaliburnParserRULE_untyped_assign_declarations)
	var _la int

	localctx = NewUntypedAssignDeclarationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Untyped_assign_declaration()
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(280)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Untyped_assign_declaration()
		}

		p.SetState(286)
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
	p.EnterRule(localctx, 54, CaliburnParserRULE_untyped_assign_declaration)
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifier:
		localctx = NewUntypedIdentifierAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(287)

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
			p.SetState(288)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(289)
			p.Untyped_assign_declarations()
		}
		{
			p.SetState(290)
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
			p.SetState(292)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(293)
			p.Untyped_assign_declarations()
		}
		{
			p.SetState(294)
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
	p.EnterRule(localctx, 56, CaliburnParserRULE_expression_statement)
	localctx = NewExpressionStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Expressions()
	}
	{
		p.SetState(299)
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
	p.EnterRule(localctx, 58, CaliburnParserRULE_expressions)
	var _la int

	localctx = NewExpressionsRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(301)
		p.expression(0)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(302)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(303)
			p.expression(0)
		}

		p.SetState(308)
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
	_startState := 60
	p.EnterRecursionRule(localctx, 60, CaliburnParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCastExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(310)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)

			var _x = p.type_expression(0)

			localctx.(*CastExpressionContext).type_ = _x
		}
		{
			p.SetState(312)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)

			var _x = p.expression(20)

			localctx.(*CastExpressionContext).exp = _x
		}

	case 2:
		localctx = NewBracketedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)

			var _x = p.expression(0)

			localctx.(*BracketedExpressionContext).exp = _x
		}
		{
			p.SetState(317)
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
			p.SetState(319)

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
			p.SetState(320)

			var _x = p.expression(17)

			localctx.(*UnaryExpressionContext).exp = _x
		}

	case 4:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(321)

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
			p.SetState(322)

			var _x = p.Literal_atom()

			localctx.(*LiteralExpressionContext).literal = _x
		}

	case 6:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)

			var _x = p.Function_type()

			localctx.(*FunctionExpressionContext).type_ = _x
		}
		{
			p.SetState(324)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432346067917340672) != 0 {
			{
				p.SetState(325)
				p.Assign_declarations()
			}

		}
		{
			p.SetState(328)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(330)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CaliburnParserIdentifier {
			{
				p.SetState(329)

				var _x = p.type_expression(0)

				localctx.(*FunctionExpressionContext).return_type = _x
			}

		}
		{
			p.SetState(332)
			p.Block()
		}

	case 7:
		localctx = NewStructExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(334)

			var _x = p.Struct_type()

			localctx.(*StructExpressionContext).type_ = _x
		}
		{
			p.SetState(335)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(336)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(337)
			p.Match(CaliburnParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.expression(0)
		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(339)
					p.Match(CaliburnParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(340)
					p.Match(CaliburnParserIdentifier)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

				{
					p.SetState(341)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(342)
					p.expression(0)
				}

			}
			p.SetState(347)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(349)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == CaliburnParserCOMMA {
			{
				p.SetState(348)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(351)
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
			p.SetState(353)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(354)
			p.expression(0)
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(355)
				p.Match(CaliburnParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case 2:
			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == CaliburnParserCOMMA {
				{
					p.SetState(356)
					p.Match(CaliburnParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(357)
					p.expression(0)
				}

				p.SetState(360)
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
			p.SetState(364)
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
	p.SetState(418)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(416)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(368)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(369)

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
					p.SetState(370)

					var _x = p.expression(19)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(371)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(372)

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
					p.SetState(373)

					var _x = p.expression(17)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 3:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(374)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(375)

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
					p.SetState(376)

					var _x = p.expression(16)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 4:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(378)

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
					p.SetState(379)

					var _x = p.expression(15)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 5:
				localctx = NewBooleanBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BooleanBinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(380)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(381)

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
					p.SetState(382)

					var _x = p.expression(14)

					localctx.(*BooleanBinaryExpressionContext).rhs_exp = _x
				}

			case 6:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(383)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(384)
					p.Match(CaliburnParserOP_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(385)

					var _x = p.expression(13)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 7:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(386)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(387)
					p.Match(CaliburnParserOP_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(388)

					var _x = p.expression(12)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 8:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).lhs_exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(389)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(390)
					p.Match(CaliburnParserOP_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(391)

					var _x = p.expression(11)

					localctx.(*BinaryExpressionContext).rhs_exp = _x
				}

			case 9:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*CallExpressionContext).exp = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(392)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(393)
					p.Match(CaliburnParserL_PAREN)
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

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432346046438309888) != 0 {
					{
						p.SetState(394)

						var _x = p.Expressions()

						localctx.(*CallExpressionContext).args = _x
					}

				}
				{
					p.SetState(397)
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
				p.SetState(398)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(399)
					p.Match(CaliburnParserPERIOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(400)

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
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(402)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(403)

					var _x = p.expression(0)

					localctx.(*IndexExpressionContext).index = _x
				}
				{
					p.SetState(404)
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
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(407)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(409)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432346046438309888) != 0 {
					{
						p.SetState(408)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).start_index = _x
					}

				}
				{
					p.SetState(411)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(413)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&432346046438309888) != 0 {
					{
						p.SetState(412)

						var _x = p.expression(0)

						localctx.(*SliceExpressionContext).end_index = _x
					}

				}
				{
					p.SetState(415)
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
		p.SetState(420)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 39, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 62, CaliburnParserRULE_function_type)
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserFUNC:
		localctx = NewFunctionTypeFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
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
			p.SetState(422)
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
	p.EnterRule(localctx, 64, CaliburnParserRULE_struct_type)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserSTRUCT:
		localctx = NewStructTypeStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
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
			p.SetState(426)
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
	p.EnterRule(localctx, 66, CaliburnParserRULE_tuple_type)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserTUPLE:
		localctx = NewTupleTypeTupleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(429)
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
			p.SetState(430)
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
	_startState := 68
	p.EnterRecursionRule(localctx, 68, CaliburnParserRULE_type_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewIdentifierTypeExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(434)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
			p.SetState(436)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				goto errorExit
			}
			{
				p.SetState(437)
				p.Match(CaliburnParserPERIOD)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(438)
				p.Match(CaliburnParserIdentifier)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 70, CaliburnParserRULE_identifiers)
	var _la int

	localctx = NewIdentifiersRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(444)
		p.Match(CaliburnParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(445)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(446)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(451)
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

func (s *UntypedLiteralAtomContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
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

func (s *TypedLiteralAtomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *TypedLiteralAtomContext) Literal() antlr.TerminalNode {
	return s.GetToken(CaliburnParserLiteral, 0)
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
	p.EnterRule(localctx, 72, CaliburnParserRULE_literal_atom)
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserLiteral:
		localctx = NewUntypedLiteralAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Match(CaliburnParserLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserIdentifier:
		localctx = NewTypedLiteralAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.Match(CaliburnParserIdentifier)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(454)
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
	case 30:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 34:
		var t *Type_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Type_expressionContext)
		}
		return p.Type_expression_Sempred(t, predIndex)

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

func (p *CaliburnParser) Type_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

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
		"", "'.'", "','", "':'", "';'", "'?'", "'='", "'if'", "'else'", "'for'",
		"'switch'", "'case'", "'default'", "'fallthrough'", "'break'", "'continue'",
		"'return'", "'using'", "'as'", "'in'", "'null'", "'var'", "'const'",
		"'type'", "'func'", "'struct'", "'tuple'", "'class'", "'interface'",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'+'", "'-'", "'!'", "'*'",
		"'/'", "'%'", "'^'", "'~'", "'|'", "'|!&'", "'&'", "'<<'", "'>>'", "'=='",
		"'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "PERIOD", "COMMA", "COLON", "SEMICOLON", "QUESTION", "ASSIGN", "IF",
		"ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", "BREAK",
		"CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", "TYPE",
		"FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN",
		"L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD", "OP_SUB",
		"OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", "OP_OR",
		"OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ", "OP_GTE",
		"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier",
		"Literal", "Terminator",
	}
	staticData.RuleNames = []string{
		"scoped_block", "statement", "control_statement", "if_statement", "for_statement",
		"switch_statement", "inline_statement", "assign_statement", "assign_declarations",
		"assign_declaration", "declared_assign_declarations", "declared_assign_declaration",
		"undeclared_assign_declarations", "undeclared_assign_declaration", "operator_assign_statement",
		"expression_statement", "expressions", "expression", "expression_atom",
		"value_atoms", "value_atom", "type_atom", "literal_atom", "untyped_literal_atom",
		"typed_literal_atom",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 288, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 5, 0,
		53, 8, 0, 10, 0, 12, 0, 56, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 63,
		8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 68, 8, 2, 1, 3, 1, 3, 5, 3, 72, 8, 3, 10,
		3, 12, 3, 75, 9, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 82, 8, 3, 3, 3,
		84, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 5, 5, 92, 8, 5, 10, 5, 12,
		5, 95, 9, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 103, 8, 5, 1, 5,
		1, 5, 5, 5, 107, 8, 5, 10, 5, 12, 5, 110, 9, 5, 1, 5, 1, 5, 3, 5, 114,
		8, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 120, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 128, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 133, 8, 8, 10, 8,
		12, 8, 136, 9, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 3, 9, 148, 8, 9, 1, 10, 1, 10, 1, 10, 5, 10, 153, 8, 10, 10, 10,
		12, 10, 156, 9, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 163, 8, 11,
		1, 12, 1, 12, 1, 12, 5, 12, 168, 8, 12, 10, 12, 12, 12, 171, 9, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 182,
		8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 5, 16, 195, 8, 16, 10, 16, 12, 16, 198, 9, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 208, 8, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 237, 8, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 251,
		8, 17, 1, 17, 1, 17, 3, 17, 255, 8, 17, 1, 17, 5, 17, 258, 8, 17, 10, 17,
		12, 17, 261, 9, 17, 1, 18, 1, 18, 3, 18, 265, 8, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 270, 8, 19, 10, 19, 12, 19, 273, 9, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 3, 22, 281, 8, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 0, 1, 34, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 7, 2, 0, 35, 36, 38, 47,
		1, 0, 35, 37, 1, 0, 41, 42, 1, 0, 38, 40, 1, 0, 35, 36, 1, 0, 46, 47, 1,
		0, 48, 53, 307, 0, 50, 1, 0, 0, 0, 2, 62, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0,
		6, 69, 1, 0, 0, 0, 8, 85, 1, 0, 0, 0, 10, 89, 1, 0, 0, 0, 12, 119, 1, 0,
		0, 0, 14, 127, 1, 0, 0, 0, 16, 129, 1, 0, 0, 0, 18, 147, 1, 0, 0, 0, 20,
		149, 1, 0, 0, 0, 22, 162, 1, 0, 0, 0, 24, 164, 1, 0, 0, 0, 26, 181, 1,
		0, 0, 0, 28, 183, 1, 0, 0, 0, 30, 188, 1, 0, 0, 0, 32, 191, 1, 0, 0, 0,
		34, 207, 1, 0, 0, 0, 36, 264, 1, 0, 0, 0, 38, 266, 1, 0, 0, 0, 40, 274,
		1, 0, 0, 0, 42, 276, 1, 0, 0, 0, 44, 280, 1, 0, 0, 0, 46, 282, 1, 0, 0,
		0, 48, 284, 1, 0, 0, 0, 50, 54, 5, 33, 0, 0, 51, 53, 3, 2, 1, 0, 52, 51,
		1, 0, 0, 0, 53, 56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 57, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 57, 58, 5, 34, 0, 0, 58, 1, 1,
		0, 0, 0, 59, 63, 3, 14, 7, 0, 60, 63, 3, 30, 15, 0, 61, 63, 3, 4, 2, 0,
		62, 59, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 3, 1, 0,
		0, 0, 64, 68, 3, 6, 3, 0, 65, 68, 3, 8, 4, 0, 66, 68, 3, 10, 5, 0, 67,
		64, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 66, 1, 0, 0, 0, 68, 5, 1, 0, 0,
		0, 69, 73, 5, 7, 0, 0, 70, 72, 3, 12, 6, 0, 71, 70, 1, 0, 0, 0, 72, 75,
		1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 76, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 76, 77, 3, 34, 17, 0, 77, 83, 3, 0, 0, 0, 78, 81, 5,
		8, 0, 0, 79, 82, 3, 0, 0, 0, 80, 82, 3, 6, 3, 0, 81, 79, 1, 0, 0, 0, 81,
		80, 1, 0, 0, 0, 82, 84, 1, 0, 0, 0, 83, 78, 1, 0, 0, 0, 83, 84, 1, 0, 0,
		0, 84, 7, 1, 0, 0, 0, 85, 86, 5, 9, 0, 0, 86, 87, 3, 34, 17, 0, 87, 88,
		3, 0, 0, 0, 88, 9, 1, 0, 0, 0, 89, 93, 5, 10, 0, 0, 90, 92, 3, 12, 6, 0,
		91, 90, 1, 0, 0, 0, 92, 95, 1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 93, 94, 1,
		0, 0, 0, 94, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 97, 3, 34, 17, 0,
		97, 108, 5, 33, 0, 0, 98, 102, 5, 11, 0, 0, 99, 103, 3, 40, 20, 0, 100,
		103, 3, 44, 22, 0, 101, 103, 3, 42, 21, 0, 102, 99, 1, 0, 0, 0, 102, 100,
		1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 3, 0,
		0, 0, 105, 107, 1, 0, 0, 0, 106, 98, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0,
		108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 113, 1, 0, 0, 0, 110,
		108, 1, 0, 0, 0, 111, 112, 5, 12, 0, 0, 112, 114, 3, 0, 0, 0, 113, 111,
		1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 116, 5, 34,
		0, 0, 116, 11, 1, 0, 0, 0, 117, 120, 3, 14, 7, 0, 118, 120, 3, 30, 15,
		0, 119, 117, 1, 0, 0, 0, 119, 118, 1, 0, 0, 0, 120, 13, 1, 0, 0, 0, 121,
		122, 3, 16, 8, 0, 122, 123, 5, 6, 0, 0, 123, 124, 3, 32, 16, 0, 124, 125,
		5, 58, 0, 0, 125, 128, 1, 0, 0, 0, 126, 128, 3, 28, 14, 0, 127, 121, 1,
		0, 0, 0, 127, 126, 1, 0, 0, 0, 128, 15, 1, 0, 0, 0, 129, 134, 3, 18, 9,
		0, 130, 131, 5, 2, 0, 0, 131, 133, 3, 14, 7, 0, 132, 130, 1, 0, 0, 0, 133,
		136, 1, 0, 0, 0, 134, 132, 1, 0, 0, 0, 134, 135, 1, 0, 0, 0, 135, 17, 1,
		0, 0, 0, 136, 134, 1, 0, 0, 0, 137, 148, 3, 22, 11, 0, 138, 148, 3, 26,
		13, 0, 139, 140, 5, 33, 0, 0, 140, 141, 3, 16, 8, 0, 141, 142, 5, 34, 0,
		0, 142, 148, 1, 0, 0, 0, 143, 144, 5, 31, 0, 0, 144, 145, 3, 16, 8, 0,
		145, 146, 5, 31, 0, 0, 146, 148, 1, 0, 0, 0, 147, 137, 1, 0, 0, 0, 147,
		138, 1, 0, 0, 0, 147, 139, 1, 0, 0, 0, 147, 143, 1, 0, 0, 0, 148, 19, 1,
		0, 0, 0, 149, 154, 3, 22, 11, 0, 150, 151, 5, 2, 0, 0, 151, 153, 3, 22,
		11, 0, 152, 150, 1, 0, 0, 0, 153, 156, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0,
		154, 155, 1, 0, 0, 0, 155, 21, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 157, 158,
		3, 42, 21, 0, 158, 159, 3, 26, 13, 0, 159, 163, 1, 0, 0, 0, 160, 161, 5,
		21, 0, 0, 161, 163, 3, 26, 13, 0, 162, 157, 1, 0, 0, 0, 162, 160, 1, 0,
		0, 0, 163, 23, 1, 0, 0, 0, 164, 169, 3, 26, 13, 0, 165, 166, 5, 2, 0, 0,
		166, 168, 3, 26, 13, 0, 167, 165, 1, 0, 0, 0, 168, 171, 1, 0, 0, 0, 169,
		167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 25, 1, 0, 0, 0, 171, 169, 1,
		0, 0, 0, 172, 182, 3, 40, 20, 0, 173, 174, 5, 33, 0, 0, 174, 175, 3, 24,
		12, 0, 175, 176, 5, 34, 0, 0, 176, 182, 1, 0, 0, 0, 177, 178, 5, 31, 0,
		0, 178, 179, 3, 24, 12, 0, 179, 180, 5, 32, 0, 0, 180, 182, 1, 0, 0, 0,
		181, 172, 1, 0, 0, 0, 181, 173, 1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 182,
		27, 1, 0, 0, 0, 183, 184, 3, 38, 19, 0, 184, 185, 7, 0, 0, 0, 185, 186,
		3, 32, 16, 0, 186, 187, 5, 58, 0, 0, 187, 29, 1, 0, 0, 0, 188, 189, 3,
		32, 16, 0, 189, 190, 5, 58, 0, 0, 190, 31, 1, 0, 0, 0, 191, 196, 3, 34,
		17, 0, 192, 193, 5, 2, 0, 0, 193, 195, 3, 34, 17, 0, 194, 192, 1, 0, 0,
		0, 195, 198, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197,
		33, 1, 0, 0, 0, 198, 196, 1, 0, 0, 0, 199, 200, 6, 17, -1, 0, 200, 201,
		5, 29, 0, 0, 201, 202, 3, 34, 17, 0, 202, 203, 5, 30, 0, 0, 203, 208, 1,
		0, 0, 0, 204, 205, 7, 1, 0, 0, 205, 208, 3, 34, 17, 13, 206, 208, 3, 36,
		18, 0, 207, 199, 1, 0, 0, 0, 207, 204, 1, 0, 0, 0, 207, 206, 1, 0, 0, 0,
		208, 259, 1, 0, 0, 0, 209, 210, 10, 14, 0, 0, 210, 211, 7, 2, 0, 0, 211,
		258, 3, 34, 17, 15, 212, 213, 10, 12, 0, 0, 213, 214, 7, 3, 0, 0, 214,
		258, 3, 34, 17, 13, 215, 216, 10, 11, 0, 0, 216, 217, 7, 4, 0, 0, 217,
		258, 3, 34, 17, 12, 218, 219, 10, 10, 0, 0, 219, 220, 7, 5, 0, 0, 220,
		258, 3, 34, 17, 11, 221, 222, 10, 9, 0, 0, 222, 223, 7, 6, 0, 0, 223, 258,
		3, 34, 17, 10, 224, 225, 10, 8, 0, 0, 225, 226, 5, 45, 0, 0, 226, 258,
		3, 34, 17, 9, 227, 228, 10, 7, 0, 0, 228, 229, 5, 44, 0, 0, 229, 258, 3,
		34, 17, 8, 230, 231, 10, 6, 0, 0, 231, 232, 5, 43, 0, 0, 232, 258, 3, 34,
		17, 7, 233, 234, 10, 5, 0, 0, 234, 236, 5, 29, 0, 0, 235, 237, 3, 32, 16,
		0, 236, 235, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238,
		258, 5, 30, 0, 0, 239, 240, 10, 4, 0, 0, 240, 241, 5, 1, 0, 0, 241, 258,
		3, 40, 20, 0, 242, 243, 10, 3, 0, 0, 243, 244, 5, 31, 0, 0, 244, 245, 3,
		34, 17, 0, 245, 246, 5, 32, 0, 0, 246, 258, 1, 0, 0, 0, 247, 248, 10, 2,
		0, 0, 248, 250, 5, 31, 0, 0, 249, 251, 3, 34, 17, 0, 250, 249, 1, 0, 0,
		0, 250, 251, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 254, 5, 3, 0, 0, 253,
		255, 3, 34, 17, 0, 254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256,
		1, 0, 0, 0, 256, 258, 5, 32, 0, 0, 257, 209, 1, 0, 0, 0, 257, 212, 1, 0,
		0, 0, 257, 215, 1, 0, 0, 0, 257, 218, 1, 0, 0, 0, 257, 221, 1, 0, 0, 0,
		257, 224, 1, 0, 0, 0, 257, 227, 1, 0, 0, 0, 257, 230, 1, 0, 0, 0, 257,
		233, 1, 0, 0, 0, 257, 239, 1, 0, 0, 0, 257, 242, 1, 0, 0, 0, 257, 247,
		1, 0, 0, 0, 258, 261, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259, 260, 1, 0,
		0, 0, 260, 35, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 262, 265, 3, 40, 20, 0,
		263, 265, 3, 44, 22, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265,
		37, 1, 0, 0, 0, 266, 271, 3, 40, 20, 0, 267, 268, 5, 2, 0, 0, 268, 270,
		3, 40, 20, 0, 269, 267, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1,
		0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 39, 1, 0, 0, 0, 273, 271, 1, 0, 0,
		0, 274, 275, 5, 56, 0, 0, 275, 41, 1, 0, 0, 0, 276, 277, 5, 56, 0, 0, 277,
		43, 1, 0, 0, 0, 278, 281, 3, 46, 23, 0, 279, 281, 3, 48, 24, 0, 280, 278,
		1, 0, 0, 0, 280, 279, 1, 0, 0, 0, 281, 45, 1, 0, 0, 0, 282, 283, 5, 57,
		0, 0, 283, 47, 1, 0, 0, 0, 284, 285, 3, 42, 21, 0, 285, 286, 5, 57, 0,
		0, 286, 49, 1, 0, 0, 0, 28, 54, 62, 67, 73, 81, 83, 93, 102, 108, 113,
		119, 127, 134, 147, 154, 162, 169, 181, 196, 207, 236, 250, 254, 257, 259,
		264, 271, 280,
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
	CaliburnParserSEMICOLON   = 4
	CaliburnParserQUESTION    = 5
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
	CaliburnParserTerminator  = 58
)

// CaliburnParser rules.
const (
	CaliburnParserRULE_scoped_block                   = 0
	CaliburnParserRULE_statement                      = 1
	CaliburnParserRULE_control_statement              = 2
	CaliburnParserRULE_if_statement                   = 3
	CaliburnParserRULE_for_statement                  = 4
	CaliburnParserRULE_switch_statement               = 5
	CaliburnParserRULE_inline_statement               = 6
	CaliburnParserRULE_assign_statement               = 7
	CaliburnParserRULE_assign_declarations            = 8
	CaliburnParserRULE_assign_declaration             = 9
	CaliburnParserRULE_declared_assign_declarations   = 10
	CaliburnParserRULE_declared_assign_declaration    = 11
	CaliburnParserRULE_undeclared_assign_declarations = 12
	CaliburnParserRULE_undeclared_assign_declaration  = 13
	CaliburnParserRULE_operator_assign_statement      = 14
	CaliburnParserRULE_expression_statement           = 15
	CaliburnParserRULE_expressions                    = 16
	CaliburnParserRULE_expression                     = 17
	CaliburnParserRULE_expression_atom                = 18
	CaliburnParserRULE_value_atoms                    = 19
	CaliburnParserRULE_value_atom                     = 20
	CaliburnParserRULE_type_atom                      = 21
	CaliburnParserRULE_literal_atom                   = 22
	CaliburnParserRULE_untyped_literal_atom           = 23
	CaliburnParserRULE_typed_literal_atom             = 24
)

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

func (p *CaliburnParser) Scoped_block() (localctx IScoped_blockContext) {
	localctx = NewScoped_blockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CaliburnParserRULE_scoped_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173033908340352) != 0 {
		{
			p.SetState(51)
			p.Statement()
		}

		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(57)
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

func (p *CaliburnParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CaliburnParserRULE_statement)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Expression_statement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(61)
			p.Control_statement()
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

func (p *CaliburnParser) Control_statement() (localctx IControl_statementContext) {
	localctx = NewControl_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CaliburnParserRULE_control_statement)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.If_statement()
		}

	case CaliburnParserFOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.For_statement()
		}

	case CaliburnParserSWITCH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
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

func (p *CaliburnParser) If_statement() (localctx IIf_statementContext) {
	localctx = NewIf_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CaliburnParserRULE_if_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(CaliburnParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(73)
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
			{
				p.SetState(70)
				p.Inline_statement()
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(76)
		p.expression(0)
	}
	{
		p.SetState(77)
		p.Scoped_block()
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserELSE {
		{
			p.SetState(78)
			p.Match(CaliburnParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case CaliburnParserL_C_BRACK:
			{
				p.SetState(79)
				p.Scoped_block()
			}

		case CaliburnParserIF:
			{
				p.SetState(80)
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

func (p *CaliburnParser) For_statement() (localctx IFor_statementContext) {
	localctx = NewFor_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CaliburnParserRULE_for_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(CaliburnParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(86)
		p.expression(0)
	}
	{
		p.SetState(87)
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
	Expression() IExpressionContext
	L_C_BRACK() antlr.TerminalNode
	R_C_BRACK() antlr.TerminalNode
	AllInline_statement() []IInline_statementContext
	Inline_statement(i int) IInline_statementContext
	AllCASE() []antlr.TerminalNode
	CASE(i int) antlr.TerminalNode
	AllScoped_block() []IScoped_blockContext
	Scoped_block(i int) IScoped_blockContext
	DEFAULT() antlr.TerminalNode
	AllValue_atom() []IValue_atomContext
	Value_atom(i int) IValue_atomContext
	AllLiteral_atom() []ILiteral_atomContext
	Literal_atom(i int) ILiteral_atomContext
	AllType_atom() []IType_atomContext
	Type_atom(i int) IType_atomContext

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

func (s *Switch_statementContext) Expression() IExpressionContext {
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

func (s *Switch_statementContext) AllValue_atom() []IValue_atomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValue_atomContext); ok {
			len++
		}
	}

	tst := make([]IValue_atomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValue_atomContext); ok {
			tst[i] = t.(IValue_atomContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Value_atom(i int) IValue_atomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
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

	return t.(IValue_atomContext)
}

func (s *Switch_statementContext) AllLiteral_atom() []ILiteral_atomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteral_atomContext); ok {
			len++
		}
	}

	tst := make([]ILiteral_atomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteral_atomContext); ok {
			tst[i] = t.(ILiteral_atomContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Literal_atom(i int) ILiteral_atomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_atomContext); ok {
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

	return t.(ILiteral_atomContext)
}

func (s *Switch_statementContext) AllType_atom() []IType_atomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_atomContext); ok {
			len++
		}
	}

	tst := make([]IType_atomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_atomContext); ok {
			tst[i] = t.(IType_atomContext)
			i++
		}
	}

	return tst
}

func (s *Switch_statementContext) Type_atom(i int) IType_atomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_atomContext); ok {
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

	return t.(IType_atomContext)
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

func (p *CaliburnParser) Switch_statement() (localctx ISwitch_statementContext) {
	localctx = NewSwitch_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CaliburnParserRULE_switch_statement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Match(CaliburnParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(93)
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
				p.SetState(90)
				p.Inline_statement()
			}

		}
		p.SetState(95)
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
		p.SetState(96)
		p.expression(0)
	}
	{
		p.SetState(97)
		p.Match(CaliburnParserL_C_BRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCASE {
		{
			p.SetState(98)
			p.Match(CaliburnParserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(99)
				p.Value_atom()
			}

		case 2:
			{
				p.SetState(100)
				p.Literal_atom()
			}

		case 3:
			{
				p.SetState(101)
				p.Type_atom()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		{
			p.SetState(104)
			p.Scoped_block()
		}

		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == CaliburnParserDEFAULT {
		{
			p.SetState(111)
			p.Match(CaliburnParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Scoped_block()
		}

	}
	{
		p.SetState(115)
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

func (p *CaliburnParser) Inline_statement() (localctx IInline_statementContext) {
	localctx = NewInline_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CaliburnParserRULE_inline_statement)
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Assign_statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
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

func (p *CaliburnParser) Assign_statement() (localctx IAssign_statementContext) {
	localctx = NewAssign_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CaliburnParserRULE_assign_statement)
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Assign_declarations()
		}
		{
			p.SetState(122)
			p.Match(CaliburnParserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Expressions()
		}
		{
			p.SetState(124)
			p.Match(CaliburnParserTerminator)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(126)
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

func (p *CaliburnParser) Assign_declarations() (localctx IAssign_declarationsContext) {
	localctx = NewAssign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CaliburnParserRULE_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(129)
		p.Assign_declaration()
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(130)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Assign_statement()
		}

		p.SetState(136)
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

func (p *CaliburnParser) Assign_declaration() (localctx IAssign_declarationContext) {
	localctx = NewAssign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CaliburnParserRULE_assign_declaration)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Declared_assign_declaration()
		}

	case 2:
		localctx = NewUndeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Undeclared_assign_declaration()
		}

	case 3:
		localctx = NewStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Assign_declarations()
		}
		{
			p.SetState(141)
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
			p.SetState(143)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Assign_declarations()
		}
		{
			p.SetState(145)
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

func (p *CaliburnParser) Declared_assign_declarations() (localctx IDeclared_assign_declarationsContext) {
	localctx = NewDeclared_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CaliburnParserRULE_declared_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Declared_assign_declaration()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(150)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Declared_assign_declaration()
		}

		p.SetState(156)
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

type TypedDeclaredAssignDeclarationContext struct {
	Declared_assign_declarationContext
}

func NewTypedDeclaredAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedDeclaredAssignDeclarationContext {
	var p = new(TypedDeclaredAssignDeclarationContext)

	InitEmptyDeclared_assign_declarationContext(&p.Declared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declared_assign_declarationContext))

	return p
}

func (s *TypedDeclaredAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedDeclaredAssignDeclarationContext) Type_atom() IType_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_atomContext)
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

func (p *CaliburnParser) Declared_assign_declaration() (localctx IDeclared_assign_declarationContext) {
	localctx = NewDeclared_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CaliburnParserRULE_declared_assign_declaration)
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifier:
		localctx = NewTypedDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Type_atom()
		}
		{
			p.SetState(158)
			p.Undeclared_assign_declaration()
		}

	case CaliburnParserVAR:
		localctx = NewUntypedDeclaredAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Match(CaliburnParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
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

func (p *CaliburnParser) Undeclared_assign_declarations() (localctx IUndeclared_assign_declarationsContext) {
	localctx = NewUndeclared_assign_declarationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CaliburnParserRULE_undeclared_assign_declarations)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Undeclared_assign_declaration()
	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(165)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Undeclared_assign_declaration()
		}

		p.SetState(171)
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
}

func NewUndeclaredAtomAssignDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UndeclaredAtomAssignDeclarationContext {
	var p = new(UndeclaredAtomAssignDeclarationContext)

	InitEmptyUndeclared_assign_declarationContext(&p.Undeclared_assign_declarationContext)
	p.parser = parser
	p.CopyAll(ctx.(*Undeclared_assign_declarationContext))

	return p
}

func (s *UndeclaredAtomAssignDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UndeclaredAtomAssignDeclarationContext) Value_atom() IValue_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_atomContext)
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

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) L_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserL_S_BRACK, 0)
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

func (s *UndeclaredTupleDestrutureAssignDeclarationContext) R_S_BRACK() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_S_BRACK, 0)
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

func (p *CaliburnParser) Undeclared_assign_declaration() (localctx IUndeclared_assign_declarationContext) {
	localctx = NewUndeclared_assign_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CaliburnParserRULE_undeclared_assign_declaration)
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserIdentifier:
		localctx = NewUndeclaredAtomAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Value_atom()
		}

	case CaliburnParserL_C_BRACK:
		localctx = NewUndeclaredStructDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Match(CaliburnParserL_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Undeclared_assign_declarations()
		}
		{
			p.SetState(175)
			p.Match(CaliburnParserR_C_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserL_S_BRACK:
		localctx = NewUndeclaredTupleDestrutureAssignDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.Match(CaliburnParserL_S_BRACK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.Undeclared_assign_declarations()
		}
		{
			p.SetState(179)
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
	op antlr.Token
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

func (s *OperatorAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorAssignmentContext) Value_atoms() IValue_atomsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_atomsContext)
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

func (p *CaliburnParser) Operator_assign_statement() (localctx IOperator_assign_statementContext) {
	localctx = NewOperator_assign_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CaliburnParserRULE_operator_assign_statement)
	var _la int

	localctx = NewOperatorAssignmentContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Value_atoms()
	}
	{
		p.SetState(184)

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
		p.SetState(185)
		p.Expressions()
	}
	{
		p.SetState(186)
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

func (p *CaliburnParser) Expression_statement() (localctx IExpression_statementContext) {
	localctx = NewExpression_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CaliburnParserRULE_expression_statement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(188)
		p.Expressions()
	}
	{
		p.SetState(189)
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

func (p *CaliburnParser) Expressions() (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CaliburnParserRULE_expressions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.expression(0)
	}
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == CaliburnParserCOMMA {
		{
			p.SetState(192)
			p.Match(CaliburnParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.expression(0)
		}

		p.SetState(198)
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

type ExponentExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewExponentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExponentExpressionContext {
	var p = new(ExponentExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExponentExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ExponentExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExponentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExponentExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ExponentExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ExponentExpressionContext) OP_POW() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_POW, 0)
}

func (s *ExponentExpressionContext) OP_ROOT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ROOT, 0)
}

func (s *ExponentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExponentExpression(s)
	}
}

func (s *ExponentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExponentExpression(s)
	}
}

type AdditionExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAdditionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AdditionExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AdditionExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AdditionExpressionContext) Expression(i int) IExpressionContext {
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

func (s *AdditionExpressionContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_ADD, 0)
}

func (s *AdditionExpressionContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_SUB, 0)
}

func (s *AdditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAdditionExpression(s)
	}
}

func (s *AdditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAdditionExpression(s)
	}
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

type AtomExpressionContext struct {
	ExpressionContext
}

func NewAtomExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtomExpressionContext {
	var p = new(AtomExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AtomExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomExpressionContext) Expression_atom() IExpression_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpression_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpression_atomContext)
}

func (s *AtomExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAtomExpression(s)
	}
}

func (s *AtomExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAtomExpression(s)
	}
}

type ShiftExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftExpressionContext {
	var p = new(ShiftExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ShiftExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ShiftExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ShiftExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ShiftExpressionContext) OP_LSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LSHIFT, 0)
}

func (s *ShiftExpressionContext) OP_RSHIFT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_RSHIFT, 0)
}

func (s *ShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterShiftExpression(s)
	}
}

func (s *ShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitShiftExpression(s)
	}
}

type OrExpressionContext struct {
	ExpressionContext
}

func NewOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExpressionContext {
	var p = new(OrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExpressionContext) AllExpression() []IExpressionContext {
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

func (s *OrExpressionContext) Expression(i int) IExpressionContext {
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

func (s *OrExpressionContext) OP_OR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_OR, 0)
}

func (s *OrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterOrExpression(s)
	}
}

func (s *OrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitOrExpression(s)
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

func (s *AccessExpressionContext) Value_atom() IValue_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_atomContext)
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

type ComparisonExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewComparisonExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExpressionContext {
	var p = new(ComparisonExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonExpressionContext) GetOp() antlr.Token { return s.op }

func (s *ComparisonExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ComparisonExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExpressionContext) AllExpression() []IExpressionContext {
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

func (s *ComparisonExpressionContext) Expression(i int) IExpressionContext {
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

func (s *ComparisonExpressionContext) OP_EQU() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_EQU, 0)
}

func (s *ComparisonExpressionContext) OP_NEQ() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_NEQ, 0)
}

func (s *ComparisonExpressionContext) OP_LTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LTE, 0)
}

func (s *ComparisonExpressionContext) OP_GTE() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GTE, 0)
}

func (s *ComparisonExpressionContext) OP_LST() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_LST, 0)
}

func (s *ComparisonExpressionContext) OP_GRT() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_GRT, 0)
}

func (s *ComparisonExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterComparisonExpression(s)
	}
}

func (s *ComparisonExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitComparisonExpression(s)
	}
}

type XOrExpressionContext struct {
	ExpressionContext
}

func NewXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XOrExpressionContext {
	var p = new(XOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *XOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XOrExpressionContext) AllExpression() []IExpressionContext {
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

func (s *XOrExpressionContext) Expression(i int) IExpressionContext {
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

func (s *XOrExpressionContext) OP_XOR() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_XOR, 0)
}

func (s *XOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterXOrExpression(s)
	}
}

func (s *XOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitXOrExpression(s)
	}
}

type AndExpressionContext struct {
	ExpressionContext
}

func NewAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExpressionContext {
	var p = new(AndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExpressionContext) AllExpression() []IExpressionContext {
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

func (s *AndExpressionContext) Expression(i int) IExpressionContext {
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

func (s *AndExpressionContext) OP_AND() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_AND, 0)
}

func (s *AndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterAndExpression(s)
	}
}

func (s *AndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitAndExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultiplicativeExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllExpression() []IExpressionContext {
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

func (s *MultiplicativeExpressionContext) Expression(i int) IExpressionContext {
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

func (s *MultiplicativeExpressionContext) OP_MUL() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MUL, 0)
}

func (s *MultiplicativeExpressionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_DIV, 0)
}

func (s *MultiplicativeExpressionContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(CaliburnParserOP_MOD, 0)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
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

func (s *CallExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(CaliburnParserR_PAREN, 0)
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

func (p *CaliburnParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *CaliburnParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, CaliburnParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserL_PAREN:
		localctx = NewBracketedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(200)
			p.Match(CaliburnParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
			p.expression(0)
		}
		{
			p.SetState(202)
			p.Match(CaliburnParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CaliburnParserOP_ADD, CaliburnParserOP_SUB, CaliburnParserOP_NOT:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(204)

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
			p.SetState(205)
			p.expression(13)
		}

	case CaliburnParserIdentifier, CaliburnParserLiteral:
		localctx = NewAtomExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(206)
			p.Expression_atom()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExponentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(210)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExponentExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_POW || _la == CaliburnParserOP_ROOT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExponentExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(211)
					p.expression(15)
				}

			case 2:
				localctx = NewMultiplicativeExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(213)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultiplicativeExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1924145348608) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultiplicativeExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(214)
					p.expression(13)
				}

			case 3:
				localctx = NewAdditionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(216)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AdditionExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_ADD || _la == CaliburnParserOP_SUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AdditionExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(217)
					p.expression(12)
				}

			case 4:
				localctx = NewShiftExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(219)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ShiftExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CaliburnParserOP_LSHIFT || _la == CaliburnParserOP_RSHIFT) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ShiftExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(220)
					p.expression(11)
				}

			case 5:
				localctx = NewComparisonExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(222)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ComparisonExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17732923532771328) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ComparisonExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(223)
					p.expression(10)
				}

			case 6:
				localctx = NewAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(225)
					p.Match(CaliburnParserOP_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(226)
					p.expression(9)
				}

			case 7:
				localctx = NewXOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(228)
					p.Match(CaliburnParserOP_XOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(229)
					p.expression(8)
				}

			case 8:
				localctx = NewOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(231)
					p.Match(CaliburnParserOP_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(232)
					p.expression(7)
				}

			case 9:
				localctx = NewCallExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(234)
					p.Match(CaliburnParserL_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(236)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023168823296) != 0 {
					{
						p.SetState(235)
						p.Expressions()
					}

				}
				{
					p.SetState(238)
					p.Match(CaliburnParserR_PAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 10:
				localctx = NewAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(240)
					p.Match(CaliburnParserPERIOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(241)
					p.Value_atom()
				}

			case 11:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(243)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(244)
					p.expression(0)
				}
				{
					p.SetState(245)
					p.Match(CaliburnParserR_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case 12:
				localctx = NewSliceExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CaliburnParserRULE_expression)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(248)
					p.Match(CaliburnParserL_S_BRACK)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(250)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023168823296) != 0 {
					{
						p.SetState(249)
						p.expression(0)
					}

				}
				{
					p.SetState(252)
					p.Match(CaliburnParserCOLON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216173023168823296) != 0 {
					{
						p.SetState(253)
						p.expression(0)
					}

				}
				{
					p.SetState(256)
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
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext())
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

// IExpression_atomContext is an interface to support dynamic dispatch.
type IExpression_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value_atom() IValue_atomContext
	Literal_atom() ILiteral_atomContext

	// IsExpression_atomContext differentiates from other interfaces.
	IsExpression_atomContext()
}

type Expression_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_atomContext() *Expression_atomContext {
	var p = new(Expression_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_atom
	return p
}

func InitEmptyExpression_atomContext(p *Expression_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_expression_atom
}

func (*Expression_atomContext) IsExpression_atomContext() {}

func NewExpression_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_atomContext {
	var p = new(Expression_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_expression_atom

	return p
}

func (s *Expression_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_atomContext) Value_atom() IValue_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValue_atomContext)
}

func (s *Expression_atomContext) Literal_atom() ILiteral_atomContext {
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

func (s *Expression_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterExpression_atom(s)
	}
}

func (s *Expression_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitExpression_atom(s)
	}
}

func (p *CaliburnParser) Expression_atom() (localctx IExpression_atomContext) {
	localctx = NewExpression_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CaliburnParserRULE_expression_atom)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Value_atom()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Literal_atom()
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

// IValue_atomsContext is an interface to support dynamic dispatch.
type IValue_atomsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValue_atom() []IValue_atomContext
	Value_atom(i int) IValue_atomContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsValue_atomsContext differentiates from other interfaces.
	IsValue_atomsContext()
}

type Value_atomsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_atomsContext() *Value_atomsContext {
	var p = new(Value_atomsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atoms
	return p
}

func InitEmptyValue_atomsContext(p *Value_atomsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atoms
}

func (*Value_atomsContext) IsValue_atomsContext() {}

func NewValue_atomsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_atomsContext {
	var p = new(Value_atomsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_value_atoms

	return p
}

func (s *Value_atomsContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_atomsContext) AllValue_atom() []IValue_atomContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValue_atomContext); ok {
			len++
		}
	}

	tst := make([]IValue_atomContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValue_atomContext); ok {
			tst[i] = t.(IValue_atomContext)
			i++
		}
	}

	return tst
}

func (s *Value_atomsContext) Value_atom(i int) IValue_atomContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValue_atomContext); ok {
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

	return t.(IValue_atomContext)
}

func (s *Value_atomsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CaliburnParserCOMMA)
}

func (s *Value_atomsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CaliburnParserCOMMA, i)
}

func (s *Value_atomsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_atomsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_atomsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterValue_atoms(s)
	}
}

func (s *Value_atomsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitValue_atoms(s)
	}
}

func (p *CaliburnParser) Value_atoms() (localctx IValue_atomsContext) {
	localctx = NewValue_atomsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CaliburnParserRULE_value_atoms)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Value_atom()
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
			p.Value_atom()
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

// IValue_atomContext is an interface to support dynamic dispatch.
type IValue_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsValue_atomContext differentiates from other interfaces.
	IsValue_atomContext()
}

type Value_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_atomContext() *Value_atomContext {
	var p = new(Value_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atom
	return p
}

func InitEmptyValue_atomContext(p *Value_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_value_atom
}

func (*Value_atomContext) IsValue_atomContext() {}

func NewValue_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_atomContext {
	var p = new(Value_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_value_atom

	return p
}

func (s *Value_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_atomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Value_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterValue_atom(s)
	}
}

func (s *Value_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitValue_atom(s)
	}
}

func (p *CaliburnParser) Value_atom() (localctx IValue_atomContext) {
	localctx = NewValue_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CaliburnParserRULE_value_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(274)
		p.Match(CaliburnParserIdentifier)
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

// IType_atomContext is an interface to support dynamic dispatch.
type IType_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsType_atomContext differentiates from other interfaces.
	IsType_atomContext()
}

type Type_atomContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_atomContext() *Type_atomContext {
	var p = new(Type_atomContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_atom
	return p
}

func InitEmptyType_atomContext(p *Type_atomContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CaliburnParserRULE_type_atom
}

func (*Type_atomContext) IsType_atomContext() {}

func NewType_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_atomContext {
	var p = new(Type_atomContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CaliburnParserRULE_type_atom

	return p
}

func (s *Type_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_atomContext) Identifier() antlr.TerminalNode {
	return s.GetToken(CaliburnParserIdentifier, 0)
}

func (s *Type_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterType_atom(s)
	}
}

func (s *Type_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitType_atom(s)
	}
}

func (p *CaliburnParser) Type_atom() (localctx IType_atomContext) {
	localctx = NewType_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CaliburnParserRULE_type_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(276)
		p.Match(CaliburnParserIdentifier)
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

// ILiteral_atomContext is an interface to support dynamic dispatch.
type ILiteral_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Untyped_literal_atom() IUntyped_literal_atomContext
	Typed_literal_atom() ITyped_literal_atomContext

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

func (s *Literal_atomContext) Untyped_literal_atom() IUntyped_literal_atomContext {
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

func (s *Literal_atomContext) Typed_literal_atom() ITyped_literal_atomContext {
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

func (s *Literal_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.EnterLiteral_atom(s)
	}
}

func (s *Literal_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CaliburnParserListener); ok {
		listenerT.ExitLiteral_atom(s)
	}
}

func (p *CaliburnParser) Literal_atom() (localctx ILiteral_atomContext) {
	localctx = NewLiteral_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CaliburnParserRULE_literal_atom)
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CaliburnParserLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(278)
			p.Untyped_literal_atom()
		}

	case CaliburnParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
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

func (p *CaliburnParser) Untyped_literal_atom() (localctx IUntyped_literal_atomContext) {
	localctx = NewUntyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CaliburnParserRULE_untyped_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
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
	Type_atom() IType_atomContext
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

func (s *Typed_literal_atomContext) Type_atom() IType_atomContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_atomContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_atomContext)
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

func (p *CaliburnParser) Typed_literal_atom() (localctx ITyped_literal_atomContext) {
	localctx = NewTyped_literal_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CaliburnParserRULE_typed_literal_atom)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Type_atom()
	}
	{
		p.SetState(285)
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
	case 17:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CaliburnParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

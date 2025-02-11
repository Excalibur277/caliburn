// Code generated from CaliburnLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parsing

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CaliburnLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CaliburnLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func caliburnlexerLexerInit() {
	staticData := &CaliburnLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
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
		"PERIOD", "COMMA", "COLON", "QUESTION", "Terminator", "ASSIGN", "IF",
		"ELSE", "FOR", "SWITCH", "CASE", "DEFAULT", "FALLTHROUGH", "BREAK",
		"CONTINUE", "RETURN", "USING", "AS", "IN", "NULL", "VAR", "CONST", "TYPE",
		"FUNC", "STRUCT", "TUPLE", "CLASS", "INTERFACE", "L_PAREN", "R_PAREN",
		"L_S_BRACK", "R_S_BRACK", "L_C_BRACK", "R_C_BRACK", "OP_ADD", "OP_SUB",
		"OP_NOT", "OP_MUL", "OP_DIV", "OP_MOD", "OP_POW", "OP_ROOT", "OP_OR",
		"OP_XOR", "OP_AND", "OP_LSHIFT", "OP_RSHIFT", "OP_EQU", "OP_NEQ", "OP_GTE",
		"OP_GRT", "OP_LTE", "OP_LST", "WhiteSpace", "Comment", "Identifier",
		"EOL", "ESCAPE_SEQUENCE", "LITERAL_TEXT", "SQUOTE", "DQUOTE", "Literal",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 57, 389, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31,
		1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42,
		1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1,
		46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 53, 4, 53, 329,
		8, 53, 11, 53, 12, 53, 330, 1, 53, 1, 53, 1, 54, 1, 54, 5, 54, 337, 8,
		54, 10, 54, 12, 54, 340, 9, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 346,
		8, 54, 1, 54, 3, 54, 349, 8, 54, 1, 55, 1, 55, 5, 55, 353, 8, 55, 10, 55,
		12, 55, 356, 9, 55, 1, 56, 3, 56, 359, 8, 56, 1, 56, 1, 56, 1, 57, 1, 57,
		1, 57, 1, 58, 1, 58, 4, 58, 368, 8, 58, 11, 58, 12, 58, 369, 1, 59, 1,
		59, 1, 60, 1, 60, 1, 61, 1, 61, 3, 61, 378, 8, 61, 1, 61, 1, 61, 1, 61,
		1, 61, 3, 61, 384, 8, 61, 1, 61, 1, 61, 3, 61, 388, 8, 61, 1, 338, 0, 62,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29,
		59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38,
		77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47,
		95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111,
		56, 113, 0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 57, 1, 0, 4, 10, 0, 9,
		13, 32, 32, 133, 133, 160, 160, 5760, 5760, 8192, 8202, 8232, 8233, 8239,
		8239, 8287, 8287, 12288, 12288, 5, 0, 10, 10, 13, 13, 34, 34, 39, 39, 92,
		92, 660, 0, 65, 90, 95, 95, 97, 122, 170, 170, 181, 181, 186, 186, 192,
		214, 216, 246, 248, 705, 710, 721, 736, 740, 748, 748, 750, 750, 880, 884,
		886, 887, 890, 893, 895, 895, 902, 902, 904, 906, 908, 908, 910, 929, 931,
		1013, 1015, 1153, 1162, 1327, 1329, 1366, 1369, 1369, 1376, 1416, 1488,
		1514, 1519, 1522, 1568, 1610, 1646, 1647, 1649, 1747, 1749, 1749, 1765,
		1766, 1774, 1775, 1786, 1788, 1791, 1791, 1808, 1808, 1810, 1839, 1869,
		1957, 1969, 1969, 1994, 2026, 2036, 2037, 2042, 2042, 2048, 2069, 2074,
		2074, 2084, 2084, 2088, 2088, 2112, 2136, 2144, 2154, 2160, 2183, 2185,
		2190, 2208, 2249, 2308, 2361, 2365, 2365, 2384, 2384, 2392, 2401, 2417,
		2432, 2437, 2444, 2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482, 2486,
		2489, 2493, 2493, 2510, 2510, 2524, 2525, 2527, 2529, 2544, 2545, 2556,
		2556, 2565, 2570, 2575, 2576, 2579, 2600, 2602, 2608, 2610, 2611, 2613,
		2614, 2616, 2617, 2649, 2652, 2654, 2654, 2674, 2676, 2693, 2701, 2703,
		2705, 2707, 2728, 2730, 2736, 2738, 2739, 2741, 2745, 2749, 2749, 2768,
		2768, 2784, 2785, 2809, 2809, 2821, 2828, 2831, 2832, 2835, 2856, 2858,
		2864, 2866, 2867, 2869, 2873, 2877, 2877, 2908, 2909, 2911, 2913, 2929,
		2929, 2947, 2947, 2949, 2954, 2958, 2960, 2962, 2965, 2969, 2970, 2972,
		2972, 2974, 2975, 2979, 2980, 2984, 2986, 2990, 3001, 3024, 3024, 3077,
		3084, 3086, 3088, 3090, 3112, 3114, 3129, 3133, 3133, 3160, 3162, 3165,
		3165, 3168, 3169, 3200, 3200, 3205, 3212, 3214, 3216, 3218, 3240, 3242,
		3251, 3253, 3257, 3261, 3261, 3293, 3294, 3296, 3297, 3313, 3314, 3332,
		3340, 3342, 3344, 3346, 3386, 3389, 3389, 3406, 3406, 3412, 3414, 3423,
		3425, 3450, 3455, 3461, 3478, 3482, 3505, 3507, 3515, 3517, 3517, 3520,
		3526, 3585, 3632, 3634, 3635, 3648, 3654, 3713, 3714, 3716, 3716, 3718,
		3722, 3724, 3747, 3749, 3749, 3751, 3760, 3762, 3763, 3773, 3773, 3776,
		3780, 3782, 3782, 3804, 3807, 3840, 3840, 3904, 3911, 3913, 3948, 3976,
		3980, 4096, 4138, 4159, 4159, 4176, 4181, 4186, 4189, 4193, 4193, 4197,
		4198, 4206, 4208, 4213, 4225, 4238, 4238, 4256, 4293, 4295, 4295, 4301,
		4301, 4304, 4346, 4348, 4680, 4682, 4685, 4688, 4694, 4696, 4696, 4698,
		4701, 4704, 4744, 4746, 4749, 4752, 4784, 4786, 4789, 4792, 4798, 4800,
		4800, 4802, 4805, 4808, 4822, 4824, 4880, 4882, 4885, 4888, 4954, 4992,
		5007, 5024, 5109, 5112, 5117, 5121, 5740, 5743, 5759, 5761, 5786, 5792,
		5866, 5873, 5880, 5888, 5905, 5919, 5937, 5952, 5969, 5984, 5996, 5998,
		6000, 6016, 6067, 6103, 6103, 6108, 6108, 6176, 6264, 6272, 6276, 6279,
		6312, 6314, 6314, 6320, 6389, 6400, 6430, 6480, 6509, 6512, 6516, 6528,
		6571, 6576, 6601, 6656, 6678, 6688, 6740, 6823, 6823, 6917, 6963, 6981,
		6988, 7043, 7072, 7086, 7087, 7098, 7141, 7168, 7203, 7245, 7247, 7258,
		7293, 7296, 7304, 7312, 7354, 7357, 7359, 7401, 7404, 7406, 7411, 7413,
		7414, 7418, 7418, 7424, 7615, 7680, 7957, 7960, 7965, 7968, 8005, 8008,
		8013, 8016, 8023, 8025, 8025, 8027, 8027, 8029, 8029, 8031, 8061, 8064,
		8116, 8118, 8124, 8126, 8126, 8130, 8132, 8134, 8140, 8144, 8147, 8150,
		8155, 8160, 8172, 8178, 8180, 8182, 8188, 8305, 8305, 8319, 8319, 8336,
		8348, 8450, 8450, 8455, 8455, 8458, 8467, 8469, 8469, 8473, 8477, 8484,
		8484, 8486, 8486, 8488, 8488, 8490, 8493, 8495, 8505, 8508, 8511, 8517,
		8521, 8526, 8526, 8579, 8580, 11264, 11492, 11499, 11502, 11506, 11507,
		11520, 11557, 11559, 11559, 11565, 11565, 11568, 11623, 11631, 11631, 11648,
		11670, 11680, 11686, 11688, 11694, 11696, 11702, 11704, 11710, 11712, 11718,
		11720, 11726, 11728, 11734, 11736, 11742, 11823, 11823, 12293, 12294, 12337,
		12341, 12347, 12348, 12353, 12438, 12445, 12447, 12449, 12538, 12540, 12543,
		12549, 12591, 12593, 12686, 12704, 12735, 12784, 12799, 13312, 19903, 19968,
		42124, 42192, 42237, 42240, 42508, 42512, 42527, 42538, 42539, 42560, 42606,
		42623, 42653, 42656, 42725, 42775, 42783, 42786, 42888, 42891, 42954, 42960,
		42961, 42963, 42963, 42965, 42969, 42994, 43009, 43011, 43013, 43015, 43018,
		43020, 43042, 43072, 43123, 43138, 43187, 43250, 43255, 43259, 43259, 43261,
		43262, 43274, 43301, 43312, 43334, 43360, 43388, 43396, 43442, 43471, 43471,
		43488, 43492, 43494, 43503, 43514, 43518, 43520, 43560, 43584, 43586, 43588,
		43595, 43616, 43638, 43642, 43642, 43646, 43695, 43697, 43697, 43701, 43702,
		43705, 43709, 43712, 43712, 43714, 43714, 43739, 43741, 43744, 43754, 43762,
		43764, 43777, 43782, 43785, 43790, 43793, 43798, 43808, 43814, 43816, 43822,
		43824, 43866, 43868, 43881, 43888, 44002, 44032, 55203, 55216, 55238, 55243,
		55291, 63744, 64109, 64112, 64217, 64256, 64262, 64275, 64279, 64285, 64285,
		64287, 64296, 64298, 64310, 64312, 64316, 64318, 64318, 64320, 64321, 64323,
		64324, 64326, 64433, 64467, 64829, 64848, 64911, 64914, 64967, 65008, 65019,
		65136, 65140, 65142, 65276, 65313, 65338, 65345, 65370, 65382, 65470, 65474,
		65479, 65482, 65487, 65490, 65495, 65498, 65500, 65536, 65547, 65549, 65574,
		65576, 65594, 65596, 65597, 65599, 65613, 65616, 65629, 65664, 65786, 66176,
		66204, 66208, 66256, 66304, 66335, 66349, 66368, 66370, 66377, 66384, 66421,
		66432, 66461, 66464, 66499, 66504, 66511, 66560, 66717, 66736, 66771, 66776,
		66811, 66816, 66855, 66864, 66915, 66928, 66938, 66940, 66954, 66956, 66962,
		66964, 66965, 66967, 66977, 66979, 66993, 66995, 67001, 67003, 67004, 67072,
		67382, 67392, 67413, 67424, 67431, 67456, 67461, 67463, 67504, 67506, 67514,
		67584, 67589, 67592, 67592, 67594, 67637, 67639, 67640, 67644, 67644, 67647,
		67669, 67680, 67702, 67712, 67742, 67808, 67826, 67828, 67829, 67840, 67861,
		67872, 67897, 67968, 68023, 68030, 68031, 68096, 68096, 68112, 68115, 68117,
		68119, 68121, 68149, 68192, 68220, 68224, 68252, 68288, 68295, 68297, 68324,
		68352, 68405, 68416, 68437, 68448, 68466, 68480, 68497, 68608, 68680, 68736,
		68786, 68800, 68850, 68864, 68899, 69248, 69289, 69296, 69297, 69376, 69404,
		69415, 69415, 69424, 69445, 69488, 69505, 69552, 69572, 69600, 69622, 69635,
		69687, 69745, 69746, 69749, 69749, 69763, 69807, 69840, 69864, 69891, 69926,
		69956, 69956, 69959, 69959, 69968, 70002, 70006, 70006, 70019, 70066, 70081,
		70084, 70106, 70106, 70108, 70108, 70144, 70161, 70163, 70187, 70207, 70208,
		70272, 70278, 70280, 70280, 70282, 70285, 70287, 70301, 70303, 70312, 70320,
		70366, 70405, 70412, 70415, 70416, 70419, 70440, 70442, 70448, 70450, 70451,
		70453, 70457, 70461, 70461, 70480, 70480, 70493, 70497, 70656, 70708, 70727,
		70730, 70751, 70753, 70784, 70831, 70852, 70853, 70855, 70855, 71040, 71086,
		71128, 71131, 71168, 71215, 71236, 71236, 71296, 71338, 71352, 71352, 71424,
		71450, 71488, 71494, 71680, 71723, 71840, 71903, 71935, 71942, 71945, 71945,
		71948, 71955, 71957, 71958, 71960, 71983, 71999, 71999, 72001, 72001, 72096,
		72103, 72106, 72144, 72161, 72161, 72163, 72163, 72192, 72192, 72203, 72242,
		72250, 72250, 72272, 72272, 72284, 72329, 72349, 72349, 72368, 72440, 72704,
		72712, 72714, 72750, 72768, 72768, 72818, 72847, 72960, 72966, 72968, 72969,
		72971, 73008, 73030, 73030, 73056, 73061, 73063, 73064, 73066, 73097, 73112,
		73112, 73440, 73458, 73474, 73474, 73476, 73488, 73490, 73523, 73648, 73648,
		73728, 74649, 74880, 75075, 77712, 77808, 77824, 78895, 78913, 78918, 82944,
		83526, 92160, 92728, 92736, 92766, 92784, 92862, 92880, 92909, 92928, 92975,
		92992, 92995, 93027, 93047, 93053, 93071, 93760, 93823, 93952, 94026, 94032,
		94032, 94099, 94111, 94176, 94177, 94179, 94179, 94208, 100343, 100352,
		101589, 101632, 101640, 110576, 110579, 110581, 110587, 110589, 110590,
		110592, 110882, 110898, 110898, 110928, 110930, 110933, 110933, 110948,
		110951, 110960, 111355, 113664, 113770, 113776, 113788, 113792, 113800,
		113808, 113817, 119808, 119892, 119894, 119964, 119966, 119967, 119970,
		119970, 119973, 119974, 119977, 119980, 119982, 119993, 119995, 119995,
		119997, 120003, 120005, 120069, 120071, 120074, 120077, 120084, 120086,
		120092, 120094, 120121, 120123, 120126, 120128, 120132, 120134, 120134,
		120138, 120144, 120146, 120485, 120488, 120512, 120514, 120538, 120540,
		120570, 120572, 120596, 120598, 120628, 120630, 120654, 120656, 120686,
		120688, 120712, 120714, 120744, 120746, 120770, 120772, 120779, 122624,
		122654, 122661, 122666, 122928, 122989, 123136, 123180, 123191, 123197,
		123214, 123214, 123536, 123565, 123584, 123627, 124112, 124139, 124896,
		124902, 124904, 124907, 124909, 124910, 124912, 124926, 124928, 125124,
		125184, 125251, 125259, 125259, 126464, 126467, 126469, 126495, 126497,
		126498, 126500, 126500, 126503, 126503, 126505, 126514, 126516, 126519,
		126521, 126521, 126523, 126523, 126530, 126530, 126535, 126535, 126537,
		126537, 126539, 126539, 126541, 126543, 126545, 126546, 126548, 126548,
		126551, 126551, 126553, 126553, 126555, 126555, 126557, 126557, 126559,
		126559, 126561, 126562, 126564, 126564, 126567, 126570, 126572, 126578,
		126580, 126583, 126585, 126588, 126590, 126590, 126592, 126601, 126603,
		126619, 126625, 126627, 126629, 126633, 126635, 126651, 131072, 173791,
		173824, 177977, 177984, 178205, 178208, 183969, 183984, 191456, 194560,
		195101, 196608, 201546, 201552, 205743, 748, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 170, 170, 178, 179, 181, 181, 185, 186, 188, 190, 192, 214, 216,
		246, 248, 705, 710, 721, 736, 740, 748, 748, 750, 750, 880, 884, 886, 887,
		890, 893, 895, 895, 902, 902, 904, 906, 908, 908, 910, 929, 931, 1013,
		1015, 1153, 1162, 1327, 1329, 1366, 1369, 1369, 1376, 1416, 1488, 1514,
		1519, 1522, 1568, 1610, 1632, 1641, 1646, 1647, 1649, 1747, 1749, 1749,
		1765, 1766, 1774, 1788, 1791, 1791, 1808, 1808, 1810, 1839, 1869, 1957,
		1969, 1969, 1984, 2026, 2036, 2037, 2042, 2042, 2048, 2069, 2074, 2074,
		2084, 2084, 2088, 2088, 2112, 2136, 2144, 2154, 2160, 2183, 2185, 2190,
		2208, 2249, 2308, 2361, 2365, 2365, 2384, 2384, 2392, 2401, 2406, 2415,
		2417, 2432, 2437, 2444, 2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482,
		2486, 2489, 2493, 2493, 2510, 2510, 2524, 2525, 2527, 2529, 2534, 2545,
		2548, 2553, 2556, 2556, 2565, 2570, 2575, 2576, 2579, 2600, 2602, 2608,
		2610, 2611, 2613, 2614, 2616, 2617, 2649, 2652, 2654, 2654, 2662, 2671,
		2674, 2676, 2693, 2701, 2703, 2705, 2707, 2728, 2730, 2736, 2738, 2739,
		2741, 2745, 2749, 2749, 2768, 2768, 2784, 2785, 2790, 2799, 2809, 2809,
		2821, 2828, 2831, 2832, 2835, 2856, 2858, 2864, 2866, 2867, 2869, 2873,
		2877, 2877, 2908, 2909, 2911, 2913, 2918, 2927, 2929, 2935, 2947, 2947,
		2949, 2954, 2958, 2960, 2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975,
		2979, 2980, 2984, 2986, 2990, 3001, 3024, 3024, 3046, 3058, 3077, 3084,
		3086, 3088, 3090, 3112, 3114, 3129, 3133, 3133, 3160, 3162, 3165, 3165,
		3168, 3169, 3174, 3183, 3192, 3198, 3200, 3200, 3205, 3212, 3214, 3216,
		3218, 3240, 3242, 3251, 3253, 3257, 3261, 3261, 3293, 3294, 3296, 3297,
		3302, 3311, 3313, 3314, 3332, 3340, 3342, 3344, 3346, 3386, 3389, 3389,
		3406, 3406, 3412, 3414, 3416, 3425, 3430, 3448, 3450, 3455, 3461, 3478,
		3482, 3505, 3507, 3515, 3517, 3517, 3520, 3526, 3558, 3567, 3585, 3632,
		3634, 3635, 3648, 3654, 3664, 3673, 3713, 3714, 3716, 3716, 3718, 3722,
		3724, 3747, 3749, 3749, 3751, 3760, 3762, 3763, 3773, 3773, 3776, 3780,
		3782, 3782, 3792, 3801, 3804, 3807, 3840, 3840, 3872, 3891, 3904, 3911,
		3913, 3948, 3976, 3980, 4096, 4138, 4159, 4169, 4176, 4181, 4186, 4189,
		4193, 4193, 4197, 4198, 4206, 4208, 4213, 4225, 4238, 4238, 4240, 4249,
		4256, 4293, 4295, 4295, 4301, 4301, 4304, 4346, 4348, 4680, 4682, 4685,
		4688, 4694, 4696, 4696, 4698, 4701, 4704, 4744, 4746, 4749, 4752, 4784,
		4786, 4789, 4792, 4798, 4800, 4800, 4802, 4805, 4808, 4822, 4824, 4880,
		4882, 4885, 4888, 4954, 4969, 4988, 4992, 5007, 5024, 5109, 5112, 5117,
		5121, 5740, 5743, 5759, 5761, 5786, 5792, 5866, 5870, 5880, 5888, 5905,
		5919, 5937, 5952, 5969, 5984, 5996, 5998, 6000, 6016, 6067, 6103, 6103,
		6108, 6108, 6112, 6121, 6128, 6137, 6160, 6169, 6176, 6264, 6272, 6276,
		6279, 6312, 6314, 6314, 6320, 6389, 6400, 6430, 6470, 6509, 6512, 6516,
		6528, 6571, 6576, 6601, 6608, 6618, 6656, 6678, 6688, 6740, 6784, 6793,
		6800, 6809, 6823, 6823, 6917, 6963, 6981, 6988, 6992, 7001, 7043, 7072,
		7086, 7141, 7168, 7203, 7232, 7241, 7245, 7293, 7296, 7304, 7312, 7354,
		7357, 7359, 7401, 7404, 7406, 7411, 7413, 7414, 7418, 7418, 7424, 7615,
		7680, 7957, 7960, 7965, 7968, 8005, 8008, 8013, 8016, 8023, 8025, 8025,
		8027, 8027, 8029, 8029, 8031, 8061, 8064, 8116, 8118, 8124, 8126, 8126,
		8130, 8132, 8134, 8140, 8144, 8147, 8150, 8155, 8160, 8172, 8178, 8180,
		8182, 8188, 8304, 8305, 8308, 8313, 8319, 8329, 8336, 8348, 8450, 8450,
		8455, 8455, 8458, 8467, 8469, 8469, 8473, 8477, 8484, 8484, 8486, 8486,
		8488, 8488, 8490, 8493, 8495, 8505, 8508, 8511, 8517, 8521, 8526, 8526,
		8528, 8585, 9312, 9371, 9450, 9471, 10102, 10131, 11264, 11492, 11499,
		11502, 11506, 11507, 11517, 11517, 11520, 11557, 11559, 11559, 11565, 11565,
		11568, 11623, 11631, 11631, 11648, 11670, 11680, 11686, 11688, 11694, 11696,
		11702, 11704, 11710, 11712, 11718, 11720, 11726, 11728, 11734, 11736, 11742,
		11823, 11823, 12293, 12295, 12321, 12329, 12337, 12341, 12344, 12348, 12353,
		12438, 12445, 12447, 12449, 12538, 12540, 12543, 12549, 12591, 12593, 12686,
		12690, 12693, 12704, 12735, 12784, 12799, 12832, 12841, 12872, 12879, 12881,
		12895, 12928, 12937, 12977, 12991, 13312, 19903, 19968, 42124, 42192, 42237,
		42240, 42508, 42512, 42539, 42560, 42606, 42623, 42653, 42656, 42735, 42775,
		42783, 42786, 42888, 42891, 42954, 42960, 42961, 42963, 42963, 42965, 42969,
		42994, 43009, 43011, 43013, 43015, 43018, 43020, 43042, 43056, 43061, 43072,
		43123, 43138, 43187, 43216, 43225, 43250, 43255, 43259, 43259, 43261, 43262,
		43264, 43301, 43312, 43334, 43360, 43388, 43396, 43442, 43471, 43481, 43488,
		43492, 43494, 43518, 43520, 43560, 43584, 43586, 43588, 43595, 43600, 43609,
		43616, 43638, 43642, 43642, 43646, 43695, 43697, 43697, 43701, 43702, 43705,
		43709, 43712, 43712, 43714, 43714, 43739, 43741, 43744, 43754, 43762, 43764,
		43777, 43782, 43785, 43790, 43793, 43798, 43808, 43814, 43816, 43822, 43824,
		43866, 43868, 43881, 43888, 44002, 44016, 44025, 44032, 55203, 55216, 55238,
		55243, 55291, 63744, 64109, 64112, 64217, 64256, 64262, 64275, 64279, 64285,
		64285, 64287, 64296, 64298, 64310, 64312, 64316, 64318, 64318, 64320, 64321,
		64323, 64324, 64326, 64433, 64467, 64829, 64848, 64911, 64914, 64967, 65008,
		65019, 65136, 65140, 65142, 65276, 65296, 65305, 65313, 65338, 65345, 65370,
		65382, 65470, 65474, 65479, 65482, 65487, 65490, 65495, 65498, 65500, 65536,
		65547, 65549, 65574, 65576, 65594, 65596, 65597, 65599, 65613, 65616, 65629,
		65664, 65786, 65799, 65843, 65856, 65912, 65930, 65931, 66176, 66204, 66208,
		66256, 66273, 66299, 66304, 66339, 66349, 66378, 66384, 66421, 66432, 66461,
		66464, 66499, 66504, 66511, 66513, 66517, 66560, 66717, 66720, 66729, 66736,
		66771, 66776, 66811, 66816, 66855, 66864, 66915, 66928, 66938, 66940, 66954,
		66956, 66962, 66964, 66965, 66967, 66977, 66979, 66993, 66995, 67001, 67003,
		67004, 67072, 67382, 67392, 67413, 67424, 67431, 67456, 67461, 67463, 67504,
		67506, 67514, 67584, 67589, 67592, 67592, 67594, 67637, 67639, 67640, 67644,
		67644, 67647, 67669, 67672, 67702, 67705, 67742, 67751, 67759, 67808, 67826,
		67828, 67829, 67835, 67867, 67872, 67897, 67968, 68023, 68028, 68047, 68050,
		68096, 68112, 68115, 68117, 68119, 68121, 68149, 68160, 68168, 68192, 68222,
		68224, 68255, 68288, 68295, 68297, 68324, 68331, 68335, 68352, 68405, 68416,
		68437, 68440, 68466, 68472, 68497, 68521, 68527, 68608, 68680, 68736, 68786,
		68800, 68850, 68858, 68899, 68912, 68921, 69216, 69246, 69248, 69289, 69296,
		69297, 69376, 69415, 69424, 69445, 69457, 69460, 69488, 69505, 69552, 69579,
		69600, 69622, 69635, 69687, 69714, 69743, 69745, 69746, 69749, 69749, 69763,
		69807, 69840, 69864, 69872, 69881, 69891, 69926, 69942, 69951, 69956, 69956,
		69959, 69959, 69968, 70002, 70006, 70006, 70019, 70066, 70081, 70084, 70096,
		70106, 70108, 70108, 70113, 70132, 70144, 70161, 70163, 70187, 70207, 70208,
		70272, 70278, 70280, 70280, 70282, 70285, 70287, 70301, 70303, 70312, 70320,
		70366, 70384, 70393, 70405, 70412, 70415, 70416, 70419, 70440, 70442, 70448,
		70450, 70451, 70453, 70457, 70461, 70461, 70480, 70480, 70493, 70497, 70656,
		70708, 70727, 70730, 70736, 70745, 70751, 70753, 70784, 70831, 70852, 70853,
		70855, 70855, 70864, 70873, 71040, 71086, 71128, 71131, 71168, 71215, 71236,
		71236, 71248, 71257, 71296, 71338, 71352, 71352, 71360, 71369, 71424, 71450,
		71472, 71483, 71488, 71494, 71680, 71723, 71840, 71922, 71935, 71942, 71945,
		71945, 71948, 71955, 71957, 71958, 71960, 71983, 71999, 71999, 72001, 72001,
		72016, 72025, 72096, 72103, 72106, 72144, 72161, 72161, 72163, 72163, 72192,
		72192, 72203, 72242, 72250, 72250, 72272, 72272, 72284, 72329, 72349, 72349,
		72368, 72440, 72704, 72712, 72714, 72750, 72768, 72768, 72784, 72812, 72818,
		72847, 72960, 72966, 72968, 72969, 72971, 73008, 73030, 73030, 73040, 73049,
		73056, 73061, 73063, 73064, 73066, 73097, 73112, 73112, 73120, 73129, 73440,
		73458, 73474, 73474, 73476, 73488, 73490, 73523, 73552, 73561, 73648, 73648,
		73664, 73684, 73728, 74649, 74752, 74862, 74880, 75075, 77712, 77808, 77824,
		78895, 78913, 78918, 82944, 83526, 92160, 92728, 92736, 92766, 92768, 92777,
		92784, 92862, 92864, 92873, 92880, 92909, 92928, 92975, 92992, 92995, 93008,
		93017, 93019, 93025, 93027, 93047, 93053, 93071, 93760, 93846, 93952, 94026,
		94032, 94032, 94099, 94111, 94176, 94177, 94179, 94179, 94208, 100343,
		100352, 101589, 101632, 101640, 110576, 110579, 110581, 110587, 110589,
		110590, 110592, 110882, 110898, 110898, 110928, 110930, 110933, 110933,
		110948, 110951, 110960, 111355, 113664, 113770, 113776, 113788, 113792,
		113800, 113808, 113817, 119488, 119507, 119520, 119539, 119648, 119672,
		119808, 119892, 119894, 119964, 119966, 119967, 119970, 119970, 119973,
		119974, 119977, 119980, 119982, 119993, 119995, 119995, 119997, 120003,
		120005, 120069, 120071, 120074, 120077, 120084, 120086, 120092, 120094,
		120121, 120123, 120126, 120128, 120132, 120134, 120134, 120138, 120144,
		120146, 120485, 120488, 120512, 120514, 120538, 120540, 120570, 120572,
		120596, 120598, 120628, 120630, 120654, 120656, 120686, 120688, 120712,
		120714, 120744, 120746, 120770, 120772, 120779, 120782, 120831, 122624,
		122654, 122661, 122666, 122928, 122989, 123136, 123180, 123191, 123197,
		123200, 123209, 123214, 123214, 123536, 123565, 123584, 123627, 123632,
		123641, 124112, 124139, 124144, 124153, 124896, 124902, 124904, 124907,
		124909, 124910, 124912, 124926, 124928, 125124, 125127, 125135, 125184,
		125251, 125259, 125259, 125264, 125273, 126065, 126123, 126125, 126127,
		126129, 126132, 126209, 126253, 126255, 126269, 126464, 126467, 126469,
		126495, 126497, 126498, 126500, 126500, 126503, 126503, 126505, 126514,
		126516, 126519, 126521, 126521, 126523, 126523, 126530, 126530, 126535,
		126535, 126537, 126537, 126539, 126539, 126541, 126543, 126545, 126546,
		126548, 126548, 126551, 126551, 126553, 126553, 126555, 126555, 126557,
		126557, 126559, 126559, 126561, 126562, 126564, 126564, 126567, 126570,
		126572, 126578, 126580, 126583, 126585, 126588, 126590, 126590, 126592,
		126601, 126603, 126619, 126625, 126627, 126629, 126633, 126635, 126651,
		127232, 127244, 130032, 130041, 131072, 173791, 173824, 177977, 177984,
		178205, 178208, 183969, 183984, 191456, 194560, 195101, 196608, 201546,
		201552, 205743, 394, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0,
		0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0,
		0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1,
		0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75,
		1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0,
		83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0,
		0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0,
		0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1,
		0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0,
		123, 1, 0, 0, 0, 1, 125, 1, 0, 0, 0, 3, 127, 1, 0, 0, 0, 5, 129, 1, 0,
		0, 0, 7, 131, 1, 0, 0, 0, 9, 133, 1, 0, 0, 0, 11, 135, 1, 0, 0, 0, 13,
		137, 1, 0, 0, 0, 15, 140, 1, 0, 0, 0, 17, 145, 1, 0, 0, 0, 19, 149, 1,
		0, 0, 0, 21, 156, 1, 0, 0, 0, 23, 161, 1, 0, 0, 0, 25, 169, 1, 0, 0, 0,
		27, 181, 1, 0, 0, 0, 29, 187, 1, 0, 0, 0, 31, 196, 1, 0, 0, 0, 33, 203,
		1, 0, 0, 0, 35, 209, 1, 0, 0, 0, 37, 212, 1, 0, 0, 0, 39, 215, 1, 0, 0,
		0, 41, 220, 1, 0, 0, 0, 43, 224, 1, 0, 0, 0, 45, 230, 1, 0, 0, 0, 47, 235,
		1, 0, 0, 0, 49, 240, 1, 0, 0, 0, 51, 247, 1, 0, 0, 0, 53, 253, 1, 0, 0,
		0, 55, 259, 1, 0, 0, 0, 57, 269, 1, 0, 0, 0, 59, 271, 1, 0, 0, 0, 61, 273,
		1, 0, 0, 0, 63, 275, 1, 0, 0, 0, 65, 277, 1, 0, 0, 0, 67, 279, 1, 0, 0,
		0, 69, 281, 1, 0, 0, 0, 71, 283, 1, 0, 0, 0, 73, 285, 1, 0, 0, 0, 75, 287,
		1, 0, 0, 0, 77, 289, 1, 0, 0, 0, 79, 291, 1, 0, 0, 0, 81, 293, 1, 0, 0,
		0, 83, 295, 1, 0, 0, 0, 85, 297, 1, 0, 0, 0, 87, 299, 1, 0, 0, 0, 89, 303,
		1, 0, 0, 0, 91, 305, 1, 0, 0, 0, 93, 308, 1, 0, 0, 0, 95, 311, 1, 0, 0,
		0, 97, 314, 1, 0, 0, 0, 99, 317, 1, 0, 0, 0, 101, 320, 1, 0, 0, 0, 103,
		322, 1, 0, 0, 0, 105, 325, 1, 0, 0, 0, 107, 328, 1, 0, 0, 0, 109, 348,
		1, 0, 0, 0, 111, 350, 1, 0, 0, 0, 113, 358, 1, 0, 0, 0, 115, 362, 1, 0,
		0, 0, 117, 367, 1, 0, 0, 0, 119, 371, 1, 0, 0, 0, 121, 373, 1, 0, 0, 0,
		123, 387, 1, 0, 0, 0, 125, 126, 5, 46, 0, 0, 126, 2, 1, 0, 0, 0, 127, 128,
		5, 44, 0, 0, 128, 4, 1, 0, 0, 0, 129, 130, 5, 58, 0, 0, 130, 6, 1, 0, 0,
		0, 131, 132, 5, 63, 0, 0, 132, 8, 1, 0, 0, 0, 133, 134, 5, 59, 0, 0, 134,
		10, 1, 0, 0, 0, 135, 136, 5, 61, 0, 0, 136, 12, 1, 0, 0, 0, 137, 138, 5,
		105, 0, 0, 138, 139, 5, 102, 0, 0, 139, 14, 1, 0, 0, 0, 140, 141, 5, 101,
		0, 0, 141, 142, 5, 108, 0, 0, 142, 143, 5, 115, 0, 0, 143, 144, 5, 101,
		0, 0, 144, 16, 1, 0, 0, 0, 145, 146, 5, 102, 0, 0, 146, 147, 5, 111, 0,
		0, 147, 148, 5, 114, 0, 0, 148, 18, 1, 0, 0, 0, 149, 150, 5, 115, 0, 0,
		150, 151, 5, 119, 0, 0, 151, 152, 5, 105, 0, 0, 152, 153, 5, 116, 0, 0,
		153, 154, 5, 99, 0, 0, 154, 155, 5, 104, 0, 0, 155, 20, 1, 0, 0, 0, 156,
		157, 5, 99, 0, 0, 157, 158, 5, 97, 0, 0, 158, 159, 5, 115, 0, 0, 159, 160,
		5, 101, 0, 0, 160, 22, 1, 0, 0, 0, 161, 162, 5, 100, 0, 0, 162, 163, 5,
		101, 0, 0, 163, 164, 5, 102, 0, 0, 164, 165, 5, 97, 0, 0, 165, 166, 5,
		117, 0, 0, 166, 167, 5, 108, 0, 0, 167, 168, 5, 116, 0, 0, 168, 24, 1,
		0, 0, 0, 169, 170, 5, 102, 0, 0, 170, 171, 5, 97, 0, 0, 171, 172, 5, 108,
		0, 0, 172, 173, 5, 108, 0, 0, 173, 174, 5, 116, 0, 0, 174, 175, 5, 104,
		0, 0, 175, 176, 5, 114, 0, 0, 176, 177, 5, 111, 0, 0, 177, 178, 5, 117,
		0, 0, 178, 179, 5, 103, 0, 0, 179, 180, 5, 104, 0, 0, 180, 26, 1, 0, 0,
		0, 181, 182, 5, 98, 0, 0, 182, 183, 5, 114, 0, 0, 183, 184, 5, 101, 0,
		0, 184, 185, 5, 97, 0, 0, 185, 186, 5, 107, 0, 0, 186, 28, 1, 0, 0, 0,
		187, 188, 5, 99, 0, 0, 188, 189, 5, 111, 0, 0, 189, 190, 5, 110, 0, 0,
		190, 191, 5, 116, 0, 0, 191, 192, 5, 105, 0, 0, 192, 193, 5, 110, 0, 0,
		193, 194, 5, 117, 0, 0, 194, 195, 5, 101, 0, 0, 195, 30, 1, 0, 0, 0, 196,
		197, 5, 114, 0, 0, 197, 198, 5, 101, 0, 0, 198, 199, 5, 116, 0, 0, 199,
		200, 5, 117, 0, 0, 200, 201, 5, 114, 0, 0, 201, 202, 5, 110, 0, 0, 202,
		32, 1, 0, 0, 0, 203, 204, 5, 117, 0, 0, 204, 205, 5, 115, 0, 0, 205, 206,
		5, 105, 0, 0, 206, 207, 5, 110, 0, 0, 207, 208, 5, 103, 0, 0, 208, 34,
		1, 0, 0, 0, 209, 210, 5, 97, 0, 0, 210, 211, 5, 115, 0, 0, 211, 36, 1,
		0, 0, 0, 212, 213, 5, 105, 0, 0, 213, 214, 5, 110, 0, 0, 214, 38, 1, 0,
		0, 0, 215, 216, 5, 110, 0, 0, 216, 217, 5, 117, 0, 0, 217, 218, 5, 108,
		0, 0, 218, 219, 5, 108, 0, 0, 219, 40, 1, 0, 0, 0, 220, 221, 5, 118, 0,
		0, 221, 222, 5, 97, 0, 0, 222, 223, 5, 114, 0, 0, 223, 42, 1, 0, 0, 0,
		224, 225, 5, 99, 0, 0, 225, 226, 5, 111, 0, 0, 226, 227, 5, 110, 0, 0,
		227, 228, 5, 115, 0, 0, 228, 229, 5, 116, 0, 0, 229, 44, 1, 0, 0, 0, 230,
		231, 5, 116, 0, 0, 231, 232, 5, 121, 0, 0, 232, 233, 5, 112, 0, 0, 233,
		234, 5, 101, 0, 0, 234, 46, 1, 0, 0, 0, 235, 236, 5, 102, 0, 0, 236, 237,
		5, 117, 0, 0, 237, 238, 5, 110, 0, 0, 238, 239, 5, 99, 0, 0, 239, 48, 1,
		0, 0, 0, 240, 241, 5, 115, 0, 0, 241, 242, 5, 116, 0, 0, 242, 243, 5, 114,
		0, 0, 243, 244, 5, 117, 0, 0, 244, 245, 5, 99, 0, 0, 245, 246, 5, 116,
		0, 0, 246, 50, 1, 0, 0, 0, 247, 248, 5, 116, 0, 0, 248, 249, 5, 117, 0,
		0, 249, 250, 5, 112, 0, 0, 250, 251, 5, 108, 0, 0, 251, 252, 5, 101, 0,
		0, 252, 52, 1, 0, 0, 0, 253, 254, 5, 99, 0, 0, 254, 255, 5, 108, 0, 0,
		255, 256, 5, 97, 0, 0, 256, 257, 5, 115, 0, 0, 257, 258, 5, 115, 0, 0,
		258, 54, 1, 0, 0, 0, 259, 260, 5, 105, 0, 0, 260, 261, 5, 110, 0, 0, 261,
		262, 5, 116, 0, 0, 262, 263, 5, 101, 0, 0, 263, 264, 5, 114, 0, 0, 264,
		265, 5, 102, 0, 0, 265, 266, 5, 97, 0, 0, 266, 267, 5, 99, 0, 0, 267, 268,
		5, 101, 0, 0, 268, 56, 1, 0, 0, 0, 269, 270, 5, 40, 0, 0, 270, 58, 1, 0,
		0, 0, 271, 272, 5, 41, 0, 0, 272, 60, 1, 0, 0, 0, 273, 274, 5, 91, 0, 0,
		274, 62, 1, 0, 0, 0, 275, 276, 5, 93, 0, 0, 276, 64, 1, 0, 0, 0, 277, 278,
		5, 123, 0, 0, 278, 66, 1, 0, 0, 0, 279, 280, 5, 125, 0, 0, 280, 68, 1,
		0, 0, 0, 281, 282, 5, 43, 0, 0, 282, 70, 1, 0, 0, 0, 283, 284, 5, 45, 0,
		0, 284, 72, 1, 0, 0, 0, 285, 286, 5, 33, 0, 0, 286, 74, 1, 0, 0, 0, 287,
		288, 5, 42, 0, 0, 288, 76, 1, 0, 0, 0, 289, 290, 5, 47, 0, 0, 290, 78,
		1, 0, 0, 0, 291, 292, 5, 37, 0, 0, 292, 80, 1, 0, 0, 0, 293, 294, 5, 94,
		0, 0, 294, 82, 1, 0, 0, 0, 295, 296, 5, 126, 0, 0, 296, 84, 1, 0, 0, 0,
		297, 298, 5, 124, 0, 0, 298, 86, 1, 0, 0, 0, 299, 300, 5, 124, 0, 0, 300,
		301, 5, 33, 0, 0, 301, 302, 5, 38, 0, 0, 302, 88, 1, 0, 0, 0, 303, 304,
		5, 38, 0, 0, 304, 90, 1, 0, 0, 0, 305, 306, 5, 60, 0, 0, 306, 307, 5, 60,
		0, 0, 307, 92, 1, 0, 0, 0, 308, 309, 5, 62, 0, 0, 309, 310, 5, 62, 0, 0,
		310, 94, 1, 0, 0, 0, 311, 312, 5, 61, 0, 0, 312, 313, 5, 61, 0, 0, 313,
		96, 1, 0, 0, 0, 314, 315, 5, 33, 0, 0, 315, 316, 5, 61, 0, 0, 316, 98,
		1, 0, 0, 0, 317, 318, 5, 62, 0, 0, 318, 319, 5, 61, 0, 0, 319, 100, 1,
		0, 0, 0, 320, 321, 5, 62, 0, 0, 321, 102, 1, 0, 0, 0, 322, 323, 5, 60,
		0, 0, 323, 324, 5, 61, 0, 0, 324, 104, 1, 0, 0, 0, 325, 326, 5, 60, 0,
		0, 326, 106, 1, 0, 0, 0, 327, 329, 7, 0, 0, 0, 328, 327, 1, 0, 0, 0, 329,
		330, 1, 0, 0, 0, 330, 328, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 332,
		1, 0, 0, 0, 332, 333, 6, 53, 0, 0, 333, 108, 1, 0, 0, 0, 334, 338, 5, 96,
		0, 0, 335, 337, 9, 0, 0, 0, 336, 335, 1, 0, 0, 0, 337, 340, 1, 0, 0, 0,
		338, 339, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 339, 341, 1, 0, 0, 0, 340,
		338, 1, 0, 0, 0, 341, 349, 5, 96, 0, 0, 342, 343, 5, 35, 0, 0, 343, 345,
		8, 1, 0, 0, 344, 346, 5, 13, 0, 0, 345, 344, 1, 0, 0, 0, 345, 346, 1, 0,
		0, 0, 346, 347, 1, 0, 0, 0, 347, 349, 5, 10, 0, 0, 348, 334, 1, 0, 0, 0,
		348, 342, 1, 0, 0, 0, 349, 110, 1, 0, 0, 0, 350, 354, 7, 2, 0, 0, 351,
		353, 7, 3, 0, 0, 352, 351, 1, 0, 0, 0, 353, 356, 1, 0, 0, 0, 354, 352,
		1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 112, 1, 0, 0, 0, 356, 354, 1, 0,
		0, 0, 357, 359, 5, 13, 0, 0, 358, 357, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0,
		359, 360, 1, 0, 0, 0, 360, 361, 5, 10, 0, 0, 361, 114, 1, 0, 0, 0, 362,
		363, 5, 92, 0, 0, 363, 364, 9, 0, 0, 0, 364, 116, 1, 0, 0, 0, 365, 368,
		8, 1, 0, 0, 366, 368, 3, 115, 57, 0, 367, 365, 1, 0, 0, 0, 367, 366, 1,
		0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370, 1, 0, 0,
		0, 370, 118, 1, 0, 0, 0, 371, 372, 5, 39, 0, 0, 372, 120, 1, 0, 0, 0, 373,
		374, 5, 34, 0, 0, 374, 122, 1, 0, 0, 0, 375, 377, 3, 119, 59, 0, 376, 378,
		3, 117, 58, 0, 377, 376, 1, 0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 379, 1,
		0, 0, 0, 379, 380, 3, 119, 59, 0, 380, 388, 1, 0, 0, 0, 381, 383, 3, 121,
		60, 0, 382, 384, 3, 117, 58, 0, 383, 382, 1, 0, 0, 0, 383, 384, 1, 0, 0,
		0, 384, 385, 1, 0, 0, 0, 385, 386, 3, 121, 60, 0, 386, 388, 1, 0, 0, 0,
		387, 375, 1, 0, 0, 0, 387, 381, 1, 0, 0, 0, 388, 124, 1, 0, 0, 0, 12, 0,
		330, 338, 345, 348, 354, 358, 367, 369, 377, 383, 387, 1, 6, 0, 0,
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

// CaliburnLexerInit initializes any static state used to implement CaliburnLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCaliburnLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CaliburnLexerInit() {
	staticData := &CaliburnLexerLexerStaticData
	staticData.once.Do(caliburnlexerLexerInit)
}

// NewCaliburnLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCaliburnLexer(input antlr.CharStream) *CaliburnLexer {
	CaliburnLexerInit()
	l := new(CaliburnLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CaliburnLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "CaliburnLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CaliburnLexer tokens.
const (
	CaliburnLexerPERIOD      = 1
	CaliburnLexerCOMMA       = 2
	CaliburnLexerCOLON       = 3
	CaliburnLexerQUESTION    = 4
	CaliburnLexerTerminator  = 5
	CaliburnLexerASSIGN      = 6
	CaliburnLexerIF          = 7
	CaliburnLexerELSE        = 8
	CaliburnLexerFOR         = 9
	CaliburnLexerSWITCH      = 10
	CaliburnLexerCASE        = 11
	CaliburnLexerDEFAULT     = 12
	CaliburnLexerFALLTHROUGH = 13
	CaliburnLexerBREAK       = 14
	CaliburnLexerCONTINUE    = 15
	CaliburnLexerRETURN      = 16
	CaliburnLexerUSING       = 17
	CaliburnLexerAS          = 18
	CaliburnLexerIN          = 19
	CaliburnLexerNULL        = 20
	CaliburnLexerVAR         = 21
	CaliburnLexerCONST       = 22
	CaliburnLexerTYPE        = 23
	CaliburnLexerFUNC        = 24
	CaliburnLexerSTRUCT      = 25
	CaliburnLexerTUPLE       = 26
	CaliburnLexerCLASS       = 27
	CaliburnLexerINTERFACE   = 28
	CaliburnLexerL_PAREN     = 29
	CaliburnLexerR_PAREN     = 30
	CaliburnLexerL_S_BRACK   = 31
	CaliburnLexerR_S_BRACK   = 32
	CaliburnLexerL_C_BRACK   = 33
	CaliburnLexerR_C_BRACK   = 34
	CaliburnLexerOP_ADD      = 35
	CaliburnLexerOP_SUB      = 36
	CaliburnLexerOP_NOT      = 37
	CaliburnLexerOP_MUL      = 38
	CaliburnLexerOP_DIV      = 39
	CaliburnLexerOP_MOD      = 40
	CaliburnLexerOP_POW      = 41
	CaliburnLexerOP_ROOT     = 42
	CaliburnLexerOP_OR       = 43
	CaliburnLexerOP_XOR      = 44
	CaliburnLexerOP_AND      = 45
	CaliburnLexerOP_LSHIFT   = 46
	CaliburnLexerOP_RSHIFT   = 47
	CaliburnLexerOP_EQU      = 48
	CaliburnLexerOP_NEQ      = 49
	CaliburnLexerOP_GTE      = 50
	CaliburnLexerOP_GRT      = 51
	CaliburnLexerOP_LTE      = 52
	CaliburnLexerOP_LST      = 53
	CaliburnLexerWhiteSpace  = 54
	CaliburnLexerComment     = 55
	CaliburnLexerIdentifier  = 56
	CaliburnLexerLiteral     = 57
)

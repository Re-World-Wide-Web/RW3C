// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // RW3C

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type RW3CParser struct {
	*antlr.BaseParser
}

var rw3cParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func rw3cParserInit() {
	staticData := &rw3cParserStaticData
	staticData.literalNames = []string{
		"", "'break'", "';'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'>'",
		"'<'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'", "'++'", "'--'",
		"'!'", "'if '", "'else '", "'switch '", "'{'", "'case '", "':'", "'default'",
		"'}'", "'while '", "'ret'", "' '", "'let '", "'?'", "'='", "'mut '",
		"'type '", "'struct '", "','", "'.'", "'['", "']'", "'fn '", "'fn'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "RESERVED", "IDENTIF_PART", "IDENTIF_START",
		"IDENTIF", "NEWLINE", "LONG", "DOUBLE", "STR", "ESC", "BOOL", "MULTI_COMMENT",
		"SINGLE_COMMENT",
	}
	staticData.ruleNames = []string{
		"prog", "run", "expr", "chain_expr", "atomic_expr", "stmt", "if_stmt",
		"if_expr", "switch_stmt", "switch_expr", "while_stmt", "while_expr",
		"ret_stmt", "def_var_stmt", "def_type_stmt", "struct_stmt", "assign_stmt",
		"call_expr", "access_prop_expr", "fn_expr", "scope", "access_prop_type_expr",
		"type_expr", "type_arg", "name_expr", "arg", "lit_expr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 54, 393, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 1, 0, 5, 0, 56, 8, 0, 10, 0, 12, 0, 59, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 3, 1, 66, 8, 1, 1, 1, 5, 1, 69, 8, 1, 10, 1, 12, 1, 72, 9,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 3, 2, 86, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 102, 8, 2, 10, 2, 12, 2, 105, 9, 2,
		1, 3, 1, 3, 1, 3, 5, 3, 110, 8, 3, 10, 3, 12, 3, 113, 9, 3, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 121, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 5, 6, 140, 8, 6, 10, 6, 12, 6, 143, 9, 6, 1, 6, 1, 6, 3, 6, 147,
		8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 154, 8, 7, 10, 7, 12, 7, 157,
		9, 7, 1, 7, 1, 7, 3, 7, 161, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 174, 8, 8, 10, 8, 12, 8, 177, 9, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 5, 9, 192, 8, 9, 10, 9, 12, 9, 195, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 3, 12, 210,
		8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 216, 8, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 222, 8, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 228, 8,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 235, 8, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 3, 14, 241, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 253, 8, 15, 10, 15, 12, 15, 256,
		9, 15, 1, 15, 3, 15, 259, 8, 15, 3, 15, 261, 8, 15, 1, 15, 3, 15, 264,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 274,
		8, 17, 10, 17, 12, 17, 277, 9, 17, 1, 17, 3, 17, 280, 8, 17, 3, 17, 282,
		8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 292,
		8, 18, 1, 19, 1, 19, 3, 19, 296, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5,
		19, 302, 8, 19, 10, 19, 12, 19, 305, 9, 19, 1, 19, 3, 19, 308, 8, 19, 3,
		19, 310, 8, 19, 1, 19, 1, 19, 1, 19, 3, 19, 315, 8, 19, 1, 19, 3, 19, 318,
		8, 19, 1, 19, 1, 19, 3, 19, 322, 8, 19, 1, 20, 1, 20, 5, 20, 326, 8, 20,
		10, 20, 12, 20, 329, 9, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 3, 21, 339, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		5, 22, 347, 8, 22, 10, 22, 12, 22, 350, 9, 22, 1, 22, 3, 22, 353, 8, 22,
		3, 22, 355, 8, 22, 1, 22, 1, 22, 1, 22, 3, 22, 360, 8, 22, 1, 22, 3, 22,
		363, 8, 22, 3, 22, 365, 8, 22, 1, 22, 5, 22, 368, 8, 22, 10, 22, 12, 22,
		371, 9, 22, 1, 23, 1, 23, 3, 23, 375, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 3, 25, 382, 8, 25, 1, 25, 3, 25, 385, 8, 25, 1, 25, 1, 25, 3, 25,
		389, 8, 25, 1, 26, 1, 26, 1, 26, 0, 1, 4, 27, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 0, 6, 1, 0, 17, 18, 1, 0, 7, 8, 1, 0, 5, 6, 1, 0, 9, 14, 1, 0, 15,
		16, 3, 0, 46, 46, 48, 50, 52, 52, 436, 0, 57, 1, 0, 0, 0, 2, 65, 1, 0,
		0, 0, 4, 85, 1, 0, 0, 0, 6, 106, 1, 0, 0, 0, 8, 120, 1, 0, 0, 0, 10, 132,
		1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 148, 1, 0, 0, 0, 16, 162, 1, 0, 0,
		0, 18, 180, 1, 0, 0, 0, 20, 198, 1, 0, 0, 0, 22, 202, 1, 0, 0, 0, 24, 206,
		1, 0, 0, 0, 26, 234, 1, 0, 0, 0, 28, 236, 1, 0, 0, 0, 30, 242, 1, 0, 0,
		0, 32, 265, 1, 0, 0, 0, 34, 269, 1, 0, 0, 0, 36, 291, 1, 0, 0, 0, 38, 293,
		1, 0, 0, 0, 40, 323, 1, 0, 0, 0, 42, 338, 1, 0, 0, 0, 44, 364, 1, 0, 0,
		0, 46, 372, 1, 0, 0, 0, 48, 376, 1, 0, 0, 0, 50, 378, 1, 0, 0, 0, 52, 390,
		1, 0, 0, 0, 54, 56, 3, 2, 1, 0, 55, 54, 1, 0, 0, 0, 56, 59, 1, 0, 0, 0,
		57, 55, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 60, 1, 0, 0, 0, 59, 57, 1,
		0, 0, 0, 60, 61, 5, 0, 0, 1, 61, 1, 1, 0, 0, 0, 62, 66, 3, 10, 5, 0, 63,
		66, 3, 4, 2, 0, 64, 66, 5, 1, 0, 0, 65, 62, 1, 0, 0, 0, 65, 63, 1, 0, 0,
		0, 65, 64, 1, 0, 0, 0, 66, 70, 1, 0, 0, 0, 67, 69, 5, 2, 0, 0, 68, 67,
		1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 3, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74, 6, 2, -1, 0, 74, 75, 5,
		3, 0, 0, 75, 76, 3, 4, 2, 0, 76, 77, 5, 4, 0, 0, 77, 86, 1, 0, 0, 0, 78,
		79, 7, 0, 0, 0, 79, 86, 3, 4, 2, 5, 80, 81, 7, 1, 0, 0, 81, 86, 3, 4, 2,
		4, 82, 83, 5, 19, 0, 0, 83, 86, 3, 4, 2, 3, 84, 86, 3, 6, 3, 0, 85, 73,
		1, 0, 0, 0, 85, 78, 1, 0, 0, 0, 85, 80, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0,
		85, 84, 1, 0, 0, 0, 86, 103, 1, 0, 0, 0, 87, 88, 10, 9, 0, 0, 88, 89, 7,
		2, 0, 0, 89, 102, 3, 4, 2, 10, 90, 91, 10, 8, 0, 0, 91, 92, 7, 1, 0, 0,
		92, 102, 3, 4, 2, 9, 93, 94, 10, 7, 0, 0, 94, 95, 7, 3, 0, 0, 95, 102,
		3, 4, 2, 8, 96, 97, 10, 6, 0, 0, 97, 98, 7, 4, 0, 0, 98, 102, 3, 4, 2,
		7, 99, 100, 10, 2, 0, 0, 100, 102, 7, 0, 0, 0, 101, 87, 1, 0, 0, 0, 101,
		90, 1, 0, 0, 0, 101, 93, 1, 0, 0, 0, 101, 96, 1, 0, 0, 0, 101, 99, 1, 0,
		0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0,
		104, 5, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 111, 3, 8, 4, 0, 107, 110,
		3, 34, 17, 0, 108, 110, 3, 36, 18, 0, 109, 107, 1, 0, 0, 0, 109, 108, 1,
		0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0,
		0, 112, 7, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 121, 3, 38, 19, 0, 115,
		121, 3, 52, 26, 0, 116, 121, 3, 40, 20, 0, 117, 121, 3, 18, 9, 0, 118,
		121, 3, 14, 7, 0, 119, 121, 3, 22, 11, 0, 120, 114, 1, 0, 0, 0, 120, 115,
		1, 0, 0, 0, 120, 116, 1, 0, 0, 0, 120, 117, 1, 0, 0, 0, 120, 118, 1, 0,
		0, 0, 120, 119, 1, 0, 0, 0, 121, 9, 1, 0, 0, 0, 122, 133, 5, 53, 0, 0,
		123, 133, 5, 54, 0, 0, 124, 133, 3, 32, 16, 0, 125, 133, 3, 26, 13, 0,
		126, 133, 3, 28, 14, 0, 127, 133, 3, 30, 15, 0, 128, 133, 3, 24, 12, 0,
		129, 133, 3, 12, 6, 0, 130, 133, 3, 20, 10, 0, 131, 133, 3, 16, 8, 0, 132,
		122, 1, 0, 0, 0, 132, 123, 1, 0, 0, 0, 132, 124, 1, 0, 0, 0, 132, 125,
		1, 0, 0, 0, 132, 126, 1, 0, 0, 0, 132, 127, 1, 0, 0, 0, 132, 128, 1, 0,
		0, 0, 132, 129, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0,
		133, 11, 1, 0, 0, 0, 134, 135, 5, 20, 0, 0, 135, 136, 3, 4, 2, 0, 136,
		141, 3, 2, 1, 0, 137, 138, 5, 21, 0, 0, 138, 140, 3, 12, 6, 0, 139, 137,
		1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0,
		0, 0, 142, 146, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 21, 0, 0,
		145, 147, 3, 2, 1, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147,
		13, 1, 0, 0, 0, 148, 149, 5, 20, 0, 0, 149, 150, 3, 4, 2, 0, 150, 155,
		3, 4, 2, 0, 151, 152, 5, 21, 0, 0, 152, 154, 3, 14, 7, 0, 153, 151, 1,
		0, 0, 0, 154, 157, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 155, 156, 1, 0, 0,
		0, 156, 160, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 158, 159, 5, 21, 0, 0, 159,
		161, 3, 4, 2, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 15, 1,
		0, 0, 0, 162, 163, 5, 22, 0, 0, 163, 164, 3, 4, 2, 0, 164, 175, 5, 23,
		0, 0, 165, 166, 5, 24, 0, 0, 166, 167, 3, 4, 2, 0, 167, 168, 5, 25, 0,
		0, 168, 169, 3, 2, 1, 0, 169, 174, 1, 0, 0, 0, 170, 171, 5, 26, 0, 0, 171,
		172, 5, 25, 0, 0, 172, 174, 3, 2, 1, 0, 173, 165, 1, 0, 0, 0, 173, 170,
		1, 0, 0, 0, 174, 177, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 175, 176, 1, 0,
		0, 0, 176, 178, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 178, 179, 5, 27, 0, 0,
		179, 17, 1, 0, 0, 0, 180, 181, 5, 22, 0, 0, 181, 182, 3, 4, 2, 0, 182,
		193, 5, 23, 0, 0, 183, 184, 5, 24, 0, 0, 184, 185, 3, 4, 2, 0, 185, 186,
		5, 25, 0, 0, 186, 187, 3, 4, 2, 0, 187, 192, 1, 0, 0, 0, 188, 189, 5, 26,
		0, 0, 189, 190, 5, 25, 0, 0, 190, 192, 3, 4, 2, 0, 191, 183, 1, 0, 0, 0,
		191, 188, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193,
		194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 196, 197,
		5, 27, 0, 0, 197, 19, 1, 0, 0, 0, 198, 199, 5, 28, 0, 0, 199, 200, 3, 4,
		2, 0, 200, 201, 3, 2, 1, 0, 201, 21, 1, 0, 0, 0, 202, 203, 5, 28, 0, 0,
		203, 204, 3, 4, 2, 0, 204, 205, 3, 4, 2, 0, 205, 23, 1, 0, 0, 0, 206, 209,
		5, 29, 0, 0, 207, 208, 5, 30, 0, 0, 208, 210, 3, 4, 2, 0, 209, 207, 1,
		0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 25, 1, 0, 0, 0, 211, 212, 5, 31, 0,
		0, 212, 213, 3, 48, 24, 0, 213, 215, 5, 25, 0, 0, 214, 216, 5, 32, 0, 0,
		215, 214, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217,
		218, 3, 44, 22, 0, 218, 221, 1, 0, 0, 0, 219, 220, 5, 33, 0, 0, 220, 222,
		3, 4, 2, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222, 235, 1, 0,
		0, 0, 223, 224, 5, 34, 0, 0, 224, 225, 3, 48, 24, 0, 225, 227, 5, 25, 0,
		0, 226, 228, 5, 32, 0, 0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 230, 3, 44, 22, 0, 230, 231, 1, 0, 0, 0, 231, 232,
		5, 33, 0, 0, 232, 233, 3, 4, 2, 0, 233, 235, 1, 0, 0, 0, 234, 211, 1, 0,
		0, 0, 234, 223, 1, 0, 0, 0, 235, 27, 1, 0, 0, 0, 236, 237, 5, 35, 0, 0,
		237, 240, 5, 46, 0, 0, 238, 239, 5, 33, 0, 0, 239, 241, 3, 44, 22, 0, 240,
		238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 29, 1, 0, 0, 0, 242, 243, 5,
		36, 0, 0, 243, 263, 5, 46, 0, 0, 244, 260, 5, 23, 0, 0, 245, 246, 5, 46,
		0, 0, 246, 247, 5, 25, 0, 0, 247, 254, 3, 46, 23, 0, 248, 249, 5, 37, 0,
		0, 249, 250, 5, 46, 0, 0, 250, 251, 5, 25, 0, 0, 251, 253, 3, 46, 23, 0,
		252, 248, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 254,
		255, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1, 0, 0, 0, 257, 259,
		5, 37, 0, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 261, 1, 0,
		0, 0, 260, 245, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0,
		262, 264, 5, 27, 0, 0, 263, 244, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264,
		31, 1, 0, 0, 0, 265, 266, 3, 4, 2, 0, 266, 267, 5, 33, 0, 0, 267, 268,
		3, 4, 2, 0, 268, 33, 1, 0, 0, 0, 269, 281, 5, 3, 0, 0, 270, 275, 3, 4,
		2, 0, 271, 272, 5, 37, 0, 0, 272, 274, 3, 4, 2, 0, 273, 271, 1, 0, 0, 0,
		274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0, 0, 0, 276,
		279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 280, 5, 37, 0, 0, 279, 278,
		1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 270, 1, 0,
		0, 0, 281, 282, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 284, 5, 4, 0, 0,
		284, 35, 1, 0, 0, 0, 285, 286, 5, 38, 0, 0, 286, 292, 5, 46, 0, 0, 287,
		288, 5, 39, 0, 0, 288, 289, 3, 4, 2, 0, 289, 290, 5, 40, 0, 0, 290, 292,
		1, 0, 0, 0, 291, 285, 1, 0, 0, 0, 291, 287, 1, 0, 0, 0, 292, 37, 1, 0,
		0, 0, 293, 295, 5, 41, 0, 0, 294, 296, 5, 46, 0, 0, 295, 294, 1, 0, 0,
		0, 295, 296, 1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 309, 5, 3, 0, 0, 298,
		303, 3, 50, 25, 0, 299, 300, 5, 37, 0, 0, 300, 302, 3, 50, 25, 0, 301,
		299, 1, 0, 0, 0, 302, 305, 1, 0, 0, 0, 303, 301, 1, 0, 0, 0, 303, 304,
		1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 306, 308, 5, 37,
		0, 0, 307, 306, 1, 0, 0, 0, 307, 308, 1, 0, 0, 0, 308, 310, 1, 0, 0, 0,
		309, 298, 1, 0, 0, 0, 309, 310, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311,
		317, 5, 4, 0, 0, 312, 314, 5, 25, 0, 0, 313, 315, 5, 32, 0, 0, 314, 313,
		1, 0, 0, 0, 314, 315, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 318, 3, 44,
		22, 0, 317, 312, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 321, 1, 0, 0, 0,
		319, 322, 3, 2, 1, 0, 320, 322, 3, 40, 20, 0, 321, 319, 1, 0, 0, 0, 321,
		320, 1, 0, 0, 0, 322, 39, 1, 0, 0, 0, 323, 327, 5, 23, 0, 0, 324, 326,
		3, 2, 1, 0, 325, 324, 1, 0, 0, 0, 326, 329, 1, 0, 0, 0, 327, 325, 1, 0,
		0, 0, 327, 328, 1, 0, 0, 0, 328, 330, 1, 0, 0, 0, 329, 327, 1, 0, 0, 0,
		330, 331, 5, 27, 0, 0, 331, 41, 1, 0, 0, 0, 332, 333, 5, 38, 0, 0, 333,
		339, 5, 46, 0, 0, 334, 335, 5, 39, 0, 0, 335, 336, 3, 44, 22, 0, 336, 337,
		5, 40, 0, 0, 337, 339, 1, 0, 0, 0, 338, 332, 1, 0, 0, 0, 338, 334, 1, 0,
		0, 0, 339, 43, 1, 0, 0, 0, 340, 365, 3, 52, 26, 0, 341, 342, 5, 42, 0,
		0, 342, 354, 5, 3, 0, 0, 343, 348, 3, 46, 23, 0, 344, 345, 5, 37, 0, 0,
		345, 347, 3, 46, 23, 0, 346, 344, 1, 0, 0, 0, 347, 350, 1, 0, 0, 0, 348,
		346, 1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 352, 1, 0, 0, 0, 350, 348,
		1, 0, 0, 0, 351, 353, 5, 37, 0, 0, 352, 351, 1, 0, 0, 0, 352, 353, 1, 0,
		0, 0, 353, 355, 1, 0, 0, 0, 354, 343, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0,
		355, 356, 1, 0, 0, 0, 356, 362, 5, 4, 0, 0, 357, 359, 5, 25, 0, 0, 358,
		360, 5, 32, 0, 0, 359, 358, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 361,
		1, 0, 0, 0, 361, 363, 3, 44, 22, 0, 362, 357, 1, 0, 0, 0, 362, 363, 1,
		0, 0, 0, 363, 365, 1, 0, 0, 0, 364, 340, 1, 0, 0, 0, 364, 341, 1, 0, 0,
		0, 365, 369, 1, 0, 0, 0, 366, 368, 3, 42, 21, 0, 367, 366, 1, 0, 0, 0,
		368, 371, 1, 0, 0, 0, 369, 367, 1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370,
		45, 1, 0, 0, 0, 371, 369, 1, 0, 0, 0, 372, 374, 3, 44, 22, 0, 373, 375,
		5, 32, 0, 0, 374, 373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 47, 1, 0,
		0, 0, 376, 377, 5, 46, 0, 0, 377, 49, 1, 0, 0, 0, 378, 384, 3, 48, 24,
		0, 379, 381, 5, 25, 0, 0, 380, 382, 5, 32, 0, 0, 381, 380, 1, 0, 0, 0,
		381, 382, 1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 385, 3, 46, 23, 0, 384,
		379, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 385, 388, 1, 0, 0, 0, 386, 387,
		5, 33, 0, 0, 387, 389, 3, 4, 2, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0,
		0, 0, 389, 51, 1, 0, 0, 0, 390, 391, 7, 5, 0, 0, 391, 53, 1, 0, 0, 0, 52,
		57, 65, 70, 85, 101, 103, 109, 111, 120, 132, 141, 146, 155, 160, 173,
		175, 191, 193, 209, 215, 221, 227, 234, 240, 254, 258, 260, 263, 275, 279,
		281, 291, 295, 303, 307, 309, 314, 317, 321, 327, 338, 348, 352, 354, 359,
		362, 364, 369, 374, 381, 384, 388,
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

// RW3CParserInit initializes any static state used to implement RW3CParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewRW3CParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func RW3CParserInit() {
	staticData := &rw3cParserStaticData
	staticData.once.Do(rw3cParserInit)
}

// NewRW3CParser produces a new parser instance for the optional input antlr.TokenStream.
func NewRW3CParser(input antlr.TokenStream) *RW3CParser {
	RW3CParserInit()
	this := new(RW3CParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &rw3cParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// RW3CParser tokens.
const (
	RW3CParserEOF            = antlr.TokenEOF
	RW3CParserT__0           = 1
	RW3CParserT__1           = 2
	RW3CParserT__2           = 3
	RW3CParserT__3           = 4
	RW3CParserT__4           = 5
	RW3CParserT__5           = 6
	RW3CParserT__6           = 7
	RW3CParserT__7           = 8
	RW3CParserT__8           = 9
	RW3CParserT__9           = 10
	RW3CParserT__10          = 11
	RW3CParserT__11          = 12
	RW3CParserT__12          = 13
	RW3CParserT__13          = 14
	RW3CParserT__14          = 15
	RW3CParserT__15          = 16
	RW3CParserT__16          = 17
	RW3CParserT__17          = 18
	RW3CParserT__18          = 19
	RW3CParserT__19          = 20
	RW3CParserT__20          = 21
	RW3CParserT__21          = 22
	RW3CParserT__22          = 23
	RW3CParserT__23          = 24
	RW3CParserT__24          = 25
	RW3CParserT__25          = 26
	RW3CParserT__26          = 27
	RW3CParserT__27          = 28
	RW3CParserT__28          = 29
	RW3CParserT__29          = 30
	RW3CParserT__30          = 31
	RW3CParserT__31          = 32
	RW3CParserT__32          = 33
	RW3CParserT__33          = 34
	RW3CParserT__34          = 35
	RW3CParserT__35          = 36
	RW3CParserT__36          = 37
	RW3CParserT__37          = 38
	RW3CParserT__38          = 39
	RW3CParserT__39          = 40
	RW3CParserT__40          = 41
	RW3CParserT__41          = 42
	RW3CParserRESERVED       = 43
	RW3CParserIDENTIF_PART   = 44
	RW3CParserIDENTIF_START  = 45
	RW3CParserIDENTIF        = 46
	RW3CParserNEWLINE        = 47
	RW3CParserLONG           = 48
	RW3CParserDOUBLE         = 49
	RW3CParserSTR            = 50
	RW3CParserESC            = 51
	RW3CParserBOOL           = 52
	RW3CParserMULTI_COMMENT  = 53
	RW3CParserSINGLE_COMMENT = 54
)

// RW3CParser rules.
const (
	RW3CParserRULE_prog                  = 0
	RW3CParserRULE_run                   = 1
	RW3CParserRULE_expr                  = 2
	RW3CParserRULE_chain_expr            = 3
	RW3CParserRULE_atomic_expr           = 4
	RW3CParserRULE_stmt                  = 5
	RW3CParserRULE_if_stmt               = 6
	RW3CParserRULE_if_expr               = 7
	RW3CParserRULE_switch_stmt           = 8
	RW3CParserRULE_switch_expr           = 9
	RW3CParserRULE_while_stmt            = 10
	RW3CParserRULE_while_expr            = 11
	RW3CParserRULE_ret_stmt              = 12
	RW3CParserRULE_def_var_stmt          = 13
	RW3CParserRULE_def_type_stmt         = 14
	RW3CParserRULE_struct_stmt           = 15
	RW3CParserRULE_assign_stmt           = 16
	RW3CParserRULE_call_expr             = 17
	RW3CParserRULE_access_prop_expr      = 18
	RW3CParserRULE_fn_expr               = 19
	RW3CParserRULE_scope                 = 20
	RW3CParserRULE_access_prop_type_expr = 21
	RW3CParserRULE_type_expr             = 22
	RW3CParserRULE_type_arg              = 23
	RW3CParserRULE_name_expr             = 24
	RW3CParserRULE_arg                   = 25
	RW3CParserRULE_lit_expr              = 26
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_run returns the _run rule contexts.
	Get_run() IRunContext

	// Set_run sets the _run rule contexts.
	Set_run(IRunContext)

	// GetExpr_list returns the expr_list rule context list.
	GetExpr_list() []IRunContext

	// SetExpr_list sets the expr_list rule context list.
	SetExpr_list([]IRunContext)

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_run      IRunContext
	expr_list []IRunContext
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Get_run() IRunContext { return s._run }

func (s *ProgContext) Set_run(v IRunContext) { s._run = v }

func (s *ProgContext) GetExpr_list() []IRunContext { return s.expr_list }

func (s *ProgContext) SetExpr_list(v []IRunContext) { s.expr_list = v }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(RW3CParserEOF, 0)
}

func (s *ProgContext) AllRun() []IRunContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunContext); ok {
			len++
		}
	}

	tst := make([]IRunContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunContext); ok {
			tst[i] = t.(IRunContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Run(i int) IRunContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
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

	return t.(IRunContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *RW3CParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, RW3CParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33568213222424970) != 0 {
		{
			p.SetState(54)

			var _x = p.Run()

			localctx.(*ProgContext)._run = _x
		}
		localctx.(*ProgContext).expr_list = append(localctx.(*ProgContext).expr_list, localctx.(*ProgContext)._run)

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
		p.Match(RW3CParserEOF)
	}

	return localctx
}

// IRunContext is an interface to support dynamic dispatch.
type IRunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRunContext differentiates from other interfaces.
	IsRunContext()
}

type RunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunContext() *RunContext {
	var p = new(RunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_run
	return p
}

func (*RunContext) IsRunContext() {}

func NewRunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunContext {
	var p = new(RunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_run

	return p
}

func (s *RunContext) GetParser() antlr.Parser { return s.parser }

func (s *RunContext) Stmt() IStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStmtContext)
}

func (s *RunContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterRun(s)
	}
}

func (s *RunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitRun(s)
	}
}

func (p *RW3CParser) Run() (localctx IRunContext) {
	this := p
	_ = this

	localctx = NewRunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, RW3CParserRULE_run)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(62)
			p.Stmt()
		}

	case 2:
		{
			p.SetState(63)
			p.expr(0)
		}

	case 3:
		{
			p.SetState(64)
			p.Match(RW3CParserT__0)
		}

	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(67)
				p.Match(RW3CParserT__1)
			}

		}
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) Chain_expr() IChain_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChain_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChain_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *RW3CParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *RW3CParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, RW3CParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__2:
		{
			p.SetState(74)
			p.Match(RW3CParserT__2)
		}
		{
			p.SetState(75)
			p.expr(0)
		}
		{
			p.SetState(76)
			p.Match(RW3CParserT__3)
		}

	case RW3CParserT__16, RW3CParserT__17:
		{
			p.SetState(78)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RW3CParserT__16 || _la == RW3CParserT__17) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(79)
			p.expr(5)
		}

	case RW3CParserT__6, RW3CParserT__7:
		{
			p.SetState(80)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RW3CParserT__6 || _la == RW3CParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(81)
			p.expr(4)
		}

	case RW3CParserT__18:
		{
			p.SetState(82)
			p.Match(RW3CParserT__18)
		}
		{
			p.SetState(83)
			p.expr(3)
		}

	case RW3CParserT__19, RW3CParserT__21, RW3CParserT__22, RW3CParserT__27, RW3CParserT__40, RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL:
		{
			p.SetState(84)
			p.Chain_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(88)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__4 || _la == RW3CParserT__5) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(89)
					p.expr(10)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(91)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__6 || _la == RW3CParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(92)
					p.expr(9)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(94)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)
					p.expr(8)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(97)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__14 || _la == RW3CParserT__15) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)
					p.expr(7)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(100)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__16 || _la == RW3CParserT__17) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}

		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IChain_exprContext is an interface to support dynamic dispatch.
type IChain_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChain_exprContext differentiates from other interfaces.
	IsChain_exprContext()
}

type Chain_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChain_exprContext() *Chain_exprContext {
	var p = new(Chain_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_chain_expr
	return p
}

func (*Chain_exprContext) IsChain_exprContext() {}

func NewChain_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chain_exprContext {
	var p = new(Chain_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_chain_expr

	return p
}

func (s *Chain_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Chain_exprContext) Atomic_expr() IAtomic_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomic_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomic_exprContext)
}

func (s *Chain_exprContext) AllCall_expr() []ICall_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICall_exprContext); ok {
			len++
		}
	}

	tst := make([]ICall_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICall_exprContext); ok {
			tst[i] = t.(ICall_exprContext)
			i++
		}
	}

	return tst
}

func (s *Chain_exprContext) Call_expr(i int) ICall_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICall_exprContext); ok {
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

	return t.(ICall_exprContext)
}

func (s *Chain_exprContext) AllAccess_prop_expr() []IAccess_prop_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccess_prop_exprContext); ok {
			len++
		}
	}

	tst := make([]IAccess_prop_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccess_prop_exprContext); ok {
			tst[i] = t.(IAccess_prop_exprContext)
			i++
		}
	}

	return tst
}

func (s *Chain_exprContext) Access_prop_expr(i int) IAccess_prop_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccess_prop_exprContext); ok {
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

	return t.(IAccess_prop_exprContext)
}

func (s *Chain_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chain_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chain_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterChain_expr(s)
	}
}

func (s *Chain_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitChain_expr(s)
	}
}

func (p *RW3CParser) Chain_expr() (localctx IChain_exprContext) {
	this := p
	_ = this

	localctx = NewChain_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RW3CParserRULE_chain_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Atomic_expr()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(109)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RW3CParserT__2:
				{
					p.SetState(107)
					p.Call_expr()
				}

			case RW3CParserT__37, RW3CParserT__38:
				{
					p.SetState(108)
					p.Access_prop_expr()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomic_exprContext is an interface to support dynamic dispatch.
type IAtomic_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomic_exprContext differentiates from other interfaces.
	IsAtomic_exprContext()
}

type Atomic_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomic_exprContext() *Atomic_exprContext {
	var p = new(Atomic_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_atomic_expr
	return p
}

func (*Atomic_exprContext) IsAtomic_exprContext() {}

func NewAtomic_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atomic_exprContext {
	var p = new(Atomic_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_atomic_expr

	return p
}

func (s *Atomic_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Atomic_exprContext) Fn_expr() IFn_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFn_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFn_exprContext)
}

func (s *Atomic_exprContext) Lit_expr() ILit_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILit_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILit_exprContext)
}

func (s *Atomic_exprContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *Atomic_exprContext) Switch_expr() ISwitch_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_exprContext)
}

func (s *Atomic_exprContext) If_expr() IIf_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_exprContext)
}

func (s *Atomic_exprContext) While_expr() IWhile_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_exprContext)
}

func (s *Atomic_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atomic_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atomic_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAtomic_expr(s)
	}
}

func (s *Atomic_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAtomic_expr(s)
	}
}

func (p *RW3CParser) Atomic_expr() (localctx IAtomic_exprContext) {
	this := p
	_ = this

	localctx = NewAtomic_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, RW3CParserRULE_atomic_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__40:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Fn_expr()
		}

	case RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Lit_expr()
		}

	case RW3CParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(116)
			p.Scope()
		}

	case RW3CParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(117)
			p.Switch_expr()
		}

	case RW3CParserT__19:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(118)
			p.If_expr()
		}

	case RW3CParserT__27:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			p.While_expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_stmt
	return p
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) MULTI_COMMENT() antlr.TerminalNode {
	return s.GetToken(RW3CParserMULTI_COMMENT, 0)
}

func (s *StmtContext) SINGLE_COMMENT() antlr.TerminalNode {
	return s.GetToken(RW3CParserSINGLE_COMMENT, 0)
}

func (s *StmtContext) Assign_stmt() IAssign_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssign_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssign_stmtContext)
}

func (s *StmtContext) Def_var_stmt() IDef_var_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_var_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_var_stmtContext)
}

func (s *StmtContext) Def_type_stmt() IDef_type_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_type_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_type_stmtContext)
}

func (s *StmtContext) Struct_stmt() IStruct_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStruct_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStruct_stmtContext)
}

func (s *StmtContext) Ret_stmt() IRet_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRet_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRet_stmtContext)
}

func (s *StmtContext) If_stmt() IIf_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_stmtContext)
}

func (s *StmtContext) While_stmt() IWhile_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StmtContext) Switch_stmt() ISwitch_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (p *RW3CParser) Stmt() (localctx IStmtContext) {
	this := p
	_ = this

	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RW3CParserRULE_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(RW3CParserMULTI_COMMENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(RW3CParserSINGLE_COMMENT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Assign_stmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.Def_var_stmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.Def_type_stmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
			p.Struct_stmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(128)
			p.Ret_stmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(129)
			p.If_stmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(130)
			p.While_stmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(131)
			p.Switch_stmt()
		}

	}

	return localctx
}

// IIf_stmtContext is an interface to support dynamic dispatch.
type IIf_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_stmtContext differentiates from other interfaces.
	IsIf_stmtContext()
}

type If_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_stmtContext() *If_stmtContext {
	var p = new(If_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_if_stmt
	return p
}

func (*If_stmtContext) IsIf_stmtContext() {}

func NewIf_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_stmtContext {
	var p = new(If_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_if_stmt

	return p
}

func (s *If_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *If_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *If_stmtContext) AllRun() []IRunContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunContext); ok {
			len++
		}
	}

	tst := make([]IRunContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunContext); ok {
			tst[i] = t.(IRunContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) Run(i int) IRunContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
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

	return t.(IRunContext)
}

func (s *If_stmtContext) AllIf_stmt() []IIf_stmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_stmtContext); ok {
			len++
		}
	}

	tst := make([]IIf_stmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_stmtContext); ok {
			tst[i] = t.(IIf_stmtContext)
			i++
		}
	}

	return tst
}

func (s *If_stmtContext) If_stmt(i int) IIf_stmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_stmtContext); ok {
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

	return t.(IIf_stmtContext)
}

func (s *If_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterIf_stmt(s)
	}
}

func (s *If_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitIf_stmt(s)
	}
}

func (p *RW3CParser) If_stmt() (localctx IIf_stmtContext) {
	this := p
	_ = this

	localctx = NewIf_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RW3CParserRULE_if_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(RW3CParserT__19)
	}
	{
		p.SetState(135)
		p.expr(0)
	}
	{
		p.SetState(136)
		p.Run()
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(137)
				p.Match(RW3CParserT__20)
			}
			{
				p.SetState(138)
				p.If_stmt()
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(144)
			p.Match(RW3CParserT__20)
		}
		{
			p.SetState(145)
			p.Run()
		}

	}

	return localctx
}

// IIf_exprContext is an interface to support dynamic dispatch.
type IIf_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_exprContext differentiates from other interfaces.
	IsIf_exprContext()
}

type If_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_exprContext() *If_exprContext {
	var p = new(If_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_if_expr
	return p
}

func (*If_exprContext) IsIf_exprContext() {}

func NewIf_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_exprContext {
	var p = new(If_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_if_expr

	return p
}

func (s *If_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *If_exprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *If_exprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *If_exprContext) AllIf_expr() []IIf_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIf_exprContext); ok {
			len++
		}
	}

	tst := make([]IIf_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIf_exprContext); ok {
			tst[i] = t.(IIf_exprContext)
			i++
		}
	}

	return tst
}

func (s *If_exprContext) If_expr(i int) IIf_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_exprContext); ok {
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

	return t.(IIf_exprContext)
}

func (s *If_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterIf_expr(s)
	}
}

func (s *If_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitIf_expr(s)
	}
}

func (p *RW3CParser) If_expr() (localctx IIf_exprContext) {
	this := p
	_ = this

	localctx = NewIf_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RW3CParserRULE_if_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(RW3CParserT__19)
	}
	{
		p.SetState(149)
		p.expr(0)
	}
	{
		p.SetState(150)
		p.expr(0)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(151)
				p.Match(RW3CParserT__20)
			}
			{
				p.SetState(152)
				p.If_expr()
			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(158)
			p.Match(RW3CParserT__20)
		}
		{
			p.SetState(159)
			p.expr(0)
		}

	}

	return localctx
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_switch_stmt
	return p
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Switch_stmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Switch_stmtContext) AllRun() []IRunContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunContext); ok {
			len++
		}
	}

	tst := make([]IRunContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunContext); ok {
			tst[i] = t.(IRunContext)
			i++
		}
	}

	return tst
}

func (s *Switch_stmtContext) Run(i int) IRunContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
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

	return t.(IRunContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (p *RW3CParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	this := p
	_ = this

	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RW3CParserRULE_switch_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(RW3CParserT__21)
	}
	{
		p.SetState(163)
		p.expr(0)
	}
	{
		p.SetState(164)
		p.Match(RW3CParserT__22)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RW3CParserT__23 || _la == RW3CParserT__25 {
		p.SetState(173)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RW3CParserT__23:
			{
				p.SetState(165)
				p.Match(RW3CParserT__23)
			}
			{
				p.SetState(166)
				p.expr(0)
			}
			{
				p.SetState(167)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(168)
				p.Run()
			}

		case RW3CParserT__25:
			{
				p.SetState(170)
				p.Match(RW3CParserT__25)
			}
			{
				p.SetState(171)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(172)
				p.Run()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(RW3CParserT__26)
	}

	return localctx
}

// ISwitch_exprContext is an interface to support dynamic dispatch.
type ISwitch_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_exprContext differentiates from other interfaces.
	IsSwitch_exprContext()
}

type Switch_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_exprContext() *Switch_exprContext {
	var p = new(Switch_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_switch_expr
	return p
}

func (*Switch_exprContext) IsSwitch_exprContext() {}

func NewSwitch_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_exprContext {
	var p = new(Switch_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_switch_expr

	return p
}

func (s *Switch_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_exprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Switch_exprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Switch_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterSwitch_expr(s)
	}
}

func (s *Switch_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitSwitch_expr(s)
	}
}

func (p *RW3CParser) Switch_expr() (localctx ISwitch_exprContext) {
	this := p
	_ = this

	localctx = NewSwitch_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, RW3CParserRULE_switch_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(RW3CParserT__21)
	}
	{
		p.SetState(181)
		p.expr(0)
	}
	{
		p.SetState(182)
		p.Match(RW3CParserT__22)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RW3CParserT__23 || _la == RW3CParserT__25 {
		p.SetState(191)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RW3CParserT__23:
			{
				p.SetState(183)
				p.Match(RW3CParserT__23)
			}
			{
				p.SetState(184)
				p.expr(0)
			}
			{
				p.SetState(185)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(186)
				p.expr(0)
			}

		case RW3CParserT__25:
			{
				p.SetState(188)
				p.Match(RW3CParserT__25)
			}
			{
				p.SetState(189)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(190)
				p.expr(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(196)
		p.Match(RW3CParserT__26)
	}

	return localctx
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_while_stmt
	return p
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *While_stmtContext) Run() IRunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *RW3CParser) While_stmt() (localctx IWhile_stmtContext) {
	this := p
	_ = this

	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RW3CParserRULE_while_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(RW3CParserT__27)
	}
	{
		p.SetState(199)
		p.expr(0)
	}
	{
		p.SetState(200)
		p.Run()
	}

	return localctx
}

// IWhile_exprContext is an interface to support dynamic dispatch.
type IWhile_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_exprContext differentiates from other interfaces.
	IsWhile_exprContext()
}

type While_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_exprContext() *While_exprContext {
	var p = new(While_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_while_expr
	return p
}

func (*While_exprContext) IsWhile_exprContext() {}

func NewWhile_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_exprContext {
	var p = new(While_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_while_expr

	return p
}

func (s *While_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *While_exprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *While_exprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *While_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhile_expr(s)
	}
}

func (s *While_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhile_expr(s)
	}
}

func (p *RW3CParser) While_expr() (localctx IWhile_exprContext) {
	this := p
	_ = this

	localctx = NewWhile_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RW3CParserRULE_while_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(RW3CParserT__27)
	}
	{
		p.SetState(203)
		p.expr(0)
	}
	{
		p.SetState(204)
		p.expr(0)
	}

	return localctx
}

// IRet_stmtContext is an interface to support dynamic dispatch.
type IRet_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRet_stmtContext differentiates from other interfaces.
	IsRet_stmtContext()
}

type Ret_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRet_stmtContext() *Ret_stmtContext {
	var p = new(Ret_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_ret_stmt
	return p
}

func (*Ret_stmtContext) IsRet_stmtContext() {}

func NewRet_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ret_stmtContext {
	var p = new(Ret_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_ret_stmt

	return p
}

func (s *Ret_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Ret_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Ret_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ret_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ret_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterRet_stmt(s)
	}
}

func (s *Ret_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitRet_stmt(s)
	}
}

func (p *RW3CParser) Ret_stmt() (localctx IRet_stmtContext) {
	this := p
	_ = this

	localctx = NewRet_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RW3CParserRULE_ret_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(RW3CParserT__28)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(207)
			p.Match(RW3CParserT__29)
		}
		{
			p.SetState(208)
			p.expr(0)
		}

	}

	return localctx
}

// IDef_var_stmtContext is an interface to support dynamic dispatch.
type IDef_var_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDef_var_stmtContext differentiates from other interfaces.
	IsDef_var_stmtContext()
}

type Def_var_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_var_stmtContext() *Def_var_stmtContext {
	var p = new(Def_var_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_def_var_stmt
	return p
}

func (*Def_var_stmtContext) IsDef_var_stmtContext() {}

func NewDef_var_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_var_stmtContext {
	var p = new(Def_var_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_def_var_stmt

	return p
}

func (s *Def_var_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_var_stmtContext) Name_expr() IName_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IName_exprContext)
}

func (s *Def_var_stmtContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Def_var_stmtContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Def_var_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_var_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Def_var_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterDef_var_stmt(s)
	}
}

func (s *Def_var_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitDef_var_stmt(s)
	}
}

func (p *RW3CParser) Def_var_stmt() (localctx IDef_var_stmtContext) {
	this := p
	_ = this

	localctx = NewDef_var_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RW3CParserRULE_def_var_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(211)
			p.Match(RW3CParserT__30)
		}
		{
			p.SetState(212)
			p.Name_expr()
		}

		{
			p.SetState(213)
			p.Match(RW3CParserT__24)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__31 {
			{
				p.SetState(214)
				p.Match(RW3CParserT__31)
			}

		}
		{
			p.SetState(217)
			p.Type_expr()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(219)
				p.Match(RW3CParserT__32)
			}
			{
				p.SetState(220)
				p.expr(0)
			}

		}

	case RW3CParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(RW3CParserT__33)
		}
		{
			p.SetState(224)
			p.Name_expr()
		}

		{
			p.SetState(225)
			p.Match(RW3CParserT__24)
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__31 {
			{
				p.SetState(226)
				p.Match(RW3CParserT__31)
			}

		}
		{
			p.SetState(229)
			p.Type_expr()
		}

		{
			p.SetState(231)
			p.Match(RW3CParserT__32)
		}
		{
			p.SetState(232)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDef_type_stmtContext is an interface to support dynamic dispatch.
type IDef_type_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDef_type_stmtContext differentiates from other interfaces.
	IsDef_type_stmtContext()
}

type Def_type_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_type_stmtContext() *Def_type_stmtContext {
	var p = new(Def_type_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_def_type_stmt
	return p
}

func (*Def_type_stmtContext) IsDef_type_stmtContext() {}

func NewDef_type_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_type_stmtContext {
	var p = new(Def_type_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_def_type_stmt

	return p
}

func (s *Def_type_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_type_stmtContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Def_type_stmtContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Def_type_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_type_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Def_type_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterDef_type_stmt(s)
	}
}

func (s *Def_type_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitDef_type_stmt(s)
	}
}

func (p *RW3CParser) Def_type_stmt() (localctx IDef_type_stmtContext) {
	this := p
	_ = this

	localctx = NewDef_type_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RW3CParserRULE_def_type_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(RW3CParserT__34)
	}
	{
		p.SetState(237)
		p.Match(RW3CParserIDENTIF)
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(238)
			p.Match(RW3CParserT__32)
		}
		{
			p.SetState(239)
			p.Type_expr()
		}

	}

	return localctx
}

// IStruct_stmtContext is an interface to support dynamic dispatch.
type IStruct_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStruct_stmtContext differentiates from other interfaces.
	IsStruct_stmtContext()
}

type Struct_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_stmtContext() *Struct_stmtContext {
	var p = new(Struct_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_struct_stmt
	return p
}

func (*Struct_stmtContext) IsStruct_stmtContext() {}

func NewStruct_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_stmtContext {
	var p = new(Struct_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_struct_stmt

	return p
}

func (s *Struct_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Struct_stmtContext) AllIDENTIF() []antlr.TerminalNode {
	return s.GetTokens(RW3CParserIDENTIF)
}

func (s *Struct_stmtContext) IDENTIF(i int) antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, i)
}

func (s *Struct_stmtContext) AllType_arg() []IType_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_argContext); ok {
			len++
		}
	}

	tst := make([]IType_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_argContext); ok {
			tst[i] = t.(IType_argContext)
			i++
		}
	}

	return tst
}

func (s *Struct_stmtContext) Type_arg(i int) IType_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_argContext); ok {
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

	return t.(IType_argContext)
}

func (s *Struct_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterStruct_stmt(s)
	}
}

func (s *Struct_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitStruct_stmt(s)
	}
}

func (p *RW3CParser) Struct_stmt() (localctx IStruct_stmtContext) {
	this := p
	_ = this

	localctx = NewStruct_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RW3CParserRULE_struct_stmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(RW3CParserT__35)
	}
	{
		p.SetState(243)
		p.Match(RW3CParserIDENTIF)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(244)
			p.Match(RW3CParserT__22)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserIDENTIF {
			{
				p.SetState(245)
				p.Match(RW3CParserIDENTIF)
			}
			{
				p.SetState(246)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(247)
				p.Type_arg()
			}
			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(248)
						p.Match(RW3CParserT__36)
					}
					{
						p.SetState(249)
						p.Match(RW3CParserIDENTIF)
					}
					{
						p.SetState(250)
						p.Match(RW3CParserT__24)
					}
					{
						p.SetState(251)
						p.Type_arg()
					}

				}
				p.SetState(256)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
			}
			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__36 {
				{
					p.SetState(257)
					p.Match(RW3CParserT__36)
				}

			}

		}
		{
			p.SetState(262)
			p.Match(RW3CParserT__26)
		}

	}

	return localctx
}

// IAssign_stmtContext is an interface to support dynamic dispatch.
type IAssign_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssign_stmtContext differentiates from other interfaces.
	IsAssign_stmtContext()
}

type Assign_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssign_stmtContext() *Assign_stmtContext {
	var p = new(Assign_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_assign_stmt
	return p
}

func (*Assign_stmtContext) IsAssign_stmtContext() {}

func NewAssign_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Assign_stmtContext {
	var p = new(Assign_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_assign_stmt

	return p
}

func (s *Assign_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Assign_stmtContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Assign_stmtContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Assign_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Assign_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Assign_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAssign_stmt(s)
	}
}

func (s *Assign_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAssign_stmt(s)
	}
}

func (p *RW3CParser) Assign_stmt() (localctx IAssign_stmtContext) {
	this := p
	_ = this

	localctx = NewAssign_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RW3CParserRULE_assign_stmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(265)
		p.expr(0)
	}
	{
		p.SetState(266)
		p.Match(RW3CParserT__32)
	}
	{
		p.SetState(267)
		p.expr(0)
	}

	return localctx
}

// ICall_exprContext is an interface to support dynamic dispatch.
type ICall_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCall_exprContext differentiates from other interfaces.
	IsCall_exprContext()
}

type Call_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCall_exprContext() *Call_exprContext {
	var p = new(Call_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_call_expr
	return p
}

func (*Call_exprContext) IsCall_exprContext() {}

func NewCall_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Call_exprContext {
	var p = new(Call_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_call_expr

	return p
}

func (s *Call_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Call_exprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Call_exprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *Call_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Call_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Call_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterCall_expr(s)
	}
}

func (s *Call_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitCall_expr(s)
	}
}

func (p *RW3CParser) Call_expr() (localctx ICall_exprContext) {
	this := p
	_ = this

	localctx = NewCall_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RW3CParserRULE_call_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(RW3CParserT__2)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6546492514763144) != 0 {
		{
			p.SetState(270)
			p.expr(0)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(271)
					p.Match(RW3CParserT__36)
				}
				{
					p.SetState(272)
					p.expr(0)
				}

			}
			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
		}
		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__36 {
			{
				p.SetState(278)
				p.Match(RW3CParserT__36)
			}

		}

	}
	{
		p.SetState(283)
		p.Match(RW3CParserT__3)
	}

	return localctx
}

// IAccess_prop_exprContext is an interface to support dynamic dispatch.
type IAccess_prop_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccess_prop_exprContext differentiates from other interfaces.
	IsAccess_prop_exprContext()
}

type Access_prop_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccess_prop_exprContext() *Access_prop_exprContext {
	var p = new(Access_prop_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_access_prop_expr
	return p
}

func (*Access_prop_exprContext) IsAccess_prop_exprContext() {}

func NewAccess_prop_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Access_prop_exprContext {
	var p = new(Access_prop_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_access_prop_expr

	return p
}

func (s *Access_prop_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Access_prop_exprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Access_prop_exprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Access_prop_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Access_prop_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Access_prop_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAccess_prop_expr(s)
	}
}

func (s *Access_prop_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAccess_prop_expr(s)
	}
}

func (p *RW3CParser) Access_prop_expr() (localctx IAccess_prop_exprContext) {
	this := p
	_ = this

	localctx = NewAccess_prop_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RW3CParserRULE_access_prop_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(291)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__37:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(RW3CParserT__37)
		}
		{
			p.SetState(286)
			p.Match(RW3CParserIDENTIF)
		}

	case RW3CParserT__38:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(RW3CParserT__38)
		}
		{
			p.SetState(288)
			p.expr(0)
		}
		{
			p.SetState(289)
			p.Match(RW3CParserT__39)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFn_exprContext is an interface to support dynamic dispatch.
type IFn_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFn_exprContext differentiates from other interfaces.
	IsFn_exprContext()
}

type Fn_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFn_exprContext() *Fn_exprContext {
	var p = new(Fn_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_fn_expr
	return p
}

func (*Fn_exprContext) IsFn_exprContext() {}

func NewFn_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fn_exprContext {
	var p = new(Fn_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_fn_expr

	return p
}

func (s *Fn_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Fn_exprContext) Run() IRunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunContext)
}

func (s *Fn_exprContext) Scope() IScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *Fn_exprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Fn_exprContext) AllArg() []IArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArgContext); ok {
			len++
		}
	}

	tst := make([]IArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArgContext); ok {
			tst[i] = t.(IArgContext)
			i++
		}
	}

	return tst
}

func (s *Fn_exprContext) Arg(i int) IArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgContext); ok {
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

	return t.(IArgContext)
}

func (s *Fn_exprContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Fn_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fn_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fn_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterFn_expr(s)
	}
}

func (s *Fn_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitFn_expr(s)
	}
}

func (p *RW3CParser) Fn_expr() (localctx IFn_exprContext) {
	this := p
	_ = this

	localctx = NewFn_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RW3CParserRULE_fn_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(RW3CParserT__40)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserIDENTIF {
		{
			p.SetState(294)
			p.Match(RW3CParserIDENTIF)
		}

	}
	{
		p.SetState(297)
		p.Match(RW3CParserT__2)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserIDENTIF {
		{
			p.SetState(298)
			p.Arg()
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(299)
					p.Match(RW3CParserT__36)
				}
				{
					p.SetState(300)
					p.Arg()
				}

			}
			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__36 {
			{
				p.SetState(306)
				p.Match(RW3CParserT__36)
			}

		}

	}
	{
		p.SetState(311)
		p.Match(RW3CParserT__3)
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__24 {
		{
			p.SetState(312)
			p.Match(RW3CParserT__24)
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__31 {
			{
				p.SetState(313)
				p.Match(RW3CParserT__31)
			}

		}
		{
			p.SetState(316)
			p.Type_expr()
		}

	}
	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(319)
			p.Run()
		}

	case 2:
		{
			p.SetState(320)
			p.Scope()
		}

	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) AllRun() []IRunContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunContext); ok {
			len++
		}
	}

	tst := make([]IRunContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunContext); ok {
			tst[i] = t.(IRunContext)
			i++
		}
	}

	return tst
}

func (s *ScopeContext) Run(i int) IRunContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunContext); ok {
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

	return t.(IRunContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitScope(s)
	}
}

func (p *RW3CParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RW3CParserRULE_scope)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(RW3CParserT__22)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33568213222424970) != 0 {
		{
			p.SetState(324)
			p.Run()
		}

		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(330)
		p.Match(RW3CParserT__26)
	}

	return localctx
}

// IAccess_prop_type_exprContext is an interface to support dynamic dispatch.
type IAccess_prop_type_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccess_prop_type_exprContext differentiates from other interfaces.
	IsAccess_prop_type_exprContext()
}

type Access_prop_type_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccess_prop_type_exprContext() *Access_prop_type_exprContext {
	var p = new(Access_prop_type_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_access_prop_type_expr
	return p
}

func (*Access_prop_type_exprContext) IsAccess_prop_type_exprContext() {}

func NewAccess_prop_type_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Access_prop_type_exprContext {
	var p = new(Access_prop_type_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_access_prop_type_expr

	return p
}

func (s *Access_prop_type_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Access_prop_type_exprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Access_prop_type_exprContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Access_prop_type_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Access_prop_type_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Access_prop_type_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAccess_prop_type_expr(s)
	}
}

func (s *Access_prop_type_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAccess_prop_type_expr(s)
	}
}

func (p *RW3CParser) Access_prop_type_expr() (localctx IAccess_prop_type_exprContext) {
	this := p
	_ = this

	localctx = NewAccess_prop_type_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RW3CParserRULE_access_prop_type_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(338)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__37:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(332)
			p.Match(RW3CParserT__37)
		}
		{
			p.SetState(333)
			p.Match(RW3CParserIDENTIF)
		}

	case RW3CParserT__38:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(RW3CParserT__38)
		}
		{
			p.SetState(335)
			p.Type_expr()
		}
		{
			p.SetState(336)
			p.Match(RW3CParserT__39)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_exprContext is an interface to support dynamic dispatch.
type IType_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_exprContext differentiates from other interfaces.
	IsType_exprContext()
}

type Type_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_exprContext() *Type_exprContext {
	var p = new(Type_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_type_expr
	return p
}

func (*Type_exprContext) IsType_exprContext() {}

func NewType_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_exprContext {
	var p = new(Type_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_type_expr

	return p
}

func (s *Type_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_exprContext) Lit_expr() ILit_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILit_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILit_exprContext)
}

func (s *Type_exprContext) AllAccess_prop_type_expr() []IAccess_prop_type_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccess_prop_type_exprContext); ok {
			len++
		}
	}

	tst := make([]IAccess_prop_type_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccess_prop_type_exprContext); ok {
			tst[i] = t.(IAccess_prop_type_exprContext)
			i++
		}
	}

	return tst
}

func (s *Type_exprContext) Access_prop_type_expr(i int) IAccess_prop_type_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccess_prop_type_exprContext); ok {
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

	return t.(IAccess_prop_type_exprContext)
}

func (s *Type_exprContext) AllType_arg() []IType_argContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IType_argContext); ok {
			len++
		}
	}

	tst := make([]IType_argContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IType_argContext); ok {
			tst[i] = t.(IType_argContext)
			i++
		}
	}

	return tst
}

func (s *Type_exprContext) Type_arg(i int) IType_argContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_argContext); ok {
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

	return t.(IType_argContext)
}

func (s *Type_exprContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Type_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterType_expr(s)
	}
}

func (s *Type_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitType_expr(s)
	}
}

func (p *RW3CParser) Type_expr() (localctx IType_exprContext) {
	this := p
	_ = this

	localctx = NewType_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, RW3CParserRULE_type_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(364)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL:
		{
			p.SetState(340)
			p.Lit_expr()
		}

	case RW3CParserT__41:
		{
			p.SetState(341)
			p.Match(RW3CParserT__41)
		}
		{
			p.SetState(342)
			p.Match(RW3CParserT__2)
		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6548691255033856) != 0 {
			{
				p.SetState(343)
				p.Type_arg()
			}
			p.SetState(348)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(344)
						p.Match(RW3CParserT__36)
					}
					{
						p.SetState(345)
						p.Type_arg()
					}

				}
				p.SetState(350)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
			}
			p.SetState(352)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__36 {
				{
					p.SetState(351)
					p.Match(RW3CParserT__36)
				}

			}

		}
		{
			p.SetState(356)
			p.Match(RW3CParserT__3)
		}
		p.SetState(362)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(357)
				p.Match(RW3CParserT__24)
			}
			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__31 {
				{
					p.SetState(358)
					p.Match(RW3CParserT__31)
				}

			}
			{
				p.SetState(361)
				p.Type_expr()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(366)
				p.Access_prop_type_expr()
			}

		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}

	return localctx
}

// IType_argContext is an interface to support dynamic dispatch.
type IType_argContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_argContext differentiates from other interfaces.
	IsType_argContext()
}

type Type_argContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_argContext() *Type_argContext {
	var p = new(Type_argContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_type_arg
	return p
}

func (*Type_argContext) IsType_argContext() {}

func NewType_argContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_argContext {
	var p = new(Type_argContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_type_arg

	return p
}

func (s *Type_argContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_argContext) Type_expr() IType_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Type_argContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_argContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_argContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterType_arg(s)
	}
}

func (s *Type_argContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitType_arg(s)
	}
}

func (p *RW3CParser) Type_arg() (localctx IType_argContext) {
	this := p
	_ = this

	localctx = NewType_argContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RW3CParserRULE_type_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		p.Type_expr()
	}
	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__31 {
		{
			p.SetState(373)
			p.Match(RW3CParserT__31)
		}

	}

	return localctx
}

// IName_exprContext is an interface to support dynamic dispatch.
type IName_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsName_exprContext differentiates from other interfaces.
	IsName_exprContext()
}

type Name_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyName_exprContext() *Name_exprContext {
	var p = new(Name_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_name_expr
	return p
}

func (*Name_exprContext) IsName_exprContext() {}

func NewName_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Name_exprContext {
	var p = new(Name_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_name_expr

	return p
}

func (s *Name_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Name_exprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Name_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Name_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Name_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterName_expr(s)
	}
}

func (s *Name_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitName_expr(s)
	}
}

func (p *RW3CParser) Name_expr() (localctx IName_exprContext) {
	this := p
	_ = this

	localctx = NewName_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RW3CParserRULE_name_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(RW3CParserIDENTIF)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) Name_expr() IName_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IName_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IName_exprContext)
}

func (s *ArgContext) Type_arg() IType_argContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_argContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_argContext)
}

func (s *ArgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *RW3CParser) Arg() (localctx IArgContext) {
	this := p
	_ = this

	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RW3CParserRULE_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Name_expr()
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__24 {
		{
			p.SetState(379)
			p.Match(RW3CParserT__24)
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__31 {
			{
				p.SetState(380)
				p.Match(RW3CParserT__31)
			}

		}
		{
			p.SetState(383)
			p.Type_arg()
		}

	}
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__32 {
		{
			p.SetState(386)
			p.Match(RW3CParserT__32)
		}
		{
			p.SetState(387)
			p.expr(0)
		}

	}

	return localctx
}

// ILit_exprContext is an interface to support dynamic dispatch.
type ILit_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLit_exprContext differentiates from other interfaces.
	IsLit_exprContext()
}

type Lit_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLit_exprContext() *Lit_exprContext {
	var p = new(Lit_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_lit_expr
	return p
}

func (*Lit_exprContext) IsLit_exprContext() {}

func NewLit_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lit_exprContext {
	var p = new(Lit_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_lit_expr

	return p
}

func (s *Lit_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Lit_exprContext) LONG() antlr.TerminalNode {
	return s.GetToken(RW3CParserLONG, 0)
}

func (s *Lit_exprContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(RW3CParserDOUBLE, 0)
}

func (s *Lit_exprContext) STR() antlr.TerminalNode {
	return s.GetToken(RW3CParserSTR, 0)
}

func (s *Lit_exprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *Lit_exprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RW3CParserBOOL, 0)
}

func (s *Lit_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lit_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lit_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterLit_expr(s)
	}
}

func (s *Lit_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitLit_expr(s)
	}
}

func (p *RW3CParser) Lit_expr() (localctx ILit_exprContext) {
	this := p
	_ = this

	localctx = NewLit_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RW3CParserRULE_lit_expr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&6544293208522752) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *RW3CParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *RW3CParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

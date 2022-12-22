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
		"", "';'", "'break'", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'>'",
		"'<'", "'<='", "'>='", "'=='", "'!='", "'&&'", "'||'", "'++'", "'--'",
		"'!'", "','", "'.'", "'['", "']'", "'if '", "'else '", "'switch '",
		"'{'", "'case '", "':'", "'default'", "'}'", "'while '", "'ret'", "' '",
		"'let '", "'?'", "'='", "'mut '", "'type '", "'struct '", "'fn '", "'fn'",
		"", "", "", "", "", "", "", "", "", "", "'null'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "RESERVED", "IDENTIF_PART", "IDENTIF_START",
		"IDENTIF", "NEWLINE", "LONG", "DOUBLE", "STR", "ESC", "BOOL", "NULL",
		"MULTI_COMMENT", "SINGLE_COMMENT",
	}
	staticData.ruleNames = []string{
		"prog", "run", "whileRun", "whileExprBody", "expr", "atomicExpr", "chainExpr",
		"callExpr", "accessPropExpr", "stmt", "ifStmt", "ifExpr", "switchStmt",
		"switchExpr", "whileStmt", "whileExpr", "retStmt", "defVarStmt", "defTypeStmt",
		"structStmt", "assignStmt", "fnExpr", "scope", "accessPropTypeExpr",
		"typeExpr", "typeArg", "nameExpr", "arg", "litExpr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 408, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 5, 0, 60, 8, 0, 10, 0, 12, 0,
		63, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 69, 8, 1, 1, 1, 5, 1, 72, 8, 1,
		10, 1, 12, 1, 75, 9, 1, 1, 2, 1, 2, 5, 2, 79, 8, 2, 10, 2, 12, 2, 82, 9,
		2, 1, 3, 1, 3, 5, 3, 86, 8, 3, 10, 3, 12, 3, 89, 9, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 5, 4, 120, 8, 4, 10, 4, 12, 4, 123, 9, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 131, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 136, 8,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 1,
		7, 3, 7, 148, 8, 7, 3, 7, 150, 8, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 160, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 3, 9, 172, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5,
		10, 179, 8, 10, 10, 10, 12, 10, 182, 9, 10, 1, 10, 1, 10, 3, 10, 186, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 193, 8, 11, 10, 11, 12, 11,
		196, 9, 11, 1, 11, 1, 11, 3, 11, 200, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 213, 8, 12, 10,
		12, 12, 12, 216, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 231, 8, 13, 10, 13, 12,
		13, 234, 9, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 3, 16, 249, 8, 16, 1, 17, 1, 17, 1,
		17, 1, 17, 3, 17, 255, 8, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 261, 8,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 267, 8, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 3, 17, 274, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 280,
		8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1,
		19, 5, 19, 292, 8, 19, 10, 19, 12, 19, 295, 9, 19, 1, 19, 3, 19, 298, 8,
		19, 3, 19, 300, 8, 19, 1, 19, 3, 19, 303, 8, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 3, 21, 311, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21,
		317, 8, 21, 10, 21, 12, 21, 320, 9, 21, 1, 21, 3, 21, 323, 8, 21, 3, 21,
		325, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 330, 8, 21, 1, 21, 3, 21, 333,
		8, 21, 1, 21, 1, 21, 3, 21, 337, 8, 21, 1, 22, 1, 22, 5, 22, 341, 8, 22,
		10, 22, 12, 22, 344, 9, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 3, 23, 354, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 362, 8, 24, 10, 24, 12, 24, 365, 9, 24, 1, 24, 3, 24, 368, 8, 24,
		3, 24, 370, 8, 24, 1, 24, 1, 24, 1, 24, 3, 24, 375, 8, 24, 1, 24, 3, 24,
		378, 8, 24, 3, 24, 380, 8, 24, 1, 24, 5, 24, 383, 8, 24, 10, 24, 12, 24,
		386, 9, 24, 1, 25, 1, 25, 3, 25, 390, 8, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 3, 27, 397, 8, 27, 1, 27, 3, 27, 400, 8, 27, 1, 27, 1, 27, 3, 27,
		404, 8, 27, 1, 28, 1, 28, 1, 28, 0, 1, 8, 29, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 0, 6, 1, 0, 17, 18, 1, 0, 7, 8, 1, 0, 5, 6, 1, 0, 9, 14, 1,
		0, 15, 16, 3, 0, 46, 46, 48, 50, 52, 53, 452, 0, 61, 1, 0, 0, 0, 2, 68,
		1, 0, 0, 0, 4, 80, 1, 0, 0, 0, 6, 87, 1, 0, 0, 0, 8, 103, 1, 0, 0, 0, 10,
		130, 1, 0, 0, 0, 12, 132, 1, 0, 0, 0, 14, 137, 1, 0, 0, 0, 16, 159, 1,
		0, 0, 0, 18, 171, 1, 0, 0, 0, 20, 173, 1, 0, 0, 0, 22, 187, 1, 0, 0, 0,
		24, 201, 1, 0, 0, 0, 26, 219, 1, 0, 0, 0, 28, 237, 1, 0, 0, 0, 30, 241,
		1, 0, 0, 0, 32, 245, 1, 0, 0, 0, 34, 273, 1, 0, 0, 0, 36, 275, 1, 0, 0,
		0, 38, 281, 1, 0, 0, 0, 40, 304, 1, 0, 0, 0, 42, 308, 1, 0, 0, 0, 44, 338,
		1, 0, 0, 0, 46, 353, 1, 0, 0, 0, 48, 379, 1, 0, 0, 0, 50, 387, 1, 0, 0,
		0, 52, 391, 1, 0, 0, 0, 54, 393, 1, 0, 0, 0, 56, 405, 1, 0, 0, 0, 58, 60,
		3, 2, 1, 0, 59, 58, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0,
		61, 62, 1, 0, 0, 0, 62, 64, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 65, 5,
		0, 0, 1, 65, 1, 1, 0, 0, 0, 66, 69, 3, 18, 9, 0, 67, 69, 3, 8, 4, 0, 68,
		66, 1, 0, 0, 0, 68, 67, 1, 0, 0, 0, 69, 73, 1, 0, 0, 0, 70, 72, 5, 1, 0,
		0, 71, 70, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74,
		1, 0, 0, 0, 74, 3, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 79, 3, 2, 1, 0,
		77, 79, 5, 2, 0, 0, 78, 76, 1, 0, 0, 0, 78, 77, 1, 0, 0, 0, 79, 82, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82,
		80, 1, 0, 0, 0, 83, 86, 3, 8, 4, 0, 84, 86, 5, 2, 0, 0, 85, 83, 1, 0, 0,
		0, 85, 84, 1, 0, 0, 0, 86, 89, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 87, 88,
		1, 0, 0, 0, 88, 7, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 90, 91, 6, 4, -1, 0,
		91, 92, 5, 3, 0, 0, 92, 93, 3, 8, 4, 0, 93, 94, 5, 4, 0, 0, 94, 104, 1,
		0, 0, 0, 95, 96, 7, 0, 0, 0, 96, 104, 3, 8, 4, 6, 97, 98, 7, 1, 0, 0, 98,
		104, 3, 8, 4, 5, 99, 100, 5, 19, 0, 0, 100, 104, 3, 8, 4, 4, 101, 104,
		3, 12, 6, 0, 102, 104, 3, 10, 5, 0, 103, 90, 1, 0, 0, 0, 103, 95, 1, 0,
		0, 0, 103, 97, 1, 0, 0, 0, 103, 99, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103,
		102, 1, 0, 0, 0, 104, 121, 1, 0, 0, 0, 105, 106, 10, 10, 0, 0, 106, 107,
		7, 2, 0, 0, 107, 120, 3, 8, 4, 11, 108, 109, 10, 9, 0, 0, 109, 110, 7,
		1, 0, 0, 110, 120, 3, 8, 4, 10, 111, 112, 10, 8, 0, 0, 112, 113, 7, 3,
		0, 0, 113, 120, 3, 8, 4, 9, 114, 115, 10, 7, 0, 0, 115, 116, 7, 4, 0, 0,
		116, 120, 3, 8, 4, 8, 117, 118, 10, 3, 0, 0, 118, 120, 7, 0, 0, 0, 119,
		105, 1, 0, 0, 0, 119, 108, 1, 0, 0, 0, 119, 111, 1, 0, 0, 0, 119, 114,
		1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 120, 123, 1, 0, 0, 0, 121, 119, 1, 0,
		0, 0, 121, 122, 1, 0, 0, 0, 122, 9, 1, 0, 0, 0, 123, 121, 1, 0, 0, 0, 124,
		131, 3, 42, 21, 0, 125, 131, 3, 56, 28, 0, 126, 131, 3, 44, 22, 0, 127,
		131, 3, 26, 13, 0, 128, 131, 3, 22, 11, 0, 129, 131, 3, 30, 15, 0, 130,
		124, 1, 0, 0, 0, 130, 125, 1, 0, 0, 0, 130, 126, 1, 0, 0, 0, 130, 127,
		1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 130, 129, 1, 0, 0, 0, 131, 11, 1, 0,
		0, 0, 132, 135, 3, 10, 5, 0, 133, 136, 3, 14, 7, 0, 134, 136, 3, 16, 8,
		0, 135, 133, 1, 0, 0, 0, 135, 134, 1, 0, 0, 0, 136, 13, 1, 0, 0, 0, 137,
		149, 5, 3, 0, 0, 138, 143, 3, 8, 4, 0, 139, 140, 5, 20, 0, 0, 140, 142,
		3, 8, 4, 0, 141, 139, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 143, 144, 1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		146, 148, 5, 20, 0, 0, 147, 146, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148,
		150, 1, 0, 0, 0, 149, 138, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 151,
		1, 0, 0, 0, 151, 152, 5, 4, 0, 0, 152, 15, 1, 0, 0, 0, 153, 154, 5, 21,
		0, 0, 154, 160, 5, 46, 0, 0, 155, 156, 5, 22, 0, 0, 156, 157, 3, 8, 4,
		0, 157, 158, 5, 23, 0, 0, 158, 160, 1, 0, 0, 0, 159, 153, 1, 0, 0, 0, 159,
		155, 1, 0, 0, 0, 160, 17, 1, 0, 0, 0, 161, 172, 5, 54, 0, 0, 162, 172,
		5, 55, 0, 0, 163, 172, 3, 40, 20, 0, 164, 172, 3, 34, 17, 0, 165, 172,
		3, 36, 18, 0, 166, 172, 3, 38, 19, 0, 167, 172, 3, 32, 16, 0, 168, 172,
		3, 20, 10, 0, 169, 172, 3, 28, 14, 0, 170, 172, 3, 24, 12, 0, 171, 161,
		1, 0, 0, 0, 171, 162, 1, 0, 0, 0, 171, 163, 1, 0, 0, 0, 171, 164, 1, 0,
		0, 0, 171, 165, 1, 0, 0, 0, 171, 166, 1, 0, 0, 0, 171, 167, 1, 0, 0, 0,
		171, 168, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171, 170, 1, 0, 0, 0, 172,
		19, 1, 0, 0, 0, 173, 174, 5, 24, 0, 0, 174, 175, 3, 8, 4, 0, 175, 180,
		3, 2, 1, 0, 176, 177, 5, 25, 0, 0, 177, 179, 3, 20, 10, 0, 178, 176, 1,
		0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0, 0,
		0, 181, 185, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 184, 5, 25, 0, 0, 184,
		186, 3, 2, 1, 0, 185, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 21, 1,
		0, 0, 0, 187, 188, 5, 24, 0, 0, 188, 189, 3, 8, 4, 0, 189, 194, 3, 8, 4,
		0, 190, 191, 5, 25, 0, 0, 191, 193, 3, 22, 11, 0, 192, 190, 1, 0, 0, 0,
		193, 196, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195,
		199, 1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 197, 198, 5, 25, 0, 0, 198, 200,
		3, 8, 4, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 23, 1, 0,
		0, 0, 201, 202, 5, 26, 0, 0, 202, 203, 3, 8, 4, 0, 203, 214, 5, 27, 0,
		0, 204, 205, 5, 28, 0, 0, 205, 206, 3, 8, 4, 0, 206, 207, 5, 29, 0, 0,
		207, 208, 3, 2, 1, 0, 208, 213, 1, 0, 0, 0, 209, 210, 5, 30, 0, 0, 210,
		211, 5, 29, 0, 0, 211, 213, 3, 2, 1, 0, 212, 204, 1, 0, 0, 0, 212, 209,
		1, 0, 0, 0, 213, 216, 1, 0, 0, 0, 214, 212, 1, 0, 0, 0, 214, 215, 1, 0,
		0, 0, 215, 217, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 217, 218, 5, 31, 0, 0,
		218, 25, 1, 0, 0, 0, 219, 220, 5, 26, 0, 0, 220, 221, 3, 8, 4, 0, 221,
		232, 5, 27, 0, 0, 222, 223, 5, 28, 0, 0, 223, 224, 3, 8, 4, 0, 224, 225,
		5, 29, 0, 0, 225, 226, 3, 8, 4, 0, 226, 231, 1, 0, 0, 0, 227, 228, 5, 30,
		0, 0, 228, 229, 5, 29, 0, 0, 229, 231, 3, 8, 4, 0, 230, 222, 1, 0, 0, 0,
		230, 227, 1, 0, 0, 0, 231, 234, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 232,
		233, 1, 0, 0, 0, 233, 235, 1, 0, 0, 0, 234, 232, 1, 0, 0, 0, 235, 236,
		5, 31, 0, 0, 236, 27, 1, 0, 0, 0, 237, 238, 5, 32, 0, 0, 238, 239, 3, 8,
		4, 0, 239, 240, 3, 4, 2, 0, 240, 29, 1, 0, 0, 0, 241, 242, 5, 32, 0, 0,
		242, 243, 3, 8, 4, 0, 243, 244, 3, 6, 3, 0, 244, 31, 1, 0, 0, 0, 245, 248,
		5, 33, 0, 0, 246, 247, 5, 34, 0, 0, 247, 249, 3, 8, 4, 0, 248, 246, 1,
		0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 33, 1, 0, 0, 0, 250, 251, 5, 35, 0,
		0, 251, 252, 3, 52, 26, 0, 252, 254, 5, 29, 0, 0, 253, 255, 5, 36, 0, 0,
		254, 253, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		257, 3, 48, 24, 0, 257, 260, 1, 0, 0, 0, 258, 259, 5, 37, 0, 0, 259, 261,
		3, 8, 4, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 274, 1, 0,
		0, 0, 262, 263, 5, 38, 0, 0, 263, 264, 3, 52, 26, 0, 264, 266, 5, 29, 0,
		0, 265, 267, 5, 36, 0, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267,
		268, 1, 0, 0, 0, 268, 269, 3, 48, 24, 0, 269, 270, 1, 0, 0, 0, 270, 271,
		5, 37, 0, 0, 271, 272, 3, 8, 4, 0, 272, 274, 1, 0, 0, 0, 273, 250, 1, 0,
		0, 0, 273, 262, 1, 0, 0, 0, 274, 35, 1, 0, 0, 0, 275, 276, 5, 39, 0, 0,
		276, 279, 5, 46, 0, 0, 277, 278, 5, 37, 0, 0, 278, 280, 3, 48, 24, 0, 279,
		277, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 37, 1, 0, 0, 0, 281, 282, 5,
		40, 0, 0, 282, 302, 5, 46, 0, 0, 283, 299, 5, 27, 0, 0, 284, 285, 5, 46,
		0, 0, 285, 286, 5, 29, 0, 0, 286, 293, 3, 50, 25, 0, 287, 288, 5, 20, 0,
		0, 288, 289, 5, 46, 0, 0, 289, 290, 5, 29, 0, 0, 290, 292, 3, 50, 25, 0,
		291, 287, 1, 0, 0, 0, 292, 295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 297, 1, 0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 298,
		5, 20, 0, 0, 297, 296, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 1, 0,
		0, 0, 299, 284, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0,
		301, 303, 5, 31, 0, 0, 302, 283, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303,
		39, 1, 0, 0, 0, 304, 305, 3, 8, 4, 0, 305, 306, 5, 37, 0, 0, 306, 307,
		3, 8, 4, 0, 307, 41, 1, 0, 0, 0, 308, 310, 5, 41, 0, 0, 309, 311, 5, 46,
		0, 0, 310, 309, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0,
		312, 324, 5, 3, 0, 0, 313, 318, 3, 54, 27, 0, 314, 315, 5, 20, 0, 0, 315,
		317, 3, 54, 27, 0, 316, 314, 1, 0, 0, 0, 317, 320, 1, 0, 0, 0, 318, 316,
		1, 0, 0, 0, 318, 319, 1, 0, 0, 0, 319, 322, 1, 0, 0, 0, 320, 318, 1, 0,
		0, 0, 321, 323, 5, 20, 0, 0, 322, 321, 1, 0, 0, 0, 322, 323, 1, 0, 0, 0,
		323, 325, 1, 0, 0, 0, 324, 313, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325,
		326, 1, 0, 0, 0, 326, 332, 5, 4, 0, 0, 327, 329, 5, 29, 0, 0, 328, 330,
		5, 36, 0, 0, 329, 328, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 331, 1, 0,
		0, 0, 331, 333, 3, 48, 24, 0, 332, 327, 1, 0, 0, 0, 332, 333, 1, 0, 0,
		0, 333, 336, 1, 0, 0, 0, 334, 337, 3, 2, 1, 0, 335, 337, 3, 44, 22, 0,
		336, 334, 1, 0, 0, 0, 336, 335, 1, 0, 0, 0, 337, 43, 1, 0, 0, 0, 338, 342,
		5, 27, 0, 0, 339, 341, 3, 2, 1, 0, 340, 339, 1, 0, 0, 0, 341, 344, 1, 0,
		0, 0, 342, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 345, 1, 0, 0, 0,
		344, 342, 1, 0, 0, 0, 345, 346, 5, 31, 0, 0, 346, 45, 1, 0, 0, 0, 347,
		348, 5, 21, 0, 0, 348, 354, 5, 46, 0, 0, 349, 350, 5, 22, 0, 0, 350, 351,
		3, 48, 24, 0, 351, 352, 5, 23, 0, 0, 352, 354, 1, 0, 0, 0, 353, 347, 1,
		0, 0, 0, 353, 349, 1, 0, 0, 0, 354, 47, 1, 0, 0, 0, 355, 380, 3, 56, 28,
		0, 356, 357, 5, 42, 0, 0, 357, 369, 5, 3, 0, 0, 358, 363, 3, 50, 25, 0,
		359, 360, 5, 20, 0, 0, 360, 362, 3, 50, 25, 0, 361, 359, 1, 0, 0, 0, 362,
		365, 1, 0, 0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 367,
		1, 0, 0, 0, 365, 363, 1, 0, 0, 0, 366, 368, 5, 20, 0, 0, 367, 366, 1, 0,
		0, 0, 367, 368, 1, 0, 0, 0, 368, 370, 1, 0, 0, 0, 369, 358, 1, 0, 0, 0,
		369, 370, 1, 0, 0, 0, 370, 371, 1, 0, 0, 0, 371, 377, 5, 4, 0, 0, 372,
		374, 5, 29, 0, 0, 373, 375, 5, 36, 0, 0, 374, 373, 1, 0, 0, 0, 374, 375,
		1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 378, 3, 48, 24, 0, 377, 372, 1,
		0, 0, 0, 377, 378, 1, 0, 0, 0, 378, 380, 1, 0, 0, 0, 379, 355, 1, 0, 0,
		0, 379, 356, 1, 0, 0, 0, 380, 384, 1, 0, 0, 0, 381, 383, 3, 46, 23, 0,
		382, 381, 1, 0, 0, 0, 383, 386, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384,
		385, 1, 0, 0, 0, 385, 49, 1, 0, 0, 0, 386, 384, 1, 0, 0, 0, 387, 389, 3,
		48, 24, 0, 388, 390, 5, 36, 0, 0, 389, 388, 1, 0, 0, 0, 389, 390, 1, 0,
		0, 0, 390, 51, 1, 0, 0, 0, 391, 392, 5, 46, 0, 0, 392, 53, 1, 0, 0, 0,
		393, 399, 3, 52, 26, 0, 394, 396, 5, 29, 0, 0, 395, 397, 5, 36, 0, 0, 396,
		395, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 1, 0, 0, 0, 398, 400,
		3, 50, 25, 0, 399, 394, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 403, 1,
		0, 0, 0, 401, 402, 5, 37, 0, 0, 402, 404, 3, 8, 4, 0, 403, 401, 1, 0, 0,
		0, 403, 404, 1, 0, 0, 0, 404, 55, 1, 0, 0, 0, 405, 406, 7, 5, 0, 0, 406,
		57, 1, 0, 0, 0, 55, 61, 68, 73, 78, 80, 85, 87, 103, 119, 121, 130, 135,
		143, 147, 149, 159, 171, 180, 185, 194, 199, 212, 214, 230, 232, 248, 254,
		260, 266, 273, 279, 293, 297, 299, 302, 310, 318, 322, 324, 329, 332, 336,
		342, 353, 363, 367, 369, 374, 377, 379, 384, 389, 396, 399, 403,
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
	RW3CParserNULL           = 53
	RW3CParserMULTI_COMMENT  = 54
	RW3CParserSINGLE_COMMENT = 55
)

// RW3CParser rules.
const (
	RW3CParserRULE_prog               = 0
	RW3CParserRULE_run                = 1
	RW3CParserRULE_whileRun           = 2
	RW3CParserRULE_whileExprBody      = 3
	RW3CParserRULE_expr               = 4
	RW3CParserRULE_atomicExpr         = 5
	RW3CParserRULE_chainExpr          = 6
	RW3CParserRULE_callExpr           = 7
	RW3CParserRULE_accessPropExpr     = 8
	RW3CParserRULE_stmt               = 9
	RW3CParserRULE_ifStmt             = 10
	RW3CParserRULE_ifExpr             = 11
	RW3CParserRULE_switchStmt         = 12
	RW3CParserRULE_switchExpr         = 13
	RW3CParserRULE_whileStmt          = 14
	RW3CParserRULE_whileExpr          = 15
	RW3CParserRULE_retStmt            = 16
	RW3CParserRULE_defVarStmt         = 17
	RW3CParserRULE_defTypeStmt        = 18
	RW3CParserRULE_structStmt         = 19
	RW3CParserRULE_assignStmt         = 20
	RW3CParserRULE_fnExpr             = 21
	RW3CParserRULE_scope              = 22
	RW3CParserRULE_accessPropTypeExpr = 23
	RW3CParserRULE_typeExpr           = 24
	RW3CParserRULE_typeArg            = 25
	RW3CParserRULE_nameExpr           = 26
	RW3CParserRULE_arg                = 27
	RW3CParserRULE_litExpr            = 28
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

	// GetExprList returns the exprList rule context list.
	GetExprList() []IRunContext

	// SetExprList sets the exprList rule context list.
	SetExprList([]IRunContext)

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_run     IRunContext
	exprList []IRunContext
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

func (s *ProgContext) GetExprList() []IRunContext { return s.exprList }

func (s *ProgContext) SetExprList(v []IRunContext) { s.exprList = v }

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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69598858623975816) != 0 {
		{
			p.SetState(58)

			var _x = p.Run()

			localctx.(*ProgContext)._run = _x
		}
		localctx.(*ProgContext).exprList = append(localctx.(*ProgContext).exprList, localctx.(*ProgContext)._run)

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
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
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(66)
			p.Stmt()
		}

	case 2:
		{
			p.SetState(67)
			p.expr(0)
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(70)
				p.Match(RW3CParserT__0)
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IWhileRunContext is an interface to support dynamic dispatch.
type IWhileRunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileRunContext differentiates from other interfaces.
	IsWhileRunContext()
}

type WhileRunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileRunContext() *WhileRunContext {
	var p = new(WhileRunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_whileRun
	return p
}

func (*WhileRunContext) IsWhileRunContext() {}

func NewWhileRunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileRunContext {
	var p = new(WhileRunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_whileRun

	return p
}

func (s *WhileRunContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileRunContext) AllRun() []IRunContext {
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

func (s *WhileRunContext) Run(i int) IRunContext {
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

func (s *WhileRunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileRunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileRunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhileRun(s)
	}
}

func (s *WhileRunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhileRun(s)
	}
}

func (p *RW3CParser) WhileRun() (localctx IWhileRunContext) {
	this := p
	_ = this

	localctx = NewWhileRunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, RW3CParserRULE_whileRun)

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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(78)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RW3CParserT__2, RW3CParserT__6, RW3CParserT__7, RW3CParserT__16, RW3CParserT__17, RW3CParserT__18, RW3CParserT__23, RW3CParserT__25, RW3CParserT__26, RW3CParserT__31, RW3CParserT__32, RW3CParserT__34, RW3CParserT__37, RW3CParserT__38, RW3CParserT__39, RW3CParserT__40, RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL, RW3CParserNULL, RW3CParserMULTI_COMMENT, RW3CParserSINGLE_COMMENT:
				{
					p.SetState(76)
					p.Run()
				}

			case RW3CParserT__1:
				{
					p.SetState(77)
					p.Match(RW3CParserT__1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IWhileExprBodyContext is an interface to support dynamic dispatch.
type IWhileExprBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileExprBodyContext differentiates from other interfaces.
	IsWhileExprBodyContext()
}

type WhileExprBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprBodyContext() *WhileExprBodyContext {
	var p = new(WhileExprBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_whileExprBody
	return p
}

func (*WhileExprBodyContext) IsWhileExprBodyContext() {}

func NewWhileExprBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprBodyContext {
	var p = new(WhileExprBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_whileExprBody

	return p
}

func (s *WhileExprBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileExprBodyContext) AllExpr() []IExprContext {
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

func (s *WhileExprBodyContext) Expr(i int) IExprContext {
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

func (s *WhileExprBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileExprBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhileExprBody(s)
	}
}

func (s *WhileExprBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhileExprBody(s)
	}
}

func (p *RW3CParser) WhileExprBody() (localctx IWhileExprBodyContext) {
	this := p
	_ = this

	localctx = NewWhileExprBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, RW3CParserRULE_whileExprBody)

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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(85)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case RW3CParserT__2, RW3CParserT__6, RW3CParserT__7, RW3CParserT__16, RW3CParserT__17, RW3CParserT__18, RW3CParserT__23, RW3CParserT__25, RW3CParserT__26, RW3CParserT__31, RW3CParserT__40, RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL, RW3CParserNULL:
				{
					p.SetState(83)
					p.expr(0)
				}

			case RW3CParserT__1:
				{
					p.SetState(84)
					p.Match(RW3CParserT__1)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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

func (s *ExprContext) ChainExpr() IChainExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChainExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChainExprContext)
}

func (s *ExprContext) AtomicExpr() IAtomicExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomicExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomicExprContext)
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
	_startState := 8
	p.EnterRecursionRule(localctx, 8, RW3CParserRULE_expr, _p)
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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(91)
			p.Match(RW3CParserT__2)
		}
		{
			p.SetState(92)
			p.expr(0)
		}
		{
			p.SetState(93)
			p.Match(RW3CParserT__3)
		}

	case 2:
		{
			p.SetState(95)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RW3CParserT__16 || _la == RW3CParserT__17) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(96)
			p.expr(6)
		}

	case 3:
		{
			p.SetState(97)
			_la = p.GetTokenStream().LA(1)

			if !(_la == RW3CParserT__6 || _la == RW3CParserT__7) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(98)
			p.expr(5)
		}

	case 4:
		{
			p.SetState(99)
			p.Match(RW3CParserT__18)
		}
		{
			p.SetState(100)
			p.expr(4)
		}

	case 5:
		{
			p.SetState(101)
			p.ChainExpr()
		}

	case 6:
		{
			p.SetState(102)
			p.AtomicExpr()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(106)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__4 || _la == RW3CParserT__5) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)
					p.expr(11)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(109)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__6 || _la == RW3CParserT__7) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(110)
					p.expr(10)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(111)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(112)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&32256) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(113)
					p.expr(9)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(115)
					_la = p.GetTokenStream().LA(1)

					if !(_la == RW3CParserT__14 || _la == RW3CParserT__15) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(116)
					p.expr(8)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, RW3CParserRULE_expr)
				p.SetState(117)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(118)
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
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomicExprContext is an interface to support dynamic dispatch.
type IAtomicExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicExprContext differentiates from other interfaces.
	IsAtomicExprContext()
}

type AtomicExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicExprContext() *AtomicExprContext {
	var p = new(AtomicExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_atomicExpr
	return p
}

func (*AtomicExprContext) IsAtomicExprContext() {}

func NewAtomicExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicExprContext {
	var p = new(AtomicExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_atomicExpr

	return p
}

func (s *AtomicExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicExprContext) FnExpr() IFnExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFnExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFnExprContext)
}

func (s *AtomicExprContext) LitExpr() ILitExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILitExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILitExprContext)
}

func (s *AtomicExprContext) Scope() IScopeContext {
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

func (s *AtomicExprContext) SwitchExpr() ISwitchExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchExprContext)
}

func (s *AtomicExprContext) IfExpr() IIfExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfExprContext)
}

func (s *AtomicExprContext) WhileExpr() IWhileExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileExprContext)
}

func (s *AtomicExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAtomicExpr(s)
	}
}

func (s *AtomicExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAtomicExpr(s)
	}
}

func (p *RW3CParser) AtomicExpr() (localctx IAtomicExprContext) {
	this := p
	_ = this

	localctx = NewAtomicExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, RW3CParserRULE_atomicExpr)

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

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__40:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.FnExpr()
		}

	case RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL, RW3CParserNULL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.LitExpr()
		}

	case RW3CParserT__26:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.Scope()
		}

	case RW3CParserT__25:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.SwitchExpr()
		}

	case RW3CParserT__23:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.IfExpr()
		}

	case RW3CParserT__31:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.WhileExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IChainExprContext is an interface to support dynamic dispatch.
type IChainExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChainExprContext differentiates from other interfaces.
	IsChainExprContext()
}

type ChainExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainExprContext() *ChainExprContext {
	var p = new(ChainExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_chainExpr
	return p
}

func (*ChainExprContext) IsChainExprContext() {}

func NewChainExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainExprContext {
	var p = new(ChainExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_chainExpr

	return p
}

func (s *ChainExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainExprContext) AtomicExpr() IAtomicExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAtomicExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAtomicExprContext)
}

func (s *ChainExprContext) CallExpr() ICallExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallExprContext)
}

func (s *ChainExprContext) AccessPropExpr() IAccessPropExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessPropExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessPropExprContext)
}

func (s *ChainExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterChainExpr(s)
	}
}

func (s *ChainExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitChainExpr(s)
	}
}

func (p *RW3CParser) ChainExpr() (localctx IChainExprContext) {
	this := p
	_ = this

	localctx = NewChainExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, RW3CParserRULE_chainExpr)

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
		p.SetState(132)
		p.AtomicExpr()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__2:
		{
			p.SetState(133)
			p.CallExpr()
		}

	case RW3CParserT__20, RW3CParserT__21:
		{
			p.SetState(134)
			p.AccessPropExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICallExprContext is an interface to support dynamic dispatch.
type ICallExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCallExprContext differentiates from other interfaces.
	IsCallExprContext()
}

type CallExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallExprContext() *CallExprContext {
	var p = new(CallExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_callExpr
	return p
}

func (*CallExprContext) IsCallExprContext() {}

func NewCallExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallExprContext {
	var p = new(CallExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_callExpr

	return p
}

func (s *CallExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CallExprContext) AllExpr() []IExprContext {
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

func (s *CallExprContext) Expr(i int) IExprContext {
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

func (s *CallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterCallExpr(s)
	}
}

func (s *CallExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitCallExpr(s)
	}
}

func (p *RW3CParser) CallExpr() (localctx ICallExprContext) {
	this := p
	_ = this

	localctx = NewCallExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, RW3CParserRULE_callExpr)
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
		p.SetState(137)
		p.Match(RW3CParserT__2)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15553696000508296) != 0 {
		{
			p.SetState(138)
			p.expr(0)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(139)
					p.Match(RW3CParserT__19)
				}
				{
					p.SetState(140)
					p.expr(0)
				}

			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__19 {
			{
				p.SetState(146)
				p.Match(RW3CParserT__19)
			}

		}

	}
	{
		p.SetState(151)
		p.Match(RW3CParserT__3)
	}

	return localctx
}

// IAccessPropExprContext is an interface to support dynamic dispatch.
type IAccessPropExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessPropExprContext differentiates from other interfaces.
	IsAccessPropExprContext()
}

type AccessPropExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessPropExprContext() *AccessPropExprContext {
	var p = new(AccessPropExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_accessPropExpr
	return p
}

func (*AccessPropExprContext) IsAccessPropExprContext() {}

func NewAccessPropExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessPropExprContext {
	var p = new(AccessPropExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_accessPropExpr

	return p
}

func (s *AccessPropExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessPropExprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *AccessPropExprContext) Expr() IExprContext {
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

func (s *AccessPropExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessPropExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessPropExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAccessPropExpr(s)
	}
}

func (s *AccessPropExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAccessPropExpr(s)
	}
}

func (p *RW3CParser) AccessPropExpr() (localctx IAccessPropExprContext) {
	this := p
	_ = this

	localctx = NewAccessPropExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, RW3CParserRULE_accessPropExpr)

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

	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(RW3CParserT__20)
		}
		{
			p.SetState(154)
			p.Match(RW3CParserIDENTIF)
		}

	case RW3CParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(RW3CParserT__21)
		}
		{
			p.SetState(156)
			p.expr(0)
		}
		{
			p.SetState(157)
			p.Match(RW3CParserT__22)
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

func (s *StmtContext) AssignStmt() IAssignStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignStmtContext)
}

func (s *StmtContext) DefVarStmt() IDefVarStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefVarStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefVarStmtContext)
}

func (s *StmtContext) DefTypeStmt() IDefTypeStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefTypeStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefTypeStmtContext)
}

func (s *StmtContext) StructStmt() IStructStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructStmtContext)
}

func (s *StmtContext) RetStmt() IRetStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRetStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRetStmtContext)
}

func (s *StmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StmtContext) WhileStmt() IWhileStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStmtContext)
}

func (s *StmtContext) SwitchStmt() ISwitchStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStmtContext)
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
	p.EnterRule(localctx, 18, RW3CParserRULE_stmt)

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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(RW3CParserMULTI_COMMENT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(RW3CParserSINGLE_COMMENT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.AssignStmt()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.DefVarStmt()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.DefTypeStmt()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.StructStmt()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(167)
			p.RetStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)
			p.IfStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.WhileStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(170)
			p.SwitchStmt()
		}

	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) Expr() IExprContext {
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

func (s *IfStmtContext) AllRun() []IRunContext {
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

func (s *IfStmtContext) Run(i int) IRunContext {
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

func (s *IfStmtContext) AllIfStmt() []IIfStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfStmtContext); ok {
			len++
		}
	}

	tst := make([]IIfStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfStmtContext); ok {
			tst[i] = t.(IIfStmtContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) IfStmt(i int) IIfStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
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

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (p *RW3CParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, RW3CParserRULE_ifStmt)

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
		p.SetState(173)
		p.Match(RW3CParserT__23)
	}
	{
		p.SetState(174)
		p.expr(0)
	}
	{
		p.SetState(175)
		p.Run()
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(176)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(177)
				p.IfStmt()
			}

		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)
			p.Match(RW3CParserT__24)
		}
		{
			p.SetState(184)
			p.Run()
		}

	}

	return localctx
}

// IIfExprContext is an interface to support dynamic dispatch.
type IIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExprContext differentiates from other interfaces.
	IsIfExprContext()
}

type IfExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExprContext() *IfExprContext {
	var p = new(IfExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_ifExpr
	return p
}

func (*IfExprContext) IsIfExprContext() {}

func NewIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExprContext {
	var p = new(IfExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_ifExpr

	return p
}

func (s *IfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExprContext) AllExpr() []IExprContext {
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

func (s *IfExprContext) Expr(i int) IExprContext {
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

func (s *IfExprContext) AllIfExpr() []IIfExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfExprContext); ok {
			len++
		}
	}

	tst := make([]IIfExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfExprContext); ok {
			tst[i] = t.(IIfExprContext)
			i++
		}
	}

	return tst
}

func (s *IfExprContext) IfExpr(i int) IIfExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfExprContext); ok {
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

	return t.(IIfExprContext)
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (p *RW3CParser) IfExpr() (localctx IIfExprContext) {
	this := p
	_ = this

	localctx = NewIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, RW3CParserRULE_ifExpr)

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
		p.SetState(187)
		p.Match(RW3CParserT__23)
	}
	{
		p.SetState(188)
		p.expr(0)
	}
	{
		p.SetState(189)
		p.expr(0)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(190)
				p.Match(RW3CParserT__24)
			}
			{
				p.SetState(191)
				p.IfExpr()
			}

		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(197)
			p.Match(RW3CParserT__24)
		}
		{
			p.SetState(198)
			p.expr(0)
		}

	}

	return localctx
}

// ISwitchStmtContext is an interface to support dynamic dispatch.
type ISwitchStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStmtContext differentiates from other interfaces.
	IsSwitchStmtContext()
}

type SwitchStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStmtContext() *SwitchStmtContext {
	var p = new(SwitchStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_switchStmt
	return p
}

func (*SwitchStmtContext) IsSwitchStmtContext() {}

func NewSwitchStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStmtContext {
	var p = new(SwitchStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_switchStmt

	return p
}

func (s *SwitchStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStmtContext) AllExpr() []IExprContext {
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

func (s *SwitchStmtContext) Expr(i int) IExprContext {
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

func (s *SwitchStmtContext) AllRun() []IRunContext {
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

func (s *SwitchStmtContext) Run(i int) IRunContext {
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

func (s *SwitchStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterSwitchStmt(s)
	}
}

func (s *SwitchStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitSwitchStmt(s)
	}
}

func (p *RW3CParser) SwitchStmt() (localctx ISwitchStmtContext) {
	this := p
	_ = this

	localctx = NewSwitchStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, RW3CParserRULE_switchStmt)
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
		p.SetState(201)
		p.Match(RW3CParserT__25)
	}
	{
		p.SetState(202)
		p.expr(0)
	}
	{
		p.SetState(203)
		p.Match(RW3CParserT__26)
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RW3CParserT__27 || _la == RW3CParserT__29 {
		p.SetState(212)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RW3CParserT__27:
			{
				p.SetState(204)
				p.Match(RW3CParserT__27)
			}
			{
				p.SetState(205)
				p.expr(0)
			}
			{
				p.SetState(206)
				p.Match(RW3CParserT__28)
			}
			{
				p.SetState(207)
				p.Run()
			}

		case RW3CParserT__29:
			{
				p.SetState(209)
				p.Match(RW3CParserT__29)
			}
			{
				p.SetState(210)
				p.Match(RW3CParserT__28)
			}
			{
				p.SetState(211)
				p.Run()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
		p.Match(RW3CParserT__30)
	}

	return localctx
}

// ISwitchExprContext is an interface to support dynamic dispatch.
type ISwitchExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchExprContext differentiates from other interfaces.
	IsSwitchExprContext()
}

type SwitchExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchExprContext() *SwitchExprContext {
	var p = new(SwitchExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_switchExpr
	return p
}

func (*SwitchExprContext) IsSwitchExprContext() {}

func NewSwitchExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchExprContext {
	var p = new(SwitchExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_switchExpr

	return p
}

func (s *SwitchExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchExprContext) AllExpr() []IExprContext {
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

func (s *SwitchExprContext) Expr(i int) IExprContext {
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

func (s *SwitchExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterSwitchExpr(s)
	}
}

func (s *SwitchExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitSwitchExpr(s)
	}
}

func (p *RW3CParser) SwitchExpr() (localctx ISwitchExprContext) {
	this := p
	_ = this

	localctx = NewSwitchExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, RW3CParserRULE_switchExpr)
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
		p.SetState(219)
		p.Match(RW3CParserT__25)
	}
	{
		p.SetState(220)
		p.expr(0)
	}
	{
		p.SetState(221)
		p.Match(RW3CParserT__26)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == RW3CParserT__27 || _la == RW3CParserT__29 {
		p.SetState(230)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case RW3CParserT__27:
			{
				p.SetState(222)
				p.Match(RW3CParserT__27)
			}
			{
				p.SetState(223)
				p.expr(0)
			}
			{
				p.SetState(224)
				p.Match(RW3CParserT__28)
			}
			{
				p.SetState(225)
				p.expr(0)
			}

		case RW3CParserT__29:
			{
				p.SetState(227)
				p.Match(RW3CParserT__29)
			}
			{
				p.SetState(228)
				p.Match(RW3CParserT__28)
			}
			{
				p.SetState(229)
				p.expr(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(235)
		p.Match(RW3CParserT__30)
	}

	return localctx
}

// IWhileStmtContext is an interface to support dynamic dispatch.
type IWhileStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStmtContext differentiates from other interfaces.
	IsWhileStmtContext()
}

type WhileStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStmtContext() *WhileStmtContext {
	var p = new(WhileStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_whileStmt
	return p
}

func (*WhileStmtContext) IsWhileStmtContext() {}

func NewWhileStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStmtContext {
	var p = new(WhileStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_whileStmt

	return p
}

func (s *WhileStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStmtContext) Expr() IExprContext {
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

func (s *WhileStmtContext) WhileRun() IWhileRunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileRunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileRunContext)
}

func (s *WhileStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhileStmt(s)
	}
}

func (s *WhileStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhileStmt(s)
	}
}

func (p *RW3CParser) WhileStmt() (localctx IWhileStmtContext) {
	this := p
	_ = this

	localctx = NewWhileStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, RW3CParserRULE_whileStmt)

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
		p.SetState(237)
		p.Match(RW3CParserT__31)
	}
	{
		p.SetState(238)
		p.expr(0)
	}
	{
		p.SetState(239)
		p.WhileRun()
	}

	return localctx
}

// IWhileExprContext is an interface to support dynamic dispatch.
type IWhileExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileExprContext differentiates from other interfaces.
	IsWhileExprContext()
}

type WhileExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprContext() *WhileExprContext {
	var p = new(WhileExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_whileExpr
	return p
}

func (*WhileExprContext) IsWhileExprContext() {}

func NewWhileExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprContext {
	var p = new(WhileExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_whileExpr

	return p
}

func (s *WhileExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileExprContext) Expr() IExprContext {
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

func (s *WhileExprContext) WhileExprBody() IWhileExprBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileExprBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileExprBodyContext)
}

func (s *WhileExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterWhileExpr(s)
	}
}

func (s *WhileExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitWhileExpr(s)
	}
}

func (p *RW3CParser) WhileExpr() (localctx IWhileExprContext) {
	this := p
	_ = this

	localctx = NewWhileExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, RW3CParserRULE_whileExpr)

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
		p.SetState(241)
		p.Match(RW3CParserT__31)
	}
	{
		p.SetState(242)
		p.expr(0)
	}
	{
		p.SetState(243)
		p.WhileExprBody()
	}

	return localctx
}

// IRetStmtContext is an interface to support dynamic dispatch.
type IRetStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetStmtContext differentiates from other interfaces.
	IsRetStmtContext()
}

type RetStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetStmtContext() *RetStmtContext {
	var p = new(RetStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_retStmt
	return p
}

func (*RetStmtContext) IsRetStmtContext() {}

func NewRetStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RetStmtContext {
	var p = new(RetStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_retStmt

	return p
}

func (s *RetStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *RetStmtContext) Expr() IExprContext {
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

func (s *RetStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RetStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RetStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterRetStmt(s)
	}
}

func (s *RetStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitRetStmt(s)
	}
}

func (p *RW3CParser) RetStmt() (localctx IRetStmtContext) {
	this := p
	_ = this

	localctx = NewRetStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, RW3CParserRULE_retStmt)

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
		p.SetState(245)
		p.Match(RW3CParserT__32)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(246)
			p.Match(RW3CParserT__33)
		}
		{
			p.SetState(247)
			p.expr(0)
		}

	}

	return localctx
}

// IDefVarStmtContext is an interface to support dynamic dispatch.
type IDefVarStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefVarStmtContext differentiates from other interfaces.
	IsDefVarStmtContext()
}

type DefVarStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefVarStmtContext() *DefVarStmtContext {
	var p = new(DefVarStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_defVarStmt
	return p
}

func (*DefVarStmtContext) IsDefVarStmtContext() {}

func NewDefVarStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefVarStmtContext {
	var p = new(DefVarStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_defVarStmt

	return p
}

func (s *DefVarStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefVarStmtContext) NameExpr() INameExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameExprContext)
}

func (s *DefVarStmtContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DefVarStmtContext) Expr() IExprContext {
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

func (s *DefVarStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefVarStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefVarStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterDefVarStmt(s)
	}
}

func (s *DefVarStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitDefVarStmt(s)
	}
}

func (p *RW3CParser) DefVarStmt() (localctx IDefVarStmtContext) {
	this := p
	_ = this

	localctx = NewDefVarStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, RW3CParserRULE_defVarStmt)
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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__34:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(RW3CParserT__34)
		}
		{
			p.SetState(251)
			p.NameExpr()
		}

		{
			p.SetState(252)
			p.Match(RW3CParserT__28)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__35 {
			{
				p.SetState(253)
				p.Match(RW3CParserT__35)
			}

		}
		{
			p.SetState(256)
			p.TypeExpr()
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(258)
				p.Match(RW3CParserT__36)
			}
			{
				p.SetState(259)
				p.expr(0)
			}

		}

	case RW3CParserT__37:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(RW3CParserT__37)
		}
		{
			p.SetState(263)
			p.NameExpr()
		}

		{
			p.SetState(264)
			p.Match(RW3CParserT__28)
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__35 {
			{
				p.SetState(265)
				p.Match(RW3CParserT__35)
			}

		}
		{
			p.SetState(268)
			p.TypeExpr()
		}

		{
			p.SetState(270)
			p.Match(RW3CParserT__36)
		}
		{
			p.SetState(271)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefTypeStmtContext is an interface to support dynamic dispatch.
type IDefTypeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefTypeStmtContext differentiates from other interfaces.
	IsDefTypeStmtContext()
}

type DefTypeStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefTypeStmtContext() *DefTypeStmtContext {
	var p = new(DefTypeStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_defTypeStmt
	return p
}

func (*DefTypeStmtContext) IsDefTypeStmtContext() {}

func NewDefTypeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefTypeStmtContext {
	var p = new(DefTypeStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_defTypeStmt

	return p
}

func (s *DefTypeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *DefTypeStmtContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *DefTypeStmtContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *DefTypeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefTypeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefTypeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterDefTypeStmt(s)
	}
}

func (s *DefTypeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitDefTypeStmt(s)
	}
}

func (p *RW3CParser) DefTypeStmt() (localctx IDefTypeStmtContext) {
	this := p
	_ = this

	localctx = NewDefTypeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, RW3CParserRULE_defTypeStmt)

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
		p.SetState(275)
		p.Match(RW3CParserT__38)
	}
	{
		p.SetState(276)
		p.Match(RW3CParserIDENTIF)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(277)
			p.Match(RW3CParserT__36)
		}
		{
			p.SetState(278)
			p.TypeExpr()
		}

	}

	return localctx
}

// IStructStmtContext is an interface to support dynamic dispatch.
type IStructStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructStmtContext differentiates from other interfaces.
	IsStructStmtContext()
}

type StructStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructStmtContext() *StructStmtContext {
	var p = new(StructStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_structStmt
	return p
}

func (*StructStmtContext) IsStructStmtContext() {}

func NewStructStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructStmtContext {
	var p = new(StructStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_structStmt

	return p
}

func (s *StructStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StructStmtContext) AllIDENTIF() []antlr.TerminalNode {
	return s.GetTokens(RW3CParserIDENTIF)
}

func (s *StructStmtContext) IDENTIF(i int) antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, i)
}

func (s *StructStmtContext) AllTypeArg() []ITypeArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeArgContext); ok {
			len++
		}
	}

	tst := make([]ITypeArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeArgContext); ok {
			tst[i] = t.(ITypeArgContext)
			i++
		}
	}

	return tst
}

func (s *StructStmtContext) TypeArg(i int) ITypeArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeArgContext); ok {
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

	return t.(ITypeArgContext)
}

func (s *StructStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterStructStmt(s)
	}
}

func (s *StructStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitStructStmt(s)
	}
}

func (p *RW3CParser) StructStmt() (localctx IStructStmtContext) {
	this := p
	_ = this

	localctx = NewStructStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, RW3CParserRULE_structStmt)
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
		p.SetState(281)
		p.Match(RW3CParserT__39)
	}
	{
		p.SetState(282)
		p.Match(RW3CParserIDENTIF)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
			p.Match(RW3CParserT__26)
		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserIDENTIF {
			{
				p.SetState(284)
				p.Match(RW3CParserIDENTIF)
			}
			{
				p.SetState(285)
				p.Match(RW3CParserT__28)
			}
			{
				p.SetState(286)
				p.TypeArg()
			}
			p.SetState(293)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(287)
						p.Match(RW3CParserT__19)
					}
					{
						p.SetState(288)
						p.Match(RW3CParserIDENTIF)
					}
					{
						p.SetState(289)
						p.Match(RW3CParserT__28)
					}
					{
						p.SetState(290)
						p.TypeArg()
					}

				}
				p.SetState(295)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
			}
			p.SetState(297)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__19 {
				{
					p.SetState(296)
					p.Match(RW3CParserT__19)
				}

			}

		}
		{
			p.SetState(301)
			p.Match(RW3CParserT__30)
		}

	}

	return localctx
}

// IAssignStmtContext is an interface to support dynamic dispatch.
type IAssignStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignStmtContext differentiates from other interfaces.
	IsAssignStmtContext()
}

type AssignStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignStmtContext() *AssignStmtContext {
	var p = new(AssignStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_assignStmt
	return p
}

func (*AssignStmtContext) IsAssignStmtContext() {}

func NewAssignStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignStmtContext {
	var p = new(AssignStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_assignStmt

	return p
}

func (s *AssignStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignStmtContext) AllExpr() []IExprContext {
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

func (s *AssignStmtContext) Expr(i int) IExprContext {
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

func (s *AssignStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAssignStmt(s)
	}
}

func (s *AssignStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAssignStmt(s)
	}
}

func (p *RW3CParser) AssignStmt() (localctx IAssignStmtContext) {
	this := p
	_ = this

	localctx = NewAssignStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, RW3CParserRULE_assignStmt)

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
		p.SetState(304)
		p.expr(0)
	}
	{
		p.SetState(305)
		p.Match(RW3CParserT__36)
	}
	{
		p.SetState(306)
		p.expr(0)
	}

	return localctx
}

// IFnExprContext is an interface to support dynamic dispatch.
type IFnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFnExprContext differentiates from other interfaces.
	IsFnExprContext()
}

type FnExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFnExprContext() *FnExprContext {
	var p = new(FnExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_fnExpr
	return p
}

func (*FnExprContext) IsFnExprContext() {}

func NewFnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FnExprContext {
	var p = new(FnExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_fnExpr

	return p
}

func (s *FnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FnExprContext) Run() IRunContext {
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

func (s *FnExprContext) Scope() IScopeContext {
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

func (s *FnExprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *FnExprContext) AllArg() []IArgContext {
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

func (s *FnExprContext) Arg(i int) IArgContext {
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

func (s *FnExprContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *FnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterFnExpr(s)
	}
}

func (s *FnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitFnExpr(s)
	}
}

func (p *RW3CParser) FnExpr() (localctx IFnExprContext) {
	this := p
	_ = this

	localctx = NewFnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, RW3CParserRULE_fnExpr)
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
		p.SetState(308)
		p.Match(RW3CParserT__40)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserIDENTIF {
		{
			p.SetState(309)
			p.Match(RW3CParserIDENTIF)
		}

	}
	{
		p.SetState(312)
		p.Match(RW3CParserT__2)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserIDENTIF {
		{
			p.SetState(313)
			p.Arg()
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(314)
					p.Match(RW3CParserT__19)
				}
				{
					p.SetState(315)
					p.Arg()
				}

			}
			p.SetState(320)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext())
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__19 {
			{
				p.SetState(321)
				p.Match(RW3CParserT__19)
			}

		}

	}
	{
		p.SetState(326)
		p.Match(RW3CParserT__3)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__28 {
		{
			p.SetState(327)
			p.Match(RW3CParserT__28)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__35 {
			{
				p.SetState(328)
				p.Match(RW3CParserT__35)
			}

		}
		{
			p.SetState(331)
			p.TypeExpr()
		}

	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(334)
			p.Run()
		}

	case 2:
		{
			p.SetState(335)
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
	p.EnterRule(localctx, 44, RW3CParserRULE_scope)
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
		p.SetState(338)
		p.Match(RW3CParserT__26)
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&69598858623975816) != 0 {
		{
			p.SetState(339)
			p.Run()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.Match(RW3CParserT__30)
	}

	return localctx
}

// IAccessPropTypeExprContext is an interface to support dynamic dispatch.
type IAccessPropTypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessPropTypeExprContext differentiates from other interfaces.
	IsAccessPropTypeExprContext()
}

type AccessPropTypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessPropTypeExprContext() *AccessPropTypeExprContext {
	var p = new(AccessPropTypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_accessPropTypeExpr
	return p
}

func (*AccessPropTypeExprContext) IsAccessPropTypeExprContext() {}

func NewAccessPropTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessPropTypeExprContext {
	var p = new(AccessPropTypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_accessPropTypeExpr

	return p
}

func (s *AccessPropTypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessPropTypeExprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *AccessPropTypeExprContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *AccessPropTypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessPropTypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessPropTypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterAccessPropTypeExpr(s)
	}
}

func (s *AccessPropTypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitAccessPropTypeExpr(s)
	}
}

func (p *RW3CParser) AccessPropTypeExpr() (localctx IAccessPropTypeExprContext) {
	this := p
	_ = this

	localctx = NewAccessPropTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, RW3CParserRULE_accessPropTypeExpr)

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

	p.SetState(353)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Match(RW3CParserT__20)
		}
		{
			p.SetState(348)
			p.Match(RW3CParserIDENTIF)
		}

	case RW3CParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.Match(RW3CParserT__21)
		}
		{
			p.SetState(350)
			p.TypeExpr()
		}
		{
			p.SetState(351)
			p.Match(RW3CParserT__22)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeExprContext is an interface to support dynamic dispatch.
type ITypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExprContext differentiates from other interfaces.
	IsTypeExprContext()
}

type TypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_typeExpr
	return p
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExprContext) LitExpr() ILitExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILitExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILitExprContext)
}

func (s *TypeExprContext) AllAccessPropTypeExpr() []IAccessPropTypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAccessPropTypeExprContext); ok {
			len++
		}
	}

	tst := make([]IAccessPropTypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAccessPropTypeExprContext); ok {
			tst[i] = t.(IAccessPropTypeExprContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) AccessPropTypeExpr(i int) IAccessPropTypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessPropTypeExprContext); ok {
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

	return t.(IAccessPropTypeExprContext)
}

func (s *TypeExprContext) AllTypeArg() []ITypeArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeArgContext); ok {
			len++
		}
	}

	tst := make([]ITypeArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeArgContext); ok {
			tst[i] = t.(ITypeArgContext)
			i++
		}
	}

	return tst
}

func (s *TypeExprContext) TypeArg(i int) ITypeArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeArgContext); ok {
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

	return t.(ITypeArgContext)
}

func (s *TypeExprContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterTypeExpr(s)
	}
}

func (s *TypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitTypeExpr(s)
	}
}

func (p *RW3CParser) TypeExpr() (localctx ITypeExprContext) {
	this := p
	_ = this

	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, RW3CParserRULE_typeExpr)
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
	p.SetState(379)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case RW3CParserIDENTIF, RW3CParserLONG, RW3CParserDOUBLE, RW3CParserSTR, RW3CParserBOOL, RW3CParserNULL:
		{
			p.SetState(355)
			p.LitExpr()
		}

	case RW3CParserT__41:
		{
			p.SetState(356)
			p.Match(RW3CParserT__41)
		}
		{
			p.SetState(357)
			p.Match(RW3CParserT__2)
		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15555890509774848) != 0 {
			{
				p.SetState(358)
				p.TypeArg()
			}
			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())

			for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				if _alt == 1 {
					{
						p.SetState(359)
						p.Match(RW3CParserT__19)
					}
					{
						p.SetState(360)
						p.TypeArg()
					}

				}
				p.SetState(365)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
			}
			p.SetState(367)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__19 {
				{
					p.SetState(366)
					p.Match(RW3CParserT__19)
				}

			}

		}
		{
			p.SetState(371)
			p.Match(RW3CParserT__3)
		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(372)
				p.Match(RW3CParserT__28)
			}
			p.SetState(374)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == RW3CParserT__35 {
				{
					p.SetState(373)
					p.Match(RW3CParserT__35)
				}

			}
			{
				p.SetState(376)
				p.TypeExpr()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(381)
				p.AccessPropTypeExpr()
			}

		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeArgContext is an interface to support dynamic dispatch.
type ITypeArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeArgContext differentiates from other interfaces.
	IsTypeArgContext()
}

type TypeArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeArgContext() *TypeArgContext {
	var p = new(TypeArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_typeArg
	return p
}

func (*TypeArgContext) IsTypeArgContext() {}

func NewTypeArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeArgContext {
	var p = new(TypeArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_typeArg

	return p
}

func (s *TypeArgContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeArgContext) TypeExpr() ITypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterTypeArg(s)
	}
}

func (s *TypeArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitTypeArg(s)
	}
}

func (p *RW3CParser) TypeArg() (localctx ITypeArgContext) {
	this := p
	_ = this

	localctx = NewTypeArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, RW3CParserRULE_typeArg)
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
		p.SetState(387)
		p.TypeExpr()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__35 {
		{
			p.SetState(388)
			p.Match(RW3CParserT__35)
		}

	}

	return localctx
}

// INameExprContext is an interface to support dynamic dispatch.
type INameExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameExprContext differentiates from other interfaces.
	IsNameExprContext()
}

type NameExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameExprContext() *NameExprContext {
	var p = new(NameExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_nameExpr
	return p
}

func (*NameExprContext) IsNameExprContext() {}

func NewNameExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameExprContext {
	var p = new(NameExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_nameExpr

	return p
}

func (s *NameExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NameExprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *NameExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterNameExpr(s)
	}
}

func (s *NameExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitNameExpr(s)
	}
}

func (p *RW3CParser) NameExpr() (localctx INameExprContext) {
	this := p
	_ = this

	localctx = NewNameExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, RW3CParserRULE_nameExpr)

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
		p.SetState(391)
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

func (s *ArgContext) NameExpr() INameExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameExprContext)
}

func (s *ArgContext) TypeArg() ITypeArgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeArgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeArgContext)
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
	p.EnterRule(localctx, 54, RW3CParserRULE_arg)
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
		p.SetState(393)
		p.NameExpr()
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__28 {
		{
			p.SetState(394)
			p.Match(RW3CParserT__28)
		}
		p.SetState(396)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == RW3CParserT__35 {
			{
				p.SetState(395)
				p.Match(RW3CParserT__35)
			}

		}
		{
			p.SetState(398)
			p.TypeArg()
		}

	}
	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == RW3CParserT__36 {
		{
			p.SetState(401)
			p.Match(RW3CParserT__36)
		}
		{
			p.SetState(402)
			p.expr(0)
		}

	}

	return localctx
}

// ILitExprContext is an interface to support dynamic dispatch.
type ILitExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLitExprContext differentiates from other interfaces.
	IsLitExprContext()
}

type LitExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLitExprContext() *LitExprContext {
	var p = new(LitExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = RW3CParserRULE_litExpr
	return p
}

func (*LitExprContext) IsLitExprContext() {}

func NewLitExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LitExprContext {
	var p = new(LitExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = RW3CParserRULE_litExpr

	return p
}

func (s *LitExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LitExprContext) LONG() antlr.TerminalNode {
	return s.GetToken(RW3CParserLONG, 0)
}

func (s *LitExprContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(RW3CParserDOUBLE, 0)
}

func (s *LitExprContext) STR() antlr.TerminalNode {
	return s.GetToken(RW3CParserSTR, 0)
}

func (s *LitExprContext) IDENTIF() antlr.TerminalNode {
	return s.GetToken(RW3CParserIDENTIF, 0)
}

func (s *LitExprContext) BOOL() antlr.TerminalNode {
	return s.GetToken(RW3CParserBOOL, 0)
}

func (s *LitExprContext) NULL() antlr.TerminalNode {
	return s.GetToken(RW3CParserNULL, 0)
}

func (s *LitExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LitExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.EnterLitExpr(s)
	}
}

func (s *LitExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(RW3CListener); ok {
		listenerT.ExitLitExpr(s)
	}
}

func (p *RW3CParser) LitExpr() (localctx ILitExprContext) {
	this := p
	_ = this

	localctx = NewLitExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, RW3CParserRULE_litExpr)
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
		p.SetState(405)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15551492463263744) != 0) {
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
	case 4:
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
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

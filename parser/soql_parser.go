// Code generated from soql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // soql

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 65, 263,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 3, 2, 3, 2, 5, 2, 56, 10,
	2, 3, 2, 5, 2, 59, 10, 2, 3, 2, 5, 2, 62, 10, 2, 3, 2, 5, 2, 65, 10, 2,
	3, 2, 5, 2, 68, 10, 2, 3, 2, 5, 2, 71, 10, 2, 3, 2, 5, 2, 74, 10, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 82, 10, 4, 12, 4, 14, 4, 85, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 96, 10,
	5, 13, 5, 14, 5, 97, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 104, 10, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 111, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7,
	8, 118, 10, 8, 12, 8, 14, 8, 121, 11, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 7, 8, 129, 10, 8, 12, 8, 14, 8, 132, 11, 8, 5, 8, 134, 10, 8, 3,
	8, 3, 8, 5, 8, 138, 10, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 151, 10, 11, 12, 11, 14, 11, 154, 11,
	11, 3, 12, 5, 12, 157, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 5, 12, 167, 10, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 7, 14, 177, 10, 14, 12, 14, 14, 14, 180, 11, 14, 3, 14,
	5, 14, 183, 10, 14, 3, 14, 3, 14, 5, 14, 187, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 194, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 208, 10, 18, 12,
	18, 14, 18, 211, 11, 18, 3, 18, 3, 18, 5, 18, 215, 10, 18, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 226, 10, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 7, 23, 241, 10, 23, 12, 23, 14, 23, 244, 11, 23, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 7, 24, 251, 10, 24, 12, 24, 14, 24, 254, 11, 24,
	3, 24, 5, 24, 257, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 2, 3, 20,
	27, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 2, 10, 3, 2, 59, 60, 5, 2, 3, 9, 44, 44, 57,
	57, 3, 2, 35, 36, 3, 2, 52, 53, 3, 2, 39, 40, 3, 2, 41, 42, 10, 2, 16,
	16, 39, 39, 43, 43, 47, 49, 54, 54, 58, 58, 62, 62, 64, 65, 4, 2, 13, 15,
	17, 18, 2, 266, 2, 52, 3, 2, 2, 2, 4, 75, 3, 2, 2, 2, 6, 78, 3, 2, 2, 2,
	8, 103, 3, 2, 2, 2, 10, 105, 3, 2, 2, 2, 12, 112, 3, 2, 2, 2, 14, 137,
	3, 2, 2, 2, 16, 139, 3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 144, 3, 2, 2,
	2, 22, 166, 3, 2, 2, 2, 24, 168, 3, 2, 2, 2, 26, 171, 3, 2, 2, 2, 28, 193,
	3, 2, 2, 2, 30, 195, 3, 2, 2, 2, 32, 200, 3, 2, 2, 2, 34, 202, 3, 2, 2,
	2, 36, 216, 3, 2, 2, 2, 38, 218, 3, 2, 2, 2, 40, 221, 3, 2, 2, 2, 42, 227,
	3, 2, 2, 2, 44, 231, 3, 2, 2, 2, 46, 245, 3, 2, 2, 2, 48, 258, 3, 2, 2,
	2, 50, 260, 3, 2, 2, 2, 52, 53, 5, 4, 3, 2, 53, 55, 5, 10, 6, 2, 54, 56,
	5, 18, 10, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58, 3, 2, 2,
	2, 57, 59, 5, 30, 16, 2, 58, 57, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61,
	3, 2, 2, 2, 60, 62, 5, 34, 18, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2,
	2, 62, 64, 3, 2, 2, 2, 63, 65, 5, 26, 14, 2, 64, 63, 3, 2, 2, 2, 64, 65,
	3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66, 68, 5, 24, 13, 2, 67, 66, 3, 2, 2,
	2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 71, 5, 38, 20, 2, 70, 69,
	3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 74, 5, 40, 21,
	2, 73, 72, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 3, 3, 2, 2, 2, 75, 76, 7,
	29, 2, 2, 76, 77, 5, 6, 4, 2, 77, 5, 3, 2, 2, 2, 78, 83, 5, 8, 5, 2, 79,
	80, 7, 24, 2, 2, 80, 82, 5, 8, 5, 2, 81, 79, 3, 2, 2, 2, 82, 85, 3, 2,
	2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 7, 3, 2, 2, 2, 85, 83,
	3, 2, 2, 2, 86, 104, 5, 14, 8, 2, 87, 104, 5, 16, 9, 2, 88, 89, 7, 38,
	2, 2, 89, 95, 5, 14, 8, 2, 90, 91, 7, 22, 2, 2, 91, 92, 5, 48, 25, 2, 92,
	93, 7, 58, 2, 2, 93, 94, 5, 6, 4, 2, 94, 96, 3, 2, 2, 2, 95, 90, 3, 2,
	2, 2, 96, 97, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99,
	3, 2, 2, 2, 99, 100, 7, 23, 2, 2, 100, 101, 5, 6, 4, 2, 101, 102, 7, 45,
	2, 2, 102, 104, 3, 2, 2, 2, 103, 86, 3, 2, 2, 2, 103, 87, 3, 2, 2, 2, 103,
	88, 3, 2, 2, 2, 104, 9, 3, 2, 2, 2, 105, 106, 7, 30, 2, 2, 106, 110, 5,
	48, 25, 2, 107, 108, 7, 46, 2, 2, 108, 109, 7, 54, 2, 2, 109, 111, 5, 12,
	7, 2, 110, 107, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 11, 3, 2, 2, 2,
	112, 113, 3, 2, 2, 2, 113, 13, 3, 2, 2, 2, 114, 115, 5, 48, 25, 2, 115,
	116, 7, 25, 2, 2, 116, 118, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2, 118, 121,
	3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 122, 3, 2,
	2, 2, 121, 119, 3, 2, 2, 2, 122, 138, 5, 48, 25, 2, 123, 124, 5, 48, 25,
	2, 124, 133, 7, 20, 2, 2, 125, 130, 5, 14, 8, 2, 126, 127, 7, 24, 2, 2,
	127, 129, 5, 14, 8, 2, 128, 126, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130,
	128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130,
	3, 2, 2, 2, 133, 125, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2,
	2, 2, 135, 136, 7, 21, 2, 2, 136, 138, 3, 2, 2, 2, 137, 119, 3, 2, 2, 2,
	137, 123, 3, 2, 2, 2, 138, 15, 3, 2, 2, 2, 139, 140, 5, 2, 2, 2, 140, 17,
	3, 2, 2, 2, 141, 142, 7, 31, 2, 2, 142, 143, 5, 20, 11, 2, 143, 19, 3,
	2, 2, 2, 144, 145, 8, 11, 1, 2, 145, 146, 5, 22, 12, 2, 146, 152, 3, 2,
	2, 2, 147, 148, 12, 3, 2, 2, 148, 149, 9, 2, 2, 2, 149, 151, 5, 20, 11,
	4, 150, 147, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 21, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 157, 7,
	61, 2, 2, 156, 155, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 158, 3, 2, 2,
	2, 158, 159, 5, 14, 8, 2, 159, 160, 9, 3, 2, 2, 160, 161, 5, 28, 15, 2,
	161, 167, 3, 2, 2, 2, 162, 163, 7, 20, 2, 2, 163, 164, 5, 20, 11, 2, 164,
	165, 7, 21, 2, 2, 165, 167, 3, 2, 2, 2, 166, 156, 3, 2, 2, 2, 166, 162,
	3, 2, 2, 2, 167, 23, 3, 2, 2, 2, 168, 169, 7, 32, 2, 2, 169, 170, 7, 17,
	2, 2, 170, 25, 3, 2, 2, 2, 171, 172, 7, 33, 2, 2, 172, 173, 7, 34, 2, 2,
	173, 178, 5, 14, 8, 2, 174, 175, 7, 10, 2, 2, 175, 177, 5, 14, 8, 2, 176,
	174, 3, 2, 2, 2, 177, 180, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179,
	3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 181, 183, 9, 4,
	2, 2, 182, 181, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2,
	184, 185, 7, 51, 2, 2, 185, 187, 9, 5, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 27, 3, 2, 2, 2, 188, 194, 5, 50, 26, 2, 189, 190,
	5, 48, 25, 2, 190, 191, 7, 28, 2, 2, 191, 192, 5, 50, 26, 2, 192, 194,
	3, 2, 2, 2, 193, 188, 3, 2, 2, 2, 193, 189, 3, 2, 2, 2, 194, 29, 3, 2,
	2, 2, 195, 196, 7, 37, 2, 2, 196, 197, 7, 47, 2, 2, 197, 198, 7, 48, 2,
	2, 198, 199, 5, 32, 17, 2, 199, 31, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	33, 3, 2, 2, 2, 202, 203, 7, 49, 2, 2, 203, 204, 7, 34, 2, 2, 204, 209,
	5, 14, 8, 2, 205, 206, 7, 10, 2, 2, 206, 208, 5, 14, 8, 2, 207, 205, 3,
	2, 2, 2, 208, 211, 3, 2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2,
	2, 210, 214, 3, 2, 2, 2, 211, 209, 3, 2, 2, 2, 212, 213, 7, 50, 2, 2, 213,
	215, 5, 36, 19, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 35,
	3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 37, 3, 2, 2, 2, 218, 219, 7, 43,
	2, 2, 219, 220, 7, 17, 2, 2, 220, 39, 3, 2, 2, 2, 221, 222, 7, 26, 2, 2,
	222, 225, 9, 6, 2, 2, 223, 224, 7, 27, 2, 2, 224, 226, 9, 7, 2, 2, 225,
	223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 41, 3, 2, 2, 2, 227, 228, 7,
	11, 2, 2, 228, 229, 5, 44, 23, 2, 229, 230, 7, 12, 2, 2, 230, 43, 3, 2,
	2, 2, 231, 232, 7, 62, 2, 2, 232, 233, 5, 50, 26, 2, 233, 234, 7, 44, 2,
	2, 234, 235, 7, 65, 2, 2, 235, 236, 7, 63, 2, 2, 236, 237, 7, 64, 2, 2,
	237, 242, 5, 46, 24, 2, 238, 239, 7, 10, 2, 2, 239, 241, 5, 46, 24, 2,
	240, 238, 3, 2, 2, 2, 241, 244, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242,
	243, 3, 2, 2, 2, 243, 45, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 245, 256, 7,
	16, 2, 2, 246, 247, 7, 20, 2, 2, 247, 252, 7, 16, 2, 2, 248, 249, 7, 10,
	2, 2, 249, 251, 7, 16, 2, 2, 250, 248, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254,
	252, 3, 2, 2, 2, 255, 257, 7, 21, 2, 2, 256, 246, 3, 2, 2, 2, 256, 257,
	3, 2, 2, 2, 257, 47, 3, 2, 2, 2, 258, 259, 9, 8, 2, 2, 259, 49, 3, 2, 2,
	2, 260, 261, 9, 9, 2, 2, 261, 51, 3, 2, 2, 2, 30, 55, 58, 61, 64, 67, 70,
	73, 83, 97, 103, 110, 119, 130, 133, 137, 152, 156, 166, 178, 182, 186,
	193, 209, 214, 225, 242, 252, 256,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "'<'", "'>'", "'<='", "'>='", "'!='", "'<>'", "','", "'['",
	"']'", "", "", "", "", "", "", "'''", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "StringLiteral", "BooleanLiteral",
	"NullLiteral", "Identifier", "IntegerLiteral", "FloatingPointLiteral",
	"QUOTE", "LPAREN", "RPAREN", "WHEN", "ELSE", "COMMA", "DOT", "FOR", "UPDATE",
	"COLON", "SELECT", "FROM", "WHERE", "LIMIT", "ORDER", "BY", "ASC", "DESC",
	"WITH", "TYPEOF", "REFERENCE", "VIEW", "VIEWSTAT", "TRACKING", "OFFSET",
	"IN", "END", "USING", "DATA", "CATEGORY", "GROUP", "HAVING", "NULLS", "FIRST",
	"LAST", "SCOPE", "ROLLUP", "CUBE", "LIKE", "THEN", "SOQL_AND", "SOQL_OR",
	"SOQL_NOT", "FIND", "FIELDS", "RETURNING", "ALL",
}

var ruleNames = []string{
	"query", "selectClause", "fieldList", "selectField", "fromClause", "filterScope",
	"soqlField", "subquery", "whereClause", "whereFields", "whereField", "limitClause",
	"orderClause", "soqlValue", "withClause", "soqlFilteringExpression", "groupClause",
	"havingConditionExpression", "offsetClause", "viewClause", "soslLiteral",
	"soslQuery", "soslReturningObject", "apexIdentifier", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type soqlParser struct {
	*antlr.BaseParser
}

func NewsoqlParser(input antlr.TokenStream) *soqlParser {
	this := new(soqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "soql.g4"

	return this
}

// soqlParser tokens.
const (
	soqlParserEOF                  = antlr.TokenEOF
	soqlParserT__0                 = 1
	soqlParserT__1                 = 2
	soqlParserT__2                 = 3
	soqlParserT__3                 = 4
	soqlParserT__4                 = 5
	soqlParserT__5                 = 6
	soqlParserT__6                 = 7
	soqlParserT__7                 = 8
	soqlParserT__8                 = 9
	soqlParserT__9                 = 10
	soqlParserStringLiteral        = 11
	soqlParserBooleanLiteral       = 12
	soqlParserNullLiteral          = 13
	soqlParserIdentifier           = 14
	soqlParserIntegerLiteral       = 15
	soqlParserFloatingPointLiteral = 16
	soqlParserQUOTE                = 17
	soqlParserLPAREN               = 18
	soqlParserRPAREN               = 19
	soqlParserWHEN                 = 20
	soqlParserELSE                 = 21
	soqlParserCOMMA                = 22
	soqlParserDOT                  = 23
	soqlParserFOR                  = 24
	soqlParserUPDATE               = 25
	soqlParserCOLON                = 26
	soqlParserSELECT               = 27
	soqlParserFROM                 = 28
	soqlParserWHERE                = 29
	soqlParserLIMIT                = 30
	soqlParserORDER                = 31
	soqlParserBY                   = 32
	soqlParserASC                  = 33
	soqlParserDESC                 = 34
	soqlParserWITH                 = 35
	soqlParserTYPEOF               = 36
	soqlParserREFERENCE            = 37
	soqlParserVIEW                 = 38
	soqlParserVIEWSTAT             = 39
	soqlParserTRACKING             = 40
	soqlParserOFFSET               = 41
	soqlParserIN                   = 42
	soqlParserEND                  = 43
	soqlParserUSING                = 44
	soqlParserDATA                 = 45
	soqlParserCATEGORY             = 46
	soqlParserGROUP                = 47
	soqlParserHAVING               = 48
	soqlParserNULLS                = 49
	soqlParserFIRST                = 50
	soqlParserLAST                 = 51
	soqlParserSCOPE                = 52
	soqlParserROLLUP               = 53
	soqlParserCUBE                 = 54
	soqlParserLIKE                 = 55
	soqlParserTHEN                 = 56
	soqlParserSOQL_AND             = 57
	soqlParserSOQL_OR              = 58
	soqlParserSOQL_NOT             = 59
	soqlParserFIND                 = 60
	soqlParserFIELDS               = 61
	soqlParserRETURNING            = 62
	soqlParserALL                  = 63
)

// soqlParser rules.
const (
	soqlParserRULE_query                     = 0
	soqlParserRULE_selectClause              = 1
	soqlParserRULE_fieldList                 = 2
	soqlParserRULE_selectField               = 3
	soqlParserRULE_fromClause                = 4
	soqlParserRULE_filterScope               = 5
	soqlParserRULE_soqlField                 = 6
	soqlParserRULE_subquery                  = 7
	soqlParserRULE_whereClause               = 8
	soqlParserRULE_whereFields               = 9
	soqlParserRULE_whereField                = 10
	soqlParserRULE_limitClause               = 11
	soqlParserRULE_orderClause               = 12
	soqlParserRULE_soqlValue                 = 13
	soqlParserRULE_withClause                = 14
	soqlParserRULE_soqlFilteringExpression   = 15
	soqlParserRULE_groupClause               = 16
	soqlParserRULE_havingConditionExpression = 17
	soqlParserRULE_offsetClause              = 18
	soqlParserRULE_viewClause                = 19
	soqlParserRULE_soslLiteral               = 20
	soqlParserRULE_soslQuery                 = 21
	soqlParserRULE_soslReturningObject       = 22
	soqlParserRULE_apexIdentifier            = 23
	soqlParserRULE_literal                   = 24
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) SelectClause() ISelectClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectClauseContext)
}

func (s *QueryContext) FromClause() IFromClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFromClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *QueryContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *QueryContext) WithClause() IWithClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithClauseContext)
}

func (s *QueryContext) GroupClause() IGroupClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupClauseContext)
}

func (s *QueryContext) OrderClause() IOrderClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderClauseContext)
}

func (s *QueryContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *QueryContext) OffsetClause() IOffsetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffsetClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOffsetClauseContext)
}

func (s *QueryContext) ViewClause() IViewClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IViewClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IViewClauseContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, soqlParserRULE_query)
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
		p.SetState(50)
		p.SelectClause()
	}
	{
		p.SetState(51)
		p.FromClause()
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserWHERE {
		{
			p.SetState(52)
			p.WhereClause()
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserWITH {
		{
			p.SetState(55)
			p.WithClause()
		}

	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserGROUP {
		{
			p.SetState(58)
			p.GroupClause()
		}

	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserORDER {
		{
			p.SetState(61)
			p.OrderClause()
		}

	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserLIMIT {
		{
			p.SetState(64)
			p.LimitClause()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserOFFSET {
		{
			p.SetState(67)
			p.OffsetClause()
		}

	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserFOR {
		{
			p.SetState(70)
			p.ViewClause()
		}

	}

	return localctx
}

// ISelectClauseContext is an interface to support dynamic dispatch.
type ISelectClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectClauseContext differentiates from other interfaces.
	IsSelectClauseContext()
}

type SelectClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectClauseContext() *SelectClauseContext {
	var p = new(SelectClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_selectClause
	return p
}

func (*SelectClauseContext) IsSelectClauseContext() {}

func NewSelectClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectClauseContext {
	var p = new(SelectClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_selectClause

	return p
}

func (s *SelectClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectClauseContext) SELECT() antlr.TerminalNode {
	return s.GetToken(soqlParserSELECT, 0)
}

func (s *SelectClauseContext) FieldList() IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *SelectClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSelectClause(s)
	}
}

func (s *SelectClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSelectClause(s)
	}
}

func (s *SelectClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSelectClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SelectClause() (localctx ISelectClauseContext) {
	localctx = NewSelectClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, soqlParserRULE_selectClause)

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
		p.SetState(73)
		p.Match(soqlParserSELECT)
	}
	{
		p.SetState(74)
		p.FieldList()
	}

	return localctx
}

// IFieldListContext is an interface to support dynamic dispatch.
type IFieldListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldListContext differentiates from other interfaces.
	IsFieldListContext()
}

type FieldListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldListContext() *FieldListContext {
	var p = new(FieldListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_fieldList
	return p
}

func (*FieldListContext) IsFieldListContext() {}

func NewFieldListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldListContext {
	var p = new(FieldListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_fieldList

	return p
}

func (s *FieldListContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldListContext) AllSelectField() []ISelectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectFieldContext)(nil)).Elem())
	var tst = make([]ISelectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectFieldContext)
		}
	}

	return tst
}

func (s *FieldListContext) SelectField(i int) ISelectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectFieldContext)
}

func (s *FieldListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(soqlParserCOMMA)
}

func (s *FieldListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserCOMMA, i)
}

func (s *FieldListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterFieldList(s)
	}
}

func (s *FieldListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitFieldList(s)
	}
}

func (s *FieldListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitFieldList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) FieldList() (localctx IFieldListContext) {
	localctx = NewFieldListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, soqlParserRULE_fieldList)
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
		p.SetState(76)
		p.SelectField()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == soqlParserCOMMA {
		{
			p.SetState(77)
			p.Match(soqlParserCOMMA)
		}
		{
			p.SetState(78)
			p.SelectField()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectFieldContext is an interface to support dynamic dispatch.
type ISelectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectFieldContext differentiates from other interfaces.
	IsSelectFieldContext()
}

type SelectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectFieldContext() *SelectFieldContext {
	var p = new(SelectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_selectField
	return p
}

func (*SelectFieldContext) IsSelectFieldContext() {}

func NewSelectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectFieldContext {
	var p = new(SelectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_selectField

	return p
}

func (s *SelectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectFieldContext) SoqlField() ISoqlFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoqlFieldContext)
}

func (s *SelectFieldContext) Subquery() ISubqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
}

func (s *SelectFieldContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(soqlParserTYPEOF, 0)
}

func (s *SelectFieldContext) ELSE() antlr.TerminalNode {
	return s.GetToken(soqlParserELSE, 0)
}

func (s *SelectFieldContext) AllFieldList() []IFieldListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldListContext)(nil)).Elem())
	var tst = make([]IFieldListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldListContext)
		}
	}

	return tst
}

func (s *SelectFieldContext) FieldList(i int) IFieldListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldListContext)
}

func (s *SelectFieldContext) END() antlr.TerminalNode {
	return s.GetToken(soqlParserEND, 0)
}

func (s *SelectFieldContext) AllWHEN() []antlr.TerminalNode {
	return s.GetTokens(soqlParserWHEN)
}

func (s *SelectFieldContext) WHEN(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserWHEN, i)
}

func (s *SelectFieldContext) AllApexIdentifier() []IApexIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem())
	var tst = make([]IApexIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IApexIdentifierContext)
		}
	}

	return tst
}

func (s *SelectFieldContext) ApexIdentifier(i int) IApexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IApexIdentifierContext)
}

func (s *SelectFieldContext) AllTHEN() []antlr.TerminalNode {
	return s.GetTokens(soqlParserTHEN)
}

func (s *SelectFieldContext) THEN(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserTHEN, i)
}

func (s *SelectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSelectField(s)
	}
}

func (s *SelectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSelectField(s)
	}
}

func (s *SelectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSelectField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SelectField() (localctx ISelectFieldContext) {
	localctx = NewSelectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, soqlParserRULE_selectField)
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

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case soqlParserIdentifier, soqlParserREFERENCE, soqlParserOFFSET, soqlParserDATA, soqlParserCATEGORY, soqlParserGROUP, soqlParserSCOPE, soqlParserTHEN, soqlParserFIND, soqlParserRETURNING, soqlParserALL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.SoqlField()
		}

	case soqlParserSELECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Subquery()
		}

	case soqlParserTYPEOF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Match(soqlParserTYPEOF)
		}
		{
			p.SetState(87)
			p.SoqlField()
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == soqlParserWHEN {
			{
				p.SetState(88)
				p.Match(soqlParserWHEN)
			}
			{
				p.SetState(89)
				p.ApexIdentifier()
			}
			{
				p.SetState(90)
				p.Match(soqlParserTHEN)
			}
			{
				p.SetState(91)
				p.FieldList()
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(soqlParserELSE)
		}
		{
			p.SetState(98)
			p.FieldList()
		}
		{
			p.SetState(99)
			p.Match(soqlParserEND)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(soqlParserFROM, 0)
}

func (s *FromClauseContext) ApexIdentifier() IApexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IApexIdentifierContext)
}

func (s *FromClauseContext) USING() antlr.TerminalNode {
	return s.GetToken(soqlParserUSING, 0)
}

func (s *FromClauseContext) SCOPE() antlr.TerminalNode {
	return s.GetToken(soqlParserSCOPE, 0)
}

func (s *FromClauseContext) FilterScope() IFilterScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilterScopeContext)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitFromClause(s)
	}
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) FromClause() (localctx IFromClauseContext) {
	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, soqlParserRULE_fromClause)
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
		p.SetState(103)
		p.Match(soqlParserFROM)
	}
	{
		p.SetState(104)
		p.ApexIdentifier()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserUSING {
		{
			p.SetState(105)
			p.Match(soqlParserUSING)
		}
		{
			p.SetState(106)
			p.Match(soqlParserSCOPE)
		}
		{
			p.SetState(107)
			p.FilterScope()
		}

	}

	return localctx
}

// IFilterScopeContext is an interface to support dynamic dispatch.
type IFilterScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterScopeContext differentiates from other interfaces.
	IsFilterScopeContext()
}

type FilterScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterScopeContext() *FilterScopeContext {
	var p = new(FilterScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_filterScope
	return p
}

func (*FilterScopeContext) IsFilterScopeContext() {}

func NewFilterScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterScopeContext {
	var p = new(FilterScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_filterScope

	return p
}

func (s *FilterScopeContext) GetParser() antlr.Parser { return s.parser }
func (s *FilterScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterFilterScope(s)
	}
}

func (s *FilterScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitFilterScope(s)
	}
}

func (s *FilterScopeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitFilterScope(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) FilterScope() (localctx IFilterScopeContext) {
	localctx = NewFilterScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, soqlParserRULE_filterScope)

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

	return localctx
}

// ISoqlFieldContext is an interface to support dynamic dispatch.
type ISoqlFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoqlFieldContext differentiates from other interfaces.
	IsSoqlFieldContext()
}

type SoqlFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoqlFieldContext() *SoqlFieldContext {
	var p = new(SoqlFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soqlField
	return p
}

func (*SoqlFieldContext) IsSoqlFieldContext() {}

func NewSoqlFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoqlFieldContext {
	var p = new(SoqlFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soqlField

	return p
}

func (s *SoqlFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *SoqlFieldContext) CopyFrom(ctx *SoqlFieldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SoqlFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoqlFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SoqlFieldReferenceContext struct {
	*SoqlFieldContext
}

func NewSoqlFieldReferenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SoqlFieldReferenceContext {
	var p = new(SoqlFieldReferenceContext)

	p.SoqlFieldContext = NewEmptySoqlFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SoqlFieldContext))

	return p
}

func (s *SoqlFieldReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoqlFieldReferenceContext) AllApexIdentifier() []IApexIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem())
	var tst = make([]IApexIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IApexIdentifierContext)
		}
	}

	return tst
}

func (s *SoqlFieldReferenceContext) ApexIdentifier(i int) IApexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IApexIdentifierContext)
}

func (s *SoqlFieldReferenceContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(soqlParserDOT)
}

func (s *SoqlFieldReferenceContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserDOT, i)
}

func (s *SoqlFieldReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoqlFieldReference(s)
	}
}

func (s *SoqlFieldReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoqlFieldReference(s)
	}
}

func (s *SoqlFieldReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoqlFieldReference(s)

	default:
		return t.VisitChildren(s)
	}
}

type SoqlFunctionCallContext struct {
	*SoqlFieldContext
}

func NewSoqlFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SoqlFunctionCallContext {
	var p = new(SoqlFunctionCallContext)

	p.SoqlFieldContext = NewEmptySoqlFieldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SoqlFieldContext))

	return p
}

func (s *SoqlFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoqlFunctionCallContext) ApexIdentifier() IApexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IApexIdentifierContext)
}

func (s *SoqlFunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(soqlParserLPAREN, 0)
}

func (s *SoqlFunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(soqlParserRPAREN, 0)
}

func (s *SoqlFunctionCallContext) AllSoqlField() []ISoqlFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem())
	var tst = make([]ISoqlFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoqlFieldContext)
		}
	}

	return tst
}

func (s *SoqlFunctionCallContext) SoqlField(i int) ISoqlFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoqlFieldContext)
}

func (s *SoqlFunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(soqlParserCOMMA)
}

func (s *SoqlFunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserCOMMA, i)
}

func (s *SoqlFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoqlFunctionCall(s)
	}
}

func (s *SoqlFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoqlFunctionCall(s)
	}
}

func (s *SoqlFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoqlFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoqlField() (localctx ISoqlFieldContext) {
	localctx = NewSoqlFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, soqlParserRULE_soqlField)
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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSoqlFieldReferenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(112)
					p.ApexIdentifier()
				}
				{
					p.SetState(113)
					p.Match(soqlParserDOT)
				}

			}
			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}
		{
			p.SetState(120)
			p.ApexIdentifier()
		}

	case 2:
		localctx = NewSoqlFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.ApexIdentifier()
		}
		{
			p.SetState(122)
			p.Match(soqlParserLPAREN)
		}
		p.SetState(131)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == soqlParserIdentifier || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(soqlParserREFERENCE-37))|(1<<(soqlParserOFFSET-37))|(1<<(soqlParserDATA-37))|(1<<(soqlParserCATEGORY-37))|(1<<(soqlParserGROUP-37))|(1<<(soqlParserSCOPE-37))|(1<<(soqlParserTHEN-37))|(1<<(soqlParserFIND-37))|(1<<(soqlParserRETURNING-37))|(1<<(soqlParserALL-37)))) != 0) {
			{
				p.SetState(123)
				p.SoqlField()
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == soqlParserCOMMA {
				{
					p.SetState(124)
					p.Match(soqlParserCOMMA)
				}
				{
					p.SetState(125)
					p.SoqlField()
				}

				p.SetState(130)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(133)
			p.Match(soqlParserRPAREN)
		}

	}

	return localctx
}

// ISubqueryContext is an interface to support dynamic dispatch.
type ISubqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubqueryContext differentiates from other interfaces.
	IsSubqueryContext()
}

type SubqueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubqueryContext() *SubqueryContext {
	var p = new(SubqueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_subquery
	return p
}

func (*SubqueryContext) IsSubqueryContext() {}

func NewSubqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryContext {
	var p = new(SubqueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_subquery

	return p
}

func (s *SubqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *SubqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubqueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSubquery(s)
	}
}

func (s *SubqueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSubquery(s)
	}
}

func (s *SubqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSubquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) Subquery() (localctx ISubqueryContext) {
	localctx = NewSubqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, soqlParserRULE_subquery)

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
		p.SetState(137)
		p.Query()
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(soqlParserWHERE, 0)
}

func (s *WhereClauseContext) WhereFields() IWhereFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereFieldsContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, soqlParserRULE_whereClause)

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
		p.SetState(139)
		p.Match(soqlParserWHERE)
	}
	{
		p.SetState(140)
		p.whereFields(0)
	}

	return localctx
}

// IWhereFieldsContext is an interface to support dynamic dispatch.
type IWhereFieldsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAnd_or returns the and_or token.
	GetAnd_or() antlr.Token

	// SetAnd_or sets the and_or token.
	SetAnd_or(antlr.Token)

	// IsWhereFieldsContext differentiates from other interfaces.
	IsWhereFieldsContext()
}

type WhereFieldsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	and_or antlr.Token
}

func NewEmptyWhereFieldsContext() *WhereFieldsContext {
	var p = new(WhereFieldsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_whereFields
	return p
}

func (*WhereFieldsContext) IsWhereFieldsContext() {}

func NewWhereFieldsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereFieldsContext {
	var p = new(WhereFieldsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_whereFields

	return p
}

func (s *WhereFieldsContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereFieldsContext) GetAnd_or() antlr.Token { return s.and_or }

func (s *WhereFieldsContext) SetAnd_or(v antlr.Token) { s.and_or = v }

func (s *WhereFieldsContext) WhereField() IWhereFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereFieldContext)
}

func (s *WhereFieldsContext) AllWhereFields() []IWhereFieldsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereFieldsContext)(nil)).Elem())
	var tst = make([]IWhereFieldsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereFieldsContext)
		}
	}

	return tst
}

func (s *WhereFieldsContext) WhereFields(i int) IWhereFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereFieldsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereFieldsContext)
}

func (s *WhereFieldsContext) SOQL_AND() antlr.TerminalNode {
	return s.GetToken(soqlParserSOQL_AND, 0)
}

func (s *WhereFieldsContext) SOQL_OR() antlr.TerminalNode {
	return s.GetToken(soqlParserSOQL_OR, 0)
}

func (s *WhereFieldsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereFieldsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereFieldsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterWhereFields(s)
	}
}

func (s *WhereFieldsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitWhereFields(s)
	}
}

func (s *WhereFieldsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitWhereFields(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) WhereFields() (localctx IWhereFieldsContext) {
	return p.whereFields(0)
}

func (p *soqlParser) whereFields(_p int) (localctx IWhereFieldsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewWhereFieldsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IWhereFieldsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, soqlParserRULE_whereFields, _p)
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
	{
		p.SetState(143)
		p.WhereField()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewWhereFieldsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, soqlParserRULE_whereFields)
			p.SetState(145)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(146)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*WhereFieldsContext).and_or = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == soqlParserSOQL_AND || _la == soqlParserSOQL_OR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*WhereFieldsContext).and_or = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(147)
				p.whereFields(2)
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IWhereFieldContext is an interface to support dynamic dispatch.
type IWhereFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// IsWhereFieldContext differentiates from other interfaces.
	IsWhereFieldContext()
}

type WhereFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
}

func NewEmptyWhereFieldContext() *WhereFieldContext {
	var p = new(WhereFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_whereField
	return p
}

func (*WhereFieldContext) IsWhereFieldContext() {}

func NewWhereFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereFieldContext {
	var p = new(WhereFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_whereField

	return p
}

func (s *WhereFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereFieldContext) GetOp() antlr.Token { return s.op }

func (s *WhereFieldContext) SetOp(v antlr.Token) { s.op = v }

func (s *WhereFieldContext) SoqlField() ISoqlFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoqlFieldContext)
}

func (s *WhereFieldContext) SoqlValue() ISoqlValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoqlValueContext)
}

func (s *WhereFieldContext) LIKE() antlr.TerminalNode {
	return s.GetToken(soqlParserLIKE, 0)
}

func (s *WhereFieldContext) IN() antlr.TerminalNode {
	return s.GetToken(soqlParserIN, 0)
}

func (s *WhereFieldContext) SOQL_NOT() antlr.TerminalNode {
	return s.GetToken(soqlParserSOQL_NOT, 0)
}

func (s *WhereFieldContext) WhereFields() IWhereFieldsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereFieldsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereFieldsContext)
}

func (s *WhereFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterWhereField(s)
	}
}

func (s *WhereFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitWhereField(s)
	}
}

func (s *WhereFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitWhereField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) WhereField() (localctx IWhereFieldContext) {
	localctx = NewWhereFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, soqlParserRULE_whereField)
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

	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case soqlParserIdentifier, soqlParserREFERENCE, soqlParserOFFSET, soqlParserDATA, soqlParserCATEGORY, soqlParserGROUP, soqlParserSCOPE, soqlParserTHEN, soqlParserSOQL_NOT, soqlParserFIND, soqlParserRETURNING, soqlParserALL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == soqlParserSOQL_NOT {
			{
				p.SetState(153)
				p.Match(soqlParserSOQL_NOT)
			}

		}
		{
			p.SetState(156)
			p.SoqlField()
		}
		{
			p.SetState(157)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*WhereFieldContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<soqlParserT__0)|(1<<soqlParserT__1)|(1<<soqlParserT__2)|(1<<soqlParserT__3)|(1<<soqlParserT__4)|(1<<soqlParserT__5)|(1<<soqlParserT__6))) != 0) || _la == soqlParserIN || _la == soqlParserLIKE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*WhereFieldContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(158)
			p.SoqlValue()
		}

	case soqlParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Match(soqlParserLPAREN)
		}
		{
			p.SetState(161)
			p.whereFields(0)
		}
		{
			p.SetState(162)
			p.Match(soqlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(soqlParserLIMIT, 0)
}

func (s *LimitClauseContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserIntegerLiteral, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, soqlParserRULE_limitClause)

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
		p.SetState(166)
		p.Match(soqlParserLIMIT)
	}
	{
		p.SetState(167)
		p.Match(soqlParserIntegerLiteral)
	}

	return localctx
}

// IOrderClauseContext is an interface to support dynamic dispatch.
type IOrderClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsc_desc returns the asc_desc token.
	GetAsc_desc() antlr.Token

	// GetNulls returns the nulls token.
	GetNulls() antlr.Token

	// SetAsc_desc sets the asc_desc token.
	SetAsc_desc(antlr.Token)

	// SetNulls sets the nulls token.
	SetNulls(antlr.Token)

	// IsOrderClauseContext differentiates from other interfaces.
	IsOrderClauseContext()
}

type OrderClauseContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	asc_desc antlr.Token
	nulls    antlr.Token
}

func NewEmptyOrderClauseContext() *OrderClauseContext {
	var p = new(OrderClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_orderClause
	return p
}

func (*OrderClauseContext) IsOrderClauseContext() {}

func NewOrderClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderClauseContext {
	var p = new(OrderClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_orderClause

	return p
}

func (s *OrderClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderClauseContext) GetAsc_desc() antlr.Token { return s.asc_desc }

func (s *OrderClauseContext) GetNulls() antlr.Token { return s.nulls }

func (s *OrderClauseContext) SetAsc_desc(v antlr.Token) { s.asc_desc = v }

func (s *OrderClauseContext) SetNulls(v antlr.Token) { s.nulls = v }

func (s *OrderClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(soqlParserORDER, 0)
}

func (s *OrderClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(soqlParserBY, 0)
}

func (s *OrderClauseContext) AllSoqlField() []ISoqlFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem())
	var tst = make([]ISoqlFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoqlFieldContext)
		}
	}

	return tst
}

func (s *OrderClauseContext) SoqlField(i int) ISoqlFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoqlFieldContext)
}

func (s *OrderClauseContext) NULLS() antlr.TerminalNode {
	return s.GetToken(soqlParserNULLS, 0)
}

func (s *OrderClauseContext) ASC() antlr.TerminalNode {
	return s.GetToken(soqlParserASC, 0)
}

func (s *OrderClauseContext) DESC() antlr.TerminalNode {
	return s.GetToken(soqlParserDESC, 0)
}

func (s *OrderClauseContext) LAST() antlr.TerminalNode {
	return s.GetToken(soqlParserLAST, 0)
}

func (s *OrderClauseContext) FIRST() antlr.TerminalNode {
	return s.GetToken(soqlParserFIRST, 0)
}

func (s *OrderClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterOrderClause(s)
	}
}

func (s *OrderClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitOrderClause(s)
	}
}

func (s *OrderClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitOrderClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) OrderClause() (localctx IOrderClauseContext) {
	localctx = NewOrderClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, soqlParserRULE_orderClause)
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
		p.SetState(169)
		p.Match(soqlParserORDER)
	}
	{
		p.SetState(170)
		p.Match(soqlParserBY)
	}
	{
		p.SetState(171)
		p.SoqlField()
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == soqlParserT__7 {
		{
			p.SetState(172)
			p.Match(soqlParserT__7)
		}
		{
			p.SetState(173)
			p.SoqlField()
		}

		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserASC || _la == soqlParserDESC {
		{
			p.SetState(179)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderClauseContext).asc_desc = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == soqlParserASC || _la == soqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderClauseContext).asc_desc = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserNULLS {
		{
			p.SetState(182)
			p.Match(soqlParserNULLS)
		}
		{
			p.SetState(183)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderClauseContext).nulls = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == soqlParserFIRST || _la == soqlParserLAST) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderClauseContext).nulls = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ISoqlValueContext is an interface to support dynamic dispatch.
type ISoqlValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoqlValueContext differentiates from other interfaces.
	IsSoqlValueContext()
}

type SoqlValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoqlValueContext() *SoqlValueContext {
	var p = new(SoqlValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soqlValue
	return p
}

func (*SoqlValueContext) IsSoqlValueContext() {}

func NewSoqlValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoqlValueContext {
	var p = new(SoqlValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soqlValue

	return p
}

func (s *SoqlValueContext) GetParser() antlr.Parser { return s.parser }

func (s *SoqlValueContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SoqlValueContext) ApexIdentifier() IApexIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IApexIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IApexIdentifierContext)
}

func (s *SoqlValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(soqlParserCOLON, 0)
}

func (s *SoqlValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoqlValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoqlValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoqlValue(s)
	}
}

func (s *SoqlValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoqlValue(s)
	}
}

func (s *SoqlValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoqlValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoqlValue() (localctx ISoqlValueContext) {
	localctx = NewSoqlValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, soqlParserRULE_soqlValue)

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

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case soqlParserStringLiteral, soqlParserBooleanLiteral, soqlParserNullLiteral, soqlParserIntegerLiteral, soqlParserFloatingPointLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Literal()
		}

	case soqlParserIdentifier, soqlParserREFERENCE, soqlParserOFFSET, soqlParserDATA, soqlParserCATEGORY, soqlParserGROUP, soqlParserSCOPE, soqlParserTHEN, soqlParserFIND, soqlParserRETURNING, soqlParserALL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			p.ApexIdentifier()
		}
		{
			p.SetState(188)
			p.Match(soqlParserCOLON)
		}
		{
			p.SetState(189)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWithClauseContext is an interface to support dynamic dispatch.
type IWithClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithClauseContext differentiates from other interfaces.
	IsWithClauseContext()
}

type WithClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithClauseContext() *WithClauseContext {
	var p = new(WithClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_withClause
	return p
}

func (*WithClauseContext) IsWithClauseContext() {}

func NewWithClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithClauseContext {
	var p = new(WithClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_withClause

	return p
}

func (s *WithClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WithClauseContext) WITH() antlr.TerminalNode {
	return s.GetToken(soqlParserWITH, 0)
}

func (s *WithClauseContext) DATA() antlr.TerminalNode {
	return s.GetToken(soqlParserDATA, 0)
}

func (s *WithClauseContext) CATEGORY() antlr.TerminalNode {
	return s.GetToken(soqlParserCATEGORY, 0)
}

func (s *WithClauseContext) SoqlFilteringExpression() ISoqlFilteringExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFilteringExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoqlFilteringExpressionContext)
}

func (s *WithClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterWithClause(s)
	}
}

func (s *WithClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitWithClause(s)
	}
}

func (s *WithClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitWithClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) WithClause() (localctx IWithClauseContext) {
	localctx = NewWithClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, soqlParserRULE_withClause)

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
		p.SetState(193)
		p.Match(soqlParserWITH)
	}
	{
		p.SetState(194)
		p.Match(soqlParserDATA)
	}
	{
		p.SetState(195)
		p.Match(soqlParserCATEGORY)
	}
	{
		p.SetState(196)
		p.SoqlFilteringExpression()
	}

	return localctx
}

// ISoqlFilteringExpressionContext is an interface to support dynamic dispatch.
type ISoqlFilteringExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoqlFilteringExpressionContext differentiates from other interfaces.
	IsSoqlFilteringExpressionContext()
}

type SoqlFilteringExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoqlFilteringExpressionContext() *SoqlFilteringExpressionContext {
	var p = new(SoqlFilteringExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soqlFilteringExpression
	return p
}

func (*SoqlFilteringExpressionContext) IsSoqlFilteringExpressionContext() {}

func NewSoqlFilteringExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoqlFilteringExpressionContext {
	var p = new(SoqlFilteringExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soqlFilteringExpression

	return p
}

func (s *SoqlFilteringExpressionContext) GetParser() antlr.Parser { return s.parser }
func (s *SoqlFilteringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoqlFilteringExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoqlFilteringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoqlFilteringExpression(s)
	}
}

func (s *SoqlFilteringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoqlFilteringExpression(s)
	}
}

func (s *SoqlFilteringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoqlFilteringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoqlFilteringExpression() (localctx ISoqlFilteringExpressionContext) {
	localctx = NewSoqlFilteringExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, soqlParserRULE_soqlFilteringExpression)

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

	return localctx
}

// IGroupClauseContext is an interface to support dynamic dispatch.
type IGroupClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupClauseContext differentiates from other interfaces.
	IsGroupClauseContext()
}

type GroupClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupClauseContext() *GroupClauseContext {
	var p = new(GroupClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_groupClause
	return p
}

func (*GroupClauseContext) IsGroupClauseContext() {}

func NewGroupClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupClauseContext {
	var p = new(GroupClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_groupClause

	return p
}

func (s *GroupClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(soqlParserGROUP, 0)
}

func (s *GroupClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(soqlParserBY, 0)
}

func (s *GroupClauseContext) AllSoqlField() []ISoqlFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem())
	var tst = make([]ISoqlFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoqlFieldContext)
		}
	}

	return tst
}

func (s *GroupClauseContext) SoqlField(i int) ISoqlFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoqlFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoqlFieldContext)
}

func (s *GroupClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(soqlParserHAVING, 0)
}

func (s *GroupClauseContext) HavingConditionExpression() IHavingConditionExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHavingConditionExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHavingConditionExpressionContext)
}

func (s *GroupClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterGroupClause(s)
	}
}

func (s *GroupClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitGroupClause(s)
	}
}

func (s *GroupClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitGroupClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) GroupClause() (localctx IGroupClauseContext) {
	localctx = NewGroupClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, soqlParserRULE_groupClause)
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
		p.SetState(200)
		p.Match(soqlParserGROUP)
	}
	{
		p.SetState(201)
		p.Match(soqlParserBY)
	}
	{
		p.SetState(202)
		p.SoqlField()
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == soqlParserT__7 {
		{
			p.SetState(203)
			p.Match(soqlParserT__7)
		}
		{
			p.SetState(204)
			p.SoqlField()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserHAVING {
		{
			p.SetState(210)
			p.Match(soqlParserHAVING)
		}
		{
			p.SetState(211)
			p.HavingConditionExpression()
		}

	}

	return localctx
}

// IHavingConditionExpressionContext is an interface to support dynamic dispatch.
type IHavingConditionExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHavingConditionExpressionContext differentiates from other interfaces.
	IsHavingConditionExpressionContext()
}

type HavingConditionExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingConditionExpressionContext() *HavingConditionExpressionContext {
	var p = new(HavingConditionExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_havingConditionExpression
	return p
}

func (*HavingConditionExpressionContext) IsHavingConditionExpressionContext() {}

func NewHavingConditionExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingConditionExpressionContext {
	var p = new(HavingConditionExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_havingConditionExpression

	return p
}

func (s *HavingConditionExpressionContext) GetParser() antlr.Parser { return s.parser }
func (s *HavingConditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingConditionExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingConditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterHavingConditionExpression(s)
	}
}

func (s *HavingConditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitHavingConditionExpression(s)
	}
}

func (s *HavingConditionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitHavingConditionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) HavingConditionExpression() (localctx IHavingConditionExpressionContext) {
	localctx = NewHavingConditionExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, soqlParserRULE_havingConditionExpression)

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

	return localctx
}

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_offsetClause
	return p
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_offsetClause

	return p
}

func (s *OffsetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(soqlParserOFFSET, 0)
}

func (s *OffsetClauseContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserIntegerLiteral, 0)
}

func (s *OffsetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterOffsetClause(s)
	}
}

func (s *OffsetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitOffsetClause(s)
	}
}

func (s *OffsetClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitOffsetClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) OffsetClause() (localctx IOffsetClauseContext) {
	localctx = NewOffsetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, soqlParserRULE_offsetClause)

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
		p.SetState(216)
		p.Match(soqlParserOFFSET)
	}
	{
		p.SetState(217)
		p.Match(soqlParserIntegerLiteral)
	}

	return localctx
}

// IViewClauseContext is an interface to support dynamic dispatch.
type IViewClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsViewClauseContext differentiates from other interfaces.
	IsViewClauseContext()
}

type ViewClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyViewClauseContext() *ViewClauseContext {
	var p = new(ViewClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_viewClause
	return p
}

func (*ViewClauseContext) IsViewClauseContext() {}

func NewViewClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ViewClauseContext {
	var p = new(ViewClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_viewClause

	return p
}

func (s *ViewClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ViewClauseContext) FOR() antlr.TerminalNode {
	return s.GetToken(soqlParserFOR, 0)
}

func (s *ViewClauseContext) VIEW() antlr.TerminalNode {
	return s.GetToken(soqlParserVIEW, 0)
}

func (s *ViewClauseContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(soqlParserREFERENCE, 0)
}

func (s *ViewClauseContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(soqlParserUPDATE, 0)
}

func (s *ViewClauseContext) TRACKING() antlr.TerminalNode {
	return s.GetToken(soqlParserTRACKING, 0)
}

func (s *ViewClauseContext) VIEWSTAT() antlr.TerminalNode {
	return s.GetToken(soqlParserVIEWSTAT, 0)
}

func (s *ViewClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ViewClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ViewClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterViewClause(s)
	}
}

func (s *ViewClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitViewClause(s)
	}
}

func (s *ViewClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitViewClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) ViewClause() (localctx IViewClauseContext) {
	localctx = NewViewClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, soqlParserRULE_viewClause)
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
		p.Match(soqlParserFOR)
	}
	{
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !(_la == soqlParserREFERENCE || _la == soqlParserVIEW) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserUPDATE {
		{
			p.SetState(221)
			p.Match(soqlParserUPDATE)
		}
		{
			p.SetState(222)
			_la = p.GetTokenStream().LA(1)

			if !(_la == soqlParserVIEWSTAT || _la == soqlParserTRACKING) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ISoslLiteralContext is an interface to support dynamic dispatch.
type ISoslLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoslLiteralContext differentiates from other interfaces.
	IsSoslLiteralContext()
}

type SoslLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoslLiteralContext() *SoslLiteralContext {
	var p = new(SoslLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soslLiteral
	return p
}

func (*SoslLiteralContext) IsSoslLiteralContext() {}

func NewSoslLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoslLiteralContext {
	var p = new(SoslLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soslLiteral

	return p
}

func (s *SoslLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *SoslLiteralContext) SoslQuery() ISoslQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoslQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoslQueryContext)
}

func (s *SoslLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoslLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoslLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoslLiteral(s)
	}
}

func (s *SoslLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoslLiteral(s)
	}
}

func (s *SoslLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoslLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoslLiteral() (localctx ISoslLiteralContext) {
	localctx = NewSoslLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, soqlParserRULE_soslLiteral)

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
		p.SetState(225)
		p.Match(soqlParserT__8)
	}
	{
		p.SetState(226)
		p.SoslQuery()
	}
	{
		p.SetState(227)
		p.Match(soqlParserT__9)
	}

	return localctx
}

// ISoslQueryContext is an interface to support dynamic dispatch.
type ISoslQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoslQueryContext differentiates from other interfaces.
	IsSoslQueryContext()
}

type SoslQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoslQueryContext() *SoslQueryContext {
	var p = new(SoslQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soslQuery
	return p
}

func (*SoslQueryContext) IsSoslQueryContext() {}

func NewSoslQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoslQueryContext {
	var p = new(SoslQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soslQuery

	return p
}

func (s *SoslQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SoslQueryContext) FIND() antlr.TerminalNode {
	return s.GetToken(soqlParserFIND, 0)
}

func (s *SoslQueryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *SoslQueryContext) IN() antlr.TerminalNode {
	return s.GetToken(soqlParserIN, 0)
}

func (s *SoslQueryContext) ALL() antlr.TerminalNode {
	return s.GetToken(soqlParserALL, 0)
}

func (s *SoslQueryContext) FIELDS() antlr.TerminalNode {
	return s.GetToken(soqlParserFIELDS, 0)
}

func (s *SoslQueryContext) RETURNING() antlr.TerminalNode {
	return s.GetToken(soqlParserRETURNING, 0)
}

func (s *SoslQueryContext) AllSoslReturningObject() []ISoslReturningObjectContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISoslReturningObjectContext)(nil)).Elem())
	var tst = make([]ISoslReturningObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoslReturningObjectContext)
		}
	}

	return tst
}

func (s *SoslQueryContext) SoslReturningObject(i int) ISoslReturningObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoslReturningObjectContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoslReturningObjectContext)
}

func (s *SoslQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoslQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoslQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoslQuery(s)
	}
}

func (s *SoslQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoslQuery(s)
	}
}

func (s *SoslQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoslQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoslQuery() (localctx ISoslQueryContext) {
	localctx = NewSoslQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, soqlParserRULE_soslQuery)
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
		p.SetState(229)
		p.Match(soqlParserFIND)
	}
	{
		p.SetState(230)
		p.Literal()
	}
	{
		p.SetState(231)
		p.Match(soqlParserIN)
	}
	{
		p.SetState(232)
		p.Match(soqlParserALL)
	}
	{
		p.SetState(233)
		p.Match(soqlParserFIELDS)
	}
	{
		p.SetState(234)
		p.Match(soqlParserRETURNING)
	}
	{
		p.SetState(235)
		p.SoslReturningObject()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == soqlParserT__7 {
		{
			p.SetState(236)
			p.Match(soqlParserT__7)
		}
		{
			p.SetState(237)
			p.SoslReturningObject()
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISoslReturningObjectContext is an interface to support dynamic dispatch.
type ISoslReturningObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoslReturningObjectContext differentiates from other interfaces.
	IsSoslReturningObjectContext()
}

type SoslReturningObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoslReturningObjectContext() *SoslReturningObjectContext {
	var p = new(SoslReturningObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_soslReturningObject
	return p
}

func (*SoslReturningObjectContext) IsSoslReturningObjectContext() {}

func NewSoslReturningObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoslReturningObjectContext {
	var p = new(SoslReturningObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_soslReturningObject

	return p
}

func (s *SoslReturningObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SoslReturningObjectContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(soqlParserIdentifier)
}

func (s *SoslReturningObjectContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(soqlParserIdentifier, i)
}

func (s *SoslReturningObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoslReturningObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SoslReturningObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterSoslReturningObject(s)
	}
}

func (s *SoslReturningObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitSoslReturningObject(s)
	}
}

func (s *SoslReturningObjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitSoslReturningObject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) SoslReturningObject() (localctx ISoslReturningObjectContext) {
	localctx = NewSoslReturningObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, soqlParserRULE_soslReturningObject)
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
		p.SetState(243)
		p.Match(soqlParserIdentifier)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == soqlParserLPAREN {
		{
			p.SetState(244)
			p.Match(soqlParserLPAREN)
		}
		{
			p.SetState(245)
			p.Match(soqlParserIdentifier)
		}
		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == soqlParserT__7 {
			{
				p.SetState(246)
				p.Match(soqlParserT__7)
			}
			{
				p.SetState(247)
				p.Match(soqlParserIdentifier)
			}

			p.SetState(252)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(253)
			p.Match(soqlParserRPAREN)
		}

	}

	return localctx
}

// IApexIdentifierContext is an interface to support dynamic dispatch.
type IApexIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsApexIdentifierContext differentiates from other interfaces.
	IsApexIdentifierContext()
}

type ApexIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApexIdentifierContext() *ApexIdentifierContext {
	var p = new(ApexIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_apexIdentifier
	return p
}

func (*ApexIdentifierContext) IsApexIdentifierContext() {}

func NewApexIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApexIdentifierContext {
	var p = new(ApexIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_apexIdentifier

	return p
}

func (s *ApexIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ApexIdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(soqlParserIdentifier, 0)
}

func (s *ApexIdentifierContext) DATA() antlr.TerminalNode {
	return s.GetToken(soqlParserDATA, 0)
}

func (s *ApexIdentifierContext) GROUP() antlr.TerminalNode {
	return s.GetToken(soqlParserGROUP, 0)
}

func (s *ApexIdentifierContext) SCOPE() antlr.TerminalNode {
	return s.GetToken(soqlParserSCOPE, 0)
}

func (s *ApexIdentifierContext) CATEGORY() antlr.TerminalNode {
	return s.GetToken(soqlParserCATEGORY, 0)
}

func (s *ApexIdentifierContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(soqlParserREFERENCE, 0)
}

func (s *ApexIdentifierContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(soqlParserOFFSET, 0)
}

func (s *ApexIdentifierContext) THEN() antlr.TerminalNode {
	return s.GetToken(soqlParserTHEN, 0)
}

func (s *ApexIdentifierContext) FIND() antlr.TerminalNode {
	return s.GetToken(soqlParserFIND, 0)
}

func (s *ApexIdentifierContext) RETURNING() antlr.TerminalNode {
	return s.GetToken(soqlParserRETURNING, 0)
}

func (s *ApexIdentifierContext) ALL() antlr.TerminalNode {
	return s.GetToken(soqlParserALL, 0)
}

func (s *ApexIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApexIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApexIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterApexIdentifier(s)
	}
}

func (s *ApexIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitApexIdentifier(s)
	}
}

func (s *ApexIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitApexIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) ApexIdentifier() (localctx IApexIdentifierContext) {
	localctx = NewApexIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, soqlParserRULE_apexIdentifier)
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
		p.SetState(256)
		_la = p.GetTokenStream().LA(1)

		if !(_la == soqlParserIdentifier || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(soqlParserREFERENCE-37))|(1<<(soqlParserOFFSET-37))|(1<<(soqlParserDATA-37))|(1<<(soqlParserCATEGORY-37))|(1<<(soqlParserGROUP-37))|(1<<(soqlParserSCOPE-37))|(1<<(soqlParserTHEN-37))|(1<<(soqlParserFIND-37))|(1<<(soqlParserRETURNING-37))|(1<<(soqlParserALL-37)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = soqlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = soqlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserIntegerLiteral, 0)
}

func (s *LiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserFloatingPointLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserStringLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserBooleanLiteral, 0)
}

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(soqlParserNullLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(soqlListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case soqlVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *soqlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, soqlParserRULE_literal)
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
		p.SetState(258)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<soqlParserStringLiteral)|(1<<soqlParserBooleanLiteral)|(1<<soqlParserNullLiteral)|(1<<soqlParserIntegerLiteral)|(1<<soqlParserFloatingPointLiteral))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *soqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *WhereFieldsContext = nil
		if localctx != nil {
			t = localctx.(*WhereFieldsContext)
		}
		return p.WhereFields_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *soqlParser) WhereFields_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

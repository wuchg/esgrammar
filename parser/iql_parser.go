// Code generated from parser/Iql.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Iql

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

type IqlParser struct {
	*antlr.BaseParser
}

var IqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func iqlParserInit() {
	staticData := &IqlParserStaticData
	staticData.LiteralNames = []string{
		"", "'['", "']'", "", "", "", "", "", "' '", "';'", "'.'", "'('", "')'",
		"','", "'='", "'*'", "'~'", "'!~'", "'<'", "'<='", "'>'", "'>='", "'!='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "FN_SUM", "FN_COUNT", "FN_STARTOFWEEK", "DATETIME", "NUMBER",
		"WHITESPACE", "SCOL", "DOT", "OPEN_PAR", "CLOSE_PAR", "COMMA", "EQ",
		"STAR", "CONTAINS", "NOT_CONTAINS", "LT", "LT_EQ", "GT", "GT_EQ", "NOT_EQ",
		"K_AFTER", "K_AND", "K_ASC", "K_BEFORE", "K_BY", "K_CHANGED", "K_DESC",
		"K_EMPTY", "K_IN", "K_IS", "K_NOT", "K_NULL", "K_ON", "K_OR", "K_ORDER",
		"K_TO", "K_WAS", "F_AFFECTED_VERSION", "F_APPROVALS", "F_ASSIGNEE",
		"F_ATTACHMENTS", "F_CATEGORY", "F_COMMENT", "F_COMPONENT", "F_CREATED",
		"F_CREATED_DATE", "F_CREATOR", "F_CUSTOM_FIELD", "F_CUSTOMER_REQUEST_TYPE",
		"F_DATE", "F_DESCRIPTION", "F_DUE", "F_DURATION", "F_ENVIRONMENT", "F_EPIC_LINK",
		"F_FILTER", "F_FIX_VERSION", "F_ISSUE", "F_ISSUE_KEY", "F_ISSUE_TYPE",
		"F_KEY", "F_LABEL", "F_LABELS", "F_LAST_VIEWED", "F_LEVEL", "F_NUMBER",
		"F_ORGANIZATION", "F_ORIGINAL_ESTIMATE", "F_PARENT", "F_PRIORITY", "F_PROJECT",
		"F_RANK", "F_REMAINING_ESTIMATE", "F_REPORTER", "F_REQUEST_CHANNEL_TYPE",
		"F_REQUEST_LAST_ACTIVITY_TIME", "F_RESOLUTION", "F_RESOLUTION_DATE",
		"F_RESOLVED", "F_SLA", "F_SPRINT", "F_STATUS", "F_SUMMARY", "F_TEXT",
		"F_TIME_SPENT", "F_TYPE", "F_UPDATED", "F_USER", "F_VERSION", "F_VOTER",
		"F_VOTES", "F_WATCHER", "F_WATCHERS", "F_WORK_RATIO", "IDENTIFIER",
		"STRING_LITERAL", "COMMENT", "LINE_COMMENT", "SPACES",
	}
	staticData.RuleNames = []string{
		"query", "iql_stmt", "expr", "ordering_term", "operator", "literal_value",
		"function", "function_name", "function_param", "literal_list", "keyword",
		"state_name", "field", "compare_dates", "dates",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 101, 177, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 3, 0,
		32, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 5, 2, 46, 8, 2, 10, 2, 12, 2, 49, 9, 2, 1, 2, 1, 2, 3, 2,
		53, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 2, 3, 2, 62, 8, 2,
		3, 2, 64, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 72, 8, 2, 10,
		2, 12, 2, 75, 9, 2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 81, 8, 3, 1, 3, 1, 3,
		1, 3, 3, 3, 86, 8, 3, 5, 3, 88, 8, 3, 10, 3, 12, 3, 91, 9, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 112, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 118, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 125, 8, 6, 10,
		6, 12, 6, 128, 9, 6, 3, 6, 130, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 3, 8, 138, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 144, 8, 9, 10, 9, 12,
		9, 147, 9, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 5, 9, 155, 8, 9, 10,
		9, 12, 9, 158, 9, 9, 1, 9, 1, 9, 3, 9, 162, 8, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 3, 13, 171, 8, 13, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 0, 1, 4, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 0, 5, 2, 0, 25, 25, 29, 29, 2, 0, 3, 3, 5, 5, 2, 0, 23, 29, 31,
		39, 1, 0, 40, 96, 3, 0, 23, 23, 26, 26, 35, 35, 198, 0, 31, 1, 0, 0, 0,
		2, 35, 1, 0, 0, 0, 4, 63, 1, 0, 0, 0, 6, 76, 1, 0, 0, 0, 8, 111, 1, 0,
		0, 0, 10, 117, 1, 0, 0, 0, 12, 119, 1, 0, 0, 0, 14, 133, 1, 0, 0, 0, 16,
		137, 1, 0, 0, 0, 18, 161, 1, 0, 0, 0, 20, 163, 1, 0, 0, 0, 22, 165, 1,
		0, 0, 0, 24, 167, 1, 0, 0, 0, 26, 170, 1, 0, 0, 0, 28, 174, 1, 0, 0, 0,
		30, 32, 3, 2, 1, 0, 31, 30, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32, 33, 1,
		0, 0, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0, 0, 0, 35, 37, 3, 4, 2, 0, 36,
		38, 3, 6, 3, 0, 37, 36, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 3, 1, 0, 0,
		0, 39, 40, 6, 2, -1, 0, 40, 41, 5, 11, 0, 0, 41, 42, 3, 4, 2, 0, 42, 43,
		5, 12, 0, 0, 43, 64, 1, 0, 0, 0, 44, 46, 5, 33, 0, 0, 45, 44, 1, 0, 0,
		0, 46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 52,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 53, 3, 24, 12, 0, 51, 53, 5, 98, 0,
		0, 52, 50, 1, 0, 0, 0, 52, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 58,
		3, 8, 4, 0, 55, 59, 3, 10, 5, 0, 56, 59, 3, 18, 9, 0, 57, 59, 3, 12, 6,
		0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 61,
		1, 0, 0, 0, 60, 62, 3, 6, 3, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 64, 1, 0, 0, 0, 63, 39, 1, 0, 0, 0, 63, 47, 1, 0, 0, 0, 64, 73, 1,
		0, 0, 0, 65, 66, 10, 4, 0, 0, 66, 67, 5, 24, 0, 0, 67, 72, 3, 4, 2, 5,
		68, 69, 10, 3, 0, 0, 69, 70, 5, 36, 0, 0, 70, 72, 3, 4, 2, 4, 71, 65, 1,
		0, 0, 0, 71, 68, 1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73,
		74, 1, 0, 0, 0, 74, 5, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 37, 0,
		0, 77, 78, 5, 27, 0, 0, 78, 80, 3, 10, 5, 0, 79, 81, 7, 0, 0, 0, 80, 79,
		1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 89, 1, 0, 0, 0, 82, 83, 5, 13, 0, 0,
		83, 85, 3, 10, 5, 0, 84, 86, 7, 0, 0, 0, 85, 84, 1, 0, 0, 0, 85, 86, 1,
		0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 82, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89,
		87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 7, 1, 0, 0, 0, 91, 89, 1, 0, 0,
		0, 92, 112, 5, 14, 0, 0, 93, 112, 5, 22, 0, 0, 94, 112, 5, 16, 0, 0, 95,
		112, 5, 17, 0, 0, 96, 112, 5, 19, 0, 0, 97, 112, 5, 18, 0, 0, 98, 112,
		5, 20, 0, 0, 99, 112, 5, 21, 0, 0, 100, 112, 5, 31, 0, 0, 101, 102, 5,
		33, 0, 0, 102, 112, 5, 31, 0, 0, 103, 112, 5, 32, 0, 0, 104, 112, 5, 39,
		0, 0, 105, 106, 5, 32, 0, 0, 106, 112, 5, 33, 0, 0, 107, 108, 5, 39, 0,
		0, 108, 112, 5, 33, 0, 0, 109, 110, 5, 28, 0, 0, 110, 112, 5, 38, 0, 0,
		111, 92, 1, 0, 0, 0, 111, 93, 1, 0, 0, 0, 111, 94, 1, 0, 0, 0, 111, 95,
		1, 0, 0, 0, 111, 96, 1, 0, 0, 0, 111, 97, 1, 0, 0, 0, 111, 98, 1, 0, 0,
		0, 111, 99, 1, 0, 0, 0, 111, 100, 1, 0, 0, 0, 111, 101, 1, 0, 0, 0, 111,
		103, 1, 0, 0, 0, 111, 104, 1, 0, 0, 0, 111, 105, 1, 0, 0, 0, 111, 107,
		1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112, 9, 1, 0, 0, 0, 113, 118, 5, 98,
		0, 0, 114, 118, 5, 97, 0, 0, 115, 118, 5, 7, 0, 0, 116, 118, 3, 28, 14,
		0, 117, 113, 1, 0, 0, 0, 117, 114, 1, 0, 0, 0, 117, 115, 1, 0, 0, 0, 117,
		116, 1, 0, 0, 0, 118, 11, 1, 0, 0, 0, 119, 120, 3, 14, 7, 0, 120, 129,
		5, 11, 0, 0, 121, 126, 3, 16, 8, 0, 122, 123, 5, 13, 0, 0, 123, 125, 3,
		16, 8, 0, 124, 122, 1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0,
		0, 126, 127, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129,
		121, 1, 0, 0, 0, 129, 130, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 132,
		5, 12, 0, 0, 132, 13, 1, 0, 0, 0, 133, 134, 7, 1, 0, 0, 134, 15, 1, 0,
		0, 0, 135, 138, 3, 10, 5, 0, 136, 138, 3, 12, 6, 0, 137, 135, 1, 0, 0,
		0, 137, 136, 1, 0, 0, 0, 138, 17, 1, 0, 0, 0, 139, 140, 5, 11, 0, 0, 140,
		145, 3, 10, 5, 0, 141, 142, 5, 13, 0, 0, 142, 144, 3, 10, 5, 0, 143, 141,
		1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 145, 146, 1, 0,
		0, 0, 146, 148, 1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 148, 149, 5, 12, 0, 0,
		149, 162, 1, 0, 0, 0, 150, 151, 5, 1, 0, 0, 151, 156, 3, 10, 5, 0, 152,
		153, 5, 13, 0, 0, 153, 155, 3, 10, 5, 0, 154, 152, 1, 0, 0, 0, 155, 158,
		1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159, 1, 0,
		0, 0, 158, 156, 1, 0, 0, 0, 159, 160, 5, 2, 0, 0, 160, 162, 1, 0, 0, 0,
		161, 139, 1, 0, 0, 0, 161, 150, 1, 0, 0, 0, 162, 19, 1, 0, 0, 0, 163, 164,
		7, 2, 0, 0, 164, 21, 1, 0, 0, 0, 165, 166, 5, 30, 0, 0, 166, 23, 1, 0,
		0, 0, 167, 168, 7, 3, 0, 0, 168, 25, 1, 0, 0, 0, 169, 171, 7, 4, 0, 0,
		170, 169, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172,
		173, 3, 28, 14, 0, 173, 27, 1, 0, 0, 0, 174, 175, 5, 6, 0, 0, 175, 29,
		1, 0, 0, 0, 21, 31, 37, 47, 52, 58, 61, 63, 71, 73, 80, 85, 89, 111, 117,
		126, 129, 137, 145, 156, 161, 170,
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

// IqlParserInit initializes any static state used to implement IqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewIqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func IqlParserInit() {
	staticData := &IqlParserStaticData
	staticData.once.Do(iqlParserInit)
}

// NewIqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewIqlParser(input antlr.TokenStream) *IqlParser {
	IqlParserInit()
	this := new(IqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &IqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Iql.g4"

	return this
}

// IqlParser tokens.
const (
	IqlParserEOF                          = antlr.TokenEOF
	IqlParserT__0                         = 1
	IqlParserT__1                         = 2
	IqlParserFN_SUM                       = 3
	IqlParserFN_COUNT                     = 4
	IqlParserFN_STARTOFWEEK               = 5
	IqlParserDATETIME                     = 6
	IqlParserNUMBER                       = 7
	IqlParserWHITESPACE                   = 8
	IqlParserSCOL                         = 9
	IqlParserDOT                          = 10
	IqlParserOPEN_PAR                     = 11
	IqlParserCLOSE_PAR                    = 12
	IqlParserCOMMA                        = 13
	IqlParserEQ                           = 14
	IqlParserSTAR                         = 15
	IqlParserCONTAINS                     = 16
	IqlParserNOT_CONTAINS                 = 17
	IqlParserLT                           = 18
	IqlParserLT_EQ                        = 19
	IqlParserGT                           = 20
	IqlParserGT_EQ                        = 21
	IqlParserNOT_EQ                       = 22
	IqlParserK_AFTER                      = 23
	IqlParserK_AND                        = 24
	IqlParserK_ASC                        = 25
	IqlParserK_BEFORE                     = 26
	IqlParserK_BY                         = 27
	IqlParserK_CHANGED                    = 28
	IqlParserK_DESC                       = 29
	IqlParserK_EMPTY                      = 30
	IqlParserK_IN                         = 31
	IqlParserK_IS                         = 32
	IqlParserK_NOT                        = 33
	IqlParserK_NULL                       = 34
	IqlParserK_ON                         = 35
	IqlParserK_OR                         = 36
	IqlParserK_ORDER                      = 37
	IqlParserK_TO                         = 38
	IqlParserK_WAS                        = 39
	IqlParserF_AFFECTED_VERSION           = 40
	IqlParserF_APPROVALS                  = 41
	IqlParserF_ASSIGNEE                   = 42
	IqlParserF_ATTACHMENTS                = 43
	IqlParserF_CATEGORY                   = 44
	IqlParserF_COMMENT                    = 45
	IqlParserF_COMPONENT                  = 46
	IqlParserF_CREATED                    = 47
	IqlParserF_CREATED_DATE               = 48
	IqlParserF_CREATOR                    = 49
	IqlParserF_CUSTOM_FIELD               = 50
	IqlParserF_CUSTOMER_REQUEST_TYPE      = 51
	IqlParserF_DATE                       = 52
	IqlParserF_DESCRIPTION                = 53
	IqlParserF_DUE                        = 54
	IqlParserF_DURATION                   = 55
	IqlParserF_ENVIRONMENT                = 56
	IqlParserF_EPIC_LINK                  = 57
	IqlParserF_FILTER                     = 58
	IqlParserF_FIX_VERSION                = 59
	IqlParserF_ISSUE                      = 60
	IqlParserF_ISSUE_KEY                  = 61
	IqlParserF_ISSUE_TYPE                 = 62
	IqlParserF_KEY                        = 63
	IqlParserF_LABEL                      = 64
	IqlParserF_LABELS                     = 65
	IqlParserF_LAST_VIEWED                = 66
	IqlParserF_LEVEL                      = 67
	IqlParserF_NUMBER                     = 68
	IqlParserF_ORGANIZATION               = 69
	IqlParserF_ORIGINAL_ESTIMATE          = 70
	IqlParserF_PARENT                     = 71
	IqlParserF_PRIORITY                   = 72
	IqlParserF_PROJECT                    = 73
	IqlParserF_RANK                       = 74
	IqlParserF_REMAINING_ESTIMATE         = 75
	IqlParserF_REPORTER                   = 76
	IqlParserF_REQUEST_CHANNEL_TYPE       = 77
	IqlParserF_REQUEST_LAST_ACTIVITY_TIME = 78
	IqlParserF_RESOLUTION                 = 79
	IqlParserF_RESOLUTION_DATE            = 80
	IqlParserF_RESOLVED                   = 81
	IqlParserF_SLA                        = 82
	IqlParserF_SPRINT                     = 83
	IqlParserF_STATUS                     = 84
	IqlParserF_SUMMARY                    = 85
	IqlParserF_TEXT                       = 86
	IqlParserF_TIME_SPENT                 = 87
	IqlParserF_TYPE                       = 88
	IqlParserF_UPDATED                    = 89
	IqlParserF_USER                       = 90
	IqlParserF_VERSION                    = 91
	IqlParserF_VOTER                      = 92
	IqlParserF_VOTES                      = 93
	IqlParserF_WATCHER                    = 94
	IqlParserF_WATCHERS                   = 95
	IqlParserF_WORK_RATIO                 = 96
	IqlParserIDENTIFIER                   = 97
	IqlParserSTRING_LITERAL               = 98
	IqlParserCOMMENT                      = 99
	IqlParserLINE_COMMENT                 = 100
	IqlParserSPACES                       = 101
)

// IqlParser rules.
const (
	IqlParserRULE_query          = 0
	IqlParserRULE_iql_stmt       = 1
	IqlParserRULE_expr           = 2
	IqlParserRULE_ordering_term  = 3
	IqlParserRULE_operator       = 4
	IqlParserRULE_literal_value  = 5
	IqlParserRULE_function       = 6
	IqlParserRULE_function_name  = 7
	IqlParserRULE_function_param = 8
	IqlParserRULE_literal_list   = 9
	IqlParserRULE_keyword        = 10
	IqlParserRULE_state_name     = 11
	IqlParserRULE_field          = 12
	IqlParserRULE_compare_dates  = 13
	IqlParserRULE_dates          = 14
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Iql_stmt() IIql_stmtContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(IqlParserEOF, 0)
}

func (s *QueryContext) Iql_stmt() IIql_stmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIql_stmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIql_stmtContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *IqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, IqlParserRULE_query)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-1090921691136) != 0) || ((int64((_la-64)) & ^0x3f) == 0 && ((int64(1)<<(_la-64))&25769803775) != 0) {
		{
			p.SetState(30)
			p.Iql_stmt()
		}

	}
	{
		p.SetState(33)
		p.Match(IqlParserEOF)
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

// IIql_stmtContext is an interface to support dynamic dispatch.
type IIql_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Ordering_term() IOrdering_termContext

	// IsIql_stmtContext differentiates from other interfaces.
	IsIql_stmtContext()
}

type Iql_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIql_stmtContext() *Iql_stmtContext {
	var p = new(Iql_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_iql_stmt
	return p
}

func InitEmptyIql_stmtContext(p *Iql_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_iql_stmt
}

func (*Iql_stmtContext) IsIql_stmtContext() {}

func NewIql_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iql_stmtContext {
	var p = new(Iql_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_iql_stmt

	return p
}

func (s *Iql_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Iql_stmtContext) Expr() IExprContext {
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

func (s *Iql_stmtContext) Ordering_term() IOrdering_termContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrdering_termContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrdering_termContext)
}

func (s *Iql_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iql_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iql_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterIql_stmt(s)
	}
}

func (s *Iql_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitIql_stmt(s)
	}
}

func (p *IqlParser) Iql_stmt() (localctx IIql_stmtContext) {
	localctx = NewIql_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, IqlParserRULE_iql_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.expr(0)
	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == IqlParserK_ORDER {
		{
			p.SetState(36)
			p.Ordering_term()
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAR() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	CLOSE_PAR() antlr.TerminalNode
	Operator() IOperatorContext
	Field() IFieldContext
	STRING_LITERAL() antlr.TerminalNode
	Literal_value() ILiteral_valueContext
	Literal_list() ILiteral_listContext
	Function() IFunctionContext
	AllK_NOT() []antlr.TerminalNode
	K_NOT(i int) antlr.TerminalNode
	Ordering_term() IOrdering_termContext
	K_AND() antlr.TerminalNode
	K_OR() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserOPEN_PAR, 0)
}

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

func (s *ExprContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserCLOSE_PAR, 0)
}

func (s *ExprContext) Operator() IOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperatorContext)
}

func (s *ExprContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExprContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(IqlParserSTRING_LITERAL, 0)
}

func (s *ExprContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *ExprContext) Literal_list() ILiteral_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_listContext)
}

func (s *ExprContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ExprContext) AllK_NOT() []antlr.TerminalNode {
	return s.GetTokens(IqlParserK_NOT)
}

func (s *ExprContext) K_NOT(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserK_NOT, i)
}

func (s *ExprContext) Ordering_term() IOrdering_termContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrdering_termContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrdering_termContext)
}

func (s *ExprContext) K_AND() antlr.TerminalNode {
	return s.GetToken(IqlParserK_AND, 0)
}

func (s *ExprContext) K_OR() antlr.TerminalNode {
	return s.GetToken(IqlParserK_OR, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *IqlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *IqlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, IqlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IqlParserOPEN_PAR:
		{
			p.SetState(40)
			p.Match(IqlParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(41)
			p.expr(0)
		}
		{
			p.SetState(42)
			p.Match(IqlParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IqlParserK_NOT, IqlParserF_AFFECTED_VERSION, IqlParserF_APPROVALS, IqlParserF_ASSIGNEE, IqlParserF_ATTACHMENTS, IqlParserF_CATEGORY, IqlParserF_COMMENT, IqlParserF_COMPONENT, IqlParserF_CREATED, IqlParserF_CREATED_DATE, IqlParserF_CREATOR, IqlParserF_CUSTOM_FIELD, IqlParserF_CUSTOMER_REQUEST_TYPE, IqlParserF_DATE, IqlParserF_DESCRIPTION, IqlParserF_DUE, IqlParserF_DURATION, IqlParserF_ENVIRONMENT, IqlParserF_EPIC_LINK, IqlParserF_FILTER, IqlParserF_FIX_VERSION, IqlParserF_ISSUE, IqlParserF_ISSUE_KEY, IqlParserF_ISSUE_TYPE, IqlParserF_KEY, IqlParserF_LABEL, IqlParserF_LABELS, IqlParserF_LAST_VIEWED, IqlParserF_LEVEL, IqlParserF_NUMBER, IqlParserF_ORGANIZATION, IqlParserF_ORIGINAL_ESTIMATE, IqlParserF_PARENT, IqlParserF_PRIORITY, IqlParserF_PROJECT, IqlParserF_RANK, IqlParserF_REMAINING_ESTIMATE, IqlParserF_REPORTER, IqlParserF_REQUEST_CHANNEL_TYPE, IqlParserF_REQUEST_LAST_ACTIVITY_TIME, IqlParserF_RESOLUTION, IqlParserF_RESOLUTION_DATE, IqlParserF_RESOLVED, IqlParserF_SLA, IqlParserF_SPRINT, IqlParserF_STATUS, IqlParserF_SUMMARY, IqlParserF_TEXT, IqlParserF_TIME_SPENT, IqlParserF_TYPE, IqlParserF_UPDATED, IqlParserF_USER, IqlParserF_VERSION, IqlParserF_VOTER, IqlParserF_VOTES, IqlParserF_WATCHER, IqlParserF_WATCHERS, IqlParserF_WORK_RATIO, IqlParserSTRING_LITERAL:
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IqlParserK_NOT {
			{
				p.SetState(44)
				p.Match(IqlParserK_NOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case IqlParserF_AFFECTED_VERSION, IqlParserF_APPROVALS, IqlParserF_ASSIGNEE, IqlParserF_ATTACHMENTS, IqlParserF_CATEGORY, IqlParserF_COMMENT, IqlParserF_COMPONENT, IqlParserF_CREATED, IqlParserF_CREATED_DATE, IqlParserF_CREATOR, IqlParserF_CUSTOM_FIELD, IqlParserF_CUSTOMER_REQUEST_TYPE, IqlParserF_DATE, IqlParserF_DESCRIPTION, IqlParserF_DUE, IqlParserF_DURATION, IqlParserF_ENVIRONMENT, IqlParserF_EPIC_LINK, IqlParserF_FILTER, IqlParserF_FIX_VERSION, IqlParserF_ISSUE, IqlParserF_ISSUE_KEY, IqlParserF_ISSUE_TYPE, IqlParserF_KEY, IqlParserF_LABEL, IqlParserF_LABELS, IqlParserF_LAST_VIEWED, IqlParserF_LEVEL, IqlParserF_NUMBER, IqlParserF_ORGANIZATION, IqlParserF_ORIGINAL_ESTIMATE, IqlParserF_PARENT, IqlParserF_PRIORITY, IqlParserF_PROJECT, IqlParserF_RANK, IqlParserF_REMAINING_ESTIMATE, IqlParserF_REPORTER, IqlParserF_REQUEST_CHANNEL_TYPE, IqlParserF_REQUEST_LAST_ACTIVITY_TIME, IqlParserF_RESOLUTION, IqlParserF_RESOLUTION_DATE, IqlParserF_RESOLVED, IqlParserF_SLA, IqlParserF_SPRINT, IqlParserF_STATUS, IqlParserF_SUMMARY, IqlParserF_TEXT, IqlParserF_TIME_SPENT, IqlParserF_TYPE, IqlParserF_UPDATED, IqlParserF_USER, IqlParserF_VERSION, IqlParserF_VOTER, IqlParserF_VOTES, IqlParserF_WATCHER, IqlParserF_WATCHERS, IqlParserF_WORK_RATIO:
			{
				p.SetState(50)
				p.Field()
			}

		case IqlParserSTRING_LITERAL:
			{
				p.SetState(51)
				p.Match(IqlParserSTRING_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(54)
			p.Operator()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case IqlParserDATETIME, IqlParserNUMBER, IqlParserIDENTIFIER, IqlParserSTRING_LITERAL:
			{
				p.SetState(55)
				p.Literal_value()
			}

		case IqlParserT__0, IqlParserOPEN_PAR:
			{
				p.SetState(56)
				p.Literal_list()
			}

		case IqlParserFN_SUM, IqlParserFN_STARTOFWEEK:
			{
				p.SetState(57)
				p.Function()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(60)
				p.Ordering_term()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IqlParserRULE_expr)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(66)
					p.Match(IqlParserK_AND)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(67)
					p.expr(5)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, IqlParserRULE_expr)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(69)
					p.Match(IqlParserK_OR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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

// IOrdering_termContext is an interface to support dynamic dispatch.
type IOrdering_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_ORDER() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	AllLiteral_value() []ILiteral_valueContext
	Literal_value(i int) ILiteral_valueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllK_ASC() []antlr.TerminalNode
	K_ASC(i int) antlr.TerminalNode
	AllK_DESC() []antlr.TerminalNode
	K_DESC(i int) antlr.TerminalNode

	// IsOrdering_termContext differentiates from other interfaces.
	IsOrdering_termContext()
}

type Ordering_termContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdering_termContext() *Ordering_termContext {
	var p = new(Ordering_termContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_ordering_term
	return p
}

func InitEmptyOrdering_termContext(p *Ordering_termContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_ordering_term
}

func (*Ordering_termContext) IsOrdering_termContext() {}

func NewOrdering_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ordering_termContext {
	var p = new(Ordering_termContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_ordering_term

	return p
}

func (s *Ordering_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ordering_termContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(IqlParserK_ORDER, 0)
}

func (s *Ordering_termContext) K_BY() antlr.TerminalNode {
	return s.GetToken(IqlParserK_BY, 0)
}

func (s *Ordering_termContext) AllLiteral_value() []ILiteral_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			len++
		}
	}

	tst := make([]ILiteral_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteral_valueContext); ok {
			tst[i] = t.(ILiteral_valueContext)
			i++
		}
	}

	return tst
}

func (s *Ordering_termContext) Literal_value(i int) ILiteral_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
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

	return t.(ILiteral_valueContext)
}

func (s *Ordering_termContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IqlParserCOMMA)
}

func (s *Ordering_termContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserCOMMA, i)
}

func (s *Ordering_termContext) AllK_ASC() []antlr.TerminalNode {
	return s.GetTokens(IqlParserK_ASC)
}

func (s *Ordering_termContext) K_ASC(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserK_ASC, i)
}

func (s *Ordering_termContext) AllK_DESC() []antlr.TerminalNode {
	return s.GetTokens(IqlParserK_DESC)
}

func (s *Ordering_termContext) K_DESC(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserK_DESC, i)
}

func (s *Ordering_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ordering_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ordering_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterOrdering_term(s)
	}
}

func (s *Ordering_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitOrdering_term(s)
	}
}

func (p *IqlParser) Ordering_term() (localctx IOrdering_termContext) {
	localctx = NewOrdering_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, IqlParserRULE_ordering_term)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(IqlParserK_ORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(IqlParserK_BY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.Literal_value()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(79)
			_la = p.GetTokenStream().LA(1)

			if !(_la == IqlParserK_ASC || _la == IqlParserK_DESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(82)
				p.Match(IqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(83)
				p.Literal_value()
			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(84)
					_la = p.GetTokenStream().LA(1)

					if !(_la == IqlParserK_ASC || _la == IqlParserK_DESC) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperatorContext is an interface to support dynamic dispatch.
type IOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NOT_EQ() antlr.TerminalNode
	CONTAINS() antlr.TerminalNode
	NOT_CONTAINS() antlr.TerminalNode
	LT_EQ() antlr.TerminalNode
	LT() antlr.TerminalNode
	GT() antlr.TerminalNode
	GT_EQ() antlr.TerminalNode
	K_IN() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_IS() antlr.TerminalNode
	K_WAS() antlr.TerminalNode
	K_CHANGED() antlr.TerminalNode
	K_TO() antlr.TerminalNode

	// IsOperatorContext differentiates from other interfaces.
	IsOperatorContext()
}

type OperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperatorContext() *OperatorContext {
	var p = new(OperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_operator
	return p
}

func InitEmptyOperatorContext(p *OperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_operator
}

func (*OperatorContext) IsOperatorContext() {}

func NewOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperatorContext {
	var p = new(OperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_operator

	return p
}

func (s *OperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *OperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(IqlParserEQ, 0)
}

func (s *OperatorContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(IqlParserNOT_EQ, 0)
}

func (s *OperatorContext) CONTAINS() antlr.TerminalNode {
	return s.GetToken(IqlParserCONTAINS, 0)
}

func (s *OperatorContext) NOT_CONTAINS() antlr.TerminalNode {
	return s.GetToken(IqlParserNOT_CONTAINS, 0)
}

func (s *OperatorContext) LT_EQ() antlr.TerminalNode {
	return s.GetToken(IqlParserLT_EQ, 0)
}

func (s *OperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(IqlParserLT, 0)
}

func (s *OperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(IqlParserGT, 0)
}

func (s *OperatorContext) GT_EQ() antlr.TerminalNode {
	return s.GetToken(IqlParserGT_EQ, 0)
}

func (s *OperatorContext) K_IN() antlr.TerminalNode {
	return s.GetToken(IqlParserK_IN, 0)
}

func (s *OperatorContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(IqlParserK_NOT, 0)
}

func (s *OperatorContext) K_IS() antlr.TerminalNode {
	return s.GetToken(IqlParserK_IS, 0)
}

func (s *OperatorContext) K_WAS() antlr.TerminalNode {
	return s.GetToken(IqlParserK_WAS, 0)
}

func (s *OperatorContext) K_CHANGED() antlr.TerminalNode {
	return s.GetToken(IqlParserK_CHANGED, 0)
}

func (s *OperatorContext) K_TO() antlr.TerminalNode {
	return s.GetToken(IqlParserK_TO, 0)
}

func (s *OperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterOperator(s)
	}
}

func (s *OperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitOperator(s)
	}
}

func (p *IqlParser) Operator() (localctx IOperatorContext) {
	localctx = NewOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, IqlParserRULE_operator)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Match(IqlParserEQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Match(IqlParserNOT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Match(IqlParserCONTAINS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(95)
			p.Match(IqlParserNOT_CONTAINS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(96)
			p.Match(IqlParserLT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(97)
			p.Match(IqlParserLT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(98)
			p.Match(IqlParserGT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(99)
			p.Match(IqlParserGT_EQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(100)
			p.Match(IqlParserK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(101)
			p.Match(IqlParserK_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.Match(IqlParserK_IN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(103)
			p.Match(IqlParserK_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(104)
			p.Match(IqlParserK_WAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(105)
			p.Match(IqlParserK_IS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(IqlParserK_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(107)
			p.Match(IqlParserK_WAS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Match(IqlParserK_NOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(109)
			p.Match(IqlParserK_CHANGED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(110)
			p.Match(IqlParserK_TO)
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

// ILiteral_valueContext is an interface to support dynamic dispatch.
type ILiteral_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	Dates() IDatesContext

	// IsLiteral_valueContext differentiates from other interfaces.
	IsLiteral_valueContext()
}

type Literal_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_valueContext() *Literal_valueContext {
	var p = new(Literal_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_literal_value
	return p
}

func InitEmptyLiteral_valueContext(p *Literal_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_literal_value
}

func (*Literal_valueContext) IsLiteral_valueContext() {}

func NewLiteral_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_valueContext {
	var p = new(Literal_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_literal_value

	return p
}

func (s *Literal_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_valueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(IqlParserSTRING_LITERAL, 0)
}

func (s *Literal_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(IqlParserIDENTIFIER, 0)
}

func (s *Literal_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(IqlParserNUMBER, 0)
}

func (s *Literal_valueContext) Dates() IDatesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatesContext)
}

func (s *Literal_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterLiteral_value(s)
	}
}

func (s *Literal_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitLiteral_value(s)
	}
}

func (p *IqlParser) Literal_value() (localctx ILiteral_valueContext) {
	localctx = NewLiteral_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, IqlParserRULE_literal_value)
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(113)
			p.Match(IqlParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.Match(IqlParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IqlParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			p.Match(IqlParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IqlParserDATETIME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(116)
			p.Dates()
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

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_name() IFunction_nameContext
	OPEN_PAR() antlr.TerminalNode
	CLOSE_PAR() antlr.TerminalNode
	AllFunction_param() []IFunction_paramContext
	Function_param(i int) IFunction_paramContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function
	return p
}

func InitEmptyFunctionContext(p *FunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Function_name() IFunction_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *FunctionContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserOPEN_PAR, 0)
}

func (s *FunctionContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserCLOSE_PAR, 0)
}

func (s *FunctionContext) AllFunction_param() []IFunction_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunction_paramContext); ok {
			len++
		}
	}

	tst := make([]IFunction_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunction_paramContext); ok {
			tst[i] = t.(IFunction_paramContext)
			i++
		}
	}

	return tst
}

func (s *FunctionContext) Function_param(i int) IFunction_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_paramContext); ok {
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

	return t.(IFunction_paramContext)
}

func (s *FunctionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IqlParserCOMMA)
}

func (s *FunctionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserCOMMA, i)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *IqlParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, IqlParserRULE_function)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Function_name()
	}
	{
		p.SetState(120)
		p.Match(IqlParserOPEN_PAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&232) != 0) || _la == IqlParserIDENTIFIER || _la == IqlParserSTRING_LITERAL {
		{
			p.SetState(121)
			p.Function_param()
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IqlParserCOMMA {
			{
				p.SetState(122)
				p.Match(IqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(123)
				p.Function_param()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(131)
		p.Match(IqlParserCLOSE_PAR)
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

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FN_SUM() antlr.TerminalNode
	FN_STARTOFWEEK() antlr.TerminalNode

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function_name
	return p
}

func InitEmptyFunction_nameContext(p *Function_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function_name
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) FN_SUM() antlr.TerminalNode {
	return s.GetToken(IqlParserFN_SUM, 0)
}

func (s *Function_nameContext) FN_STARTOFWEEK() antlr.TerminalNode {
	return s.GetToken(IqlParserFN_STARTOFWEEK, 0)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (p *IqlParser) Function_name() (localctx IFunction_nameContext) {
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, IqlParserRULE_function_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(_la == IqlParserFN_SUM || _la == IqlParserFN_STARTOFWEEK) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IFunction_paramContext is an interface to support dynamic dispatch.
type IFunction_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal_value() ILiteral_valueContext
	Function() IFunctionContext

	// IsFunction_paramContext differentiates from other interfaces.
	IsFunction_paramContext()
}

type Function_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_paramContext() *Function_paramContext {
	var p = new(Function_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function_param
	return p
}

func InitEmptyFunction_paramContext(p *Function_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_function_param
}

func (*Function_paramContext) IsFunction_paramContext() {}

func NewFunction_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_paramContext {
	var p = new(Function_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_function_param

	return p
}

func (s *Function_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_paramContext) Literal_value() ILiteral_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteral_valueContext)
}

func (s *Function_paramContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Function_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterFunction_param(s)
	}
}

func (s *Function_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitFunction_param(s)
	}
}

func (p *IqlParser) Function_param() (localctx IFunction_paramContext) {
	localctx = NewFunction_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, IqlParserRULE_function_param)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IqlParserDATETIME, IqlParserNUMBER, IqlParserIDENTIFIER, IqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Literal_value()
		}

	case IqlParserFN_SUM, IqlParserFN_STARTOFWEEK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.Function()
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

// ILiteral_listContext is an interface to support dynamic dispatch.
type ILiteral_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPEN_PAR() antlr.TerminalNode
	AllLiteral_value() []ILiteral_valueContext
	Literal_value(i int) ILiteral_valueContext
	CLOSE_PAR() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLiteral_listContext differentiates from other interfaces.
	IsLiteral_listContext()
}

type Literal_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_listContext() *Literal_listContext {
	var p = new(Literal_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_literal_list
	return p
}

func InitEmptyLiteral_listContext(p *Literal_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_literal_list
}

func (*Literal_listContext) IsLiteral_listContext() {}

func NewLiteral_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_listContext {
	var p = new(Literal_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_literal_list

	return p
}

func (s *Literal_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_listContext) OPEN_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserOPEN_PAR, 0)
}

func (s *Literal_listContext) AllLiteral_value() []ILiteral_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILiteral_valueContext); ok {
			len++
		}
	}

	tst := make([]ILiteral_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILiteral_valueContext); ok {
			tst[i] = t.(ILiteral_valueContext)
			i++
		}
	}

	return tst
}

func (s *Literal_listContext) Literal_value(i int) ILiteral_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteral_valueContext); ok {
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

	return t.(ILiteral_valueContext)
}

func (s *Literal_listContext) CLOSE_PAR() antlr.TerminalNode {
	return s.GetToken(IqlParserCLOSE_PAR, 0)
}

func (s *Literal_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(IqlParserCOMMA)
}

func (s *Literal_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(IqlParserCOMMA, i)
}

func (s *Literal_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterLiteral_list(s)
	}
}

func (s *Literal_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitLiteral_list(s)
	}
}

func (p *IqlParser) Literal_list() (localctx ILiteral_listContext) {
	localctx = NewLiteral_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, IqlParserRULE_literal_list)
	var _la int

	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case IqlParserOPEN_PAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(IqlParserOPEN_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.Literal_value()
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IqlParserCOMMA {
			{
				p.SetState(141)
				p.Match(IqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(142)
				p.Literal_value()
			}

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(148)
			p.Match(IqlParserCLOSE_PAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case IqlParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(IqlParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Literal_value()
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == IqlParserCOMMA {
			{
				p.SetState(152)
				p.Match(IqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(153)
				p.Literal_value()
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(159)
			p.Match(IqlParserT__1)
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

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_AFTER() antlr.TerminalNode
	K_AND() antlr.TerminalNode
	K_ASC() antlr.TerminalNode
	K_BEFORE() antlr.TerminalNode
	K_BY() antlr.TerminalNode
	K_CHANGED() antlr.TerminalNode
	K_DESC() antlr.TerminalNode
	K_IN() antlr.TerminalNode
	K_IS() antlr.TerminalNode
	K_NOT() antlr.TerminalNode
	K_NULL() antlr.TerminalNode
	K_ON() antlr.TerminalNode
	K_OR() antlr.TerminalNode
	K_ORDER() antlr.TerminalNode
	K_TO() antlr.TerminalNode
	K_WAS() antlr.TerminalNode

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_keyword
	return p
}

func InitEmptyKeywordContext(p *KeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_keyword
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) K_AFTER() antlr.TerminalNode {
	return s.GetToken(IqlParserK_AFTER, 0)
}

func (s *KeywordContext) K_AND() antlr.TerminalNode {
	return s.GetToken(IqlParserK_AND, 0)
}

func (s *KeywordContext) K_ASC() antlr.TerminalNode {
	return s.GetToken(IqlParserK_ASC, 0)
}

func (s *KeywordContext) K_BEFORE() antlr.TerminalNode {
	return s.GetToken(IqlParserK_BEFORE, 0)
}

func (s *KeywordContext) K_BY() antlr.TerminalNode {
	return s.GetToken(IqlParserK_BY, 0)
}

func (s *KeywordContext) K_CHANGED() antlr.TerminalNode {
	return s.GetToken(IqlParserK_CHANGED, 0)
}

func (s *KeywordContext) K_DESC() antlr.TerminalNode {
	return s.GetToken(IqlParserK_DESC, 0)
}

func (s *KeywordContext) K_IN() antlr.TerminalNode {
	return s.GetToken(IqlParserK_IN, 0)
}

func (s *KeywordContext) K_IS() antlr.TerminalNode {
	return s.GetToken(IqlParserK_IS, 0)
}

func (s *KeywordContext) K_NOT() antlr.TerminalNode {
	return s.GetToken(IqlParserK_NOT, 0)
}

func (s *KeywordContext) K_NULL() antlr.TerminalNode {
	return s.GetToken(IqlParserK_NULL, 0)
}

func (s *KeywordContext) K_ON() antlr.TerminalNode {
	return s.GetToken(IqlParserK_ON, 0)
}

func (s *KeywordContext) K_OR() antlr.TerminalNode {
	return s.GetToken(IqlParserK_OR, 0)
}

func (s *KeywordContext) K_ORDER() antlr.TerminalNode {
	return s.GetToken(IqlParserK_ORDER, 0)
}

func (s *KeywordContext) K_TO() antlr.TerminalNode {
	return s.GetToken(IqlParserK_TO, 0)
}

func (s *KeywordContext) K_WAS() antlr.TerminalNode {
	return s.GetToken(IqlParserK_WAS, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *IqlParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, IqlParserRULE_keyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1098429497344) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// IState_nameContext is an interface to support dynamic dispatch.
type IState_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	K_EMPTY() antlr.TerminalNode

	// IsState_nameContext differentiates from other interfaces.
	IsState_nameContext()
}

type State_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyState_nameContext() *State_nameContext {
	var p = new(State_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_state_name
	return p
}

func InitEmptyState_nameContext(p *State_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_state_name
}

func (*State_nameContext) IsState_nameContext() {}

func NewState_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *State_nameContext {
	var p = new(State_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_state_name

	return p
}

func (s *State_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *State_nameContext) K_EMPTY() antlr.TerminalNode {
	return s.GetToken(IqlParserK_EMPTY, 0)
}

func (s *State_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *State_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *State_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterState_name(s)
	}
}

func (s *State_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitState_name(s)
	}
}

func (p *IqlParser) State_name() (localctx IState_nameContext) {
	localctx = NewState_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, IqlParserRULE_state_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(IqlParserK_EMPTY)
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	F_AFFECTED_VERSION() antlr.TerminalNode
	F_APPROVALS() antlr.TerminalNode
	F_ASSIGNEE() antlr.TerminalNode
	F_ATTACHMENTS() antlr.TerminalNode
	F_CATEGORY() antlr.TerminalNode
	F_COMMENT() antlr.TerminalNode
	F_COMPONENT() antlr.TerminalNode
	F_CREATED() antlr.TerminalNode
	F_CREATED_DATE() antlr.TerminalNode
	F_CREATOR() antlr.TerminalNode
	F_CUSTOM_FIELD() antlr.TerminalNode
	F_CUSTOMER_REQUEST_TYPE() antlr.TerminalNode
	F_DATE() antlr.TerminalNode
	F_DESCRIPTION() antlr.TerminalNode
	F_DUE() antlr.TerminalNode
	F_DURATION() antlr.TerminalNode
	F_ENVIRONMENT() antlr.TerminalNode
	F_EPIC_LINK() antlr.TerminalNode
	F_FILTER() antlr.TerminalNode
	F_FIX_VERSION() antlr.TerminalNode
	F_ISSUE() antlr.TerminalNode
	F_ISSUE_KEY() antlr.TerminalNode
	F_ISSUE_TYPE() antlr.TerminalNode
	F_KEY() antlr.TerminalNode
	F_LABEL() antlr.TerminalNode
	F_LABELS() antlr.TerminalNode
	F_LAST_VIEWED() antlr.TerminalNode
	F_LEVEL() antlr.TerminalNode
	F_NUMBER() antlr.TerminalNode
	F_ORGANIZATION() antlr.TerminalNode
	F_ORIGINAL_ESTIMATE() antlr.TerminalNode
	F_PARENT() antlr.TerminalNode
	F_PRIORITY() antlr.TerminalNode
	F_PROJECT() antlr.TerminalNode
	F_RANK() antlr.TerminalNode
	F_REMAINING_ESTIMATE() antlr.TerminalNode
	F_REPORTER() antlr.TerminalNode
	F_REQUEST_CHANNEL_TYPE() antlr.TerminalNode
	F_REQUEST_LAST_ACTIVITY_TIME() antlr.TerminalNode
	F_RESOLUTION() antlr.TerminalNode
	F_RESOLUTION_DATE() antlr.TerminalNode
	F_RESOLVED() antlr.TerminalNode
	F_SLA() antlr.TerminalNode
	F_SPRINT() antlr.TerminalNode
	F_STATUS() antlr.TerminalNode
	F_SUMMARY() antlr.TerminalNode
	F_TEXT() antlr.TerminalNode
	F_TIME_SPENT() antlr.TerminalNode
	F_TYPE() antlr.TerminalNode
	F_UPDATED() antlr.TerminalNode
	F_USER() antlr.TerminalNode
	F_VERSION() antlr.TerminalNode
	F_VOTER() antlr.TerminalNode
	F_VOTES() antlr.TerminalNode
	F_WATCHER() antlr.TerminalNode
	F_WATCHERS() antlr.TerminalNode
	F_WORK_RATIO() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) F_AFFECTED_VERSION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_AFFECTED_VERSION, 0)
}

func (s *FieldContext) F_APPROVALS() antlr.TerminalNode {
	return s.GetToken(IqlParserF_APPROVALS, 0)
}

func (s *FieldContext) F_ASSIGNEE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ASSIGNEE, 0)
}

func (s *FieldContext) F_ATTACHMENTS() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ATTACHMENTS, 0)
}

func (s *FieldContext) F_CATEGORY() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CATEGORY, 0)
}

func (s *FieldContext) F_COMMENT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_COMMENT, 0)
}

func (s *FieldContext) F_COMPONENT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_COMPONENT, 0)
}

func (s *FieldContext) F_CREATED() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CREATED, 0)
}

func (s *FieldContext) F_CREATED_DATE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CREATED_DATE, 0)
}

func (s *FieldContext) F_CREATOR() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CREATOR, 0)
}

func (s *FieldContext) F_CUSTOM_FIELD() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CUSTOM_FIELD, 0)
}

func (s *FieldContext) F_CUSTOMER_REQUEST_TYPE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_CUSTOMER_REQUEST_TYPE, 0)
}

func (s *FieldContext) F_DATE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_DATE, 0)
}

func (s *FieldContext) F_DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_DESCRIPTION, 0)
}

func (s *FieldContext) F_DUE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_DUE, 0)
}

func (s *FieldContext) F_DURATION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_DURATION, 0)
}

func (s *FieldContext) F_ENVIRONMENT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ENVIRONMENT, 0)
}

func (s *FieldContext) F_EPIC_LINK() antlr.TerminalNode {
	return s.GetToken(IqlParserF_EPIC_LINK, 0)
}

func (s *FieldContext) F_FILTER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_FILTER, 0)
}

func (s *FieldContext) F_FIX_VERSION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_FIX_VERSION, 0)
}

func (s *FieldContext) F_ISSUE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ISSUE, 0)
}

func (s *FieldContext) F_ISSUE_KEY() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ISSUE_KEY, 0)
}

func (s *FieldContext) F_ISSUE_TYPE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ISSUE_TYPE, 0)
}

func (s *FieldContext) F_KEY() antlr.TerminalNode {
	return s.GetToken(IqlParserF_KEY, 0)
}

func (s *FieldContext) F_LABEL() antlr.TerminalNode {
	return s.GetToken(IqlParserF_LABEL, 0)
}

func (s *FieldContext) F_LABELS() antlr.TerminalNode {
	return s.GetToken(IqlParserF_LABELS, 0)
}

func (s *FieldContext) F_LAST_VIEWED() antlr.TerminalNode {
	return s.GetToken(IqlParserF_LAST_VIEWED, 0)
}

func (s *FieldContext) F_LEVEL() antlr.TerminalNode {
	return s.GetToken(IqlParserF_LEVEL, 0)
}

func (s *FieldContext) F_NUMBER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_NUMBER, 0)
}

func (s *FieldContext) F_ORGANIZATION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ORGANIZATION, 0)
}

func (s *FieldContext) F_ORIGINAL_ESTIMATE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_ORIGINAL_ESTIMATE, 0)
}

func (s *FieldContext) F_PARENT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_PARENT, 0)
}

func (s *FieldContext) F_PRIORITY() antlr.TerminalNode {
	return s.GetToken(IqlParserF_PRIORITY, 0)
}

func (s *FieldContext) F_PROJECT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_PROJECT, 0)
}

func (s *FieldContext) F_RANK() antlr.TerminalNode {
	return s.GetToken(IqlParserF_RANK, 0)
}

func (s *FieldContext) F_REMAINING_ESTIMATE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_REMAINING_ESTIMATE, 0)
}

func (s *FieldContext) F_REPORTER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_REPORTER, 0)
}

func (s *FieldContext) F_REQUEST_CHANNEL_TYPE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_REQUEST_CHANNEL_TYPE, 0)
}

func (s *FieldContext) F_REQUEST_LAST_ACTIVITY_TIME() antlr.TerminalNode {
	return s.GetToken(IqlParserF_REQUEST_LAST_ACTIVITY_TIME, 0)
}

func (s *FieldContext) F_RESOLUTION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_RESOLUTION, 0)
}

func (s *FieldContext) F_RESOLUTION_DATE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_RESOLUTION_DATE, 0)
}

func (s *FieldContext) F_RESOLVED() antlr.TerminalNode {
	return s.GetToken(IqlParserF_RESOLVED, 0)
}

func (s *FieldContext) F_SLA() antlr.TerminalNode {
	return s.GetToken(IqlParserF_SLA, 0)
}

func (s *FieldContext) F_SPRINT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_SPRINT, 0)
}

func (s *FieldContext) F_STATUS() antlr.TerminalNode {
	return s.GetToken(IqlParserF_STATUS, 0)
}

func (s *FieldContext) F_SUMMARY() antlr.TerminalNode {
	return s.GetToken(IqlParserF_SUMMARY, 0)
}

func (s *FieldContext) F_TEXT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_TEXT, 0)
}

func (s *FieldContext) F_TIME_SPENT() antlr.TerminalNode {
	return s.GetToken(IqlParserF_TIME_SPENT, 0)
}

func (s *FieldContext) F_TYPE() antlr.TerminalNode {
	return s.GetToken(IqlParserF_TYPE, 0)
}

func (s *FieldContext) F_UPDATED() antlr.TerminalNode {
	return s.GetToken(IqlParserF_UPDATED, 0)
}

func (s *FieldContext) F_USER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_USER, 0)
}

func (s *FieldContext) F_VERSION() antlr.TerminalNode {
	return s.GetToken(IqlParserF_VERSION, 0)
}

func (s *FieldContext) F_VOTER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_VOTER, 0)
}

func (s *FieldContext) F_VOTES() antlr.TerminalNode {
	return s.GetToken(IqlParserF_VOTES, 0)
}

func (s *FieldContext) F_WATCHER() antlr.TerminalNode {
	return s.GetToken(IqlParserF_WATCHER, 0)
}

func (s *FieldContext) F_WATCHERS() antlr.TerminalNode {
	return s.GetToken(IqlParserF_WATCHERS, 0)
}

func (s *FieldContext) F_WORK_RATIO() antlr.TerminalNode {
	return s.GetToken(IqlParserF_WORK_RATIO, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *IqlParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, IqlParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-40)) & ^0x3f) == 0 && ((int64(1)<<(_la-40))&144115188075855871) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ICompare_datesContext is an interface to support dynamic dispatch.
type ICompare_datesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Dates() IDatesContext
	K_ON() antlr.TerminalNode
	K_AFTER() antlr.TerminalNode
	K_BEFORE() antlr.TerminalNode

	// IsCompare_datesContext differentiates from other interfaces.
	IsCompare_datesContext()
}

type Compare_datesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompare_datesContext() *Compare_datesContext {
	var p = new(Compare_datesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_compare_dates
	return p
}

func InitEmptyCompare_datesContext(p *Compare_datesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_compare_dates
}

func (*Compare_datesContext) IsCompare_datesContext() {}

func NewCompare_datesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compare_datesContext {
	var p = new(Compare_datesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_compare_dates

	return p
}

func (s *Compare_datesContext) GetParser() antlr.Parser { return s.parser }

func (s *Compare_datesContext) Dates() IDatesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatesContext)
}

func (s *Compare_datesContext) K_ON() antlr.TerminalNode {
	return s.GetToken(IqlParserK_ON, 0)
}

func (s *Compare_datesContext) K_AFTER() antlr.TerminalNode {
	return s.GetToken(IqlParserK_AFTER, 0)
}

func (s *Compare_datesContext) K_BEFORE() antlr.TerminalNode {
	return s.GetToken(IqlParserK_BEFORE, 0)
}

func (s *Compare_datesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compare_datesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compare_datesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterCompare_dates(s)
	}
}

func (s *Compare_datesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitCompare_dates(s)
	}
}

func (p *IqlParser) Compare_dates() (localctx ICompare_datesContext) {
	localctx = NewCompare_datesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, IqlParserRULE_compare_dates)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34435235840) != 0 {
		{
			p.SetState(169)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34435235840) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(172)
		p.Dates()
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

// IDatesContext is an interface to support dynamic dispatch.
type IDatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATETIME() antlr.TerminalNode

	// IsDatesContext differentiates from other interfaces.
	IsDatesContext()
}

type DatesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatesContext() *DatesContext {
	var p = new(DatesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_dates
	return p
}

func InitEmptyDatesContext(p *DatesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = IqlParserRULE_dates
}

func (*DatesContext) IsDatesContext() {}

func NewDatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatesContext {
	var p = new(DatesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = IqlParserRULE_dates

	return p
}

func (s *DatesContext) GetParser() antlr.Parser { return s.parser }

func (s *DatesContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(IqlParserDATETIME, 0)
}

func (s *DatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.EnterDates(s)
	}
}

func (s *DatesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(IqlListener); ok {
		listenerT.ExitDates(s)
	}
}

func (p *IqlParser) Dates() (localctx IDatesContext) {
	localctx = NewDatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, IqlParserRULE_dates)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(IqlParserDATETIME)
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

func (p *IqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
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

func (p *IqlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

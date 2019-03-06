// Generated from Gherkin.g4 by ANTLR 4.7.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 139,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 52, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 5, 8, 111, 10, 8, 3, 9, 6, 9, 114, 10, 9, 13, 9, 14,
	9, 115, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 6, 17, 136,
	10, 17, 13, 17, 14, 17, 137, 2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 3, 2, 4, 4, 2, 11, 11, 34, 34, 10, 2, 11, 12, 15, 15, 34, 34, 36,
	37, 60, 60, 66, 66, 94, 94, 126, 126, 2, 144, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 3, 51, 3, 2, 2,
	2, 5, 53, 3, 2, 2, 2, 7, 61, 3, 2, 2, 2, 9, 72, 3, 2, 2, 2, 11, 81, 3,
	2, 2, 2, 13, 98, 3, 2, 2, 2, 15, 110, 3, 2, 2, 2, 17, 113, 3, 2, 2, 2,
	19, 117, 3, 2, 2, 2, 21, 119, 3, 2, 2, 2, 23, 122, 3, 2, 2, 2, 25, 124,
	3, 2, 2, 2, 27, 126, 3, 2, 2, 2, 29, 128, 3, 2, 2, 2, 31, 132, 3, 2, 2,
	2, 33, 135, 3, 2, 2, 2, 35, 36, 7, 73, 2, 2, 36, 37, 7, 107, 2, 2, 37,
	38, 7, 120, 2, 2, 38, 39, 7, 103, 2, 2, 39, 52, 7, 112, 2, 2, 40, 41, 7,
	89, 2, 2, 41, 42, 7, 106, 2, 2, 42, 43, 7, 103, 2, 2, 43, 52, 7, 112, 2,
	2, 44, 45, 7, 86, 2, 2, 45, 46, 7, 106, 2, 2, 46, 47, 7, 103, 2, 2, 47,
	52, 7, 112, 2, 2, 48, 49, 7, 67, 2, 2, 49, 50, 7, 112, 2, 2, 50, 52, 7,
	102, 2, 2, 51, 35, 3, 2, 2, 2, 51, 40, 3, 2, 2, 2, 51, 44, 3, 2, 2, 2,
	51, 48, 3, 2, 2, 2, 52, 4, 3, 2, 2, 2, 53, 54, 7, 72, 2, 2, 54, 55, 7,
	103, 2, 2, 55, 56, 7, 99, 2, 2, 56, 57, 7, 118, 2, 2, 57, 58, 7, 119, 2,
	2, 58, 59, 7, 116, 2, 2, 59, 60, 7, 103, 2, 2, 60, 6, 3, 2, 2, 2, 61, 62,
	7, 68, 2, 2, 62, 63, 7, 99, 2, 2, 63, 64, 7, 101, 2, 2, 64, 65, 7, 109,
	2, 2, 65, 66, 7, 105, 2, 2, 66, 67, 7, 116, 2, 2, 67, 68, 7, 113, 2, 2,
	68, 69, 7, 119, 2, 2, 69, 70, 7, 112, 2, 2, 70, 71, 7, 102, 2, 2, 71, 8,
	3, 2, 2, 2, 72, 73, 7, 85, 2, 2, 73, 74, 7, 101, 2, 2, 74, 75, 7, 103,
	2, 2, 75, 76, 7, 112, 2, 2, 76, 77, 7, 99, 2, 2, 77, 78, 7, 116, 2, 2,
	78, 79, 7, 107, 2, 2, 79, 80, 7, 113, 2, 2, 80, 10, 3, 2, 2, 2, 81, 82,
	7, 85, 2, 2, 82, 83, 7, 101, 2, 2, 83, 84, 7, 103, 2, 2, 84, 85, 7, 112,
	2, 2, 85, 86, 7, 99, 2, 2, 86, 87, 7, 116, 2, 2, 87, 88, 7, 107, 2, 2,
	88, 89, 7, 113, 2, 2, 89, 90, 7, 34, 2, 2, 90, 91, 7, 81, 2, 2, 91, 92,
	7, 119, 2, 2, 92, 93, 7, 118, 2, 2, 93, 94, 7, 110, 2, 2, 94, 95, 7, 107,
	2, 2, 95, 96, 7, 112, 2, 2, 96, 97, 7, 103, 2, 2, 97, 12, 3, 2, 2, 2, 98,
	99, 7, 71, 2, 2, 99, 100, 7, 122, 2, 2, 100, 101, 7, 99, 2, 2, 101, 102,
	7, 111, 2, 2, 102, 103, 7, 114, 2, 2, 103, 104, 7, 110, 2, 2, 104, 105,
	7, 103, 2, 2, 105, 106, 7, 117, 2, 2, 106, 14, 3, 2, 2, 2, 107, 111, 7,
	12, 2, 2, 108, 109, 7, 15, 2, 2, 109, 111, 7, 12, 2, 2, 110, 107, 3, 2,
	2, 2, 110, 108, 3, 2, 2, 2, 111, 16, 3, 2, 2, 2, 112, 114, 9, 2, 2, 2,
	113, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 18, 3, 2, 2, 2, 117, 118, 7, 126, 2, 2, 118, 20,
	3, 2, 2, 2, 119, 120, 7, 94, 2, 2, 120, 121, 7, 126, 2, 2, 121, 22, 3,
	2, 2, 2, 122, 123, 7, 60, 2, 2, 123, 24, 3, 2, 2, 2, 124, 125, 7, 66, 2,
	2, 125, 26, 3, 2, 2, 2, 126, 127, 7, 37, 2, 2, 127, 28, 3, 2, 2, 2, 128,
	129, 7, 36, 2, 2, 129, 130, 7, 36, 2, 2, 130, 131, 7, 36, 2, 2, 131, 30,
	3, 2, 2, 2, 132, 133, 7, 36, 2, 2, 133, 32, 3, 2, 2, 2, 134, 136, 10, 3,
	2, 2, 135, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2,
	137, 138, 3, 2, 2, 2, 138, 34, 3, 2, 2, 2, 7, 2, 51, 110, 115, 137, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "'Feature'", "'Background'", "'Scenario'", "'Scenario Outline'",
	"'Examples'", "", "", "'|'", "'\\|'", "':'", "'@'", "'#'", "'\"\"\"'",
	"'\"'",
}

var lexerSymbolicNames = []string{
	"", "STEP_KW", "FEATURE_KW", "BACKGROUND_KW", "SCENARIO_KW", "OUTLINE_KW",
	"EXAMPLES_KW", "EOL", "WS", "PIPE", "ESCAPED_PIPE", "COLON", "AT", "HASH",
	"TRIPLE_QUOTE", "QUOTE", "CHAR",
}

var lexerRuleNames = []string{
	"STEP_KW", "FEATURE_KW", "BACKGROUND_KW", "SCENARIO_KW", "OUTLINE_KW",
	"EXAMPLES_KW", "EOL", "WS", "PIPE", "ESCAPED_PIPE", "COLON", "AT", "HASH",
	"TRIPLE_QUOTE", "QUOTE", "CHAR",
}

type GherkinLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGherkinLexer(input antlr.CharStream) *GherkinLexer {

	l := new(GherkinLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gherkin.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GherkinLexer tokens.
const (
	GherkinLexerSTEP_KW       = 1
	GherkinLexerFEATURE_KW    = 2
	GherkinLexerBACKGROUND_KW = 3
	GherkinLexerSCENARIO_KW   = 4
	GherkinLexerOUTLINE_KW    = 5
	GherkinLexerEXAMPLES_KW   = 6
	GherkinLexerEOL           = 7
	GherkinLexerWS            = 8
	GherkinLexerPIPE          = 9
	GherkinLexerESCAPED_PIPE  = 10
	GherkinLexerCOLON         = 11
	GherkinLexerAT            = 12
	GherkinLexerHASH          = 13
	GherkinLexerTRIPLE_QUOTE  = 14
	GherkinLexerQUOTE         = 15
	GherkinLexerCHAR          = 16
)

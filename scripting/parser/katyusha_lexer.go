// Code generated from Katyusha.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 21, 185,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14, 3, 56,
	11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 10, 5, 10, 95, 10, 10, 3, 10, 6, 10, 98, 10, 10, 13, 10, 14, 10,
	99, 3, 10, 3, 10, 6, 10, 104, 10, 10, 13, 10, 14, 10, 105, 5, 10, 108,
	10, 10, 3, 10, 6, 10, 111, 10, 10, 13, 10, 14, 10, 112, 3, 11, 3, 11, 5,
	11, 117, 10, 11, 3, 11, 6, 11, 120, 10, 11, 13, 11, 14, 11, 121, 3, 12,
	3, 12, 6, 12, 126, 10, 12, 13, 12, 14, 12, 127, 3, 13, 6, 13, 131, 10,
	13, 13, 13, 14, 13, 132, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 151,
	10, 18, 12, 18, 14, 18, 154, 11, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 7, 19, 162, 10, 19, 12, 19, 14, 19, 165, 11, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 6, 20, 171, 10, 20, 13, 20, 14, 20, 172, 3, 21, 6, 21, 176,
	10, 21, 13, 21, 14, 21, 177, 3, 22, 3, 22, 5, 22, 182, 10, 22, 3, 22, 3,
	22, 2, 2, 23, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 2, 41, 2, 43, 21, 3, 2, 10, 3, 2, 12, 12, 3, 2, 50, 59, 5, 2, 106,
	106, 111, 111, 117, 117, 5, 2, 67, 92, 97, 97, 99, 124, 7, 2, 47, 48, 50,
	60, 67, 92, 97, 97, 99, 124, 8, 2, 44, 44, 47, 48, 50, 59, 65, 65, 67,
	92, 99, 124, 6, 2, 12, 12, 14, 15, 36, 36, 94, 94, 6, 2, 12, 12, 14, 15,
	49, 49, 94, 94, 2, 199, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2,
	2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3,
	2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23,
	3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2,
	31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 3, 45, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 57, 3, 2, 2,
	2, 9, 61, 3, 2, 2, 2, 11, 65, 3, 2, 2, 2, 13, 72, 3, 2, 2, 2, 15, 77, 3,
	2, 2, 2, 17, 83, 3, 2, 2, 2, 19, 110, 3, 2, 2, 2, 21, 116, 3, 2, 2, 2,
	23, 123, 3, 2, 2, 2, 25, 130, 3, 2, 2, 2, 27, 134, 3, 2, 2, 2, 29, 137,
	3, 2, 2, 2, 31, 140, 3, 2, 2, 2, 33, 143, 3, 2, 2, 2, 35, 146, 3, 2, 2,
	2, 37, 157, 3, 2, 2, 2, 39, 168, 3, 2, 2, 2, 41, 175, 3, 2, 2, 2, 43, 181,
	3, 2, 2, 2, 45, 46, 7, 12, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 116, 2,
	2, 48, 49, 7, 119, 2, 2, 49, 50, 7, 112, 2, 2, 50, 54, 3, 2, 2, 2, 51,
	53, 10, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2,
	2, 2, 54, 55, 3, 2, 2, 2, 55, 6, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 58,
	7, 117, 2, 2, 58, 59, 7, 103, 2, 2, 59, 60, 7, 118, 2, 2, 60, 8, 3, 2,
	2, 2, 61, 62, 7, 99, 2, 2, 62, 63, 7, 102, 2, 2, 63, 64, 7, 102, 2, 2,
	64, 10, 3, 2, 2, 2, 65, 66, 7, 116, 2, 2, 66, 67, 7, 103, 2, 2, 67, 68,
	7, 111, 2, 2, 68, 69, 7, 113, 2, 2, 69, 70, 7, 120, 2, 2, 70, 71, 7, 103,
	2, 2, 71, 12, 3, 2, 2, 2, 72, 73, 7, 110, 2, 2, 73, 74, 7, 107, 2, 2, 74,
	75, 7, 117, 2, 2, 75, 76, 7, 118, 2, 2, 76, 14, 3, 2, 2, 2, 77, 78, 7,
	106, 2, 2, 78, 79, 7, 113, 2, 2, 79, 80, 7, 117, 2, 2, 80, 81, 7, 118,
	2, 2, 81, 82, 7, 117, 2, 2, 82, 16, 3, 2, 2, 2, 83, 84, 7, 47, 2, 2, 84,
	85, 7, 47, 2, 2, 85, 86, 7, 113, 2, 2, 86, 87, 7, 112, 2, 2, 87, 88, 7,
	103, 2, 2, 88, 89, 7, 110, 2, 2, 89, 90, 7, 107, 2, 2, 90, 91, 7, 112,
	2, 2, 91, 92, 7, 103, 2, 2, 92, 18, 3, 2, 2, 2, 93, 95, 7, 47, 2, 2, 94,
	93, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 98, 9, 3, 2,
	2, 97, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100,
	3, 2, 2, 2, 100, 107, 3, 2, 2, 2, 101, 103, 7, 48, 2, 2, 102, 104, 9, 3,
	2, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2,
	105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 101, 3, 2, 2, 2, 107,
	108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 111, 9, 4, 2, 2, 110, 94, 3,
	2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2,
	2, 113, 20, 3, 2, 2, 2, 114, 115, 7, 50, 2, 2, 115, 117, 7, 122, 2, 2,
	116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118,
	120, 9, 3, 2, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 119,
	3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 22, 3, 2, 2, 2, 123, 125, 9, 5,
	2, 2, 124, 126, 9, 6, 2, 2, 125, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2,
	127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 24, 3, 2, 2, 2, 129, 131,
	9, 7, 2, 2, 130, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 130, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 26, 3, 2, 2, 2, 134, 135, 7, 63, 2, 2,
	135, 136, 7, 63, 2, 2, 136, 28, 3, 2, 2, 2, 137, 138, 7, 63, 2, 2, 138,
	139, 7, 128, 2, 2, 139, 30, 3, 2, 2, 2, 140, 141, 7, 35, 2, 2, 141, 142,
	7, 63, 2, 2, 142, 32, 3, 2, 2, 2, 143, 144, 7, 35, 2, 2, 144, 145, 7, 128,
	2, 2, 145, 34, 3, 2, 2, 2, 146, 152, 7, 36, 2, 2, 147, 148, 7, 94, 2, 2,
	148, 151, 11, 2, 2, 2, 149, 151, 10, 8, 2, 2, 150, 147, 3, 2, 2, 2, 150,
	149, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153,
	3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 155, 156, 7, 36,
	2, 2, 156, 36, 3, 2, 2, 2, 157, 163, 7, 49, 2, 2, 158, 159, 7, 94, 2, 2,
	159, 162, 11, 2, 2, 2, 160, 162, 10, 9, 2, 2, 161, 158, 3, 2, 2, 2, 161,
	160, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164,
	3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 167, 7, 49,
	2, 2, 167, 38, 3, 2, 2, 2, 168, 170, 7, 37, 2, 2, 169, 171, 10, 2, 2, 2,
	170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172,
	173, 3, 2, 2, 2, 173, 40, 3, 2, 2, 2, 174, 176, 7, 34, 2, 2, 175, 174,
	3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2,
	2, 2, 178, 42, 3, 2, 2, 2, 179, 182, 5, 41, 21, 2, 180, 182, 5, 39, 20,
	2, 181, 179, 3, 2, 2, 2, 181, 180, 3, 2, 2, 2, 182, 183, 3, 2, 2, 2, 183,
	184, 8, 22, 2, 2, 184, 44, 3, 2, 2, 2, 20, 2, 54, 94, 99, 105, 107, 112,
	116, 121, 127, 132, 150, 152, 161, 163, 172, 177, 181, 3, 8, 2, 2,
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
	"", "'\n'", "", "'set'", "'add'", "'remove'", "'list'", "'hosts'", "'--oneline'",
	"", "", "", "", "'=='", "'=~'", "'!='", "'!~'",
}

var lexerSymbolicNames = []string{
	"", "", "RUN", "SET", "ADD", "REMOVE", "LIST", "HOSTS", "ONELINE", "DURATION",
	"NUMBER", "IDENTIFIER", "GLOB", "EQUALS", "MATCHES", "NOT_EQUALS", "NOT_MATCHES",
	"STRING", "REGEXP", "SKIP_",
}

var lexerRuleNames = []string{
	"T__0", "RUN", "SET", "ADD", "REMOVE", "LIST", "HOSTS", "ONELINE", "DURATION",
	"NUMBER", "IDENTIFIER", "GLOB", "EQUALS", "MATCHES", "NOT_EQUALS", "NOT_MATCHES",
	"STRING", "REGEXP", "COMMENT", "SPACES", "SKIP_",
}

type KatyushaLexer struct {
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

func NewKatyushaLexer(input antlr.CharStream) *KatyushaLexer {

	l := new(KatyushaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Katyusha.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// KatyushaLexer tokens.
const (
	KatyushaLexerT__0        = 1
	KatyushaLexerRUN         = 2
	KatyushaLexerSET         = 3
	KatyushaLexerADD         = 4
	KatyushaLexerREMOVE      = 5
	KatyushaLexerLIST        = 6
	KatyushaLexerHOSTS       = 7
	KatyushaLexerONELINE     = 8
	KatyushaLexerDURATION    = 9
	KatyushaLexerNUMBER      = 10
	KatyushaLexerIDENTIFIER  = 11
	KatyushaLexerGLOB        = 12
	KatyushaLexerEQUALS      = 13
	KatyushaLexerMATCHES     = 14
	KatyushaLexerNOT_EQUALS  = 15
	KatyushaLexerNOT_MATCHES = 16
	KatyushaLexerSTRING      = 17
	KatyushaLexerREGEXP      = 18
	KatyushaLexerSKIP_       = 19
)

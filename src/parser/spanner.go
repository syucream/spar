// Code generated by goyacc -o src/parser/spanner.go src/parser/spanner.go.y. DO NOT EDIT.
//line src/parser/spanner.go.y:2
package parser

import __yyfmt__ "fmt"

//line src/parser/spanner.go.y:2
import (
	"strconv"

	"github.com/syucream/spar/src/types"
)

//line src/parser/spanner.go.y:11
type yySymType struct {
	yys       int
	empty     struct{}
	flag      bool
	i64       int64
	str       string
	strs      []string
	col       types.Column
	cols      []types.Column
	coltype   types.ColumnType
	key       types.Key
	keys      []types.Key
	keyorder  types.KeyOrder
	clstr     types.Cluster
	ondelete  types.OnDelete
	stcls     types.StoringClause
	intlr     types.Interleave
	intlrs    []types.Interleave
	alt       types.Alteration
	LastToken int
}

const PRIMARY = 57346
const KEY = 57347
const ASC = 57348
const DESC = 57349
const INTERLEAVE = 57350
const IN = 57351
const PARENT = 57352
const ARRAY = 57353
const OPTIONS = 57354
const NOT = 57355
const NULL = 57356
const ON = 57357
const DELETE = 57358
const CASCADE = 57359
const NO = 57360
const ACTION = 57361
const MAX = 57362
const UNIQUE = 57363
const NULL_FILTERED = 57364
const STORING = 57365
const ADD = 57366
const COLUMN = 57367
const SET = 57368
const true = 57369
const null = 57370
const allow_commit_timestamp = 57371
const CREATE = 57372
const ALTER = 57373
const DROP = 57374
const DATABASE = 57375
const TABLE = 57376
const INDEX = 57377
const BOOL = 57378
const INT64 = 57379
const FLOAT64 = 57380
const STRING = 57381
const BYTES = 57382
const DATE = 57383
const TIMESTAMP = 57384
const database_id = 57385
const table_name = 57386
const column_name = 57387
const index_name = 57388
const decimal_value = 57389
const hex_value = 57390

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PRIMARY",
	"KEY",
	"ASC",
	"DESC",
	"INTERLEAVE",
	"IN",
	"PARENT",
	"ARRAY",
	"OPTIONS",
	"NOT",
	"NULL",
	"ON",
	"DELETE",
	"CASCADE",
	"NO",
	"ACTION",
	"MAX",
	"UNIQUE",
	"NULL_FILTERED",
	"STORING",
	"ADD",
	"COLUMN",
	"SET",
	"true",
	"null",
	"allow_commit_timestamp",
	"'('",
	"','",
	"')'",
	"';'",
	"'='",
	"CREATE",
	"ALTER",
	"DROP",
	"DATABASE",
	"TABLE",
	"INDEX",
	"BOOL",
	"INT64",
	"FLOAT64",
	"STRING",
	"BYTES",
	"DATE",
	"TIMESTAMP",
	"database_id",
	"table_name",
	"column_name",
	"index_name",
	"decimal_value",
	"hex_value",
	"'<'",
	"'>'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 157

var yyAct = [...]int{

	127, 47, 106, 105, 84, 53, 71, 87, 52, 104,
	89, 75, 41, 44, 62, 32, 141, 136, 107, 43,
	67, 62, 55, 56, 57, 58, 59, 60, 61, 80,
	65, 138, 123, 76, 31, 30, 27, 26, 42, 24,
	25, 22, 90, 91, 55, 56, 57, 58, 59, 60,
	61, 55, 56, 57, 58, 59, 60, 61, 19, 20,
	37, 34, 39, 23, 70, 9, 10, 11, 118, 140,
	139, 132, 40, 38, 111, 117, 79, 111, 112, 133,
	18, 93, 92, 17, 64, 97, 96, 16, 15, 14,
	13, 131, 103, 102, 50, 82, 51, 129, 101, 100,
	94, 74, 73, 33, 109, 124, 125, 49, 110, 46,
	45, 122, 29, 95, 119, 77, 78, 66, 48, 63,
	86, 72, 85, 116, 134, 130, 108, 128, 114, 115,
	83, 99, 69, 2, 137, 12, 8, 7, 6, 5,
	4, 3, 1, 36, 35, 126, 120, 121, 135, 28,
	21, 98, 81, 68, 113, 88, 54,
}
var yyPact = [...]int{

	30, 30, -1000, 57, 56, 55, 54, 50, 47, 20,
	24, 0, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -11,
	-13, 90, -1000, -14, -15, -36, -1000, 73, 21, -1000,
	36, -1000, -1000, -31, -38, -1000, -1000, 85, 84, 103,
	82, 62, 65, 10, 104, -31, -20, -1000, 101, -30,
	128, -31, 108, -1000, -1000, -1000, -1000, -1000, 72, 71,
	-1000, -1000, -43, -16, -1000, -1000, 98, 3, 64, 125,
	-1000, 110, 106, -10, -10, -19, 70, -1000, 94, 108,
	110, -1000, 123, 69, -1000, 68, -1000, 61, -1000, -1000,
	-1000, -1000, 60, -46, -32, -1000, -1000, -1000, -1000, 117,
	-32, 79, -1000, -1000, -1000, 46, -1000, 122, 113, 43,
	34, -32, 88, -1000, -1000, -1000, -17, -1000, 78, -1000,
	119, -1000, 67, 103, 59, 39, 48, -1000, 115, -33,
	-1000, -1000, -1000, 119, -18, 38, -1000, -1000, -1000, -1000,
	-34, -1000,
}
var yyPgo = [...]int{

	0, 38, 12, 8, 5, 156, 7, 155, 4, 154,
	2, 3, 153, 152, 151, 1, 6, 150, 149, 148,
	147, 146, 0, 145, 144, 143, 142, 133, 141, 140,
	139, 138, 137, 136,
}
var yyR1 = [...]int{

	0, 26, 26, 27, 27, 27, 27, 27, 27, 28,
	29, 2, 2, 2, 1, 12, 11, 11, 10, 9,
	9, 9, 13, 13, 14, 15, 15, 15, 3, 3,
	4, 4, 4, 4, 4, 4, 4, 6, 6, 5,
	8, 8, 8, 16, 16, 30, 17, 17, 18, 18,
	21, 21, 20, 19, 19, 23, 23, 23, 22, 31,
	31, 24, 24, 24, 25, 25, 32, 33, 7, 7,
}
var yyR2 = [...]int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 3,
	8, 0, 1, 3, 4, 5, 1, 3, 2, 0,
	1, 1, 0, 2, 5, 0, 3, 4, 1, 1,
	1, 1, 1, 4, 4, 1, 1, 1, 1, 4,
	0, 6, 6, 0, 2, 12, 0, 1, 0, 1,
	0, 1, 4, 1, 3, 0, 1, 3, 3, 4,
	4, 3, 3, 2, 5, 5, 3, 3, 1, 1,
}
var yyChk = [...]int{

	-1000, -26, -27, -28, -29, -30, -31, -32, -33, 35,
	36, 37, -27, 33, 33, 33, 33, 33, 33, 38,
	39, -17, 21, 39, 39, 40, 48, 49, -18, 22,
	49, 49, 51, 30, 40, -24, -25, 24, 37, 26,
	36, -2, -1, 50, 51, 25, 25, -15, 15, 25,
	32, 31, -3, -4, -5, 41, 42, 43, 44, 45,
	46, 47, 11, 15, -1, 50, 16, 50, -12, 4,
	-2, -16, 13, 30, 30, 54, 49, 17, 18, -3,
	26, -13, 31, 5, -8, 12, 14, -6, -7, 20,
	52, 53, -6, -4, 30, 19, -16, -8, -14, 8,
	30, 30, 32, 32, 55, -11, -10, 50, 9, -11,
	29, 31, 32, -9, 6, 7, 10, 32, 34, -10,
	-21, -20, 23, 49, 27, 28, -23, -22, 8, 30,
	-15, 32, 32, 31, 9, -19, 50, -22, 49, 32,
	31, 50,
}
var yyDef = [...]int{

	0, -2, 1, 0, 0, 0, 0, 0, 0, 46,
	0, 0, 2, 3, 4, 5, 6, 7, 8, 0,
	0, 48, 47, 0, 0, 0, 9, 0, 0, 49,
	0, 66, 67, 11, 0, 59, 60, 0, 0, 25,
	0, 0, 12, 0, 0, 0, 0, 63, 0, 0,
	0, 11, 43, 28, 29, 30, 31, 32, 0, 0,
	35, 36, 0, 0, 61, 62, 0, 0, 22, 0,
	13, 40, 0, 0, 0, 0, 0, 26, 0, 43,
	40, 10, 0, 0, 14, 0, 44, 0, 37, 38,
	68, 69, 0, 0, 0, 27, 64, 65, 23, 0,
	0, 0, 33, 34, 39, 0, 16, 19, 0, 0,
	0, 0, 50, 18, 20, 21, 0, 15, 0, 17,
	55, 51, 0, 25, 0, 0, 45, 56, 0, 0,
	24, 41, 42, 0, 0, 0, 53, 57, 58, 52,
	0, 54,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	30, 32, 3, 3, 31, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 33,
	54, 34, 55,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 35, 36,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:97
		{
			s := types.CreateDatabaseStatement{
				DatabaseId: yyDollar[3].str,
			}
			yylex.(*lexerWrapper).result.CreateDatabases = append(yylex.(*lexerWrapper).result.CreateDatabases, s)
		}
	case 10:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line src/parser/spanner.go.y:106
		{
			s := types.CreateTableStatement{
				TableName:   yyDollar[3].str,
				Columns:     yyDollar[5].cols,
				PrimaryKeys: yyDollar[7].keys,
				Cluster:     yyDollar[8].clstr,
			}
			yylex.(*lexerWrapper).result.CreateTables = append(yylex.(*lexerWrapper).result.CreateTables, s)
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:118
		{
			yyVAL.cols = make([]types.Column, 0, 0)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:122
		{
			yyVAL.cols = make([]types.Column, 0, 1)
			yyVAL.cols = append(yyVAL.cols, yyDollar[1].col)
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:127
		{
			yyVAL.cols = append(yyDollar[3].cols, yyDollar[1].col)
		}
	case 14:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:133
		{
			yyVAL.col = types.Column{Name: yyDollar[1].str, Type: yyDollar[2].coltype, NotNull: yyDollar[3].flag, Options: yyDollar[4].str}
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:139
		{
			yyVAL.keys = yyDollar[4].keys
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:145
		{
			yyVAL.keys = make([]types.Key, 0, 1)
			yyVAL.keys = append(yyVAL.keys, yyDollar[1].key)
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:150
		{
			yyVAL.keys = append(yyDollar[1].keys, yyDollar[3].key)
		}
	case 18:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:156
		{
			yyVAL.key = types.Key{Name: yyDollar[1].str, KeyOrder: yyDollar[2].keyorder}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:162
		{
			yyVAL.keyorder = types.Asc
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:166
		{
			yyVAL.keyorder = types.Asc
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:170
		{
			yyVAL.keyorder = types.Desc
		}
	case 22:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:176
		{
			yyVAL.clstr = types.Cluster{}
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:180
		{
			yyVAL.clstr = yyDollar[2].clstr
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:186
		{
			yyVAL.clstr = types.Cluster{TableName: yyDollar[4].str, OnDelete: yyDollar[5].ondelete}
		}
	case 25:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:192
		{
			// default
			yyVAL.ondelete = types.NoAction
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:197
		{
			yyVAL.ondelete = types.Cascade
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:201
		{
			yyVAL.ondelete = types.NoAction
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:207
		{
			yyVAL.coltype = yyDollar[1].coltype
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:211
		{
			yyVAL.coltype = yyDollar[1].coltype
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:217
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Bool}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:221
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Int64}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:225
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Float64}
		}
	case 33:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:229
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.String, Length: yyDollar[3].i64}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:233
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Bytes, Length: yyDollar[3].i64}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:237
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Date}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:241
		{
			yyVAL.coltype = types.ColumnType{TypeTag: types.Timestamp}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:247
		{
			yyVAL.i64 = yyDollar[1].i64
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:251
		{
			yyVAL.i64 = 10485760
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:257
		{
			yyVAL.coltype = yyDollar[3].coltype
			yyVAL.coltype.IsArray = types.True
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:264
		{
			yyVAL.str = ""
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line src/parser/spanner.go.y:268
		{
			yyVAL.str = yyDollar[3].str + "=" + yyDollar[5].str
		}
	case 42:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line src/parser/spanner.go.y:272
		{
			yyVAL.str = yyDollar[3].str + "=" + yyDollar[5].str
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:278
		{
			yyVAL.flag = types.False
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:282
		{
			yyVAL.flag = types.True
		}
	case 45:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line src/parser/spanner.go.y:288
		{
			s := types.CreateIndexStatement{
				Unique:        yyDollar[2].flag,
				NullFiltered:  yyDollar[3].flag,
				IndexName:     yyDollar[5].str,
				TableName:     yyDollar[7].str,
				Keys:          yyDollar[9].keys,
				StoringClause: yyDollar[11].stcls,
				Interleaves:   yyDollar[12].intlrs,
			}
			yylex.(*lexerWrapper).result.CreateIndexes = append(yylex.(*lexerWrapper).result.CreateIndexes, s)
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:303
		{
			yyVAL.flag = types.False
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:307
		{
			yyVAL.flag = types.True
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:313
		{
			yyVAL.flag = types.False
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:317
		{
			yyVAL.flag = types.True
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:323
		{
			yyVAL.stcls = types.StoringClause{}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:327
		{
			yyVAL.stcls = yyDollar[1].stcls
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:333
		{
			yyVAL.stcls = types.StoringClause{ColumnNames: yyDollar[3].strs}
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:339
		{
			yyVAL.strs = make([]string, 0, 1)
			yyVAL.strs = append(yyVAL.strs, yyDollar[1].str)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:344
		{
			yyVAL.strs = append(yyDollar[1].strs, yyDollar[3].str)
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line src/parser/spanner.go.y:350
		{
			yyVAL.intlrs = make([]types.Interleave, 0, 0)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:354
		{
			yyVAL.intlrs = make([]types.Interleave, 0, 1)
			yyVAL.intlrs = append(yyVAL.intlrs, yyDollar[1].intlr)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:359
		{
			yyVAL.intlrs = append(yyDollar[1].intlrs, yyDollar[3].intlr)
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:365
		{
			yyVAL.intlr = types.Interleave{TableName: yyDollar[3].str}
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:371
		{
			s := types.AlterTableStatement{
				TableName:  yyDollar[3].str,
				Alteration: yyDollar[4].alt,
			}
			yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line src/parser/spanner.go.y:379
		{
			s := types.AlterTableStatement{
				TableName:  yyDollar[3].str,
				Alteration: yyDollar[4].alt,
			}
			yylex.(*lexerWrapper).result.AlterTables = append(yylex.(*lexerWrapper).result.AlterTables, s)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:389
		{
			yyVAL.alt = &types.AddColumnTableAlteration{
				Column: yyDollar[3].col,
			}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:395
		{
			yyVAL.alt = &types.DropColumnTableAlteration{
				ColumnName: yyDollar[3].str,
			}
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line src/parser/spanner.go.y:401
		{
			yyVAL.alt = &types.SetTableAlteration{
				OnDelete: yyDollar[2].ondelete,
			}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:409
		{
			yyVAL.alt = &types.AlterColumnTypesAlteration{
				ColumnName: yyDollar[3].str,
				ColumnType: yyDollar[4].coltype,
				NotNull:    yyDollar[5].flag,
			}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line src/parser/spanner.go.y:417
		{
			yyVAL.alt = &types.AlterColumnSetAlteration{
				ColumnName: yyDollar[3].str,
				Options:    yyDollar[5].str,
			}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:426
		{
			s := types.DropTableStatement{
				TableName: yyDollar[3].str,
			}
			yylex.(*lexerWrapper).result.DropTables = append(yylex.(*lexerWrapper).result.DropTables, s)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line src/parser/spanner.go.y:435
		{
			s := types.DropIndexStatement{
				IndexName: yyDollar[3].str,
			}
			yylex.(*lexerWrapper).result.DropIndexes = append(yylex.(*lexerWrapper).result.DropIndexes, s)
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:444
		{
			v, _ := strconv.ParseInt(yyDollar[1].str, 10, 64)
			yyVAL.i64 = v
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line src/parser/spanner.go.y:449
		{
			v, _ := strconv.ParseInt(yyDollar[1].str, 16, 64)
			yyVAL.i64 = v
		}
	}
	goto yystack /* stack new state and value */
}

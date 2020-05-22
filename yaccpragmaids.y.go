// Code generated by goyacc -o ./yaccpragmaids.y.go -p YaccPragmaIds ./yaccpragmaids.y. DO NOT EDIT.

//line ./yaccpragmaids.y:2
package yaccpragmaids

import __yyfmt__ "fmt"

//line ./yaccpragmaids.y:2

import "strconv"

//line ./yaccpragmaids.y:13
type YaccPragmaIdsSymType struct {
	yys     int
	Value   string
	Version IVersion
}

const RWIdl = 57346
const RWLocal = 57347
const RWRmi = 57348
const RWDce = 57349
const Value = 57350

var YaccPragmaIdsToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"RWIdl",
	"RWLocal",
	"RWRmi",
	"RWDce",
	"Value",
	"':'",
	"'.'",
}
var YaccPragmaIdsStatenames = [...]string{}

const YaccPragmaIdsEofCode = 1
const YaccPragmaIdsErrCode = 2
const YaccPragmaIdsInitialStackSize = 16

//line ./yaccpragmaids.y:123

//line yacctab:1
var YaccPragmaIdsExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const YaccPragmaIdsPrivate = 57344

const YaccPragmaIdsLast = 24

var YaccPragmaIdsAct = [...]int{

	17, 21, 2, 3, 5, 4, 22, 1, 16, 15,
	14, 9, 8, 7, 6, 24, 19, 23, 20, 18,
	13, 12, 11, 10,
}
var YaccPragmaIdsPact = [...]int{

	-2, -1000, 5, 4, 3, 2, 15, 14, 13, 12,
	1, -1000, 0, -1, 11, 11, 10, -1000, -9, -1000,
	-3, 9, 7, -1000, -1000,
}
var YaccPragmaIdsPgo = [...]int{

	0, 0, 7,
}
var YaccPragmaIdsR1 = [...]int{

	0, 2, 2, 2, 2, 2, 1, 1, 1,
}
var YaccPragmaIdsR2 = [...]int{

	0, 5, 3, 5, 5, 7, 1, 2, 3,
}
var YaccPragmaIdsChk = [...]int{

	-1000, -2, 4, 5, 7, 6, 9, 9, 9, 9,
	8, 8, 8, 8, 9, 9, 9, -1, 8, -1,
	8, 10, 9, 8, 8,
}
var YaccPragmaIdsDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 2, 0, 0, 0, 0, 0, 1, 6, 3,
	4, 7, 0, 8, 5,
}
var YaccPragmaIdsTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 10, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 9,
}
var YaccPragmaIdsTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8,
}
var YaccPragmaIdsTok3 = [...]int{
	0,
}

var YaccPragmaIdsErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	YaccPragmaIdsDebug        = 0
	YaccPragmaIdsErrorVerbose = false
)

type YaccPragmaIdsLexer interface {
	Lex(lval *YaccPragmaIdsSymType) int
	Error(s string)
}

type YaccPragmaIdsParser interface {
	Parse(YaccPragmaIdsLexer) int
	Lookahead() int
}

type YaccPragmaIdsParserImpl struct {
	lval  YaccPragmaIdsSymType
	stack [YaccPragmaIdsInitialStackSize]YaccPragmaIdsSymType
	char  int
}

func (p *YaccPragmaIdsParserImpl) Lookahead() int {
	return p.char
}

func YaccPragmaIdsNewParser() YaccPragmaIdsParser {
	return &YaccPragmaIdsParserImpl{}
}

const YaccPragmaIdsFlag = -1000

func YaccPragmaIdsTokname(c int) string {
	if c >= 1 && c-1 < len(YaccPragmaIdsToknames) {
		if YaccPragmaIdsToknames[c-1] != "" {
			return YaccPragmaIdsToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func YaccPragmaIdsStatname(s int) string {
	if s >= 0 && s < len(YaccPragmaIdsStatenames) {
		if YaccPragmaIdsStatenames[s] != "" {
			return YaccPragmaIdsStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func YaccPragmaIdsErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !YaccPragmaIdsErrorVerbose {
		return "syntax error"
	}

	for _, e := range YaccPragmaIdsErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + YaccPragmaIdsTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := YaccPragmaIdsPact[state]
	for tok := TOKSTART; tok-1 < len(YaccPragmaIdsToknames); tok++ {
		if n := base + tok; n >= 0 && n < YaccPragmaIdsLast && YaccPragmaIdsChk[YaccPragmaIdsAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if YaccPragmaIdsDef[state] == -2 {
		i := 0
		for YaccPragmaIdsExca[i] != -1 || YaccPragmaIdsExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; YaccPragmaIdsExca[i] >= 0; i += 2 {
			tok := YaccPragmaIdsExca[i]
			if tok < TOKSTART || YaccPragmaIdsExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if YaccPragmaIdsExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += YaccPragmaIdsTokname(tok)
	}
	return res
}

func YaccPragmaIdslex1(lex YaccPragmaIdsLexer, lval *YaccPragmaIdsSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = YaccPragmaIdsTok1[0]
		goto out
	}
	if char < len(YaccPragmaIdsTok1) {
		token = YaccPragmaIdsTok1[char]
		goto out
	}
	if char >= YaccPragmaIdsPrivate {
		if char < YaccPragmaIdsPrivate+len(YaccPragmaIdsTok2) {
			token = YaccPragmaIdsTok2[char-YaccPragmaIdsPrivate]
			goto out
		}
	}
	for i := 0; i < len(YaccPragmaIdsTok3); i += 2 {
		token = YaccPragmaIdsTok3[i+0]
		if token == char {
			token = YaccPragmaIdsTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = YaccPragmaIdsTok2[1] /* unknown char */
	}
	if YaccPragmaIdsDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", YaccPragmaIdsTokname(token), uint(char))
	}
	return char, token
}

func YaccPragmaIdsParse(YaccPragmaIdslex YaccPragmaIdsLexer) int {
	return YaccPragmaIdsNewParser().Parse(YaccPragmaIdslex)
}

func (YaccPragmaIdsrcvr *YaccPragmaIdsParserImpl) Parse(YaccPragmaIdslex YaccPragmaIdsLexer) int {
	var YaccPragmaIdsn int
	var YaccPragmaIdsVAL YaccPragmaIdsSymType
	var YaccPragmaIdsDollar []YaccPragmaIdsSymType
	_ = YaccPragmaIdsDollar // silence set and not used
	YaccPragmaIdsS := YaccPragmaIdsrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	YaccPragmaIdsstate := 0
	YaccPragmaIdsrcvr.char = -1
	YaccPragmaIdstoken := -1 // YaccPragmaIdsrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		YaccPragmaIdsstate = -1
		YaccPragmaIdsrcvr.char = -1
		YaccPragmaIdstoken = -1
	}()
	YaccPragmaIdsp := -1
	goto YaccPragmaIdsstack

ret0:
	return 0

ret1:
	return 1

YaccPragmaIdsstack:
	/* put a state and value onto the stack */
	if YaccPragmaIdsDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", YaccPragmaIdsTokname(YaccPragmaIdstoken), YaccPragmaIdsStatname(YaccPragmaIdsstate))
	}

	YaccPragmaIdsp++
	if YaccPragmaIdsp >= len(YaccPragmaIdsS) {
		nyys := make([]YaccPragmaIdsSymType, len(YaccPragmaIdsS)*2)
		copy(nyys, YaccPragmaIdsS)
		YaccPragmaIdsS = nyys
	}
	YaccPragmaIdsS[YaccPragmaIdsp] = YaccPragmaIdsVAL
	YaccPragmaIdsS[YaccPragmaIdsp].yys = YaccPragmaIdsstate

YaccPragmaIdsnewstate:
	YaccPragmaIdsn = YaccPragmaIdsPact[YaccPragmaIdsstate]
	if YaccPragmaIdsn <= YaccPragmaIdsFlag {
		goto YaccPragmaIdsdefault /* simple state */
	}
	if YaccPragmaIdsrcvr.char < 0 {
		YaccPragmaIdsrcvr.char, YaccPragmaIdstoken = YaccPragmaIdslex1(YaccPragmaIdslex, &YaccPragmaIdsrcvr.lval)
	}
	YaccPragmaIdsn += YaccPragmaIdstoken
	if YaccPragmaIdsn < 0 || YaccPragmaIdsn >= YaccPragmaIdsLast {
		goto YaccPragmaIdsdefault
	}
	YaccPragmaIdsn = YaccPragmaIdsAct[YaccPragmaIdsn]
	if YaccPragmaIdsChk[YaccPragmaIdsn] == YaccPragmaIdstoken { /* valid shift */
		YaccPragmaIdsrcvr.char = -1
		YaccPragmaIdstoken = -1
		YaccPragmaIdsVAL = YaccPragmaIdsrcvr.lval
		YaccPragmaIdsstate = YaccPragmaIdsn
		if Errflag > 0 {
			Errflag--
		}
		goto YaccPragmaIdsstack
	}

YaccPragmaIdsdefault:
	/* default state action */
	YaccPragmaIdsn = YaccPragmaIdsDef[YaccPragmaIdsstate]
	if YaccPragmaIdsn == -2 {
		if YaccPragmaIdsrcvr.char < 0 {
			YaccPragmaIdsrcvr.char, YaccPragmaIdstoken = YaccPragmaIdslex1(YaccPragmaIdslex, &YaccPragmaIdsrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if YaccPragmaIdsExca[xi+0] == -1 && YaccPragmaIdsExca[xi+1] == YaccPragmaIdsstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			YaccPragmaIdsn = YaccPragmaIdsExca[xi+0]
			if YaccPragmaIdsn < 0 || YaccPragmaIdsn == YaccPragmaIdstoken {
				break
			}
		}
		YaccPragmaIdsn = YaccPragmaIdsExca[xi+1]
		if YaccPragmaIdsn < 0 {
			goto ret0
		}
	}
	if YaccPragmaIdsn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			YaccPragmaIdslex.Error(YaccPragmaIdsErrorMessage(YaccPragmaIdsstate, YaccPragmaIdstoken))
			Nerrs++
			if YaccPragmaIdsDebug >= 1 {
				__yyfmt__.Printf("%s", YaccPragmaIdsStatname(YaccPragmaIdsstate))
				__yyfmt__.Printf(" saw %s\n", YaccPragmaIdsTokname(YaccPragmaIdstoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for YaccPragmaIdsp >= 0 {
				YaccPragmaIdsn = YaccPragmaIdsPact[YaccPragmaIdsS[YaccPragmaIdsp].yys] + YaccPragmaIdsErrCode
				if YaccPragmaIdsn >= 0 && YaccPragmaIdsn < YaccPragmaIdsLast {
					YaccPragmaIdsstate = YaccPragmaIdsAct[YaccPragmaIdsn] /* simulate a shift of "error" */
					if YaccPragmaIdsChk[YaccPragmaIdsstate] == YaccPragmaIdsErrCode {
						goto YaccPragmaIdsstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if YaccPragmaIdsDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", YaccPragmaIdsS[YaccPragmaIdsp].yys)
				}
				YaccPragmaIdsp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if YaccPragmaIdsDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", YaccPragmaIdsTokname(YaccPragmaIdstoken))
			}
			if YaccPragmaIdstoken == YaccPragmaIdsEofCode {
				goto ret1
			}
			YaccPragmaIdsrcvr.char = -1
			YaccPragmaIdstoken = -1
			goto YaccPragmaIdsnewstate /* try again in the same state */
		}
	}

	/* reduction by production YaccPragmaIdsn */
	if YaccPragmaIdsDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", YaccPragmaIdsn, YaccPragmaIdsStatname(YaccPragmaIdsstate))
	}

	YaccPragmaIdsnt := YaccPragmaIdsn
	YaccPragmaIdspt := YaccPragmaIdsp
	_ = YaccPragmaIdspt // guard against "declared and not used"

	YaccPragmaIdsp -= YaccPragmaIdsR2[YaccPragmaIdsn]
	// YaccPragmaIdsp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if YaccPragmaIdsp+1 >= len(YaccPragmaIdsS) {
		nyys := make([]YaccPragmaIdsSymType, len(YaccPragmaIdsS)*2)
		copy(nyys, YaccPragmaIdsS)
		YaccPragmaIdsS = nyys
	}
	YaccPragmaIdsVAL = YaccPragmaIdsS[YaccPragmaIdsp+1]

	/* consult goto table to find next state */
	YaccPragmaIdsn = YaccPragmaIdsR1[YaccPragmaIdsn]
	YaccPragmaIdsg := YaccPragmaIdsPgo[YaccPragmaIdsn]
	YaccPragmaIdsj := YaccPragmaIdsg + YaccPragmaIdsS[YaccPragmaIdsp].yys + 1

	if YaccPragmaIdsj >= YaccPragmaIdsLast {
		YaccPragmaIdsstate = YaccPragmaIdsAct[YaccPragmaIdsg]
	} else {
		YaccPragmaIdsstate = YaccPragmaIdsAct[YaccPragmaIdsj]
		if YaccPragmaIdsChk[YaccPragmaIdsstate] != -YaccPragmaIdsn {
			YaccPragmaIdsstate = YaccPragmaIdsAct[YaccPragmaIdsg]
		}
	}
	// dummy call; replaced with literal code
	switch YaccPragmaIdsnt {

	case 1:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-5 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:28
		{
			if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
				v := NewIdlVersion(YaccPragmaIdsDollar[3].Value, YaccPragmaIdsDollar[5].Version)
				setTypeVersion.SetTypeValue(v)
			}
		}
	case 2:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-3 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:35
		{
			if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
				v := NewLocalVersion(YaccPragmaIdsDollar[3].Value)
				setTypeVersion.SetTypeValue(v)
			}
		}
	case 3:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-5 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:42
		{
			if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
				v, e := NewDceVersion(YaccPragmaIdsDollar[3].Value, YaccPragmaIdsDollar[5].Version)
				if e != nil {
					YaccPragmaIdslex.Error(e.Error())
					return 1
				}
				setTypeVersion.SetTypeValue(v)
			}
		}
	case 4:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-5 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:53
		{
			if len(YaccPragmaIdsDollar[5].Value) != 16 {
				YaccPragmaIdslex.Error("hash value incorrect length")
				return 1
			}
			intValue05, err := strconv.ParseUint(YaccPragmaIdsDollar[5].Value, 16, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
				v := NewRmiVersion(YaccPragmaIdsDollar[3].Value, uint64(intValue05), 0)
				setTypeVersion.SetTypeValue(v)
			}
		}
	case 5:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-7 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:69
		{
			if len(YaccPragmaIdsDollar[5].Value) != 16 {
				YaccPragmaIdslex.Error("hash value incorrect length")
				return 1
			}
			intValue05, err := strconv.ParseUint(YaccPragmaIdsDollar[5].Value, 16, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			intValue07, err := strconv.ParseUint(YaccPragmaIdsDollar[7].Value, 16, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}

			if setTypeVersion, ok := YaccPragmaIdslex.(ISetTypeVersion); ok {
				v := NewRmiVersion(YaccPragmaIdsDollar[3].Value, uint64(intValue05), intValue07)
				setTypeVersion.SetTypeValue(v)
			}
		}
	case 6:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-1 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:92
		{
			intValue01, err := strconv.ParseUint(YaccPragmaIdsDollar[1].Value, 10, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			YaccPragmaIdsVAL.Version = NewVersion(intValue01, 0)
		}
	case 7:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-2 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:101
		{
			intValue01, err := strconv.ParseUint(YaccPragmaIdsDollar[1].Value, 10, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			YaccPragmaIdsVAL.Version = NewVersion(intValue01, 0)
		}
	case 8:
		YaccPragmaIdsDollar = YaccPragmaIdsS[YaccPragmaIdspt-3 : YaccPragmaIdspt+1]
//line ./yaccpragmaids.y:110
		{
			intValue01, err := strconv.ParseUint(YaccPragmaIdsDollar[1].Value, 10, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			intValue03, err := strconv.ParseUint(YaccPragmaIdsDollar[3].Value, 10, 64)
			if err != nil {
				YaccPragmaIdslex.Error(err.Error())
				return 1
			}
			YaccPragmaIdsVAL.Version = NewVersion(intValue01, intValue03)
		}
	}
	goto YaccPragmaIdsstack /* stack new state and value */
}
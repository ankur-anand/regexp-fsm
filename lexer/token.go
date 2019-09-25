package lexer

const (
	eos int = iota
	any
	startOfLine
	endOfLine
	characterCLassStart
	characterClassEnd
	curlyOpen
	curlyClose
	parenOpen
	parenClose
	or
	oneOrMore
	zeroOrMore
	zeroOrOne
	dash
)

// Token for regexp machine
var Token = map[string]int{
	".":   any,
	"^":   startOfLine,
	"$":   endOfLine,
	"[":   characterCLassStart,
	"]":   characterClassEnd,
	"{":   curlyOpen,
	"}":   curlyClose,
	"(":   parenOpen,
	")":   parenClose,
	"*":   zeroOrMore,
	"+":   oneOrMore,
	"?":   zeroOrOne,
	"|":   or,
	"-":   dash,
	"EOF": eos,
}

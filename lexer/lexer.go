package lexer

var escapeVector = map[byte]byte{
	'b': '\b',
	'f': '\f',
	'n': '\n',
	's': ' ',
	't': '\t',
	'e': '\033',
}

// Lexer will take the source code as input
// and output the tokens that represent
// the source code.
type Lexer struct {
	input  []byte
	lexeme byte
	// current position in input(points to current char)
	position int
	// current token under examination
	currentToken int
	isEscape     bool
}

// New constructor to return new initialized Lexer
func New(input []byte) *Lexer {
	l := &Lexer{
		input: input,
	}
	return l
}

func (l *Lexer) Advance() int {
	if l.position > len(l.input)-1 {
		l.currentToken = Token["EOL"]
		return Token["EOL"]
	}

	l.lexeme = l.input[l.position]
	// check if the byte is an escape sequence
	if l.lexeme == '\\' {
		l.isEscape = !l.isEscape
		l.position++
		l.currentToken = l.handleEscape()
	}
	return Token["EOL"]
}

func (l *Lexer) handleEscape() int {
	if l.position > len(l.input)-1 {
		l.currentToken = Token["EOL"]
		return Token["EOL"]
	}
	switch l.input[l.position] {
	case '\\':
		l.Advance()
	case 'b', 'f', 'n', 'r', 't', 's':
	case :
		
	}
}

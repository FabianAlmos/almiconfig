package lexer

import "almi/consts"

const (
	COMMA      = 44
	LSQBRACKET = 91
)

type Lexer struct {
	Line  string
	Char  rune
	Index int

	Token  string
	Tokens []string
}

func NewLexer(line string) *Lexer {
	char := rune(0)
	if 0 < len(line) {
		char = rune(line[0])
	}

	return &Lexer{
		Line: line,
		Char: char,
	}
}

func (l *Lexer) HasNext() bool {
	return l.Index < len(l.Line)
}

func (l *Lexer) Next() rune {
	if l.HasNext() {
		l.Char = rune(l.Line[l.Index])
		l.Index++
		return l.Char
	}

	return 0
}

func (l *Lexer) Tokenize() []string {
	for l.HasNext() {
		l.Next()
		if l.Char == COMMA && l.Token[len(l.Token)-1] != LSQBRACKET {
			l.Next()
			l.Tokens = append(l.Tokens, l.Token)
			l.Token = consts.EMPTY
		}
		l.Token += string(l.Char)
	}

	l.Tokens = append(l.Tokens, l.Token)
	l.Token = consts.EMPTY

	return l.Tokens
}

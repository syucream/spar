package parser

import (
	"io"
)

type Lexer struct {
	scanner Scanner
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		scanner: NewScanner(r),
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit := l.scanner.Scan()

	if tok == IDENT {
	}
}

func (l *Lexer) Error(e string) {
	panic(e)
}

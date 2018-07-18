package parser

import (
	"io"
)

type Lexer struct {
	scanner *Scanner

	Stmt Statement
}

func NewLexer(r io.Reader) *Lexer {
	return &Lexer{
		scanner: NewScanner(r),
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
SCAN:
	tok, lit := l.scanner.Scan()

	switch tok {
	case IDENT:
		lval.str = lit
	default:
		goto SCAN
	}

	return len(lit)
}

func (l *Lexer) Error(e string) {
	panic(e)
}

package parser

import (
	"io"
	"log"
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
	case EOF:
		// Stop lex
		return 0
	case IDENT:
		lval.str = lit
	case WS:
		// Skip
		goto SCAN
	default:
		log.Fatalf("Unexpected token: token: %d, literal: %s\n", tok, lit)
	}
	log.Println(lit)

	// TODO easy to tokenize
	// Use prepared map?
	tokVal := 0
	switch lit {
	case "CREATE":
		tokVal = CREATE
	case "DATABASE":
		tokVal = DATABASE
	default:
		// TODO check last token is DATABASE!!!
		tokVal = database_id
	}

	return tokVal
}

func (l *Lexer) Error(e string) {
	panic(e)
}

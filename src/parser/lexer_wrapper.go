package parser

import (
	"fmt"
	"log"

	"github.com/syucream/spar/src/lexer"
	"github.com/syucream/spar/src/types"
)

type lexerWrapper struct {
	impl   *lexer.LexerImpl
	result types.DDStatements
	err    error
}

func newLexerWrapper(li *lexer.LexerImpl) *lexerWrapper {
	return &lexerWrapper{
		impl: li,
	}
}

func (l *lexerWrapper) Lex(lval *yySymType) int {
	r, err := l.impl.Lex(lval.LastToken)
	if err != nil {
		log.Fatal(err)
	}
	// To debug lexer
	// log.Print(r)

	tokVal := r.Token
	lval.str = r.Literal
	lval.LastToken = tokVal

	return tokVal
}

func (l *lexerWrapper) Error(e string) {
	l.err = fmt.Errorf(e)
}

package parser

import (
	"fmt"
	"log"

	"github.com/syucream/spar/src/lexer"
	"github.com/syucream/spar/src/types"
)

type LexerWrapper struct {
	impl   *lexer.LexerImpl
	Result types.DDStatements
	Err    error
}

func NewLexerWrapper(li *lexer.LexerImpl) *LexerWrapper {
	return &LexerWrapper{
		impl: li,
		Result: types.DDStatements{
			CreateDatabases: []types.CreateDatabaseStatement{},
			CreateTables:    []types.CreateTableStatement{},
		},
	}
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
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

func (l *LexerWrapper) Error(e string) {
	l.Err = fmt.Errorf(e)
}

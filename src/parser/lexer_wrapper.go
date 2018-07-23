package parser

import (
	"log"

	"github.com/syucream/jack/src/lexer"
)

type LexerWrapper struct {
	impl   *lexer.LexerImpl
	Result DDStatements
}

func NewLexerWrapper(li *lexer.LexerImpl) *LexerWrapper {
	return &LexerWrapper{
		impl: li,
		Result: DDStatements{
			CreateDatabases: []CreateDatabaseStatement{},
			CreateTables:    []CreateTableStatement{},
		},
	}
}

func (l *LexerWrapper) Lex(lval *yySymType) int {
	r, err := l.impl.Lex(lval.LastToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(r)

	tokVal := r.Token
	lval.str = r.Literal
	lval.LastToken = tokVal

	return tokVal
}

func (l *LexerWrapper) Error(e string) {
	panic(e)
}

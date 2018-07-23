package parser

import (
	"io"

	"github.com/syucream/jack/src/lexer"
)

func Parse(r io.Reader) *DDStatements {
	impl := lexer.NewLexerImpl(r, &KeywordTokenizer{})
	l := NewLexerWrapper(impl)

	yyParse(l)

	return l.Result
}

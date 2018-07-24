package parser

import (
	"io"

	"github.com/syucream/spar/src/lexer"
)

func Parse(r io.Reader) DDStatements {
	impl := lexer.NewLexerImpl(r, &KeywordTokenizer{})
	l := NewLexerWrapper(impl)

	yyParse(l)

	return l.Result
}

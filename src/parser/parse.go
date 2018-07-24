package parser

import (
	"io"

	"github.com/syucream/spar/src/lexer"
	"github.com/syucream/spar/src/types"
)

func Parse(r io.Reader) types.DDStatements {
	impl := lexer.NewLexerImpl(r, &KeywordTokenizer{})
	l := NewLexerWrapper(impl)

	yyParse(l)

	return l.Result
}

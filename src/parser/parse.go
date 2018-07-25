package parser

import (
	"io"

	"github.com/syucream/spar/src/lexer"
	"github.com/syucream/spar/src/types"
)

func Parse(r io.Reader) (*types.DDStatements, error) {
	impl := lexer.NewLexerImpl(r, &KeywordTokenizer{})
	l := NewLexerWrapper(impl)

	yyParse(l)

	if l.Err != nil {
		return nil, l.Err
	} else {
		return &l.Result, nil
	}
}

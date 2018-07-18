package parser

import (
	"io"
)

func ParseSpannerDDL(r io.Reader) Statement {
	l := NewLexer(r)
	yyParse(l)
	return l.Stmt
}

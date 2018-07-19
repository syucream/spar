package lexer

import (
	"fmt"
	"io"
)

var (
	UnexpectedTokenErr = fmt.Errorf("unexpected token")
)

type Tokenizer interface {
	FromStrLit(lit string, lastToken int) int
}

type LexerImpl struct {
	scanner   *Scanner
	tokenizer Tokenizer
	Result    interface{}
}

type LexerResult struct {
	Token   int
	Literal string
}

func NewLexerImpl(r io.Reader, t Tokenizer) *LexerImpl {
	return &LexerImpl{
		scanner:   NewScanner(r),
		tokenizer: t,
	}
}

func (li *LexerImpl) Lex(lastToken int) (*LexerResult, error) {
	result := &LexerResult{}

SCAN:
	tok, lit := li.scanner.Scan()

	switch tok {
	case EOF:
		// Stop lex
	case IDENT, LEFT_PARENTHESIS, RIGHT_PARENTHESIS:
		result.Literal = lit
	case WS:
		// Skip
		goto SCAN
	default:
		return nil, UnexpectedTokenErr
	}

	result.Token = li.tokenizer.FromStrLit(lit, lastToken)

	return result, nil
}

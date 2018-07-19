package parser

import (
	"io"
	"log"
)

var keywords = map[string]int{
	"CREATE":     CREATE,
	"ALTER":      ALTER,
	"DROP":       DROP,
	"DATABASE":   DATABASE,
	"TABLE":      TABLE,
	"INDEX":      INDEX,
	"PRIMARY":    PRIMARY,
	"KEY":        KEY,
	"ASC":        ASC,
	"DESC":       DESC,
	"INTERLEAVE": INTERLEAVE,
	"IN":         IN,
	"PARENT":     PARENT,
	"ARRAY":      ARRAY,
	"OPTIONS":    OPTIONS,
	"NOT":        NOT,
	"NULL":       NULL,
	"ON":         ON,
	"DELETE":     DELETE,
	"CASCADE":    CASCADE,
	"NO":         NO,
	"ACTION":     ACTION,
	"MAX":        MAX,
	"true":       true,
	"null":       null,
	"allow_commit_timestamp": allow_commit_timestamp,
	"BOOL":          BOOL,
	"INT64":         INT64,
	"FLOAT64":       FLOAT64,
	"STRING":        STRING,
	"BYTES":         BYTES,
	"DATE":          DATE,
	"TIMESTAMP":     TIMESTAMP,
	"database_id":   database_id,
	"decimal_value": decimal_value,
	"hex_value":     hex_value,
	"table_name":    table_name,
	"column_name":   column_name,
	"index_name":    index_name,
}

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
		// case LEFT_PARENTHESIS:
		// case RIGHT_PARENTHESIS:
		lval.str = lit
	case WS:
		// Skip
		goto SCAN
	default:
		log.Fatalf("Unexpected token: token: %d, literal: %s\n", tok, lit)
	}
	log.Println(lit)

	tokVal := 0
	if v, ok := keywords[lit]; ok {
		tokVal = v
	} else {
		switch lval.LastToken {
		case DATABASE:
			tokVal = database_id
		case TABLE:
			tokVal = table_name
		}
	}
	log.Println(tokVal)

	lval.LastToken = tokVal
	return tokVal
}

func (l *Lexer) Error(e string) {
	panic(e)
}

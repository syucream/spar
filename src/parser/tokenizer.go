package parser

import (
	"strconv"
)

const (
	LEFT_PARENTHESIS_TOKEN  = int('(')
	RIGHT_PARENTHESIS_TOKEN = int(')')
	COMMA_TOKEN             = int(',')
	SEMICOLON_TOKEN         = int(';')
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

var symbols = map[string]int{
	"(": LEFT_PARENTHESIS_TOKEN,
	")": RIGHT_PARENTHESIS_TOKEN,
	",": COMMA_TOKEN,
	";": SEMICOLON_TOKEN,
}

type KeywordTokenizer struct{}

func (kt *KeywordTokenizer) FromStrLit(lit string, lastToken int) int {
	tokVal := 0

	if v, ok := keywords[lit]; ok {
		tokVal = v
	} else if v, ok := symbols[lit]; ok {
		tokVal = v
	} else if _, err := strconv.ParseInt(lit, 10, 0); err == nil {
		tokVal = decimal_value
	} else {
		switch lastToken {
		case DATABASE:
			tokVal = database_id
		case TABLE:
			tokVal = table_name
		case LEFT_PARENTHESIS_TOKEN, COMMA_TOKEN:
			tokVal = column_name
		}
	}

	return tokVal
}

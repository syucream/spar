package parser

import (
	"regexp"
	"strconv"
)

const (
	LEFT_PARENTHESIS_TOKEN  = int('(')
	RIGHT_PARENTHESIS_TOKEN = int(')')
	COMMA_TOKEN             = int(',')
	SEMICOLON_TOKEN         = int(';')
	EQUAL_TOKEN             = int('=')
	ANGLE_LEFT_TOKEN        = int('<')
	ANGLE_RIGHT_TOKEN       = int('>')
)

var keywords = map[string]int{
	"CREATE":                 CREATE,
	"ALTER":                  ALTER,
	"DROP":                   DROP,
	"DATABASE":               DATABASE,
	"TABLE":                  TABLE,
	"INDEX":                  INDEX,
	"PRIMARY":                PRIMARY,
	"KEY":                    KEY,
	"ASC":                    ASC,
	"DESC":                   DESC,
	"INTERLEAVE":             INTERLEAVE,
	"IN":                     IN,
	"PARENT":                 PARENT,
	"ARRAY":                  ARRAY,
	"OPTIONS":                OPTIONS,
	"NOT":                    NOT,
	"NULL":                   NULL,
	"ON":                     ON,
	"DELETE":                 DELETE,
	"CASCADE":                CASCADE,
	"NO":                     NO,
	"ACTION":                 ACTION,
	"MAX":                    MAX,
	"UNIQUE":                 UNIQUE,
	"NULL_FILTERED":          NULL_FILTERED,
	"STORING":                STORING,
	"ADD":                    ADD,
	"COLUMN":                 COLUMN,
	"SET":                    SET,
	"true":                   true,
	"null":                   null,
	"allow_commit_timestamp": allow_commit_timestamp,
	"BOOL":                   BOOL,
	"INT64":                  INT64,
	"FLOAT64":                FLOAT64,
	"STRING":                 STRING,
	"BYTES":                  BYTES,
	"DATE":                   DATE,
	"TIMESTAMP":              TIMESTAMP,
	"database_id":            database_id,
	"decimal_value":          decimal_value,
	"hex_value":              hex_value,
	"table_name":             table_name,
	"column_name":            column_name,
	"index_name":             index_name,
}

var symbols = map[string]int{
	"(": LEFT_PARENTHESIS_TOKEN,
	")": RIGHT_PARENTHESIS_TOKEN,
	",": COMMA_TOKEN,
	";": SEMICOLON_TOKEN,
	"=": EQUAL_TOKEN,
	"<": ANGLE_LEFT_TOKEN,
	">": ANGLE_RIGHT_TOKEN,
}

var (
	databaseIdRegexp = regexp.MustCompile(`[a-z][a-z0-9_\-]*[a-z0-9]`)
	nameAttrRegexp   = regexp.MustCompile(`[a-zA-Z][a-zA-Z0-9_]+`)
)

type keywordTokenizer struct{}

// FromStrLit tokenize lit to a token pre-defined by goyacc with last token as a hint.
// TODO Check some literals satisfy regexp specs.
func (kt *keywordTokenizer) FromStrLit(lit string, lastToken int) int {
	tokVal := 0

	if v, ok := keywords[lit]; ok {
		tokVal = v
	} else if v, ok := symbols[lit]; ok {
		tokVal = v
	} else if _, err := strconv.ParseInt(lit, 10, 0); err == nil {
		tokVal = decimal_value
	} else if len(lit) >= 3 && lit[:2] == "0x" {
		if _, err := strconv.ParseInt(lit[2:], 16, 0); err == nil {
			tokVal = hex_value
		}
	} else {
		switch lastToken {
		case DATABASE:
			if databaseIdRegexp.MatchString(lit) {
				tokVal = database_id
			}
		case TABLE:
			if nameAttrRegexp.MatchString(lit) {
				tokVal = table_name
			}
		case INDEX:
			if nameAttrRegexp.MatchString(lit) {
				tokVal = index_name
			}
		case ON: // TODO duplicated token! Check pre-parsed tokens more!
			if nameAttrRegexp.MatchString(lit) {
				tokVal = table_name
			}
		case PARENT:
			if nameAttrRegexp.MatchString(lit) {
				tokVal = table_name
			}
		case COLUMN, LEFT_PARENTHESIS_TOKEN, COMMA_TOKEN:
			if nameAttrRegexp.MatchString(lit) {
				tokVal = column_name
			}
		}
	}

	return tokVal
}

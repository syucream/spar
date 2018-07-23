package spanner2mysql

import (
	"fmt"
	"strings"

	"github.com/syucream/jack/src/parser"
)

var toMysqlType = map[string]string{
	"BOOL":      "TINYINT(1)",
	"INT64":     "BIGINT",
	"FLOAT64":   "DOUBLE",
	"STRING":    "VARCHAR",
	"BYTES":     "BLOB",
	"DATE":      "DATE",
	"TIMESTAMP": "TIMESTAMP",
}

func getMysqlType(t string) (string, error) {
	origType := t

	index := strings.Index(t, "(")
	if index != -1 {
		// Either STRING or BYTES
		origType = t[:index]
	}

	convertedType := ""
	if v, ok := toMysqlType[origType]; ok {
		convertedType = v

		// Append length attribute for VARCHAR
		if convertedType == "VARCHAR" {
			convertedType += t[index:]
		}
	} else {
		return "", fmt.Errorf("Invalid type: %s\n", origType)
	}

	return convertedType, nil
}

func GetMysqlCreateTables(statements parser.DDStatements) (string, error) {
	converted := ""

	for _, ct := range statements.CreateTables {
		converted += fmt.Sprintf("CREATE TABLE %s (\n", ct.Statement.Id)

		var defs []string
		for _, col := range ct.Columns {
			convertedType, err := getMysqlType(col.Type)
			if err != nil {
				return "", err
			}

			attr := "NULL"
			if col.NotNull {
				attr = "NOT NULL"
			}

			defs = append(defs, fmt.Sprintf("  %s %s %s", col.Name, convertedType, attr))
		}
		for _, pk := range ct.PrimaryKeys {
			defs = append(defs, fmt.Sprintf("  PRIMARY KEY(%s)", pk.Name))
		}

		converted += strings.Join(defs, ",\n") + "\n);\n"
	}

	return converted, nil
}

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

		// Replace too big VARCHAR to TEXT or append length attribute for VARCHAR
		// TODO more strict check
		if convertedType == "VARCHAR" {
			if t[index:] == "(MAX)" {
				convertedType = "TEXT"
			} else {
				convertedType += t[index:]
			}
		}
	} else {
		return "", fmt.Errorf("Invalid type: %s\n", origType)
	}

	return convertedType, nil
}

func getRelation(child parser.CreateTableStatement, maybeParents []parser.CreateTableStatement) (string, error) {
	var parent *parser.CreateTableStatement
	for _, s := range maybeParents {
		if child.Cluster.TableName == s.Id {
			parent = &s
			break
		}
	}

	if parent == nil {
		return "", fmt.Errorf("Invalid interleave: %v\n", child.Cluster)
	}

	var keyCol *parser.Column
	for _, cc := range child.Columns {
		for _, pc := range parent.Columns {
			if cc.Name == pc.Name && cc.Type == pc.Type {
				keyCol = &cc
				break
			}
		}
	}

	if keyCol == nil {
		return "", fmt.Errorf("Invalid interleave: %v\n", child.Cluster)
	}

	return fmt.Sprintf("  FOREIGN KEY (%s) REFERENCES %s(%s)", keyCol.Name, parent.Id, keyCol.Name), nil
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

			defs = append(defs, fmt.Sprintf("  %s %s %s", col.Name, convertedType, col.Nullability))
		}

		keyNames := make([]string, 0, len(ct.PrimaryKeys))
		for _, pk := range ct.PrimaryKeys {
			keyNames = append(keyNames, pk.Name)
		}
		defs = append(defs, fmt.Sprintf("  PRIMARY KEY(%s)", strings.Join(keyNames, ", ")))

		if ct.Cluster.TableName != "" {
			// Convert interleave to foreign key
			relation, err := getRelation(ct, statements.CreateTables)
			if err != nil {
				return "", err
			}
			defs = append(defs, relation)
		}

		converted += strings.Join(defs, ",\n") + "\n);\n"
	}

	return converted, nil
}

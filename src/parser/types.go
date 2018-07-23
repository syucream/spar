package parser

type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
}

type Statement struct {
	Action string // CREATE, ALTER, DROP
	Target string // DATABASE, TABLE, INDEX
	Id     string
}

type Column struct {
	Name    string
	Type    string // BOOL, INT64, ...
	NotNull bool
}

type CreateDatabaseStatement struct {
	Statement
}

func SetCreateDatabaseStatement(yylex interface{}, action string, target string, id string) {
	s := CreateDatabaseStatement{
		Statement: Statement{
			Action: action,
			Target: target,
			Id:     id,
		},
	}
	yylex.(*LexerWrapper).Result.CreateDatabases = append(yylex.(*LexerWrapper).Result.CreateDatabases, s)
}

type CreateTableStatement struct {
	Statement

	Columns     []Column
	PrimaryKeys []string
}

func SetCreateTableStatement(yylex interface{}, action string, target string, id string, cols []Column, pks []string) {
	s := CreateTableStatement{
		Statement: Statement{
			Action: action,
			Target: target,
			Id:     id,
		},
		Columns:     cols,
		PrimaryKeys: pks,
	}
	yylex.(*LexerWrapper).Result.CreateTables = append(yylex.(*LexerWrapper).Result.CreateTables, s)
}

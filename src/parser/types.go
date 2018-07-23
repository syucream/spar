package parser

type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
	CreateIndexes   []CreateIndexStatement
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

type Key struct {
	Name  string
	Order string
}

type CreateDatabaseStatement struct {
	Statement
}

type CreateTableStatement struct {
	Statement

	Columns     []Column
	PrimaryKeys []Key
}

type CreateIndexStatement struct {
	Statement

	Unique       bool
	NullFiltered bool
	TableName    string
	Keys         []Key
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

func SetCreateTableStatement(yylex interface{}, action string, target string, id string, cols []Column, keys []Key) {
	s := CreateTableStatement{
		Statement: Statement{
			Action: action,
			Target: target,
			Id:     id,
		},
		Columns:     cols,
		PrimaryKeys: keys,
	}
	yylex.(*LexerWrapper).Result.CreateTables = append(yylex.(*LexerWrapper).Result.CreateTables, s)
}

func SetCreateIndexStatement(yylex interface{}, action string, target string, id string, unique bool, nullFiltered bool, tableName string, keys []Key) {
	s := CreateIndexStatement{
		Statement: Statement{
			Action: action,
			Target: target,
			Id:     id,
		},
		TableName: tableName,
		Keys:      keys,
	}
	yylex.(*LexerWrapper).Result.CreateIndexes = append(yylex.(*LexerWrapper).Result.CreateIndexes, s)
}

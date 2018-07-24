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
	Name        string
	Type        string // BOOL, INT64, ...
	Nullability string
	Options     string
}

type Key struct {
	Name  string
	Order string
}

type Cluster struct {
	TableName string
	OnDelete  string
}

type CreateDatabaseStatement struct {
	Statement
}

type CreateTableStatement struct {
	Statement

	Columns     []Column
	PrimaryKeys []Key
	Cluster     Cluster
}

type CreateIndexStatement struct {
	Statement

	Unique       string
	NullFiltered string
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

func SetCreateTableStatement(yylex interface{}, action string, target string, id string, cols []Column, keys []Key, cluster Cluster) {
	s := CreateTableStatement{
		Statement: Statement{
			Action: action,
			Target: target,
			Id:     id,
		},
		Columns:     cols,
		PrimaryKeys: keys,
		Cluster:     cluster,
	}
	yylex.(*LexerWrapper).Result.CreateTables = append(yylex.(*LexerWrapper).Result.CreateTables, s)
}

func SetCreateIndexStatement(yylex interface{}, action string, target string, id string, unique string, nullFiltered string, tableName string, keys []Key) {
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

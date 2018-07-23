package parser

type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
}

type DDStatement struct {
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
	DDStatement
}

func SetCreateDatabaseStatement(yylex interface{}, action string, target string, id string) {
	s := &CreateDatabaseStatement{
		Action: action,
		Target: target,
		Id:     id,
	}
	yylex.(*LexerWrapper).Result.CreateDatabases = append(yylex.(*LexerWrapper).Result.CreateDatabases, s)
}

type CreateTableStatement struct {
	DDStatement

	Columns     []Column
	PrimaryKeys []string
}

func SetCreateTableStatement(yylex interface{}, action string, target string, id string, cols []Columns, pks []string) {
	s := &CreateTableStatement{
		Action:      action,
		Target:      target,
		Id:          id,
		Columns:     cols,
		PrimaryKeys: pks,
	}
	yylex.(*LexerWrapper).Result.CreateTables = append(yylex.(*LexerWrapper).Result.CreateTables, s)
}

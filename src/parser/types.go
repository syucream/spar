package parser

// TODO DDStatements

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
	yylex.(*LexerWrapper).Result = &CreateDatabaseStatement{
		Action: action,
		Target: target,
		Id:     id,
	}
}

type CreateTableStatement struct {
	DDStatement

	Columns     []Column
	PrimaryKeys []string
}

func SetCreateTableStatement(yylex interface{}, action string, target string, id string, cols []Columns, pks []string) {
	yylex.(*LexerWrapper).Result = &CreateTableStatement{
		Action:      action,
		Target:      target,
		Id:          id,
		Columns:     cols,
		PrimaryKeys: pks,
	}
}

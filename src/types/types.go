package types

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

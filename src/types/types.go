package types

type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
	CreateIndexes   []CreateIndexStatement
	DropTables      []DropTableStatement
	DropIndexes     []DropIndexStatement
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
	DatabaseId string
}

type CreateTableStatement struct {
	TableName   string
	Columns     []Column
	PrimaryKeys []Key
	Cluster     Cluster
}

type CreateIndexStatement struct {
	IndexName    string
	Unique       string
	NullFiltered string
	TableName    string
	Keys         []Key
}

type DropTableStatement struct {
	TableName string
}

type DropIndexStatement struct {
	IndexName string
}

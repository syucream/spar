package types

// DDStatements has parsed statements.
type DDStatements struct {
	CreateDatabases []CreateDatabaseStatement
	CreateTables    []CreateTableStatement
	CreateIndexes   []CreateIndexStatement
	AlterTables     []AlterTableStatement
	DropTables      []DropTableStatement
	DropIndexes     []DropIndexStatement
}

// Column is a table column.
type Column struct {
	Name        string
	Type        string // BOOL, INT64, ...
	Nullability string
	Options     string
}

// Key is a table key.
type Key struct {
	Name  string
	Order string
}

// Cluster is a Spanner table cluster.
type Cluster struct {
	TableName string
	OnDelete  string
}

// StoringClause is a storing clause info.
type StoringClause struct {
	ColumnNames []string
}

// Interleave is a interlive.
type Interleave struct {
	TableName string
}

// Alteration is a alter table operation.
type Alteration interface {
	Alter()
}

// AddColumnTableAlteration is a alter table operation, corresponding to 'ADD COLUMN'.
type AddColumnTableAlteration struct {
	Column Column
}

// Alter satisfies Alteration interface
func (self *AddColumnTableAlteration) Alter() {}

// DropColumnTableAlteration is a alter table operation, corresponding to 'DROP COLUMN'.
type DropColumnTableAlteration struct {
	ColumnName string
}

// Alter satisfies Alteration interface
func (self *DropColumnTableAlteration) Alter() {}

// SetTableAlteration is a alter table operation, corresponding to 'SET'.
type SetTableAlteration struct {
	OnDelete string
}

// Alter satisfies Alteration interface
func (self *SetTableAlteration) Alter() {}

// AlterColumnTypesAlteration is a alter table operation, corresponding to 'ALTER COLUMN'.
type AlterColumnTypesAlteration struct {
	ColumnName  string
	ColumnType  string
	Nullability string
}

// Alter satisfies Alteration interface
func (self *AlterColumnTypesAlteration) Alter() {}

// AlterColumnSetAlteration is a alter table operation, corresponding to 'ALTER COLUMN'.
type AlterColumnSetAlteration struct {
	ColumnName string
	Options    string
}

// Alter satisfies Alteration interface
func (self *AlterColumnSetAlteration) Alter() {}

// CreateDatabaseStatement is a 'CREATE DATABASE' statement info.
type CreateDatabaseStatement struct {
	DatabaseId string
}

// CreateTableStatement is a 'CREATE TABLE' statement info.
type CreateTableStatement struct {
	TableName   string
	Columns     []Column
	PrimaryKeys []Key
	Cluster     Cluster
}

// CreateIndexStatement is a 'CREATE INDEX' statement info.
type CreateIndexStatement struct {
	IndexName     string
	Unique        string
	NullFiltered  string
	TableName     string
	Keys          []Key
	StoringClause StoringClause
	Interleaves   []Interleave
}

// AlterTableStatement is a 'ALTER TABLE' statement info.
type AlterTableStatement struct {
	TableName  string
	Alteration Alteration // Get concrete type yourself
}

// DropTableStatement is a 'DROP TABLE' statement info.
type DropTableStatement struct {
	TableName string
}

// DropIndexStatement is a 'DROP INDEX' statement info.
type DropIndexStatement struct {
	IndexName string
}
